package indexservice

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/allocator"
	"github.com/zilliztech/milvus-distributed/internal/errors"
	"github.com/zilliztech/milvus-distributed/internal/kv"
	etcdkv "github.com/zilliztech/milvus-distributed/internal/kv/etcd"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/indexpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

type IndexService struct {
	// implement Service

	//nodeClients [] .Interface
	// factory method
	loopCtx    context.Context
	loopCancel func()
	loopWg     sync.WaitGroup

	grpcServer *grpc.Server
	nodeNum    int64

	idAllocator *allocator.IDAllocator

	kv kv.Base

	metaTable *metaTable
	// Add callback functions at different stages
	startCallbacks []func()
	closeCallbacks []func()
}

type UniqueID = typeutil.UniqueID
type Timestamp = typeutil.Timestamp

func (i *IndexService) Init() {
	panic("implement me")
}

func (i *IndexService) Start() {
	panic("implement me")
}

func (i *IndexService) Stop() {
	panic("implement me")
}

func (i *IndexService) GetComponentStates() (*internalpb2.ComponentStates, error) {
	panic("implement me")
}

func (i *IndexService) GetTimeTickChannel() (string, error) {
	panic("implement me")
}

func (i *IndexService) GetStatisticsChannel() (string, error) {
	panic("implement me")
}

func (i *IndexService) RegisterNode(req *indexpb.RegisterNodeRequest) (*indexpb.RegisterNodeResponse, error) {
	nodeID := i.nodeNum + 1

	//TODO: update meta table
	_, ok := i.metaTable.nodeID2Address[nodeID]
	if ok {
		log.Fatalf("Register IndexNode fatal, IndexNode has already exists with nodeID=%d", nodeID)
	}

	log.Println("this is register indexNode func")
	i.metaTable.nodeID2Address[nodeID] = req.Address

	//TODO: register index node params?
	var params []*commonpb.KeyValuePair
	minioAddress, err := Params.Load("minio.address")
	if err != nil {
		return nil, err
	}
	minioPort, err := Params.Load("minio.port")
	if err != nil {
		return nil, err
	}
	params = append(params, &commonpb.KeyValuePair{Key: "minio.address", Value: minioAddress})
	params = append(params, &commonpb.KeyValuePair{Key: "minio.port", Value: minioPort})
	params = append(params, &commonpb.KeyValuePair{Key: "minio.accessKeyID", Value: Params.MinIOAccessKeyID})
	params = append(params, &commonpb.KeyValuePair{Key: "minio.secretAccessKey", Value: Params.MinIOSecretAccessKey})
	params = append(params, &commonpb.KeyValuePair{Key: "minio.useSSL", Value: strconv.FormatBool(Params.MinIOUseSSL)})
	params = append(params, &commonpb.KeyValuePair{Key: "minio.bucketName", Value: Params.MinioBucketName})

	i.nodeNum++

	return &indexpb.RegisterNodeResponse{
		InitParams: &internalpb2.InitParams{
			NodeID:      nodeID,
			StartParams: params,
		},
	}, nil
}

func (i *IndexService) BuildIndex(req *indexpb.BuildIndexRequest) (*indexpb.BuildIndexResponse, error) {
	//TODO: Multiple indexes will build at same time.
	//ctx := context.Background()
	//indexNodeClient := indexnode.NewIndexNode(ctx, rand.Int63n(i.nodeNum))
	//
	////TODO: Allocator index ID
	//indexID := int64(0)
	//
	//request := &indexpb.BuildIndexCmd{
	//	IndexID: indexID,
	//	Req:     req,
	//}
	//
	//status, err := indexNodeClient.BuildIndex(request)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &indexpb.BuildIndexResponse{
	//	Status:  status,
	//	IndexID: indexID,
	//}, nil
	return nil, nil
}

func (i *IndexService) GetIndexStates(req *indexpb.IndexStatesRequest) (*indexpb.IndexStatesResponse, error) {

	ret := &indexpb.IndexStatesResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		State:   commonpb.IndexState_FINISHED,
		IndexID: req.IndexID,
	}

	meta, ok := i.metaTable.indexID2Meta[req.IndexID]
	if !ok {
		ret.Status = &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_BUILD_INDEX_ERROR,
			Reason:    "index does not exists with ID = " + strconv.FormatInt(req.IndexID, 10),
		}
		ret.State = commonpb.IndexState_NONE

		return ret, errors.Errorf("index already exists with ID = " + strconv.FormatInt(req.IndexID, 10))
	}

	ret.State = meta.State
	ret.IndexID = meta.IndexID

	return ret, nil
}

func (i *IndexService) GetIndexFilePaths(req *indexpb.IndexFilePathRequest) (*indexpb.IndexFilePathsResponse, error) {
	panic("implement me")
}

func (i *IndexService) NotifyBuildIndex(nty *indexpb.BuildIndexNotification) (*commonpb.Status, error) {
	//TODO: Multiple indexes are building successfully at same time.
	meta, ok := i.metaTable.indexID2Meta[nty.IndexID]
	if !ok {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_BUILD_INDEX_ERROR,
			Reason:    "index already exists with ID = " + strconv.FormatInt(nty.IndexID, 10),
		}, errors.Errorf("index already exists with ID = " + strconv.FormatInt(nty.IndexID, 10))
	}

	meta.State = commonpb.IndexState_FINISHED
	meta.IndexFilePaths = nty.IndexFilePaths
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
		Reason:    "",
	}, nil
}

func NewIndexServiceImpl(ctx context.Context) *IndexService {
	Params.Init()
	ctx1, cancel := context.WithCancel(ctx)
	s := &IndexService{
		loopCtx:    ctx1,
		loopCancel: cancel,
	}

	connectEtcdFn := func() error {
		etcdAddress := Params.EtcdAddress
		etcdClient, err := clientv3.New(clientv3.Config{Endpoints: []string{etcdAddress}})
		if err != nil {
			return err
		}
		etcdKV := etcdkv.NewEtcdKV(etcdClient, Params.MetaRootPath)
		metakv, err := NewMetaTable(etcdKV)
		if err != nil {
			return err
		}
		s.metaTable = metakv
		return nil
	}
	err := Retry(10, time.Millisecond*200, connectEtcdFn)
	if err != nil {
		return nil
	}

	s.nodeNum = 0
	return s
}
