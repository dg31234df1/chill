package typeutil

import (
	"testing"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestNewKvPairs(t *testing.T) {
	kvPairs := []*commonpb.KeyValuePair{
		{Key: common.DimKey, Value: "128"},
	}
	h := NewKvPairs(kvPairs)
	v, err := h.Get(common.DimKey)
	assert.NoError(t, err)
	assert.Equal(t, "128", v)
	_, err = h.Get("not_exist")
	assert.Error(t, err)
}
