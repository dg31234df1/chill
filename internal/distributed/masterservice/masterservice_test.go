// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package grpcmasterservice

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	grpcmasterserviceclient "github.com/milvus-io/milvus/internal/distributed/masterservice/client"

	"github.com/golang/protobuf/proto"
	cms "github.com/milvus-io/milvus/internal/masterservice"
	"github.com/milvus-io/milvus/internal/msgstream"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/etcdpb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/proto/masterpb"
	"github.com/milvus-io/milvus/internal/proto/milvuspb"
	"github.com/milvus-io/milvus/internal/proto/schemapb"
	"github.com/milvus-io/milvus/internal/types"
	"github.com/milvus-io/milvus/internal/util/retry"
	"github.com/milvus-io/milvus/internal/util/sessionutil"
	"github.com/milvus-io/milvus/internal/util/typeutil"
	"github.com/stretchr/testify/assert"
	"go.etcd.io/etcd/clientv3"
)

func GenSegInfoMsgPack(seg *datapb.SegmentInfo) *msgstream.MsgPack {
	msgPack := msgstream.MsgPack{}
	baseMsg := msgstream.BaseMsg{
		BeginTimestamp: 0,
		EndTimestamp:   0,
		HashValues:     []uint32{0},
	}
	segMsg := &msgstream.SegmentInfoMsg{
		BaseMsg: baseMsg,
		SegmentMsg: datapb.SegmentMsg{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_SegmentInfo,
				MsgID:     0,
				Timestamp: 0,
				SourceID:  0,
			},
			Segment: seg,
		},
	}
	msgPack.Msgs = append(msgPack.Msgs, segMsg)
	return &msgPack
}

func GenFlushedSegMsgPack(segID typeutil.UniqueID) *msgstream.MsgPack {
	msgPack := msgstream.MsgPack{}
	baseMsg := msgstream.BaseMsg{
		BeginTimestamp: 0,
		EndTimestamp:   0,
		HashValues:     []uint32{0},
	}
	segMsg := &msgstream.FlushCompletedMsg{
		BaseMsg: baseMsg,
		SegmentFlushCompletedMsg: internalpb.SegmentFlushCompletedMsg{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_SegmentFlushDone,
				MsgID:     0,
				Timestamp: 0,
				SourceID:  0,
			},
			SegmentID: segID,
		},
	}
	msgPack.Msgs = append(msgPack.Msgs, segMsg)
	return &msgPack
}

func TestGrpcService(t *testing.T) {
	const (
		dbName    = "testDB"
		collName  = "testColl"
		collName2 = "testColl-again"
		partName  = "testPartition"
		fieldName = "vector"
		fieldID   = 100
		segID     = 1001
	)
	rand.Seed(time.Now().UnixNano())
	randVal := rand.Int()

	Params.Init()
	Params.Port = (randVal % 100) + 10000
	parts := strings.Split(Params.Address, ":")
	if len(parts) == 2 {
		Params.Address = parts[0] + ":" + strconv.Itoa(Params.Port)
		t.Log("newParams.Address:", Params.Address)
	}

	ctx := context.Background()
	msFactory := msgstream.NewPmsFactory()
	svr, err := NewServer(ctx, msFactory)
	assert.Nil(t, err)

	cms.Params.Init()
	cms.Params.MetaRootPath = fmt.Sprintf("/%d/test/meta", randVal)
	cms.Params.KvRootPath = fmt.Sprintf("/%d/test/kv", randVal)
	cms.Params.ProxyTimeTickChannel = fmt.Sprintf("proxyTimeTick%d", randVal)
	cms.Params.MsgChannelSubName = fmt.Sprintf("msgChannel%d", randVal)
	cms.Params.TimeTickChannel = fmt.Sprintf("timeTick%d", randVal)
	cms.Params.DdChannel = fmt.Sprintf("ddChannel%d", randVal)
	cms.Params.StatisticsChannel = fmt.Sprintf("stateChannel%d", randVal)
	cms.Params.DataServiceSegmentChannel = fmt.Sprintf("segmentChannel%d", randVal)

	cms.Params.MaxPartitionNum = 64
	cms.Params.DefaultPartitionName = "_default"
	cms.Params.DefaultIndexName = "_default"

	t.Logf("master service port = %d", Params.Port)

	core, ok := (svr.masterService).(*cms.Core)
	assert.True(t, ok)

	err = core.Register()
	assert.Nil(t, err)

	err = svr.startGrpc()
	assert.Nil(t, err)
	svr.masterService.UpdateStateCode(internalpb.StateCode_Initializing)

	etcdCli, err := initEtcd(cms.Params.EtcdAddress)
	assert.Nil(t, err)
	_, err = etcdCli.Delete(ctx, sessionutil.DefaultServiceRoot, clientv3.WithPrefix())
	assert.Nil(t, err)

	err = core.Init()
	assert.Nil(t, err)

	core.ProxyTimeTickChan = make(chan typeutil.Timestamp, 8)
	FlushedSegmentChan := make(chan *msgstream.MsgPack, 8)
	core.DataNodeFlushedSegmentChan = FlushedSegmentChan
	SegmentInfoChan := make(chan *msgstream.MsgPack, 8)
	core.DataServiceSegmentChan = SegmentInfoChan

	timeTickArray := make([]typeutil.Timestamp, 0, 16)
	core.SendTimeTick = func(ts typeutil.Timestamp) error {
		t.Logf("send time tick %d", ts)
		timeTickArray = append(timeTickArray, ts)
		return nil
	}
	createCollectionArray := make([]*internalpb.CreateCollectionRequest, 0, 16)
	core.SendDdCreateCollectionReq = func(ctx context.Context, req *internalpb.CreateCollectionRequest) error {
		t.Logf("Create Colllection %s", req.CollectionName)
		createCollectionArray = append(createCollectionArray, req)
		return nil
	}

	dropCollectionArray := make([]*internalpb.DropCollectionRequest, 0, 16)
	core.SendDdDropCollectionReq = func(ctx context.Context, req *internalpb.DropCollectionRequest) error {
		t.Logf("Drop Collection %s", req.CollectionName)
		dropCollectionArray = append(dropCollectionArray, req)
		return nil
	}

	createPartitionArray := make([]*internalpb.CreatePartitionRequest, 0, 16)
	core.SendDdCreatePartitionReq = func(ctx context.Context, req *internalpb.CreatePartitionRequest) error {
		t.Logf("Create Partition %s", req.PartitionName)
		createPartitionArray = append(createPartitionArray, req)
		return nil
	}

	dropPartitionArray := make([]*internalpb.DropPartitionRequest, 0, 16)
	core.SendDdDropPartitionReq = func(ctx context.Context, req *internalpb.DropPartitionRequest) error {
		t.Logf("Drop Partition %s", req.PartitionName)
		dropPartitionArray = append(dropPartitionArray, req)
		return nil
	}

	core.CallGetBinlogFilePathsService = func(segID typeutil.UniqueID, fieldID typeutil.UniqueID) ([]string, error) {
		return []string{"file1", "file2", "file3"}, nil
	}
	core.CallGetNumRowsService = func(segID typeutil.UniqueID, isFromFlushedChan bool) (int64, error) {
		return cms.Params.MinSegmentSizeToEnableIndex, nil
	}

	var binlogLock sync.Mutex
	binlogPathArray := make([]string, 0, 16)
	core.CallBuildIndexService = func(ctx context.Context, binlog []string, field *schemapb.FieldSchema, idxInfo *etcdpb.IndexInfo) (typeutil.UniqueID, error) {
		binlogLock.Lock()
		defer binlogLock.Unlock()
		binlogPathArray = append(binlogPathArray, binlog...)
		return 2000, nil
	}

	var dropIDLock sync.Mutex
	dropID := make([]typeutil.UniqueID, 0, 16)
	core.CallDropIndexService = func(ctx context.Context, indexID typeutil.UniqueID) error {
		dropIDLock.Lock()
		defer dropIDLock.Unlock()
		dropID = append(dropID, indexID)
		return nil
	}

	collectionMetaCache := make([]string, 0, 16)
	core.CallInvalidateCollectionMetaCacheService = func(ctx context.Context, ts typeutil.Timestamp, dbName string, collectionName string) error {
		collectionMetaCache = append(collectionMetaCache, collectionName)
		return nil
	}

	core.CallReleaseCollectionService = func(ctx context.Context, ts typeutil.Timestamp, dbID typeutil.UniqueID, collectionID typeutil.UniqueID) error {
		return nil
	}

	err = svr.start()
	assert.Nil(t, err)

	svr.masterService.UpdateStateCode(internalpb.StateCode_Healthy)

	cli, err := grpcmasterserviceclient.NewClient(Params.Address, []string{cms.Params.EtcdAddress}, 3*time.Second)
	assert.Nil(t, err)

	err = cli.Init()
	assert.Nil(t, err)

	err = cli.Start()
	assert.Nil(t, err)

	t.Run("get component states", func(t *testing.T) {
		req := &internalpb.GetComponentStatesRequest{}
		rsp, err := svr.GetComponentStates(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
	})

	t.Run("get time tick channel", func(t *testing.T) {
		req := &internalpb.GetTimeTickChannelRequest{}
		rsp, err := svr.GetTimeTickChannel(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
	})

	t.Run("get statistics channel", func(t *testing.T) {
		req := &internalpb.GetStatisticsChannelRequest{}
		rsp, err := svr.GetStatisticsChannel(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
	})

	t.Run("get dd channel", func(t *testing.T) {
		req := &internalpb.GetDdChannelRequest{}
		rsp, err := svr.GetDdChannel(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
	})

	t.Run("alloc time stamp", func(t *testing.T) {
		req := &masterpb.AllocTimestampRequest{
			Count: 1,
		}
		rsp, err := svr.AllocTimestamp(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
	})

	t.Run("alloc id", func(t *testing.T) {
		req := &masterpb.AllocIDRequest{
			Count: 1,
		}
		rsp, err := svr.AllocID(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
	})

	t.Run("create collection", func(t *testing.T) {
		schema := schemapb.CollectionSchema{
			Name:   collName,
			AutoID: true,
			Fields: []*schemapb.FieldSchema{
				{
					FieldID:      fieldID,
					Name:         fieldName,
					IsPrimaryKey: false,
					DataType:     schemapb.DataType_FloatVector,
					TypeParams:   nil,
					IndexParams: []*commonpb.KeyValuePair{
						{
							Key:   "ik1",
							Value: "iv1",
						},
					},
				},
			},
		}

		sbf, err := proto.Marshal(&schema)
		assert.Nil(t, err)

		req := &milvuspb.CreateCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_CreateCollection,
				MsgID:     100,
				Timestamp: 100,
				SourceID:  100,
			},
			DbName:         dbName,
			CollectionName: collName,
			Schema:         sbf,
		}

		status, err := cli.CreateCollection(ctx, req)
		assert.Nil(t, err)

		assert.Equal(t, 1, len(createCollectionArray))
		assert.Equal(t, commonpb.ErrorCode_Success, status.ErrorCode)
		assert.Equal(t, commonpb.MsgType_CreateCollection, createCollectionArray[0].Base.MsgType)
		assert.Equal(t, collName, createCollectionArray[0].CollectionName)

		req.Base.MsgID = 101
		req.Base.Timestamp = 101
		req.Base.SourceID = 101
		status, err = cli.CreateCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_UnexpectedError, status.ErrorCode)

		req.Base.MsgID = 102
		req.Base.Timestamp = 102
		req.Base.SourceID = 102
		req.CollectionName = collName2
		status, err = cli.CreateCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_UnexpectedError, status.ErrorCode)

		schema.Name = req.CollectionName
		sbf, err = proto.Marshal(&schema)
		assert.Nil(t, err)
		req.Schema = sbf
		req.Base.MsgID = 103
		req.Base.Timestamp = 103
		req.Base.SourceID = 103
		status, err = cli.CreateCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, status.ErrorCode)
		assert.Equal(t, 2, len(createCollectionArray))
		assert.Equal(t, commonpb.MsgType_CreateCollection, createCollectionArray[1].Base.MsgType)
		assert.Equal(t, collName2, createCollectionArray[1].CollectionName)

		//time stamp go back, master response to add the timestamp, so the time tick will never go back
		//schema.Name = "testColl-goback"
		//sbf, err = proto.Marshal(&schema)
		//assert.Nil(t, err)
		//req.CollectionName = schema.Name
		//req.Schema = sbf
		//req.Base.MsgID = 103
		//req.Base.Timestamp = 103
		//req.Base.SourceID = 103
		//status, err = cli.CreateCollection(ctx, req)
		//assert.Nil(t, err)
		//assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_UnexpectedError)
		//matched, err := regexp.MatchString("input timestamp = [0-9]+, last dd time stamp = [0-9]+", status.Reason)
		//assert.Nil(t, err)
		//assert.True(t, matched)
	})

	t.Run("has collection", func(t *testing.T) {
		req := &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_HasCollection,
				MsgID:     110,
				Timestamp: 110,
				SourceID:  110,
			},
			DbName:         "testDb",
			CollectionName: collName,
		}
		rsp, err := cli.HasCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, true, rsp.Value)

		req = &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_HasCollection,
				MsgID:     111,
				Timestamp: 111,
				SourceID:  111,
			},
			DbName:         "testDb",
			CollectionName: "testColl2",
		}
		rsp, err = cli.HasCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, false, rsp.Value)

		// test time stamp go back
		req = &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_HasCollection,
				MsgID:     111,
				Timestamp: 111,
				SourceID:  111,
			},
			DbName:         "testDb",
			CollectionName: "testColl2",
		}
		rsp, err = cli.HasCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, false, rsp.Value)
	})

	t.Run("describe collection", func(t *testing.T) {
		collMeta, err := core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)
		req := &milvuspb.DescribeCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DescribeCollection,
				MsgID:     120,
				Timestamp: 120,
				SourceID:  120,
			},
			DbName:         "testDb",
			CollectionName: collName,
		}
		rsp, err := cli.DescribeCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, collName, rsp.Schema.Name)
		assert.Equal(t, collMeta.ID, rsp.CollectionID)
	})

	t.Run("show collection", func(t *testing.T) {
		req := &milvuspb.ShowCollectionsRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_ShowCollections,
				MsgID:     130,
				Timestamp: 130,
				SourceID:  130,
			},
			DbName: "testDb",
		}
		rsp, err := cli.ShowCollections(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.ElementsMatch(t, rsp.CollectionNames, []string{collName, collName2})
		assert.Equal(t, 2, len(rsp.CollectionNames))
	})

	t.Run("create partition", func(t *testing.T) {
		req := &milvuspb.CreatePartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_CreatePartition,
				MsgID:     140,
				Timestamp: 140,
				SourceID:  140,
			},
			DbName:         dbName,
			CollectionName: collName,
			PartitionName:  partName,
		}
		status, err := cli.CreatePartition(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, status.ErrorCode)
		collMeta, err := core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(collMeta.PartitionIDs))
		partMeta, err := core.MetaTable.GetPartitionByID(1, collMeta.PartitionIDs[1], 0)
		assert.Nil(t, err)
		assert.Equal(t, partName, partMeta.PartitionName)
		assert.Equal(t, 1, len(collectionMetaCache))
	})

	t.Run("has partition", func(t *testing.T) {
		req := &milvuspb.HasPartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_HasPartition,
				MsgID:     150,
				Timestamp: 150,
				SourceID:  150,
			},
			DbName:         dbName,
			CollectionName: collName,
			PartitionName:  partName,
		}
		rsp, err := cli.HasPartition(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, true, rsp.Value)
	})

	t.Run("show partition", func(t *testing.T) {
		coll, err := core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)
		req := &milvuspb.ShowPartitionsRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_ShowPartitions,
				MsgID:     160,
				Timestamp: 160,
				SourceID:  160,
			},
			DbName:         "testDb",
			CollectionName: collName,
			CollectionID:   coll.ID,
		}
		rsp, err := cli.ShowPartitions(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, 2, len(rsp.PartitionNames))
		assert.Equal(t, 2, len(rsp.PartitionIDs))
	})

	t.Run("show segment", func(t *testing.T) {
		coll, err := core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)
		partID := coll.PartitionIDs[1]
		part, err := core.MetaTable.GetPartitionByID(1, partID, 0)
		assert.Nil(t, err)
		assert.Zero(t, len(part.SegmentIDs))
		seg := &datapb.SegmentInfo{
			ID:           1000,
			CollectionID: coll.ID,
			PartitionID:  part.PartitionID,
		}
		segInfoMsgPack := GenSegInfoMsgPack(seg)
		SegmentInfoChan <- segInfoMsgPack
		time.Sleep(time.Millisecond * 100)
		part, err = core.MetaTable.GetPartitionByID(1, partID, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(part.SegmentIDs))

		// send msg twice, partition still contains 1 segment
		segInfoMsgPack1 := GenSegInfoMsgPack(seg)
		SegmentInfoChan <- segInfoMsgPack1
		time.Sleep(time.Millisecond * 100)
		part1, err := core.MetaTable.GetPartitionByID(1, partID, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(part1.SegmentIDs))

		req := &milvuspb.ShowSegmentsRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_ShowSegments,
				MsgID:     170,
				Timestamp: 170,
				SourceID:  170,
			},
			CollectionID: coll.ID,
			PartitionID:  partID,
		}
		rsp, err := cli.ShowSegments(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, int64(1000), rsp.SegmentIDs[0])
		assert.Equal(t, 1, len(rsp.SegmentIDs))
	})

	t.Run("create index", func(t *testing.T) {
		req := &milvuspb.CreateIndexRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_CreateIndex,
				MsgID:     180,
				Timestamp: 180,
				SourceID:  180,
			},
			DbName:         dbName,
			CollectionName: collName,
			FieldName:      fieldName,
			ExtraParams: []*commonpb.KeyValuePair{
				{
					Key:   "ik1",
					Value: "iv1",
				},
			},
		}
		collMeta, err := core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)
		assert.Zero(t, len(collMeta.FieldIndexes))
		rsp, err := cli.CreateIndex(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.ErrorCode)
		collMeta, err = core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(collMeta.FieldIndexes))

		binlogLock.Lock()
		defer binlogLock.Unlock()
		assert.Equal(t, 3, len(binlogPathArray))
		assert.ElementsMatch(t, binlogPathArray, []string{"file1", "file2", "file3"})

		req.FieldName = "no field"
		rsp, err = cli.CreateIndex(ctx, req)
		assert.Nil(t, err)
		assert.NotEqual(t, commonpb.ErrorCode_Success, rsp.ErrorCode)
	})

	t.Run("describe segment", func(t *testing.T) {
		coll, err := core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)

		req := &milvuspb.DescribeSegmentRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DescribeSegment,
				MsgID:     190,
				Timestamp: 190,
				SourceID:  190,
			},
			CollectionID: coll.ID,
			SegmentID:    1000,
		}
		rsp, err := cli.DescribeSegment(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		t.Logf("index id = %d", rsp.IndexID)
	})

	t.Run("describe index", func(t *testing.T) {
		req := &milvuspb.DescribeIndexRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DescribeIndex,
				MsgID:     200,
				Timestamp: 200,
				SourceID:  200,
			},
			DbName:         dbName,
			CollectionName: collName,
			FieldName:      fieldName,
			IndexName:      "",
		}
		rsp, err := cli.DescribeIndex(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, 1, len(rsp.IndexDescriptions))
		assert.Equal(t, cms.Params.DefaultIndexName, rsp.IndexDescriptions[0].IndexName)
	})

	t.Run("flush segment", func(t *testing.T) {
		coll, err := core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)
		partID := coll.PartitionIDs[1]
		part, err := core.MetaTable.GetPartitionByID(1, partID, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(part.SegmentIDs))
		seg := &datapb.SegmentInfo{
			ID:           segID,
			CollectionID: coll.ID,
			PartitionID:  part.PartitionID,
		}
		segInfoMsgPack := GenSegInfoMsgPack(seg)
		SegmentInfoChan <- segInfoMsgPack
		time.Sleep(time.Millisecond * 100)
		part, err = core.MetaTable.GetPartitionByID(1, partID, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(part.SegmentIDs))
		flushedSegMsgPack := GenFlushedSegMsgPack(segID)
		FlushedSegmentChan <- flushedSegMsgPack
		time.Sleep(time.Millisecond * 100)
		segIdxInfo, err := core.MetaTable.GetSegmentIndexInfoByID(segID, -1, "")
		assert.Nil(t, err)

		// send msg twice, segIdxInfo should not change
		flushedSegMsgPack1 := GenFlushedSegMsgPack(segID)
		FlushedSegmentChan <- flushedSegMsgPack1
		time.Sleep(time.Millisecond * 100)
		segIdxInfo1, err := core.MetaTable.GetSegmentIndexInfoByID(segID, -1, "")
		assert.Nil(t, err)
		assert.Equal(t, segIdxInfo, segIdxInfo1)

		req := &milvuspb.DescribeIndexRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DescribeIndex,
				MsgID:     210,
				Timestamp: 210,
				SourceID:  210,
			},
			DbName:         dbName,
			CollectionName: collName,
			FieldName:      fieldName,
			IndexName:      "",
		}
		rsp, err := cli.DescribeIndex(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.Status.ErrorCode)
		assert.Equal(t, 1, len(rsp.IndexDescriptions))
		assert.Equal(t, cms.Params.DefaultIndexName, rsp.IndexDescriptions[0].IndexName)

	})

	t.Run("drop index", func(t *testing.T) {
		req := &milvuspb.DropIndexRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DropIndex,
				MsgID:     215,
				Timestamp: 215,
				SourceID:  215,
			},
			DbName:         dbName,
			CollectionName: collName,
			FieldName:      fieldName,
			IndexName:      cms.Params.DefaultIndexName,
		}
		_, idx, err := core.MetaTable.GetIndexByName(collName, cms.Params.DefaultIndexName)
		assert.Nil(t, err)
		assert.Equal(t, len(idx), 1)
		rsp, err := cli.DropIndex(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, rsp.ErrorCode)

		dropIDLock.Lock()
		assert.Equal(t, 1, len(dropID))
		assert.Equal(t, idx[0].IndexID, dropID[0])
		dropIDLock.Unlock()
	})

	t.Run("drop partition", func(t *testing.T) {
		req := &milvuspb.DropPartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DropPartition,
				MsgID:     220,
				Timestamp: 220,
				SourceID:  220,
			},
			DbName:         dbName,
			CollectionName: collName,
			PartitionName:  partName,
		}
		status, err := cli.DropPartition(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, status.ErrorCode)
		collMeta, err := core.MetaTable.GetCollectionByName(collName, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(collMeta.PartitionIDs))
		partMeta, err := core.MetaTable.GetPartitionByID(1, collMeta.PartitionIDs[0], 0)
		assert.Nil(t, err)
		assert.Equal(t, cms.Params.DefaultPartitionName, partMeta.PartitionName)
		assert.Equal(t, 2, len(collectionMetaCache))
	})

	t.Run("drop collection", func(t *testing.T) {
		req := &milvuspb.DropCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DropCollection,
				MsgID:     230,
				Timestamp: 230,
				SourceID:  230,
			},
			DbName:         "testDb",
			CollectionName: collName,
		}

		status, err := cli.DropCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(dropCollectionArray))
		assert.Equal(t, commonpb.MsgType_DropCollection, dropCollectionArray[0].Base.MsgType)
		assert.Equal(t, commonpb.ErrorCode_Success, status.ErrorCode)
		assert.Equal(t, collName, dropCollectionArray[0].CollectionName)
		assert.Equal(t, 3, len(collectionMetaCache))
		assert.Equal(t, collName, collectionMetaCache[0])

		req = &milvuspb.DropCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_DropCollection,
				MsgID:     231,
				Timestamp: 231,
				SourceID:  231,
			},
			DbName:         "testDb",
			CollectionName: collName,
		}
		status, err = cli.DropCollection(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(dropCollectionArray))
		assert.Equal(t, commonpb.ErrorCode_UnexpectedError, status.ErrorCode)
	})

	err = cli.Stop()
	assert.Nil(t, err)

	err = svr.Stop()
	assert.Nil(t, err)
}

type mockCore struct {
	types.MasterComponent
}

func (m *mockCore) UpdateStateCode(internalpb.StateCode) {
}
func (m *mockCore) SetProxyService(context.Context, types.ProxyService) error {
	return nil
}
func (m *mockCore) SetDataService(context.Context, types.DataService) error {
	return nil
}
func (m *mockCore) SetIndexService(types.IndexService) error {
	return nil
}

func (m *mockCore) SetQueryService(types.QueryService) error {
	return nil
}

func (m *mockCore) Register() error {
	return nil
}

func (m *mockCore) Init() error {
	return nil
}

func (m *mockCore) Start() error {
	return nil
}

func (m *mockCore) Stop() error {
	return fmt.Errorf("stop error")
}

type mockProxy struct {
	types.ProxyService
}

func (m *mockProxy) Init() error {
	return nil
}
func (m *mockProxy) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	return &internalpb.ComponentStates{
		State: &internalpb.ComponentInfo{
			StateCode: internalpb.StateCode_Healthy,
		},
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		SubcomponentStates: []*internalpb.ComponentInfo{
			{
				StateCode: internalpb.StateCode_Healthy,
			},
		},
	}, nil
}
func (m *mockProxy) Stop() error {
	return fmt.Errorf("stop error")
}

type mockDataService struct {
	types.DataService
}

func (m *mockDataService) Init() error {
	return nil
}
func (m *mockDataService) Start() error {
	return nil
}
func (m *mockDataService) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	return &internalpb.ComponentStates{
		State: &internalpb.ComponentInfo{
			StateCode: internalpb.StateCode_Healthy,
		},
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		SubcomponentStates: []*internalpb.ComponentInfo{
			{
				StateCode: internalpb.StateCode_Healthy,
			},
		},
	}, nil
}
func (m *mockDataService) Stop() error {
	return fmt.Errorf("stop error")
}

type mockIndex struct {
	types.IndexService
}

func (m *mockIndex) Init() error {
	return nil
}

func (m *mockIndex) Stop() error {
	return fmt.Errorf("stop error")
}

type mockQuery struct {
	types.QueryService
}

func (m *mockQuery) Init() error {
	return nil
}

func (m *mockQuery) Start() error {
	return nil
}

func (m *mockQuery) Stop() error {
	return fmt.Errorf("stop error")
}

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	svr := Server{
		masterService: &mockCore{},
		ctx:           ctx,
		cancel:        cancel,
		grpcErrChan:   make(chan error),
	}
	Params.Init()
	Params.Port = 1000000
	err := svr.Run()
	assert.NotNil(t, err)
	assert.EqualError(t, err, "listen tcp: address 1000000: invalid port")

	svr.newProxyServiceClient = func(s string) types.ProxyService {
		return &mockProxy{}
	}
	svr.newDataServiceClient = func(s, address string, timeout time.Duration) types.DataService {
		return &mockDataService{}
	}
	svr.newIndexServiceClient = func(s string) types.IndexService {
		return &mockIndex{}
	}
	svr.newQueryServiceClient = func(s string) (types.QueryService, error) {
		return &mockQuery{}, nil
	}

	Params.Port = rand.Int()%100 + 10000

	rand.Seed(time.Now().UnixNano())
	randVal := rand.Int()
	cms.Params.Init()
	cms.Params.MetaRootPath = fmt.Sprintf("/%d/test/meta", randVal)

	etcdCli, err := initEtcd(cms.Params.EtcdAddress)
	assert.Nil(t, err)
	_, err = etcdCli.Delete(ctx, sessionutil.DefaultServiceRoot, clientv3.WithPrefix())
	assert.Nil(t, err)
	err = svr.Run()
	assert.Nil(t, err)

	err = svr.Stop()
	assert.Nil(t, err)

}

func initEtcd(etcdAddress string) (*clientv3.Client, error) {
	var etcdCli *clientv3.Client
	connectEtcdFn := func() error {
		etcd, err := clientv3.New(clientv3.Config{Endpoints: []string{etcdAddress}, DialTimeout: 5 * time.Second})
		if err != nil {
			return err
		}
		etcdCli = etcd
		return nil
	}
	err := retry.Retry(100000, time.Millisecond*200, connectEtcdFn)
	if err != nil {
		return nil, err
	}
	return etcdCli, nil
}
