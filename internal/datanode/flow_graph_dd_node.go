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

package datanode

import (
	"sync"

	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/msgstream"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/util/flowgraph"
	"github.com/milvus-io/milvus/internal/util/trace"
	"github.com/opentracing/opentracing-go"
)

// make sure ddNode implements flowgraph.Node
var _ flowgraph.Node = (*ddNode)(nil)

// ddNode filter messages from message streams.
//
// ddNode recives all the messages from message stream dml channels, including insert messages,
//  delete messages and ddl messages like CreateCollectionMsg.
//
// ddNode filters insert messages according to the `flushedSegment` and `FilterThreshold`.
//  If the timestamp of the insert message is earlier than `FilterThreshold`, ddNode will
//  filter out the insert message for those who belong to `flushedSegment`
//
// When receiving a `DropCollection` message, ddNode will send a signal to DataNode `BackgroundGC`
//  goroutinue, telling DataNode to release the resources of this perticular flow graph.
//
// After the filtering process, ddNode passes all the valid insert messages and delete message
//  to the following flow graph node, which in DataNode is `insertBufferNode`
type ddNode struct {
	BaseNode

	clearSignal  chan<- UniqueID
	collectionID UniqueID

	segID2SegInfo   sync.Map // segment ID to *SegmentInfo
	flushedSegments []*datapb.SegmentInfo
}

// Name returns node name, implementing flowgraph.Node
func (ddn *ddNode) Name() string {
	return "ddNode"
}

// Operate handles input messages, implementing flowgrpah.Node
func (ddn *ddNode) Operate(in []Msg) []Msg {
	// log.Debug("DDNode Operating")

	if len(in) != 1 {
		log.Warn("Invalid operate message input in ddNode", zap.Int("input length", len(in)))
		return []Msg{}
	}

	msMsg, ok := in[0].(*MsgStreamMsg)
	if !ok {
		log.Warn("Type assertion failed for MsgStreamMsg")
		return []Msg{}
	}

	var spans []opentracing.Span
	for _, msg := range msMsg.TsMessages() {
		sp, ctx := trace.StartSpanFromContext(msg.TraceCtx())
		spans = append(spans, sp)
		msg.SetTraceCtx(ctx)
	}

	var fgMsg = flowGraphMsg{
		insertMessages: make([]*msgstream.InsertMsg, 0),
		timeRange: TimeRange{
			timestampMin: msMsg.TimestampMin(),
			timestampMax: msMsg.TimestampMax(),
		},
		startPositions: make([]*internalpb.MsgPosition, 0),
		endPositions:   make([]*internalpb.MsgPosition, 0),
	}

	for _, msg := range msMsg.TsMessages() {
		switch msg.Type() {
		case commonpb.MsgType_DropCollection:
			if msg.(*msgstream.DropCollectionMsg).GetCollectionID() == ddn.collectionID {
				log.Info("Destroying current flowgraph", zap.Any("collectionID", ddn.collectionID))
				ddn.clearSignal <- ddn.collectionID
				return []Msg{}
			}
		case commonpb.MsgType_Insert:
			log.Debug("DDNode receive insert messages")
			imsg := msg.(*msgstream.InsertMsg)
			if imsg.CollectionID != ddn.collectionID {
				//log.Debug("filter invalid InsertMsg, collection mis-match",
				//	zap.Int64("Get msg collID", imsg.CollectionID),
				//	zap.Int64("Expected collID", ddn.collectionID))
				continue
			}
			if msg.EndTs() < FilterThreshold {
				log.Info("Filtering Insert Messages",
					zap.Uint64("Message endts", msg.EndTs()),
					zap.Uint64("FilterThreshold", FilterThreshold),
				)
				if ddn.filterFlushedSegmentInsertMessages(imsg) {
					continue
				}
			}
			fgMsg.insertMessages = append(fgMsg.insertMessages, imsg)
		case commonpb.MsgType_Delete:
			log.Debug("DDNode receive delete messages")
			dmsg := msg.(*msgstream.DeleteMsg)
			if dmsg.CollectionID != ddn.collectionID {
				//log.Debug("filter invalid DeleteMsg, collection mis-match",
				//	zap.Int64("Get msg collID", dmsg.CollectionID),
				//	zap.Int64("Expected collID", ddn.collectionID))
				continue
			}
			fgMsg.deleteMessages = append(fgMsg.deleteMessages, dmsg)
		}
	}

	fgMsg.startPositions = append(fgMsg.startPositions, msMsg.StartPositions()...)
	fgMsg.endPositions = append(fgMsg.endPositions, msMsg.EndPositions()...)

	for _, sp := range spans {
		sp.Finish()
	}

	return []Msg{&fgMsg}
}

func (ddn *ddNode) filterFlushedSegmentInsertMessages(msg *msgstream.InsertMsg) bool {
	if ddn.isFlushed(msg.GetSegmentID()) {
		return true
	}

	if si, ok := ddn.segID2SegInfo.Load(msg.GetSegmentID()); ok {
		if msg.EndTs() <= si.(*datapb.SegmentInfo).GetDmlPosition().GetTimestamp() {
			return true
		}

		ddn.segID2SegInfo.Delete(msg.GetSegmentID())
	}
	return false
}

func (ddn *ddNode) isFlushed(segmentID UniqueID) bool {
	for _, s := range ddn.flushedSegments {
		if s.ID == segmentID {
			return true
		}
	}
	return false
}

func newDDNode(clearSignal chan<- UniqueID, collID UniqueID, vchanInfo *datapb.VchannelInfo) *ddNode {
	baseNode := BaseNode{}
	baseNode.SetMaxParallelism(Params.FlowGraphMaxQueueLength)

	fs := make([]*datapb.SegmentInfo, 0, len(vchanInfo.GetFlushedSegments()))
	fs = append(fs, vchanInfo.GetFlushedSegments()...)
	log.Debug("ddNode add flushed segment",
		zap.Int64("collectionID", vchanInfo.GetCollectionID()),
		zap.Int("No. Segment", len(vchanInfo.GetFlushedSegments())),
	)

	dd := &ddNode{
		BaseNode:        baseNode,
		clearSignal:     clearSignal,
		collectionID:    collID,
		flushedSegments: fs,
	}

	for _, us := range vchanInfo.GetUnflushedSegments() {
		dd.segID2SegInfo.Store(us.GetID(), us)
	}

	log.Debug("ddNode add unflushed segment",
		zap.Int64("collectionID", collID),
		zap.Int("No. Segment", len(vchanInfo.GetUnflushedSegments())),
	)

	return dd
}
