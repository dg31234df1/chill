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

package compaction

import (
	"context"
	sio "io"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/allocator"
	"github.com/milvus-io/milvus/internal/flushcommon/io"
	"github.com/milvus-io/milvus/internal/metastore/kv/binlog"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/tsoutil"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

func isExpiredEntity(ttl int64, now, ts typeutil.Timestamp) bool {
	// entity expire is not enabled if duration <= 0
	if ttl <= 0 {
		return false
	}

	pts, _ := tsoutil.ParseTS(ts)
	pnow, _ := tsoutil.ParseTS(now)
	expireTime := pts.Add(time.Duration(ttl))
	return expireTime.Before(pnow)
}

func mergeDeltalogs(ctx context.Context, io io.BinlogIO, dpaths map[typeutil.UniqueID][]string) (map[interface{}]typeutil.Timestamp, error) {
	pk2ts := make(map[interface{}]typeutil.Timestamp)

	if len(dpaths) == 0 {
		log.Info("compact with no deltalogs, skip merge deltalogs")
		return pk2ts, nil
	}

	blobs := make([]*storage.Blob, 0)
	for segID, paths := range dpaths {
		if len(paths) == 0 {
			continue
		}
		binaries, err := io.Download(ctx, paths)
		if err != nil {
			log.Warn("compact wrong, fail to download deltalogs",
				zap.Int64("segment", segID),
				zap.Strings("path", paths),
				zap.Error(err))
			return nil, err
		}

		for i := range binaries {
			blobs = append(blobs, &storage.Blob{Value: binaries[i]})
		}
	}
	reader, err := storage.CreateDeltalogReader(blobs)
	if err != nil {
		log.Error("malformed delta file", zap.Error(err))
		return nil, err
	}
	defer reader.Close()

	for {
		err := reader.Next()
		if err != nil {
			if err == sio.EOF {
				break
			}
			log.Error("compact wrong, fail to read deltalogs", zap.Error(err))
			return nil, err
		}

		dl := reader.Value()
		// If pk already exists in pk2ts, record the later one.
		if ts, ok := pk2ts[dl.Pk.GetValue()]; ok && ts > dl.Ts {
			continue
		}
		pk2ts[dl.Pk.GetValue()] = dl.Ts
	}

	log.Info("compact mergeDeltalogs end",
		zap.Int("deleted pk counts", len(pk2ts)))

	return pk2ts, nil
}

func composePaths(segments []*datapb.CompactionSegmentBinlogs) (map[typeutil.UniqueID][]string, [][]string, error) {
	if err := binlog.DecompressCompactionBinlogs(segments); err != nil {
		log.Warn("compact wrong, fail to decompress compaction binlogs", zap.Error(err))
		return nil, nil, err
	}

	deltaPaths := make(map[typeutil.UniqueID][]string) // segmentID to deltalog paths
	allPath := make([][]string, 0)                     // group by binlog batch
	for _, s := range segments {
		// Get the batch count of field binlog files from non-empty segment
		// each segment might contain different batches
		var binlogBatchCount int
		for _, b := range s.GetFieldBinlogs() {
			if b != nil {
				binlogBatchCount = len(b.GetBinlogs())
				break
			}
		}
		if binlogBatchCount == 0 {
			log.Warn("compacting empty segment", zap.Int64("segmentID", s.GetSegmentID()))
			continue
		}

		for idx := 0; idx < binlogBatchCount; idx++ {
			var batchPaths []string
			for _, f := range s.GetFieldBinlogs() {
				batchPaths = append(batchPaths, f.GetBinlogs()[idx].GetLogPath())
			}
			allPath = append(allPath, batchPaths)
		}

		deltaPaths[s.GetSegmentID()] = []string{}
		for _, d := range s.GetDeltalogs() {
			for _, l := range d.GetBinlogs() {
				deltaPaths[s.GetSegmentID()] = append(deltaPaths[s.GetSegmentID()], l.GetLogPath())
			}
		}
	}
	return deltaPaths, allPath, nil
}

func serializeWrite(ctx context.Context, allocator allocator.Interface, writer *SegmentWriter) (kvs map[string][]byte, fieldBinlogs map[int64]*datapb.FieldBinlog, err error) {
	_, span := otel.Tracer(typeutil.DataNodeRole).Start(ctx, "serializeWrite")
	defer span.End()

	blobs, tr, err := writer.SerializeYield()
	startID, _, err := allocator.Alloc(uint32(len(blobs)))
	if err != nil {
		return nil, nil, err
	}

	kvs = make(map[string][]byte)
	fieldBinlogs = make(map[int64]*datapb.FieldBinlog)
	for i := range blobs {
		// Blob Key is generated by Serialize from int64 fieldID in collection schema, which won't raise error in ParseInt
		fID, _ := strconv.ParseInt(blobs[i].GetKey(), 10, 64)
		key, _ := binlog.BuildLogPath(storage.InsertBinlog, writer.GetCollectionID(), writer.GetPartitionID(), writer.GetSegmentID(), fID, startID+int64(i))

		kvs[key] = blobs[i].GetValue()
		fieldBinlogs[fID] = &datapb.FieldBinlog{
			FieldID: fID,
			Binlogs: []*datapb.Binlog{
				{
					LogSize:       int64(len(blobs[i].GetValue())),
					MemorySize:    blobs[i].GetMemorySize(),
					LogPath:       key,
					EntriesNum:    blobs[i].RowNum,
					TimestampFrom: tr.GetMinTimestamp(),
					TimestampTo:   tr.GetMaxTimestamp(),
				},
			},
		}
	}

	return
}

func statSerializeWrite(ctx context.Context, io io.BinlogIO, allocator allocator.Interface, writer *SegmentWriter) (*datapb.FieldBinlog, error) {
	ctx, span := otel.Tracer(typeutil.DataNodeRole).Start(ctx, "statslog serializeWrite")
	defer span.End()
	sblob, err := writer.Finish()
	if err != nil {
		return nil, err
	}

	return uploadStatsBlobs(ctx, writer.GetCollectionID(), writer.GetPartitionID(), writer.GetSegmentID(), writer.GetPkID(), writer.GetRowNum(), io, allocator, sblob)
}

func uploadStatsBlobs(ctx context.Context, collectionID, partitionID, segmentID, pkID, numRows int64,
	io io.BinlogIO, allocator allocator.Interface, blob *storage.Blob,
) (*datapb.FieldBinlog, error) {
	logID, err := allocator.AllocOne()
	if err != nil {
		return nil, err
	}

	key, _ := binlog.BuildLogPath(storage.StatsBinlog, collectionID, partitionID, segmentID, pkID, logID)
	kvs := map[string][]byte{key: blob.GetValue()}
	statFieldLog := &datapb.FieldBinlog{
		FieldID: pkID,
		Binlogs: []*datapb.Binlog{
			{
				LogSize:    int64(len(blob.GetValue())),
				MemorySize: int64(len(blob.GetValue())),
				LogPath:    key,
				EntriesNum: numRows,
			},
		},
	}
	if err := io.Upload(ctx, kvs); err != nil {
		log.Warn("failed to upload stats log", zap.Error(err))
		return nil, err
	}

	return statFieldLog, nil
}

func mergeFieldBinlogs(base, paths map[typeutil.UniqueID]*datapb.FieldBinlog) {
	for fID, fpath := range paths {
		if _, ok := base[fID]; !ok {
			base[fID] = &datapb.FieldBinlog{FieldID: fID, Binlogs: make([]*datapb.Binlog, 0)}
		}
		base[fID].Binlogs = append(base[fID].Binlogs, fpath.GetBinlogs()...)
	}
}

func bm25SerializeWrite(ctx context.Context, io io.BinlogIO, allocator allocator.Interface, writer *SegmentWriter) ([]*datapb.FieldBinlog, error) {
	ctx, span := otel.Tracer(typeutil.DataNodeRole).Start(ctx, "bm25 stats log serializeWrite")
	defer span.End()

	stats, err := writer.GetBm25StatsBlob()
	if err != nil {
		return nil, err
	}

	logID, _, err := allocator.Alloc(uint32(len(stats)))
	if err != nil {
		return nil, err
	}

	kvs := make(map[string][]byte)
	binlogs := []*datapb.FieldBinlog{}
	for fieldID, blob := range stats {
		key, _ := binlog.BuildLogPath(storage.BM25Binlog, writer.GetCollectionID(), writer.GetPartitionID(), writer.GetSegmentID(), fieldID, logID)
		kvs[key] = blob.GetValue()
		fieldLog := &datapb.FieldBinlog{
			FieldID: fieldID,
			Binlogs: []*datapb.Binlog{
				{
					LogSize:    int64(len(blob.GetValue())),
					MemorySize: int64(len(blob.GetValue())),
					LogPath:    key,
					EntriesNum: writer.GetRowNum(),
				},
			},
		}

		binlogs = append(binlogs, fieldLog)
	}

	if err := io.Upload(ctx, kvs); err != nil {
		log.Warn("failed to upload bm25 log", zap.Error(err))
		return nil, err
	}

	return binlogs, nil
}
