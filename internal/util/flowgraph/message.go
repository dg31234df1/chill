package flowgraph

import "github.com/zilliztech/milvus-distributed/internal/msgstream"

type Msg interface {
	TimeTick() Timestamp
}

type MsgStreamMsg struct {
	tsMessages     []msgstream.TsMsg
	timestampMin   Timestamp
	timestampMax   Timestamp
	startPositions []*MsgPosition
}

func GenerateMsgStreamMsg(tsMessages []msgstream.TsMsg, timestampMin, timestampMax Timestamp, positions []*MsgPosition) *MsgStreamMsg {
	return &MsgStreamMsg{
		tsMessages:     tsMessages,
		timestampMin:   timestampMin,
		timestampMax:   timestampMax,
		startPositions: positions,
	}
}

func (msMsg *MsgStreamMsg) TimeTick() Timestamp {
	return msMsg.timestampMax
}

func (msMsg *MsgStreamMsg) DownStreamNodeIdx() int {
	return 0
}

func (msMsg *MsgStreamMsg) TsMessages() []msgstream.TsMsg {
	return msMsg.tsMessages
}

func (msMsg *MsgStreamMsg) TimestampMin() Timestamp {
	return msMsg.timestampMin
}

func (msMsg *MsgStreamMsg) TimestampMax() Timestamp {
	return msMsg.timestampMax
}

func (msMsg *MsgStreamMsg) StartPositions() []*MsgPosition {
	return msMsg.startPositions
}
