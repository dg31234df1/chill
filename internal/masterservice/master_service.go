package masterservice

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/errors"
	etcdkv "github.com/zilliztech/milvus-distributed/internal/kv/etcd"
	ms "github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/msgstream/pulsarms"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/indexpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/proto/masterpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	"github.com/zilliztech/milvus-distributed/internal/proto/proxypb"
	"github.com/zilliztech/milvus-distributed/internal/proto/querypb"
	"github.com/zilliztech/milvus-distributed/internal/util/tsoutil"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
	"go.etcd.io/etcd/clientv3"
)

//  internalpb2 -> internalpb
//  proxypb(proxy_service)
//  querypb(query_service)
//  datapb(data_service)
//  indexpb(index_service)
//  milvuspb -> milvuspb
//  masterpb2 -> masterpb (master_service)

type ProxyServiceInterface interface {
	GetTimeTickChannel() (*milvuspb.StringResponse, error)
	InvalidateCollectionMetaCache(request *proxypb.InvalidateCollMetaCacheRequest) (*commonpb.Status, error)
}

type DataServiceInterface interface {
	GetInsertBinlogPaths(req *datapb.InsertBinlogPathRequest) (*datapb.InsertBinlogPathsResponse, error)
	GetSegmentInfoChannel() (*milvuspb.StringResponse, error)
}

type IndexServiceInterface interface {
	BuildIndex(req *indexpb.BuildIndexRequest) (*indexpb.BuildIndexResponse, error)
}

type QueryServiceInterface interface {
	ReleaseCollection(req *querypb.ReleaseCollectionRequest) (*commonpb.Status, error)
}

type Interface interface {
	//service
	Init() error
	Start() error
	Stop() error
	GetComponentStates() (*internalpb2.ComponentStates, error)

	//DDL request
	CreateCollection(in *milvuspb.CreateCollectionRequest) (*commonpb.Status, error)
	DropCollection(in *milvuspb.DropCollectionRequest) (*commonpb.Status, error)
	HasCollection(in *milvuspb.HasCollectionRequest) (*milvuspb.BoolResponse, error)
	DescribeCollection(in *milvuspb.DescribeCollectionRequest) (*milvuspb.DescribeCollectionResponse, error)
	ShowCollections(in *milvuspb.ShowCollectionRequest) (*milvuspb.ShowCollectionResponse, error)
	CreatePartition(in *milvuspb.CreatePartitionRequest) (*commonpb.Status, error)
	DropPartition(in *milvuspb.DropPartitionRequest) (*commonpb.Status, error)
	HasPartition(in *milvuspb.HasPartitionRequest) (*milvuspb.BoolResponse, error)
	ShowPartitions(in *milvuspb.ShowPartitionRequest) (*milvuspb.ShowPartitionResponse, error)

	//index builder service
	CreateIndex(in *milvuspb.CreateIndexRequest) (*commonpb.Status, error)
	DescribeIndex(in *milvuspb.DescribeIndexRequest) (*milvuspb.DescribeIndexResponse, error)

	//global timestamp allocator
	AllocTimestamp(in *masterpb.TsoRequest) (*masterpb.TsoResponse, error)
	AllocID(in *masterpb.IDRequest) (*masterpb.IDResponse, error)

	//TODO, master load these channel form config file ?

	//receiver time tick from proxy service, and put it into this channel
	GetTimeTickChannel() (string, error)

	//receive ddl from rpc and time tick from proxy service, and put them into this channel
	GetDdChannel() (string, error)

	//just define a channel, not used currently
	GetStatisticsChannel() (string, error)

	//segment
	DescribeSegment(in *milvuspb.DescribeSegmentRequest) (*milvuspb.DescribeSegmentResponse, error)
	ShowSegments(in *milvuspb.ShowSegmentRequest) (*milvuspb.ShowSegmentResponse, error)

	//get system config from master, not used currently
	//GetSysConfigs(in *milvuspb.SysConfigRequest)

	//GetIndexState(ctx context.Context, request *milvuspb.IndexStateRequest) (*milvuspb.IndexStateResponse, error)
}

// ------------------ struct -----------------------

// master core
type Core struct {
	/*
		ProxyServiceClient Interface:
		get proxy service time tick channel,InvalidateCollectionMetaCache

		DataService Interface:
		Segment States Channel, from DataService, if create new segment, data service should put the segment id into this channel, and let the master add the segment id to the collection meta
		Segment Flush Watcher, monitor if segment has flushed into disk

		IndexService Interface:
		indexBuilder Sch, tell index service to build index
	*/

	MetaTable *metaTable
	//id allocator
	idAllocator *GlobalIDAllocator
	//tso allocator
	tsoAllocator *GlobalTSOAllocator

	//inner members
	ctx     context.Context
	cancel  context.CancelFunc
	etcdCli *clientv3.Client
	kvBase  *etcdkv.EtcdKV
	metaKV  *etcdkv.EtcdKV

	//setMsgStreams, receive time tick from proxy service time tick channel
	ProxyTimeTickChan chan typeutil.Timestamp

	//setMsgStreams, send time tick into dd channel and time tick channel
	SendTimeTick func(t typeutil.Timestamp) error

	//setMsgStreams, send create collection into dd channel
	DdCreateCollectionReq func(req *internalpb2.CreateCollectionRequest) error

	//setMsgStreams, send drop collection into dd channel, and notify the proxy to delete this collection
	DdDropCollectionReq func(req *internalpb2.DropCollectionRequest) error

	//setMsgStreams, send create partition into dd channel
	DdCreatePartitionReq func(req *internalpb2.CreatePartitionRequest) error

	//setMsgStreams, send drop partition into dd channel
	DdDropPartitionReq func(req *internalpb2.DropPartitionRequest) error

	//setMsgStreams segment channel, receive segment info from data service, if master create segment
	DataServiceSegmentChan chan *datapb.SegmentInfo

	//setMsgStreams ,if segment flush completed, data node would put segment id into msg stream
	DataNodeSegmentFlushCompletedChan chan typeutil.UniqueID

	//get binlog file path from data service,
	GetBinlogFilePathsFromDataServiceReq func(segID typeutil.UniqueID, fieldID typeutil.UniqueID) ([]string, error)

	//call index builder's client to build index, return build id
	BuildIndexReq func(binlog []string, typeParams []*commonpb.KeyValuePair, indexParams []*commonpb.KeyValuePair, indexID typeutil.UniqueID, indexName string) (typeutil.UniqueID, error)

	//proxy service interface, notify proxy service to drop collection
	InvalidateCollectionMetaCache func(ts typeutil.Timestamp, dbName string, collectionName string) error

	//query service interface, notify query service to release collection
	ReleaseCollection func(ts typeutil.Timestamp, dbID typeutil.UniqueID, collectionID typeutil.UniqueID) error

	// put create index task into this chan
	indexTaskQueue chan *CreateIndexTask

	//dd request scheduler
	ddReqQueue      chan reqTask //dd request will be push into this chan
	lastDdTimeStamp typeutil.Timestamp

	//time tick loop
	lastTimeTick typeutil.Timestamp

	//states code
	stateCode atomic.Value

	//call once
	initOnce  sync.Once
	startOnce sync.Once
	isInit    atomic.Value
}

// --------------------- function --------------------------

func NewCore(c context.Context) (*Core, error) {
	ctx, cancel := context.WithCancel(c)
	rand.Seed(time.Now().UnixNano())
	Params.Init()
	core := &Core{
		ctx:    ctx,
		cancel: cancel,
	}
	core.stateCode.Store(internalpb2.StateCode_INITIALIZING)
	core.isInit.Store(false)
	return core, nil
}

func (c *Core) checkInit() error {
	if c.MetaTable == nil {
		return errors.Errorf("MetaTable is nil")
	}
	if c.idAllocator == nil {
		return errors.Errorf("idAllocator is nil")
	}
	if c.tsoAllocator == nil {
		return errors.Errorf("tsoAllocator is nil")
	}
	if c.etcdCli == nil {
		return errors.Errorf("etcdCli is nil")
	}
	if c.metaKV == nil {
		return errors.Errorf("metaKV is nil")
	}
	if c.kvBase == nil {
		return errors.Errorf("kvBase is nil")
	}
	if c.ProxyTimeTickChan == nil {
		return errors.Errorf("ProxyTimeTickChan is nil")
	}
	if c.ddReqQueue == nil {
		return errors.Errorf("ddReqQueue is nil")
	}
	if c.DdCreateCollectionReq == nil {
		return errors.Errorf("DdCreateCollectionReq is nil")
	}
	if c.DdDropCollectionReq == nil {
		return errors.Errorf("DdDropCollectionReq is nil")
	}
	if c.DdCreatePartitionReq == nil {
		return errors.Errorf("DdCreatePartitionReq is nil")
	}
	if c.DdDropPartitionReq == nil {
		return errors.Errorf("DdDropPartitionReq is nil")
	}
	if c.DataServiceSegmentChan == nil {
		return errors.Errorf("DataServiceSegmentChan is nil")
	}
	if c.GetBinlogFilePathsFromDataServiceReq == nil {
		return errors.Errorf("GetBinlogFilePathsFromDataServiceReq is nil")
	}
	if c.BuildIndexReq == nil {
		return errors.Errorf("BuildIndexReq is nil")
	}
	if c.InvalidateCollectionMetaCache == nil {
		return errors.Errorf("InvalidateCollectionMetaCache is nil")
	}
	if c.indexTaskQueue == nil {
		return errors.Errorf("indexTaskQueue is nil")
	}
	if c.DataNodeSegmentFlushCompletedChan == nil {
		return errors.Errorf("DataNodeSegmentFlushCompletedChan is nil")
	}
	if c.ReleaseCollection == nil {
		return errors.Errorf("ReleaseCollection is nil")
	}

	log.Printf("master node id = %d", Params.NodeID)
	log.Printf("master dd channel name = %s", Params.DdChannel)
	log.Printf("master time ticke channel name = %s", Params.TimeTickChannel)
	return nil
}

func (c *Core) startDdScheduler() {
	for {
		select {
		case <-c.ctx.Done():
			log.Printf("close dd scheduler, exit task execution loop")
			return
		case task, ok := <-c.ddReqQueue:
			if !ok {
				log.Printf("dd chan is closed, exit task execution loopo")
				return
			}
			ts, err := task.Ts()
			if err != nil {
				task.Notify(err)
				break
			}
			if !task.IgnoreTimeStamp() && ts <= c.lastDdTimeStamp {
				task.Notify(errors.Errorf("input timestamp = %d, last dd time stamp = %d", ts, c.lastDdTimeStamp))
				break
			}
			err = task.Execute()
			task.Notify(err)
			if ts > c.lastDdTimeStamp {
				c.lastDdTimeStamp = ts
			}
		}
	}
}

func (c *Core) startTimeTickLoop() {
	for {
		select {
		case <-c.ctx.Done():
			log.Printf("close master time tick loop")
			return
		case tt, ok := <-c.ProxyTimeTickChan:
			if !ok {
				log.Printf("proxyTimeTickStream is closed, exit time tick loop")
				return
			}
			if tt <= c.lastTimeTick {
				log.Printf("master time tick go back, last time tick = %d, input time tick = %d", c.lastTimeTick, tt)
			}
			if err := c.SendTimeTick(tt); err != nil {
				log.Printf("master send time tick into dd and time_tick channel failed: %s", err.Error())
			}
			c.lastTimeTick = tt
		}
	}
}

//data service send segment info to master when create segment
func (c *Core) startDataServiceSegmentLoop() {
	for {
		select {
		case <-c.ctx.Done():
			log.Printf("close data service segment loop")
			return
		case seg, ok := <-c.DataServiceSegmentChan:
			if !ok {
				log.Printf("data service segment is closed, exit loop")
				return
			}
			if seg == nil {
				log.Printf("segment from data service is nill")
			} else if err := c.MetaTable.AddSegment(seg); err != nil {
				//what if master add segment failed, but data service success?
				log.Printf("add segment info meta table failed ")
			} else {
				log.Printf("add segment, collection id = %d, partition id = %d, segment id = %d", seg.CollectionID, seg.PartitionID, seg.SegmentID)
			}
		}
	}
}

//create index loop
func (c *Core) startCreateIndexLoop() {
	for {
		select {
		case <-c.ctx.Done():
			log.Printf("close create index loop")
			return
		case t, ok := <-c.indexTaskQueue:
			if !ok {
				log.Printf("index task chan has closed, exit loop")
				return
			}
			if err := t.BuildIndex(); err != nil {
				log.Printf("create index failed, error = %s", err.Error())
			} else {
				log.Printf("create index,index name = %s, field name = %s, segment id = %d", t.indexName, t.fieldSchema.Name, t.segmentID)
			}
		}
	}
}

func (c *Core) startSegmentFlushCompletedLoop() {
	for {
		select {
		case <-c.ctx.Done():
			log.Printf("close segment flush completed loop")
			return
		case seg, ok := <-c.DataNodeSegmentFlushCompletedChan:
			if !ok {
				log.Printf("data node segment flush completed chan has colsed, exit loop")
			}
			coll, err := c.MetaTable.GetCollectionBySegmentID(seg)
			if err != nil {
				log.Printf("GetCollectionBySegmentID, error = %s ", err.Error())
			}
			for i, f := range coll.IndexParams {
				fieldSch, err := GetFieldSchemaByID(coll, f.FiledID)
				if err == nil {
					t := &CreateIndexTask{
						core:        c,
						segmentID:   seg,
						indexName:   fmt.Sprintf("index_field_%d_%d", f.FiledID, i),
						fieldSchema: fieldSch,
						indexParams: nil,
					}
					c.indexTaskQueue <- t
				}
			}
		}
	}
}

func (c *Core) tsLoop() {
	tsoTicker := time.NewTicker(UpdateTimestampStep)
	defer tsoTicker.Stop()
	ctx, cancel := context.WithCancel(c.ctx)
	defer cancel()
	for {
		select {
		case <-tsoTicker.C:
			if err := c.tsoAllocator.UpdateTSO(); err != nil {
				log.Println("failed to update timestamp", err)
				return
			}
			if err := c.idAllocator.UpdateID(); err != nil {
				log.Println("failed to update id", err)
				return
			}
		case <-ctx.Done():
			// Server is closed and it should return nil.
			log.Println("tsLoop is closed")
			return
		}
	}
}
func (c *Core) setMsgStreams() error {

	if Params.PulsarAddress == "" {
		return errors.Errorf("PulsarAddress is empty")
	}
	if Params.MsgChannelSubName == "" {
		return errors.Errorf("MsgChannelSubName is emptyr")
	}

	//proxy time tick stream,
	if Params.ProxyTimeTickChannel == "" {
		return errors.Errorf("ProxyTimeTickChannel is empty")
	}

	factory := pulsarms.NewFactory(Params.PulsarAddress, 1024, 1024)
	proxyTimeTickStream, _ := factory.NewMsgStream(c.ctx)
	proxyTimeTickStream.AsConsumer([]string{Params.ProxyTimeTickChannel}, Params.MsgChannelSubName)
	proxyTimeTickStream.Start()

	// master time tick channel
	if Params.TimeTickChannel == "" {
		return errors.Errorf("TimeTickChannel is empty")
	}
	timeTickStream, _ := factory.NewMsgStream(c.ctx)
	timeTickStream.AsProducer([]string{Params.TimeTickChannel})

	// master dd channel
	if Params.DdChannel == "" {
		return errors.Errorf("DdChannel is empty")
	}
	ddStream, _ := factory.NewMsgStream(c.ctx)
	ddStream.AsProducer([]string{Params.DdChannel})

	c.SendTimeTick = func(t typeutil.Timestamp) error {
		msgPack := ms.MsgPack{}
		baseMsg := ms.BaseMsg{
			MsgCtx:         nil,
			BeginTimestamp: t,
			EndTimestamp:   t,
			HashValues:     []uint32{0},
		}
		timeTickResult := internalpb2.TimeTickMsg{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kTimeTick,
				MsgID:     0,
				Timestamp: t,
				SourceID:  int64(Params.NodeID),
			},
		}
		timeTickMsg := &ms.TimeTickMsg{
			BaseMsg:     baseMsg,
			TimeTickMsg: timeTickResult,
		}
		msgPack.Msgs = append(msgPack.Msgs, timeTickMsg)
		if err := timeTickStream.Broadcast(&msgPack); err != nil {
			return err
		}
		if err := ddStream.Broadcast(&msgPack); err != nil {
			return err
		}
		return nil
	}

	c.DdCreateCollectionReq = func(req *internalpb2.CreateCollectionRequest) error {
		msgPack := ms.MsgPack{}
		baseMsg := ms.BaseMsg{
			BeginTimestamp: req.Base.Timestamp,
			EndTimestamp:   req.Base.Timestamp,
			HashValues:     []uint32{0},
		}
		collMsg := &ms.CreateCollectionMsg{
			BaseMsg:                 baseMsg,
			CreateCollectionRequest: *req,
		}
		msgPack.Msgs = append(msgPack.Msgs, collMsg)
		if err := ddStream.Broadcast(&msgPack); err != nil {
			return err
		}
		return nil
	}

	c.DdDropCollectionReq = func(req *internalpb2.DropCollectionRequest) error {
		msgPack := ms.MsgPack{}
		baseMsg := ms.BaseMsg{
			BeginTimestamp: req.Base.Timestamp,
			EndTimestamp:   req.Base.Timestamp,
			HashValues:     []uint32{0},
		}
		collMsg := &ms.DropCollectionMsg{
			BaseMsg:               baseMsg,
			DropCollectionRequest: *req,
		}
		msgPack.Msgs = append(msgPack.Msgs, collMsg)
		if err := ddStream.Broadcast(&msgPack); err != nil {
			return err
		}
		return nil
	}

	c.DdCreatePartitionReq = func(req *internalpb2.CreatePartitionRequest) error {
		msgPack := ms.MsgPack{}
		baseMsg := ms.BaseMsg{
			BeginTimestamp: req.Base.Timestamp,
			EndTimestamp:   req.Base.Timestamp,
			HashValues:     []uint32{0},
		}
		collMsg := &ms.CreatePartitionMsg{
			BaseMsg:                baseMsg,
			CreatePartitionRequest: *req,
		}
		msgPack.Msgs = append(msgPack.Msgs, collMsg)
		if err := ddStream.Broadcast(&msgPack); err != nil {
			return err
		}
		return nil
	}

	c.DdDropPartitionReq = func(req *internalpb2.DropPartitionRequest) error {
		msgPack := ms.MsgPack{}
		baseMsg := ms.BaseMsg{
			BeginTimestamp: req.Base.Timestamp,
			EndTimestamp:   req.Base.Timestamp,
			HashValues:     []uint32{0},
		}
		collMsg := &ms.DropPartitionMsg{
			BaseMsg:              baseMsg,
			DropPartitionRequest: *req,
		}
		msgPack.Msgs = append(msgPack.Msgs, collMsg)
		if err := ddStream.Broadcast(&msgPack); err != nil {
			return err
		}
		return nil
	}

	// receive time tick from msg stream
	c.ProxyTimeTickChan = make(chan typeutil.Timestamp, 1024)
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			case ttmsgs, ok := <-proxyTimeTickStream.Chan():
				if !ok {
					log.Printf("proxy time tick msg stream closed")
					return
				}
				if len(ttmsgs.Msgs) > 0 {
					for _, ttm := range ttmsgs.Msgs {
						ttmsg, ok := ttm.(*ms.TimeTickMsg)
						if !ok {
							continue
						}
						c.ProxyTimeTickChan <- ttmsg.Base.Timestamp
					}
				}
			}
		}
	}()

	//segment channel, data service create segment,or data node flush segment will put msg in this channel
	if Params.DataServiceSegmentChannel == "" {
		return errors.Errorf("DataServiceSegmentChannel is empty")
	}
	dataServiceStream, _ := factory.NewMsgStream(c.ctx)
	dataServiceStream.AsConsumer([]string{Params.DataServiceSegmentChannel}, Params.MsgChannelSubName)
	dataServiceStream.Start()
	c.DataServiceSegmentChan = make(chan *datapb.SegmentInfo, 1024)
	c.DataNodeSegmentFlushCompletedChan = make(chan typeutil.UniqueID, 1024)

	// receive segment info from msg stream
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			case segMsg, ok := <-dataServiceStream.Chan():
				if !ok {
					log.Printf("data service segment msg closed")
				}
				if len(segMsg.Msgs) > 0 {
					for _, segm := range segMsg.Msgs {
						segInfoMsg, ok := segm.(*ms.SegmentInfoMsg)
						if ok {
							c.DataServiceSegmentChan <- segInfoMsg.Segment
						} else {
							flushMsg, ok := segm.(*ms.FlushCompletedMsg)
							if ok {
								c.DataNodeSegmentFlushCompletedChan <- flushMsg.SegmentFlushCompletedMsg.SegmentID
							} else {
								log.Printf("receive unexpected msg from data service stream, value = %v", segm)
							}
						}
					}
				}
			}
		}
	}()

	return nil
}

func (c *Core) SetProxyService(s ProxyServiceInterface) error {
	rsp, err := s.GetTimeTickChannel()
	if err != nil {
		return err
	}
	Params.ProxyTimeTickChannel = rsp.Value
	log.Printf("proxy time tick channel name = %s", Params.ProxyTimeTickChannel)

	c.InvalidateCollectionMetaCache = func(ts typeutil.Timestamp, dbName string, collectionName string) error {
		status, _ := s.InvalidateCollectionMetaCache(&proxypb.InvalidateCollMetaCacheRequest{
			Base: &commonpb.MsgBase{
				MsgType:   0, //TODO,MsgType
				MsgID:     0,
				Timestamp: ts,
				SourceID:  int64(Params.NodeID),
			},
			DbName:         dbName,
			CollectionName: collectionName,
		})
		if status == nil {
			return errors.New("invalidate collection metacache resp is nil")
		}
		if status.ErrorCode != commonpb.ErrorCode_SUCCESS {
			return errors.New(status.Reason)
		}
		return nil
	}
	return nil
}

func (c *Core) SetDataService(s DataServiceInterface) error {
	rsp, err := s.GetSegmentInfoChannel()
	if err != nil {
		return err
	}
	Params.DataServiceSegmentChannel = rsp.Value
	log.Printf("data service segment channel name = %s", Params.DataServiceSegmentChannel)

	c.GetBinlogFilePathsFromDataServiceReq = func(segID typeutil.UniqueID, fieldID typeutil.UniqueID) ([]string, error) {
		ts, err := c.tsoAllocator.Alloc(1)
		if err != nil {
			return nil, err
		}
		binlog, err := s.GetInsertBinlogPaths(&datapb.InsertBinlogPathRequest{
			Base: &commonpb.MsgBase{
				MsgType:   0, //TODO, msy type
				MsgID:     0,
				Timestamp: ts,
				SourceID:  int64(Params.NodeID),
			},
			SegmentID: segID,
		})
		if err != nil {
			return nil, err
		}
		if binlog.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
			return nil, errors.Errorf("GetInsertBinlogPaths from data service failed, error = %s", binlog.Status.Reason)
		}
		for i := range binlog.FieldIDs {
			if binlog.FieldIDs[i] == fieldID {
				return binlog.Paths[i].Values, nil
			}
		}
		return nil, errors.Errorf("binlog file not exist, segment id = %d, field id = %d", segID, fieldID)
	}
	return nil
}

func (c *Core) SetIndexService(s IndexServiceInterface) error {
	c.BuildIndexReq = func(binlog []string, typeParams []*commonpb.KeyValuePair, indexParams []*commonpb.KeyValuePair, indexID typeutil.UniqueID, indexName string) (typeutil.UniqueID, error) {
		rsp, err := s.BuildIndex(&indexpb.BuildIndexRequest{
			DataPaths:   binlog,
			TypeParams:  typeParams,
			IndexParams: indexParams,
			IndexID:     indexID,
			IndexName:   indexName,
		})
		if err != nil {
			return 0, err
		}
		if rsp.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
			return 0, errors.Errorf("BuildIndex from index service failed, error = %s", rsp.Status.Reason)
		}
		return rsp.IndexBuildID, nil
	}
	return nil
}

func (c *Core) SetQueryService(s QueryServiceInterface) error {
	c.ReleaseCollection = func(ts typeutil.Timestamp, dbID typeutil.UniqueID, collectionID typeutil.UniqueID) error {
		req := &querypb.ReleaseCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kReleaseCollection,
				MsgID:     0, //TODO, msg ID
				Timestamp: ts,
				SourceID:  int64(Params.NodeID),
			},
			DbID:         dbID,
			CollectionID: collectionID,
		}
		rsp, err := s.ReleaseCollection(req)
		if err != nil {
			return err
		}
		if rsp.ErrorCode != commonpb.ErrorCode_SUCCESS {
			return errors.Errorf("ReleaseCollection from query service failed, error = %s", rsp.Reason)
		}
		return nil
	}
	return nil
}

func (c *Core) Init() error {
	var initError error = nil
	c.initOnce.Do(func() {
		if c.etcdCli, initError = clientv3.New(clientv3.Config{Endpoints: []string{Params.EtcdAddress}, DialTimeout: 5 * time.Second}); initError != nil {
			return
		}
		c.metaKV = etcdkv.NewEtcdKV(c.etcdCli, Params.MetaRootPath)
		if c.MetaTable, initError = NewMetaTable(c.metaKV); initError != nil {
			return
		}

		c.kvBase = etcdkv.NewEtcdKV(c.etcdCli, Params.KvRootPath)

		c.idAllocator = NewGlobalIDAllocator("idTimestamp", tsoutil.NewTSOKVBase([]string{Params.EtcdAddress}, Params.KvRootPath, "gid"))
		if initError = c.idAllocator.Initialize(); initError != nil {
			return
		}
		c.tsoAllocator = NewGlobalTSOAllocator("timestamp", tsoutil.NewTSOKVBase([]string{Params.EtcdAddress}, Params.KvRootPath, "tso"))
		if initError = c.tsoAllocator.Initialize(); initError != nil {
			return
		}
		c.ddReqQueue = make(chan reqTask, 1024)
		c.indexTaskQueue = make(chan *CreateIndexTask, 1024)
		initError = c.setMsgStreams()
		c.isInit.Store(true)
	})
	if initError == nil {
		log.Printf("Master service State Code = %s", internalpb2.StateCode_name[int32(internalpb2.StateCode_INITIALIZING)])
	}
	return initError
}

func (c *Core) Start() error {
	isInit := c.isInit.Load().(bool)
	if !isInit {
		return errors.Errorf("call init before start")
	}
	if err := c.checkInit(); err != nil {
		return err
	}
	c.startOnce.Do(func() {
		go c.startDdScheduler()
		go c.startTimeTickLoop()
		go c.startDataServiceSegmentLoop()
		go c.startCreateIndexLoop()
		go c.startSegmentFlushCompletedLoop()
		go c.tsLoop()
		c.stateCode.Store(internalpb2.StateCode_HEALTHY)
	})
	log.Printf("Master service State Code = %s", internalpb2.StateCode_name[int32(internalpb2.StateCode_HEALTHY)])
	return nil
}

func (c *Core) Stop() error {
	c.cancel()
	c.stateCode.Store(internalpb2.StateCode_ABNORMAL)
	return nil
}

func (c *Core) GetComponentStates() (*internalpb2.ComponentStates, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	log.Printf("GetComponentStates : %s", internalpb2.StateCode_name[int32(code)])

	return &internalpb2.ComponentStates{
		State: &internalpb2.ComponentInfo{
			NodeID:    int64(Params.NodeID),
			Role:      typeutil.MasterServiceRole,
			StateCode: code,
			ExtraInfo: nil,
		},
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		SubcomponentStates: []*internalpb2.ComponentInfo{
			{
				NodeID:    int64(Params.NodeID),
				Role:      typeutil.MasterServiceRole,
				StateCode: code,
				ExtraInfo: nil,
			},
		},
	}, nil
}

func (c *Core) GetTimeTickChannel() (string, error) {
	return Params.TimeTickChannel, nil
}

func (c *Core) GetDdChannel() (string, error) {
	return Params.DdChannel, nil
}

func (c *Core) GetStatisticsChannel() (string, error) {
	return Params.StatisticsChannel, nil
}

func (c *Core) CreateCollection(in *milvuspb.CreateCollectionRequest) (*commonpb.Status, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
		}, nil
	}
	log.Printf("CreateCollection : %s", in.CollectionName)
	t := &CreateCollectionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    "Create collection failed: " + err.Error(),
		}, nil
	}
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}, nil
}

func (c *Core) DropCollection(in *milvuspb.DropCollectionRequest) (*commonpb.Status, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
		}, nil
	}
	log.Printf("DropCollection : %s", in.CollectionName)
	t := &DropCollectionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    "Create collection failed: " + err.Error(),
		}, nil
	}
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}, nil
}

func (c *Core) HasCollection(in *milvuspb.HasCollectionRequest) (*milvuspb.BoolResponse, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &milvuspb.BoolResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
			},
			Value: false,
		}, nil
	}
	log.Printf("HasCollection : %s", in.CollectionName)
	t := &HasCollectionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req:           in,
		HasCollection: false,
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &milvuspb.BoolResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "Has collection failed: " + err.Error(),
			},
			Value: false,
		}, nil
	}
	return &milvuspb.BoolResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		Value: t.HasCollection,
	}, nil
}

func (c *Core) DescribeCollection(in *milvuspb.DescribeCollectionRequest) (*milvuspb.DescribeCollectionResponse, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &milvuspb.DescribeCollectionResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
			},
			Schema:       nil,
			CollectionID: 0,
		}, nil
	}
	log.Printf("DescribeCollection : %s", in.CollectionName)
	t := &DescribeCollectionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
		Rsp: &milvuspb.DescribeCollectionResponse{},
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &milvuspb.DescribeCollectionResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "describe collection failed: " + err.Error(),
			},
			Schema: nil,
		}, nil
	}
	t.Rsp.Status = &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}
	return t.Rsp, nil
}

func (c *Core) ShowCollections(in *milvuspb.ShowCollectionRequest) (*milvuspb.ShowCollectionResponse, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &milvuspb.ShowCollectionResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
			},
			CollectionNames: nil,
		}, nil
	}
	log.Printf("ShowCollections : %s", in.DbName)
	t := &ShowCollectionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
		Rsp: &milvuspb.ShowCollectionResponse{
			CollectionNames: nil,
		},
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &milvuspb.ShowCollectionResponse{
			CollectionNames: nil,
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "ShowCollections failed: " + err.Error(),
			},
		}, nil
	}
	t.Rsp.Status = &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}
	return t.Rsp, nil
}

func (c *Core) CreatePartition(in *milvuspb.CreatePartitionRequest) (*commonpb.Status, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
		}, nil
	}
	log.Printf("CreatePartition : %s - %s", in.CollectionName, in.PartitionName)
	t := &CreatePartitionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    "create partition failed: " + err.Error(),
		}, nil
	}
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}, nil
}

func (c *Core) DropPartition(in *milvuspb.DropPartitionRequest) (*commonpb.Status, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
		}, nil
	}
	log.Printf("DropPartition : %s - %s", in.CollectionName, in.PartitionName)
	t := &DropPartitionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    "DropPartition failed: " + err.Error(),
		}, nil
	}
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}, nil
}

func (c *Core) HasPartition(in *milvuspb.HasPartitionRequest) (*milvuspb.BoolResponse, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &milvuspb.BoolResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
			},
			Value: false,
		}, nil
	}
	log.Printf("HasPartition : %s - %s", in.CollectionName, in.PartitionName)
	t := &HasPartitionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req:          in,
		HasPartition: false,
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &milvuspb.BoolResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "HasPartition failed: " + err.Error(),
			},
			Value: false,
		}, nil
	}
	return &milvuspb.BoolResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		Value: t.HasPartition,
	}, nil
}

func (c *Core) ShowPartitions(in *milvuspb.ShowPartitionRequest) (*milvuspb.ShowPartitionResponse, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &milvuspb.ShowPartitionResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
			},
			PartitionNames: nil,
			PartitionIDs:   nil,
		}, nil
	}
	log.Printf("ShowPartitions : %s", in.CollectionName)
	t := &ShowPartitionReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
		Rsp: &milvuspb.ShowPartitionResponse{
			PartitionNames: nil,
			Status:         nil,
		},
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &milvuspb.ShowPartitionResponse{
			PartitionNames: nil,
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "ShowPartitions failed: " + err.Error(),
			},
		}, nil
	}
	t.Rsp.Status = &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}
	return t.Rsp, nil
}

func (c *Core) CreateIndex(in *milvuspb.CreateIndexRequest) (*commonpb.Status, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
		}, nil
	}
	log.Printf("CreateIndex : %s - %s ", in.CollectionName, in.FieldName)
	t := &CreateIndexReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    "CreateIndex failed, error = " + err.Error(),
		}, nil
	}
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}, nil
}

func (c *Core) DescribeIndex(in *milvuspb.DescribeIndexRequest) (*milvuspb.DescribeIndexResponse, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &milvuspb.DescribeIndexResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
			},
			IndexDescriptions: nil,
		}, nil
	}
	log.Printf("DescribeIndex : %s - %s", in.CollectionName, in.FieldName)
	t := &DescribeIndexReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
		Rsp: &milvuspb.DescribeIndexResponse{
			Status:            nil,
			IndexDescriptions: nil,
		},
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &milvuspb.DescribeIndexResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "DescribeIndex failed, error = " + err.Error(),
			},
			IndexDescriptions: nil,
		}, nil
	}
	t.Rsp.Status = &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}
	return t.Rsp, nil
}

func (c *Core) DescribeSegment(in *milvuspb.DescribeSegmentRequest) (*milvuspb.DescribeSegmentResponse, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &milvuspb.DescribeSegmentResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
			},
			IndexID: 0,
		}, nil
	}
	log.Printf("DescribeSegment : %d - %d", in.CollectionID, in.SegmentID)
	t := &DescribeSegmentReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
		Rsp: &milvuspb.DescribeSegmentResponse{
			Status:  nil,
			IndexID: 0,
		},
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &milvuspb.DescribeSegmentResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "DescribeSegment failed, error = " + err.Error(),
			},
			IndexID: 0,
		}, nil
	}
	t.Rsp.Status = &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}
	return t.Rsp, nil
}

func (c *Core) ShowSegments(in *milvuspb.ShowSegmentRequest) (*milvuspb.ShowSegmentResponse, error) {
	code := c.stateCode.Load().(internalpb2.StateCode)
	if code != internalpb2.StateCode_HEALTHY {
		return &milvuspb.ShowSegmentResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("state code = %s", internalpb2.StateCode_name[int32(code)]),
			},
			SegmentIDs: nil,
		}, nil
	}
	log.Printf("ShowSegments : %d - %d", in.CollectionID, in.PartitionID)
	t := &ShowSegmentReqTask{
		baseReqTask: baseReqTask{
			cv:   make(chan error),
			core: c,
		},
		Req: in,
		Rsp: &milvuspb.ShowSegmentResponse{
			Status:     nil,
			SegmentIDs: nil,
		},
	}
	c.ddReqQueue <- t
	err := t.WaitToFinish()
	if err != nil {
		return &milvuspb.ShowSegmentResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "ShowSegments failed, error: " + err.Error(),
			},
			SegmentIDs: nil,
		}, nil
	}
	t.Rsp.Status = &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}
	return t.Rsp, nil
}

func (c *Core) AllocTimestamp(in *masterpb.TsoRequest) (*masterpb.TsoResponse, error) {
	ts, err := c.tsoAllocator.Alloc(in.Count)
	if err != nil {
		return &masterpb.TsoResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "AllocTimestamp failed: " + err.Error(),
			},
			Timestamp: 0,
			Count:     0,
		}, nil
	}
	// log.Printf("AllocTimestamp : %d", ts)
	return &masterpb.TsoResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		Timestamp: ts,
		Count:     in.Count,
	}, nil
}

func (c *Core) AllocID(in *masterpb.IDRequest) (*masterpb.IDResponse, error) {
	start, _, err := c.idAllocator.Alloc(in.Count)
	if err != nil {
		return &masterpb.IDResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "AllocID failed: " + err.Error(),
			},
			ID:    0,
			Count: in.Count,
		}, nil
	}
	log.Printf("AllocID : %d", start)
	return &masterpb.IDResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		ID:    start,
		Count: in.Count,
	}, nil
}
