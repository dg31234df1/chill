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

package querynode

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strconv"

	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/common"
	"github.com/milvus-io/milvus/internal/kv"
	etcdkv "github.com/milvus-io/milvus/internal/kv/etcd"
	minioKV "github.com/milvus-io/milvus/internal/kv/minio"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/querypb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/internal/types"
	"github.com/milvus-io/milvus/internal/util/funcutil"
)

const (
	queryCoordSegmentMetaPrefix = "queryCoord-segmentMeta"
	queryNodeSegmentMetaPrefix  = "queryNode-segmentMeta"
)

// segmentLoader is only responsible for loading the field data from binlog
type segmentLoader struct {
	historicalReplica ReplicaInterface

	dataCoord types.DataCoord

	minioKV kv.BaseKV // minio minioKV
	etcdKV  *etcdkv.EtcdKV

	indexLoader *indexLoader
}

func (loader *segmentLoader) loadSegmentOfConditionHandOff(req *querypb.LoadSegmentsRequest) error {
	return errors.New("TODO: implement hand off")
}

func (loader *segmentLoader) loadSegmentOfConditionLoadBalance(req *querypb.LoadSegmentsRequest) error {
	return loader.loadSegment(req, false)
}

func (loader *segmentLoader) loadSegmentOfConditionGRPC(req *querypb.LoadSegmentsRequest) error {
	return loader.loadSegment(req, true)
}

func (loader *segmentLoader) loadSegmentOfConditionNodeDown(req *querypb.LoadSegmentsRequest) error {
	return loader.loadSegment(req, true)
}

func (loader *segmentLoader) loadSegment(req *querypb.LoadSegmentsRequest, onService bool) error {
	// no segment needs to load, return
	if len(req.Infos) == 0 {
		return nil
	}

	newSegments := make([]*Segment, 0)
	segmentGC := func() {
		for _, s := range newSegments {
			deleteSegment(s)
		}
	}
	setSegments := func() error {
		for _, s := range newSegments {
			err := loader.historicalReplica.setSegment(s)
			if err != nil {
				segmentGC()
				return err
			}
		}
		return nil
	}

	// start to load
	for _, info := range req.Infos {
		segmentID := info.SegmentID
		partitionID := info.PartitionID
		collectionID := info.CollectionID

		collection, err := loader.historicalReplica.getCollectionByID(collectionID)
		if err != nil {
			log.Warn(err.Error())
			segmentGC()
			return err
		}
		segment := newSegment(collection, segmentID, partitionID, collectionID, "", segmentTypeSealed, onService)
		err = loader.loadSegmentInternal(collectionID, segment, info)
		if err != nil {
			deleteSegment(segment)
			log.Warn(err.Error())
			segmentGC()
			return err
		}
		if onService {
			key := fmt.Sprintf("%s/%d", queryCoordSegmentMetaPrefix, segmentID)
			value, err := loader.etcdKV.Load(key)
			if err != nil {
				deleteSegment(segment)
				log.Warn("error when load segment info from etcd", zap.Any("error", err.Error()))
				segmentGC()
				return err
			}
			segmentInfo := &querypb.SegmentInfo{}
			err = proto.Unmarshal([]byte(value), segmentInfo)
			if err != nil {
				deleteSegment(segment)
				log.Warn("error when unmarshal segment info from etcd", zap.Any("error", err.Error()))
				segmentGC()
				return err
			}
			segmentInfo.SegmentState = querypb.SegmentState_sealed
			newKey := fmt.Sprintf("%s/%d", queryNodeSegmentMetaPrefix, segmentID)
			newValue, err := proto.Marshal(segmentInfo)
			if err != nil {
				deleteSegment(segment)
				log.Warn("error when marshal segment info", zap.Error(err))
				segmentGC()
				return err
			}
			err = loader.etcdKV.Save(newKey, string(newValue))
			if err != nil {
				deleteSegment(segment)
				log.Warn("error when update segment info to etcd", zap.Any("error", err.Error()))
				segmentGC()
				return err
			}
		}
		newSegments = append(newSegments, segment)
	}

	return setSegments()
}

func (loader *segmentLoader) loadSegmentInternal(collectionID UniqueID, segment *Segment, segmentLoadInfo *querypb.SegmentLoadInfo) error {
	vectorFieldIDs, err := loader.historicalReplica.getVecFieldIDsByCollectionID(collectionID)
	if err != nil {
		return err
	}
	if len(vectorFieldIDs) <= 0 {
		return fmt.Errorf("no vector field in collection %d", collectionID)
	}

	// add VectorFieldInfo for vector fields
	for _, fieldBinlog := range segmentLoadInfo.BinlogPaths {
		if funcutil.SliceContain(vectorFieldIDs, fieldBinlog.FieldID) {
			vectorFieldInfo := newVectorFieldInfo(fieldBinlog)
			segment.setVectorFieldInfo(fieldBinlog.FieldID, vectorFieldInfo)
		}
	}

	indexedFieldIDs := make([]FieldID, 0)
	for _, vecFieldID := range vectorFieldIDs {
		err = loader.indexLoader.setIndexInfo(collectionID, segment, vecFieldID)
		if err != nil {
			log.Warn(err.Error())
			continue
		}
		indexedFieldIDs = append(indexedFieldIDs, vecFieldID)
	}

	// we don't need to load raw data for indexed vector field
	fieldBinlogs := loader.filterFieldBinlogs(segmentLoadInfo.BinlogPaths, indexedFieldIDs)

	log.Debug("loading insert...")
	err = loader.loadSegmentFieldsData(segment, fieldBinlogs)
	if err != nil {
		return err
	}

	pkIDField, err := loader.historicalReplica.getPKFieldIDByCollectionID(collectionID)
	if err != nil {
		return err
	}
	if pkIDField == common.InvalidFieldID {
		log.Warn("segment primary key field doesn't exist when load segment")
	} else {
		log.Debug("loading bloom filter...")
		pkStatsBinlogs := loader.filterPKStatsBinlogs(segmentLoadInfo.Statslogs, pkIDField)
		err = loader.loadSegmentBloomFilter(segment, pkStatsBinlogs)
		if err != nil {
			return err
		}
	}

	log.Debug("loading delta...")
	err = loader.loadDeltaLogs(segment, segmentLoadInfo.Deltalogs)
	if err != nil {
		return err
	}

	for _, id := range indexedFieldIDs {
		log.Debug("loading index...")
		err = loader.indexLoader.loadIndex(segment, id)
		if err != nil {
			return err
		}
	}

	return nil
}

//func (loader *segmentLoader) GetSegmentStates(segmentID UniqueID) (*datapb.GetSegmentStatesResponse, error) {
//	ctx := context.TODO()
//	if loader.dataCoord == nil {
//		return nil, errors.New("null data service client")
//	}
//
//	segmentStatesRequest := &datapb.GetSegmentStatesRequest{
//		SegmentIDs: []int64{segmentID},
//	}
//	statesResponse, err := loader.dataCoord.GetSegmentStates(ctx, segmentStatesRequest)
//	if err != nil || statesResponse.Status.ErrorCode != commonpb.ErrorCode_Success {
//		return nil, err
//	}
//	if len(statesResponse.States) != 1 {
//		return nil, errors.New("segment states' len should be 1")
//	}
//
//	return statesResponse, nil
//}

func (loader *segmentLoader) filterPKStatsBinlogs(fieldBinlogs []*datapb.FieldBinlog, pkFieldID int64) []string {
	result := make([]string, 0)
	for _, fieldBinlog := range fieldBinlogs {
		if fieldBinlog.FieldID == pkFieldID {
			result = append(result, fieldBinlog.Binlogs...)
		}
	}
	return result
}

func (loader *segmentLoader) filterFieldBinlogs(fieldBinlogs []*datapb.FieldBinlog, skipFieldIDs []int64) []*datapb.FieldBinlog {
	result := make([]*datapb.FieldBinlog, 0)
	for _, fieldBinlog := range fieldBinlogs {
		if !funcutil.SliceContain(skipFieldIDs, fieldBinlog.FieldID) {
			result = append(result, fieldBinlog)
		}
	}
	return result
}

func (loader *segmentLoader) loadSegmentFieldsData(segment *Segment, fieldBinlogs []*datapb.FieldBinlog) error {
	iCodec := storage.InsertCodec{}
	defer func() {
		err := iCodec.Close()
		if err != nil {
			log.Warn(err.Error())
		}
	}()
	blobs := make([]*storage.Blob, 0)
	for _, fb := range fieldBinlogs {
		log.Debug("load segment fields data",
			zap.Int64("segmentID", segment.segmentID),
			zap.Any("fieldID", fb.FieldID),
			zap.String("paths", fmt.Sprintln(fb.Binlogs)),
		)
		for _, path := range fb.Binlogs {
			p := path
			binLog, err := loader.minioKV.Load(path)
			if err != nil {
				// TODO: return or continue?
				return err
			}
			blob := &storage.Blob{
				Key:   p,
				Value: []byte(binLog),
			}
			blobs = append(blobs, blob)
		}
	}

	_, _, insertData, err := iCodec.Deserialize(blobs)
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	for fieldID, value := range insertData.Data {
		var numRows []int64
		var data interface{}
		switch fieldData := value.(type) {
		case *storage.BoolFieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.Int8FieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.Int16FieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.Int32FieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.Int64FieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.FloatFieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.DoubleFieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.StringFieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.FloatVectorFieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		case *storage.BinaryVectorFieldData:
			numRows = fieldData.NumRows
			data = fieldData.Data
		default:
			return errors.New("unexpected field data type")
		}
		if fieldID == common.TimeStampField {
			segment.setIDBinlogRowSizes(numRows)
		}
		totalNumRows := int64(0)
		for _, numRow := range numRows {
			totalNumRows += numRow
		}
		err = segment.segmentLoadFieldData(fieldID, int(totalNumRows), data)
		if err != nil {
			// TODO: return or continue?
			return err
		}
	}

	return nil
}

func (loader *segmentLoader) loadSegmentBloomFilter(segment *Segment, binlogPaths []string) error {
	if len(binlogPaths) == 0 {
		log.Info("there are no stats logs saved with segment", zap.Any("segmentID", segment.segmentID))
		return nil
	}

	values, err := loader.minioKV.MultiLoad(binlogPaths)
	if err != nil {
		return err
	}
	blobs := make([]*storage.Blob, 0)
	for i := 0; i < len(values); i++ {
		blobs = append(blobs, &storage.Blob{Value: []byte(values[i])})
	}

	stats, err := storage.DeserializeStats(blobs)
	if err != nil {
		return err
	}
	for _, stat := range stats {
		if stat.BF == nil {
			log.Warn("stat log with nil bloom filter", zap.Int64("segmentID", segment.segmentID), zap.Any("stat", stat))
			continue
		}
		err = segment.pkFilter.Merge(stat.BF)
		if err != nil {
			return err
		}
	}
	return nil
}

func (loader *segmentLoader) loadDeltaLogs(segment *Segment, deltaLogs []*datapb.DeltaLogInfo) error {
	if len(deltaLogs) == 0 {
		log.Info("there are no delta logs saved with segment", zap.Any("segmentID", segment.segmentID))
		return nil
	}
	dCodec := storage.DeleteCodec{}
	blobs := make([]*storage.Blob, 0)
	for _, deltaLog := range deltaLogs {
		value, err := loader.minioKV.Load(deltaLog.DeltaLogPath)
		if err != nil {
			return err
		}
		blob := &storage.Blob{
			Key:   deltaLog.DeltaLogPath,
			Value: []byte(value),
		}
		blobs = append(blobs, blob)
	}
	_, _, deltaData, err := dCodec.Deserialize(blobs)
	if err != nil {
		return err
	}

	//rowCount := len(deltaData.Data)
	pks := make([]int64, 0)
	tss := make([]int64, 0)
	for pk, ts := range deltaData.Data {
		pks = append(pks, pk)
		tss = append(tss, ts)
	}
	//segment.Delete(pks, tss, rowCount)
	return nil
}

// JoinIDPath joins ids to path format.
func JoinIDPath(ids ...UniqueID) string {
	idStr := make([]string, len(ids))
	for _, id := range ids {
		idStr = append(idStr, strconv.FormatInt(id, 10))
	}
	return path.Join(idStr...)
}

func newSegmentLoader(ctx context.Context, rootCoord types.RootCoord, indexCoord types.IndexCoord, replica ReplicaInterface, etcdKV *etcdkv.EtcdKV) *segmentLoader {
	option := &minioKV.Option{
		Address:           Params.MinioEndPoint,
		AccessKeyID:       Params.MinioAccessKeyID,
		SecretAccessKeyID: Params.MinioSecretAccessKey,
		UseSSL:            Params.MinioUseSSLStr,
		CreateBucket:      true,
		BucketName:        Params.MinioBucketName,
	}

	client, err := minioKV.NewMinIOKV(ctx, option)
	if err != nil {
		panic(err)
	}

	iLoader := newIndexLoader(ctx, rootCoord, indexCoord, replica)
	return &segmentLoader{
		historicalReplica: replica,

		minioKV: client,
		etcdKV:  etcdKV,

		indexLoader: iLoader,
	}
}
