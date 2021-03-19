package memms

import (
	"context"
	"errors"
	"log"
	"strconv"
	"sync"

	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/msgstream/util"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
)

type TsMsg = msgstream.TsMsg
type MsgPack = msgstream.MsgPack
type MsgType = msgstream.MsgType
type UniqueID = msgstream.UniqueID
type BaseMsg = msgstream.BaseMsg
type Timestamp = msgstream.Timestamp
type IntPrimaryKey = msgstream.IntPrimaryKey
type TimeTickMsg = msgstream.TimeTickMsg
type QueryNodeStatsMsg = msgstream.QueryNodeStatsMsg
type RepackFunc = msgstream.RepackFunc

type MemMsgStream struct {
	ctx          context.Context
	streamCancel func()

	repackFunc msgstream.RepackFunc

	consumers []*Consumer
	producers []string

	receiveBuf chan *msgstream.MsgPack

	wait sync.WaitGroup
}

func NewMemMsgStream(ctx context.Context, receiveBufSize int64) (*MemMsgStream, error) {
	streamCtx, streamCancel := context.WithCancel(ctx)
	receiveBuf := make(chan *msgstream.MsgPack, receiveBufSize)
	channels := make([]string, 0)
	consumers := make([]*Consumer, 0)

	stream := &MemMsgStream{
		ctx:          streamCtx,
		streamCancel: streamCancel,
		receiveBuf:   receiveBuf,
		consumers:    consumers,
		producers:    channels,
	}

	return stream, nil
}

func (mms *MemMsgStream) Start() {
}

func (mms *MemMsgStream) Close() {
	for _, consumer := range mms.consumers {
		Mmq.DestroyConsumerGroup(consumer.GroupName, consumer.ChannelName)
	}

	mms.streamCancel()
	mms.wait.Wait()
}

func (mms *MemMsgStream) SetRepackFunc(repackFunc RepackFunc) {
	mms.repackFunc = repackFunc
}

func (mms *MemMsgStream) AsProducer(channels []string) {
	for _, channel := range channels {
		err := Mmq.CreateChannel(channel)
		if err == nil {
			mms.producers = append(mms.producers, channel)
		} else {
			errMsg := "Failed to create producer " + channel + ", error = " + err.Error()
			panic(errMsg)
		}
	}
}

func (mms *MemMsgStream) AsConsumer(channels []string, groupName string) {
	for _, channelName := range channels {
		consumer, err := Mmq.CreateConsumerGroup(groupName, channelName)
		if err == nil {
			mms.consumers = append(mms.consumers, consumer)

			mms.wait.Add(1)
			go mms.receiveMsg(*consumer)
		}
	}
}

func (mms *MemMsgStream) Produce(ctx context.Context, pack *msgstream.MsgPack) error {
	tsMsgs := pack.Msgs
	if len(tsMsgs) <= 0 {
		log.Printf("Warning: Receive empty msgPack")
		return nil
	}
	if len(mms.producers) <= 0 {
		return errors.New("nil producer in msg stream")
	}
	reBucketValues := make([][]int32, len(tsMsgs))
	for channelID, tsMsg := range tsMsgs {
		hashValues := tsMsg.HashKeys()
		bucketValues := make([]int32, len(hashValues))
		for index, hashValue := range hashValues {
			if tsMsg.Type() == commonpb.MsgType_SearchResult {
				searchResult := tsMsg.(*msgstream.SearchResultMsg)
				channelID := searchResult.ResultChannelID
				channelIDInt, _ := strconv.ParseInt(channelID, 10, 64)
				if channelIDInt >= int64(len(mms.producers)) {
					return errors.New("Failed to produce rmq msg to unKnow channel")
				}
				bucketValues[index] = int32(channelIDInt)
				continue
			}
			bucketValues[index] = int32(hashValue % uint32(len(mms.producers)))
		}
		reBucketValues[channelID] = bucketValues
	}

	var result map[int32]*msgstream.MsgPack
	var err error
	if mms.repackFunc != nil {
		result, err = mms.repackFunc(tsMsgs, reBucketValues)
	} else {
		msgType := (tsMsgs[0]).Type()
		switch msgType {
		case commonpb.MsgType_Insert:
			result, err = util.InsertRepackFunc(tsMsgs, reBucketValues)
		case commonpb.MsgType_Delete:
			result, err = util.DeleteRepackFunc(tsMsgs, reBucketValues)
		default:
			result, err = util.DefaultRepackFunc(tsMsgs, reBucketValues)
		}
	}
	if err != nil {
		return err
	}
	for k, v := range result {
		err := Mmq.Produce(mms.producers[k], v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mms *MemMsgStream) Broadcast(ctx context.Context, msgPack *MsgPack) error {
	for _, channelName := range mms.producers {
		err := Mmq.Produce(channelName, msgPack)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mms *MemMsgStream) Consume() (*msgstream.MsgPack, context.Context) {
	for {
		select {
		case cm, ok := <-mms.receiveBuf:
			if !ok {
				log.Println("buf chan closed")
				return nil, nil
			}
			return cm, nil
		case <-mms.ctx.Done():
			log.Printf("context closed")
			return nil, nil
		}
	}
}

/**
receiveMsg func is used to solve search timeout problem
which is caused by selectcase
*/
func (mms *MemMsgStream) receiveMsg(consumer Consumer) {
	defer mms.wait.Done()
	for {
		select {
		case <-mms.ctx.Done():
			return
		case msg := <-consumer.MsgChan:
			if msg == nil {
				return
			}

			mms.receiveBuf <- msg
		}
	}
}

func (mms *MemMsgStream) Chan() <-chan *msgstream.MsgPack {
	return mms.receiveBuf
}

func (mms *MemMsgStream) Seek(offset *msgstream.MsgPosition) error {
	return errors.New("MemMsgStream seek not implemented")
}
