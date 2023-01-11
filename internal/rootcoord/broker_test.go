package rootcoord

import (
	"context"
	"errors"
	"testing"

	"github.com/milvus-io/milvus/internal/proto/indexpb"

	"github.com/milvus-io/milvus/internal/metastore/model"

	"github.com/milvus-io/milvus-proto/go-api/milvuspb"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/stretchr/testify/assert"
)

func TestServerBroker_ReleaseCollection(t *testing.T) {
	t.Run("failed to execute", func(t *testing.T) {
		c := newTestCore(withInvalidQueryCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.ReleaseCollection(ctx, 1)
		assert.Error(t, err)
	})

	t.Run("non success error code on execute", func(t *testing.T) {
		c := newTestCore(withFailedQueryCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.ReleaseCollection(ctx, 1)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		c := newTestCore(withValidQueryCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.ReleaseCollection(ctx, 1)
		assert.NoError(t, err)
	})
}

func TestServerBroker_GetSegmentInfo(t *testing.T) {
	t.Run("failed to execute", func(t *testing.T) {
		c := newTestCore(withInvalidQueryCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		_, err := b.GetQuerySegmentInfo(ctx, 1, []int64{1, 2})
		assert.Error(t, err)
	})

	t.Run("non success error code on execute", func(t *testing.T) {
		c := newTestCore(withFailedQueryCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		resp, err := b.GetQuerySegmentInfo(ctx, 1, []int64{1, 2})
		assert.NoError(t, err)
		assert.Equal(t, commonpb.ErrorCode_UnexpectedError, resp.GetStatus().GetErrorCode())
	})

	t.Run("success", func(t *testing.T) {
		c := newTestCore(withValidQueryCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		resp, err := b.GetQuerySegmentInfo(ctx, 1, []int64{1, 2})
		assert.NoError(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.GetStatus().GetErrorCode())
	})
}

func TestServerBroker_WatchChannels(t *testing.T) {
	t.Run("failed to execute", func(t *testing.T) {
		defer cleanTestEnv()

		c := newTestCore(withInvalidDataCoord(), withRocksMqTtSynchronizer())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.WatchChannels(ctx, &watchInfo{})
		assert.Error(t, err)
	})

	t.Run("non success error code on execute", func(t *testing.T) {
		defer cleanTestEnv()

		c := newTestCore(withFailedDataCoord(), withRocksMqTtSynchronizer())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.WatchChannels(ctx, &watchInfo{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		defer cleanTestEnv()

		c := newTestCore(withValidDataCoord(), withRocksMqTtSynchronizer())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.WatchChannels(ctx, &watchInfo{})
		assert.NoError(t, err)
	})
}

func TestServerBroker_UnwatchChannels(t *testing.T) {
	// TODO: implement
	b := newServerBroker(newTestCore())
	ctx := context.Background()
	b.UnwatchChannels(ctx, &watchInfo{})
}

func TestServerBroker_Flush(t *testing.T) {
	t.Run("failed to execute", func(t *testing.T) {
		c := newTestCore(withInvalidDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.Flush(ctx, 1, []int64{1, 2})
		assert.Error(t, err)
	})

	t.Run("non success error code on execute", func(t *testing.T) {
		c := newTestCore(withFailedDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.Flush(ctx, 1, []int64{1, 2})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		c := newTestCore(withValidDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.Flush(ctx, 1, []int64{1, 2})
		assert.NoError(t, err)
	})
}

func TestServerBroker_Import(t *testing.T) {
	t.Run("failed to execute", func(t *testing.T) {
		c := newTestCore(withInvalidDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		resp, err := b.Import(ctx, &datapb.ImportTaskRequest{})
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("non success error code on execute", func(t *testing.T) {
		c := newTestCore(withFailedDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		resp, err := b.Import(ctx, &datapb.ImportTaskRequest{})
		assert.NoError(t, err)
		assert.Equal(t, commonpb.ErrorCode_UnexpectedError, resp.GetStatus().GetErrorCode())
	})

	t.Run("success", func(t *testing.T) {
		c := newTestCore(withValidDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		resp, err := b.Import(ctx, &datapb.ImportTaskRequest{})
		assert.NoError(t, err)
		assert.Equal(t, commonpb.ErrorCode_Success, resp.GetStatus().GetErrorCode())
	})
}

func TestServerBroker_DropCollectionIndex(t *testing.T) {
	t.Run("failed to execute", func(t *testing.T) {
		c := newTestCore(withInvalidDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.DropCollectionIndex(ctx, 1, nil)
		assert.Error(t, err)
	})

	t.Run("non success error code on execute", func(t *testing.T) {
		c := newTestCore(withFailedDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.DropCollectionIndex(ctx, 1, nil)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		c := newTestCore(withValidDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.DropCollectionIndex(ctx, 1, nil)
		assert.NoError(t, err)
	})
}

func TestServerBroker_GetSegmentIndexState(t *testing.T) {
	t.Run("failed to execute", func(t *testing.T) {
		c := newTestCore(withInvalidDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		_, err := b.GetSegmentIndexState(ctx, 1, "index_name", []UniqueID{1, 2})
		assert.Error(t, err)
	})

	t.Run("non success error code on execute", func(t *testing.T) {
		c := newTestCore(withFailedDataCoord())
		b := newServerBroker(c)
		ctx := context.Background()
		_, err := b.GetSegmentIndexState(ctx, 1, "index_name", []UniqueID{1, 2})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		c := newTestCore(withValidDataCoord())
		c.dataCoord.(*mockDataCoord).GetSegmentIndexStateFunc = func(ctx context.Context, req *indexpb.GetSegmentIndexStateRequest) (*indexpb.GetSegmentIndexStateResponse, error) {
			return &indexpb.GetSegmentIndexStateResponse{
				Status: succStatus(),
				States: []*indexpb.SegmentIndexState{
					{
						SegmentID:  1,
						State:      commonpb.IndexState_Finished,
						FailReason: "",
					},
				},
			}, nil
		}
		b := newServerBroker(c)
		ctx := context.Background()
		states, err := b.GetSegmentIndexState(ctx, 1, "index_name", []UniqueID{1})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(states))
		assert.Equal(t, commonpb.IndexState_Finished, states[0].GetState())
	})
}

func TestServerBroker_BroadcastAlteredCollection(t *testing.T) {
	collMeta := &model.Collection{
		CollectionID: 1,
		StartPositions: []*commonpb.KeyDataPair{
			{
				Key:  "0",
				Data: []byte("0"),
			},
		},
		Partitions: []*model.Partition{
			{
				PartitionID:               2,
				PartitionName:             "test_partition_name_1",
				PartitionCreatedTimestamp: 0,
			},
		},
	}

	t.Run("get meta fail", func(t *testing.T) {
		c := newTestCore(withInvalidDataCoord())
		c.meta = &mockMetaTable{
			GetCollectionByIDFunc: func(ctx context.Context, collectionID UniqueID, ts Timestamp, allowUnavailable bool) (*model.Collection, error) {
				return nil, errors.New("err")
			},
		}
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.BroadcastAlteredCollection(ctx, &milvuspb.AlterCollectionRequest{})
		assert.Error(t, err)
	})

	t.Run("failed to execute", func(t *testing.T) {
		c := newTestCore(withInvalidDataCoord())
		c.meta = &mockMetaTable{
			GetCollectionByIDFunc: func(ctx context.Context, collectionID UniqueID, ts Timestamp, allowUnavailable bool) (*model.Collection, error) {
				return collMeta, nil
			},
		}
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.BroadcastAlteredCollection(ctx, &milvuspb.AlterCollectionRequest{})
		assert.Error(t, err)
	})

	t.Run("non success error code on execute", func(t *testing.T) {
		c := newTestCore(withFailedDataCoord())
		c.meta = &mockMetaTable{
			GetCollectionByIDFunc: func(ctx context.Context, collectionID UniqueID, ts Timestamp, allowUnavailable bool) (*model.Collection, error) {
				return collMeta, nil
			},
		}
		b := newServerBroker(c)
		ctx := context.Background()
		err := b.BroadcastAlteredCollection(ctx, &milvuspb.AlterCollectionRequest{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		c := newTestCore(withValidDataCoord())
		c.meta = &mockMetaTable{
			GetCollectionByIDFunc: func(ctx context.Context, collectionID UniqueID, ts Timestamp, allowUnavailable bool) (*model.Collection, error) {
				return collMeta, nil
			},
		}
		b := newServerBroker(c)
		ctx := context.Background()

		req := &milvuspb.AlterCollectionRequest{
			CollectionID: 1,
		}
		err := b.BroadcastAlteredCollection(ctx, req)
		assert.NoError(t, err)
	})
}
