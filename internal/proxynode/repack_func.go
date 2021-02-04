package proxynode

import (
	"log"
	"sort"

	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"

	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"

	"github.com/zilliztech/milvus-distributed/internal/errors"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
)

func insertRepackFunc(tsMsgs []msgstream.TsMsg,
	hashKeys [][]int32,
	segIDAssigner *SegIDAssigner,
	together bool) (map[int32]*msgstream.MsgPack, error) {

	result := make(map[int32]*msgstream.MsgPack)

	channelCountMap := make(map[UniqueID]map[int32]uint32)    //  reqID --> channelID to count
	channelMaxTSMap := make(map[UniqueID]map[int32]Timestamp) //  reqID --> channelID to max Timestamp
	reqSchemaMap := make(map[UniqueID][]UniqueID)             //  reqID --> channelID [2]UniqueID {CollectionID, PartitionID}
	channelNamesMap := make(map[UniqueID][]string)            // collectionID --> channelNames

	for i, request := range tsMsgs {
		if request.Type() != commonpb.MsgType_kInsert {
			return nil, errors.New(string("msg's must be Insert"))
		}
		insertRequest, ok := request.(*msgstream.InsertMsg)
		if !ok {
			return nil, errors.New(string("msg's must be Insert"))
		}

		keys := hashKeys[i]
		timestampLen := len(insertRequest.Timestamps)
		rowIDLen := len(insertRequest.RowIDs)
		rowDataLen := len(insertRequest.RowData)
		keysLen := len(keys)

		if keysLen != timestampLen || keysLen != rowIDLen || keysLen != rowDataLen {
			return nil, errors.New(string("the length of hashValue, timestamps, rowIDs, RowData are not equal"))
		}

		reqID := insertRequest.Base.MsgID
		if _, ok := channelCountMap[reqID]; !ok {
			channelCountMap[reqID] = make(map[int32]uint32)
		}

		if _, ok := channelMaxTSMap[reqID]; !ok {
			channelMaxTSMap[reqID] = make(map[int32]Timestamp)
		}

		if _, ok := reqSchemaMap[reqID]; !ok {
			reqSchemaMap[reqID] = []UniqueID{insertRequest.CollectionID, insertRequest.PartitionID}
		}

		for idx, channelID := range keys {
			channelCountMap[reqID][channelID]++
			if _, ok := channelMaxTSMap[reqID][channelID]; !ok {
				channelMaxTSMap[reqID][channelID] = typeutil.ZeroTimestamp
			}
			ts := insertRequest.Timestamps[idx]
			if channelMaxTSMap[reqID][channelID] < ts {
				channelMaxTSMap[reqID][channelID] = ts
			}
		}

		collID := insertRequest.CollectionID
		if _, ok := channelNamesMap[collID]; !ok {
			channelNames, err := globalInsertChannelsMap.getInsertChannels(collID)
			if err != nil {
				return nil, err
			}
			channelNamesMap[collID] = channelNames
		}
	}

	var getChannelName = func(collID UniqueID, channelID int32) string {
		if _, ok := channelNamesMap[collID]; !ok {
			return ""
		}
		names := channelNamesMap[collID]
		return names[channelID]
	}

	reqSegCountMap := make(map[UniqueID]map[int32]map[UniqueID]uint32)

	for reqID, countInfo := range channelCountMap {
		if _, ok := reqSegCountMap[reqID]; !ok {
			reqSegCountMap[reqID] = make(map[int32]map[UniqueID]uint32)
		}
		schema := reqSchemaMap[reqID]
		collID, partitionID := schema[0], schema[1]
		for channelID, count := range countInfo {
			ts, ok := channelMaxTSMap[reqID][channelID]
			if !ok {
				ts = typeutil.ZeroTimestamp
				log.Println("Warning: did not get max Timstamp!")
			}
			channelName := getChannelName(collID, channelID)
			if channelName == "" {
				return nil, errors.New("ProxyNode, repack_func, can not found channelName")
			}
			mapInfo, err := segIDAssigner.GetSegmentID(collID, partitionID, channelName, count, ts)
			if err != nil {
				return nil, err
			}
			reqSegCountMap[reqID][channelID] = make(map[UniqueID]uint32)
			reqSegCountMap[reqID][channelID] = mapInfo
		}
	}

	reqSegAccumulateCountMap := make(map[UniqueID]map[int32][]uint32)
	reqSegIDMap := make(map[UniqueID]map[int32][]UniqueID)
	reqSegAllocateCounter := make(map[UniqueID]map[int32]uint32)

	for reqID, channelInfo := range reqSegCountMap {
		if _, ok := reqSegAccumulateCountMap[reqID]; !ok {
			reqSegAccumulateCountMap[reqID] = make(map[int32][]uint32)
		}
		if _, ok := reqSegIDMap[reqID]; !ok {
			reqSegIDMap[reqID] = make(map[int32][]UniqueID)
		}
		if _, ok := reqSegAllocateCounter[reqID]; !ok {
			reqSegAllocateCounter[reqID] = make(map[int32]uint32)
		}
		for channelID, segInfo := range channelInfo {
			reqSegAllocateCounter[reqID][channelID] = 0
			keys := make([]UniqueID, len(segInfo))
			i := 0
			for key := range segInfo {
				keys[i] = key
				i++
			}
			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
			accumulate := uint32(0)
			for _, key := range keys {
				accumulate += segInfo[key]
				if _, ok := reqSegAccumulateCountMap[reqID][channelID]; !ok {
					reqSegAccumulateCountMap[reqID][channelID] = make([]uint32, 0)
				}
				reqSegAccumulateCountMap[reqID][channelID] = append(
					reqSegAccumulateCountMap[reqID][channelID],
					accumulate,
				)
				if _, ok := reqSegIDMap[reqID][channelID]; !ok {
					reqSegIDMap[reqID][channelID] = make([]UniqueID, 0)
				}
				reqSegIDMap[reqID][channelID] = append(
					reqSegIDMap[reqID][channelID],
					key,
				)
			}
		}
	}

	var getSegmentID = func(reqID UniqueID, channelID int32) UniqueID {
		reqSegAllocateCounter[reqID][channelID]++
		cur := reqSegAllocateCounter[reqID][channelID]
		accumulateSlice := reqSegAccumulateCountMap[reqID][channelID]
		segIDSlice := reqSegIDMap[reqID][channelID]
		for index, count := range accumulateSlice {
			if cur <= count {
				return segIDSlice[index]
			}
		}
		log.Panic("Can't Found SegmentID")
		return 0
	}

	for i, request := range tsMsgs {
		insertRequest := request.(*msgstream.InsertMsg)
		keys := hashKeys[i]
		reqID := insertRequest.Base.MsgID
		collectionName := insertRequest.CollectionName
		collectionID := insertRequest.CollectionID
		partitionID := insertRequest.PartitionID
		partitionName := insertRequest.PartitionName
		proxyID := insertRequest.Base.SourceID
		channelNames := channelNamesMap[collectionID]
		for index, key := range keys {
			ts := insertRequest.Timestamps[index]
			rowID := insertRequest.RowIDs[index]
			row := insertRequest.RowData[index]
			_, ok := result[key]
			if !ok {
				msgPack := msgstream.MsgPack{}
				result[key] = &msgPack
			}
			segmentID := getSegmentID(reqID, key)
			channelID := channelNames[int(key)%len(channelNames)]
			sliceRequest := internalpb2.InsertRequest{
				Base: &commonpb.MsgBase{
					MsgType:   commonpb.MsgType_kInsert,
					MsgID:     reqID,
					Timestamp: ts,
					SourceID:  proxyID,
				},
				CollectionID:   collectionID,
				PartitionID:    partitionID,
				CollectionName: collectionName,
				PartitionName:  partitionName,
				SegmentID:      segmentID,
				// todo rename to ChannelName
				// ChannelID:  strconv.FormatInt(int64(key), 10),
				ChannelID:  channelID,
				Timestamps: []uint64{ts},
				RowIDs:     []int64{rowID},
				RowData:    []*commonpb.Blob{row},
			}
			insertMsg := &msgstream.InsertMsg{
				InsertRequest: sliceRequest,
			}
			if together { // all rows with same hash value are accumulated to only one message
				if len(result[key].Msgs) <= 0 {
					result[key].Msgs = append(result[key].Msgs, insertMsg)
				} else {
					accMsgs, _ := result[key].Msgs[0].(*msgstream.InsertMsg)
					accMsgs.Timestamps = append(accMsgs.Timestamps, ts)
					accMsgs.RowIDs = append(accMsgs.RowIDs, rowID)
					accMsgs.RowData = append(accMsgs.RowData, row)
				}
			} else { // every row is a message
				result[key].Msgs = append(result[key].Msgs, insertMsg)
			}
		}
	}

	return result, nil
}
