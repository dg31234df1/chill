package indexnode

import (
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/indexpb"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
)

type ServiceBase = typeutil.Component

type Interface interface {
	BuildIndex(req *indexpb.BuildIndexCmd) (*commonpb.Status, error)
}
