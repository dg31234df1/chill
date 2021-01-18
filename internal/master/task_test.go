package master

import (
	"testing"

	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"

	"github.com/stretchr/testify/assert"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
)

func TestMaster_CreateCollectionTask(t *testing.T) {
	req := milvuspb.CreateCollectionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kCreateCollection,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
		Schema: nil,
	}
	var collectionTask task = &createCollectionTask{
		req:      &req,
		baseTask: baseTask{},
	}
	assert.Equal(t, commonpb.MsgType_kCreateCollection, collectionTask.Type())
	ts, err := collectionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	collectionTask = &createCollectionTask{
		req:      nil,
		baseTask: baseTask{},
	}

	assert.Equal(t, commonpb.MsgType_kNone, collectionTask.Type())
	ts, err = collectionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = collectionTask.Execute()
	assert.NotNil(t, err)
}

func TestMaster_DropCollectionTask(t *testing.T) {
	req := milvuspb.DropCollectionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kDropPartition,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
	}
	var collectionTask task = &dropCollectionTask{
		req:        &req,
		baseTask:   baseTask{},
		segManager: NewMockSegmentManager(),
	}
	assert.Equal(t, commonpb.MsgType_kDropPartition, collectionTask.Type())
	ts, err := collectionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	collectionTask = &dropCollectionTask{
		req:        nil,
		baseTask:   baseTask{},
		segManager: NewMockSegmentManager(),
	}

	assert.Equal(t, commonpb.MsgType_kNone, collectionTask.Type())
	ts, err = collectionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = collectionTask.Execute()
	assert.NotNil(t, err)
}

func TestMaster_HasCollectionTask(t *testing.T) {
	req := milvuspb.HasCollectionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kHasCollection,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
	}
	var collectionTask task = &hasCollectionTask{
		req:      &req,
		baseTask: baseTask{},
	}
	assert.Equal(t, commonpb.MsgType_kHasCollection, collectionTask.Type())
	ts, err := collectionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	collectionTask = &hasCollectionTask{
		req:      nil,
		baseTask: baseTask{},
	}

	assert.Equal(t, commonpb.MsgType_kNone, collectionTask.Type())
	ts, err = collectionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = collectionTask.Execute()
	assert.NotNil(t, err)
}

func TestMaster_ShowCollectionTask(t *testing.T) {
	req := milvuspb.ShowCollectionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kShowCollections,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
	}
	var collectionTask task = &showCollectionsTask{
		req:      &req,
		baseTask: baseTask{},
	}
	assert.Equal(t, commonpb.MsgType_kShowCollections, collectionTask.Type())
	ts, err := collectionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	collectionTask = &showCollectionsTask{
		req:      nil,
		baseTask: baseTask{},
	}

	assert.Equal(t, commonpb.MsgType_kNone, collectionTask.Type())
	ts, err = collectionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = collectionTask.Execute()
	assert.NotNil(t, err)
}

func TestMaster_DescribeCollectionTask(t *testing.T) {
	req := milvuspb.DescribeCollectionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kDescribeCollection,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
	}
	var collectionTask task = &describeCollectionTask{
		req:      &req,
		baseTask: baseTask{},
	}
	assert.Equal(t, commonpb.MsgType_kDescribeCollection, collectionTask.Type())
	ts, err := collectionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	collectionTask = &describeCollectionTask{
		req:      nil,
		baseTask: baseTask{},
	}

	assert.Equal(t, commonpb.MsgType_kNone, collectionTask.Type())
	ts, err = collectionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = collectionTask.Execute()
	assert.NotNil(t, err)
}

func TestMaster_CreatePartitionTask(t *testing.T) {
	req := milvuspb.CreatePartitionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kCreatePartition,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
	}
	var partitionTask task = &createPartitionTask{
		req:      &req,
		baseTask: baseTask{},
	}
	assert.Equal(t, commonpb.MsgType_kCreatePartition, partitionTask.Type())
	ts, err := partitionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	partitionTask = &createPartitionTask{
		req:      nil,
		baseTask: baseTask{},
	}

	assert.Equal(t, commonpb.MsgType_kNone, partitionTask.Type())
	ts, err = partitionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = partitionTask.Execute()
	assert.NotNil(t, err)
}
func TestMaster_DropPartitionTask(t *testing.T) {
	req := milvuspb.DropPartitionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kDropPartition,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
	}
	var partitionTask task = &dropPartitionTask{
		req:      &req,
		baseTask: baseTask{},
	}
	assert.Equal(t, commonpb.MsgType_kDropPartition, partitionTask.Type())
	ts, err := partitionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	partitionTask = &dropPartitionTask{
		req:      nil,
		baseTask: baseTask{},
	}

	assert.Equal(t, commonpb.MsgType_kNone, partitionTask.Type())
	ts, err = partitionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = partitionTask.Execute()
	assert.NotNil(t, err)
}
func TestMaster_HasPartitionTask(t *testing.T) {
	req := milvuspb.HasPartitionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kHasPartition,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
	}
	var partitionTask task = &hasPartitionTask{
		req:      &req,
		baseTask: baseTask{},
	}
	assert.Equal(t, commonpb.MsgType_kHasPartition, partitionTask.Type())
	ts, err := partitionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	partitionTask = &hasPartitionTask{
		req:      nil,
		baseTask: baseTask{},
	}

	assert.Equal(t, commonpb.MsgType_kNone, partitionTask.Type())
	ts, err = partitionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = partitionTask.Execute()
	assert.NotNil(t, err)
}

//func TestMaster_DescribePartitionTask(t *testing.T) {
//	req := milvuspb.DescribePartitionRequest{
//		MsgType:       commonpb.MsgType_kDescribePartition,
//		ReqID:         1,
//		Timestamp:     11,
//		ProxyID:       1,
//		PartitionName: nil,
//	}
//	var partitionTask task = &describePartitionTask{
//		req:      &req,
//		baseTask: baseTask{},
//	}
//	assert.Equal(t, commonpb.MsgType_kDescribePartition, partitionTask.Type())
//	ts, err := partitionTask.Ts()
//	assert.Equal(t, uint64(11), ts)
//	assert.Nil(t, err)
//
//	partitionTask = &describePartitionTask{
//		req:      nil,
//		baseTask: baseTask{},
//	}
//
//	assert.Equal(t, commonpb.MsgType_kNone, partitionTask.Type())
//	ts, err = partitionTask.Ts()
//	assert.Equal(t, uint64(0), ts)
//	assert.NotNil(t, err)
//	err = partitionTask.Execute()
//	assert.NotNil(t, err)
//}

func TestMaster_ShowPartitionTask(t *testing.T) {
	req := milvuspb.ShowPartitionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kShowPartitions,
			MsgID:     1,
			Timestamp: 11,
			SourceID:  1,
		},
	}
	var partitionTask task = &showPartitionTask{
		req:      &req,
		baseTask: baseTask{},
	}
	assert.Equal(t, commonpb.MsgType_kShowPartitions, partitionTask.Type())
	ts, err := partitionTask.Ts()
	assert.Equal(t, uint64(11), ts)
	assert.Nil(t, err)

	partitionTask = &showPartitionTask{
		req:      nil,
		baseTask: baseTask{},
	}

	assert.Equal(t, commonpb.MsgType_kNone, partitionTask.Type())
	ts, err = partitionTask.Ts()
	assert.Equal(t, uint64(0), ts)
	assert.NotNil(t, err)
	err = partitionTask.Execute()
	assert.NotNil(t, err)
}
