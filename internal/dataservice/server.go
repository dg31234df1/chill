package dataservice

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/logutil"

	"github.com/golang/protobuf/proto"
	grpcdatanodeclient "github.com/zilliztech/milvus-distributed/internal/distributed/datanode/client"
	etcdkv "github.com/zilliztech/milvus-distributed/internal/kv/etcd"
	"github.com/zilliztech/milvus-distributed/internal/log"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/timesync"
	"github.com/zilliztech/milvus-distributed/internal/types"
	"github.com/zilliztech/milvus-distributed/internal/util/retry"
	"github.com/zilliztech/milvus-distributed/internal/util/trace"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"

	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
)

const role = "dataservice"

type (
	UniqueID  = typeutil.UniqueID
	Timestamp = typeutil.Timestamp
)
type Server struct {
	ctx               context.Context
	serverLoopCtx     context.Context
	serverLoopCancel  context.CancelFunc
	serverLoopWg      sync.WaitGroup
	state             atomic.Value
	client            *etcdkv.EtcdKV
	meta              *meta
	segAllocator      segmentAllocatorInterface
	statsHandler      *statsHandler
	ddHandler         *ddHandler
	allocator         allocatorInterface
	cluster           *dataNodeCluster
	msgProducer       *timesync.MsgProducer
	registerFinishCh  chan struct{}
	masterClient      types.MasterService
	ttMsgStream       msgstream.MsgStream
	k2sMsgStream      msgstream.MsgStream
	ddChannelName     string
	segmentInfoStream msgstream.MsgStream
	insertChannels    []string
	msFactory         msgstream.Factory
	ttBarrier         timesync.TimeTickBarrier
	allocMu           sync.RWMutex
}

func CreateServer(ctx context.Context, factory msgstream.Factory) (*Server, error) {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan struct{})
	s := &Server{
		ctx:              ctx,
		registerFinishCh: ch,
		cluster:          newDataNodeCluster(ch),
		msFactory:        factory,
	}
	s.insertChannels = s.getInsertChannels()
	s.UpdateStateCode(internalpb.StateCode_Abnormal)
	return s, nil
}

func (s *Server) getInsertChannels() []string {
	channels := make([]string, Params.InsertChannelNum)
	var i int64 = 0
	for ; i < Params.InsertChannelNum; i++ {
		channels[i] = Params.InsertChannelPrefixName + strconv.FormatInt(i, 10)
	}
	return channels
}

func (s *Server) SetMasterClient(masterClient types.MasterService) {
	s.masterClient = masterClient
}

func (s *Server) Init() error {
	return nil
}

func (s *Server) Start() error {
	var err error
	m := map[string]interface{}{
		"PulsarAddress":  Params.PulsarAddress,
		"ReceiveBufSize": 1024,
		"PulsarBufSize":  1024}
	err = s.msFactory.SetParams(m)
	if err != nil {
		return err
	}

	s.allocator = newAllocator(s.masterClient)
	if err = s.initMeta(); err != nil {
		return err
	}
	s.statsHandler = newStatsHandler(s.meta)
	s.segAllocator = newSegmentAllocator(s.meta, s.allocator)
	s.ddHandler = newDDHandler(s.meta, s.segAllocator)
	s.initSegmentInfoChannel()
	if err = s.loadMetaFromMaster(); err != nil {
		return err
	}
	s.waitDataNodeRegister()
	s.cluster.WatchInsertChannels(s.insertChannels)
	if err = s.initMsgProducer(); err != nil {
		return err
	}
	s.startServerLoop()
	s.UpdateStateCode(internalpb.StateCode_Healthy)
	log.Debug("start success")
	return nil
}

func (s *Server) UpdateStateCode(code internalpb.StateCode) {
	s.state.Store(code)
}

func (s *Server) checkStateIsHealthy() bool {
	return s.state.Load().(internalpb.StateCode) == internalpb.StateCode_Healthy
}

func (s *Server) initMeta() error {
	connectEtcdFn := func() error {
		etcdClient, err := clientv3.New(clientv3.Config{Endpoints: []string{Params.EtcdAddress}})
		if err != nil {
			return err
		}
		etcdKV := etcdkv.NewEtcdKV(etcdClient, Params.MetaRootPath)
		s.client = etcdKV
		s.meta, err = newMeta(etcdKV)
		if err != nil {
			return err
		}
		return nil
	}
	err := retry.Retry(100000, time.Millisecond*200, connectEtcdFn)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) initSegmentInfoChannel() {
	segmentInfoStream, _ := s.msFactory.NewMsgStream(s.ctx)
	segmentInfoStream.AsProducer([]string{Params.SegmentInfoChannelName})
	log.Debug("dataservice AsProducer: " + Params.SegmentInfoChannelName)
	s.segmentInfoStream = segmentInfoStream
	s.segmentInfoStream.Start()
}
func (s *Server) initMsgProducer() error {
	var err error
	if s.ttMsgStream, err = s.msFactory.NewMsgStream(s.ctx); err != nil {
		return err
	}
	s.ttMsgStream.AsConsumer([]string{Params.TimeTickChannelName}, Params.DataServiceSubscriptionName)
	log.Debug("dataservice AsConsumer: " + Params.TimeTickChannelName + " : " + Params.DataServiceSubscriptionName)
	s.ttMsgStream.Start()
	s.ttBarrier = timesync.NewHardTimeTickBarrier(s.ctx, s.ttMsgStream, s.cluster.GetNodeIDs())
	s.ttBarrier.Start()
	if s.k2sMsgStream, err = s.msFactory.NewMsgStream(s.ctx); err != nil {
		return err
	}
	s.k2sMsgStream.AsProducer(Params.K2SChannelNames)
	log.Debug("dataservice AsProducer: " + strings.Join(Params.K2SChannelNames, ", "))
	s.k2sMsgStream.Start()
	dataNodeTTWatcher := newDataNodeTimeTickWatcher(s.meta, s.segAllocator, s.cluster)
	k2sMsgWatcher := timesync.NewMsgTimeTickWatcher(s.k2sMsgStream)
	if s.msgProducer, err = timesync.NewTimeSyncMsgProducer(s.ttBarrier, dataNodeTTWatcher, k2sMsgWatcher); err != nil {
		return err
	}
	s.msgProducer.Start(s.ctx)
	return nil
}

func (s *Server) loadMetaFromMaster() error {
	ctx := context.Background()
	log.Debug("loading collection meta from master")
	var err error
	if err = s.checkMasterIsHealthy(); err != nil {
		return err
	}
	if s.ddChannelName == "" {
		channel, err := s.masterClient.GetDdChannel(ctx)
		if err != nil {
			return err
		}
		s.ddChannelName = channel.Value
	}
	collections, err := s.masterClient.ShowCollections(ctx, &milvuspb.ShowCollectionsRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_ShowCollections,
			MsgID:     -1, // todo add msg id
			Timestamp: 0,  // todo
			SourceID:  Params.NodeID,
		},
		DbName: "",
	})
	if err = VerifyResponse(collections, err); err != nil {
		return err
	}
	for _, collectionName := range collections.CollectionNames {
		collection, err := s.masterClient.DescribeCollection(ctx, &milvuspb.DescribeCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DescribeCollection,
				MsgID:     -1, // todo
				Timestamp: 0,  // todo
				SourceID:  Params.NodeID,
			},
			DbName:         "",
			CollectionName: collectionName,
		})
		if err = VerifyResponse(collection, err); err != nil {
			log.Error("describe collection error", zap.String("collectionName", collectionName), zap.Error(err))
			continue
		}
		partitions, err := s.masterClient.ShowPartitions(ctx, &milvuspb.ShowPartitionsRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_ShowPartitions,
				MsgID:     -1, // todo
				Timestamp: 0,  // todo
				SourceID:  Params.NodeID,
			},
			DbName:         "",
			CollectionName: collectionName,
			CollectionID:   collection.CollectionID,
		})
		if err = VerifyResponse(partitions, err); err != nil {
			log.Error("show partitions error", zap.String("collectionName", collectionName), zap.Int64("collectionID", collection.CollectionID), zap.Error(err))
			continue
		}
		err = s.meta.AddCollection(&collectionInfo{
			ID:         collection.CollectionID,
			Schema:     collection.Schema,
			Partitions: partitions.PartitionIDs,
		})
		if err != nil {
			log.Error("add collection to meta error", zap.Int64("collectionID", collection.CollectionID), zap.Error(err))
			continue
		}
	}
	log.Debug("load collection meta from master complete")
	return nil
}

func (s *Server) checkMasterIsHealthy() error {
	ticker := time.NewTicker(300 * time.Millisecond)
	ctx, cancel := context.WithTimeout(s.ctx, 30*time.Second)
	defer func() {
		ticker.Stop()
		cancel()
	}()
	for {
		var resp *internalpb.ComponentStates
		var err error
		select {
		case <-ctx.Done():
			return errors.New("master is not healthy")
		case <-ticker.C:
			resp, err = s.masterClient.GetComponentStates(ctx)
			if err = VerifyResponse(resp, err); err != nil {
				return err
			}
		}
		if resp.State.StateCode == internalpb.StateCode_Healthy {
			break
		}
	}
	return nil
}

func (s *Server) startServerLoop() {
	s.serverLoopCtx, s.serverLoopCancel = context.WithCancel(s.ctx)
	s.serverLoopWg.Add(3)
	go s.startStatsChannel(s.serverLoopCtx)
	go s.startSegmentFlushChannel(s.serverLoopCtx)
	go s.startProxyServiceTimeTickLoop(s.serverLoopCtx)
}

func (s *Server) startStatsChannel(ctx context.Context) {
	defer logutil.LogPanic()
	defer s.serverLoopWg.Done()
	statsStream, _ := s.msFactory.NewMsgStream(ctx)
	statsStream.AsConsumer([]string{Params.StatisticsChannelName}, Params.DataServiceSubscriptionName)
	log.Debug("dataservice AsConsumer: " + Params.StatisticsChannelName + " : " + Params.DataServiceSubscriptionName)
	statsStream.Start()
	defer statsStream.Close()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		msgPack := statsStream.Consume()
		for _, msg := range msgPack.Msgs {
			statistics, ok := msg.(*msgstream.SegmentStatisticsMsg)
			if !ok {
				log.Error("receive unknown type msg from stats channel", zap.Stringer("msgType", msg.Type()))
			}
			for _, stat := range statistics.SegStats {
				if err := s.statsHandler.HandleSegmentStat(stat); err != nil {
					log.Error("handle segment stat error", zap.Int64("segmentID", stat.SegmentID), zap.Error(err))
					continue
				}
			}
		}
	}
}

func (s *Server) startSegmentFlushChannel(ctx context.Context) {
	defer logutil.LogPanic()
	defer s.serverLoopWg.Done()
	flushStream, _ := s.msFactory.NewMsgStream(ctx)
	flushStream.AsConsumer([]string{Params.SegmentInfoChannelName}, Params.DataServiceSubscriptionName)
	log.Debug("dataservice AsConsumer: " + Params.SegmentInfoChannelName + " : " + Params.DataServiceSubscriptionName)
	flushStream.Start()
	defer flushStream.Close()
	for {
		select {
		case <-ctx.Done():
			log.Debug("segment flush channel shut down")
			return
		default:
		}
		msgPack := flushStream.Consume()
		for _, msg := range msgPack.Msgs {
			if msg.Type() != commonpb.MsgType_SegmentFlushDone {
				continue
			}
			realMsg := msg.(*msgstream.FlushCompletedMsg)
			err := s.meta.FlushSegment(realMsg.SegmentID, realMsg.BeginTimestamp)
			log.Debug("dataservice flushed segment", zap.Any("segmentID", realMsg.SegmentID), zap.Error(err))
			if err != nil {
				log.Error("get segment from meta error", zap.Int64("segmentID", realMsg.SegmentID), zap.Error(err))
				continue
			}
		}
	}
}

func (s *Server) startProxyServiceTimeTickLoop(ctx context.Context) {
	defer logutil.LogPanic()
	defer s.serverLoopWg.Done()
	flushStream, _ := s.msFactory.NewMsgStream(ctx)
	flushStream.AsConsumer([]string{Params.ProxyTimeTickChannelName}, Params.DataServiceSubscriptionName)
	flushStream.Start()
	defer flushStream.Close()
	for {
		select {
		case <-ctx.Done():
			log.Debug("Proxy service timetick loop shut down")
		default:
		}
		msgPack := flushStream.Consume()
		s.allocMu.Lock()
		for _, msg := range msgPack.Msgs {
			if msg.Type() != commonpb.MsgType_TimeTick {
				log.Warn("receive unknown msg from proxy service timetick", zap.Stringer("msgType", msg.Type()))
				continue
			}
			tMsg := msg.(*msgstream.TimeTickMsg)
			traceCtx := context.TODO()
			if err := s.segAllocator.ExpireAllocations(traceCtx, tMsg.Base.Timestamp); err != nil {
				log.Error("expire allocations error", zap.Error(err))
			}
		}
		s.allocMu.Unlock()
	}
}

func (s *Server) startDDChannel(ctx context.Context) {
	defer s.serverLoopWg.Done()
	ddStream, _ := s.msFactory.NewMsgStream(ctx)
	ddStream.AsConsumer([]string{s.ddChannelName}, Params.DataServiceSubscriptionName)
	log.Debug("dataservice AsConsumer: " + s.ddChannelName + " : " + Params.DataServiceSubscriptionName)
	ddStream.Start()
	defer ddStream.Close()
	for {
		select {
		case <-ctx.Done():
			log.Debug("dd channel shut down")
			return
		default:
		}
		msgPack := ddStream.Consume()
		for _, msg := range msgPack.Msgs {
			if err := s.ddHandler.HandleDDMsg(ctx, msg); err != nil {
				log.Error("handle dd msg error", zap.Error(err))
				continue
			}
		}
	}
}

func (s *Server) waitDataNodeRegister() {
	log.Debug("waiting data node to register")
	<-s.registerFinishCh
	log.Debug("all data nodes register")
}

func (s *Server) Stop() error {
	s.cluster.ShutDownClients()
	s.ttBarrier.Close()
	s.ttMsgStream.Close()
	s.k2sMsgStream.Close()
	s.msgProducer.Close()
	s.segmentInfoStream.Close()
	s.stopServerLoop()
	return nil
}

func (s *Server) stopServerLoop() {
	s.serverLoopCancel()
	s.serverLoopWg.Wait()
}

func (s *Server) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	resp := &internalpb.ComponentStates{
		State: &internalpb.ComponentInfo{
			NodeID:    Params.NodeID,
			Role:      role,
			StateCode: s.state.Load().(internalpb.StateCode),
		},
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
		},
	}
	dataNodeStates, err := s.cluster.GetDataNodeStates(ctx)
	if err != nil {
		resp.Status.Reason = err.Error()
		return resp, nil
	}
	resp.SubcomponentStates = dataNodeStates
	resp.Status.ErrorCode = commonpb.ErrorCode_Success
	return resp, nil
}

func (s *Server) GetTimeTickChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return &milvuspb.StringResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		Value: Params.TimeTickChannelName,
	}, nil
}

func (s *Server) GetStatisticsChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return &milvuspb.StringResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		Value: Params.StatisticsChannelName,
	}, nil
}

func (s *Server) RegisterNode(ctx context.Context, req *datapb.RegisterNodeRequest) (*datapb.RegisterNodeResponse, error) {
	ret := &datapb.RegisterNodeResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
		},
	}
	log.Debug("DataService: RegisterNode:", zap.String("IP", req.Address.Ip), zap.Int64("Port", req.Address.Port))
	node, err := s.newDataNode(req.Address.Ip, req.Address.Port, req.Base.SourceID)
	if err != nil {
		return nil, err
	}

	s.cluster.Register(node)

	if s.ddChannelName == "" {
		resp, err := s.masterClient.GetDdChannel(ctx)
		if err = VerifyResponse(resp, err); err != nil {
			ret.Status.Reason = err.Error()
			return ret, err
		}
		s.ddChannelName = resp.Value
	}
	ret.Status.ErrorCode = commonpb.ErrorCode_Success
	ret.InitParams = &internalpb.InitParams{
		NodeID: Params.NodeID,
		StartParams: []*commonpb.KeyValuePair{
			{Key: "DDChannelName", Value: s.ddChannelName},
			{Key: "SegmentStatisticsChannelName", Value: Params.StatisticsChannelName},
			{Key: "TimeTickChannelName", Value: Params.TimeTickChannelName},
			{Key: "CompleteFlushChannelName", Value: Params.SegmentInfoChannelName},
		},
	}
	return ret, nil
}

func (s *Server) newDataNode(ip string, port int64, id UniqueID) (*dataNode, error) {
	client := grpcdatanodeclient.NewClient(fmt.Sprintf("%s:%d", ip, port))
	if err := client.Init(); err != nil {
		return nil, err
	}

	if err := client.Start(); err != nil {
		return nil, err
	}
	return &dataNode{
		id: id,
		address: struct {
			ip   string
			port int64
		}{ip: ip, port: port},
		client:     client,
		channelNum: 0,
	}, nil
}

func (s *Server) Flush(ctx context.Context, req *datapb.FlushRequest) (*commonpb.Status, error) {
	s.allocMu.Lock()
	defer s.allocMu.Unlock()
	if !s.checkStateIsHealthy() {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
			Reason:    "server is initializing",
		}, nil
	}
	s.segAllocator.SealAllSegments(ctx, req.CollectionID)
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}, nil
}

func (s *Server) AssignSegmentID(ctx context.Context, req *datapb.AssignSegmentIDRequest) (*datapb.AssignSegmentIDResponse, error) {
	s.allocMu.Lock()
	defer s.allocMu.Unlock()
	resp := &datapb.AssignSegmentIDResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		SegIDAssignments: make([]*datapb.SegmentIDAssignment, 0),
	}
	if !s.checkStateIsHealthy() {
		resp.Status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		resp.Status.Reason = "server is initializing"
		return resp, nil
	}
	for _, r := range req.SegmentIDRequests {
		if !s.meta.HasCollection(r.CollectionID) {
			if err := s.loadCollectionFromMaster(ctx, r.CollectionID); err != nil {
				log.Error("load collection from master error", zap.Int64("collectionID", r.CollectionID), zap.Error(err))
				continue
			}
		}
		result := &datapb.SegmentIDAssignment{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UnexpectedError,
			},
		}
		segmentID, retCount, expireTs, err := s.segAllocator.AllocSegment(ctx, r.CollectionID, r.PartitionID, r.ChannelName, int(r.Count))
		if err != nil {
			if _, ok := err.(errRemainInSufficient); !ok {
				result.Status.Reason = fmt.Sprintf("allocation of Collection %d, Partition %d, Channel %s, Count %d error:  %s",
					r.CollectionID, r.PartitionID, r.ChannelName, r.Count, err.Error())
				resp.SegIDAssignments = append(resp.SegIDAssignments, result)
				continue
			}

			if err = s.openNewSegment(ctx, r.CollectionID, r.PartitionID, r.ChannelName); err != nil {
				result.Status.Reason = fmt.Sprintf("open new segment of Collection %d, Partition %d, Channel %s, Count %d error:  %s",
					r.CollectionID, r.PartitionID, r.ChannelName, r.Count, err.Error())
				resp.SegIDAssignments = append(resp.SegIDAssignments, result)
				continue
			}

			segmentID, retCount, expireTs, err = s.segAllocator.AllocSegment(ctx, r.CollectionID, r.PartitionID, r.ChannelName, int(r.Count))
			if err != nil {
				result.Status.Reason = fmt.Sprintf("retry allocation of Collection %d, Partition %d, Channel %s, Count %d error:  %s",
					r.CollectionID, r.PartitionID, r.ChannelName, r.Count, err.Error())
				resp.SegIDAssignments = append(resp.SegIDAssignments, result)
				continue
			}
		}

		result.Status.ErrorCode = commonpb.ErrorCode_Success
		result.CollectionID = r.CollectionID
		result.SegID = segmentID
		result.PartitionID = r.PartitionID
		result.Count = uint32(retCount)
		result.ExpireTime = expireTs
		result.ChannelName = r.ChannelName
		resp.SegIDAssignments = append(resp.SegIDAssignments, result)
	}
	return resp, nil
}

func (s *Server) loadCollectionFromMaster(ctx context.Context, collectionID int64) error {
	resp, err := s.masterClient.DescribeCollection(ctx, &milvuspb.DescribeCollectionRequest{
		Base: &commonpb.MsgBase{
			MsgType:  commonpb.MsgType_DescribeCollection,
			SourceID: Params.NodeID,
		},
		DbName:       "",
		CollectionID: collectionID,
	})
	if err = VerifyResponse(resp, err); err != nil {
		return err
	}
	collInfo := &collectionInfo{
		ID:     resp.CollectionID,
		Schema: resp.Schema,
	}
	return s.meta.AddCollection(collInfo)
}

func (s *Server) openNewSegment(ctx context.Context, collectionID UniqueID, partitionID UniqueID, channelName string) error {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	id, err := s.allocator.allocID()
	if err != nil {
		return err
	}
	segmentInfo, err := BuildSegment(collectionID, partitionID, id, channelName)
	if err != nil {
		return err
	}
	if err = s.meta.AddSegment(segmentInfo); err != nil {
		return err
	}
	if err = s.segAllocator.OpenSegment(ctx, segmentInfo); err != nil {
		return err
	}
	infoMsg := &msgstream.SegmentInfoMsg{
		BaseMsg: msgstream.BaseMsg{
			HashValues: []uint32{0},
		},
		SegmentMsg: datapb.SegmentMsg{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_SegmentInfo,
				MsgID:     0,
				Timestamp: 0,
				SourceID:  Params.NodeID,
			},
			Segment: segmentInfo,
		},
	}
	msgPack := msgstream.MsgPack{
		Msgs: []msgstream.TsMsg{infoMsg},
	}
	if err = s.segmentInfoStream.Produce(&msgPack); err != nil {
		return err
	}
	return nil
}

func (s *Server) ShowSegments(ctx context.Context, req *datapb.ShowSegmentsRequest) (*datapb.ShowSegmentsResponse, error) {
	resp := &datapb.ShowSegmentsResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
		},
	}
	if !s.checkStateIsHealthy() {
		resp.Status.Reason = "server is initializing"
		return resp, nil
	}
	ids := s.meta.GetSegmentsOfPartition(req.CollectionID, req.PartitionID)
	resp.Status.ErrorCode = commonpb.ErrorCode_Success
	resp.SegmentIDs = ids
	return resp, nil
}

func (s *Server) GetSegmentStates(ctx context.Context, req *datapb.GetSegmentStatesRequest) (*datapb.GetSegmentStatesResponse, error) {
	resp := &datapb.GetSegmentStatesResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
		},
	}
	if !s.checkStateIsHealthy() {
		resp.Status.Reason = "server is initializing"
		return resp, nil
	}

	for _, segmentID := range req.SegmentIDs {
		state := &datapb.SegmentStateInfo{
			Status:    &commonpb.Status{},
			SegmentID: segmentID,
		}
		segmentInfo, err := s.meta.GetSegment(segmentID)
		if err != nil {
			state.Status.ErrorCode = commonpb.ErrorCode_UnexpectedError
			state.Status.Reason = "get segment states error: " + err.Error()
		} else {
			state.Status.ErrorCode = commonpb.ErrorCode_Success
			state.State = segmentInfo.State
			state.CreateTime = segmentInfo.OpenTime
			state.SealedTime = segmentInfo.SealedTime
			state.FlushedTime = segmentInfo.FlushedTime
			state.StartPosition = segmentInfo.StartPosition
			state.EndPosition = segmentInfo.EndPosition
		}
		resp.States = append(resp.States, state)
	}
	resp.Status.ErrorCode = commonpb.ErrorCode_Success

	return resp, nil
}

func (s *Server) GetInsertBinlogPaths(ctx context.Context, req *datapb.GetInsertBinlogPathsRequest) (*datapb.GetInsertBinlogPathsResponse, error) {
	resp := &datapb.GetInsertBinlogPathsResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
		},
	}
	p := path.Join(Params.SegmentFlushMetaPath, strconv.FormatInt(req.SegmentID, 10))
	_, values, err := s.client.LoadWithPrefix(p)
	if err != nil {
		resp.Status.Reason = err.Error()
		return resp, nil
	}
	m := make(map[int64][]string)
	tMeta := &datapb.SegmentFieldBinlogMeta{}
	for _, v := range values {
		if err := proto.UnmarshalText(v, tMeta); err != nil {
			resp.Status.Reason = err.Error()
			return resp, nil
		}
		m[tMeta.FieldID] = append(m[tMeta.FieldID], tMeta.BinlogPath)
	}

	fids := make([]UniqueID, len(m))
	paths := make([]*internalpb.StringList, len(m))
	for k, v := range m {
		fids = append(fids, k)
		paths = append(paths, &internalpb.StringList{Values: v})
	}
	resp.Status.ErrorCode = commonpb.ErrorCode_Success
	resp.FieldIDs = fids
	resp.Paths = paths
	return resp, nil
}

func (s *Server) GetInsertChannels(ctx context.Context, req *datapb.GetInsertChannelsRequest) (*internalpb.StringList, error) {
	return &internalpb.StringList{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		Values: s.insertChannels,
	}, nil
}

func (s *Server) GetCollectionStatistics(ctx context.Context, req *datapb.GetCollectionStatisticsRequest) (*datapb.GetCollectionStatisticsResponse, error) {
	resp := &datapb.GetCollectionStatisticsResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
		},
	}
	nums, err := s.meta.GetNumRowsOfCollection(req.CollectionID)
	if err != nil {
		resp.Status.Reason = err.Error()
		return resp, nil
	}
	resp.Status.ErrorCode = commonpb.ErrorCode_Success
	resp.Stats = append(resp.Stats, &commonpb.KeyValuePair{Key: "row_count", Value: strconv.FormatInt(nums, 10)})
	return resp, nil
}

func (s *Server) GetPartitionStatistics(ctx context.Context, req *datapb.GetPartitionStatisticsRequest) (*datapb.GetPartitionStatisticsResponse, error) {
	// todo implement
	return nil, nil
}

func (s *Server) GetSegmentInfoChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return &milvuspb.StringResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		Value: Params.SegmentInfoChannelName,
	}, nil
}

func (s *Server) GetSegmentInfo(ctx context.Context, req *datapb.GetSegmentInfoRequest) (*datapb.GetSegmentInfoResponse, error) {
	resp := &datapb.GetSegmentInfoResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
		},
	}
	if !s.checkStateIsHealthy() {
		resp.Status.Reason = "data service is not healthy"
		return resp, nil
	}
	infos := make([]*datapb.SegmentInfo, len(req.SegmentIDs))
	for i, id := range req.SegmentIDs {
		segmentInfo, err := s.meta.GetSegment(id)
		if err != nil {
			resp.Status.Reason = err.Error()
			return resp, nil
		}
		infos[i] = segmentInfo
	}
	resp.Status.ErrorCode = commonpb.ErrorCode_Success
	resp.Infos = infos
	return resp, nil
}
