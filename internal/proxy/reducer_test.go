package proxy

import (
	"context"
	"testing"

	"github.com/milvus-io/milvus/internal/proto/planpb"
	"github.com/stretchr/testify/assert"
)

func Test_createMilvusReducer(t *testing.T) {
	n := &planpb.PlanNode{
		Node: &planpb.PlanNode_Query{
			Query: &planpb.QueryPlanNode{
				IsCount: false,
			},
		},
	}
	var r milvusReducer
	ctx := context.Background()

	r = createMilvusReducer(ctx, nil, nil, nil, n, "")
	_, ok := r.(*defaultLimitReducer)
	assert.True(t, ok)

	n.Node.(*planpb.PlanNode_Query).Query.IsCount = true
	r = createMilvusReducer(ctx, nil, nil, nil, n, "")
	_, ok = r.(*cntReducer)
	assert.True(t, ok)
}
