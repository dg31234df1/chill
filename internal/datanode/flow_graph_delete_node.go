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
	"fmt"
	"reflect"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/mq/msgstream"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/internal/util/retry"
	"github.com/milvus-io/milvus/internal/util/trace"
	"github.com/milvus-io/milvus/internal/util/tsoutil"
)

// DeleteNode is to process delete msg, flush delete info into storage.
type deleteNode struct {
	BaseNode
	ctx              context.Context
	channelName      string
	delBufferManager *DelBufferManager // manager of delete msg
	channel          Channel
	idAllocator      allocatorInterface
	flushManager     flushManager

	clearSignal chan<- string
}

func (dn *deleteNode) Name() string {
	return "deleteNode-" + dn.channelName
}

func (dn *deleteNode) Close() {
	log.Info("Flowgraph Delete Node closing")
}

func (dn *deleteNode) showDelBuf(segIDs []UniqueID, ts Timestamp) {
	for _, segID := range segIDs {
		if _, ok := dn.delBufferManager.Load(segID); ok {
			log.Debug("delta buffer status",
				zap.Uint64("timestamp", ts),
				zap.Int64("segment ID", segID),
				zap.Int64("entries", dn.delBufferManager.GetEntriesNum(segID)),
				zap.Int64("memory size", dn.delBufferManager.GetSegDelBufMemSize(segID)),
				zap.String("vChannel", dn.channelName))
		}
	}
}

// Operate implementing flowgraph.Node, performs delete data process
func (dn *deleteNode) Operate(in []Msg) []Msg {
	if in == nil {
		log.Debug("type assertion failed for flowGraphMsg because it's nil")
		return []Msg{}
	}

	if len(in) != 1 {
		log.Error("Invalid operate message input in deleteNode", zap.Int("input length", len(in)))
		return []Msg{}
	}

	fgMsg, ok := in[0].(*flowGraphMsg)
	if !ok {
		log.Warn("type assertion failed for flowGraphMsg", zap.String("name", reflect.TypeOf(in[0]).Name()))
		return []Msg{}
	}

	var spans []opentracing.Span
	for _, msg := range fgMsg.deleteMessages {
		sp, ctx := trace.StartSpanFromContext(msg.TraceCtx())
		spans = append(spans, sp)
		msg.SetTraceCtx(ctx)
	}

	// update compacted segment before operation
	if len(fgMsg.deleteMessages) > 0 || len(fgMsg.segmentsToSync) > 0 {
		dn.updateCompactedSegments()
	}

	// process delete messages
	var segIDs []UniqueID
	for i, msg := range fgMsg.deleteMessages {
		traceID, _, _ := trace.InfoFromSpan(spans[i])
		log.Debug("Buffer delete request in DataNode", zap.String("traceID", traceID))
		tmpSegIDs, err := dn.bufferDeleteMsg(msg, fgMsg.timeRange, fgMsg.startPositions[0], fgMsg.endPositions[0])
		if err != nil {
			// error occurs only when deleteMsg is misaligned, should not happen
			err = fmt.Errorf("buffer delete msg failed, err = %s", err)
			log.Error(err.Error())
			panic(err)
		}
		segIDs = append(segIDs, tmpSegIDs...)
	}

	// display changed segment's status in dn.delBuf of a certain ts
	if len(fgMsg.deleteMessages) != 0 {
		dn.showDelBuf(segIDs, fgMsg.timeRange.timestampMax)
	}

	//here we adopt a quite radical strategy:
	//every time we make sure that the N biggest delDataBuf can be flushed
	//when memsize usage reaches a certain level
	//then we will add all segments in the fgMsg.segmentsToFlush into the toFlushSeg and remove duplicate segments
	//the aim for taking all these actions is to guarantee that the memory consumed by delBuf will not exceed a limit
	segmentsToFlush := dn.delBufferManager.ShouldFlushSegments()
	log.Info("should flush segments, ", zap.Int("seg_count", len(segmentsToFlush)))
	for _, msgSegmentID := range fgMsg.segmentsToSync {
		existed := false
		for _, autoFlushSegment := range segmentsToFlush {
			if msgSegmentID == autoFlushSegment {
				existed = true
				break
			}
		}
		if !existed {
			segmentsToFlush = append(segmentsToFlush, msgSegmentID)
		}
	}

	// process flush messages
	if len(segmentsToFlush) > 0 {
		log.Debug("DeleteNode receives flush message",
			zap.Int64s("segIDs", segmentsToFlush),
			zap.String("vChannelName", dn.channelName),
			zap.Time("posTime", tsoutil.PhysicalTime(fgMsg.endPositions[0].Timestamp)))
		for _, segmentToFlush := range segmentsToFlush {
			buf, ok := dn.delBufferManager.Load(segmentToFlush)
			if !ok {
				// no related delta data to flush, send empty buf to complete flush life-cycle
				dn.flushManager.flushDelData(nil, segmentToFlush, fgMsg.endPositions[0])
			} else {
				err := retry.Do(dn.ctx, func() error {
					return dn.flushManager.flushDelData(buf, segmentToFlush, fgMsg.endPositions[0])
				}, getFlowGraphRetryOpt())
				if err != nil {
					err = fmt.Errorf("failed to flush delete data, err = %s", err)
					log.Error(err.Error())
					panic(err)
				}
				// remove delete buf
				dn.delBufferManager.Delete(segmentToFlush)
			}
		}
	}

	// process drop collection message, delete node shall notify flush manager all data are cleared and send signal to DataSyncService cleaner
	if fgMsg.dropCollection {
		dn.flushManager.notifyAllFlushed()
		log.Info("DeleteNode notifies BackgroundGC to release vchannel", zap.String("vChannelName", dn.channelName))
		dn.clearSignal <- dn.channelName
	}

	for _, sp := range spans {
		sp.Finish()
	}
	return in
}

// update delBuf for compacted segments
func (dn *deleteNode) updateCompactedSegments() {
	compactedTo2From := dn.channel.listCompactedSegmentIDs()

	for compactedTo, compactedFrom := range compactedTo2From {
		// if the compactedTo segment has 0 numRows, remove all segments related
		if !dn.channel.hasSegment(compactedTo, true) {
			for _, segID := range compactedFrom {
				dn.delBufferManager.Delete(segID)
			}
			dn.channel.removeSegments(compactedFrom...)
			continue
		}

		dn.delBufferManager.CompactSegBuf(compactedTo, compactedFrom)
		log.Info("update delBuf for compacted segments",
			zap.Int64("compactedTo segmentID", compactedTo),
			zap.Int64s("compactedFrom segmentIDs", compactedFrom),
		)
		dn.channel.removeSegments(compactedFrom...)
	}
}

func (dn *deleteNode) bufferDeleteMsg(msg *msgstream.DeleteMsg, tr TimeRange, startPos, endPos *internalpb.MsgPosition) ([]UniqueID, error) {
	log.Debug("bufferDeleteMsg", zap.Any("primary keys", msg.PrimaryKeys), zap.String("vChannelName", dn.channelName))

	primaryKeys := storage.ParseIDs2PrimaryKeys(msg.PrimaryKeys)
	segIDToPks, segIDToTss := dn.filterSegmentByPK(msg.PartitionID, primaryKeys, msg.Timestamps)

	segIDs := make([]UniqueID, 0, len(segIDToPks))
	for segID, pks := range segIDToPks {
		segIDs = append(segIDs, segID)

		tss, ok := segIDToTss[segID]
		if !ok || len(pks) != len(tss) {
			return nil, fmt.Errorf("primary keys and timestamp's element num mis-match, segmentID = %d", segID)
		}
		dn.delBufferManager.StoreNewDeletes(segID, pks, tss, tr, startPos, endPos)
	}

	return segIDs, nil
}

// filterSegmentByPK returns the bloom filter check result.
// If the key may exist in the segment, returns it in map.
// If the key not exist in the segment, the segment is filter out.
func (dn *deleteNode) filterSegmentByPK(partID UniqueID, pks []primaryKey, tss []Timestamp) (
	map[UniqueID][]primaryKey, map[UniqueID][]uint64) {
	segID2Pks := make(map[UniqueID][]primaryKey)
	segID2Tss := make(map[UniqueID][]uint64)
	segments := dn.channel.filterSegments(partID)
	for index, pk := range pks {
		for _, segment := range segments {
			segmentID := segment.segmentID
			if segment.isPKExist(pk) {
				segID2Pks[segmentID] = append(segID2Pks[segmentID], pk)
				segID2Tss[segmentID] = append(segID2Tss[segmentID], tss[index])
			}
		}
	}

	return segID2Pks, segID2Tss
}

func newDeleteNode(ctx context.Context, fm flushManager, sig chan<- string, config *nodeConfig) (*deleteNode, error) {
	baseNode := BaseNode{}
	baseNode.SetMaxQueueLength(config.maxQueueLength)
	baseNode.SetMaxParallelism(config.maxParallelism)

	return &deleteNode{
		ctx:      ctx,
		BaseNode: baseNode,
		delBufferManager: &DelBufferManager{
			channel:       config.channel,
			delMemorySize: 0,
			delBufHeap:    &PriorityQueue{},
		},
		channel:      config.channel,
		idAllocator:  config.allocator,
		channelName:  config.vChannelName,
		flushManager: fm,
		clearSignal:  sig,
	}, nil
}
