package timestamp

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/atomic"
	"google.golang.org/grpc"

	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus/internal/mocks"
	"github.com/milvus-io/milvus/internal/proto/rootcoordpb"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
)

func TestLocalAllocator(t *testing.T) {
	allocator := newLocalAllocator()

	ts, err := allocator.allocateOne()
	assert.Error(t, err)
	assert.Zero(t, ts)

	allocator.update(1, 100)

	counter := atomic.NewUint64(0)
	for i := 0; i < 100; i++ {
		ts, err := allocator.allocateOne()
		assert.NoError(t, err)
		assert.NotZero(t, ts)
		counter.Add(ts)
	}
	assert.Equal(t, uint64(5050), counter.Load())

	// allocator exhausted.
	ts, err = allocator.allocateOne()
	assert.Error(t, err)
	assert.Zero(t, ts)

	// allocator can not be rollback.
	allocator.update(90, 100)
	ts, err = allocator.allocateOne()
	assert.Error(t, err)
	assert.Zero(t, ts)

	// allocator can be only increasing.
	allocator.update(101, 100)
	ts, err = allocator.allocateOne()
	assert.NoError(t, err)
	assert.Equal(t, ts, uint64(101))

	// allocator can be exhausted.
	allocator.exhausted()
	ts, err = allocator.allocateOne()
	assert.Error(t, err)
	assert.Zero(t, ts)
}

func TestRemoteAllocator(t *testing.T) {
	paramtable.Init()
	paramtable.SetNodeID(1)

	client := NewMockRootCoordClient(t)

	allocator := newRemoteAllocator(client)
	ts, count, err := allocator.allocate(context.Background(), 100)
	assert.NoError(t, err)
	assert.NotZero(t, ts)
	assert.Equal(t, count, 100)

	// Test error.
	client = mocks.NewMockRootCoordClient(t)
	client.EXPECT().AllocTimestamp(mock.Anything, mock.Anything).RunAndReturn(
		func(ctx context.Context, atr *rootcoordpb.AllocTimestampRequest, co ...grpc.CallOption) (*rootcoordpb.AllocTimestampResponse, error) {
			return nil, errors.New("test")
		},
	)
	allocator = newRemoteAllocator(client)
	_, _, err = allocator.allocate(context.Background(), 100)
	assert.Error(t, err)

	client.EXPECT().AllocTimestamp(mock.Anything, mock.Anything).Unset()
	client.EXPECT().AllocTimestamp(mock.Anything, mock.Anything).RunAndReturn(
		func(ctx context.Context, atr *rootcoordpb.AllocTimestampRequest, co ...grpc.CallOption) (*rootcoordpb.AllocTimestampResponse, error) {
			return &rootcoordpb.AllocTimestampResponse{
				Status: &commonpb.Status{
					ErrorCode: commonpb.ErrorCode_ForceDeny,
				},
			}, nil
		},
	)
	allocator = newRemoteAllocator(client)
	_, _, err = allocator.allocate(context.Background(), 100)
	assert.Error(t, err)
}
