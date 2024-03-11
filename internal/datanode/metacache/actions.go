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

package metacache

import (
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

type SegmentFilter interface {
	Filter(info *SegmentInfo) bool
	SegmentIDs() ([]int64, bool)
}

type SegmentIDFilter struct {
	segmentIDs []int64
	ids        typeutil.Set[int64]
}

func WithSegmentIDs(segmentIDs ...int64) SegmentFilter {
	set := typeutil.NewSet(segmentIDs...)
	return &SegmentIDFilter{
		segmentIDs: segmentIDs,
		ids:        set,
	}
}

func (f *SegmentIDFilter) Filter(info *SegmentInfo) bool {
	return f.ids.Contain(info.segmentID)
}

func (f *SegmentIDFilter) SegmentIDs() ([]int64, bool) {
	return f.segmentIDs, true
}

type SegmentFilterFunc func(info *SegmentInfo) bool

func (f SegmentFilterFunc) Filter(info *SegmentInfo) bool {
	return f(info)
}

func (f SegmentFilterFunc) SegmentIDs() ([]int64, bool) {
	return nil, false
}

func WithPartitionID(partitionID int64) SegmentFilter {
	return SegmentFilterFunc(func(info *SegmentInfo) bool {
		return partitionID == common.InvalidPartitionID || info.partitionID == partitionID
	})
}

func WithSegmentState(states ...commonpb.SegmentState) SegmentFilter {
	set := typeutil.NewSet(states...)
	return SegmentFilterFunc(func(info *SegmentInfo) bool {
		return set.Len() > 0 && set.Contain(info.state)
	})
}

func WithStartPosNotRecorded() SegmentFilter {
	return SegmentFilterFunc(func(info *SegmentInfo) bool {
		return !info.startPosRecorded
	})
}

func WithLevel(level datapb.SegmentLevel) SegmentFilter {
	return SegmentFilterFunc(func(info *SegmentInfo) bool {
		return info.level == level
	})
}

func WithCompacted() SegmentFilter {
	return SegmentFilterFunc(func(info *SegmentInfo) bool {
		return info.compactTo != 0
	})
}

func WithNoSyncingTask() SegmentFilter {
	return SegmentFilterFunc(func(info *SegmentInfo) bool {
		return info.syncingTasks == 0
	})
}

type SegmentAction func(info *SegmentInfo)

func UpdateState(state commonpb.SegmentState) SegmentAction {
	return func(info *SegmentInfo) {
		info.state = state
	}
}

func UpdateCheckpoint(checkpoint *msgpb.MsgPosition) SegmentAction {
	return func(info *SegmentInfo) {
		info.checkpoint = checkpoint
	}
}

func UpdateNumOfRows(numOfRows int64) SegmentAction {
	return func(info *SegmentInfo) {
		info.flushedRows = numOfRows
	}
}

func UpdateBufferedRows(bufferedRows int64) SegmentAction {
	return func(info *SegmentInfo) {
		info.bufferRows = bufferedRows
	}
}

func RollStats(newStats ...*storage.PrimaryKeyStats) SegmentAction {
	return func(info *SegmentInfo) {
		info.bfs.Roll(newStats...)
	}
}

func CompactTo(compactTo int64) SegmentAction {
	return func(info *SegmentInfo) {
		info.compactTo = compactTo
	}
}

func StartSyncing(batchSize int64) SegmentAction {
	return func(info *SegmentInfo) {
		info.syncingRows += batchSize
		info.bufferRows -= batchSize
		info.syncingTasks++
	}
}

func FinishSyncing(batchSize int64) SegmentAction {
	return func(info *SegmentInfo) {
		info.flushedRows += batchSize
		info.syncingRows -= batchSize
		info.syncingTasks--
	}
}

func SetStartPosRecorded(flag bool) SegmentAction {
	return func(info *SegmentInfo) {
		info.startPosRecorded = flag
	}
}

// MergeSegmentAction is the util function to merge multiple SegmentActions into one.
func MergeSegmentAction(actions ...SegmentAction) SegmentAction {
	return func(info *SegmentInfo) {
		for _, action := range actions {
			action(info)
		}
	}
}
