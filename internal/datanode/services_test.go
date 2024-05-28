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
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	allocator2 "github.com/milvus-io/milvus/internal/allocator"
	"github.com/milvus-io/milvus/internal/datanode/allocator"
	"github.com/milvus-io/milvus/internal/datanode/broker"
	"github.com/milvus-io/milvus/internal/datanode/compaction"
	"github.com/milvus-io/milvus/internal/datanode/metacache"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/internal/util/sessionutil"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/etcd"
	"github.com/milvus-io/milvus/pkg/util/merr"
	"github.com/milvus-io/milvus/pkg/util/metricsinfo"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
)

type DataNodeServicesSuite struct {
	suite.Suite

	broker  *broker.MockBroker
	node    *DataNode
	etcdCli *clientv3.Client
	ctx     context.Context
	cancel  context.CancelFunc
}

func TestDataNodeServicesSuite(t *testing.T) {
	suite.Run(t, new(DataNodeServicesSuite))
}

func (s *DataNodeServicesSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	etcdCli, err := etcd.GetEtcdClient(
		Params.EtcdCfg.UseEmbedEtcd.GetAsBool(),
		Params.EtcdCfg.EtcdUseSSL.GetAsBool(),
		Params.EtcdCfg.Endpoints.GetAsStrings(),
		Params.EtcdCfg.EtcdTLSCert.GetValue(),
		Params.EtcdCfg.EtcdTLSKey.GetValue(),
		Params.EtcdCfg.EtcdTLSCACert.GetValue(),
		Params.EtcdCfg.EtcdTLSMinVersion.GetValue())
	s.Require().NoError(err)
	s.etcdCli = etcdCli
}

func (s *DataNodeServicesSuite) SetupTest() {
	s.node = newIDLEDataNodeMock(s.ctx, schemapb.DataType_Int64)
	s.node.SetEtcdClient(s.etcdCli)

	err := s.node.Init()
	s.Require().NoError(err)

	alloc := allocator.NewMockAllocator(s.T())
	alloc.EXPECT().Start().Return(nil).Maybe()
	alloc.EXPECT().Close().Maybe()
	alloc.EXPECT().GetIDAlloactor().Return(&allocator2.IDAllocator{}).Maybe()
	alloc.EXPECT().Alloc(mock.Anything).Call.Return(int64(22222),
		func(count uint32) int64 {
			return int64(22222 + count)
		}, nil).Maybe()
	s.node.allocator = alloc

	broker := broker.NewMockBroker(s.T())
	broker.EXPECT().GetSegmentInfo(mock.Anything, mock.Anything).
		Return([]*datapb.SegmentInfo{}, nil).Maybe()
	broker.EXPECT().ReportTimeTick(mock.Anything, mock.Anything).Return(nil).Maybe()
	broker.EXPECT().SaveBinlogPaths(mock.Anything, mock.Anything).Return(nil).Maybe()
	broker.EXPECT().UpdateChannelCheckpoint(mock.Anything, mock.Anything).Return(nil).Maybe()

	s.broker = broker
	s.node.broker = broker

	err = s.node.Start()
	s.Require().NoError(err)

	s.node.chunkManager = storage.NewLocalChunkManager(storage.RootPath("/tmp/milvus_test/datanode"))
	paramtable.SetNodeID(1)
}

func (s *DataNodeServicesSuite) TearDownTest() {
	if s.broker != nil {
		s.broker.AssertExpectations(s.T())
		s.broker = nil
	}

	if s.node != nil {
		s.node.Stop()
		s.node = nil
	}
}

func (s *DataNodeServicesSuite) TearDownSuite() {
	s.cancel()
	err := s.etcdCli.Close()
	s.Require().NoError(err)
}

func (s *DataNodeServicesSuite) TestNotInUseAPIs() {
	s.Run("WatchDmChannels", func() {
		status, err := s.node.WatchDmChannels(s.ctx, &datapb.WatchDmChannelsRequest{})
		s.Assert().NoError(err)
		s.Assert().True(merr.Ok(status))
	})
	s.Run("GetTimeTickChannel", func() {
		_, err := s.node.GetTimeTickChannel(s.ctx, nil)
		s.Assert().NoError(err)
	})

	s.Run("GetStatisticsChannel", func() {
		_, err := s.node.GetStatisticsChannel(s.ctx, nil)
		s.Assert().NoError(err)
	})
}

func (s *DataNodeServicesSuite) TestGetComponentStates() {
	resp, err := s.node.GetComponentStates(s.ctx, nil)
	s.Assert().NoError(err)
	s.Assert().True(merr.Ok(resp.GetStatus()))
	s.Assert().Equal(common.NotRegisteredID, resp.State.NodeID)

	s.node.SetSession(&sessionutil.Session{})
	s.node.session.UpdateRegistered(true)
	resp, err = s.node.GetComponentStates(context.Background(), nil)
	s.Assert().NoError(err)
	s.Assert().True(merr.Ok(resp.GetStatus()))
}

func (s *DataNodeServicesSuite) TestGetCompactionState() {
	s.Run("success", func() {
		mockC := compaction.NewMockCompactor(s.T())
		s.node.compactionExecutor.executing.Insert(int64(3), mockC)

		mockC2 := compaction.NewMockCompactor(s.T())
		s.node.compactionExecutor.executing.Insert(int64(2), mockC2)

		s.node.compactionExecutor.completed.Insert(int64(1), &datapb.CompactionPlanResult{
			PlanID: 1,
			State:  commonpb.CompactionState_Completed,
			Segments: []*datapb.CompactionSegment{
				{SegmentID: 10},
			},
		})

		s.node.compactionExecutor.completed.Insert(int64(4), &datapb.CompactionPlanResult{
			PlanID: 4,
			Type:   datapb.CompactionType_Level0DeleteCompaction,
			State:  commonpb.CompactionState_Completed,
		})

		stat, err := s.node.GetCompactionState(s.ctx, nil)
		s.Assert().NoError(err)
		s.Assert().Equal(4, len(stat.GetResults()))

		var mu sync.RWMutex
		cnt := 0
		for _, v := range stat.GetResults() {
			if v.GetState() == commonpb.CompactionState_Completed {
				mu.Lock()
				cnt++
				mu.Unlock()
			}
		}
		mu.Lock()
		s.Assert().Equal(2, cnt)
		mu.Unlock()

		s.Assert().Equal(1, s.node.compactionExecutor.completed.Len())
	})

	s.Run("unhealthy", func() {
		node := &DataNode{}
		node.UpdateStateCode(commonpb.StateCode_Abnormal)
		resp, _ := node.GetCompactionState(s.ctx, nil)
		s.Assert().Equal(merr.Code(merr.ErrServiceNotReady), resp.GetStatus().GetCode())
	})
}

func (s *DataNodeServicesSuite) TestCompaction() {
	dmChannelName := "by-dev-rootcoord-dml_0_100v0"

	s.Run("service_not_ready", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		node := &DataNode{}
		node.UpdateStateCode(commonpb.StateCode_Abnormal)
		req := &datapb.CompactionPlan{
			PlanID:  1000,
			Channel: dmChannelName,
		}

		resp, err := node.Compaction(ctx, req)
		s.NoError(err)
		s.False(merr.Ok(resp))
	})

	s.Run("unknown CompactionType", func() {
		node := s.node
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		req := &datapb.CompactionPlan{
			PlanID:  1000,
			Channel: dmChannelName,
			SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
				{SegmentID: 102, Level: datapb.SegmentLevel_L0},
				{SegmentID: 103, Level: datapb.SegmentLevel_L1},
			},
		}

		resp, err := node.Compaction(ctx, req)
		s.NoError(err)
		s.False(merr.Ok(resp))
	})
}

func (s *DataNodeServicesSuite) TestFlushSegments() {
	dmChannelName := "fake-by-dev-rootcoord-dml-channel-test-FlushSegments"
	schema := &schemapb.CollectionSchema{
		Name: "test_collection",
		Fields: []*schemapb.FieldSchema{
			{FieldID: common.RowIDField, Name: common.RowIDFieldName, DataType: schemapb.DataType_Int64},
			{FieldID: common.TimeStampField, Name: common.TimeStampFieldName, DataType: schemapb.DataType_Int64},
			{FieldID: common.StartOfUserFieldID, DataType: schemapb.DataType_Int64, IsPrimaryKey: true, Name: "pk"},
			{FieldID: common.StartOfUserFieldID + 1, DataType: schemapb.DataType_FloatVector, TypeParams: []*commonpb.KeyValuePair{
				{Key: common.DimKey, Value: "128"},
			}},
		},
	}
	segmentID := int64(100)

	vchan := &datapb.VchannelInfo{
		CollectionID:        1,
		ChannelName:         dmChannelName,
		UnflushedSegmentIds: []int64{},
		FlushedSegmentIds:   []int64{},
	}

	err := s.node.flowgraphManager.AddandStartWithEtcdTickler(s.node, vchan, schema, genTestTickler())
	s.Require().NoError(err)

	fgservice, ok := s.node.flowgraphManager.GetFlowgraphService(dmChannelName)
	s.Require().True(ok)

	metaCache := metacache.NewMockMetaCache(s.T())
	metaCache.EXPECT().Collection().Return(1).Maybe()
	metaCache.EXPECT().Schema().Return(schema).Maybe()
	s.node.writeBufferManager.Register(dmChannelName, metaCache, nil)

	fgservice.metacache.AddSegment(&datapb.SegmentInfo{
		ID:            segmentID,
		CollectionID:  1,
		PartitionID:   2,
		State:         commonpb.SegmentState_Growing,
		StartPosition: &msgpb.MsgPosition{},
	}, func(_ *datapb.SegmentInfo) *metacache.BloomFilterSet { return metacache.NewBloomFilterSet() })

	s.Run("service_not_ready", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		node := &DataNode{}
		node.UpdateStateCode(commonpb.StateCode_Abnormal)
		req := &datapb.FlushSegmentsRequest{
			Base: &commonpb.MsgBase{
				TargetID: s.node.GetSession().ServerID,
			},
			DbID:         0,
			CollectionID: 1,
			SegmentIDs:   []int64{0},
		}

		resp, err := node.FlushSegments(ctx, req)
		s.NoError(err)
		s.False(merr.Ok(resp))
	})

	s.Run("node_id_not_match", func() {
		node := s.node
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		req := &datapb.FlushSegmentsRequest{
			Base: &commonpb.MsgBase{
				TargetID: s.node.GetSession().ServerID + 1,
			},
			DbID:         0,
			CollectionID: 1,
			SegmentIDs:   []int64{0},
		}

		resp, err := node.FlushSegments(ctx, req)
		s.NoError(err)
		s.False(merr.Ok(resp))
	})

	s.Run("channel_not_found", func() {
		node := s.node
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		req := &datapb.FlushSegmentsRequest{
			Base: &commonpb.MsgBase{
				TargetID: s.node.GetSession().ServerID,
			},
			DbID:         0,
			CollectionID: 1,
			SegmentIDs:   []int64{segmentID},
		}

		resp, err := node.FlushSegments(ctx, req)
		s.NoError(err)
		s.False(merr.Ok(resp))
	})

	s.Run("normal_flush", func() {
		node := s.node
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		req := &datapb.FlushSegmentsRequest{
			Base: &commonpb.MsgBase{
				TargetID: s.node.GetSession().ServerID,
			},
			DbID:         0,
			CollectionID: 1,
			SegmentIDs:   []int64{segmentID},
			ChannelName:  dmChannelName,
		}

		resp, err := node.FlushSegments(ctx, req)
		s.NoError(err)
		s.True(merr.Ok(resp))
	})
}

func (s *DataNodeServicesSuite) TestShowConfigurations() {
	pattern := "datanode.Port"
	req := &internalpb.ShowConfigurationsRequest{
		Base: &commonpb.MsgBase{
			MsgType: commonpb.MsgType_WatchQueryChannels,
			MsgID:   rand.Int63(),
		},
		Pattern: pattern,
	}

	// test closed server
	node := &DataNode{}
	node.SetSession(&sessionutil.Session{SessionRaw: sessionutil.SessionRaw{ServerID: 1}})
	node.stateCode.Store(commonpb.StateCode_Abnormal)

	resp, err := node.ShowConfigurations(s.ctx, req)
	s.Assert().NoError(err)
	s.Assert().False(merr.Ok(resp.GetStatus()))

	node.stateCode.Store(commonpb.StateCode_Healthy)
	resp, err = node.ShowConfigurations(s.ctx, req)
	s.Assert().NoError(err)
	s.Assert().True(merr.Ok(resp.GetStatus()))
	s.Assert().Equal(1, len(resp.Configuations))
	s.Assert().Equal("datanode.port", resp.Configuations[0].Key)
}

func (s *DataNodeServicesSuite) TestGetMetrics() {
	node := &DataNode{}
	node.SetSession(&sessionutil.Session{SessionRaw: sessionutil.SessionRaw{ServerID: 1}})
	node.flowgraphManager = newFlowgraphManager()
	// server is closed
	node.stateCode.Store(commonpb.StateCode_Abnormal)
	resp, err := node.GetMetrics(s.ctx, &milvuspb.GetMetricsRequest{})
	s.Assert().NoError(err)
	s.Assert().False(merr.Ok(resp.GetStatus()))

	node.stateCode.Store(commonpb.StateCode_Healthy)

	// failed to parse metric type
	invalidRequest := "invalid request"
	resp, err = node.GetMetrics(s.ctx, &milvuspb.GetMetricsRequest{
		Request: invalidRequest,
	})
	s.Assert().NoError(err)
	s.Assert().False(merr.Ok(resp.GetStatus()))

	// unsupported metric type
	unsupportedMetricType := "unsupported"
	req, err := metricsinfo.ConstructRequestByMetricType(unsupportedMetricType)
	s.Assert().NoError(err)
	resp, err = node.GetMetrics(s.ctx, req)
	s.Assert().NoError(err)
	s.Assert().False(merr.Ok(resp.GetStatus()))

	// normal case
	req, err = metricsinfo.ConstructRequestByMetricType(metricsinfo.SystemInfoMetrics)
	s.Assert().NoError(err)
	resp, err = node.GetMetrics(node.ctx, req)
	s.Assert().NoError(err)
	s.Assert().True(merr.Ok(resp.GetStatus()))
	log.Info("Test DataNode.GetMetrics",
		zap.String("name", resp.ComponentName),
		zap.String("response", resp.Response))
}

func (s *DataNodeServicesSuite) TestResendSegmentStats() {
	req := &datapb.ResendSegmentStatsRequest{
		Base: &commonpb.MsgBase{},
	}

	resp, err := s.node.ResendSegmentStats(s.ctx, req)
	s.Assert().NoError(err, "empty call, no error")
	s.Assert().True(merr.Ok(resp.GetStatus()), "empty call, status shall be OK")
}

func (s *DataNodeServicesSuite) TestRPCWatch() {
	s.Run("node not healthy", func() {
		s.SetupTest()
		s.node.UpdateStateCode(commonpb.StateCode_Abnormal)

		ctx := context.Background()
		status, err := s.node.NotifyChannelOperation(ctx, nil)
		s.NoError(err)
		s.False(merr.Ok(status))
		s.ErrorIs(merr.Error(status), merr.ErrServiceNotReady)

		resp, err := s.node.CheckChannelOperationProgress(ctx, nil)
		s.NoError(err)
		s.False(merr.Ok(resp.GetStatus()))
		s.ErrorIs(merr.Error(status), merr.ErrServiceNotReady)
	})

	s.Run("submit error", func() {
		s.SetupTest()
		ctx := context.Background()
		status, err := s.node.NotifyChannelOperation(ctx, &datapb.ChannelOperationsRequest{Infos: []*datapb.ChannelWatchInfo{{OpID: 19530}}})
		s.NoError(err)
		s.False(merr.Ok(status))
		s.NotErrorIs(merr.Error(status), merr.ErrServiceNotReady)

		resp, err := s.node.CheckChannelOperationProgress(ctx, nil)
		s.NoError(err)
		s.False(merr.Ok(resp.GetStatus()))
	})
}

func (s *DataNodeServicesSuite) TestQuerySlot() {
	s.Run("node not healthy", func() {
		s.SetupTest()
		s.node.UpdateStateCode(commonpb.StateCode_Abnormal)

		ctx := context.Background()
		resp, err := s.node.QuerySlot(ctx, nil)
		s.NoError(err)
		s.False(merr.Ok(resp.GetStatus()))
		s.ErrorIs(merr.Error(resp.GetStatus()), merr.ErrServiceNotReady)
	})

	s.Run("normal case", func() {
		s.SetupTest()
		ctx := context.Background()
		resp, err := s.node.QuerySlot(ctx, nil)
		s.NoError(err)
		s.True(merr.Ok(resp.GetStatus()))
		s.NoError(merr.Error(resp.GetStatus()))
	})
}
