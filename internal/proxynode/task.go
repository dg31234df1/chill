package proxynode

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/zilliztech/milvus-distributed/internal/allocator"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/indexpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	"github.com/zilliztech/milvus-distributed/internal/proto/querypb"
	"github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
)

type task interface {
	ID() UniqueID       // return ReqID
	SetID(uid UniqueID) // set ReqID
	Type() commonpb.MsgType
	BeginTs() Timestamp
	EndTs() Timestamp
	SetTs(ts Timestamp)
	OnEnqueue() error
	PreExecute() error
	Execute() error
	PostExecute() error
	WaitToFinish() error
	Notify(err error)
}

type BaseInsertTask = msgstream.InsertMsg

type InsertTask struct {
	BaseInsertTask
	Condition
	dataServiceClient DataServiceClient
	result            *milvuspb.InsertResponse
	rowIDAllocator    *allocator.IDAllocator
}

func (it *InsertTask) OnEnqueue() error {
	return nil
}

func (it *InsertTask) SetID(uid UniqueID) {
	it.Base.MsgID = uid
}

func (it *InsertTask) SetTs(ts Timestamp) {
	rowNum := len(it.RowData)
	it.Timestamps = make([]uint64, rowNum)
	for index := range it.Timestamps {
		it.Timestamps[index] = ts
	}
	it.BeginTimestamp = ts
	it.EndTimestamp = ts
}

func (it *InsertTask) BeginTs() Timestamp {
	return it.BeginTimestamp
}

func (it *InsertTask) EndTs() Timestamp {
	return it.EndTimestamp
}

func (it *InsertTask) ID() UniqueID {
	return it.Base.MsgID
}

func (it *InsertTask) Type() commonpb.MsgType {
	return it.Base.MsgType
}

func (it *InsertTask) PreExecute() error {
	it.Base.MsgType = commonpb.MsgType_kInsert
	it.Base.SourceID = Params.ProxyID

	collectionName := it.BaseInsertTask.CollectionName
	if err := ValidateCollectionName(collectionName); err != nil {
		return err
	}
	partitionTag := it.BaseInsertTask.PartitionName
	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}

	return nil
}

func (it *InsertTask) Execute() error {
	collectionName := it.BaseInsertTask.CollectionName
	collSchema, err := globalMetaCache.GetCollectionSchema(collectionName)
	if err != nil {
		return err
	}
	autoID := collSchema.AutoID
	collID, err := globalMetaCache.GetCollectionID(collectionName)
	if err != nil {
		return err
	}
	it.CollectionID = collID
	var partitionID UniqueID
	if len(it.PartitionName) > 0 {
		partitionID, err = globalMetaCache.GetPartitionID(collectionName, it.PartitionName)
		if err != nil {
			return err
		}
	} else {
		partitionID, err = globalMetaCache.GetPartitionID(collectionName, Params.DefaultPartitionTag)
		if err != nil {
			return err
		}
	}
	it.PartitionID = partitionID
	var rowIDBegin UniqueID
	var rowIDEnd UniqueID
	rowNums := len(it.BaseInsertTask.RowData)
	rowIDBegin, rowIDEnd, _ = it.rowIDAllocator.Alloc(uint32(rowNums))

	it.BaseInsertTask.RowIDs = make([]UniqueID, rowNums)
	for i := rowIDBegin; i < rowIDEnd; i++ {
		offset := i - rowIDBegin
		it.BaseInsertTask.RowIDs[offset] = i
	}

	if autoID {
		if it.HashValues == nil || len(it.HashValues) == 0 {
			it.HashValues = make([]uint32, 0)
		}
		for _, rowID := range it.RowIDs {
			hashValue, _ := typeutil.Hash32Int64(rowID)
			it.HashValues = append(it.HashValues, hashValue)
		}
	}

	var tsMsg msgstream.TsMsg = &it.BaseInsertTask
	msgPack := &msgstream.MsgPack{
		BeginTs: it.BeginTs(),
		EndTs:   it.EndTs(),
		Msgs:    make([]msgstream.TsMsg, 1),
	}

	it.result = &milvuspb.InsertResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
		},
		RowIDBegin: rowIDBegin,
		RowIDEnd:   rowIDEnd,
	}

	msgPack.Msgs[0] = tsMsg

	stream, err := globalInsertChannelsMap.getInsertMsgStream(collID)
	if err != nil {
		resp, _ := it.dataServiceClient.GetInsertChannels(&datapb.InsertChannelRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kInsert, // todo
				MsgID:     it.Base.MsgID,            // todo
				Timestamp: 0,                        // todo
				SourceID:  Params.ProxyID,
			},
			DbID:         0, // todo
			CollectionID: collID,
		})
		if resp == nil {
			return errors.New("get insert channels resp is nil")
		}
		if resp.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
			return errors.New(resp.Status.Reason)
		}
		err = globalInsertChannelsMap.createInsertMsgStream(collID, resp.Values)
		if err != nil {
			return err
		}
	}
	stream, err = globalInsertChannelsMap.getInsertMsgStream(collID)
	if err != nil {
		it.result.Status.ErrorCode = commonpb.ErrorCode_UNEXPECTED_ERROR
		it.result.Status.Reason = err.Error()
		return err
	}

	err = stream.Produce(msgPack)
	if err != nil {
		it.result.Status.ErrorCode = commonpb.ErrorCode_UNEXPECTED_ERROR
		it.result.Status.Reason = err.Error()
		return err
	}

	return nil
}

func (it *InsertTask) PostExecute() error {
	return nil
}

type CreateCollectionTask struct {
	Condition
	*milvuspb.CreateCollectionRequest
	masterClient      MasterClient
	dataServiceClient DataServiceClient
	result            *commonpb.Status
	schema            *schemapb.CollectionSchema
}

func (cct *CreateCollectionTask) OnEnqueue() error {
	cct.Base = &commonpb.MsgBase{}
	return nil
}

func (cct *CreateCollectionTask) ID() UniqueID {
	return cct.Base.MsgID
}

func (cct *CreateCollectionTask) SetID(uid UniqueID) {
	cct.Base.MsgID = uid
}

func (cct *CreateCollectionTask) Type() commonpb.MsgType {
	return cct.Base.MsgType
}

func (cct *CreateCollectionTask) BeginTs() Timestamp {
	return cct.Base.Timestamp
}

func (cct *CreateCollectionTask) EndTs() Timestamp {
	return cct.Base.Timestamp
}

func (cct *CreateCollectionTask) SetTs(ts Timestamp) {
	cct.Base.Timestamp = ts
}

func (cct *CreateCollectionTask) PreExecute() error {
	cct.Base.MsgType = commonpb.MsgType_kCreateCollection
	cct.Base.SourceID = Params.ProxyID

	cct.schema = &schemapb.CollectionSchema{}
	err := proto.Unmarshal(cct.Schema, cct.schema)
	if err != nil {
		return err
	}

	if int64(len(cct.schema.Fields)) > Params.MaxFieldNum {
		return errors.New("maximum field's number should be limited to " + strconv.FormatInt(Params.MaxFieldNum, 10))
	}

	// validate collection name
	if err := ValidateCollectionName(cct.schema.Name); err != nil {
		return err
	}

	if err := ValidateDuplicatedFieldName(cct.schema.Fields); err != nil {
		return err
	}

	if err := ValidatePrimaryKey(cct.schema); err != nil {
		return err
	}

	// validate field name
	for _, field := range cct.schema.Fields {
		if err := ValidateFieldName(field.Name); err != nil {
			return err
		}
		if field.DataType == schemapb.DataType_VECTOR_FLOAT || field.DataType == schemapb.DataType_VECTOR_BINARY {
			exist := false
			var dim int64 = 0
			for _, param := range field.TypeParams {
				if param.Key == "dim" {
					exist = true
					tmp, err := strconv.ParseInt(param.Value, 10, 64)
					if err != nil {
						return err
					}
					dim = tmp
					break
				}
			}
			if !exist {
				return errors.New("dimension is not defined in field type params")
			}
			if field.DataType == schemapb.DataType_VECTOR_FLOAT {
				if err := ValidateDimension(dim, false); err != nil {
					return err
				}
			} else {
				if err := ValidateDimension(dim, true); err != nil {
					return err
				}
			}
		}
		if err := ValidateVectorFieldMetricType(field); err != nil {
			return err
		}
	}

	return nil
}

func (cct *CreateCollectionTask) Execute() error {
	var err error
	cct.result, err = cct.masterClient.CreateCollection(cct.CreateCollectionRequest)
	if err != nil {
		return err
	}
	if cct.result.ErrorCode == commonpb.ErrorCode_SUCCESS {
		collID, err := globalMetaCache.GetCollectionID(cct.CollectionName)
		if err != nil {
			return err
		}
		resp, _ := cct.dataServiceClient.GetInsertChannels(&datapb.InsertChannelRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kInsert, // todo
				MsgID:     cct.Base.MsgID,           // todo
				Timestamp: 0,                        // todo
				SourceID:  Params.ProxyID,
			},
			DbID:         0, // todo
			CollectionID: collID,
		})
		if resp == nil {
			return errors.New("get insert channels resp is nil")
		}
		if resp.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
			return errors.New(resp.Status.Reason)
		}
		err = globalInsertChannelsMap.createInsertMsgStream(collID, resp.Values)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cct *CreateCollectionTask) PostExecute() error {
	return nil
}

type DropCollectionTask struct {
	Condition
	*milvuspb.DropCollectionRequest
	masterClient MasterClient
	result       *commonpb.Status
}

func (dct *DropCollectionTask) OnEnqueue() error {
	dct.Base = &commonpb.MsgBase{}
	return nil
}

func (dct *DropCollectionTask) ID() UniqueID {
	return dct.Base.MsgID
}

func (dct *DropCollectionTask) SetID(uid UniqueID) {
	dct.Base.MsgID = uid
}

func (dct *DropCollectionTask) Type() commonpb.MsgType {
	return dct.Base.MsgType
}

func (dct *DropCollectionTask) BeginTs() Timestamp {
	return dct.Base.Timestamp
}

func (dct *DropCollectionTask) EndTs() Timestamp {
	return dct.Base.Timestamp
}

func (dct *DropCollectionTask) SetTs(ts Timestamp) {
	dct.Base.Timestamp = ts
}

func (dct *DropCollectionTask) PreExecute() error {
	dct.Base.MsgType = commonpb.MsgType_kDropCollection
	dct.Base.SourceID = Params.ProxyID

	if err := ValidateCollectionName(dct.CollectionName); err != nil {
		return err
	}
	return nil
}

func (dct *DropCollectionTask) Execute() error {
	collID, err := globalMetaCache.GetCollectionID(dct.CollectionName)
	if err != nil {
		return err
	}

	dct.result, err = dct.masterClient.DropCollection(dct.DropCollectionRequest)
	if err != nil {
		return err
	}

	err = globalInsertChannelsMap.closeInsertMsgStream(collID)
	if err != nil {
		return err
	}

	return nil
}

func (dct *DropCollectionTask) PostExecute() error {
	globalMetaCache.RemoveCollection(dct.CollectionName)
	return nil
}

type SearchTask struct {
	Condition
	internalpb2.SearchRequest
	queryMsgStream msgstream.MsgStream
	resultBuf      chan []*internalpb2.SearchResults
	result         *milvuspb.SearchResults
	query          *milvuspb.SearchRequest
}

func (st *SearchTask) OnEnqueue() error {
	return nil
}

func (st *SearchTask) ID() UniqueID {
	return st.Base.MsgID
}

func (st *SearchTask) SetID(uid UniqueID) {
	st.Base.MsgID = uid
}

func (st *SearchTask) Type() commonpb.MsgType {
	return st.Base.MsgType
}

func (st *SearchTask) BeginTs() Timestamp {
	return st.Base.Timestamp
}

func (st *SearchTask) EndTs() Timestamp {
	return st.Base.Timestamp
}

func (st *SearchTask) SetTs(ts Timestamp) {
	st.Base.Timestamp = ts
}

func (st *SearchTask) PreExecute() error {
	st.Base.MsgType = commonpb.MsgType_kSearch
	st.Base.SourceID = Params.ProxyID

	collectionName := st.query.CollectionName
	_, err := globalMetaCache.GetCollectionID(collectionName)
	if err != nil { // err is not nil if collection not exists
		return err
	}

	if err := ValidateCollectionName(st.query.CollectionName); err != nil {
		return err
	}

	for _, tag := range st.query.PartitionNames {
		if err := ValidatePartitionTag(tag, false); err != nil {
			return err
		}
	}
	st.Base.MsgType = commonpb.MsgType_kSearch
	queryBytes, err := proto.Marshal(st.query)
	if err != nil {
		return err
	}
	st.Query = &commonpb.Blob{
		Value: queryBytes,
	}

	st.ResultChannelID = Params.SearchResultChannelNames[0]
	st.DbID = 0 // todo
	collectionID, err := globalMetaCache.GetCollectionID(collectionName)
	if err != nil { // err is not nil if collection not exists
		return err
	}
	st.CollectionID = collectionID
	st.PartitionIDs = make([]UniqueID, 0)
	for _, partitionName := range st.query.PartitionNames {
		partitionID, err := globalMetaCache.GetPartitionID(collectionName, partitionName)
		if err != nil {
			return err
		}
		st.PartitionIDs = append(st.PartitionIDs, partitionID)
	}
	st.Dsl = st.query.Dsl
	st.PlaceholderGroup = st.query.PlaceholderGroup

	return nil
}

func (st *SearchTask) Execute() error {
	var tsMsg msgstream.TsMsg = &msgstream.SearchMsg{
		SearchRequest: st.SearchRequest,
		BaseMsg: msgstream.BaseMsg{
			HashValues:     []uint32{uint32(Params.ProxyID)},
			BeginTimestamp: st.Base.Timestamp,
			EndTimestamp:   st.Base.Timestamp,
		},
	}
	msgPack := &msgstream.MsgPack{
		BeginTs: st.Base.Timestamp,
		EndTs:   st.Base.Timestamp,
		Msgs:    make([]msgstream.TsMsg, 1),
	}
	msgPack.Msgs[0] = tsMsg
	err := st.queryMsgStream.Produce(msgPack)
	log.Printf("[NodeImpl] length of searchMsg: %v", len(msgPack.Msgs))
	if err != nil {
		log.Printf("[NodeImpl] send search request failed: %v", err)
	}
	return err
}

func (st *SearchTask) PostExecute() error {
	for {
		select {
		case <-st.Ctx().Done():
			log.Print("SearchTask: wait to finish failed, timeout!, taskID:", st.ID())
			return errors.New("SearchTask:wait to finish failed, timeout:" + strconv.FormatInt(st.ID(), 10))
		case searchResults := <-st.resultBuf:
			// fmt.Println("searchResults: ", searchResults)
			filterSearchResult := make([]*internalpb2.SearchResults, 0)
			var filterReason string
			for _, partialSearchResult := range searchResults {
				if partialSearchResult.Status.ErrorCode == commonpb.ErrorCode_SUCCESS {
					filterSearchResult = append(filterSearchResult, partialSearchResult)
					// For debugging, please don't delete.
					//for i := 0; i < len(partialSearchResult.Hits); i++ {
					//	testHits := milvuspb.Hits{}
					//	err := proto.Unmarshal(partialSearchResult.Hits[i], &testHits)
					//	if err != nil {
					//		panic(err)
					//	}
					//	fmt.Println(testHits.IDs)
					//	fmt.Println(testHits.Scores)
					//}
				} else {
					filterReason += partialSearchResult.Status.Reason + "\n"
				}
			}

			availableQueryNodeNum := len(filterSearchResult)
			if availableQueryNodeNum <= 0 {
				st.result = &milvuspb.SearchResults{
					Status: &commonpb.Status{
						ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
						Reason:    filterReason,
					},
				}
				return errors.New(filterReason)
			}

			hits := make([][]*milvuspb.Hits, 0)
			for _, partialSearchResult := range filterSearchResult {
				if partialSearchResult.Hits == nil || len(partialSearchResult.Hits) <= 0 {
					filterReason += "nq is zero\n"
					continue
				}
				partialHits := make([]*milvuspb.Hits, 0)
				for _, bs := range partialSearchResult.Hits {
					partialHit := &milvuspb.Hits{}
					err := proto.Unmarshal(bs, partialHit)
					if err != nil {
						log.Println("unmarshal error")
						return err
					}
					partialHits = append(partialHits, partialHit)
				}
				hits = append(hits, partialHits)
			}

			availableQueryNodeNum = len(hits)
			if availableQueryNodeNum <= 0 {
				st.result = &milvuspb.SearchResults{
					Status: &commonpb.Status{
						ErrorCode: commonpb.ErrorCode_SUCCESS,
						Reason:    filterReason,
					},
				}
				return nil
			}

			nq := len(hits[0])
			if nq <= 0 {
				st.result = &milvuspb.SearchResults{
					Status: &commonpb.Status{
						ErrorCode: commonpb.ErrorCode_SUCCESS,
						Reason:    filterReason,
					},
				}
				return nil
			}

			topk := 0
			getMax := func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}
			for _, hit := range hits {
				topk = getMax(topk, len(hit[0].IDs))
			}
			st.result = &milvuspb.SearchResults{
				Status: &commonpb.Status{
					ErrorCode: 0,
				},
				Hits: make([][]byte, 0),
			}

			const minFloat32 = -1 * float32(math.MaxFloat32)
			for i := 0; i < nq; i++ {
				locs := make([]int, availableQueryNodeNum)
				reducedHits := &milvuspb.Hits{
					IDs:     make([]int64, 0),
					RowData: make([][]byte, 0),
					Scores:  make([]float32, 0),
				}

				for j := 0; j < topk; j++ {
					valid := false
					choice, maxDistance := 0, minFloat32
					for q, loc := range locs { // query num, the number of ways to merge
						if loc >= len(hits[q][i].IDs) {
							continue
						}
						distance := hits[q][i].Scores[loc]
						if distance > maxDistance || (math.Abs(float64(distance-maxDistance)) < math.SmallestNonzeroFloat32 && choice != q) {
							choice = q
							maxDistance = distance
							valid = true
						}
					}
					if !valid {
						break
					}
					choiceOffset := locs[choice]
					// check if distance is valid, `invalid` here means very very big,
					// in this process, distance here is the smallest, so the rest of distance are all invalid
					if hits[choice][i].Scores[choiceOffset] <= minFloat32 {
						break
					}
					reducedHits.IDs = append(reducedHits.IDs, hits[choice][i].IDs[choiceOffset])
					if hits[choice][i].RowData != nil && len(hits[choice][i].RowData) > 0 {
						reducedHits.RowData = append(reducedHits.RowData, hits[choice][i].RowData[choiceOffset])
					}
					reducedHits.Scores = append(reducedHits.Scores, hits[choice][i].Scores[choiceOffset])
					locs[choice]++
				}
				if searchResults[0].MetricType != "IP" {
					for k := range reducedHits.Scores {
						reducedHits.Scores[k] *= -1
					}
				}
				reducedHitsBs, err := proto.Marshal(reducedHits)
				if err != nil {
					log.Println("marshal error")
					return err
				}
				st.result.Hits = append(st.result.Hits, reducedHitsBs)
			}
			return nil
		}
	}
}

type HasCollectionTask struct {
	Condition
	*milvuspb.HasCollectionRequest
	masterClient MasterClient
	result       *milvuspb.BoolResponse
}

func (hct *HasCollectionTask) OnEnqueue() error {
	hct.Base = &commonpb.MsgBase{}
	return nil
}

func (hct *HasCollectionTask) ID() UniqueID {
	return hct.Base.MsgID
}

func (hct *HasCollectionTask) SetID(uid UniqueID) {
	hct.Base.MsgID = uid
}

func (hct *HasCollectionTask) Type() commonpb.MsgType {
	return hct.Base.MsgType
}

func (hct *HasCollectionTask) BeginTs() Timestamp {
	return hct.Base.Timestamp
}

func (hct *HasCollectionTask) EndTs() Timestamp {
	return hct.Base.Timestamp
}

func (hct *HasCollectionTask) SetTs(ts Timestamp) {
	hct.Base.Timestamp = ts
}

func (hct *HasCollectionTask) PreExecute() error {
	hct.Base.MsgType = commonpb.MsgType_kHasCollection
	hct.Base.SourceID = Params.ProxyID

	if err := ValidateCollectionName(hct.CollectionName); err != nil {
		return err
	}
	return nil
}

func (hct *HasCollectionTask) Execute() error {
	var err error
	hct.result, err = hct.masterClient.HasCollection(hct.HasCollectionRequest)
	if hct.result == nil {
		return errors.New("has collection resp is nil")
	}
	if hct.result.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(hct.result.Status.Reason)
	}
	return err
}

func (hct *HasCollectionTask) PostExecute() error {
	return nil
}

type DescribeCollectionTask struct {
	Condition
	*milvuspb.DescribeCollectionRequest
	masterClient MasterClient
	result       *milvuspb.DescribeCollectionResponse
}

func (dct *DescribeCollectionTask) OnEnqueue() error {
	dct.Base = &commonpb.MsgBase{}
	return nil
}

func (dct *DescribeCollectionTask) ID() UniqueID {
	return dct.Base.MsgID
}

func (dct *DescribeCollectionTask) SetID(uid UniqueID) {
	dct.Base.MsgID = uid
}

func (dct *DescribeCollectionTask) Type() commonpb.MsgType {
	return dct.Base.MsgType
}

func (dct *DescribeCollectionTask) BeginTs() Timestamp {
	return dct.Base.Timestamp
}

func (dct *DescribeCollectionTask) EndTs() Timestamp {
	return dct.Base.Timestamp
}

func (dct *DescribeCollectionTask) SetTs(ts Timestamp) {
	dct.Base.Timestamp = ts
}

func (dct *DescribeCollectionTask) PreExecute() error {
	dct.Base.MsgType = commonpb.MsgType_kDescribeCollection
	dct.Base.SourceID = Params.ProxyID

	if err := ValidateCollectionName(dct.CollectionName); err != nil {
		return err
	}
	return nil
}

func (dct *DescribeCollectionTask) Execute() error {
	var err error
	dct.result, err = dct.masterClient.DescribeCollection(dct.DescribeCollectionRequest)
	if dct.result == nil {
		return errors.New("has collection resp is nil")
	}
	if dct.result.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(dct.result.Status.Reason)
	}
	return err
}

func (dct *DescribeCollectionTask) PostExecute() error {
	return nil
}

type GetCollectionsStatisticsTask struct {
	Condition
	*milvuspb.CollectionStatsRequest
	dataServiceClient DataServiceClient
	result            *milvuspb.CollectionStatsResponse
}

func (g *GetCollectionsStatisticsTask) ID() UniqueID {
	return g.Base.MsgID
}

func (g *GetCollectionsStatisticsTask) SetID(uid UniqueID) {
	g.Base.MsgID = uid
}

func (g *GetCollectionsStatisticsTask) Type() commonpb.MsgType {
	return g.Base.MsgType
}

func (g *GetCollectionsStatisticsTask) BeginTs() Timestamp {
	return g.Base.Timestamp
}

func (g *GetCollectionsStatisticsTask) EndTs() Timestamp {
	return g.Base.Timestamp
}

func (g *GetCollectionsStatisticsTask) SetTs(ts Timestamp) {
	g.Base.Timestamp = ts
}

func (g *GetCollectionsStatisticsTask) OnEnqueue() error {
	g.Base = &commonpb.MsgBase{}
	return nil
}

func (g *GetCollectionsStatisticsTask) PreExecute() error {
	g.Base.MsgType = commonpb.MsgType_kGetCollectionStatistics
	g.Base.SourceID = Params.ProxyID
	return nil
}

func (g *GetCollectionsStatisticsTask) Execute() error {
	collID, err := globalMetaCache.GetCollectionID(g.CollectionName)
	if err != nil {
		return err
	}
	req := &datapb.CollectionStatsRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kGetCollectionStatistics,
			MsgID:     g.Base.MsgID,
			Timestamp: g.Base.Timestamp,
			SourceID:  g.Base.SourceID,
		},
		CollectionID: collID,
	}

	result, _ := g.dataServiceClient.GetCollectionStatistics(req)
	if result == nil {
		return errors.New("get collection statistics resp is nil")
	}
	if result.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(result.Status.Reason)
	}
	g.result = &milvuspb.CollectionStatsResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		Stats: result.Stats,
	}
	return nil
}

func (g *GetCollectionsStatisticsTask) PostExecute() error {
	return nil
}

type ShowCollectionsTask struct {
	Condition
	*milvuspb.ShowCollectionRequest
	masterClient MasterClient
	result       *milvuspb.ShowCollectionResponse
	ctx          context.Context
}

func (sct *ShowCollectionsTask) OnEnqueue() error {
	sct.Base = &commonpb.MsgBase{}
	return nil
}

func (sct *ShowCollectionsTask) ID() UniqueID {
	return sct.Base.MsgID
}

func (sct *ShowCollectionsTask) SetID(uid UniqueID) {
	sct.Base.MsgID = uid
}

func (sct *ShowCollectionsTask) Type() commonpb.MsgType {
	return sct.Base.MsgType
}

func (sct *ShowCollectionsTask) BeginTs() Timestamp {
	return sct.Base.Timestamp
}

func (sct *ShowCollectionsTask) EndTs() Timestamp {
	return sct.Base.Timestamp
}

func (sct *ShowCollectionsTask) SetTs(ts Timestamp) {
	sct.Base.Timestamp = ts
}

func (sct *ShowCollectionsTask) PreExecute() error {
	sct.Base.MsgType = commonpb.MsgType_kShowCollections
	sct.Base.SourceID = Params.ProxyID

	return nil
}

func (sct *ShowCollectionsTask) Execute() error {
	var err error
	sct.result, err = sct.masterClient.ShowCollections(sct.ShowCollectionRequest)
	if sct.result == nil {
		return errors.New("get collection statistics resp is nil")
	}
	if sct.result.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(sct.result.Status.Reason)
	}
	return err
}

func (sct *ShowCollectionsTask) PostExecute() error {
	return nil
}

type CreatePartitionTask struct {
	Condition
	*milvuspb.CreatePartitionRequest
	masterClient MasterClient
	result       *commonpb.Status
}

func (cpt *CreatePartitionTask) OnEnqueue() error {
	cpt.Base = &commonpb.MsgBase{}
	return nil
}

func (cpt *CreatePartitionTask) ID() UniqueID {
	return cpt.Base.MsgID
}

func (cpt *CreatePartitionTask) SetID(uid UniqueID) {
	cpt.Base.MsgID = uid
}

func (cpt *CreatePartitionTask) Type() commonpb.MsgType {
	return cpt.Base.MsgType
}

func (cpt *CreatePartitionTask) BeginTs() Timestamp {
	return cpt.Base.Timestamp
}

func (cpt *CreatePartitionTask) EndTs() Timestamp {
	return cpt.Base.Timestamp
}

func (cpt *CreatePartitionTask) SetTs(ts Timestamp) {
	cpt.Base.Timestamp = ts
}

func (cpt *CreatePartitionTask) PreExecute() error {
	cpt.Base.MsgType = commonpb.MsgType_kCreatePartition
	cpt.Base.SourceID = Params.ProxyID

	collName, partitionTag := cpt.CollectionName, cpt.PartitionName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}

	return nil
}

func (cpt *CreatePartitionTask) Execute() (err error) {
	cpt.result, err = cpt.masterClient.CreatePartition(cpt.CreatePartitionRequest)
	if cpt.result == nil {
		return errors.New("get collection statistics resp is nil")
	}
	if cpt.result.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(cpt.result.Reason)
	}
	return err
}

func (cpt *CreatePartitionTask) PostExecute() error {
	return nil
}

type DropPartitionTask struct {
	Condition
	*milvuspb.DropPartitionRequest
	masterClient MasterClient
	result       *commonpb.Status
}

func (dpt *DropPartitionTask) OnEnqueue() error {
	dpt.Base = &commonpb.MsgBase{}
	return nil
}

func (dpt *DropPartitionTask) ID() UniqueID {
	return dpt.Base.MsgID
}

func (dpt *DropPartitionTask) SetID(uid UniqueID) {
	dpt.Base.MsgID = uid
}

func (dpt *DropPartitionTask) Type() commonpb.MsgType {
	return dpt.Base.MsgType
}

func (dpt *DropPartitionTask) BeginTs() Timestamp {
	return dpt.Base.Timestamp
}

func (dpt *DropPartitionTask) EndTs() Timestamp {
	return dpt.Base.Timestamp
}

func (dpt *DropPartitionTask) SetTs(ts Timestamp) {
	dpt.Base.Timestamp = ts
}

func (dpt *DropPartitionTask) PreExecute() error {
	dpt.Base.MsgType = commonpb.MsgType_kDropPartition
	dpt.Base.SourceID = Params.ProxyID

	collName, partitionTag := dpt.CollectionName, dpt.PartitionName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}

	return nil
}

func (dpt *DropPartitionTask) Execute() (err error) {
	dpt.result, err = dpt.masterClient.DropPartition(dpt.DropPartitionRequest)
	if dpt.result == nil {
		return errors.New("get collection statistics resp is nil")
	}
	if dpt.result.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(dpt.result.Reason)
	}
	return err
}

func (dpt *DropPartitionTask) PostExecute() error {
	return nil
}

type HasPartitionTask struct {
	Condition
	*milvuspb.HasPartitionRequest
	masterClient MasterClient
	result       *milvuspb.BoolResponse
}

func (hpt *HasPartitionTask) OnEnqueue() error {
	hpt.Base = &commonpb.MsgBase{}
	return nil
}

func (hpt *HasPartitionTask) ID() UniqueID {
	return hpt.Base.MsgID
}

func (hpt *HasPartitionTask) SetID(uid UniqueID) {
	hpt.Base.MsgID = uid
}

func (hpt *HasPartitionTask) Type() commonpb.MsgType {
	return hpt.Base.MsgType
}

func (hpt *HasPartitionTask) BeginTs() Timestamp {
	return hpt.Base.Timestamp
}

func (hpt *HasPartitionTask) EndTs() Timestamp {
	return hpt.Base.Timestamp
}

func (hpt *HasPartitionTask) SetTs(ts Timestamp) {
	hpt.Base.Timestamp = ts
}

func (hpt *HasPartitionTask) PreExecute() error {
	hpt.Base.MsgType = commonpb.MsgType_kHasPartition
	hpt.Base.SourceID = Params.ProxyID

	collName, partitionTag := hpt.CollectionName, hpt.PartitionName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}
	return nil
}

func (hpt *HasPartitionTask) Execute() (err error) {
	hpt.result, err = hpt.masterClient.HasPartition(hpt.HasPartitionRequest)
	if hpt.result == nil {
		return errors.New("get collection statistics resp is nil")
	}
	if hpt.result.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(hpt.result.Status.Reason)
	}
	return err
}

func (hpt *HasPartitionTask) PostExecute() error {
	return nil
}

type ShowPartitionsTask struct {
	Condition
	*milvuspb.ShowPartitionRequest
	masterClient MasterClient
	result       *milvuspb.ShowPartitionResponse
}

func (spt *ShowPartitionsTask) OnEnqueue() error {
	spt.Base = &commonpb.MsgBase{}
	return nil
}

func (spt *ShowPartitionsTask) ID() UniqueID {
	return spt.Base.MsgID
}

func (spt *ShowPartitionsTask) SetID(uid UniqueID) {
	spt.Base.MsgID = uid
}

func (spt *ShowPartitionsTask) Type() commonpb.MsgType {
	return spt.Base.MsgType
}

func (spt *ShowPartitionsTask) BeginTs() Timestamp {
	return spt.Base.Timestamp
}

func (spt *ShowPartitionsTask) EndTs() Timestamp {
	return spt.Base.Timestamp
}

func (spt *ShowPartitionsTask) SetTs(ts Timestamp) {
	spt.Base.Timestamp = ts
}

func (spt *ShowPartitionsTask) PreExecute() error {
	spt.Base.MsgType = commonpb.MsgType_kShowPartitions
	spt.Base.SourceID = Params.ProxyID

	if err := ValidateCollectionName(spt.CollectionName); err != nil {
		return err
	}
	return nil
}

func (spt *ShowPartitionsTask) Execute() error {
	var err error
	spt.result, err = spt.masterClient.ShowPartitions(spt.ShowPartitionRequest)
	if spt.result == nil {
		return errors.New("get collection statistics resp is nil")
	}
	if spt.result.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(spt.result.Status.Reason)
	}
	return err
}

func (spt *ShowPartitionsTask) PostExecute() error {
	return nil
}

type CreateIndexTask struct {
	Condition
	*milvuspb.CreateIndexRequest
	masterClient MasterClient
	result       *commonpb.Status
}

func (cit *CreateIndexTask) OnEnqueue() error {
	cit.Base = &commonpb.MsgBase{}
	return nil
}

func (cit *CreateIndexTask) ID() UniqueID {
	return cit.Base.MsgID
}

func (cit *CreateIndexTask) SetID(uid UniqueID) {
	cit.Base.MsgID = uid
}

func (cit *CreateIndexTask) Type() commonpb.MsgType {
	return cit.Base.MsgType
}

func (cit *CreateIndexTask) BeginTs() Timestamp {
	return cit.Base.Timestamp
}

func (cit *CreateIndexTask) EndTs() Timestamp {
	return cit.Base.Timestamp
}

func (cit *CreateIndexTask) SetTs(ts Timestamp) {
	cit.Base.Timestamp = ts
}

func (cit *CreateIndexTask) PreExecute() error {
	cit.Base.MsgType = commonpb.MsgType_kCreateIndex
	cit.Base.SourceID = Params.ProxyID

	collName, fieldName := cit.CollectionName, cit.FieldName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidateFieldName(fieldName); err != nil {
		return err
	}

	return nil
}

func (cit *CreateIndexTask) Execute() error {
	var err error
	cit.result, err = cit.masterClient.CreateIndex(cit.CreateIndexRequest)
	if cit.result == nil {
		return errors.New("get collection statistics resp is nil")
	}
	if cit.result.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(cit.result.Reason)
	}
	return err
}

func (cit *CreateIndexTask) PostExecute() error {
	return nil
}

type DescribeIndexTask struct {
	Condition
	*milvuspb.DescribeIndexRequest
	masterClient MasterClient
	result       *milvuspb.DescribeIndexResponse
}

func (dit *DescribeIndexTask) OnEnqueue() error {
	dit.Base = &commonpb.MsgBase{}
	return nil
}

func (dit *DescribeIndexTask) ID() UniqueID {
	return dit.Base.MsgID
}

func (dit *DescribeIndexTask) SetID(uid UniqueID) {
	dit.Base.MsgID = uid
}

func (dit *DescribeIndexTask) Type() commonpb.MsgType {
	return dit.Base.MsgType
}

func (dit *DescribeIndexTask) BeginTs() Timestamp {
	return dit.Base.Timestamp
}

func (dit *DescribeIndexTask) EndTs() Timestamp {
	return dit.Base.Timestamp
}

func (dit *DescribeIndexTask) SetTs(ts Timestamp) {
	dit.Base.Timestamp = ts
}

func (dit *DescribeIndexTask) PreExecute() error {
	dit.Base.MsgType = commonpb.MsgType_kDescribeIndex
	dit.Base.SourceID = Params.ProxyID

	collName, fieldName := dit.CollectionName, dit.FieldName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidateFieldName(fieldName); err != nil {
		return err
	}

	// only support default index name for now. @2021.02.18
	if dit.IndexName == "" {
		dit.IndexName = Params.DefaultIndexName
	}

	return nil
}

func (dit *DescribeIndexTask) Execute() error {
	var err error
	dit.result, err = dit.masterClient.DescribeIndex(dit.DescribeIndexRequest)
	log.Println("YYYYY:", dit.result)
	if dit.result == nil {
		return errors.New("get collection statistics resp is nil")
	}
	if dit.result.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(dit.result.Status.Reason)
	}
	return err
}

func (dit *DescribeIndexTask) PostExecute() error {
	return nil
}

type DropIndexTask struct {
	Condition
	*milvuspb.DropIndexRequest
	masterClient MasterClient
	result       *commonpb.Status
}

func (dit *DropIndexTask) OnEnqueue() error {
	dit.Base = &commonpb.MsgBase{}
	return nil
}

func (dit *DropIndexTask) ID() UniqueID {
	return dit.Base.MsgID
}

func (dit *DropIndexTask) SetID(uid UniqueID) {
	dit.Base.MsgID = uid
}

func (dit *DropIndexTask) Type() commonpb.MsgType {
	return dit.Base.MsgType
}

func (dit *DropIndexTask) BeginTs() Timestamp {
	return dit.Base.Timestamp
}

func (dit *DropIndexTask) EndTs() Timestamp {
	return dit.Base.Timestamp
}

func (dit *DropIndexTask) SetTs(ts Timestamp) {
	dit.Base.Timestamp = ts
}

func (dit *DropIndexTask) PreExecute() error {
	dit.Base.MsgType = commonpb.MsgType_kDropIndex
	dit.Base.SourceID = Params.ProxyID

	collName, fieldName := dit.CollectionName, dit.FieldName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidateFieldName(fieldName); err != nil {
		return err
	}

	return nil
}

func (dit *DropIndexTask) Execute() error {
	var err error
	dit.result, err = dit.masterClient.DropIndex(dit.DropIndexRequest)
	if dit.result == nil {
		return errors.New("drop index resp is nil")
	}
	if dit.result.ErrorCode != commonpb.ErrorCode_SUCCESS {
		return errors.New(dit.result.Reason)
	}
	return err
}

func (dit *DropIndexTask) PostExecute() error {
	return nil
}

type GetIndexStateTask struct {
	Condition
	*milvuspb.IndexStateRequest
	indexServiceClient IndexServiceClient
	masterClient       MasterClient
	result             *milvuspb.IndexStateResponse
}

func (dipt *GetIndexStateTask) OnEnqueue() error {
	dipt.Base = &commonpb.MsgBase{}
	return nil
}

func (dipt *GetIndexStateTask) ID() UniqueID {
	return dipt.Base.MsgID
}

func (dipt *GetIndexStateTask) SetID(uid UniqueID) {
	dipt.Base.MsgID = uid
}

func (dipt *GetIndexStateTask) Type() commonpb.MsgType {
	return dipt.Base.MsgType
}

func (dipt *GetIndexStateTask) BeginTs() Timestamp {
	return dipt.Base.Timestamp
}

func (dipt *GetIndexStateTask) EndTs() Timestamp {
	return dipt.Base.Timestamp
}

func (dipt *GetIndexStateTask) SetTs(ts Timestamp) {
	dipt.Base.Timestamp = ts
}

func (dipt *GetIndexStateTask) PreExecute() error {
	dipt.Base.MsgType = commonpb.MsgType_kGetIndexState
	dipt.Base.SourceID = Params.ProxyID

	collName, fieldName := dipt.CollectionName, dipt.FieldName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidateFieldName(fieldName); err != nil {
		return err
	}

	return nil
}

func (dipt *GetIndexStateTask) Execute() error {
	collectionName := dipt.CollectionName
	collectionID, err := globalMetaCache.GetCollectionID(collectionName)
	if err != nil { // err is not nil if collection not exists
		return err
	}

	showPartitionRequest := &milvuspb.ShowPartitionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kShowPartitions,
			MsgID:     dipt.Base.MsgID,
			Timestamp: dipt.Base.Timestamp,
			SourceID:  Params.ProxyID,
		},
		DbName:         dipt.DbName,
		CollectionName: collectionName,
		CollectionID:   collectionID,
	}
	partitions, err := dipt.masterClient.ShowPartitions(showPartitionRequest)
	if err != nil {
		return err
	}

	if dipt.IndexName == "" {
		dipt.IndexName = Params.DefaultIndexName
	}

	describeIndexReq := milvuspb.DescribeIndexRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kDescribeIndex,
			MsgID:     dipt.Base.MsgID,
			Timestamp: dipt.Base.Timestamp,
			SourceID:  Params.ProxyID,
		},
		DbName:         dipt.DbName,
		CollectionName: dipt.CollectionName,
		FieldName:      dipt.FieldName,
		IndexName:      dipt.IndexName,
	}

	indexDescriptionResp, err2 := dipt.masterClient.DescribeIndex(&describeIndexReq)
	if err2 != nil {
		return err2
	}

	matchIndexID := int64(-1)
	foundIndexID := false
	for _, desc := range indexDescriptionResp.IndexDescriptions {
		if desc.IndexName == dipt.IndexName {
			matchIndexID = desc.IndexID
			foundIndexID = true
			break
		}
	}
	if !foundIndexID {
		return errors.New(fmt.Sprint("Can't found IndexID for indexName", dipt.IndexName))
	}

	var allSegmentIDs []UniqueID
	for _, partitionID := range partitions.PartitionIDs {
		showSegmentsRequest := &milvuspb.ShowSegmentRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kShowSegment,
				MsgID:     dipt.Base.MsgID,
				Timestamp: dipt.Base.Timestamp,
				SourceID:  Params.ProxyID,
			},
			CollectionID: collectionID,
			PartitionID:  partitionID,
		}
		segments, err := dipt.masterClient.ShowSegments(showSegmentsRequest)
		if err != nil {
			return err
		}
		if segments.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
			return errors.New(segments.Status.Reason)
		}
		allSegmentIDs = append(allSegmentIDs, segments.SegmentIDs...)
	}

	getIndexStatesRequest := &indexpb.IndexStatesRequest{
		IndexBuildIDs: make([]UniqueID, 0),
	}

	for _, segmentID := range allSegmentIDs {
		describeSegmentRequest := &milvuspb.DescribeSegmentRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDescribeSegment,
				MsgID:     dipt.Base.MsgID,
				Timestamp: dipt.Base.Timestamp,
				SourceID:  Params.ProxyID,
			},
			CollectionID: collectionID,
			SegmentID:    segmentID,
		}
		segmentDesc, err := dipt.masterClient.DescribeSegment(describeSegmentRequest)
		if err != nil {
			return err
		}
		if segmentDesc.IndexID == matchIndexID {
			getIndexStatesRequest.IndexBuildIDs = append(getIndexStatesRequest.IndexBuildIDs, segmentDesc.BuildID)
		}
	}

	log.Println("GetIndexState:: len of allSegmentIDs:", len(allSegmentIDs), " len of IndexBuildIDs", len(getIndexStatesRequest.IndexBuildIDs))
	if len(allSegmentIDs) != len(getIndexStatesRequest.IndexBuildIDs) {
		dipt.result = &milvuspb.IndexStateResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_SUCCESS,
				Reason:    "",
			},
			State: commonpb.IndexState_INPROGRESS,
		}
		return err
	}

	states, err := dipt.indexServiceClient.GetIndexStates(getIndexStatesRequest)
	if err != nil {
		return err
	}

	if states.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		dipt.result = &milvuspb.IndexStateResponse{
			Status: states.Status,
			State:  commonpb.IndexState_FAILED,
		}

		return nil
	}

	for _, state := range states.States {
		if state.State != commonpb.IndexState_FINISHED {
			dipt.result = &milvuspb.IndexStateResponse{
				Status: states.Status,
				State:  state.State,
			}

			return nil
		}
	}

	dipt.result = &milvuspb.IndexStateResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		State: commonpb.IndexState_FINISHED,
	}

	return nil
}

func (dipt *GetIndexStateTask) PostExecute() error {
	return nil
}

type FlushTask struct {
	Condition
	*milvuspb.FlushRequest
	dataServiceClient DataServiceClient
	result            *commonpb.Status
}

func (ft *FlushTask) OnEnqueue() error {
	ft.Base = &commonpb.MsgBase{}
	return nil
}

func (ft *FlushTask) ID() UniqueID {
	return ft.Base.MsgID
}

func (ft *FlushTask) SetID(uid UniqueID) {
	ft.Base.MsgID = uid
}

func (ft *FlushTask) Type() commonpb.MsgType {
	return ft.Base.MsgType
}

func (ft *FlushTask) BeginTs() Timestamp {
	return ft.Base.Timestamp
}

func (ft *FlushTask) EndTs() Timestamp {
	return ft.Base.Timestamp
}

func (ft *FlushTask) SetTs(ts Timestamp) {
	ft.Base.Timestamp = ts
}

func (ft *FlushTask) PreExecute() error {
	ft.Base.MsgType = commonpb.MsgType_kFlush
	ft.Base.SourceID = Params.ProxyID
	return nil
}

func (ft *FlushTask) Execute() error {
	for _, collName := range ft.CollectionNames {
		collID, err := globalMetaCache.GetCollectionID(collName)
		if err != nil {
			return err
		}
		flushReq := &datapb.FlushRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kFlush,
				MsgID:     ft.Base.MsgID,
				Timestamp: ft.Base.Timestamp,
				SourceID:  ft.Base.SourceID,
			},
			DbID:         0,
			CollectionID: collID,
		}
		var status *commonpb.Status
		status, _ = ft.dataServiceClient.Flush(flushReq)
		if status == nil {
			return errors.New("flush resp is nil")
		}
		if status.ErrorCode != commonpb.ErrorCode_SUCCESS {
			return errors.New(status.Reason)
		}
	}
	ft.result = &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
	}
	return nil
}

func (ft *FlushTask) PostExecute() error {
	return nil
}

type LoadCollectionTask struct {
	Condition
	*milvuspb.LoadCollectionRequest
	queryserviceClient QueryServiceClient
	result             *commonpb.Status
}

func (lct *LoadCollectionTask) OnEnqueue() error {
	lct.Base = &commonpb.MsgBase{}
	return nil
}

func (lct *LoadCollectionTask) ID() UniqueID {
	return lct.Base.MsgID
}

func (lct *LoadCollectionTask) SetID(uid UniqueID) {
	lct.Base.MsgID = uid
}

func (lct *LoadCollectionTask) Type() commonpb.MsgType {
	return lct.Base.MsgType
}

func (lct *LoadCollectionTask) BeginTs() Timestamp {
	return lct.Base.Timestamp
}

func (lct *LoadCollectionTask) EndTs() Timestamp {
	return lct.Base.Timestamp
}

func (lct *LoadCollectionTask) SetTs(ts Timestamp) {
	lct.Base.Timestamp = ts
}

func (lct *LoadCollectionTask) PreExecute() error {
	lct.Base.MsgType = commonpb.MsgType_kLoadCollection
	lct.Base.SourceID = Params.ProxyID

	collName := lct.CollectionName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	return nil
}

func (lct *LoadCollectionTask) Execute() (err error) {
	collID, err := globalMetaCache.GetCollectionID(lct.CollectionName)
	if err != nil {
		return err
	}
	collSchema, err := globalMetaCache.GetCollectionSchema(lct.CollectionName)
	if err != nil {
		return err
	}

	request := &querypb.LoadCollectionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kLoadCollection,
			MsgID:     lct.Base.MsgID,
			Timestamp: lct.Base.Timestamp,
			SourceID:  lct.Base.SourceID,
		},
		DbID:         0,
		CollectionID: collID,
		Schema:       collSchema,
	}
	lct.result, err = lct.queryserviceClient.LoadCollection(request)
	return err
}

func (lct *LoadCollectionTask) PostExecute() error {
	return nil
}

type ReleaseCollectionTask struct {
	Condition
	*milvuspb.ReleaseCollectionRequest
	queryserviceClient QueryServiceClient
	result             *commonpb.Status
}

func (rct *ReleaseCollectionTask) OnEnqueue() error {
	rct.Base = &commonpb.MsgBase{}
	return nil
}

func (rct *ReleaseCollectionTask) ID() UniqueID {
	return rct.Base.MsgID
}

func (rct *ReleaseCollectionTask) SetID(uid UniqueID) {
	rct.Base.MsgID = uid
}

func (rct *ReleaseCollectionTask) Type() commonpb.MsgType {
	return rct.Base.MsgType
}

func (rct *ReleaseCollectionTask) BeginTs() Timestamp {
	return rct.Base.Timestamp
}

func (rct *ReleaseCollectionTask) EndTs() Timestamp {
	return rct.Base.Timestamp
}

func (rct *ReleaseCollectionTask) SetTs(ts Timestamp) {
	rct.Base.Timestamp = ts
}

func (rct *ReleaseCollectionTask) PreExecute() error {
	rct.Base.MsgType = commonpb.MsgType_kReleaseCollection
	rct.Base.SourceID = Params.ProxyID

	collName := rct.CollectionName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	return nil
}

func (rct *ReleaseCollectionTask) Execute() (err error) {
	collID, err := globalMetaCache.GetCollectionID(rct.CollectionName)
	if err != nil {
		return err
	}
	request := &querypb.ReleaseCollectionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kReleaseCollection,
			MsgID:     rct.Base.MsgID,
			Timestamp: rct.Base.Timestamp,
			SourceID:  rct.Base.SourceID,
		},
		DbID:         0,
		CollectionID: collID,
	}
	rct.result, err = rct.queryserviceClient.ReleaseCollection(request)
	return err
}

func (rct *ReleaseCollectionTask) PostExecute() error {
	return nil
}

type LoadPartitionTask struct {
	Condition
	*milvuspb.LoadPartitonRequest
	queryserviceClient QueryServiceClient
	result             *commonpb.Status
}

func (lpt *LoadPartitionTask) OnEnqueue() error {
	lpt.Base = &commonpb.MsgBase{}
	return nil
}

func (lpt *LoadPartitionTask) ID() UniqueID {
	return lpt.Base.MsgID
}

func (lpt *LoadPartitionTask) SetID(uid UniqueID) {
	lpt.Base.MsgID = uid
}

func (lpt *LoadPartitionTask) Type() commonpb.MsgType {
	return lpt.Base.MsgType
}

func (lpt *LoadPartitionTask) BeginTs() Timestamp {
	return lpt.Base.Timestamp
}

func (lpt *LoadPartitionTask) EndTs() Timestamp {
	return lpt.Base.Timestamp
}

func (lpt *LoadPartitionTask) SetTs(ts Timestamp) {
	lpt.Base.Timestamp = ts
}

func (lpt *LoadPartitionTask) PreExecute() error {
	lpt.Base.MsgType = commonpb.MsgType_kLoadPartition
	lpt.Base.SourceID = Params.ProxyID

	collName := lpt.CollectionName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	return nil
}

func (lpt *LoadPartitionTask) Execute() (err error) {
	var partitionIDs []int64
	collID, err := globalMetaCache.GetCollectionID(lpt.CollectionName)
	if err != nil {
		return err
	}
	collSchema, err := globalMetaCache.GetCollectionSchema(lpt.CollectionName)
	if err != nil {
		return err
	}
	for _, partitionName := range lpt.PartitionNames {
		partitionID, err := globalMetaCache.GetPartitionID(lpt.CollectionName, partitionName)
		if err != nil {
			return err
		}
		partitionIDs = append(partitionIDs, partitionID)
	}
	request := &querypb.LoadPartitionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kLoadPartition,
			MsgID:     lpt.Base.MsgID,
			Timestamp: lpt.Base.Timestamp,
			SourceID:  lpt.Base.SourceID,
		},
		DbID:         0,
		CollectionID: collID,
		PartitionIDs: partitionIDs,
		Schema:       collSchema,
	}
	lpt.result, err = lpt.queryserviceClient.LoadPartitions(request)
	return err
}

func (lpt *LoadPartitionTask) PostExecute() error {
	return nil
}

type ReleasePartitionTask struct {
	Condition
	*milvuspb.ReleasePartitionRequest
	queryserviceClient QueryServiceClient
	result             *commonpb.Status
}

func (rpt *ReleasePartitionTask) OnEnqueue() error {
	rpt.Base = &commonpb.MsgBase{}
	return nil
}

func (rpt *ReleasePartitionTask) ID() UniqueID {
	return rpt.Base.MsgID
}

func (rpt *ReleasePartitionTask) SetID(uid UniqueID) {
	rpt.Base.MsgID = uid
}

func (rpt *ReleasePartitionTask) Type() commonpb.MsgType {
	return rpt.Base.MsgType
}

func (rpt *ReleasePartitionTask) BeginTs() Timestamp {
	return rpt.Base.Timestamp
}

func (rpt *ReleasePartitionTask) EndTs() Timestamp {
	return rpt.Base.Timestamp
}

func (rpt *ReleasePartitionTask) SetTs(ts Timestamp) {
	rpt.Base.Timestamp = ts
}

func (rpt *ReleasePartitionTask) PreExecute() error {
	rpt.Base.MsgType = commonpb.MsgType_kReleasePartition
	rpt.Base.SourceID = Params.ProxyID

	collName := rpt.CollectionName

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	return nil
}

func (rpt *ReleasePartitionTask) Execute() (err error) {
	var partitionIDs []int64
	collID, err := globalMetaCache.GetCollectionID(rpt.CollectionName)
	if err != nil {
		return err
	}
	for _, partitionName := range rpt.PartitionNames {
		partitionID, err := globalMetaCache.GetPartitionID(rpt.CollectionName, partitionName)
		if err != nil {
			return err
		}
		partitionIDs = append(partitionIDs, partitionID)
	}
	request := &querypb.ReleasePartitionRequest{
		Base: &commonpb.MsgBase{
			MsgType:   commonpb.MsgType_kReleasePartition,
			MsgID:     rpt.Base.MsgID,
			Timestamp: rpt.Base.Timestamp,
			SourceID:  rpt.Base.SourceID,
		},
		DbID:         0,
		CollectionID: collID,
		PartitionIDs: partitionIDs,
	}
	rpt.result, err = rpt.queryserviceClient.ReleasePartitions(request)
	return err
}

func (rpt *ReleasePartitionTask) PostExecute() error {
	return nil
}
