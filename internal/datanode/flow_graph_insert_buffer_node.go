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
	"errors"
	"fmt"
	"math"
	"reflect"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/schemapb"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/metrics"
	"github.com/milvus-io/milvus/internal/mq/msgstream"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/internal/util/funcutil"
	"github.com/milvus-io/milvus/internal/util/retry"
	"github.com/milvus-io/milvus/internal/util/trace"
	"github.com/milvus-io/milvus/internal/util/tsoutil"
)

type insertBufferNode struct {
	BaseNode

	ctx          context.Context
	channelName  string
	insertBuffer sync.Map // SegmentID to BufferData
	channel      Channel
	idAllocator  allocatorInterface

	flushMap         sync.Map
	flushChan        <-chan flushMsg
	resendTTChan     <-chan resendTTMsg
	flushingSegCache *Cache
	flushManager     flushManager

	timeTickStream msgstream.MsgStream
	ttLogger       *timeTickLogger
	ttMerger       *mergedTimeTickerSender

	lastTimestamp Timestamp
}

type timeTickLogger struct {
	start        atomic.Uint64
	counter      atomic.Int32
	vChannelName string
}

func (l *timeTickLogger) LogTs(ts Timestamp) {
	if l.counter.Load() == 0 {
		l.start.Store(ts)
	}
	l.counter.Inc()
	if l.counter.Load() == 1000 {
		min := l.start.Load()
		l.start.Store(ts)
		l.counter.Store(0)
		go l.printLogs(min, ts)
	}
}

func (l *timeTickLogger) printLogs(start, end Timestamp) {
	t1, _ := tsoutil.ParseTS(start)
	t2, _ := tsoutil.ParseTS(end)
	log.Debug("IBN timetick log", zap.Time("from", t1), zap.Time("to", t2), zap.Duration("elapsed", t2.Sub(t1)), zap.Uint64("start", start), zap.Uint64("end", end), zap.String("vChannelName", l.vChannelName))
}

type segmentCheckPoint struct {
	numRows int64
	pos     internalpb.MsgPosition
}

func (ibNode *insertBufferNode) Name() string {
	return "ibNode-" + ibNode.channelName
}

func (ibNode *insertBufferNode) Close() {
	ibNode.ttMerger.close()

	if ibNode.timeTickStream != nil {
		ibNode.timeTickStream.Close()
	}
}

func (ibNode *insertBufferNode) Operate(in []Msg) []Msg {
	fgMsg, ok := ibNode.verifyInMsg(in)
	if !ok {
		return []Msg{}
	}

	if fgMsg.dropCollection {
		ibNode.flushManager.startDropping()
	}

	var spans []opentracing.Span
	for _, msg := range fgMsg.insertMessages {
		sp, ctx := trace.StartSpanFromContext(msg.TraceCtx())
		spans = append(spans, sp)
		msg.SetTraceCtx(ctx)
	}

	// replace pchannel with vchannel
	startPositions := make([]*internalpb.MsgPosition, 0, len(fgMsg.startPositions))
	for idx := range fgMsg.startPositions {
		pos := proto.Clone(fgMsg.startPositions[idx]).(*internalpb.MsgPosition)
		pos.ChannelName = ibNode.channelName
		startPositions = append(startPositions, pos)
	}
	endPositions := make([]*internalpb.MsgPosition, 0, len(fgMsg.endPositions))
	for idx := range fgMsg.endPositions {
		pos := proto.Clone(fgMsg.endPositions[idx]).(*internalpb.MsgPosition)
		pos.ChannelName = ibNode.channelName
		endPositions = append(endPositions, pos)
	}

	if startPositions[0].Timestamp < ibNode.lastTimestamp {
		// message stream should guarantee that this should not happen
		err := fmt.Errorf("insert buffer node consumed old messages, channel = %s, timestamp = %d, lastTimestamp = %d",
			ibNode.channelName, startPositions[0].Timestamp, ibNode.lastTimestamp)
		log.Error(err.Error())
		panic(err)
	}

	ibNode.lastTimestamp = endPositions[0].Timestamp

	// Updating segment statistics in channel
	seg2Upload, err := ibNode.updateSegmentStates(fgMsg.insertMessages, startPositions[0], endPositions[0])
	if err != nil {
		// Occurs only if the collectionID is mismatch, should not happen
		err = fmt.Errorf("update segment states in channel meta wrong, err = %s", err)
		log.Error(err.Error())
		panic(err)
	}

	// insert messages -> buffer
	for _, msg := range fgMsg.insertMessages {
		err := ibNode.bufferInsertMsg(msg, endPositions[0])
		if err != nil {
			// error occurs when missing schema info or data is misaligned, should not happen
			err = fmt.Errorf("insertBufferNode msg to buffer failed, err = %s", err)
			log.Error(err.Error())
			panic(err)
		}
	}

	ibNode.DisplayStatistics(seg2Upload)

	segmentsToFlush := ibNode.Flush(fgMsg, seg2Upload, endPositions[0])

	ibNode.WriteTimeTick(fgMsg.timeRange.timestampMax, seg2Upload)

	res := flowGraphMsg{
		deleteMessages:  fgMsg.deleteMessages,
		timeRange:       fgMsg.timeRange,
		startPositions:  fgMsg.startPositions,
		endPositions:    fgMsg.endPositions,
		segmentsToFlush: segmentsToFlush,
		dropCollection:  fgMsg.dropCollection,
	}

	for _, sp := range spans {
		sp.Finish()
	}

	// send delete msg to DeleteNode
	return []Msg{&res}
}

func (ibNode *insertBufferNode) verifyInMsg(in []Msg) (*flowGraphMsg, bool) {
	// while closing
	if in == nil {
		log.Debug("type assertion failed for flowGraphMsg because it's nil")
		return nil, false
	}

	if len(in) != 1 {
		log.Warn("Invalid operate message input in insertBufferNode", zap.Int("input length", len(in)))
		return nil, false
	}

	fgMsg, ok := in[0].(*flowGraphMsg)
	if !ok {
		log.Warn("type assertion failed for flowGraphMsg", zap.String("name", reflect.TypeOf(in[0]).Name()))
	}
	return fgMsg, ok
}

// DisplayStatistics logs the statistic changes of segment in mem
func (ibNode *insertBufferNode) DisplayStatistics(seg2Upload []UniqueID) {
	// Find and return the smaller input
	min := func(former, latter int) (smaller int) {
		if former <= latter {
			return former
		}
		return latter
	}

	displaySize := min(10, len(seg2Upload))

	// Log the segment statistics in mem
	for k, segID := range seg2Upload[:displaySize] {
		bd, ok := ibNode.insertBuffer.Load(segID)
		if !ok {
			continue
		}

		log.Info("insert seg buffer status", zap.Int("No.", k),
			zap.Int64("segmentID", segID),
			zap.String("vchannel name", ibNode.channelName),
			zap.Int64("buffer size", bd.(*BufferData).size),
			zap.Int64("buffer limit", bd.(*BufferData).limit))
	}
}

func (ibNode *insertBufferNode) Flush(fgMsg *flowGraphMsg, seg2Upload []UniqueID, endPosition *internalpb.MsgPosition) []UniqueID {
	type flushTask struct {
		buffer    *BufferData
		segmentID UniqueID
		flushed   bool
		dropped   bool
		auto      bool
	}

	var (
		flushTaskList   []flushTask
		segmentsToFlush []UniqueID
	)

	if fgMsg.dropCollection {
		segmentsToFlush := ibNode.channel.listAllSegmentIDs()
		log.Info("Receive drop collection request and flushing all segments",
			zap.Any("segments", segmentsToFlush),
			zap.String("channel", ibNode.channelName),
		)
		flushTaskList = make([]flushTask, 0, len(segmentsToFlush))

		for _, seg2Flush := range segmentsToFlush {
			var buf *BufferData
			bd, ok := ibNode.insertBuffer.Load(seg2Flush)
			if !ok {
				buf = nil
			} else {
				buf = bd.(*BufferData)
			}
			flushTaskList = append(flushTaskList, flushTask{
				buffer:    buf,
				segmentID: seg2Flush,
				flushed:   false,
				dropped:   true,
			})
		}
	} else {
		segmentsToFlush = make([]UniqueID, 0, len(seg2Upload)+1) //auto flush number + possible manual flush
		flushTaskList = make([]flushTask, 0, len(seg2Upload)+1)

		// Auto Syncing
		for _, segToFlush := range seg2Upload {
			if bd, ok := ibNode.insertBuffer.Load(segToFlush); ok && bd.(*BufferData).effectiveCap() <= 0 {
				log.Info("(Auto Flush)",
					zap.Int64("segment id", segToFlush),
					zap.String("vchannel name", ibNode.channelName),
				)
				ibuffer := bd.(*BufferData)

				flushTaskList = append(flushTaskList, flushTask{
					buffer:    ibuffer,
					segmentID: segToFlush,
					flushed:   false,
					dropped:   false,
					auto:      true,
				})
			}
		}

		mergeFlushTask := func(segmentID UniqueID, setupTask func(task *flushTask)) {
			// Merge auto & manual flush tasks with the same segment ID.
			dup := false
			for i, task := range flushTaskList {
				if task.segmentID == segmentID {
					log.Info("merging flush task, updating flushed flag",
						zap.Int64("segment ID", segmentID))
					setupTask(&flushTaskList[i])
					dup = true
					break
				}
			}
			// Load buffer and create new flush task if there's no existing flush task for this segment.
			if !dup {
				bd, ok := ibNode.insertBuffer.Load(segmentID)
				var buf *BufferData
				if ok {
					buf = bd.(*BufferData)
				}
				task := flushTask{
					buffer:    buf,
					segmentID: segmentID,
					dropped:   false,
				}
				setupTask(&task)
				flushTaskList = append(flushTaskList, task)
			}
		}

		// Manual Syncing
		select {
		case fmsg := <-ibNode.flushChan:
			log.Info("(Manual Flush) receiving flush message",
				zap.Int64("segmentID", fmsg.segmentID),
				zap.Int64("collectionID", fmsg.collectionID),
				zap.Bool("flushed", fmsg.flushed),
				zap.String("v-channel name", ibNode.channelName),
			)
			mergeFlushTask(fmsg.segmentID, func(task *flushTask) {
				task.flushed = fmsg.flushed
			})
		default:
		}

		// process drop partition
		for _, partitionDrop := range fgMsg.dropPartitions {
			segmentIDs := ibNode.channel.listPartitionSegments(partitionDrop)
			log.Info("(Drop Partition) process drop partition",
				zap.Int64("collectionID", ibNode.channel.getCollectionID()),
				zap.Int64("partitionID", partitionDrop),
				zap.Int64s("segmentIDs", segmentIDs),
				zap.String("v-channel name", ibNode.channelName),
			)
			for _, segID := range segmentIDs {
				mergeFlushTask(segID, func(task *flushTask) {
					task.flushed = true
					task.dropped = true
				})
			}
		}
	}

	for _, task := range flushTaskList {
		log.Debug("insertBufferNode flushing BufferData",
			zap.Int64("segmentID", task.segmentID),
			zap.Bool("flushed", task.flushed),
			zap.Bool("dropped", task.dropped),
			zap.Any("position", endPosition),
		)

		segStats, err := ibNode.channel.getSegmentStatslog(task.segmentID)
		if err != nil && !errors.Is(err, errSegmentStatsNotChanged) {
			log.Error("failed to get segment stats log", zap.Int64("segmentID", task.segmentID), zap.Error(err))
			panic(err)
		}

		err = retry.Do(ibNode.ctx, func() error {
			return ibNode.flushManager.flushBufferData(task.buffer,
				segStats,
				task.segmentID,
				task.flushed,
				task.dropped,
				endPosition)
		}, getFlowGraphRetryOpt())
		if err != nil {
			metrics.DataNodeFlushBufferCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.FailLabel).Inc()
			metrics.DataNodeFlushBufferCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.TotalLabel).Inc()
			if task.auto {
				metrics.DataNodeAutoFlushBufferCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.FailLabel).Inc()
				metrics.DataNodeAutoFlushBufferCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.TotalLabel).Inc()
			}
			err = fmt.Errorf("insertBufferNode flushBufferData failed, err = %s", err)
			log.Error(err.Error())
			panic(err)
		}
		segmentsToFlush = append(segmentsToFlush, task.segmentID)
		ibNode.insertBuffer.Delete(task.segmentID)
		metrics.DataNodeFlushBufferCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.SuccessLabel).Inc()
		metrics.DataNodeFlushBufferCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.TotalLabel).Inc()
		if task.auto {
			metrics.DataNodeAutoFlushBufferCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.TotalLabel).Inc()
			metrics.DataNodeAutoFlushBufferCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.FailLabel).Inc()
		}
	}
	return segmentsToFlush
}

// updateSegmentStates updates statistics in channel meta for the segments in insertMsgs.
//  If the segment doesn't exist, a new segment will be created.
//  The segment number of rows will be updated in mem, waiting to be uploaded to DataCoord.
func (ibNode *insertBufferNode) updateSegmentStates(insertMsgs []*msgstream.InsertMsg, startPos, endPos *internalpb.MsgPosition) (seg2Upload []UniqueID, err error) {
	uniqueSeg := make(map[UniqueID]int64)
	for _, msg := range insertMsgs {

		currentSegID := msg.GetSegmentID()
		collID := msg.GetCollectionID()
		partitionID := msg.GetPartitionID()

		if !ibNode.channel.hasSegment(currentSegID, true) {
			err = ibNode.channel.addSegment(
				addSegmentReq{
					segType:     datapb.SegmentType_New,
					segID:       currentSegID,
					collID:      collID,
					partitionID: partitionID,
					startPos:    startPos,
					endPos:      endPos,
				})
			if err != nil {
				log.Error("add segment wrong",
					zap.Int64("segID", currentSegID),
					zap.Int64("collID", collID),
					zap.Int64("partID", partitionID),
					zap.String("chanName", msg.GetShardName()),
					zap.Error(err))
				return
			}
		}

		segNum := uniqueSeg[currentSegID]
		uniqueSeg[currentSegID] = segNum + int64(len(msg.RowIDs))
	}

	seg2Upload = make([]UniqueID, 0, len(uniqueSeg))
	for id, num := range uniqueSeg {
		seg2Upload = append(seg2Upload, id)
		ibNode.channel.updateStatistics(id, num)
	}

	return
}

/* #nosec G103 */
// bufferInsertMsg put InsertMsg into buffer
// 	1.1 fetch related schema from channel meta
// 	1.2 Get buffer data and put data into each field buffer
// 	1.3 Put back into buffer
// 	1.4 Update related statistics
func (ibNode *insertBufferNode) bufferInsertMsg(msg *msgstream.InsertMsg, endPos *internalpb.MsgPosition) error {
	if err := msg.CheckAligned(); err != nil {
		return err
	}
	currentSegID := msg.GetSegmentID()
	collectionID := msg.GetCollectionID()

	collSchema, err := ibNode.channel.getCollectionSchema(collectionID, msg.EndTs())
	if err != nil {
		log.Error("Get schema wrong:", zap.Error(err))
		return err
	}

	// Get Dimension
	// TODO GOOSE: under assumption that there's only 1 Vector field in one collection schema
	var dimension int
	for _, field := range collSchema.Fields {
		if field.DataType == schemapb.DataType_FloatVector ||
			field.DataType == schemapb.DataType_BinaryVector {

			dimension, err = storage.GetDimFromParams(field.TypeParams)
			if err != nil {
				log.Error("failed to get dim from field", zap.Error(err))
				return err
			}
			break
		}
	}

	newbd, err := newBufferData(int64(dimension))
	if err != nil {
		return err
	}
	bd, _ := ibNode.insertBuffer.LoadOrStore(currentSegID, newbd)

	buffer := bd.(*BufferData)
	// idata := buffer.buffer

	addedBuffer, err := storage.InsertMsgToInsertData(msg, collSchema)
	if err != nil {
		log.Error("failed to transfer insert msg to insert data", zap.Error(err))
		return err
	}

	addedPfData, err := storage.GetPkFromInsertData(collSchema, addedBuffer)
	if err != nil {
		log.Warn("no primary field found in insert msg", zap.Error(err))
	} else {
		ibNode.channel.updateSegmentPKRange(currentSegID, addedPfData)
	}

	// Maybe there are large write zoom if frequent insert requests are met.
	buffer.buffer = storage.MergeInsertData(buffer.buffer, addedBuffer)

	tsData, err := storage.GetTimestampFromInsertData(addedBuffer)
	if err != nil {
		log.Warn("no timestamp field found in insert msg", zap.Error(err))
		return err
	}

	// update buffer size
	buffer.updateSize(int64(msg.NRows()))
	// update timestamp range
	buffer.updateTimeRange(ibNode.getTimestampRange(tsData))

	metrics.DataNodeConsumeMsgRowsCount.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), metrics.InsertLabel).Add(float64(len(msg.RowData)))

	// store in buffer
	ibNode.insertBuffer.Store(currentSegID, buffer)

	// store current endPositions as Segment->EndPostion
	ibNode.channel.updateSegmentEndPosition(currentSegID, endPos)

	return nil
}

func (ibNode *insertBufferNode) getTimestampRange(tsData *storage.Int64FieldData) TimeRange {
	tr := TimeRange{
		timestampMin: math.MaxUint64,
		timestampMax: 0,
	}

	for _, data := range tsData.Data {
		if uint64(data) < tr.timestampMin {
			tr.timestampMin = Timestamp(data)
		}
		if uint64(data) > tr.timestampMax {
			tr.timestampMax = Timestamp(data)
		}
	}
	return tr
}

// WriteTimeTick writes timetick once insertBufferNode operates.
func (ibNode *insertBufferNode) WriteTimeTick(ts Timestamp, segmentIDs []int64) {

	select {
	case resendTTMsg := <-ibNode.resendTTChan:
		log.Info("resend TT msg received in insertBufferNode",
			zap.Int64s("segmentIDs", resendTTMsg.segmentIDs))
		segmentIDs = append(segmentIDs, resendTTMsg.segmentIDs...)
	default:
	}

	ibNode.ttLogger.LogTs(ts)
	ibNode.ttMerger.bufferTs(ts, segmentIDs)
	rateCol.updateFlowGraphTt(ibNode.channelName, ts)
}

func (ibNode *insertBufferNode) getCollectionandPartitionIDbySegID(segmentID UniqueID) (collID, partitionID UniqueID, err error) {
	return ibNode.channel.getCollectionAndPartitionID(segmentID)
}

func newInsertBufferNode(ctx context.Context, collID UniqueID, flushCh <-chan flushMsg, resendTTCh <-chan resendTTMsg,
	fm flushManager, flushingSegCache *Cache, config *nodeConfig) (*insertBufferNode, error) {

	baseNode := BaseNode{}
	baseNode.SetMaxQueueLength(config.maxQueueLength)
	baseNode.SetMaxParallelism(config.maxParallelism)

	//input stream, data node time tick
	wTt, err := config.msFactory.NewMsgStream(ctx)
	if err != nil {
		return nil, err
	}
	wTt.AsProducer([]string{Params.CommonCfg.DataCoordTimeTick})
	metrics.DataNodeNumProducers.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID())).Inc()
	log.Debug("datanode AsProducer", zap.String("TimeTickChannelName", Params.CommonCfg.DataCoordTimeTick))
	var wTtMsgStream msgstream.MsgStream = wTt
	wTtMsgStream.Start()

	mt := newMergedTimeTickerSender(func(ts Timestamp, segmentIDs []int64) error {
		stats := make([]*datapb.SegmentStats, 0, len(segmentIDs))
		for _, sid := range segmentIDs {
			stat, err := config.channel.getSegmentStatisticsUpdates(sid)
			if err != nil {
				log.Warn("failed to get segment statistics info", zap.Int64("segmentID", sid), zap.Error(err))
				continue
			}
			stats = append(stats, stat)
		}
		msgPack := msgstream.MsgPack{}
		timeTickMsg := msgstream.DataNodeTtMsg{
			BaseMsg: msgstream.BaseMsg{
				BeginTimestamp: ts,
				EndTimestamp:   ts,
				HashValues:     []uint32{0},
			},
			DataNodeTtMsg: datapb.DataNodeTtMsg{
				Base: &commonpb.MsgBase{
					MsgType:   commonpb.MsgType_DataNodeTt,
					MsgID:     0,
					Timestamp: ts,
				},
				ChannelName:   config.vChannelName,
				Timestamp:     ts,
				SegmentsStats: stats,
			},
		}
		msgPack.Msgs = append(msgPack.Msgs, &timeTickMsg)
		pt, _ := tsoutil.ParseHybridTs(ts)
		pChan := funcutil.ToPhysicalChannel(config.vChannelName)
		metrics.DataNodeTimeSync.WithLabelValues(fmt.Sprint(Params.DataNodeCfg.GetNodeID()), pChan).Set(float64(pt))
		return wTtMsgStream.Produce(&msgPack)
	})

	return &insertBufferNode{
		ctx:          ctx,
		BaseNode:     baseNode,
		insertBuffer: sync.Map{},

		timeTickStream:   wTtMsgStream,
		flushMap:         sync.Map{},
		flushChan:        flushCh,
		resendTTChan:     resendTTCh,
		flushingSegCache: flushingSegCache,
		flushManager:     fm,

		channel:     config.channel,
		idAllocator: config.allocator,
		channelName: config.vChannelName,
		ttMerger:    mt,
		ttLogger:    &timeTickLogger{vChannelName: config.vChannelName},
	}, nil
}
