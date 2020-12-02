package proxy

import (
	"context"
	"errors"
	"log"
	"math"
	"strconv"

	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"

	"github.com/golang/protobuf/proto"

	"github.com/zilliztech/milvus-distributed/internal/allocator"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/masterpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/servicepb"
)

type task interface {
	ID() UniqueID       // return ReqID
	SetID(uid UniqueID) // set ReqID
	Type() internalpb.MsgType
	BeginTs() Timestamp
	EndTs() Timestamp
	SetTs(ts Timestamp)
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
	result                *servicepb.IntegerRangeResponse
	manipulationMsgStream *msgstream.PulsarMsgStream
	ctx                   context.Context
	rowIDAllocator        *allocator.IDAllocator
}

func (it *InsertTask) SetID(uid UniqueID) {
	it.ReqID = uid
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
	return it.ReqID
}

func (it *InsertTask) Type() internalpb.MsgType {
	return it.MsgType
}

func (it *InsertTask) PreExecute() error {
	collectionName := it.BaseInsertTask.CollectionName
	if err := ValidateCollectionName(collectionName); err != nil {
		return err
	}
	partitionTag := it.BaseInsertTask.PartitionTag
	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}

	return nil
}

func (it *InsertTask) Execute() error {
	collectionName := it.BaseInsertTask.CollectionName
	if !globalMetaCache.Hit(collectionName) {
		err := globalMetaCache.Update(collectionName)
		if err != nil {
			return err
		}
	}
	description, err := globalMetaCache.Get(collectionName)
	if err != nil || description == nil {
		return err
	}
	autoID := description.Schema.AutoID
	var rowIDBegin UniqueID
	var rowIDEnd UniqueID
	if autoID || true {
		if it.HashValues == nil || len(it.HashValues) == 0 {
			it.HashValues = make([]uint32, 0)
		}
		rowNums := len(it.BaseInsertTask.RowData)
		rowIDBegin, rowIDEnd, _ = it.rowIDAllocator.Alloc(uint32(rowNums))
		it.BaseInsertTask.RowIDs = make([]UniqueID, rowNums)
		for i := rowIDBegin; i < rowIDEnd; i++ {
			offset := i - rowIDBegin
			it.BaseInsertTask.RowIDs[offset] = i
			hashValue, _ := typeutil.Hash32Int64(i)
			it.HashValues = append(it.HashValues, hashValue)
		}
	}

	var tsMsg msgstream.TsMsg = &it.BaseInsertTask
	msgPack := &msgstream.MsgPack{
		BeginTs: it.BeginTs(),
		EndTs:   it.EndTs(),
		Msgs:    make([]msgstream.TsMsg, 1),
	}
	msgPack.Msgs[0] = tsMsg
	err = it.manipulationMsgStream.Produce(msgPack)
	it.result = &servicepb.IntegerRangeResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
		},
		Begin: rowIDBegin,
		End:   rowIDEnd,
	}
	if err != nil {
		it.result.Status.ErrorCode = commonpb.ErrorCode_UNEXPECTED_ERROR
		it.result.Status.Reason = err.Error()
	}
	return nil
}

func (it *InsertTask) PostExecute() error {
	return nil
}

type CreateCollectionTask struct {
	Condition
	internalpb.CreateCollectionRequest
	masterClient masterpb.MasterClient
	result       *commonpb.Status
	ctx          context.Context
	schema       *schemapb.CollectionSchema
}

func (cct *CreateCollectionTask) ID() UniqueID {
	return cct.ReqID
}

func (cct *CreateCollectionTask) SetID(uid UniqueID) {
	cct.ReqID = uid
}

func (cct *CreateCollectionTask) Type() internalpb.MsgType {
	return cct.MsgType
}

func (cct *CreateCollectionTask) BeginTs() Timestamp {
	return cct.Timestamp
}

func (cct *CreateCollectionTask) EndTs() Timestamp {
	return cct.Timestamp
}

func (cct *CreateCollectionTask) SetTs(ts Timestamp) {
	cct.Timestamp = ts
}

func (cct *CreateCollectionTask) PreExecute() error {
	if int64(len(cct.schema.Fields)) > Params.MaxFieldNum() {
		return errors.New("maximum field's number should be limited to " + strconv.FormatInt(Params.MaxFieldNum(), 10))
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
	schemaBytes, _ := proto.Marshal(cct.schema)
	cct.CreateCollectionRequest.Schema.Value = schemaBytes
	resp, err := cct.masterClient.CreateCollection(cct.ctx, &cct.CreateCollectionRequest)
	if err != nil {
		log.Printf("create collection failed, error= %v", err)
		cct.result = &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    err.Error(),
		}
	} else {
		cct.result = resp
	}
	return err
}

func (cct *CreateCollectionTask) PostExecute() error {
	return nil
}

type DropCollectionTask struct {
	Condition
	internalpb.DropCollectionRequest
	masterClient masterpb.MasterClient
	result       *commonpb.Status
	ctx          context.Context
}

func (dct *DropCollectionTask) ID() UniqueID {
	return dct.ReqID
}

func (dct *DropCollectionTask) SetID(uid UniqueID) {
	dct.ReqID = uid
}

func (dct *DropCollectionTask) Type() internalpb.MsgType {
	return dct.MsgType
}

func (dct *DropCollectionTask) BeginTs() Timestamp {
	return dct.Timestamp
}

func (dct *DropCollectionTask) EndTs() Timestamp {
	return dct.Timestamp
}

func (dct *DropCollectionTask) SetTs(ts Timestamp) {
	dct.Timestamp = ts
}

func (dct *DropCollectionTask) PreExecute() error {
	if err := ValidateCollectionName(dct.CollectionName.CollectionName); err != nil {
		return err
	}
	return nil
}

func (dct *DropCollectionTask) Execute() error {
	resp, err := dct.masterClient.DropCollection(dct.ctx, &dct.DropCollectionRequest)
	if err != nil {
		log.Printf("drop collection failed, error= %v", err)
		dct.result = &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    err.Error(),
		}
	} else {
		dct.result = resp
	}
	return err
}

func (dct *DropCollectionTask) PostExecute() error {
	if globalMetaCache.Hit(dct.CollectionName.CollectionName) {
		return globalMetaCache.Remove(dct.CollectionName.CollectionName)
	}
	return nil
}

type QueryTask struct {
	Condition
	internalpb.SearchRequest
	queryMsgStream *msgstream.PulsarMsgStream
	resultBuf      chan []*internalpb.SearchResult
	result         *servicepb.QueryResult
	ctx            context.Context
	query          *servicepb.Query
}

func (qt *QueryTask) ID() UniqueID {
	return qt.ReqID
}

func (qt *QueryTask) SetID(uid UniqueID) {
	qt.ReqID = uid
}

func (qt *QueryTask) Type() internalpb.MsgType {
	return qt.MsgType
}

func (qt *QueryTask) BeginTs() Timestamp {
	return qt.Timestamp
}

func (qt *QueryTask) EndTs() Timestamp {
	return qt.Timestamp
}

func (qt *QueryTask) SetTs(ts Timestamp) {
	qt.Timestamp = ts
}

func (qt *QueryTask) PreExecute() error {
	collectionName := qt.query.CollectionName
	if !globalMetaCache.Hit(collectionName) {
		err := globalMetaCache.Update(collectionName)
		if err != nil {
			return err
		}
	}
	_, err := globalMetaCache.Get(collectionName)
	if err != nil { // err is not nil if collection not exists
		return err
	}

	if err := ValidateCollectionName(qt.query.CollectionName); err != nil {
		return err
	}

	for _, tag := range qt.query.PartitionTags {
		if err := ValidatePartitionTag(tag, false); err != nil {
			return err
		}
	}
	qt.MsgType = internalpb.MsgType_kSearch
	if qt.query.PartitionTags == nil || len(qt.query.PartitionTags) <= 0 {
		qt.query.PartitionTags = []string{Params.defaultPartitionTag()}
	}
	queryBytes, err := proto.Marshal(qt.query)
	if err != nil {
		return err
	}
	qt.Query = &commonpb.Blob{
		Value: queryBytes,
	}
	return nil
}

func (qt *QueryTask) Execute() error {
	var tsMsg msgstream.TsMsg = &msgstream.SearchMsg{
		SearchRequest: qt.SearchRequest,
		BaseMsg: msgstream.BaseMsg{
			HashValues:     []uint32{uint32(Params.ProxyID())},
			BeginTimestamp: qt.Timestamp,
			EndTimestamp:   qt.Timestamp,
		},
	}
	msgPack := &msgstream.MsgPack{
		BeginTs: qt.Timestamp,
		EndTs:   qt.Timestamp,
		Msgs:    make([]msgstream.TsMsg, 1),
	}
	msgPack.Msgs[0] = tsMsg
	err := qt.queryMsgStream.Produce(msgPack)
	log.Printf("[Proxy] length of searchMsg: %v", len(msgPack.Msgs))
	if err != nil {
		log.Printf("[Proxy] send search request failed: %v", err)
	}
	return err
}

func (qt *QueryTask) PostExecute() error {
	for {
		select {
		case <-qt.ctx.Done():
			log.Print("wait to finish failed, timeout!")
			return errors.New("wait to finish failed, timeout")
		case searchResults := <-qt.resultBuf:
			filterSearchResult := make([]*internalpb.SearchResult, 0)
			var filterReason string
			for _, partialSearchResult := range searchResults {
				if partialSearchResult.Status.ErrorCode == commonpb.ErrorCode_SUCCESS {
					filterSearchResult = append(filterSearchResult, partialSearchResult)
				} else {
					filterReason += partialSearchResult.Status.Reason + "\n"
				}
			}

			rlen := len(filterSearchResult) // query node num
			if rlen <= 0 {
				qt.result = &servicepb.QueryResult{
					Status: &commonpb.Status{
						ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
						Reason:    filterReason,
					},
				}
				return errors.New(filterReason)
			}

			n := len(filterSearchResult[0].Hits) // n
			if n <= 0 {
				qt.result = &servicepb.QueryResult{}
				return nil
			}

			hits := make([][]*servicepb.Hits, rlen)
			for i, partialSearchResult := range filterSearchResult {
				hits[i] = make([]*servicepb.Hits, n)
				for j, bs := range partialSearchResult.Hits {
					hits[i][j] = &servicepb.Hits{}
					err := proto.Unmarshal(bs, hits[i][j])
					if err != nil {
						log.Println("unmarshal error")
						return err
					}
				}
			}

			k := len(hits[0][0].IDs)
			qt.result = &servicepb.QueryResult{
				Status: &commonpb.Status{
					ErrorCode: 0,
				},
				Hits: make([][]byte, 0),
			}

			for i := 0; i < n; i++ { // n
				locs := make([]int, rlen)
				reducedHits := &servicepb.Hits{
					IDs:     make([]int64, 0),
					RowData: make([][]byte, 0),
					Scores:  make([]float32, 0),
				}

				for j := 0; j < k; j++ { // k
					choice, minDistance := 0, float32(math.MaxFloat32)
					for q, loc := range locs { // query num, the number of ways to merge
						distance := hits[q][i].Scores[loc]
						if distance < minDistance {
							choice = q
							minDistance = distance
						}
					}
					choiceOffset := locs[choice]
					// check if distance is valid, `invalid` here means very very big,
					// in this process, distance here is the smallest, so the rest of distance are all invalid
					if hits[choice][i].Scores[choiceOffset] >= float32(math.MaxFloat32) {
						break
					}
					reducedHits.IDs = append(reducedHits.IDs, hits[choice][i].IDs[choiceOffset])
					if hits[choice][i].RowData != nil && len(hits[choice][i].RowData) > 0 {
						reducedHits.RowData = append(reducedHits.RowData, hits[choice][i].RowData[choiceOffset])
					}
					reducedHits.Scores = append(reducedHits.Scores, hits[choice][i].Scores[choiceOffset])
					locs[choice]++
				}
				reducedHitsBs, err := proto.Marshal(reducedHits)
				if err != nil {
					log.Println("marshal error")
					return err
				}
				qt.result.Hits = append(qt.result.Hits, reducedHitsBs)
			}
			return nil
		}
	}
}

type HasCollectionTask struct {
	Condition
	internalpb.HasCollectionRequest
	masterClient masterpb.MasterClient
	result       *servicepb.BoolResponse
	ctx          context.Context
}

func (hct *HasCollectionTask) ID() UniqueID {
	return hct.ReqID
}

func (hct *HasCollectionTask) SetID(uid UniqueID) {
	hct.ReqID = uid
}

func (hct *HasCollectionTask) Type() internalpb.MsgType {
	return hct.MsgType
}

func (hct *HasCollectionTask) BeginTs() Timestamp {
	return hct.Timestamp
}

func (hct *HasCollectionTask) EndTs() Timestamp {
	return hct.Timestamp
}

func (hct *HasCollectionTask) SetTs(ts Timestamp) {
	hct.Timestamp = ts
}

func (hct *HasCollectionTask) PreExecute() error {
	if err := ValidateCollectionName(hct.CollectionName.CollectionName); err != nil {
		return err
	}
	return nil
}

func (hct *HasCollectionTask) Execute() error {
	resp, err := hct.masterClient.HasCollection(hct.ctx, &hct.HasCollectionRequest)
	if err != nil {
		log.Printf("has collection failed, error= %v", err)
		hct.result = &servicepb.BoolResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "internal error",
			},
			Value: false,
		}
	} else {
		hct.result = resp
	}
	return err
}

func (hct *HasCollectionTask) PostExecute() error {
	return nil
}

type DescribeCollectionTask struct {
	Condition
	internalpb.DescribeCollectionRequest
	masterClient masterpb.MasterClient
	result       *servicepb.CollectionDescription
	ctx          context.Context
}

func (dct *DescribeCollectionTask) ID() UniqueID {
	return dct.ReqID
}

func (dct *DescribeCollectionTask) SetID(uid UniqueID) {
	dct.ReqID = uid
}

func (dct *DescribeCollectionTask) Type() internalpb.MsgType {
	return dct.MsgType
}

func (dct *DescribeCollectionTask) BeginTs() Timestamp {
	return dct.Timestamp
}

func (dct *DescribeCollectionTask) EndTs() Timestamp {
	return dct.Timestamp
}

func (dct *DescribeCollectionTask) SetTs(ts Timestamp) {
	dct.Timestamp = ts
}

func (dct *DescribeCollectionTask) PreExecute() error {
	if err := ValidateCollectionName(dct.CollectionName.CollectionName); err != nil {
		return err
	}
	return nil
}

func (dct *DescribeCollectionTask) Execute() error {
	if !globalMetaCache.Hit(dct.CollectionName.CollectionName) {
		err := globalMetaCache.Update(dct.CollectionName.CollectionName)
		if err != nil {
			return err
		}
	}
	var err error
	dct.result, err = globalMetaCache.Get(dct.CollectionName.CollectionName)
	return err
}

func (dct *DescribeCollectionTask) PostExecute() error {
	return nil
}

type ShowCollectionsTask struct {
	Condition
	internalpb.ShowCollectionRequest
	masterClient masterpb.MasterClient
	result       *servicepb.StringListResponse
	ctx          context.Context
}

func (sct *ShowCollectionsTask) ID() UniqueID {
	return sct.ReqID
}

func (sct *ShowCollectionsTask) SetID(uid UniqueID) {
	sct.ReqID = uid
}

func (sct *ShowCollectionsTask) Type() internalpb.MsgType {
	return sct.MsgType
}

func (sct *ShowCollectionsTask) BeginTs() Timestamp {
	return sct.Timestamp
}

func (sct *ShowCollectionsTask) EndTs() Timestamp {
	return sct.Timestamp
}

func (sct *ShowCollectionsTask) SetTs(ts Timestamp) {
	sct.Timestamp = ts
}

func (sct *ShowCollectionsTask) PreExecute() error {
	return nil
}

func (sct *ShowCollectionsTask) Execute() error {
	resp, err := sct.masterClient.ShowCollections(sct.ctx, &sct.ShowCollectionRequest)
	if err != nil {
		log.Printf("show collections failed, error= %v", err)
		sct.result = &servicepb.StringListResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "internal error",
			},
		}
	} else {
		sct.result = resp
	}
	return err
}

func (sct *ShowCollectionsTask) PostExecute() error {
	return nil
}

type CreatePartitionTask struct {
	Condition
	internalpb.CreatePartitionRequest
	masterClient masterpb.MasterClient
	result       *commonpb.Status
	ctx          context.Context
}

func (cpt *CreatePartitionTask) ID() UniqueID {
	return cpt.ReqID
}

func (cpt *CreatePartitionTask) SetID(uid UniqueID) {
	cpt.ReqID = uid
}

func (cpt *CreatePartitionTask) Type() internalpb.MsgType {
	return cpt.MsgType
}

func (cpt *CreatePartitionTask) BeginTs() Timestamp {
	return cpt.Timestamp
}

func (cpt *CreatePartitionTask) EndTs() Timestamp {
	return cpt.Timestamp
}

func (cpt *CreatePartitionTask) SetTs(ts Timestamp) {
	cpt.Timestamp = ts
}

func (cpt *CreatePartitionTask) PreExecute() error {
	collName, partitionTag := cpt.PartitionName.CollectionName, cpt.PartitionName.Tag

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}

	return nil
}

func (cpt *CreatePartitionTask) Execute() (err error) {
	cpt.result, err = cpt.masterClient.CreatePartition(cpt.ctx, &cpt.CreatePartitionRequest)
	return err
}

func (cpt *CreatePartitionTask) PostExecute() error {
	return nil
}

type DropPartitionTask struct {
	Condition
	internalpb.DropPartitionRequest
	masterClient masterpb.MasterClient
	result       *commonpb.Status
	ctx          context.Context
}

func (dpt *DropPartitionTask) ID() UniqueID {
	return dpt.ReqID
}

func (dpt *DropPartitionTask) SetID(uid UniqueID) {
	dpt.ReqID = uid
}

func (dpt *DropPartitionTask) Type() internalpb.MsgType {
	return dpt.MsgType
}

func (dpt *DropPartitionTask) BeginTs() Timestamp {
	return dpt.Timestamp
}

func (dpt *DropPartitionTask) EndTs() Timestamp {
	return dpt.Timestamp
}

func (dpt *DropPartitionTask) SetTs(ts Timestamp) {
	dpt.Timestamp = ts
}

func (dpt *DropPartitionTask) PreExecute() error {
	collName, partitionTag := dpt.PartitionName.CollectionName, dpt.PartitionName.Tag

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}

	return nil
}

func (dpt *DropPartitionTask) Execute() (err error) {
	dpt.result, err = dpt.masterClient.DropPartition(dpt.ctx, &dpt.DropPartitionRequest)
	return err
}

func (dpt *DropPartitionTask) PostExecute() error {
	return nil
}

type HasPartitionTask struct {
	Condition
	internalpb.HasPartitionRequest
	masterClient masterpb.MasterClient
	result       *servicepb.BoolResponse
	ctx          context.Context
}

func (hpt *HasPartitionTask) ID() UniqueID {
	return hpt.ReqID
}

func (hpt *HasPartitionTask) SetID(uid UniqueID) {
	hpt.ReqID = uid
}

func (hpt *HasPartitionTask) Type() internalpb.MsgType {
	return hpt.MsgType
}

func (hpt *HasPartitionTask) BeginTs() Timestamp {
	return hpt.Timestamp
}

func (hpt *HasPartitionTask) EndTs() Timestamp {
	return hpt.Timestamp
}

func (hpt *HasPartitionTask) SetTs(ts Timestamp) {
	hpt.Timestamp = ts
}

func (hpt *HasPartitionTask) PreExecute() error {
	collName, partitionTag := hpt.PartitionName.CollectionName, hpt.PartitionName.Tag

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}
	return nil
}

func (hpt *HasPartitionTask) Execute() (err error) {
	hpt.result, err = hpt.masterClient.HasPartition(hpt.ctx, &hpt.HasPartitionRequest)
	return err
}

func (hpt *HasPartitionTask) PostExecute() error {
	return nil
}

type DescribePartitionTask struct {
	Condition
	internalpb.DescribePartitionRequest
	masterClient masterpb.MasterClient
	result       *servicepb.PartitionDescription
	ctx          context.Context
}

func (dpt *DescribePartitionTask) ID() UniqueID {
	return dpt.ReqID
}

func (dpt *DescribePartitionTask) SetID(uid UniqueID) {
	dpt.ReqID = uid
}

func (dpt *DescribePartitionTask) Type() internalpb.MsgType {
	return dpt.MsgType
}

func (dpt *DescribePartitionTask) BeginTs() Timestamp {
	return dpt.Timestamp
}

func (dpt *DescribePartitionTask) EndTs() Timestamp {
	return dpt.Timestamp
}

func (dpt *DescribePartitionTask) SetTs(ts Timestamp) {
	dpt.Timestamp = ts
}

func (dpt *DescribePartitionTask) PreExecute() error {
	collName, partitionTag := dpt.PartitionName.CollectionName, dpt.PartitionName.Tag

	if err := ValidateCollectionName(collName); err != nil {
		return err
	}

	if err := ValidatePartitionTag(partitionTag, true); err != nil {
		return err
	}
	return nil
}

func (dpt *DescribePartitionTask) Execute() (err error) {
	dpt.result, err = dpt.masterClient.DescribePartition(dpt.ctx, &dpt.DescribePartitionRequest)
	return err
}

func (dpt *DescribePartitionTask) PostExecute() error {
	return nil
}

type ShowPartitionsTask struct {
	Condition
	internalpb.ShowPartitionRequest
	masterClient masterpb.MasterClient
	result       *servicepb.StringListResponse
	ctx          context.Context
}

func (spt *ShowPartitionsTask) ID() UniqueID {
	return spt.ReqID
}

func (spt *ShowPartitionsTask) SetID(uid UniqueID) {
	spt.ReqID = uid
}

func (spt *ShowPartitionsTask) Type() internalpb.MsgType {
	return spt.MsgType
}

func (spt *ShowPartitionsTask) BeginTs() Timestamp {
	return spt.Timestamp
}

func (spt *ShowPartitionsTask) EndTs() Timestamp {
	return spt.Timestamp
}

func (spt *ShowPartitionsTask) SetTs(ts Timestamp) {
	spt.Timestamp = ts
}

func (spt *ShowPartitionsTask) PreExecute() error {
	if err := ValidateCollectionName(spt.CollectionName.CollectionName); err != nil {
		return err
	}
	return nil
}

func (spt *ShowPartitionsTask) Execute() (err error) {
	spt.result, err = spt.masterClient.ShowPartitions(spt.ctx, &spt.ShowPartitionRequest)
	return err
}

func (spt *ShowPartitionsTask) PostExecute() error {
	return nil
}
