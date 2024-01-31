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

package importutilv2

import (
	"context"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/internal/util/importutilv2/binlog"
	"github.com/milvus-io/milvus/internal/util/importutilv2/json"
	"github.com/milvus-io/milvus/internal/util/importutilv2/numpy"
	"github.com/milvus-io/milvus/internal/util/importutilv2/parquet"
	"github.com/milvus-io/milvus/pkg/util/merr"
)

//go:generate mockery --name=Reader --structname=MockReader --output=./  --filename=mock_reader.go --with-expecter --inpackage
type Reader interface {
	Read() (*storage.InsertData, error)
	Close()
}

func NewReader(ctx context.Context,
	cm storage.ChunkManager,
	schema *schemapb.CollectionSchema,
	importFile *internalpb.ImportFile,
	options Options,
	bufferSize int,
) (Reader, error) {
	if IsBackup(options) {
		tsStart, tsEnd, err := ParseTimeRange(options)
		if err != nil {
			return nil, err
		}
		paths := importFile.GetPaths()
		return binlog.NewReader(ctx, cm, schema, paths, tsStart, tsEnd)
	}

	fileType, err := GetFileType(importFile)
	if err != nil {
		return nil, err
	}
	switch fileType {
	case JSON:
		reader, err := cm.Reader(ctx, importFile.GetPaths()[0])
		if err != nil {
			return nil, WrapReadFileError(importFile.GetPaths()[0], err)
		}
		return json.NewReader(reader, schema, bufferSize)
	case Numpy:
		return numpy.NewReader(ctx, schema, importFile.GetPaths(), cm, bufferSize)
	case Parquet:
		cmReader, err := cm.Reader(ctx, importFile.GetPaths()[0])
		if err != nil {
			return nil, err
		}
		return parquet.NewReader(ctx, schema, cmReader, bufferSize)
	}
	return nil, merr.WrapErrImportFailed("unexpected import file")
}
