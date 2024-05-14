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

package datacoord

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/kv"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/conc"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

type ChannelManager interface {
	Startup(ctx context.Context, legacyNodes, allNodes []int64) error
	Close()

	AddNode(nodeID UniqueID) error
	DeleteNode(nodeID UniqueID) error
	Watch(ctx context.Context, ch RWChannel) error
	Release(nodeID UniqueID, channelName string) error

	Match(nodeID UniqueID, channel string) bool
	FindWatcher(channel string) (UniqueID, error)

	GetChannel(nodeID int64, channel string) (RWChannel, bool)
	GetNodeIDByChannelName(channel string) (int64, bool)
	GetNodeChannelsByCollectionID(collectionID int64) map[int64][]string
	GetChannelsByCollectionID(collectionID int64) []RWChannel
	GetChannelNamesByCollectionID(collectionID int64) []string
}

// An interface sessionManager implments
type SubCluster interface {
	NotifyChannelOperation(ctx context.Context, nodeID int64, req *datapb.ChannelOperationsRequest) error
	CheckChannelOperationProgress(ctx context.Context, nodeID int64, info *datapb.ChannelWatchInfo) (*datapb.ChannelOperationProgressResponse, error)
}

type ChannelManagerImplV2 struct {
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.RWMutex

	h          Handler
	store      RWChannelStore
	subCluster SubCluster // sessionManager
	allocator  allocator

	factory       ChannelPolicyFactory
	balancePolicy BalanceChannelPolicy

	balanceCheckLoop ChannelBGChecker

	legacyNodes typeutil.UniqueSet

	lastActiveTimestamp time.Time
}

// ChannelBGChecker are goroutining running background
type ChannelBGChecker func(ctx context.Context)

// ChannelmanagerOptV2 is to set optional parameters in channel manager.
type ChannelmanagerOptV2 func(c *ChannelManagerImplV2)

func withFactoryV2(f ChannelPolicyFactory) ChannelmanagerOptV2 {
	return func(c *ChannelManagerImplV2) { c.factory = f }
}

func withCheckerV2() ChannelmanagerOptV2 {
	return func(c *ChannelManagerImplV2) { c.balanceCheckLoop = c.CheckLoop }
}

func NewChannelManagerV2(
	kv kv.TxnKV,
	h Handler,
	subCluster SubCluster, // sessionManager
	alloc allocator,
	options ...ChannelmanagerOptV2,
) (*ChannelManagerImplV2, error) {
	m := &ChannelManagerImplV2{
		h:          h,
		ctx:        context.TODO(), // TODO
		factory:    NewChannelPolicyFactoryV1(),
		store:      NewChannelStoreV2(kv),
		subCluster: subCluster,
		allocator:  alloc,
	}

	if err := m.store.Reload(); err != nil {
		return nil, err
	}

	for _, opt := range options {
		opt(m)
	}

	m.balancePolicy = m.factory.NewBalancePolicy()
	m.lastActiveTimestamp = time.Now()
	return m, nil
}

func (m *ChannelManagerImplV2) Startup(ctx context.Context, legacyNodes, allNodes []int64) error {
	m.ctx, m.cancel = context.WithCancel(ctx)

	m.legacyNodes = typeutil.NewUniqueSet(legacyNodes...)

	m.mu.Lock()
	m.store.SetLegacyChannelByNode(legacyNodes...)
	oNodes := m.store.GetNodes()
	m.mu.Unlock()

	// Add new online nodes to the cluster.
	offLines, newOnLines := lo.Difference(oNodes, allNodes)
	lo.ForEach(newOnLines, func(nodeID int64, _ int) {
		m.AddNode(nodeID)
	})

	// Delete offlines from the cluster
	lo.ForEach(offLines, func(nodeID int64, _ int) {
		m.DeleteNode(nodeID)
	})

	m.mu.Lock()
	nodeChannels := m.store.GetNodeChannelsBy(
		WithAllNodes(),
		func(ch *StateChannel) bool {
			return m.h.CheckShouldDropChannel(ch.GetName())
		})
	m.mu.Unlock()

	for _, info := range nodeChannels {
		m.finishRemoveChannel(info.NodeID, lo.Values(info.Channels)...)
	}

	if m.balanceCheckLoop != nil {
		log.Info("starting channel balance loop")
		go m.balanceCheckLoop(m.ctx)
	}

	log.Info("cluster start up",
		zap.Int64s("allNodes", allNodes),
		zap.Int64s("legacyNodes", legacyNodes),
		zap.Int64s("oldNodes", oNodes),
		zap.Int64s("newOnlines", newOnLines),
		zap.Int64s("offLines", offLines))
	return nil
}

func (m *ChannelManagerImplV2) Close() {
	if m.cancel != nil {
		m.cancel()
	}
}

func (m *ChannelManagerImplV2) AddNode(nodeID UniqueID) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	log.Info("register node", zap.Int64("registered node", nodeID))

	m.store.AddNode(nodeID)
	updates := AvgAssignByCountPolicy(m.store.GetNodesChannels(), m.store.GetBufferChannelInfo(), m.legacyNodes.Collect())

	if updates == nil {
		log.Info("register node with no reassignment", zap.Int64("registered node", nodeID))
		return nil
	}

	err := m.execute(updates)
	if err != nil {
		log.Warn("fail to update channel operation updates into meta", zap.Error(err))
	}
	return err
}

// Release writes ToRelease channel watch states for a channel
func (m *ChannelManagerImplV2) Release(nodeID UniqueID, channelName string) error {
	log := log.With(
		zap.Int64("nodeID", nodeID),
		zap.String("channel", channelName),
	)

	// channel in bufferID are released already
	if nodeID == bufferID {
		return nil
	}

	log.Info("Releasing channel from watched node")
	ch, found := m.GetChannel(nodeID, channelName)
	if !found {
		return fmt.Errorf("fail to find matching nodeID: %d with channelName: %s", nodeID, channelName)
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	updates := NewChannelOpSet(NewChannelOp(nodeID, Release, ch))
	return m.execute(updates)
}

func (m *ChannelManagerImplV2) Watch(ctx context.Context, ch RWChannel) error {
	log := log.Ctx(ctx).With(zap.String("channel", ch.GetName()))
	m.mu.Lock()
	defer m.mu.Unlock()

	log.Info("Add channel")
	updates := NewChannelOpSet(NewChannelOp(bufferID, Watch, ch))
	err := m.execute(updates)
	if err != nil {
		log.Warn("fail to update new channel updates into meta",
			zap.Array("updates", updates), zap.Error(err))
	}

	// channel already written into meta, try to assign it to the cluster
	// not error is returned if failed, the assignment will retry later
	updates = AvgAssignByCountPolicy(m.store.GetNodesChannels(), m.store.GetBufferChannelInfo(), m.legacyNodes.Collect())
	if updates == nil {
		return nil
	}

	if err := m.execute(updates); err != nil {
		log.Warn("fail to assign channel, will retry later", zap.Array("updates", updates), zap.Error(err))
		return nil
	}

	log.Info("Assign channel", zap.Array("updates", updates))
	return nil
}

func (m *ChannelManagerImplV2) DeleteNode(nodeID UniqueID) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.legacyNodes.Remove(nodeID)
	info := m.store.GetNode(nodeID)
	if info == nil || len(info.Channels) == 0 {
		if nodeID != bufferID {
			m.store.RemoveNode(nodeID)
		}
		return nil
	}

	updates := NewChannelOpSet(
		NewDeleteOp(info.NodeID, lo.Values(info.Channels)...),
		NewChannelOp(bufferID, Watch, lo.Values(info.Channels)...),
	)
	log.Info("deregister node", zap.Int64("nodeID", nodeID), zap.Array("updates", updates))

	err := m.execute(updates)
	if err != nil {
		log.Warn("fail to update channel operation updates into meta", zap.Error(err))
		return err
	}

	if nodeID != bufferID {
		m.store.RemoveNode(nodeID)
	}
	return nil
}

// reassign reassigns a channel to another DataNode.
func (m *ChannelManagerImplV2) reassign(original *NodeChannelInfo) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	updates := AvgAssignByCountPolicy(m.store.GetNodesChannels(), original, m.legacyNodes.Collect())
	if updates != nil {
		return m.execute(updates)
	}

	if original.NodeID != bufferID {
		log.RatedWarn(5.0, "Failed to reassign channel to other nodes, assign to the original nodes",
			zap.Any("original node", original.NodeID),
			zap.Strings("channels", lo.Keys(original.Channels)),
		)
		updates := NewChannelOpSet(NewChannelOp(original.NodeID, Watch, lo.Values(original.Channels)...))
		return m.execute(updates)
	}

	return nil
}

func (m *ChannelManagerImplV2) Balance() {
	m.mu.Lock()
	defer m.mu.Unlock()

	watchedCluster := m.store.GetNodeChannelsBy(WithoutBufferNode(), WithChannelStates(Watched))
	updates := m.balancePolicy(watchedCluster)
	if updates == nil {
		return
	}

	log.Info("Channel balancer got new reAllocations:", zap.Array("assignment", updates))
	if err := m.execute(updates); err != nil {
		log.Warn("Channel balancer fail to execute", zap.Array("assignment", updates), zap.Error(err))
	}
}

func (m *ChannelManagerImplV2) Match(nodeID UniqueID, channel string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	info := m.store.GetNode(nodeID)
	if info == nil {
		return false
	}

	_, ok := info.Channels[channel]
	return ok
}

func (m *ChannelManagerImplV2) GetChannel(nodeID int64, channelName string) (RWChannel, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if nodeChannelInfo := m.store.GetNode(nodeID); nodeChannelInfo != nil {
		if ch, ok := nodeChannelInfo.Channels[channelName]; ok {
			return ch, true
		}
	}
	return nil, false
}

func (m *ChannelManagerImplV2) GetNodeIDByChannelName(channel string) (int64, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	nodeChannels := m.store.GetNodeChannelsBy(
		WithoutBufferNode(),
		WithChannelName(channel))

	if len(nodeChannels) > 0 {
		return nodeChannels[0].NodeID, true
	}

	return 0, false
}

func (m *ChannelManagerImplV2) GetNodeChannelsByCollectionID(collectionID int64) map[int64][]string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	nodeChs := make(map[UniqueID][]string)
	nodeChannels := m.store.GetNodeChannelsBy(
		WithoutBufferNode(),
		WithCollectionIDV2(collectionID))
	lo.ForEach(nodeChannels, func(info *NodeChannelInfo, _ int) {
		nodeChs[info.NodeID] = lo.Keys(info.Channels)
	})
	return nodeChs
}

func (m *ChannelManagerImplV2) GetChannelsByCollectionID(collectionID int64) []RWChannel {
	m.mu.RLock()
	defer m.mu.RUnlock()
	channels := []RWChannel{}

	nodeChannels := m.store.GetNodeChannelsBy(
		WithAllNodes(),
		WithCollectionIDV2(collectionID))
	lo.ForEach(nodeChannels, func(info *NodeChannelInfo, _ int) {
		channels = append(channels, lo.Values(info.Channels)...)
	})
	return channels
}

func (m *ChannelManagerImplV2) GetChannelNamesByCollectionID(collectionID int64) []string {
	channels := m.GetChannelsByCollectionID(collectionID)
	return lo.Map(channels, func(ch RWChannel, _ int) string {
		return ch.GetName()
	})
}

func (m *ChannelManagerImplV2) FindWatcher(channel string) (UniqueID, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	infos := m.store.GetNodesChannels()
	for _, info := range infos {
		for _, channelInfo := range info.Channels {
			if channelInfo.GetName() == channel {
				return info.NodeID, nil
			}
		}
	}

	// channel in buffer
	bufferInfo := m.store.GetBufferChannelInfo()
	for _, channelInfo := range bufferInfo.Channels {
		if channelInfo.GetName() == channel {
			return bufferID, errChannelInBuffer
		}
	}
	return 0, errChannelNotWatched
}

// unsafe innter func
func (m *ChannelManagerImplV2) removeChannel(nodeID int64, ch RWChannel) error {
	op := NewChannelOpSet(NewChannelOp(nodeID, Delete, ch))
	log.Info("remove channel assignment",
		zap.String("channel", ch.GetName()),
		zap.Int64("assignment", nodeID),
		zap.Int64("collectionID", ch.GetCollectionID()))
	return m.store.Update(op)
}

func (m *ChannelManagerImplV2) CheckLoop(ctx context.Context) {
	balanceTicker := time.NewTicker(Params.DataCoordCfg.ChannelBalanceInterval.GetAsDuration(time.Second))
	defer balanceTicker.Stop()
	checkTicker := time.NewTicker(Params.DataCoordCfg.ChannelCheckInterval.GetAsDuration(time.Second))
	defer checkTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Info("background checking channels loop quit")
			return
		case <-balanceTicker.C:
			// balance
			if time.Since(m.lastActiveTimestamp) >= Params.DataCoordCfg.ChannelBalanceSilentDuration.GetAsDuration(time.Second) {
				m.Balance()
			}
		case <-checkTicker.C:
			m.AdvanceChannelState()
		}
	}
}

func (m *ChannelManagerImplV2) AdvanceChannelState() {
	m.mu.RLock()
	standbys := m.store.GetNodeChannelsBy(WithAllNodes(), WithChannelStates(Standby))
	toNotifies := m.store.GetNodeChannelsBy(WithoutBufferNode(), WithChannelStates(ToWatch, ToRelease))
	toChecks := m.store.GetNodeChannelsBy(WithoutBufferNode(), WithChannelStates(Watching, Releasing))
	m.mu.RUnlock()

	// Processing standby channels
	updatedStandbys := m.advanceStandbys(standbys)
	updatedToCheckes := m.advanceToChecks(toChecks)
	updatedToNotifies := m.advanceToNotifies(toNotifies)

	if updatedStandbys || updatedToCheckes || updatedToNotifies {
		m.lastActiveTimestamp = time.Now()
	}
}

func (m *ChannelManagerImplV2) finishRemoveChannel(nodeID int64, channels ...RWChannel) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, ch := range channels {
		if err := m.removeChannel(nodeID, ch); err != nil {
			log.Warn("Failed to remove channel", zap.Any("channel", ch), zap.Error(err))
			continue
		}

		if err := m.h.FinishDropChannel(ch.GetName(), ch.GetCollectionID()); err != nil {
			log.Warn("Failed to finish drop channel", zap.Any("channel", ch), zap.Error(err))
			continue
		}
	}
}

func (m *ChannelManagerImplV2) advanceStandbys(standbys []*NodeChannelInfo) bool {
	var advanced bool = false
	for _, nodeAssign := range standbys {
		validChannels := make(map[string]RWChannel)
		for chName, ch := range nodeAssign.Channels {
			// drop marked-drop channels
			if m.h.CheckShouldDropChannel(chName) {
				m.finishRemoveChannel(nodeAssign.NodeID, ch)
				continue
			}
			validChannels[chName] = ch
		}
		nodeAssign.Channels = validChannels

		if len(nodeAssign.Channels) == 0 {
			continue
		}

		chNames := lo.Keys(validChannels)
		if err := m.reassign(nodeAssign); err != nil {
			log.Warn("Reassign channels fail",
				zap.Int64("nodeID", nodeAssign.NodeID),
				zap.Strings("channels", chNames),
			)
		}

		log.Info("Reassign standby channels to node",
			zap.Int64("nodeID", nodeAssign.NodeID),
			zap.Strings("channels", chNames),
		)
		advanced = true
	}

	return advanced
}

func (m *ChannelManagerImplV2) advanceToNotifies(toNotifies []*NodeChannelInfo) bool {
	var advanced bool = false
	for _, nodeAssign := range toNotifies {
		channelCount := len(nodeAssign.Channels)
		if channelCount == 0 {
			continue
		}

		var (
			succeededChannels = make([]RWChannel, 0, channelCount)
			failedChannels    = make([]RWChannel, 0, channelCount)
			futures           = make([]*conc.Future[any], 0, channelCount)
		)

		chNames := lo.Keys(nodeAssign.Channels)
		log.Info("Notify channel operations to datanode",
			zap.Int64("assignment", nodeAssign.NodeID),
			zap.Int("total operation count", len(nodeAssign.Channels)),
			zap.Strings("channel names", chNames),
		)
		for _, ch := range nodeAssign.Channels {
			innerCh := ch

			future := getOrCreateIOPool().Submit(func() (any, error) {
				err := m.Notify(nodeAssign.NodeID, innerCh.GetWatchInfo())
				return innerCh, err
			})
			futures = append(futures, future)
		}

		for _, f := range futures {
			ch, err := f.Await()
			if err != nil {
				failedChannels = append(failedChannels, ch.(RWChannel))
			} else {
				succeededChannels = append(succeededChannels, ch.(RWChannel))
				advanced = true
			}
		}

		log.Info("Finish to notify channel operations to datanode",
			zap.Int64("assignment", nodeAssign.NodeID),
			zap.Int("operation count", channelCount),
			zap.Int("success count", len(succeededChannels)),
			zap.Int("failure count", len(failedChannels)),
		)
		m.mu.Lock()
		m.store.UpdateState(false, failedChannels...)
		m.store.UpdateState(true, succeededChannels...)
		m.mu.Unlock()
	}

	return advanced
}

type poolResult struct {
	successful bool
	ch         RWChannel
}

func (m *ChannelManagerImplV2) advanceToChecks(toChecks []*NodeChannelInfo) bool {
	var advanced bool = false
	for _, nodeAssign := range toChecks {
		if len(nodeAssign.Channels) == 0 {
			continue
		}

		futures := make([]*conc.Future[any], 0, len(nodeAssign.Channels))

		chNames := lo.Keys(nodeAssign.Channels)
		log.Info("Check ToWatch/ToRelease channel operations progress",
			zap.Int("channel count", len(nodeAssign.Channels)),
			zap.Strings("channel names", chNames),
		)

		for _, ch := range nodeAssign.Channels {
			innerCh := ch

			future := getOrCreateIOPool().Submit(func() (any, error) {
				successful, got := m.Check(nodeAssign.NodeID, innerCh.GetWatchInfo())
				if got {
					return poolResult{
						successful: successful,
						ch:         innerCh,
					}, nil
				}
				return nil, errors.New("Got results with no progress")
			})
			futures = append(futures, future)
		}

		for _, f := range futures {
			got, err := f.Await()
			if err == nil {
				m.mu.Lock()
				result := got.(poolResult)
				m.store.UpdateState(result.successful, result.ch)
				m.mu.Unlock()

				advanced = true
			}
		}

		log.Info("Finish to Check ToWatch/ToRelease channel operations progress",
			zap.Int("channel count", len(nodeAssign.Channels)),
			zap.Strings("channel names", chNames),
		)
	}
	return advanced
}

func (m *ChannelManagerImplV2) Notify(nodeID int64, info *datapb.ChannelWatchInfo) error {
	log := log.With(
		zap.String("channel", info.GetVchan().GetChannelName()),
		zap.Int64("assignment", nodeID),
		zap.String("operation", info.GetState().String()),
	)
	log.Info("Notify channel operation")
	err := m.subCluster.NotifyChannelOperation(m.ctx, nodeID, &datapb.ChannelOperationsRequest{Infos: []*datapb.ChannelWatchInfo{info}})
	if err != nil {
		log.Warn("Fail to notify channel operations", zap.Error(err))
		return err
	}
	log.Debug("Success to notify channel operations")
	return nil
}

func (m *ChannelManagerImplV2) Check(nodeID int64, info *datapb.ChannelWatchInfo) (successful bool, got bool) {
	log := log.With(
		zap.Int64("opID", info.GetOpID()),
		zap.Int64("nodeID", nodeID),
		zap.String("check operation", info.GetState().String()),
		zap.String("channel", info.GetVchan().GetChannelName()),
	)
	resp, err := m.subCluster.CheckChannelOperationProgress(m.ctx, nodeID, info)
	if err != nil {
		log.Warn("Fail to check channel operation progress")
		return false, false
	}
	log.Info("Got channel operation progress",
		zap.String("got state", resp.GetState().String()),
		zap.Int32("progress", resp.GetProgress()))
	switch info.GetState() {
	case datapb.ChannelWatchState_ToWatch:
		if resp.GetState() == datapb.ChannelWatchState_ToWatch {
			return false, false
		}
		if resp.GetState() == datapb.ChannelWatchState_WatchSuccess {
			return true, true
		}
		if resp.GetState() == datapb.ChannelWatchState_WatchFailure {
			return false, true
		}
	case datapb.ChannelWatchState_ToRelease:
		if resp.GetState() == datapb.ChannelWatchState_ToRelease {
			return false, false
		}
		if resp.GetState() == datapb.ChannelWatchState_ReleaseSuccess {
			return true, true
		}
		if resp.GetState() == datapb.ChannelWatchState_ReleaseFailure {
			return false, true
		}
	}
	return false, false
}

func (m *ChannelManagerImplV2) execute(updates *ChannelOpSet) error {
	for _, op := range updates.ops {
		if op.Type != Delete {
			if err := m.fillChannelWatchInfo(op); err != nil {
				log.Warn("fail to fill channel watch info", zap.Error(err))
				return err
			}
		}
	}

	return m.store.Update(updates)
}

// fillChannelWatchInfoWithState updates the channel op by filling in channel watch info.
func (m *ChannelManagerImplV2) fillChannelWatchInfo(op *ChannelOp) error {
	startTs := time.Now().Unix()
	for _, ch := range op.Channels {
		vcInfo := m.h.GetDataVChanPositions(ch, allPartitionID)
		opID, err := m.allocator.allocID(context.Background())
		if err != nil {
			return err
		}

		info := &datapb.ChannelWatchInfo{
			Vchan:   vcInfo,
			StartTs: startTs,
			State:   inferStateByOpType(op.Type),
			Schema:  ch.GetSchema(),
			OpID:    opID,
		}
		ch.UpdateWatchInfo(info)
	}
	return nil
}

func inferStateByOpType(opType ChannelOpType) datapb.ChannelWatchState {
	switch opType {
	case Watch:
		return datapb.ChannelWatchState_ToWatch
	case Release:
		return datapb.ChannelWatchState_ToRelease
	default:
		return datapb.ChannelWatchState_ToWatch
	}
}
