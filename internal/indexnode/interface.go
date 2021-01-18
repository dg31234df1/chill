package indexnode

import (
	"github.com/zilliztech/milvus-distributed/internal/proto/indexpb"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
)

type ServiceBase = typeutil.Component

type Interface interface {
	ServiceBase
	BuildIndex(req indexpb.BuildIndexRequest) (indexpb.BuildIndexResponse, error)
}
