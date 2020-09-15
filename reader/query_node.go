package reader

/*

#cgo CFLAGS: -I../core/include

#cgo LDFLAGS: -L../core/lib -lmilvus_dog_segment -Wl,-rpath=../core/lib

#include "collection_c.h"
#include "partition_c.h"
#include "segment_c.h"

*/
import "C"

import (
	"fmt"
	msgPb "github.com/czs007/suvlim/pkg/master/grpc/message"
	"github.com/czs007/suvlim/reader/message_client"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type InsertData struct {
	insertIDs        map[int64][]int64
	insertTimestamps map[int64][]uint64
	insertRecords    map[int64][][]byte
	insertOffset     map[int64]int64
}

type DeleteData struct {
	deleteIDs        map[int64][]int64
	deleteTimestamps map[int64][]uint64
	deleteOffset     map[int64]int64
}

type DeleteRecord struct {
	entityID  int64
	timestamp uint64
	segmentID int64
}

type DeletePreprocessData struct {
	deleteRecords []*DeleteRecord
	count         int32
}

type QueryNodeDataBuffer struct {
	InsertDeleteBuffer      []*msgPb.InsertOrDeleteMsg
	SearchBuffer            []*msgPb.SearchMsg
	validInsertDeleteBuffer []bool
	validSearchBuffer       []bool
}

type QueryNode struct {
	QueryNodeId          uint64
	Collections          []*Collection
	SegmentsMap          map[int64]*Segment
	messageClient        message_client.MessageClient
	queryNodeTimeSync    *QueryNodeTime
	buffer               QueryNodeDataBuffer
	deletePreprocessData DeletePreprocessData
	deleteData           DeleteData
	insertData           InsertData
}

func NewQueryNode(queryNodeId uint64, timeSync uint64) *QueryNode {
	mc := message_client.MessageClient{}

	queryNodeTimeSync := &QueryNodeTime{
		ReadTimeSyncMin: timeSync,
		ReadTimeSyncMax: timeSync,
		WriteTimeSync:   timeSync,
		SearchTimeSync:  timeSync,
		TSOTimeSync:     timeSync,
	}

	segmentsMap := make(map[int64]*Segment)

	buffer := QueryNodeDataBuffer{
		InsertDeleteBuffer:      make([]*msgPb.InsertOrDeleteMsg, 0),
		SearchBuffer:            make([]*msgPb.SearchMsg, 0),
		validInsertDeleteBuffer: make([]bool, 0),
		validSearchBuffer:       make([]bool, 0),
	}

	return &QueryNode{
		QueryNodeId:          queryNodeId,
		Collections:          nil,
		SegmentsMap:          segmentsMap,
		messageClient:        mc,
		queryNodeTimeSync:    queryNodeTimeSync,
		buffer:               buffer,
	}
}

func (node *QueryNode) QueryNodeDataInit() {
	deletePreprocessData := DeletePreprocessData{
		deleteRecords: make([]*DeleteRecord, 0),
		count:         0,
	}

	deleteData := DeleteData{
		deleteIDs:        make(map[int64][]int64),
		deleteTimestamps: make(map[int64][]uint64),
		deleteOffset:     make(map[int64]int64),
	}

	insertData := InsertData{
		insertIDs:        make(map[int64][]int64),
		insertTimestamps: make(map[int64][]uint64),
		insertRecords:    make(map[int64][][]byte),
		insertOffset:     make(map[int64]int64),
	}

	node.deletePreprocessData = deletePreprocessData
	node.deleteData = deleteData
	node.insertData = insertData
}

func (node *QueryNode) NewCollection(collectionName string, schemaConfig string) *Collection {
	cName := C.CString(collectionName)
	cSchema := C.CString(schemaConfig)
	collection := C.NewCollection(cName, cSchema)

	var newCollection = &Collection{CollectionPtr: collection, CollectionName: collectionName}
	node.Collections = append(node.Collections, newCollection)

	return newCollection
}

func (node *QueryNode) DeleteCollection(collection *Collection) {
	cPtr := collection.CollectionPtr
	C.DeleteCollection(cPtr)

	// TODO: remove from node.Collections
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (node *QueryNode) PrepareBatchMsg() []int {
	var msgLen = node.messageClient.PrepareBatchMsg()
	return msgLen
}

func (node *QueryNode) StartMessageClient() {
	// TODO: add consumerMsgSchema
	node.messageClient.InitClient("pulsar://192.168.2.28:6650")

	go node.messageClient.ReceiveMessage()
}

func (node *QueryNode) InitQueryNodeCollection() {
	// TODO: remove hard code, add collection creation request
	// TODO: error handle
	var newCollection = node.NewCollection("collection1", "fakeSchema")
	var newPartition = newCollection.NewPartition("partition1")
	// TODO: add segment id
	var segment = newPartition.NewSegment(0)
	node.SegmentsMap[0] = segment
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (node *QueryNode) RunInsertDelete() {
	var count = 0
	var start time.Time
	for {
		//time.Sleep(2 * 1000 * time.Millisecond)
		node.QueryNodeDataInit()
		// TODO: get timeRange from message client
		var timeRange = TimeRange{0, 0}
		var msgLen = node.PrepareBatchMsg()
		//fmt.Println("PrepareBatchMsg Done, Insert len = ", msgLen[0])
		if msgLen[0] == 0 {
			//fmt.Println("0 msg found")
			continue
		}
		if count == 0 {
			start = time.Now()
		}
		count+=msgLen[0]
		node.MessagesPreprocess(node.messageClient.InsertOrDeleteMsg, timeRange)
		//fmt.Println("MessagesPreprocess Done")
		node.WriterDelete()
		node.PreInsertAndDelete()
		//fmt.Println("PreInsertAndDelete Done")
		node.DoInsertAndDelete()
		//fmt.Println("DoInsertAndDelete Done")
		node.queryNodeTimeSync.UpdateSearchTimeSync(timeRange)
		//fmt.Print("UpdateSearchTimeSync Done\n\n\n")
		if count == 100000 - 1 {
			elapsed := time.Since(start)
			fmt.Println("Query node insert 10 × 10000 time:", elapsed)
		}
	}
}

func (node *QueryNode) RunSearch() {
	for {
		//time.Sleep(2 * 1000 * time.Millisecond)

		start := time.Now()

		if len(node.messageClient.GetSearchChan()) <= 0 {
			//fmt.Println("null Search")
			continue
		}
		node.messageClient.SearchMsg = node.messageClient.SearchMsg[:0]
		msg := <-node.messageClient.GetSearchChan()
		node.messageClient.SearchMsg = append(node.messageClient.SearchMsg, msg)
		fmt.Println("Do Search...")
		node.Search(node.messageClient.SearchMsg)

		elapsed := time.Since(start)
		fmt.Println("Query node search time:", elapsed)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (node *QueryNode) MessagesPreprocess(insertDeleteMessages []*msgPb.InsertOrDeleteMsg, timeRange TimeRange) msgPb.Status {
	//var tMax = timeRange.timestampMax

	// 1. Extract messages before readTimeSync from QueryNodeDataBuffer.
	//    Set valid bitmap to false.
	for i, msg := range node.buffer.InsertDeleteBuffer {
		//if msg.Timestamp < tMax {
		if msg.Op == msgPb.OpType_INSERT {
			if msg.RowsData == nil {
				continue
			}
			node.insertData.insertIDs[msg.SegmentId] = append(node.insertData.insertIDs[msg.SegmentId], msg.Uid)
			node.insertData.insertTimestamps[msg.SegmentId] = append(node.insertData.insertTimestamps[msg.SegmentId], msg.Timestamp)
			node.insertData.insertRecords[msg.SegmentId] = append(node.insertData.insertRecords[msg.SegmentId], msg.RowsData.Blob)
		} else if msg.Op == msgPb.OpType_DELETE {
			var r = DeleteRecord{
				entityID:  msg.Uid,
				timestamp: msg.Timestamp,
			}
			node.deletePreprocessData.deleteRecords = append(node.deletePreprocessData.deleteRecords, &r)
			atomic.AddInt32(&node.deletePreprocessData.count, 1)
		}
		node.buffer.validInsertDeleteBuffer[i] = false
		//}
	}

	// 2. Remove invalid messages from buffer.
	for i, isValid := range node.buffer.validInsertDeleteBuffer {
		if !isValid {
			copy(node.buffer.InsertDeleteBuffer[i:], node.buffer.InsertDeleteBuffer[i+1:])                          // Shift a[i+1:] left one index.
			node.buffer.InsertDeleteBuffer[len(node.buffer.InsertDeleteBuffer)-1] = nil                             // Erase last element (write zero value).
			node.buffer.InsertDeleteBuffer = node.buffer.InsertDeleteBuffer[:len(node.buffer.InsertDeleteBuffer)-1] // Truncate slice.
		}
	}

	// 3. Extract messages before readTimeSync from current messageClient.
	//    Move massages after readTimeSync to QueryNodeDataBuffer.
	//    Set valid bitmap to true.
	for _, msg := range insertDeleteMessages {
		//if msg.Timestamp < tMax {
		if msg.Op == msgPb.OpType_INSERT {
			if msg.RowsData == nil {
				continue
			}
			node.insertData.insertIDs[msg.SegmentId] = append(node.insertData.insertIDs[msg.SegmentId], msg.Uid)
			node.insertData.insertTimestamps[msg.SegmentId] = append(node.insertData.insertTimestamps[msg.SegmentId], msg.Timestamp)
			node.insertData.insertRecords[msg.SegmentId] = append(node.insertData.insertRecords[msg.SegmentId], msg.RowsData.Blob)
		} else if msg.Op == msgPb.OpType_DELETE {
			var r = DeleteRecord{
				entityID:  msg.Uid,
				timestamp: msg.Timestamp,
			}
			node.deletePreprocessData.deleteRecords = append(node.deletePreprocessData.deleteRecords, &r)
			atomic.AddInt32(&node.deletePreprocessData.count, 1)
		}
		//} else {
		//	node.buffer.InsertDeleteBuffer = append(node.buffer.InsertDeleteBuffer, msg)
		//	node.buffer.validInsertDeleteBuffer = append(node.buffer.validInsertDeleteBuffer, true)
		//}
	}

	return msgPb.Status{ErrorCode: msgPb.ErrorCode_SUCCESS}
}

func (node *QueryNode) WriterDelete() msgPb.Status {
	// TODO: set timeout
	for {
		if node.deletePreprocessData.count == 0 {
			return msgPb.Status{ErrorCode: msgPb.ErrorCode_SUCCESS}
		}
		node.messageClient.PrepareKey2SegmentMsg()
		var ids, timestamps, segmentIDs = node.GetKey2Segments()
		for i := 0; i < len(*ids); i++ {
			id := (*ids)[i]
			timestamp := (*timestamps)[i]
			segmentID := (*segmentIDs)[i]
			for _, r := range node.deletePreprocessData.deleteRecords {
				if r.timestamp == timestamp && r.entityID == id {
					r.segmentID = segmentID
					atomic.AddInt32(&node.deletePreprocessData.count, -1)
				}
			}
		}
	}
}

func (node *QueryNode) PreInsertAndDelete() msgPb.Status {
	// 1. Do PreInsert
	for segmentID := range node.insertData.insertRecords {
		var targetSegment, err = node.GetSegmentBySegmentID(segmentID)
		if err != nil {
			fmt.Println(err.Error())
			return msgPb.Status{ErrorCode: 1}
		}

		var numOfRecords = len(node.insertData.insertRecords[segmentID])
		var offset = targetSegment.SegmentPreInsert(numOfRecords)
		node.insertData.insertOffset[segmentID] = offset
	}

	// 2. Sort delete preprocess data by segment id
	for _, r := range node.deletePreprocessData.deleteRecords {
		node.deleteData.deleteIDs[r.segmentID] = append(node.deleteData.deleteIDs[r.segmentID], r.entityID)
		node.deleteData.deleteTimestamps[r.segmentID] = append(node.deleteData.deleteTimestamps[r.segmentID], r.timestamp)
	}

	// 3. Do PreDelete
	for segmentID := range node.deleteData.deleteIDs {
		var targetSegment, err = node.GetSegmentBySegmentID(segmentID)
		if err != nil {
			fmt.Println(err.Error())
			return msgPb.Status{ErrorCode: 1}
		}

		var numOfRecords = len(node.deleteData.deleteIDs[segmentID])
		var offset = targetSegment.SegmentPreDelete(numOfRecords)
		node.deleteData.deleteOffset[segmentID] = offset
	}

	return msgPb.Status{ErrorCode: msgPb.ErrorCode_SUCCESS}
}

func (node *QueryNode) DoInsertAndDelete() msgPb.Status {
	var wg sync.WaitGroup
	// Do insert
	for segmentID, records := range node.insertData.insertRecords {
		wg.Add(1)
		go node.DoInsert(segmentID, &records, &wg)
	}

	// Do delete
	for segmentID, deleteIDs := range node.deleteData.deleteIDs {
		wg.Add(1)
		var deleteTimestamps = node.deleteData.deleteTimestamps[segmentID]
		fmt.Println("Doing delete......")
		go node.DoDelete(segmentID, &deleteIDs, &deleteTimestamps, &wg)
	}

	wg.Wait()
	return msgPb.Status{ErrorCode: msgPb.ErrorCode_SUCCESS}
}

func (node *QueryNode) DoInsert(segmentID int64, records *[][]byte, wg *sync.WaitGroup) msgPb.Status {
	var targetSegment, err = node.GetSegmentBySegmentID(segmentID)
	if err != nil {
		fmt.Println(err.Error())
		return msgPb.Status{ErrorCode: 1}
	}

	ids := node.insertData.insertIDs[segmentID]
	timestamps := node.insertData.insertTimestamps[segmentID]
	offsets := node.insertData.insertOffset[segmentID]

	err = targetSegment.SegmentInsert(offsets, &ids, &timestamps, records)
	if err != nil {
		fmt.Println(err.Error())
		return msgPb.Status{ErrorCode: 1}
	}

	wg.Done()
	return msgPb.Status{ErrorCode: msgPb.ErrorCode_SUCCESS}
}

func (node *QueryNode) DoDelete(segmentID int64, deleteIDs *[]int64, deleteTimestamps *[]uint64, wg *sync.WaitGroup) msgPb.Status {
	var segment, err = node.GetSegmentBySegmentID(segmentID)
	if err != nil {
		fmt.Println(err.Error())
		return msgPb.Status{ErrorCode: 1}
	}

	offset := node.deleteData.deleteOffset[segmentID]

	err = segment.SegmentDelete(offset, deleteIDs, deleteTimestamps)
	if err != nil {
		fmt.Println(err.Error())
		return msgPb.Status{ErrorCode: 1}
	}

	wg.Done()
	return msgPb.Status{ErrorCode: msgPb.ErrorCode_SUCCESS}
}

func (node *QueryNode) Search(searchMessages []*msgPb.SearchMsg) msgPb.Status {
	var clientId = (*(searchMessages[0])).ClientId

	type SearchResultTmp struct {
		ResultId       int64
		ResultDistance float32
	}

	// Traverse all messages in the current messageClient.
	// TODO: Do not receive batched search requests
	for _, msg := range searchMessages {
		var collectionName = searchMessages[0].CollectionName
		var targetCollection, err = node.GetCollectionByCollectionName(collectionName)
		if err != nil {
			fmt.Println(err.Error())
			return msgPb.Status{ErrorCode: 1}
		}

		var resultsTmp = make([]SearchResultTmp, 0)
		// TODO: get top-k's k from queryString
		const TopK = 1

		var timestamp = msg.Timestamp
		var vector = msg.Records

		// 1. Timestamp check
		// TODO: return or wait? Or adding graceful time
		//if timestamp > node.queryNodeTimeSync.SearchTimeSync {
		//	return msgPb.Status{ErrorCode: 1}
		//}

		// 2. Do search in all segments
		for _, partition := range targetCollection.Partitions {
			for _, openSegment := range partition.OpenedSegments {
				var res, err = openSegment.SegmentSearch("", timestamp, vector)
				if err != nil {
					fmt.Println(err.Error())
					return msgPb.Status{ErrorCode: 1}
				}
				fmt.Println(res.ResultIds)
				for i := 0; i < len(res.ResultIds); i++ {
					resultsTmp = append(resultsTmp, SearchResultTmp{ResultId: res.ResultIds[i], ResultDistance: res.ResultDistances[i]})
				}
			}
			for _, closedSegment := range partition.ClosedSegments {
				var res, err = closedSegment.SegmentSearch("", timestamp, vector)
				if err != nil {
					fmt.Println(err.Error())
					return msgPb.Status{ErrorCode: 1}
				}
				for i := 0; i <= len(res.ResultIds); i++ {
					resultsTmp = append(resultsTmp, SearchResultTmp{ResultId: res.ResultIds[i], ResultDistance: res.ResultDistances[i]})
				}
			}
		}

		// 2. Reduce results
		sort.Slice(resultsTmp, func(i, j int) bool {
			return resultsTmp[i].ResultDistance < resultsTmp[j].ResultDistance
		})
		resultsTmp = resultsTmp[:TopK]
		var entities = msgPb.Entities{
			Ids: make([]int64, 0),
		}
		var results = msgPb.QueryResult{
			Entities:  &entities,
			Distances: make([]float32, 0),
			QueryId:   msg.Uid,
		}
		for _, res := range resultsTmp {
			results.Entities.Ids = append(results.Entities.Ids, res.ResultId)
			results.Distances = append(results.Distances, res.ResultDistance)
			results.Scores = append(results.Distances, float32(0))
		}

		results.RowNum = int64(len(results.Distances))

		// 3. publish result to pulsar
		node.PublishSearchResult(&results, clientId)
	}

	return msgPb.Status{ErrorCode: msgPb.ErrorCode_SUCCESS}
}
