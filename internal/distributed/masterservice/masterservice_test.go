package masterservice

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	cms "github.com/zilliztech/milvus-distributed/internal/masterservice"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	"github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
)

func TestGrpcService(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	randVal := rand.Int()

	cms.Params.Address = "127.0.0.1"
	cms.Params.Port = (randVal % 100) + 10000
	cms.Params.NodeID = 0

	cms.Params.PulsarAddress = "pulsar://127.0.0.1:6650"
	cms.Params.EtcdAddress = "127.0.0.1:2379"
	cms.Params.MetaRootPath = fmt.Sprintf("/%d/test/meta", randVal)
	cms.Params.KvRootPath = fmt.Sprintf("/%d/test/kv", randVal)
	cms.Params.ProxyTimeTickChannel = fmt.Sprintf("proxyTimeTick%d", randVal)
	cms.Params.MsgChannelSubName = fmt.Sprintf("msgChannel%d", randVal)
	cms.Params.TimeTickChannel = fmt.Sprintf("timeTick%d", randVal)
	cms.Params.DdChannel = fmt.Sprintf("ddChannel%d", randVal)
	cms.Params.StatisticsChannel = fmt.Sprintf("stateChannel%d", randVal)

	cms.Params.MaxPartitionNum = 64
	cms.Params.DefaultPartitionName = "_default"

	t.Logf("master service port = %d", cms.Params.Port)

	svr, err := NewGrpcServer()
	assert.Nil(t, err)

	core := svr.core.(*cms.Core)

	core.ProxyTimeTickChan = make(chan typeutil.Timestamp, 8)

	timeTickArray := make([]typeutil.Timestamp, 0, 16)
	core.SendTimeTick = func(ts typeutil.Timestamp) error {
		t.Logf("send time tick %d", ts)
		timeTickArray = append(timeTickArray, ts)
		return nil
	}
	createCollectionArray := make([]*internalpb2.CreateCollectionRequest, 0, 16)
	core.DdCreateCollectionReq = func(req *internalpb2.CreateCollectionRequest) error {
		t.Logf("Create Colllection %s", req.CollectionName)
		createCollectionArray = append(createCollectionArray, req)
		return nil
	}

	dropCollectionArray := make([]*internalpb2.DropCollectionRequest, 0, 16)
	core.DdDropCollectionReq = func(req *internalpb2.DropCollectionRequest) error {
		t.Logf("Drop Collection %s", req.CollectionName)
		dropCollectionArray = append(dropCollectionArray, req)
		return nil
	}

	createPartitionArray := make([]*internalpb2.CreatePartitionRequest, 0, 16)
	core.DdCreatePartitionReq = func(req *internalpb2.CreatePartitionRequest) error {
		t.Logf("Create Partition %s", req.PartitionName)
		createPartitionArray = append(createPartitionArray, req)
		return nil
	}

	dropPartitionArray := make([]*internalpb2.DropPartitionRequest, 0, 16)
	core.DdDropPartitionReq = func(req *internalpb2.DropPartitionRequest) error {
		t.Logf("Drop Partition %s", req.PartitionName)
		dropPartitionArray = append(dropPartitionArray, req)
		return nil
	}

	err = svr.Init(&cms.InitParams{ProxyTimeTickChannel: fmt.Sprintf("proxyTimeTick%d", randVal)})
	assert.Nil(t, err)
	err = svr.Start()
	assert.Nil(t, err)

	cli, err := NewGrpcClient(fmt.Sprintf("127.0.0.1:%d", cms.Params.Port), 3*time.Second)
	assert.Nil(t, err)

	err = cli.Init(&cms.InitParams{ProxyTimeTickChannel: fmt.Sprintf("proxyTimeTick%d", randVal)})
	assert.Nil(t, err)

	err = cli.Start()
	assert.Nil(t, err)

	t.Run("create collection", func(t *testing.T) {
		schema := schemapb.CollectionSchema{
			Name:        "testColl",
			Description: "testColl",
			AutoID:      true,
			Fields: []*schemapb.FieldSchema{
				{
					FieldID:      100,
					Name:         "vector",
					IsPrimaryKey: false,
					Description:  "vector",
					DataType:     schemapb.DataType_VECTOR_FLOAT,
					TypeParams:   nil,
					IndexParams:  nil,
				},
			},
		}

		sbf, err := proto.Marshal(&schema)
		assert.Nil(t, err)

		req := &milvuspb.CreateCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kCreateCollection,
				MsgID:     100,
				Timestamp: 100,
				SourceID:  100,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			Schema:         sbf,
		}

		status, err := cli.CreateCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, len(createCollectionArray), 1)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, createCollectionArray[0].Base.MsgType, commonpb.MsgType_kCreateCollection)
		assert.Equal(t, createCollectionArray[0].CollectionName, "testColl")
	})

	t.Run("has collection", func(t *testing.T) {
		req := &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kHasCollection,
				MsgID:     101,
				Timestamp: 101,
				SourceID:  101,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		rsp, err := cli.HasCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Value, true)

		req = &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kHasCollection,
				MsgID:     102,
				Timestamp: 102,
				SourceID:  102,
			},
			DbName:         "testDb",
			CollectionName: "testColl2",
		}
		rsp, err = cli.HasCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Value, false)

		req = &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kHasCollection,
				MsgID:     102,
				Timestamp: 102,
				SourceID:  102,
			},
			DbName:         "testDb",
			CollectionName: "testColl2",
		}
		rsp, err = cli.HasCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_UNEXPECTED_ERROR)

	})

	t.Run("describe collection", func(t *testing.T) {
		req := &milvuspb.DescribeCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDescribeCollection,
				MsgID:     103,
				Timestamp: 103,
				SourceID:  103,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		rsp, err := cli.DescribeCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Schema.Name, "testColl")
	})

	t.Run("show collection", func(t *testing.T) {
		req := &milvuspb.ShowCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kShowCollections,
				MsgID:     106,
				Timestamp: 106,
				SourceID:  106,
			},
			DbName: "testDb",
		}
		rsp, err := cli.ShowCollections(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.CollectionNames[0], "testColl")
		assert.Equal(t, len(rsp.CollectionNames), 1)
	})

	t.Run("create partition", func(t *testing.T) {
		req := &milvuspb.CreatePartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kCreatePartition,
				MsgID:     107,
				Timestamp: 107,
				SourceID:  107,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			PartitionName:  "testPartition",
		}
		status, err := cli.CreatePartition(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		collMeta, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		assert.Equal(t, len(collMeta.PartitionIDs), 2)
		partMeta, err := core.MetaTable.GetPartitionByID(collMeta.PartitionIDs[1])
		assert.Nil(t, err)
		assert.Equal(t, partMeta.PartitionName, "testPartition")

	})

	t.Run("has partition", func(t *testing.T) {
		req := &milvuspb.HasPartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kHasPartition,
				MsgID:     108,
				Timestamp: 108,
				SourceID:  108,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			PartitionName:  "testPartition",
		}
		rsp, err := cli.HasPartition(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Value, true)
	})

	t.Run("show partition", func(t *testing.T) {
		req := &milvuspb.ShowPartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kShowPartitions,
				MsgID:     110,
				Timestamp: 110,
				SourceID:  110,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		rsp, err := cli.ShowPartitions(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, len(rsp.PartitionNames), 2)
	})

	t.Run("drop partition", func(t *testing.T) {
		req := &milvuspb.DropPartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDropPartition,
				MsgID:     199,
				Timestamp: 199,
				SourceID:  199,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			PartitionName:  "testPartition",
		}
		status, err := cli.DropPartition(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		collMeta, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		assert.Equal(t, len(collMeta.PartitionIDs), 1)
		partMeta, err := core.MetaTable.GetPartitionByID(collMeta.PartitionIDs[0])
		assert.Nil(t, err)
		assert.Equal(t, partMeta.PartitionName, cms.Params.DefaultPartitionName)
	})

	t.Run("drop collection", func(t *testing.T) {
		req := &milvuspb.DropCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDropCollection,
				MsgID:     200,
				Timestamp: 200,
				SourceID:  200,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}

		status, err := cli.DropCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, len(dropCollectionArray), 1)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, dropCollectionArray[0].Base.MsgType, commonpb.MsgType_kDropCollection)
		assert.Equal(t, dropCollectionArray[0].CollectionName, "testColl")

		req = &milvuspb.DropCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDropCollection,
				MsgID:     200,
				Timestamp: 200,
				SourceID:  200,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		status, err = cli.DropCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, len(dropCollectionArray), 1)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_UNEXPECTED_ERROR)
	})

	err = cli.Stop()
	assert.Nil(t, err)

	err = svr.Stop()
	assert.Nil(t, err)
}
