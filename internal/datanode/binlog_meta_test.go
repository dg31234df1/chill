package datanode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	memkv "github.com/zilliztech/milvus-distributed/internal/kv/mem"
)

func TestMetaTable_Basic(t *testing.T) {

	kvMock := memkv.NewMemoryKV()
	allocMock := NewAllocatorFactory(22222)
	meta, err := NewBinlogMeta(kvMock, allocMock)
	require.NoError(t, err)
	defer meta.client.Close()

	t.Run("TestBasic_genKey", func(t *testing.T) {
		// 0/1
		alloc := true
		k, err := meta.genKey(alloc, 0)
		assert.NoError(t, err)
		assert.True(t, strings.HasPrefix(k, "0/"))

		// rand int64
		_, err = meta.genKey(alloc)
		assert.NoError(t, err)

		// 1/2/3/1
		k, err = meta.genKey(alloc, 1, 2, 3)
		assert.NoError(t, err)
		assert.True(t, strings.HasPrefix(k, "1/2/3/"))

		// 0
		alloc = false
		k, err = meta.genKey(alloc, 0)
		assert.NoError(t, err)
		assert.Equal(t, "0", k)

		// ""
		k, err = meta.genKey(alloc)
		assert.NoError(t, err)
		assert.Equal(t, "", k)

		// 1/2/3
		k, err = meta.genKey(alloc, 1, 2, 3)
		assert.NoError(t, err)
		assert.Equal(t, "1/2/3", k)
	})

	t.Run("TestBasic_SaveSegmentBinlogMetaTxn", func(t *testing.T) {
		segID := UniqueID(999999)
		fieldID2Path := map[UniqueID]string{
			100: "a",
			200: "b",
			300: "c",
		}

		err := meta.SaveSegmentBinlogMetaTxn(segID, fieldID2Path)
		assert.NoError(t, err)

		metas, err := meta.getFieldBinlogMeta(segID, 100)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(metas))
		assert.Equal(t, "a", metas[0].GetBinlogPath())

		metas, err = meta.getFieldBinlogMeta(segID, 200)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(metas))
		assert.Equal(t, "b", metas[0].GetBinlogPath())

		metas, err = meta.getFieldBinlogMeta(segID, 300)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(metas))
		assert.Equal(t, "c", metas[0].GetBinlogPath())

		fieldID2Path2 := map[UniqueID]string{
			100: "aa",
			200: "bb",
			300: "cc",
		}

		err = meta.SaveSegmentBinlogMetaTxn(segID, fieldID2Path2)
		assert.NoError(t, err)

		metas, err = meta.getSegmentBinlogMeta(segID)
		assert.NoError(t, err)
		assert.Equal(t, 6, len(metas))

		paths := make([]string, 0, 6)
		for _, meta := range metas {
			paths = append(paths, meta.GetBinlogPath())
		}

		assert.ElementsMatch(t, []string{"a", "b", "c", "aa", "bb", "cc"}, paths)
	})

	t.Run("TestBasic_SaveDDLBinlogMetaTxn", func(t *testing.T) {
		collID := UniqueID(888888)
		tsPath := "a/b/c"
		ddlPath := "c/b/a"

		err := meta.SaveDDLBinlogMetaTxn(collID, tsPath, ddlPath)
		assert.NoError(t, err)

		metas, err := meta.getDDLBinlogMete(collID)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(metas))
		assert.Equal(t, "a/b/c", metas[0].GetTsBinlogPath())
		assert.Equal(t, "c/b/a", metas[0].GetDdlBinlogPath())

		err = meta.SaveDDLBinlogMetaTxn(collID, tsPath, ddlPath)
		assert.NoError(t, err)

		metas, err = meta.getDDLBinlogMete(collID)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(metas))
		assert.Equal(t, "a/b/c", metas[0].GetTsBinlogPath())
		assert.Equal(t, "c/b/a", metas[0].GetDdlBinlogPath())
	})

}
