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
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/blang/semver/v4"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/milvuspb"
	datanodeclient "github.com/milvus-io/milvus/internal/distributed/datanode/client"
	rootcoordclient "github.com/milvus-io/milvus/internal/distributed/rootcoord/client"
	etcdkv "github.com/milvus-io/milvus/internal/kv/etcd"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/metrics"
	"github.com/milvus-io/milvus/internal/mq/msgstream"
	"github.com/milvus-io/milvus/internal/mq/msgstream/mqwrapper"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/internal/types"
	"github.com/milvus-io/milvus/internal/util/commonpbutil"
	"github.com/milvus-io/milvus/internal/util/dependency"
	"github.com/milvus-io/milvus/internal/util/funcutil"
	"github.com/milvus-io/milvus/internal/util/logutil"
	"github.com/milvus-io/milvus/internal/util/metricsinfo"
	"github.com/milvus-io/milvus/internal/util/paramtable"
	"github.com/milvus-io/milvus/internal/util/retry"
	"github.com/milvus-io/milvus/internal/util/sessionutil"
	"github.com/milvus-io/milvus/internal/util/timerecord"
	"github.com/milvus-io/milvus/internal/util/tsoutil"
	"github.com/milvus-io/milvus/internal/util/typeutil"
)

const (
	connEtcdMaxRetryTime = 100
	allPartitionID       = 0 // paritionID means no filtering
)

var (
	// TODO: sunby put to config
	enableTtChecker           = true
	ttCheckerName             = "dataTtChecker"
	ttMaxInterval             = 2 * time.Minute
	ttCheckerWarnMsg          = fmt.Sprintf("Datacoord haven't received tt for %f minutes", ttMaxInterval.Minutes())
	segmentTimedFlushDuration = 10.0
)

type (
	// UniqueID shortcut for typeutil.UniqueID
	UniqueID = typeutil.UniqueID
	// Timestamp shortcurt for typeutil.Timestamp
	Timestamp = typeutil.Timestamp
)

type dataNodeCreatorFunc func(ctx context.Context, addr string) (types.DataNode, error)
type rootCoordCreatorFunc func(ctx context.Context, metaRootPath string, etcdClient *clientv3.Client) (types.RootCoord, error)

// makes sure Server implements `DataCoord`
var _ types.DataCoord = (*Server)(nil)

var Params *paramtable.ComponentParam = paramtable.Get()

// Server implements `types.DataCoord`
// handles Data Coordinator related jobs
type Server struct {
	ctx              context.Context
	serverLoopCtx    context.Context
	serverLoopCancel context.CancelFunc
	serverLoopWg     sync.WaitGroup
	quitCh           chan struct{}
	stateCode        atomic.Value
	helper           ServerHelper

	etcdCli          *clientv3.Client
	address          string
	kvClient         *etcdkv.EtcdKV
	meta             *meta
	segmentManager   Manager
	allocator        allocator
	cluster          *Cluster
	sessionManager   *SessionManager
	channelManager   *ChannelManager
	rootCoordClient  types.RootCoord
	garbageCollector *garbageCollector
	gcOpt            GcOption
	handler          Handler

	compactionTrigger trigger
	compactionHandler compactionPlanContext

	metricsCacheManager *metricsinfo.MetricsCacheManager

	flushCh chan UniqueID
	factory dependency.Factory

	session   *sessionutil.Session
	dnEventCh <-chan *sessionutil.SessionEvent
	//icEventCh <-chan *sessionutil.SessionEvent
	qcEventCh <-chan *sessionutil.SessionEvent

	enableActiveStandBy bool
	activateFunc        func()

	dataNodeCreator        dataNodeCreatorFunc
	rootCoordClientCreator rootCoordCreatorFunc
	indexCoord             types.IndexCoord

	segReferManager *SegmentReferenceManager
}

// ServerHelper datacoord server injection helper
type ServerHelper struct {
	eventAfterHandleDataNodeTt func()
}

func defaultServerHelper() ServerHelper {
	return ServerHelper{
		eventAfterHandleDataNodeTt: func() {},
	}
}

// Option utility function signature to set DataCoord server attributes
type Option func(svr *Server)

// SetRootCoordCreator returns an `Option` setting RootCoord creator with provided parameter
func SetRootCoordCreator(creator rootCoordCreatorFunc) Option {
	return func(svr *Server) {
		svr.rootCoordClientCreator = creator
	}
}

// SetServerHelper returns an `Option` setting ServerHelp with provided parameter
func SetServerHelper(helper ServerHelper) Option {
	return func(svr *Server) {
		svr.helper = helper
	}
}

// SetCluster returns an `Option` setting Cluster with provided parameter
func SetCluster(cluster *Cluster) Option {
	return func(svr *Server) {
		svr.cluster = cluster
	}
}

// SetDataNodeCreator returns an `Option` setting DataNode create function
func SetDataNodeCreator(creator dataNodeCreatorFunc) Option {
	return func(svr *Server) {
		svr.dataNodeCreator = creator
	}
}

// SetSegmentManager returns an Option to set SegmentManager
func SetSegmentManager(manager Manager) Option {
	return func(svr *Server) {
		svr.segmentManager = manager
	}
}

// CreateServer creates a `Server` instance
func CreateServer(ctx context.Context, factory dependency.Factory, opts ...Option) *Server {
	rand.Seed(time.Now().UnixNano())
	s := &Server{
		ctx:                    ctx,
		quitCh:                 make(chan struct{}),
		factory:                factory,
		flushCh:                make(chan UniqueID, 1024),
		dataNodeCreator:        defaultDataNodeCreatorFunc,
		rootCoordClientCreator: defaultRootCoordCreatorFunc,
		helper:                 defaultServerHelper(),
		metricsCacheManager:    metricsinfo.NewMetricsCacheManager(),
		enableActiveStandBy:    Params.DataCoordCfg.EnableActiveStandby,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func defaultDataNodeCreatorFunc(ctx context.Context, addr string) (types.DataNode, error) {
	return datanodeclient.NewClient(ctx, addr)
}

func defaultRootCoordCreatorFunc(ctx context.Context, metaRootPath string, client *clientv3.Client) (types.RootCoord, error) {
	return rootcoordclient.NewClient(ctx, metaRootPath, client)
}

// QuitSignal returns signal when server quits
func (s *Server) QuitSignal() <-chan struct{} {
	return s.quitCh
}

// Register registers data service at etcd
func (s *Server) Register() error {
	s.session.Register()
	if s.enableActiveStandBy {
		s.session.ProcessActiveStandBy(s.activateFunc)
	}
	go s.session.LivenessCheck(s.serverLoopCtx, func() {
		logutil.Logger(s.ctx).Error("disconnected from etcd and exited", zap.Int64("serverID", s.session.ServerID))
		if err := s.Stop(); err != nil {
			logutil.Logger(s.ctx).Fatal("failed to stop server", zap.Error(err))
		}
		// manually send signal to starter goroutine
		if s.session.TriggerKill {
			if p, err := os.FindProcess(os.Getpid()); err == nil {
				p.Signal(syscall.SIGINT)
			}
		}
	})
	return nil
}

func (s *Server) initSession() error {
	s.session = sessionutil.NewSession(s.ctx, Params.EtcdCfg.MetaRootPath, s.etcdCli)
	if s.session == nil {
		return errors.New("failed to initialize session")
	}
	s.session.Init(typeutil.DataCoordRole, s.address, true, true)
	s.session.SetEnableActiveStandBy(s.enableActiveStandBy)
	paramtable.SetNodeID(s.session.ServerID)
	return nil
}

// Init change server state to Initializing
func (s *Server) Init() error {
	var err error
	s.stateCode.Store(commonpb.StateCode_Initializing)
	s.factory.Init(Params)

	if err = s.initRootCoordClient(); err != nil {
		return err
	}

	storageCli, err := s.newChunkManagerFactory()
	if err != nil {
		return err
	}

	if err = s.initMeta(storageCli.RootPath()); err != nil {
		return err
	}

	s.handler = newServerHandler(s)

	if err = s.initCluster(); err != nil {
		return err
	}

	s.allocator = newRootCoordAllocator(s.rootCoordClient)

	if err = s.initSession(); err != nil {
		return err
	}

	if err = s.initServiceDiscovery(); err != nil {
		return err
	}

	if Params.DataCoordCfg.EnableCompaction {
		s.createCompactionHandler()
		s.createCompactionTrigger()
	}
	s.initSegmentManager()

	s.initGarbageCollection(storageCli)

	return nil
}

// Start initialize `Server` members and start loops, follow steps are taken:
//  1. initialize message factory parameters
//  2. initialize root coord client, meta, datanode cluster, segment info channel,
//     allocator, segment manager
//  3. start service discovery and server loops, which includes message stream handler (segment statistics,datanode tt)
//     datanodes etcd watch, etcd alive check and flush completed status check
//  4. set server state to Healthy
func (s *Server) Start() error {
	s.compactionHandler.start()
	s.compactionTrigger.start()

	if s.enableActiveStandBy {
		s.activateFunc = func() {
			// todo complete the activateFunc
			log.Info("datacoord switch from standby to active, activating")
			s.startServerLoop()
			s.stateCode.Store(commonpb.StateCode_Healthy)
			logutil.Logger(s.ctx).Info("startup success")
		}
		s.stateCode.Store(commonpb.StateCode_StandBy)
		logutil.Logger(s.ctx).Info("DataCoord enter standby mode successfully")
	} else {
		s.startServerLoop()
		s.stateCode.Store(commonpb.StateCode_Healthy)
		logutil.Logger(s.ctx).Info("DataCoord startup successfully")
	}

	Params.DataCoordCfg.CreatedTime = time.Now()
	Params.DataCoordCfg.UpdatedTime = time.Now()

	// DataCoord (re)starts successfully and starts to collection segment stats
	// data from all DataNode.
	// This will prevent DataCoord from missing out any important segment stats
	// data while offline.
	log.Info("DataCoord (re)starts successfully and re-collecting segment stats from DataNodes")
	s.reCollectSegmentStats(s.ctx)

	return nil
}

func (s *Server) initCluster() error {
	if s.cluster != nil {
		return nil
	}

	var err error
	s.channelManager, err = NewChannelManager(s.kvClient, s.handler, withMsgstreamFactory(s.factory), withStateChecker())
	if err != nil {
		return err
	}
	s.sessionManager = NewSessionManager(withSessionCreator(s.dataNodeCreator))
	s.cluster = NewCluster(s.sessionManager, s.channelManager)
	return nil
}

func (s *Server) SetAddress(address string) {
	s.address = address
}

// SetEtcdClient sets etcd client for datacoord.
func (s *Server) SetEtcdClient(client *clientv3.Client) {
	s.etcdCli = client
}

func (s *Server) SetIndexCoord(indexCoord types.IndexCoord) {
	s.indexCoord = indexCoord
}

func (s *Server) createCompactionHandler() {
	s.compactionHandler = newCompactionPlanHandler(s.sessionManager, s.channelManager, s.meta, s.allocator, s.flushCh, s.segReferManager)
}

func (s *Server) stopCompactionHandler() {
	s.compactionHandler.stop()
}

func (s *Server) createCompactionTrigger() {
	s.compactionTrigger = newCompactionTrigger(s.meta, s.compactionHandler, s.allocator, s.segReferManager, s.indexCoord, s.handler)
}

func (s *Server) stopCompactionTrigger() {
	s.compactionTrigger.stop()
}

func (s *Server) newChunkManagerFactory() (storage.ChunkManager, error) {
	chunkManagerFactory := storage.NewChunkManagerFactoryWithParam(Params)
	cli, err := chunkManagerFactory.NewPersistentStorageChunkManager(s.ctx)
	if err != nil {
		log.Error("chunk manager init failed", zap.Error(err))
		return nil, err
	}
	return cli, err
}

func (s *Server) initGarbageCollection(cli storage.ChunkManager) {
	s.garbageCollector = newGarbageCollector(s.meta, s.handler, s.segReferManager, s.indexCoord, GcOption{
		cli:              cli,
		enabled:          Params.DataCoordCfg.EnableGarbageCollection,
		checkInterval:    Params.DataCoordCfg.GCInterval,
		missingTolerance: Params.DataCoordCfg.GCMissingTolerance,
		dropTolerance:    Params.DataCoordCfg.GCDropTolerance,
	})
}

func (s *Server) initServiceDiscovery() error {
	r := semver.MustParseRange(">=2.1.2")
	sessions, rev, err := s.session.GetSessionsWithVersionRange(typeutil.DataNodeRole, r)
	if err != nil {
		log.Warn("DataCoord failed to init service discovery", zap.Error(err))
		return err
	}
	log.Info("DataCoord success to get DataNode sessions", zap.Any("sessions", sessions))

	datanodes := make([]*NodeInfo, 0, len(sessions))
	for _, session := range sessions {
		info := &NodeInfo{
			NodeID:  session.ServerID,
			Address: session.Address,
		}
		datanodes = append(datanodes, info)
	}

	s.cluster.Startup(s.ctx, datanodes)

	// TODO implement rewatch logic
	s.dnEventCh = s.session.WatchServicesWithVersionRange(typeutil.DataNodeRole, r, rev+1, nil)

	//icSessions, icRevision, err := s.session.GetSessions(typeutil.IndexCoordRole)
	//if err != nil {
	//	log.Error("DataCoord get IndexCoord session failed", zap.Error(err))
	//	return err
	//}
	//serverIDs := make([]UniqueID, 0, len(icSessions))
	//for _, session := range icSessions {
	//	serverIDs = append(serverIDs, session.ServerID)
	//}
	//s.icEventCh = s.session.WatchServices(typeutil.IndexCoordRole, icRevision+1, nil)

	qcSessions, qcRevision, err := s.session.GetSessions(typeutil.QueryCoordRole)
	if err != nil {
		log.Error("DataCoord get QueryCoord session failed", zap.Error(err))
		return err
	}
	serverIDs := make([]UniqueID, 0, len(qcSessions))
	for _, session := range qcSessions {
		serverIDs = append(serverIDs, session.ServerID)
	}
	s.qcEventCh = s.session.WatchServices(typeutil.QueryCoordRole, qcRevision+1, nil)

	s.segReferManager, err = NewSegmentReferenceManager(s.kvClient, serverIDs)
	return err
}

func (s *Server) initSegmentManager() {
	if s.segmentManager == nil {
		s.segmentManager = newSegmentManager(s.meta, s.allocator, s.rootCoordClient)
	}
}

func (s *Server) initMeta(chunkManagerRootPath string) error {
	etcdKV := etcdkv.NewEtcdKV(s.etcdCli, Params.EtcdCfg.MetaRootPath)
	s.kvClient = etcdKV
	reloadEtcdFn := func() error {
		var err error
		s.meta, err = newMeta(s.ctx, s.kvClient, chunkManagerRootPath)
		if err != nil {
			return err
		}
		return nil
	}
	return retry.Do(s.ctx, reloadEtcdFn, retry.Attempts(connEtcdMaxRetryTime))
}

func (s *Server) startServerLoop() {
	s.serverLoopCtx, s.serverLoopCancel = context.WithCancel(s.ctx)
	s.serverLoopWg.Add(3)
	s.startDataNodeTtLoop(s.serverLoopCtx)
	s.startWatchService(s.serverLoopCtx)
	s.startFlushLoop(s.serverLoopCtx)
	s.garbageCollector.start()
}

// startDataNodeTtLoop start a goroutine to recv data node tt msg from msgstream
// tt msg stands for the currently consumed timestamp for each channel
func (s *Server) startDataNodeTtLoop(ctx context.Context) {
	ttMsgStream, err := s.factory.NewMsgStream(ctx)
	if err != nil {
		log.Error("DataCoord failed to create timetick channel", zap.Error(err))
		panic(err)
	}
	ttMsgStream.AsConsumer([]string{Params.CommonCfg.DataCoordTimeTick},
		Params.CommonCfg.DataCoordSubName, mqwrapper.SubscriptionPositionLatest)
	log.Info("DataCoord creates the timetick channel consumer",
		zap.String("timeTickChannel", Params.CommonCfg.DataCoordTimeTick),
		zap.String("subscription", Params.CommonCfg.DataCoordSubName))
	ttMsgStream.Start()

	go s.handleDataNodeTimetickMsgstream(ctx, ttMsgStream)
}

func (s *Server) handleDataNodeTimetickMsgstream(ctx context.Context, ttMsgStream msgstream.MsgStream) {
	var checker *timerecord.LongTermChecker
	if enableTtChecker {
		checker = timerecord.NewLongTermChecker(ctx, ttCheckerName, ttMaxInterval, ttCheckerWarnMsg)
		checker.Start()
		defer checker.Stop()
	}

	defer logutil.LogPanic()
	defer s.serverLoopWg.Done()
	defer func() {
		// https://github.com/milvus-io/milvus/issues/15659
		// msgstream service closed before datacoord quits
		defer func() {
			if x := recover(); x != nil {
				log.Error("Failed to close ttMessage", zap.Any("recovered", x))
			}
		}()
		ttMsgStream.Close()
	}()
	for {
		select {
		case <-ctx.Done():
			log.Info("DataNode timetick loop shutdown")
			return
		case msgPack, ok := <-ttMsgStream.Chan():
			if !ok || msgPack == nil || len(msgPack.Msgs) == 0 {
				log.Info("receive nil timetick msg and shutdown timetick channel")
				return
			}

			for _, msg := range msgPack.Msgs {
				ttMsg, ok := msg.(*msgstream.DataNodeTtMsg)
				if !ok {
					log.Warn("receive unexpected msg type from tt channel")
					continue
				}
				if enableTtChecker {
					checker.Check()
				}

				if err := s.handleTimetickMessage(ctx, ttMsg); err != nil {
					log.Error("failed to handle timetick message", zap.Error(err))
					continue
				}
			}
			s.helper.eventAfterHandleDataNodeTt()
		}
	}
}

func (s *Server) handleTimetickMessage(ctx context.Context, ttMsg *msgstream.DataNodeTtMsg) error {
	ch := ttMsg.GetChannelName()
	ts := ttMsg.GetTimestamp()
	physical, _ := tsoutil.ParseTS(ts)
	if time.Since(physical).Minutes() > 1 {
		// if lag behind, log every 1 mins about
		log.RatedWarn(60.0, "time tick lag behind for more than 1 minutes", zap.String("channel", ch), zap.Time("timetick", physical))
	}

	sub := tsoutil.SubByNow(ts)
	pChannelName := funcutil.ToPhysicalChannel(ch)
	metrics.DataCoordConsumeDataNodeTimeTickLag.
		WithLabelValues(fmt.Sprint(paramtable.GetNodeID()), pChannelName).
		Set(float64(sub))

	s.updateSegmentStatistics(ttMsg.GetSegmentsStats())

	if err := s.segmentManager.ExpireAllocations(ch, ts); err != nil {
		return fmt.Errorf("expire allocations: %w", err)
	}

	flushableIDs, err := s.segmentManager.GetFlushableSegments(ctx, ch, ts)
	if err != nil {
		return fmt.Errorf("get flushable segments: %w", err)
	}
	flushableSegments := s.getFlushableSegmentsInfo(flushableIDs)

	staleSegments := s.getStaleSegmentsInfo(ch)
	staleSegments = s.filterWithFlushableSegments(staleSegments, flushableIDs)

	if len(flushableSegments)+len(staleSegments) == 0 {
		return nil
	}

	log.Info("start flushing segments",
		zap.Int64s("segment IDs", flushableIDs),
		zap.Int("# of stale/mark segments", len(staleSegments)))
	// update segment last update triggered time
	// it's ok to fail flushing, since next timetick after duration will re-trigger
	s.setLastFlushTime(flushableSegments)
	s.setLastFlushTime(staleSegments)

	finfo, minfo := make([]*datapb.SegmentInfo, 0, len(flushableSegments)), make([]*datapb.SegmentInfo, 0, len(staleSegments))
	for _, info := range flushableSegments {
		finfo = append(finfo, info.SegmentInfo)
	}
	for _, info := range staleSegments {
		minfo = append(minfo, info.SegmentInfo)
	}
	err = s.cluster.Flush(s.ctx, ttMsg.GetBase().GetSourceID(), ch, finfo, minfo)
	if err != nil {
		log.Warn("handle")
		return err
	}

	return nil
}

func (s *Server) updateSegmentStatistics(stats []*datapb.SegmentStats) {
	for _, stat := range stats {
		// Log if # of rows is updated.
		if s.meta.GetSegmentUnsafe(stat.GetSegmentID()) != nil &&
			s.meta.GetSegmentUnsafe(stat.GetSegmentID()).GetNumOfRows() != stat.GetNumRows() {
			log.Info("Updating segment number of rows",
				zap.Int64("segment ID", stat.GetSegmentID()),
				zap.Int64("old value", s.meta.GetSegmentUnsafe(stat.GetSegmentID()).GetNumOfRows()),
				zap.Int64("new value", stat.GetNumRows()),
			)
		}
		s.meta.SetCurrentRows(stat.GetSegmentID(), stat.GetNumRows())
	}
}

func (s *Server) getFlushableSegmentsInfo(flushableIDs []int64) []*SegmentInfo {
	res := make([]*SegmentInfo, 0, len(flushableIDs))
	for _, id := range flushableIDs {
		sinfo := s.meta.GetSegment(id)
		if sinfo == nil {
			log.Error("get segment from meta error", zap.Int64("id", id))
			continue
		}
		res = append(res, sinfo)
	}
	return res
}

func (s *Server) getStaleSegmentsInfo(ch string) []*SegmentInfo {
	return s.meta.SelectSegments(func(info *SegmentInfo) bool {
		return isSegmentHealthy(info) &&
			info.GetInsertChannel() == ch &&
			!info.lastFlushTime.IsZero() &&
			time.Since(info.lastFlushTime).Minutes() >= segmentTimedFlushDuration &&
			info.GetNumOfRows() != 0
	})
}

func (s *Server) filterWithFlushableSegments(staleSegments []*SegmentInfo, flushableIDs []int64) []*SegmentInfo {
	filter := map[int64]struct{}{}
	for _, sid := range flushableIDs {
		filter[sid] = struct{}{}
	}

	res := make([]*SegmentInfo, 0, len(staleSegments))
	for _, sinfo := range staleSegments {
		if _, ok := filter[sinfo.GetID()]; ok {
			continue
		}
		res = append(res, sinfo)
	}
	return res
}

func (s *Server) setLastFlushTime(segments []*SegmentInfo) {
	for _, sinfo := range segments {
		s.meta.SetLastFlushTime(sinfo.GetID(), time.Now())
	}
}

// start a goroutine wto watch services
func (s *Server) startWatchService(ctx context.Context) {
	go s.watchService(ctx)
}

func (s *Server) stopServiceWatch() {
	// ErrCompacted is handled inside SessionWatcher, which means there is some other error occurred, closing server.
	logutil.Logger(s.ctx).Error("watch service channel closed", zap.Int64("serverID", s.session.ServerID))
	go s.Stop()
	if s.session.TriggerKill {
		if p, err := os.FindProcess(os.Getpid()); err == nil {
			p.Signal(syscall.SIGINT)
		}
	}
}

func (s *Server) processSessionEvent(ctx context.Context, role string, event *sessionutil.SessionEvent) {
	switch event.EventType {
	case sessionutil.SessionAddEvent:
		log.Info("there is a new service online",
			zap.String("server role", role),
			zap.Int64("server ID", event.Session.ServerID))

	case sessionutil.SessionDelEvent:
		log.Warn("there is service offline",
			zap.String("server role", role),
			zap.Int64("server ID", event.Session.ServerID))
		if err := retry.Do(ctx, func() error {
			return s.segReferManager.ReleaseSegmentsLockByNodeID(event.Session.ServerID)
		}, retry.Attempts(100)); err != nil {
			panic(err)
		}
	}
}

// watchService watches services.
func (s *Server) watchService(ctx context.Context) {
	defer logutil.LogPanic()
	defer s.serverLoopWg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Info("watch service shutdown")
			return
		case event, ok := <-s.dnEventCh:
			if !ok {
				s.stopServiceWatch()
				return
			}
			if err := s.handleSessionEvent(ctx, event); err != nil {
				go func() {
					if err := s.Stop(); err != nil {
						log.Warn("DataCoord server stop error", zap.Error(err))
					}
				}()
				return
			}
		//case event, ok := <-s.icEventCh:
		//	if !ok {
		//		s.stopServiceWatch()
		//		return
		//	}
		//	s.processSessionEvent(ctx, "IndexCoord", event)
		case event, ok := <-s.qcEventCh:
			if !ok {
				s.stopServiceWatch()
				return
			}
			s.processSessionEvent(ctx, "QueryCoord", event)
		}
	}
}

// handles session events - DataNodes Add/Del
func (s *Server) handleSessionEvent(ctx context.Context, event *sessionutil.SessionEvent) error {
	if event == nil {
		return nil
	}
	info := &datapb.DataNodeInfo{
		Address:  event.Session.Address,
		Version:  event.Session.ServerID,
		Channels: []*datapb.ChannelStatus{},
	}
	node := &NodeInfo{
		NodeID:  event.Session.ServerID,
		Address: event.Session.Address,
	}
	switch event.EventType {
	case sessionutil.SessionAddEvent:
		log.Info("received datanode register",
			zap.String("address", info.Address),
			zap.Int64("serverID", info.Version))
		if err := s.cluster.Register(node); err != nil {
			log.Warn("failed to register node", zap.Int64("id", node.NodeID), zap.String("address", node.Address), zap.Error(err))
			return err
		}
		s.metricsCacheManager.InvalidateSystemInfoMetrics()
	case sessionutil.SessionDelEvent:
		log.Info("received datanode unregister",
			zap.String("address", info.Address),
			zap.Int64("serverID", info.Version))
		if err := s.cluster.UnRegister(node); err != nil {
			log.Warn("failed to deregister node", zap.Int64("id", node.NodeID), zap.String("address", node.Address), zap.Error(err))
			return err
		}
		s.metricsCacheManager.InvalidateSystemInfoMetrics()
	default:
		log.Warn("receive unknown service event type",
			zap.Any("type", event.EventType))
	}
	return nil
}

// startFlushLoop starts a goroutine to handle post func process
// which is to notify `RootCoord` that this segment is flushed
func (s *Server) startFlushLoop(ctx context.Context) {
	go func() {
		defer logutil.LogPanic()
		defer s.serverLoopWg.Done()
		ctx2, cancel := context.WithCancel(ctx)
		defer cancel()
		// send `Flushing` segments
		go s.handleFlushingSegments(ctx2)
		for {
			select {
			case <-ctx.Done():
				logutil.Logger(s.ctx).Info("flush loop shutdown")
				return
			case segmentID := <-s.flushCh:
				//Ignore return error
				log.Info("flush successfully", zap.Any("segmentID", segmentID))
				err := s.postFlush(ctx, segmentID)
				if err != nil {
					log.Warn("failed to do post flush", zap.Any("segmentID", segmentID), zap.Error(err))
				}
			}
		}
	}()
}

// post function after flush is done
// 1. check segment id is valid
// 2. notify RootCoord segment is flushed
// 3. change segment state to `Flushed` in meta
func (s *Server) postFlush(ctx context.Context, segmentID UniqueID) error {
	segment := s.meta.GetSegment(segmentID)
	if segment == nil {
		return errors.New("segment not found, might be a faked segemnt, ignore post flush")
	}
	// set segment to SegmentState_Flushed
	if err := s.meta.SetState(segmentID, commonpb.SegmentState_Flushed); err != nil {
		log.Error("flush segment complete failed", zap.Error(err))
		return err
	}
	log.Info("flush segment complete", zap.Int64("id", segmentID))
	return nil
}

// recovery logic, fetch all Segment in `Flushing` state and do Flush notification logic
func (s *Server) handleFlushingSegments(ctx context.Context) {
	segments := s.meta.GetFlushingSegments()
	for _, segment := range segments {
		select {
		case <-ctx.Done():
			return
		case s.flushCh <- segment.ID:
		}
	}
}

func (s *Server) initRootCoordClient() error {
	var err error
	if s.rootCoordClient, err = s.rootCoordClientCreator(s.ctx, Params.EtcdCfg.MetaRootPath, s.etcdCli); err != nil {
		return err
	}
	if err = s.rootCoordClient.Init(); err != nil {
		return err
	}
	return s.rootCoordClient.Start()
}

// Stop do the Server finalize processes
// it checks the server status is healthy, if not, just quit
// if Server is healthy, set server state to stopped, release etcd session,
//
//	stop message stream client and stop server loops
func (s *Server) Stop() error {
	if !s.stateCode.CompareAndSwap(commonpb.StateCode_Healthy, commonpb.StateCode_Abnormal) {
		return nil
	}
	logutil.Logger(s.ctx).Info("server shutdown")
	s.cluster.Close()
	s.garbageCollector.close()
	s.stopServerLoop()
	s.session.Revoke(time.Second)

	if Params.DataCoordCfg.EnableCompaction {
		s.stopCompactionTrigger()
		s.stopCompactionHandler()
	}
	return nil
}

// CleanMeta only for test
func (s *Server) CleanMeta() error {
	log.Debug("clean meta", zap.Any("kv", s.kvClient))
	return s.kvClient.RemoveWithPrefix("")
}

func (s *Server) stopServerLoop() {
	s.serverLoopCancel()
	s.serverLoopWg.Wait()
}

//func (s *Server) validateAllocRequest(collID UniqueID, partID UniqueID, channelName string) error {
//	if !s.meta.HasCollection(collID) {
//		return fmt.Errorf("can not find collection %d", collID)
//	}
//	if !s.meta.HasPartition(collID, partID) {
//		return fmt.Errorf("can not find partition %d", partID)
//	}
//	for _, name := range s.insertChannels {
//		if name == channelName {
//			return nil
//		}
//	}
//	return fmt.Errorf("can not find channel %s", channelName)
//}

// loadCollectionFromRootCoord communicates with RootCoord and asks for collection information.
// collection information will be added to server meta info.
func (s *Server) loadCollectionFromRootCoord(ctx context.Context, collectionID int64) error {
	resp, err := s.rootCoordClient.DescribeCollection(ctx, &milvuspb.DescribeCollectionRequest{
		Base: commonpbutil.NewMsgBase(
			commonpbutil.WithMsgType(commonpb.MsgType_DescribeCollection),
			commonpbutil.WithSourceID(paramtable.GetNodeID()),
		),
		DbName:       "",
		CollectionID: collectionID,
	})
	if err = VerifyResponse(resp, err); err != nil {
		return err
	}
	presp, err := s.rootCoordClient.ShowPartitions(ctx, &milvuspb.ShowPartitionsRequest{
		Base: commonpbutil.NewMsgBase(
			commonpbutil.WithMsgType(commonpb.MsgType_ShowPartitions),
			commonpbutil.WithMsgID(0),
			commonpbutil.WithTimeStamp(0),
			commonpbutil.WithSourceID(paramtable.GetNodeID()),
		),
		DbName:         "",
		CollectionName: resp.Schema.Name,
		CollectionID:   resp.CollectionID,
	})
	if err = VerifyResponse(presp, err); err != nil {
		log.Error("show partitions error", zap.String("collectionName", resp.Schema.Name),
			zap.Int64("collectionID", resp.CollectionID), zap.Error(err))
		return err
	}

	properties := make(map[string]string)
	for _, pair := range resp.Properties {
		properties[pair.GetKey()] = pair.GetValue()
	}

	collInfo := &collectionInfo{
		ID:             resp.CollectionID,
		Schema:         resp.Schema,
		Partitions:     presp.PartitionIDs,
		StartPositions: resp.GetStartPositions(),
		Properties:     properties,
	}
	s.meta.AddCollection(collInfo)
	return nil
}

func (s *Server) reCollectSegmentStats(ctx context.Context) {
	if s.channelManager == nil {
		log.Error("null channel manager found, which should NOT happen in non-testing environment")
		return
	}
	nodes := s.sessionManager.getLiveNodeIDs()
	log.Info("re-collecting segment stats from DataNodes",
		zap.Int64s("DataNode IDs", nodes))
	for _, node := range nodes {
		s.cluster.ReCollectSegmentStats(ctx, node)
	}
}
