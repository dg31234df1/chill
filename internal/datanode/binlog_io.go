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
	"strconv"

	"github.com/cockroachdb/errors"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/datanode/allocator"
	"github.com/milvus-io/milvus/internal/datanode/io"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/metautil"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

var (
	errUploadToBlobStorage     = errors.New("upload to blob storage wrong")
	errDownloadFromBlobStorage = errors.New("download from blob storage wrong")
	// errStart used for retry start
	errStart = errors.New("start")
)

func downloadBlobs(ctx context.Context, b io.BinlogIO, paths []string) ([]*Blob, error) {
	ctx, span := otel.Tracer(typeutil.DataNodeRole).Start(ctx, "downloadBlobs")
	defer span.End()
	log.Debug("down load", zap.Strings("path", paths))
	bytes, err := b.Download(ctx, paths)
	if err != nil {
		log.Warn("ctx done when downloading kvs from blob storage", zap.Strings("paths", paths))
		return nil, errDownloadFromBlobStorage
	}
	resp := make([]*Blob, len(paths))
	if len(paths) == 0 {
		return resp, nil
	}
	for i := range bytes {
		resp[i] = &Blob{Key: paths[i], Value: bytes[i]}
	}
	return resp, nil
}

// genDeltaBlobs returns key, value
func genDeltaBlobs(b io.BinlogIO, allocator allocator.Allocator, data *DeleteData, collID, partID, segID UniqueID) (string, []byte, error) {
	dCodec := storage.NewDeleteCodec()

	blob, err := dCodec.Serialize(collID, partID, segID, data)
	if err != nil {
		return "", nil, err
	}

	idx, err := allocator.AllocOne()
	if err != nil {
		return "", nil, err
	}
	k := metautil.JoinIDPath(collID, partID, segID, idx)
	key := b.JoinFullPath(common.SegmentDeltaLogPath, k)

	return key, blob.GetValue(), nil
}

// genInsertBlobs returns insert-paths and save blob to kvs
func genInsertBlobs(b io.BinlogIO, allocator allocator.Allocator, data *InsertData, collectionID, partID, segID UniqueID, iCodec *storage.InsertCodec, kvs map[string][]byte) (map[UniqueID]*datapb.FieldBinlog, error) {
	inlogs, err := iCodec.Serialize(partID, segID, data)
	if err != nil {
		return nil, err
	}

	inpaths := make(map[UniqueID]*datapb.FieldBinlog)
	notifyGenIdx := make(chan struct{})
	defer close(notifyGenIdx)

	generator, err := allocator.GetGenerator(len(inlogs), notifyGenIdx)
	if err != nil {
		return nil, err
	}

	for _, blob := range inlogs {
		// Blob Key is generated by Serialize from int64 fieldID in collection schema, which won't raise error in ParseInt
		fID, _ := strconv.ParseInt(blob.GetKey(), 10, 64)
		k := metautil.JoinIDPath(collectionID, partID, segID, fID, <-generator)
		key := b.JoinFullPath(common.SegmentInsertLogPath, k)
		value := blob.GetValue()
		fileLen := len(value)

		kvs[key] = value
		inpaths[fID] = &datapb.FieldBinlog{
			FieldID: fID,
			Binlogs: []*datapb.Binlog{{LogSize: int64(fileLen), LogPath: key, EntriesNum: blob.RowNum}},
		}
	}

	return inpaths, nil
}

// genStatBlobs return stats log paths and save blob to kvs
func genStatBlobs(b io.BinlogIO, allocator allocator.Allocator, stats *storage.PrimaryKeyStats, collectionID, partID, segID UniqueID, iCodec *storage.InsertCodec, kvs map[string][]byte, totRows int64) (map[UniqueID]*datapb.FieldBinlog, error) {
	statBlob, err := iCodec.SerializePkStats(stats, totRows)
	if err != nil {
		return nil, err
	}
	statPaths := make(map[UniqueID]*datapb.FieldBinlog)

	idx, err := allocator.AllocOne()
	if err != nil {
		return nil, err
	}
	fID, _ := strconv.ParseInt(statBlob.GetKey(), 10, 64)
	k := metautil.JoinIDPath(collectionID, partID, segID, fID, idx)
	key := b.JoinFullPath(common.SegmentStatslogPath, k)
	value := statBlob.GetValue()
	fileLen := len(value)

	kvs[key] = value

	statPaths[fID] = &datapb.FieldBinlog{
		FieldID: fID,
		Binlogs: []*datapb.Binlog{{LogSize: int64(fileLen), LogPath: key, EntriesNum: totRows}},
	}
	return statPaths, nil
}

// update stats log
// also update with insert data if not nil
func uploadStatsLog(
	ctx context.Context,
	b io.BinlogIO,
	allocator allocator.Allocator,
	collectionID UniqueID,
	partID UniqueID,
	segID UniqueID,
	stats *storage.PrimaryKeyStats,
	totRows int64,
	iCodec *storage.InsertCodec,
) (map[UniqueID]*datapb.FieldBinlog, error) {
	ctx, span := otel.Tracer(typeutil.DataNodeRole).Start(ctx, "UploadStatslog")
	defer span.End()
	kvs := make(map[string][]byte)

	statPaths, err := genStatBlobs(b, allocator, stats, collectionID, partID, segID, iCodec, kvs, totRows)
	if err != nil {
		return nil, err
	}

	err = b.Upload(ctx, kvs)
	if err != nil {
		return nil, err
	}

	return statPaths, nil
}

func uploadInsertLog(
	ctx context.Context,
	b io.BinlogIO,
	allocator allocator.Allocator,
	collectionID UniqueID,
	partID UniqueID,
	segID UniqueID,
	iData *InsertData,
	iCodec *storage.InsertCodec,
) (map[UniqueID]*datapb.FieldBinlog, error) {
	ctx, span := otel.Tracer(typeutil.DataNodeRole).Start(ctx, "UploadInsertLog")
	defer span.End()
	kvs := make(map[string][]byte)

	if iData.IsEmpty() {
		log.Warn("binlog io uploading empty insert data",
			zap.Int64("segmentID", segID),
			zap.Int64("collectionID", iCodec.Schema.GetID()),
		)
		return nil, nil
	}

	inpaths, err := genInsertBlobs(b, allocator, iData, collectionID, partID, segID, iCodec, kvs)
	if err != nil {
		return nil, err
	}

	err = b.Upload(ctx, kvs)
	if err != nil {
		return nil, err
	}

	return inpaths, nil
}

func uploadDeltaLog(
	ctx context.Context,
	b io.BinlogIO,
	allocator allocator.Allocator,
	collectionID UniqueID,
	partID UniqueID,
	segID UniqueID,
	dData *DeleteData,
) ([]*datapb.FieldBinlog, error) {
	ctx, span := otel.Tracer(typeutil.DataNodeRole).Start(ctx, "UploadDeltaLog")
	defer span.End()
	var (
		deltaInfo = make([]*datapb.FieldBinlog, 0)
		kvs       = make(map[string][]byte)
	)

	if dData.RowCount > 0 {
		k, v, err := genDeltaBlobs(b, allocator, dData, collectionID, partID, segID)
		if err != nil {
			log.Warn("generate delta blobs wrong",
				zap.Int64("collectionID", collectionID),
				zap.Int64("segmentID", segID),
				zap.Error(err))
			return nil, err
		}

		kvs[k] = v
		deltaInfo = append(deltaInfo, &datapb.FieldBinlog{
			FieldID: 0, // TODO: Not useful on deltalogs, FieldID shall be ID of primary key field
			Binlogs: []*datapb.Binlog{{
				EntriesNum: dData.RowCount,
				LogPath:    k,
				LogSize:    int64(len(v)),
			}},
		})
	} else {
		return nil, nil
	}

	err := b.Upload(ctx, kvs)
	if err != nil {
		return nil, err
	}

	return deltaInfo, nil
}
