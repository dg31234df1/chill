package msgstream

import (
	"context"

	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
)

type UniqueID = typeutil.UniqueID
type Timestamp = typeutil.Timestamp
type IntPrimaryKey = typeutil.IntPrimaryKey
type MsgPosition = internalpb.MsgPosition

type MsgPack struct {
	BeginTs        Timestamp
	EndTs          Timestamp
	Msgs           []TsMsg
	StartPositions []*MsgPosition
	EndPositions   []*MsgPosition
}

type RepackFunc func(msgs []TsMsg, hashKeys [][]int32) (map[int32]*MsgPack, error)

type MsgStream interface {
	Start()
	Close()
	Chan() <-chan *MsgPack
	AsProducer(channels []string)
	AsConsumer(channels []string, subName string)
	SetRepackFunc(repackFunc RepackFunc)

	Produce(context.Context, *MsgPack) error
	Broadcast(context.Context, *MsgPack) error
	Consume() (*MsgPack, context.Context)
	Seek(offset *MsgPosition) error
}

type Factory interface {
	SetParams(params map[string]interface{}) error
	NewMsgStream(ctx context.Context) (MsgStream, error)
	NewTtMsgStream(ctx context.Context) (MsgStream, error)
	NewQueryMsgStream(ctx context.Context) (MsgStream, error)
}
