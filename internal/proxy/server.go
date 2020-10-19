package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	mpb "github.com/zilliztech/milvus-distributed/internal/proto/master"
	pb "github.com/zilliztech/milvus-distributed/internal/proto/message"
	"github.com/zilliztech/milvus-distributed/internal/master/collection"
	"github.com/golang/protobuf/proto"
	etcd "go.etcd.io/etcd/clientv3"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
	"time"
)

const (
	keyCollectionPath = "collection"
	keySegmentPath    = "segment"
)

type proxyServer struct {
	pb.UnimplementedMilvusServiceServer
	address       string
	masterAddress string
	rootPath      string   // etcd root path
	pulsarAddr    string   // pulsar address for reader
	readerTopics  []string //reader topics
	deleteTopic   string
	queryTopic    string
	resultTopic   string
	resultGroup   string
	numReaderNode int
	proxyId       int64
	getTimestamp  func(count uint32) ([]Timestamp, pb.Status)
	client        *etcd.Client
	ctx           context.Context
	wg            sync.WaitGroup
	////////////////////////////////////////////////////////////////
	masterConn   *grpc.ClientConn
	masterClient mpb.MasterClient
	grpcServer   *grpc.Server
	reqSch       *requestScheduler
	///////////////////////////////////////////////////////////////
	collectionList   map[uint64]*mpb.Collection
	nameCollectionId map[string]uint64
	segmentList      map[uint64]*mpb.Segment
	collectionMux    sync.Mutex
	queryId          atomic.Uint64
}

func (s *proxyServer) CreateCollection(ctx context.Context, req *pb.Mapping) (*pb.Status, error) {
	log.Printf("create collection %s", req.CollectionName)
	return s.masterClient.CreateCollection(ctx, req)
}

func (s *proxyServer) CountCollection(ctx context.Context, req *pb.CollectionName) (*pb.CollectionRowCount, error) {
	s.collectionMux.Lock()
	defer s.collectionMux.Unlock()

	collection_id, ok := s.nameCollectionId[req.CollectionName]
	if !ok {
		return &pb.CollectionRowCount{
			CollectionRowCount: 0,
			Status: &pb.Status{
				ErrorCode: pb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("unable to get collection %s", req.CollectionName),
			},
		}, nil
	}
	if info, ok := s.collectionList[collection_id]; ok {
		count := int64(0)
		for _, seg_id := range info.SegmentIds {
			if seg, ok := s.segmentList[seg_id]; ok {
				count += seg.Rows
			}
		}
		return &pb.CollectionRowCount{
			CollectionRowCount: count,
			Status: &pb.Status{
				ErrorCode: pb.ErrorCode_SUCCESS,
			},
		}, nil
	}
	return &pb.CollectionRowCount{
		CollectionRowCount: 0,
		Status: &pb.Status{
			ErrorCode: pb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    fmt.Sprintf("unable to get collection %s", req.CollectionName),
		},
	}, nil
}

func (s *proxyServer) CreateIndex(ctx context.Context, req *pb.IndexParam) (*pb.Status, error) {
	log.Printf("create index, collection name = %s, index name = %s, filed_name = %s", req.CollectionName, req.IndexName, req.FieldName)
	return s.masterClient.CreateIndex(ctx, req)
}

func (s *proxyServer) DeleteByID(ctx context.Context, req *pb.DeleteByIDParam) (*pb.Status, error) {
	log.Printf("delete entites, total = %d", len(req.IdArray))
	pm := &manipulationReq{
		ManipulationReqMsg: pb.ManipulationReqMsg{
			CollectionName: req.CollectionName,
			ReqType:        pb.ReqType_kDeleteEntityByID,
			ProxyId:        s.proxyId,
		},
		proxy: s,
	}
	for _, id := range req.IdArray {
		pm.PrimaryKeys = append(pm.PrimaryKeys, uint64(id))
	}
	if len(pm.PrimaryKeys) > 1 {
		if st := pm.PreExecute(); st.ErrorCode != pb.ErrorCode_SUCCESS {
			return &st, nil
		}
		if st := pm.Execute(); st.ErrorCode != pb.ErrorCode_SUCCESS {
			return &st, nil
		}
		if st := pm.PostExecute(); st.ErrorCode != pb.ErrorCode_SUCCESS {
			return &st, nil
		}
		if st := pm.WaitToFinish(); st.ErrorCode != pb.ErrorCode_SUCCESS {
			return &st, nil
		}
	}
	return &pb.Status{ErrorCode: pb.ErrorCode_SUCCESS}, nil
}

func (s *proxyServer) Insert(ctx context.Context, req *pb.InsertParam) (*pb.EntityIds, error) {
	log.Printf("Insert Entities, total =  %d", len(req.RowsData))
	ipm := make(map[uint32]*manipulationReq)

	//TODO
	if len(req.EntityIdArray) == 0 { //primary key is empty, set primary key by server
		log.Printf("Set primary key")
	}
	if len(req.EntityIdArray) != len(req.RowsData) {
		return &pb.EntityIds{
			Status: &pb.Status{
				ErrorCode: pb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    fmt.Sprintf("length of EntityIdArray not equal to lenght of RowsData"),
			},
			EntityIdArray: req.EntityIdArray,
		}, nil
	}

	for i := 0; i < len(req.EntityIdArray); i++ {
		key := uint64(req.EntityIdArray[i])
		hash, err := Hash32_Uint64(key)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "hash failed on %d", key)
		}
		hash = hash % uint32(len(s.readerTopics))
		ip, ok := ipm[hash]
		if !ok {
			segId, err := s.getSegmentId(int32(hash), req.CollectionName)
			if err != nil {
				return nil, err
			}
			ipm[hash] = &manipulationReq{
				ManipulationReqMsg: pb.ManipulationReqMsg{
					CollectionName: req.CollectionName,
					PartitionTag:   req.PartitionTag,
					SegmentId:      segId,
					ChannelId:      uint64(hash),
					ReqType:        pb.ReqType_kInsert,
					ProxyId:        s.proxyId,
					ExtraParams:    req.ExtraParams,
				},
				proxy: s,
			}
			ip = ipm[hash]
		}
		ip.PrimaryKeys = append(ip.PrimaryKeys, key)
		ip.RowsData = append(ip.RowsData, req.RowsData[i])
	}
	for _, ip := range ipm {
		if st := ip.PreExecute(); st.ErrorCode != pb.ErrorCode_SUCCESS { //do nothing
			return &pb.EntityIds{
				Status:        &st,
				EntityIdArray: req.EntityIdArray,
			}, nil
		}
		if st := ip.Execute(); st.ErrorCode != pb.ErrorCode_SUCCESS { // push into chan
			return &pb.EntityIds{
				Status:        &st,
				EntityIdArray: req.EntityIdArray,
			}, nil
		}
		if st := ip.PostExecute(); st.ErrorCode != pb.ErrorCode_SUCCESS { //post to pulsar
			return &pb.EntityIds{
				Status:        &st,
				EntityIdArray: req.EntityIdArray,
			}, nil
		}
	}
	for _, ip := range ipm {
		if st := ip.WaitToFinish(); st.ErrorCode != pb.ErrorCode_SUCCESS {
			log.Printf("Wait to finish failed, error code = %d", st.ErrorCode)
		}
	}

	return &pb.EntityIds{
		Status: &pb.Status{
			ErrorCode: pb.ErrorCode_SUCCESS,
		},
		EntityIdArray: req.EntityIdArray,
	}, nil
}

func (s *proxyServer) Search(ctx context.Context, req *pb.SearchParam) (*pb.QueryResult, error) {
	qm := &queryReq{
		QueryReqMsg: pb.QueryReqMsg{
			CollectionName: req.CollectionName,
			VectorParam:    req.VectorParam,
			PartitionTags:  req.PartitionTag,
			Dsl:            req.Dsl,
			ExtraParams:    req.ExtraParams,
			ProxyId:        s.proxyId,
			QueryId:        s.queryId.Add(1),
			ReqType:        pb.ReqType_kSearch,
		},
		proxy: s,
	}
	log.Printf("search on collection %s, proxy id = %d, query id = %d", req.CollectionName, qm.ProxyId, qm.QueryId)
	if st := qm.PreExecute(); st.ErrorCode != pb.ErrorCode_SUCCESS {
		return &pb.QueryResult{
			Status:  &st,
			QueryId: qm.QueryId,
			ProxyId: qm.ProxyId,
		}, nil
	}
	if st := qm.Execute(); st.ErrorCode != pb.ErrorCode_SUCCESS {
		return &pb.QueryResult{
			Status:  &st,
			QueryId: qm.QueryId,
			ProxyId: qm.ProxyId,
		}, nil
	}
	if st := qm.PostExecute(); st.ErrorCode != pb.ErrorCode_SUCCESS {
		return &pb.QueryResult{
			Status:  &st,
			QueryId: qm.QueryId,
			ProxyId: qm.ProxyId,
		}, nil
	}
	if st := qm.WaitToFinish(); st.ErrorCode != pb.ErrorCode_SUCCESS {
		return &pb.QueryResult{
			Status:  &st,
			QueryId: qm.QueryId,
			ProxyId: qm.ProxyId,
		}, nil
	}
	return s.reduceResult(qm), nil
}

//check if proxySerer is set correct
func (s *proxyServer) check() error {
	if len(s.address) == 0 {
		return fmt.Errorf("proxy address is unset")
	}
	if len(s.masterAddress) == 0 {
		return fmt.Errorf("master address is unset")
	}
	if len(s.rootPath) == 0 {
		return fmt.Errorf("root path for etcd is unset")
	}
	if len(s.pulsarAddr) == 0 {
		return fmt.Errorf("pulsar address is unset")
	}
	if len(s.readerTopics) == 0 {
		return fmt.Errorf("reader topics is unset")
	}
	if len(s.deleteTopic) == 0 {
		return fmt.Errorf("delete topic is unset")
	}
	if len(s.queryTopic) == 0 {
		return fmt.Errorf("query topic is unset")
	}
	if len(s.resultTopic) == 0 {
		return fmt.Errorf("result topic is unset")
	}
	if len(s.resultGroup) == 0 {
		return fmt.Errorf("result group is unset")
	}
	if s.numReaderNode <= 0 {
		return fmt.Errorf("number of reader nodes is unset")
	}
	if s.proxyId <= 0 {
		return fmt.Errorf("proxyId is unset")
	}
	log.Printf("proxy id = %d", s.proxyId)
	if s.getTimestamp == nil {
		return fmt.Errorf("getTimestamp is unset")
	}
	if s.client == nil {
		return fmt.Errorf("etcd client is unset")
	}
	if s.ctx == nil {
		return fmt.Errorf("context is unset")
	}
	return nil
}

func (s *proxyServer) getSegmentId(channelId int32, colName string) (uint64, error) {
	s.collectionMux.Lock()
	defer s.collectionMux.Unlock()
	colId, ok := s.nameCollectionId[colName]
	if !ok {
		return 0, status.Errorf(codes.Unknown, "can't get collection id of %s", colName)
	}
	colInfo, ok := s.collectionList[colId]
	if !ok {
		return 0, status.Errorf(codes.Unknown, "can't get collection, name = %s, id = %d", colName, colId)
	}
	for _, segId := range colInfo.SegmentIds {
		seg, ok := s.segmentList[segId]
		if !ok {
			return 0, status.Errorf(codes.Unknown, "can't get segment of %d", segId)
		}
		if seg.Status == mpb.SegmentStatus_OPENED {
			if seg.ChannelStart <= channelId && channelId < seg.ChannelEnd {
				return segId, nil
			}
		}
	}
	return 0, status.Errorf(codes.Unknown, "can't get segment id, channel id = %d", channelId)
}

func (s *proxyServer) connectMaster() error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	conn, err := grpc.DialContext(ctx, s.masterAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Connect to master failed, error= %v", err)
		return err
	}
	log.Printf("Connected to master, master_addr=%s", s.masterAddress)
	s.masterConn = conn
	s.masterClient = mpb.NewMasterClient(conn)
	return nil
}

func (s *proxyServer) Close() {
	s.client.Close()
	s.masterConn.Close()
	s.grpcServer.Stop()
}

func (s *proxyServer) StartGrpcServer() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	go func() {
		s.wg.Add(1)
		defer s.wg.Done()
		server := grpc.NewServer()
		pb.RegisterMilvusServiceServer(server, s)
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf("Proxy grpc server fatal error=%v", err)
		}
	}()
	return nil
}

func (s *proxyServer) WatchEtcd() error {
	s.collectionMux.Lock()
	defer s.collectionMux.Unlock()

	cos, err := s.client.Get(s.ctx, s.rootPath+"/"+keyCollectionPath, etcd.WithPrefix())
	if err != nil {
		return err
	}
	for _, cob := range cos.Kvs {
		// TODO: simplify collection struct
		var co mpb.Collection
		var mco collection.Collection
		if err := json.Unmarshal(cob.Value, &mco); err != nil {
			return err
		}
		proto.UnmarshalText(mco.GrpcMarshalString, &co)
		s.nameCollectionId[co.Name] = co.Id
		s.collectionList[co.Id] = &co
		log.Printf("watch collection, name = %s, id = %d", co.Name, co.Id)
	}
	segs, err := s.client.Get(s.ctx, s.rootPath+"/"+keySegmentPath, etcd.WithPrefix())
	if err != nil {
		return err
	}
	for _, segb := range segs.Kvs {
		var seg mpb.Segment
		if err := json.Unmarshal(segb.Value, &seg); err != nil {
			return err
		}
		s.segmentList[seg.SegmentId] = &seg
		log.Printf("watch segment id = %d\n", seg.SegmentId)
	}

	cow := s.client.Watch(s.ctx, s.rootPath+"/"+keyCollectionPath, etcd.WithPrefix(), etcd.WithRev(cos.Header.Revision+1))
	segw := s.client.Watch(s.ctx, s.rootPath+"/"+keySegmentPath, etcd.WithPrefix(), etcd.WithRev(segs.Header.Revision+1))
	go func() {
		s.wg.Add(1)
		defer s.wg.Done()
		for {
			select {
			case <-s.ctx.Done():
				return
			case coe := <-cow:
				func() {
					s.collectionMux.Lock()
					defer s.collectionMux.Unlock()
					for _, e := range coe.Events {
						var co mpb.Collection
						var mco collection.Collection
						if err := json.Unmarshal(e.Kv.Value, &mco); err != nil {
							log.Printf("unmarshal Collection failed, error = %v", err)
						} else {
							proto.UnmarshalText(mco.GrpcMarshalString, &co)
							s.nameCollectionId[co.Name] = co.Id
							s.collectionList[co.Id] = &co
							log.Printf("watch collection, name = %s, id = %d", co.Name, co.Id)
						}
					}
				}()
			case sege := <-segw:
				func() {
					s.collectionMux.Lock()
					defer s.collectionMux.Unlock()
					for _, e := range sege.Events {
						var seg mpb.Segment
						if err := json.Unmarshal(e.Kv.Value, &seg); err != nil {
							log.Printf("unmarshal Segment failed, error = %v", err)
						} else {
							s.segmentList[seg.SegmentId] = &seg
							log.Printf("watch segment id = %d\n", seg.SegmentId)
						}
					}
				}()
			}
		}
	}()

	return nil
}

func startProxyServer(srv *proxyServer) error {
	if err := srv.check(); err != nil {
		return err
	}
	srv.reqSch = &requestScheduler{}
	if err := srv.restartManipulationRoutine(1024); err != nil {
		return err
	}
	if err := srv.restartQueryRoutine(1024); err != nil {
		return err
	}

	srv.nameCollectionId = make(map[string]uint64)
	srv.collectionList = make(map[uint64]*mpb.Collection)
	srv.segmentList = make(map[uint64]*mpb.Segment)

	if err := srv.connectMaster(); err != nil {
		return err
	}
	if err := srv.WatchEtcd(); err != nil {
		return err
	}

	srv.queryId.Store(uint64(time.Now().UnixNano()))

	return srv.StartGrpcServer()
}
