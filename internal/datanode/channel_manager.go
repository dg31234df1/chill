// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datanode

import (
	"context"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/lifetime"
	"github.com/milvus-io/milvus/pkg/util/merr"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

type releaseFunc func(channel string)

type ChannelManager interface {
	Submit(info *datapb.ChannelWatchInfo) error
	GetProgress(info *datapb.ChannelWatchInfo) *datapb.ChannelOperationProgressResponse
	Close()
	Start()
}

type ChannelManagerImpl struct {
	mu sync.RWMutex
	dn *DataNode

	fgManager FlowgraphManager

	communicateCh chan *opState
	opRunners     *typeutil.ConcurrentMap[string, *opRunner] // channel -> runner
	abnormals     *typeutil.ConcurrentMap[int64, string]     // OpID -> Channel

	releaseFunc releaseFunc

	closeCh     lifetime.SafeChan
	closeWaiter sync.WaitGroup
}

func NewChannelManager(dn *DataNode) *ChannelManagerImpl {
	cm := ChannelManagerImpl{
		dn:        dn,
		fgManager: dn.flowgraphManager,

		communicateCh: make(chan *opState, 100),
		opRunners:     typeutil.NewConcurrentMap[string, *opRunner](),
		abnormals:     typeutil.NewConcurrentMap[int64, string](),

		releaseFunc: dn.flowgraphManager.RemoveFlowgraph,

		closeCh: lifetime.NewSafeChan(),
	}

	return &cm
}

func (m *ChannelManagerImpl) Submit(info *datapb.ChannelWatchInfo) error {
	channel := info.GetVchan().GetChannelName()
	runner := m.getOrCreateRunner(channel)
	return runner.Enqueue(info)
}

func (m *ChannelManagerImpl) GetProgress(info *datapb.ChannelWatchInfo) *datapb.ChannelOperationProgressResponse {
	m.mu.RLock()
	defer m.mu.RUnlock()
	resp := &datapb.ChannelOperationProgressResponse{
		Status: merr.Success(),
		OpID:   info.GetOpID(),
	}

	channel := info.GetVchan().GetChannelName()
	switch info.GetState() {
	case datapb.ChannelWatchState_ToWatch:
		// running flowgraph means watch success
		if m.fgManager.HasFlowgraphWithOpID(channel, info.GetOpID()) {
			resp.State = datapb.ChannelWatchState_WatchSuccess
			resp.Progress = 100
			return resp
		}

		if runner, ok := m.opRunners.Get(channel); ok {
			if runner.Exist(info.GetOpID()) {
				resp.State = datapb.ChannelWatchState_ToWatch
			} else {
				resp.State = datapb.ChannelWatchState_WatchFailure
			}
			return resp
		}
		resp.State = datapb.ChannelWatchState_WatchFailure
		return resp

	case datapb.ChannelWatchState_ToRelease:
		if !m.fgManager.HasFlowgraph(channel) {
			resp.State = datapb.ChannelWatchState_ReleaseSuccess
			return resp
		}
		if runner, ok := m.opRunners.Get(channel); ok && runner.Exist(info.GetOpID()) {
			resp.State = datapb.ChannelWatchState_ToRelease
			return resp
		}

		resp.State = datapb.ChannelWatchState_ReleaseFailure
		return resp
	default:
		err := merr.WrapErrParameterInvalid("ToWatch or ToRelease", info.GetState().String())
		log.Warn("fail to get progress", zap.Error(err))
		resp.Status = merr.Status(err)
		return resp
	}
}

func (m *ChannelManagerImpl) Close() {
	if m.opRunners != nil {
		m.opRunners.Range(func(channel string, runner *opRunner) bool {
			runner.Close()
			return true
		})
	}
	m.closeCh.Close()
	m.closeWaiter.Wait()
}

func (m *ChannelManagerImpl) Start() {
	m.closeWaiter.Add(1)
	go func() {
		defer m.closeWaiter.Done()
		log.Info("DataNode ChannelManager start")
		for {
			select {
			case opState := <-m.communicateCh:
				m.handleOpState(opState)
			case <-m.closeCh.CloseCh():
				log.Info("DataNode ChannelManager exit")
				return
			}
		}
	}()
}

func (m *ChannelManagerImpl) handleOpState(opState *opState) {
	m.mu.Lock()
	defer m.mu.Unlock()
	log := log.With(
		zap.Int64("opID", opState.opID),
		zap.String("channel", opState.channel),
		zap.String("State", opState.state.String()),
	)
	switch opState.state {
	case datapb.ChannelWatchState_WatchSuccess:
		log.Info("Success to watch")
		m.fgManager.AddFlowgraph(opState.fg)

	case datapb.ChannelWatchState_WatchFailure:
		log.Info("Fail to watch")

	case datapb.ChannelWatchState_ReleaseSuccess:
		log.Info("Success to release")

	case datapb.ChannelWatchState_ReleaseFailure:
		log.Info("Fail to release, add channel to abnormal lists")
		m.abnormals.Insert(opState.opID, opState.channel)
	}

	m.finishOp(opState.opID, opState.channel)
}

func (m *ChannelManagerImpl) getOrCreateRunner(channel string) *opRunner {
	runner, loaded := m.opRunners.GetOrInsert(channel, NewOpRunner(channel, m.dn, m.releaseFunc, m.communicateCh))
	if !loaded {
		runner.Start()
	}
	return runner
}

func (m *ChannelManagerImpl) finishOp(opID int64, channel string) {
	if runner, loaded := m.opRunners.GetAndRemove(channel); loaded {
		runner.FinishOp(opID)
		runner.Close()
	}
}

type opInfo struct {
	tickler *tickler
}

type opRunner struct {
	channel     string
	dn          *DataNode
	releaseFunc releaseFunc

	guard      sync.RWMutex
	allOps     map[UniqueID]*opInfo // opID -> tickler
	opsInQueue chan *datapb.ChannelWatchInfo
	resultCh   chan *opState

	closeCh lifetime.SafeChan
	closeWg sync.WaitGroup
}

func NewOpRunner(channel string, dn *DataNode, f releaseFunc, resultCh chan *opState) *opRunner {
	return &opRunner{
		channel:     channel,
		dn:          dn,
		releaseFunc: f,
		opsInQueue:  make(chan *datapb.ChannelWatchInfo, 10),
		allOps:      make(map[UniqueID]*opInfo),
		resultCh:    resultCh,
		closeCh:     lifetime.NewSafeChan(),
	}
}

func (r *opRunner) Start() {
	r.closeWg.Add(1)
	go func() {
		defer r.closeWg.Done()
		for {
			select {
			case info := <-r.opsInQueue:
				r.NotifyState(r.Execute(info))
			case <-r.closeCh.CloseCh():
				return
			}
		}
	}()
}

func (r *opRunner) FinishOp(opID UniqueID) {
	r.guard.Lock()
	defer r.guard.Unlock()
	delete(r.allOps, opID)
}

func (r *opRunner) Exist(opID UniqueID) bool {
	r.guard.RLock()
	defer r.guard.RUnlock()
	_, ok := r.allOps[opID]
	return ok
}

func (r *opRunner) Enqueue(info *datapb.ChannelWatchInfo) error {
	if info.GetState() != datapb.ChannelWatchState_ToWatch &&
		info.GetState() != datapb.ChannelWatchState_ToRelease {
		return errors.New("Invalid channel watch state")
	}

	r.guard.Lock()
	defer r.guard.Unlock()
	if _, ok := r.allOps[info.GetOpID()]; !ok {
		r.opsInQueue <- info
		r.allOps[info.GetOpID()] = &opInfo{}
	}
	return nil
}

func (r *opRunner) UnfinishedOpSize() int {
	r.guard.RLock()
	defer r.guard.RUnlock()
	return len(r.allOps)
}

// Execute excutes channel operations, channel state is validated during enqueue
func (r *opRunner) Execute(info *datapb.ChannelWatchInfo) *opState {
	log.Info("Start to execute channel operation",
		zap.String("channel", info.GetVchan().GetChannelName()),
		zap.Int64("opID", info.GetOpID()),
		zap.String("state", info.GetState().String()),
	)
	if info.GetState() == datapb.ChannelWatchState_ToWatch {
		return r.watchWithTimer(info)
	}

	// ToRelease state
	return r.releaseWithTimer(r.releaseFunc, info.GetVchan().GetChannelName(), info.GetOpID())
}

// watchWithTimer will return WatchFailure after WatchTimeoutInterval
func (r *opRunner) watchWithTimer(info *datapb.ChannelWatchInfo) *opState {
	opState := &opState{
		channel: info.GetVchan().GetChannelName(),
		opID:    info.GetOpID(),
	}
	log := log.With(zap.String("channel", opState.channel), zap.Int64("opID", opState.opID))

	r.guard.Lock()
	opInfo, ok := r.allOps[info.GetOpID()]
	r.guard.Unlock()
	if !ok {
		opState.state = datapb.ChannelWatchState_WatchFailure
		return opState
	}
	tickler := newTickler()
	opInfo.tickler = tickler

	var (
		successSig = make(chan struct{}, 1)
		waiter     sync.WaitGroup
	)

	watchTimeout := Params.DataCoordCfg.WatchTimeoutInterval.GetAsDuration(time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), watchTimeout)
	defer cancel()

	startTimer := func(wg *sync.WaitGroup) {
		defer wg.Done()

		timer := time.NewTimer(watchTimeout)
		defer timer.Stop()

		log := log.With(zap.Duration("timeout", watchTimeout))
		log.Info("Start timer for ToWatch operation")
		for {
			select {
			case <-timer.C:
				// watch timeout
				tickler.close()
				cancel()
				log.Info("Stop timer for ToWatch operation timeout")
				return

			case <-r.closeCh.CloseCh():
				// runner closed from outside
				tickler.close()
				cancel()
				log.Info("Suspend ToWatch operation from outside of opRunner")
				return

			case <-tickler.progressSig:
				log.Info("Reset timer for tickler updated")
				timer.Reset(watchTimeout)

			case <-successSig:
				// watch success
				log.Info("Stop timer for ToWatch operation succeeded")
				return
			}
		}
	}

	waiter.Add(2)
	go startTimer(&waiter)
	go func() {
		defer waiter.Done()
		fg, err := executeWatch(ctx, r.dn, info, tickler)
		if err != nil {
			opState.state = datapb.ChannelWatchState_WatchFailure
		} else {
			opState.state = datapb.ChannelWatchState_WatchSuccess
			opState.fg = fg
			successSig <- struct{}{}
		}
	}()

	waiter.Wait()
	return opState
}

// releaseWithTimer will return ReleaseFailure after WatchTimeoutInterval
func (r *opRunner) releaseWithTimer(releaseFunc releaseFunc, channel string, opID UniqueID) *opState {
	opState := &opState{
		channel: channel,
		opID:    opID,
	}
	var (
		successSig = make(chan struct{}, 1)
		waiter     sync.WaitGroup
	)

	log := log.With(zap.Int64("opID", opID), zap.String("channel", channel))
	startTimer := func(wg *sync.WaitGroup) {
		defer wg.Done()
		releaseTimeout := Params.DataCoordCfg.WatchTimeoutInterval.GetAsDuration(time.Second)
		timer := time.NewTimer(releaseTimeout)
		defer timer.Stop()

		log := log.With(zap.Duration("timeout", releaseTimeout))
		log.Info("Start ToRelease timer")
		for {
			select {
			case <-timer.C:
				log.Info("Stop timer for ToRelease operation timeout")
				opState.state = datapb.ChannelWatchState_ReleaseFailure
				return

			case <-r.closeCh.CloseCh():
				// runner closed from outside
				log.Info("Stop timer for opRunner closed")
				return

			case <-successSig:
				log.Info("Stop timer for ToRelease operation succeeded")
				opState.state = datapb.ChannelWatchState_ReleaseSuccess
				return
			}
		}
	}

	waiter.Add(1)
	go startTimer(&waiter)
	go func() {
		// TODO: failure should panic this DN, but we're not sure how
		//   to recover when releaseFunc stuck.
		// Whenever we see a stuck, it's a bug need to be fixed.
		// In case of the unknown behavior after the stuck of release,
		//   we'll mark this channel abnormal in this DN. This goroutine might never return.
		//
		// The channel can still be balanced into other DNs, but not on this one.
		// ExclusiveConsumer error happens when the same DN subscribes the same pchannel twice.
		releaseFunc(opState.channel)
		successSig <- struct{}{}
	}()

	waiter.Wait()
	return opState
}

func (r *opRunner) NotifyState(state *opState) {
	r.resultCh <- state
}

func (r *opRunner) Close() {
	r.closeCh.Close()
	r.closeWg.Wait()
}

type opState struct {
	channel string
	opID    int64
	state   datapb.ChannelWatchState
	fg      *dataSyncService
}

// executeWatch will always return, won't be stuck, either success or fail.
func executeWatch(ctx context.Context, dn *DataNode, info *datapb.ChannelWatchInfo, tickler *tickler) (*dataSyncService, error) {
	dataSyncService, err := newDataSyncService(ctx, dn, info, tickler)
	if err != nil {
		return nil, err
	}

	dataSyncService.start()

	return dataSyncService, nil
}

// tickler counts every time when called inc(),
type tickler struct {
	count     *atomic.Int32
	total     *atomic.Int32
	closedSig *atomic.Bool

	progressSig chan struct{}
}

func (t *tickler) inc() {
	t.count.Inc()
	t.progressSig <- struct{}{}
}

func (t *tickler) setTotal(total int32) {
	t.total.Store(total)
}

// progress returns the count over total if total is set
// else just return the count number.
func (t *tickler) progress() int32 {
	if t.total.Load() == 0 {
		return t.count.Load()
	}
	return (t.count.Load() / t.total.Load()) * 100
}

func (t *tickler) close() {
	t.closedSig.CompareAndSwap(false, true)
}

func (t *tickler) closed() bool {
	return t.closedSig.Load()
}

func newTickler() *tickler {
	return &tickler{
		count:       atomic.NewInt32(0),
		total:       atomic.NewInt32(0),
		closedSig:   atomic.NewBool(false),
		progressSig: make(chan struct{}, 200),
	}
}
