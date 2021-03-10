package msgstream

import (
	"errors"

	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
)

type UnmarshalFunc func(interface{}) (TsMsg, error)

type UnmarshalDispatcher interface {
	Unmarshal(input interface{}, msgType commonpb.MsgType) (TsMsg, error)
	AddMsgTemplate(msgType commonpb.MsgType, unmarshalFunc UnmarshalFunc)
}

type UnmarshalDispatcherFactory interface {
	NewUnmarshalDispatcher() *UnmarshalDispatcher
}

// ProtoUnmarshalDispatcher ant its factory

type ProtoUnmarshalDispatcher struct {
	TempMap map[commonpb.MsgType]UnmarshalFunc
}

func (p *ProtoUnmarshalDispatcher) Unmarshal(input interface{}, msgType commonpb.MsgType) (TsMsg, error) {
	unmarshalFunc, ok := p.TempMap[msgType]
	if !ok {
		return nil, errors.New("not set unmarshalFunc for this messageType")
	}
	return unmarshalFunc(input)
}

func (p *ProtoUnmarshalDispatcher) AddMsgTemplate(msgType commonpb.MsgType, unmarshalFunc UnmarshalFunc) {
	p.TempMap[msgType] = unmarshalFunc
}

type ProtoUDFactory struct{}

func (pudf *ProtoUDFactory) NewUnmarshalDispatcher() *ProtoUnmarshalDispatcher {
	insertMsg := InsertMsg{}
	deleteMsg := DeleteMsg{}
	searchMsg := SearchMsg{}
	searchResultMsg := SearchResultMsg{}
	timeTickMsg := TimeTickMsg{}
	createCollectionMsg := CreateCollectionMsg{}
	dropCollectionMsg := DropCollectionMsg{}
	createPartitionMsg := CreatePartitionMsg{}
	dropPartitionMsg := DropPartitionMsg{}
	loadIndexMsg := LoadIndexMsg{}
	flushMsg := FlushMsg{}
	segmentInfoMsg := SegmentInfoMsg{}
	flushCompletedMsg := FlushCompletedMsg{}
	queryNodeSegStatsMsg := QueryNodeStatsMsg{}
	segmentStatisticsMsg := SegmentStatisticsMsg{}

	p := &ProtoUnmarshalDispatcher{}
	p.TempMap = make(map[commonpb.MsgType]UnmarshalFunc)
	p.TempMap[commonpb.MsgType_Insert] = insertMsg.Unmarshal
	p.TempMap[commonpb.MsgType_Delete] = deleteMsg.Unmarshal
	p.TempMap[commonpb.MsgType_Search] = searchMsg.Unmarshal
	p.TempMap[commonpb.MsgType_SearchResult] = searchResultMsg.Unmarshal
	p.TempMap[commonpb.MsgType_TimeTick] = timeTickMsg.Unmarshal
	p.TempMap[commonpb.MsgType_QueryNodeStats] = queryNodeSegStatsMsg.Unmarshal
	p.TempMap[commonpb.MsgType_CreateCollection] = createCollectionMsg.Unmarshal
	p.TempMap[commonpb.MsgType_DropCollection] = dropCollectionMsg.Unmarshal
	p.TempMap[commonpb.MsgType_CreatePartition] = createPartitionMsg.Unmarshal
	p.TempMap[commonpb.MsgType_DropPartition] = dropPartitionMsg.Unmarshal
	p.TempMap[commonpb.MsgType_LoadIndex] = loadIndexMsg.Unmarshal
	p.TempMap[commonpb.MsgType_Flush] = flushMsg.Unmarshal
	p.TempMap[commonpb.MsgType_SegmentInfo] = segmentInfoMsg.Unmarshal
	p.TempMap[commonpb.MsgType_SegmentFlushDone] = flushCompletedMsg.Unmarshal
	p.TempMap[commonpb.MsgType_SegmentStatistics] = segmentStatisticsMsg.Unmarshal

	return p
}

// MemUnmarshalDispatcher ant its factory

//type MemUDFactory struct {
//
//}
//func (mudf *MemUDFactory) NewUnmarshalDispatcher() *UnmarshalDispatcher {
//
//}
