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
	"bytes"
	"context"
	"encoding/binary"
	"math"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/schemapb"
	"github.com/milvus-io/milvus/internal/common"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/mq/msgstream"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/internal/util/dependency"
	"github.com/milvus-io/milvus/internal/util/paramtable"
	"github.com/milvus-io/milvus/internal/util/tsoutil"
)

var dataSyncServiceTestDir = "/tmp/milvus_test/data_sync_service"

func getVchanInfo(info *testInfo) *datapb.VchannelInfo {
	var ufs []*datapb.SegmentInfo
	var fs []*datapb.SegmentInfo
	if info.isValidCase {
		ufs = []*datapb.SegmentInfo{{
			CollectionID:  info.ufCollID,
			PartitionID:   1,
			InsertChannel: info.ufchanName,
			ID:            info.ufSegID,
			NumOfRows:     info.ufNor,
			DmlPosition:   &internalpb.MsgPosition{},
		}}
		fs = []*datapb.SegmentInfo{{
			CollectionID:  info.fCollID,
			PartitionID:   1,
			InsertChannel: info.fchanName,
			ID:            info.fSegID,
			NumOfRows:     info.fNor,
			DmlPosition:   &internalpb.MsgPosition{},
		}}
	} else {
		ufs = []*datapb.SegmentInfo{}
	}

	var ufsIds []int64
	var fsIds []int64
	for _, segmentInfo := range ufs {
		ufsIds = append(ufsIds, segmentInfo.ID)
	}
	for _, segmentInfo := range fs {
		fsIds = append(fsIds, segmentInfo.ID)
	}
	vi := &datapb.VchannelInfo{
		CollectionID:        info.collID,
		ChannelName:         info.chanName,
		SeekPosition:        &internalpb.MsgPosition{},
		UnflushedSegmentIds: ufsIds,
		FlushedSegmentIds:   fsIds,
	}
	return vi
}

type testInfo struct {
	isValidCase  bool
	channelNil   bool
	inMsgFactory msgstream.Factory

	collID   UniqueID
	chanName string

	ufCollID   UniqueID
	ufSegID    UniqueID
	ufchanName string
	ufNor      int64

	fCollID   UniqueID
	fSegID    UniqueID
	fchanName string
	fNor      int64

	description string
}

func TestDataSyncService_newDataSyncService(te *testing.T) {

	ctx := context.Background()

	tests := []*testInfo{
		{true, false, &mockMsgStreamFactory{false, true},
			0, "by-dev-rootcoord-dml-test_v0",
			0, 0, "", 0,
			0, 0, "", 0,
			"SetParamsReturnError"},
		{true, false, &mockMsgStreamFactory{true, true},
			0, "by-dev-rootcoord-dml-test_v0",
			1, 0, "", 0,
			1, 1, "", 0,
			"CollID 0 mismach with seginfo collID 1"},
		{true, false, &mockMsgStreamFactory{true, true},
			1, "by-dev-rootcoord-dml-test_v1",
			1, 0, "by-dev-rootcoord-dml-test_v2", 0,
			1, 1, "by-dev-rootcoord-dml-test_v3", 0,
			"chanName c1 mismach with seginfo chanName c2"},
		{true, false, &mockMsgStreamFactory{true, true},
			1, "by-dev-rootcoord-dml-test_v1",
			1, 0, "by-dev-rootcoord-dml-test_v1", 0,
			1, 1, "by-dev-rootcoord-dml-test_v2", 0,
			"add normal segments"},
		{false, false, &mockMsgStreamFactory{true, false},
			0, "by-dev-rootcoord-dml-test_v0",
			0, 0, "", 0,
			0, 0, "", 0,
			"error when newinsertbufernode"},
		{false, true, &mockMsgStreamFactory{true, false},
			0, "by-dev-rootcoord-dml-test_v0",
			0, 0, "", 0,
			0, 0, "", 0,
			"channel nil"},
		{true, false, &mockMsgStreamFactory{true, true},
			1, "by-dev-rootcoord-dml-test_v1",
			1, 1, "by-dev-rootcoord-dml-test_v1", 0,
			1, 2, "by-dev-rootcoord-dml-test_v1", 0,
			"add un-flushed and flushed segments"},
	}
	cm := storage.NewLocalChunkManager(storage.RootPath(dataSyncServiceTestDir))
	defer cm.RemoveWithPrefix(ctx, cm.RootPath())

	for _, test := range tests {
		te.Run(test.description, func(t *testing.T) {
			df := &DataCoordFactory{}
			rc := &RootCoordFactory{pkType: schemapb.DataType_Int64}

			channel := newChannel("channel", test.collID, nil, rc, cm)
			if test.channelNil {
				channel = nil
			}

			ds, err := newDataSyncService(ctx,
				make(chan flushMsg),
				make(chan resendTTMsg),
				channel,
				NewAllocatorFactory(),
				test.inMsgFactory,
				getVchanInfo(test),
				make(chan string),
				df,
				newCache(),
				cm,
				newCompactionExecutor(),
			)

			if !test.isValidCase {
				assert.Error(t, err)
				assert.Nil(t, ds)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, ds)

				// start
				ds.fg = nil
				ds.start()
			}
		})
	}

}

// NOTE: start pulsar before test
func TestDataSyncService_Start(t *testing.T) {
	const ctxTimeInMillisecond = 10000

	delay := time.Now().Add(ctxTimeInMillisecond * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), delay)
	defer cancel()

	// init data node
	insertChannelName := "by-dev-rootcoord-dml"
	ddlChannelName := "by-dev-rootcoord-ddl"

	Factory := &MetaFactory{}
	collMeta := Factory.GetCollectionMeta(UniqueID(0), "coll1", schemapb.DataType_Int64)
	mockRootCoord := &RootCoordFactory{
		pkType: schemapb.DataType_Int64,
	}

	flushChan := make(chan flushMsg, 100)
	resendTTChan := make(chan resendTTMsg, 100)
	cm := storage.NewLocalChunkManager(storage.RootPath(dataSyncServiceTestDir))
	defer cm.RemoveWithPrefix(ctx, cm.RootPath())
	channel := newChannel(insertChannelName, collMeta.ID, collMeta.GetSchema(), mockRootCoord, cm)

	allocFactory := NewAllocatorFactory(1)
	factory := dependency.NewDefaultFactory(true)
	defer os.RemoveAll("/tmp/milvus")
	paramtable.Get().Save(Params.DataNodeCfg.FlushInsertBufferSize.Key, "1")

	ufs := []*datapb.SegmentInfo{{
		CollectionID:  collMeta.ID,
		PartitionID:   1,
		InsertChannel: insertChannelName,
		ID:            0,
		NumOfRows:     0,
		DmlPosition:   &internalpb.MsgPosition{},
	}}
	fs := []*datapb.SegmentInfo{{
		CollectionID:  collMeta.ID,
		PartitionID:   1,
		InsertChannel: insertChannelName,
		ID:            1,
		NumOfRows:     0,
		DmlPosition:   &internalpb.MsgPosition{},
	}}
	var ufsIds []int64
	var fsIds []int64
	for _, segmentInfo := range ufs {
		ufsIds = append(ufsIds, segmentInfo.ID)
	}
	for _, segmentInfo := range fs {
		fsIds = append(fsIds, segmentInfo.ID)
	}
	vchan := &datapb.VchannelInfo{
		CollectionID:        collMeta.ID,
		ChannelName:         insertChannelName,
		UnflushedSegmentIds: ufsIds,
		FlushedSegmentIds:   fsIds,
	}

	signalCh := make(chan string, 100)

	dataCoord := &DataCoordFactory{}
	dataCoord.UserSegmentInfo = map[int64]*datapb.SegmentInfo{
		0: {
			ID:            0,
			CollectionID:  collMeta.ID,
			PartitionID:   1,
			InsertChannel: insertChannelName,
		},

		1: {
			ID:            1,
			CollectionID:  collMeta.ID,
			PartitionID:   1,
			InsertChannel: insertChannelName,
		},
	}

	sync, err := newDataSyncService(ctx, flushChan, resendTTChan, channel, allocFactory, factory, vchan, signalCh, dataCoord, newCache(), cm, newCompactionExecutor())
	assert.Nil(t, err)

	sync.flushListener = make(chan *segmentFlushPack)
	defer close(sync.flushListener)
	sync.start()
	defer sync.close()

	timeRange := TimeRange{
		timestampMin: 0,
		timestampMax: math.MaxUint64,
	}
	dataFactory := NewDataFactory()
	insertMessages := dataFactory.GetMsgStreamTsInsertMsgs(2, insertChannelName, tsoutil.GetCurrentTime())

	msgPack := msgstream.MsgPack{
		BeginTs: timeRange.timestampMin,
		EndTs:   timeRange.timestampMax,
		Msgs:    insertMessages,
		StartPositions: []*internalpb.MsgPosition{{
			ChannelName: insertChannelName,
		}},
		EndPositions: []*internalpb.MsgPosition{{
			ChannelName: insertChannelName,
		}},
	}

	// generate timeTick
	timeTickMsgPack := msgstream.MsgPack{}

	timeTickMsg := &msgstream.TimeTickMsg{
		BaseMsg: msgstream.BaseMsg{
			BeginTimestamp: Timestamp(0),
			EndTimestamp:   Timestamp(0),
			HashValues:     []uint32{0},
		},
		TimeTickMsg: internalpb.TimeTickMsg{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_TimeTick,
				MsgID:     UniqueID(0),
				Timestamp: math.MaxUint64,
				SourceID:  0,
			},
		},
	}
	timeTickMsgPack.Msgs = append(timeTickMsgPack.Msgs, timeTickMsg)

	// pulsar produce
	assert.NoError(t, err)
	insertStream, _ := factory.NewMsgStream(ctx)
	insertStream.AsProducer([]string{insertChannelName})

	ddStream, _ := factory.NewMsgStream(ctx)
	ddStream.AsProducer([]string{ddlChannelName})

	var insertMsgStream msgstream.MsgStream = insertStream
	var ddMsgStream msgstream.MsgStream = ddStream

	err = insertMsgStream.Produce(&msgPack)
	assert.NoError(t, err)

	_, err = insertMsgStream.Broadcast(&timeTickMsgPack)
	assert.NoError(t, err)
	_, err = ddMsgStream.Broadcast(&timeTickMsgPack)
	assert.NoError(t, err)

	select {
	case flushPack := <-sync.flushListener:
		assert.True(t, flushPack.segmentID == 1)
		return
	case <-sync.ctx.Done():
		assert.Fail(t, "test timeout")
	}
}

func TestDataSyncService_Close(t *testing.T) {
	const ctxTimeInMillisecond = 1000000

	delay := time.Now().Add(ctxTimeInMillisecond * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), delay)
	defer cancel()

	os.RemoveAll("/tmp/milvus")

	// init data node
	insertChannelName := "by-dev-rootcoord-dml2"
	ddlChannelName := "by-dev-rootcoord-ddl2"

	Factory := &MetaFactory{}
	collMeta := Factory.GetCollectionMeta(UniqueID(0), "coll1", schemapb.DataType_Int64)
	mockRootCoord := &RootCoordFactory{
		pkType: schemapb.DataType_Int64,
	}

	flushChan := make(chan flushMsg, 100)
	resendTTChan := make(chan resendTTMsg, 100)
	cm := storage.NewLocalChunkManager(storage.RootPath(dataSyncServiceTestDir))
	defer cm.RemoveWithPrefix(ctx, cm.RootPath())
	channel := newChannel(insertChannelName, collMeta.ID, collMeta.GetSchema(), mockRootCoord, cm)

	allocFactory := NewAllocatorFactory(1)
	factory := dependency.NewDefaultFactory(true)
	defer os.RemoveAll("/tmp/milvus")

	paramtable.Get().Remove(Params.DataNodeCfg.FlushInsertBufferSize.Key)
	ufs := []*datapb.SegmentInfo{{
		CollectionID:  collMeta.ID,
		PartitionID:   1,
		InsertChannel: insertChannelName,
		ID:            0,
		NumOfRows:     0,
		DmlPosition:   &internalpb.MsgPosition{},
	}}
	fs := []*datapb.SegmentInfo{{
		CollectionID:  collMeta.ID,
		PartitionID:   1,
		InsertChannel: insertChannelName,
		ID:            1,
		NumOfRows:     0,
		DmlPosition:   &internalpb.MsgPosition{},
	}}
	var ufsIds []int64
	var fsIds []int64
	for _, segmentInfo := range ufs {
		ufsIds = append(ufsIds, segmentInfo.ID)
	}
	for _, segmentInfo := range fs {
		fsIds = append(fsIds, segmentInfo.ID)
	}
	vchan := &datapb.VchannelInfo{
		CollectionID:        collMeta.ID,
		ChannelName:         insertChannelName,
		UnflushedSegmentIds: ufsIds,
		FlushedSegmentIds:   fsIds,
	}

	signalCh := make(chan string, 100)

	dataCoord := &DataCoordFactory{}
	dataCoord.UserSegmentInfo = map[int64]*datapb.SegmentInfo{
		0: {
			ID:            0,
			CollectionID:  collMeta.ID,
			PartitionID:   1,
			InsertChannel: insertChannelName,
		},

		1: {
			ID:            1,
			CollectionID:  collMeta.ID,
			PartitionID:   1,
			InsertChannel: insertChannelName,
		},
	}

	sync, err := newDataSyncService(ctx, flushChan, resendTTChan, channel, allocFactory, factory, vchan, signalCh, dataCoord, newCache(), cm, newCompactionExecutor())
	assert.Nil(t, err)

	sync.flushListener = make(chan *segmentFlushPack, 10)
	defer close(sync.flushListener)
	sync.start()

	dataFactory := NewDataFactory()
	ts := tsoutil.GetCurrentTime()
	insertMessages := dataFactory.GetMsgStreamTsInsertMsgs(2, insertChannelName, ts)
	msgPack := msgstream.MsgPack{
		BeginTs: ts,
		EndTs:   ts,
		Msgs:    insertMessages,
		StartPositions: []*internalpb.MsgPosition{{
			ChannelName: insertChannelName,
		}},
		EndPositions: []*internalpb.MsgPosition{{
			ChannelName: insertChannelName,
		}},
	}

	// 400 is the actual data
	int64Pks := []primaryKey{
		newInt64PrimaryKey(400),
	}
	deleteMessages := dataFactory.GenMsgStreamDeleteMsgWithTs(0, int64Pks, insertChannelName, ts+1)
	inMsgs := make([]msgstream.TsMsg, 0)
	inMsgs = append(inMsgs, deleteMessages)

	msgPackDelete := msgstream.MsgPack{
		BeginTs: ts + 1,
		EndTs:   ts + 1,
		Msgs:    inMsgs,
		StartPositions: []*internalpb.MsgPosition{{
			ChannelName: insertChannelName,
		}},
		EndPositions: []*internalpb.MsgPosition{{
			ChannelName: insertChannelName,
		}},
	}

	// generate timeTick
	timeTickMsg := &msgstream.TimeTickMsg{
		BaseMsg: msgstream.BaseMsg{
			BeginTimestamp: ts,
			EndTimestamp:   ts + 2,
			HashValues:     []uint32{0},
		},
		TimeTickMsg: internalpb.TimeTickMsg{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_TimeTick,
				MsgID:     UniqueID(2),
				Timestamp: ts + 2,
				SourceID:  0,
			},
		},
	}

	timeTickMsgPack := msgstream.MsgPack{
		BeginTs: ts + 2,
		EndTs:   ts + 2,
		StartPositions: []*internalpb.MsgPosition{{
			ChannelName: insertChannelName,
		}},
		EndPositions: []*internalpb.MsgPosition{{
			ChannelName: insertChannelName,
		}},
	}
	timeTickMsgPack.Msgs = append(timeTickMsgPack.Msgs, timeTickMsg)

	// pulsar produce
	assert.NoError(t, err)
	insertStream, _ := factory.NewMsgStream(ctx)
	insertStream.AsProducer([]string{insertChannelName})

	ddStream, _ := factory.NewMsgStream(ctx)
	ddStream.AsProducer([]string{ddlChannelName})

	var insertMsgStream msgstream.MsgStream = insertStream
	var ddMsgStream msgstream.MsgStream = ddStream

	err = insertMsgStream.Produce(&msgPack)
	assert.NoError(t, err)

	err = insertMsgStream.Produce(&msgPackDelete)
	assert.NoError(t, err)

	_, err = insertMsgStream.Broadcast(&timeTickMsgPack)
	assert.NoError(t, err)
	_, err = ddMsgStream.Broadcast(&timeTickMsgPack)
	assert.NoError(t, err)

	// wait for delete
	for sync.delBufferManager.GetEntriesNum(1) == 0 {
		time.Sleep(100)
	}

	// close and wait for flush
	sync.close()
	for {
		select {
		case flushPack, ok := <-sync.flushListener:
			assert.True(t, ok)
			if flushPack.segmentID == 1 {
				assert.True(t, len(flushPack.insertLogs) == 12)
				assert.True(t, len(flushPack.statsLogs) == 1)
				assert.True(t, len(flushPack.deltaLogs) == 1)
				return
			}
			if flushPack.segmentID == 0 {
				assert.True(t, len(flushPack.insertLogs) == 0)
				assert.True(t, len(flushPack.statsLogs) == 0)
				assert.True(t, len(flushPack.deltaLogs) == 0)
			}
		case <-sync.ctx.Done():
		}
	}
}

func genBytes() (rawData []byte) {
	const DIM = 2
	const N = 1

	// Float vector
	var fvector = [DIM]float32{1, 2}
	for _, ele := range fvector {
		buf := make([]byte, 4)
		common.Endian.PutUint32(buf, math.Float32bits(ele))
		rawData = append(rawData, buf...)
	}

	// Binary vector
	// Dimension of binary vector is 32
	// size := 4,  = 32 / 8
	var bvector = []byte{255, 255, 255, 0}
	rawData = append(rawData, bvector...)

	// Bool
	var fieldBool = true
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, common.Endian, fieldBool); err != nil {
		panic(err)
	}

	rawData = append(rawData, buf.Bytes()...)

	// int8
	var dataInt8 int8 = 100
	bint8 := new(bytes.Buffer)
	if err := binary.Write(bint8, common.Endian, dataInt8); err != nil {
		panic(err)
	}
	rawData = append(rawData, bint8.Bytes()...)
	log.Debug("Rawdata length:", zap.Int("Length of rawData", len(rawData)))
	return
}

func TestBytesReader(t *testing.T) {
	rawData := genBytes()

	// Bytes Reader is able to recording the position
	rawDataReader := bytes.NewReader(rawData)

	var fvector = make([]float32, 2)
	err := binary.Read(rawDataReader, common.Endian, &fvector)
	assert.Nil(t, err)
	assert.ElementsMatch(t, fvector, []float32{1, 2})

	var bvector = make([]byte, 4)
	err = binary.Read(rawDataReader, common.Endian, &bvector)
	assert.Nil(t, err)
	assert.ElementsMatch(t, bvector, []byte{255, 255, 255, 0})

	var fieldBool bool
	err = binary.Read(rawDataReader, common.Endian, &fieldBool)
	assert.Nil(t, err)
	assert.Equal(t, true, fieldBool)

	var dataInt8 int8
	err = binary.Read(rawDataReader, common.Endian, &dataInt8)
	assert.Nil(t, err)
	assert.Equal(t, int8(100), dataInt8)
}

func TestGetSegmentInfos(t *testing.T) {
	dataCoord := &DataCoordFactory{}
	dsService := &dataSyncService{
		dataCoord: dataCoord,
	}
	segmentInfos, err := dsService.getSegmentInfos([]int64{1})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(segmentInfos))

	dataCoord.GetSegmentInfosError = true
	segmentInfos2, err := dsService.getSegmentInfos([]int64{1})
	assert.Error(t, err)
	assert.Empty(t, segmentInfos2)

	dataCoord.GetSegmentInfosError = false
	dataCoord.GetSegmentInfosNotSuccess = true
	segmentInfos3, err := dsService.getSegmentInfos([]int64{1})
	assert.Error(t, err)
	assert.Empty(t, segmentInfos3)

	dataCoord.GetSegmentInfosError = false
	dataCoord.GetSegmentInfosNotSuccess = false
	dataCoord.UserSegmentInfo = map[int64]*datapb.SegmentInfo{
		5: {
			ID:            100,
			CollectionID:  101,
			PartitionID:   102,
			InsertChannel: "by-dev-rootcoord-dml-test_v1",
		},
	}

	segmentInfos, err = dsService.getSegmentInfos([]int64{5})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(segmentInfos))
	assert.Equal(t, int64(100), segmentInfos[0].ID)
}

func TestClearGlobalFlushingCache(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dataCoord := &DataCoordFactory{}
	cm := storage.NewLocalChunkManager(storage.RootPath(dataSyncServiceTestDir))
	defer cm.RemoveWithPrefix(ctx, cm.RootPath())
	channel := newChannel("channel", 1, nil, &RootCoordFactory{pkType: schemapb.DataType_Int64}, cm)
	var err error

	cache := newCache()
	dsService := &dataSyncService{
		dataCoord:        dataCoord,
		channel:          channel,
		flushingSegCache: cache,
	}

	err = channel.addSegment(
		addSegmentReq{
			segType:     datapb.SegmentType_New,
			segID:       1,
			collID:      1,
			partitionID: 1,
			startPos:    &internalpb.MsgPosition{},
			endPos:      &internalpb.MsgPosition{}})
	assert.NoError(t, err)

	err = channel.addSegment(
		addSegmentReq{
			segType:      datapb.SegmentType_Flushed,
			segID:        2,
			collID:       1,
			partitionID:  1,
			numOfRows:    0,
			statsBinLogs: nil,
			recoverTs:    0,
		})
	assert.NoError(t, err)

	err = channel.addSegment(
		addSegmentReq{
			segType:      datapb.SegmentType_Normal,
			segID:        3,
			collID:       1,
			partitionID:  1,
			numOfRows:    0,
			statsBinLogs: nil,
			recoverTs:    0,
		})
	assert.NoError(t, err)

	cache.checkOrCache(1)
	cache.checkOrCache(2)
	cache.checkOrCache(4)

	dsService.clearGlobalFlushingCache()

	assert.False(t, cache.checkIfCached(1))
	assert.False(t, cache.checkIfCached(2))
	assert.False(t, cache.checkIfCached(3))
	assert.True(t, cache.checkIfCached(4))
}

func TestGetChannelLatestMsgID(t *testing.T) {
	delay := time.Now().Add(ctxTimeInMillisecond * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), delay)
	defer cancel()
	factory := dependency.NewDefaultFactory(true)

	dataCoord := &DataCoordFactory{}
	dsService := &dataSyncService{
		dataCoord: dataCoord,
		msFactory: factory,
	}

	dmlChannelName := "fake-by-dev-rootcoord-dml-channel_12345v0"

	insertStream, _ := factory.NewMsgStream(ctx)
	insertStream.AsProducer([]string{dmlChannelName})
	id, err := dsService.getChannelLatestMsgID(ctx, dmlChannelName, 0)
	assert.NoError(t, err)
	assert.NotNil(t, id)
}
