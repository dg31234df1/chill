package querynode

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/msgstream/pulsarms"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
)

type statsService struct {
	ctx context.Context

	replica collectionReplica

	fieldStatsChan chan []*internalpb2.FieldStats
	statsStream    msgstream.MsgStream
}

func newStatsService(ctx context.Context, replica collectionReplica, fieldStatsChan chan []*internalpb2.FieldStats) *statsService {

	return &statsService{
		ctx: ctx,

		replica: replica,

		fieldStatsChan: fieldStatsChan,
		statsStream:    nil,
	}
}

func (sService *statsService) start() {
	sleepTimeInterval := Params.StatsPublishInterval
	receiveBufSize := Params.StatsReceiveBufSize

	// start pulsar
	msgStreamURL := Params.PulsarAddress
	producerChannels := []string{Params.StatsChannelName}

	statsStream := pulsarms.NewPulsarMsgStream(sService.ctx, receiveBufSize)
	statsStream.SetPulsarClient(msgStreamURL)
	statsStream.CreatePulsarProducers(producerChannels)

	var statsMsgStream msgstream.MsgStream = statsStream

	sService.statsStream = statsMsgStream
	sService.statsStream.Start()

	// start service
	fmt.Println("do segments statistic in ", strconv.Itoa(sleepTimeInterval), "ms")
	for {
		select {
		case <-sService.ctx.Done():
			return
		case <-time.After(time.Duration(sleepTimeInterval) * time.Millisecond):
			sService.publicStatistic(nil)
		case fieldStats := <-sService.fieldStatsChan:
			sService.publicStatistic(fieldStats)
		}
	}
}

func (sService *statsService) close() {
	if sService.statsStream != nil {
		sService.statsStream.Close()
	}
}

func (sService *statsService) publicStatistic(fieldStats []*internalpb2.FieldStats) {
	segStats := sService.replica.getSegmentStatistics()

	queryNodeStats := internalpb2.QueryNodeStats{
		Base: &commonpb.MsgBase{
			MsgType:  commonpb.MsgType_kQueryNodeStats,
			SourceID: Params.QueryNodeID,
		},
		SegStats:   segStats,
		FieldStats: fieldStats,
	}

	var msg msgstream.TsMsg = &msgstream.QueryNodeStatsMsg{
		BaseMsg: msgstream.BaseMsg{
			HashValues: []uint32{0},
		},
		QueryNodeStats: queryNodeStats,
	}

	var msgPack = msgstream.MsgPack{
		Msgs: []msgstream.TsMsg{msg},
	}
	err := sService.statsStream.Produce(&msgPack)
	if err != nil {
		log.Println(err)
	}
}
