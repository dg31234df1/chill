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

package storage

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/cockroachdb/errors"

	"github.com/stretchr/testify/assert"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/proto/etcdpb"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

func initMeta() *etcdpb.CollectionMeta {
	meta := &etcdpb.CollectionMeta{
		ID:            1,
		CreateTime:    1,
		SegmentIDs:    []int64{0, 1},
		PartitionTags: []string{"partition_0", "partition_1"},
		Schema: &schemapb.CollectionSchema{
			Name:        "schema",
			Description: "schema",
			AutoID:      true,
			Fields: []*schemapb.FieldSchema{
				{
					FieldID:      0,
					Name:         "row_id",
					IsPrimaryKey: false,
					Description:  "row_id",
					DataType:     schemapb.DataType_Int64,
				},
				{
					FieldID:      1,
					Name:         "Ts",
					IsPrimaryKey: false,
					Description:  "Ts",
					DataType:     schemapb.DataType_Int64,
				},
				{
					FieldID:      101,
					Name:         "field_int8",
					IsPrimaryKey: false,
					Description:  "description_3",
					DataType:     schemapb.DataType_Int8,
				},
				{
					FieldID:      108,
					Name:         "field_binary_vector",
					IsPrimaryKey: false,
					Description:  "description_10",
					DataType:     schemapb.DataType_BinaryVector,
				},
				{
					FieldID:      109,
					Name:         "field_float_vector",
					IsPrimaryKey: false,
					Description:  "description_11",
					DataType:     schemapb.DataType_FloatVector,
				},
			},
		},
	}
	return meta
}

func initBinlogFile(schema *etcdpb.CollectionMeta) []*Blob {
	insertCodec := NewInsertCodecWithSchema(schema)
	insertData := &InsertData{
		Data: map[int64]FieldData{
			0: &Int64FieldData{
				Data: []int64{3, 4},
			},
			1: &Int64FieldData{
				Data: []int64{3, 4},
			},
			101: &Int8FieldData{
				Data: []int8{3, 4},
			},
			108: &BinaryVectorFieldData{
				Data: []byte{0, 255},
				Dim:  8,
			},
			109: &FloatVectorFieldData{
				Data: []float32{0, 1, 2, 3, 4, 5, 6, 7, 0, 111, 222, 333, 444, 555, 777, 666},
				Dim:  8,
			},
		},
	}

	blobs, err := insertCodec.Serialize(1, 1, insertData)
	if err != nil {
		return nil
	}
	return blobs
}

func buildVectorChunkManager(ctx context.Context, localPath string, localCacheEnable bool) (*VectorChunkManager, error) {
	bucketName := "vector-chunk-manager"

	rcm, err := newMinIOChunkManager(ctx, bucketName, "")
	if err != nil {
		return nil, err
	}
	lcm := NewLocalChunkManager(RootPath(localPath))

	vcm, err := NewVectorChunkManager(ctx, lcm, rcm, 16, localCacheEnable)
	if err != nil {
		return nil, err
	}

	return vcm, nil
}

var Params = paramtable.Get()
var localPath = "/tmp/milvus_test/chunkmanager/"

func TestMain(m *testing.M) {
	paramtable.Init()
	exitCode := m.Run()
	err := os.RemoveAll(localPath)
	if err != nil {
		return
	}
	os.Exit(exitCode)
}

func TestNewVectorChunkManager(t *testing.T) {
	ctx := context.Background()
	bucketName := "vector-chunk-manager"

	rcm, err := newMinIOChunkManager(ctx, bucketName, "")
	assert.NoError(t, err)
	assert.NotNil(t, rcm)
	lcm := NewLocalChunkManager(RootPath(localPath))

	vcm, err := NewVectorChunkManager(ctx, lcm, rcm, 16, true)
	assert.Equal(t, "", vcm.RootPath())
	assert.NoError(t, err)
	assert.NotNil(t, vcm)

	vcm, err = NewVectorChunkManager(ctx, lcm, rcm, -1, true)
	assert.Error(t, err)
	assert.Nil(t, vcm)
}

func TestVectorChunkManager_GetPath(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	localCaches := []bool{true, false}
	for _, localCache := range localCaches {
		vcm, err := buildVectorChunkManager(ctx, localPath, localCache)
		assert.NoError(t, err)
		assert.NotNil(t, vcm)

		key := path.Join(localPath, "1")
		err = vcm.Write(ctx, key, []byte{1})
		assert.NoError(t, err)
		pathGet, err := vcm.Path(ctx, key)
		assert.NoError(t, err)
		assert.Equal(t, pathGet, key)

		err = vcm.cacheStorage.Write(ctx, key, []byte{1})
		assert.NoError(t, err)
		pathGet, err = vcm.Path(ctx, key)
		assert.NoError(t, err)
		assert.Equal(t, pathGet, key)

		err = vcm.RemoveWithPrefix(ctx, localPath)
		assert.NoError(t, err)
		vcm.Close()
	}
}

func TestVectorChunkManager_GetSize(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	localCaches := []bool{true, false}
	for _, localCache := range localCaches {
		vcm, err := buildVectorChunkManager(ctx, localPath, localCache)
		assert.NoError(t, err)
		assert.NotNil(t, vcm)

		key := path.Join(localPath, "1")
		err = vcm.Write(ctx, key, []byte{1})
		assert.NoError(t, err)
		sizeGet, err := vcm.Size(ctx, key)
		assert.NoError(t, err)
		assert.EqualValues(t, sizeGet, 1)

		err = vcm.cacheStorage.Write(ctx, key, []byte{1})
		assert.NoError(t, err)
		sizeGet, err = vcm.Size(ctx, key)
		assert.NoError(t, err)
		assert.EqualValues(t, sizeGet, 1)

		err = vcm.RemoveWithPrefix(ctx, localPath)
		assert.NoError(t, err)
		vcm.Close()
	}
}

func TestVectorChunkManager_Write(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localCaches := []bool{true, false}
	for _, localCache := range localCaches {
		vcm, err := buildVectorChunkManager(ctx, localPath, localCache)
		assert.NoError(t, err)
		assert.NotNil(t, vcm)

		key1 := path.Join(localPath, "key_1")
		key2 := path.Join(localPath, "key_2")
		err = vcm.Write(ctx, key1, []byte{1})
		assert.NoError(t, err)

		exist, err := vcm.Exist(ctx, key1)
		assert.True(t, exist)
		assert.NoError(t, err)

		contents := map[string][]byte{
			key1: {111},
			key2: {222},
		}
		err = vcm.MultiWrite(ctx, contents)
		assert.NoError(t, err)

		exist, err = vcm.Exist(ctx, key1)
		assert.True(t, exist)
		assert.NoError(t, err)
		exist, err = vcm.Exist(ctx, key2)
		assert.True(t, exist)
		assert.NoError(t, err)

		err = vcm.RemoveWithPrefix(ctx, localPath)
		assert.NoError(t, err)
		vcm.Close()
	}
}

func TestVectorChunkManager_Remove(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localCaches := []bool{true, false}
	for _, localCache := range localCaches {
		vcm, err := buildVectorChunkManager(ctx, localPath, localCache)
		assert.NoError(t, err)
		assert.NotNil(t, vcm)

		key1 := path.Join(localPath, "key_1")
		key2 := path.Join(localPath, "key_2")
		err = vcm.cacheStorage.Write(ctx, key1, []byte{1})
		assert.NoError(t, err)

		err = vcm.Remove(ctx, key1)
		assert.NoError(t, err)

		exist, err := vcm.Exist(ctx, key1)
		assert.False(t, exist)
		assert.NoError(t, err)

		contents := map[string][]byte{
			key1: {111},
			key2: {222},
		}
		err = vcm.cacheStorage.MultiWrite(ctx, contents)
		assert.NoError(t, err)

		err = vcm.MultiRemove(ctx, []string{key1, key2})
		assert.NoError(t, err)

		exist, err = vcm.Exist(ctx, key1)
		assert.False(t, exist)
		assert.NoError(t, err)
		exist, err = vcm.Exist(ctx, key2)
		assert.False(t, exist)
		assert.NoError(t, err)

		err = vcm.RemoveWithPrefix(ctx, localPath)
		assert.NoError(t, err)
		vcm.Close()
	}
}

type mockFailedChunkManager struct {
	fail bool
	ChunkManager
}

func (m *mockFailedChunkManager) Remove(ctx context.Context, key string) error {
	if m.fail {
		return errors.New("remove error")
	}
	return nil
}

func (m *mockFailedChunkManager) RemoveWithPrefix(ctx context.Context, prefix string) error {
	if m.fail {
		return errors.New("remove with prefix error")
	}
	return nil
}
func (m *mockFailedChunkManager) MultiRemove(ctx context.Context, key []string) error {
	if m.fail {
		return errors.New("multi remove error")
	}
	return nil
}

func TestVectorChunkManager_Remove_Fail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vcm := &VectorChunkManager{
		vectorStorage: &mockFailedChunkManager{fail: true},
		cacheStorage:  &mockFailedChunkManager{fail: true},
	}
	assert.Error(t, vcm.Remove(ctx, "test"))
	assert.Error(t, vcm.MultiRemove(ctx, []string{"test"}))
	assert.Error(t, vcm.RemoveWithPrefix(ctx, "test"))
}

func TestVectorChunkManager_LocalPath(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	vcm, err := buildVectorChunkManager(ctx, localPath, true)
	assert.NotNil(t, vcm)
	assert.NoError(t, err)

	meta := initMeta()
	binlogs := initBinlogFile(meta)
	keys := make([]string, len(binlogs))
	for i, binlog := range binlogs {
		err = vcm.vectorStorage.Write(ctx, binlog.Key, binlog.Value)
		assert.Nil(t, err)
		keys[i] = binlog.Key
	}

	// cache file to cacheStorage
	_, err = vcm.Read(ctx, keys[0])
	assert.NoError(t, err)

	// check file under cacheStorage.rootPath
	absLocalPath := path.Join(vcm.cacheStorage.RootPath(), keys[0])
	exit, err := vcm.cacheStorage.Exist(ctx, absLocalPath)
	assert.NoError(t, err)
	assert.True(t, exit)

	err = vcm.RemoveWithPrefix(ctx, localPath)
	assert.NoError(t, err)

	vcm.Close()
}

func TestVectorChunkManager_Read(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	localCaches := []bool{true, false}
	for _, localCache := range localCaches {
		vcm, err := buildVectorChunkManager(ctx, localPath, localCache)
		assert.NotNil(t, vcm)
		assert.NoError(t, err)

		content, err := vcm.Read(ctx, "9999")
		assert.Error(t, err)
		assert.Nil(t, content)

		meta := initMeta()
		binlogs := initBinlogFile(meta)
		assert.NotNil(t, binlogs)
		for _, binlog := range binlogs {
			err := vcm.vectorStorage.Write(ctx, binlog.Key, binlog.Value)
			assert.NoError(t, err)
		}

		content, err = vcm.Read(ctx, "108")
		assert.NoError(t, err)
		assert.Equal(t, []byte{0, 255}, content)

		content, err = vcm.Read(ctx, "109")
		assert.NoError(t, err)
		floatResult := make([]float32, 0)
		for i := 0; i < len(content)/4; i++ {
			singleData := typeutil.BytesToFloat32(content[i*4 : i*4+4])
			floatResult = append(floatResult, singleData)
		}
		assert.Equal(t, []float32{0, 1, 2, 3, 4, 5, 6, 7, 0, 111, 222, 333, 444, 555, 777, 666}, floatResult)

		contents, err := vcm.MultiRead(ctx, []string{"108", "109"})
		assert.NoError(t, err)
		assert.Equal(t, []byte{0, 255}, contents[0])

		floatResult = make([]float32, 0)
		for i := 0; i < len(content)/4; i++ {
			singleData := typeutil.BytesToFloat32(contents[1][i*4 : i*4+4])
			floatResult = append(floatResult, singleData)
		}
		assert.Equal(t, []float32{0, 1, 2, 3, 4, 5, 6, 7, 0, 111, 222, 333, 444, 555, 777, 666}, floatResult)

		keys, contents, err := vcm.ReadWithPrefix(ctx, "10")
		assert.NoError(t, err)
		assert.Equal(t, "101", keys[0])
		assert.Equal(t, []byte{3, 4}, contents[0])

		assert.NoError(t, err)
		assert.Equal(t, "108", keys[1])
		assert.Equal(t, []byte{0, 255}, contents[1])

		floatResult = make([]float32, 0)
		for i := 0; i < len(content)/4; i++ {
			singleData := typeutil.BytesToFloat32(contents[2][i*4 : i*4+4])
			floatResult = append(floatResult, singleData)
		}
		assert.Equal(t, "109", keys[2])
		assert.Equal(t, []float32{0, 1, 2, 3, 4, 5, 6, 7, 0, 111, 222, 333, 444, 555, 777, 666}, floatResult)

		content, err = vcm.ReadAt(ctx, "109", 8*4, 8*4)
		assert.NoError(t, err)

		floatResult = make([]float32, 0)
		for i := 0; i < len(content)/4; i++ {
			singleData := typeutil.BytesToFloat32(content[i*4 : i*4+4])
			floatResult = append(floatResult, singleData)
		}
		assert.Equal(t, []float32{0, 111, 222, 333, 444, 555, 777, 666}, floatResult)

		content, err = vcm.ReadAt(ctx, "9999", 0, 8*4)
		assert.Error(t, err)
		assert.Nil(t, content)

		content, err = vcm.ReadAt(ctx, "109", 8*4, 8*4)
		assert.NoError(t, err)
		assert.Equal(t, 32, len(content))

		if localCache {
			r, err := vcm.Mmap(ctx, "109")
			assert.NoError(t, err)
			p := make([]byte, 32)
			n, err := r.ReadAt(p, 32)
			assert.NoError(t, err)
			assert.Equal(t, n, 32)

			r, err = vcm.Mmap(ctx, "not exist")
			assert.Error(t, err)
			assert.Nil(t, r)
		}

		content, err = vcm.ReadAt(ctx, "109", 9999, 8*4)
		assert.Error(t, err)
		assert.Nil(t, content)

		content, err = vcm.ReadAt(ctx, "9999", 0, 8*4)
		assert.Error(t, err)
		assert.Nil(t, content)

		err = vcm.RemoveWithPrefix(ctx, localPath)
		assert.NoError(t, err)

		vcm.Close()
	}
}
