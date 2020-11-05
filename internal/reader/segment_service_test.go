package reader

//import (
//	"context"
//	"strconv"
//	"testing"
//	"time"
//
//	"github.com/zilliztech/milvus-distributed/internal/conf"
//	"github.com/zilliztech/milvus-distributed/internal/msgclient"
//)
//
//// NOTE: start pulsar before test
//func TestSegmentManagement_SegmentStatistic(t *testing.T) {
//	conf.LoadConfig("config.yaml")
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	mc := msgclient.ReaderMessageClient{}
//	pulsarAddr := "pulsar://"
//	pulsarAddr += conf.Config.Pulsar.Address
//	pulsarAddr += ":"
//	pulsarAddr += strconv.FormatInt(int64(conf.Config.Pulsar.Port), 10)
//
//	mc.InitClient(ctx, pulsarAddr)
//	mc.ReceiveMessage()
//
//	node := CreateQueryNode(ctx, 0, 0, &mc)
//
//	// Construct node, collection, partition and segment
//	var collection = node.newCollection(0, "collection0", "")
//	var partition = collection.newPartition("partition0")
//	var segment = partition.newSegment(0)
//	node.SegmentsMap[0] = segment
//
//	node.SegmentStatistic(1000)
//
//	node.Close()
//}
//
//// NOTE: start pulsar before test
//func TestSegmentManagement_SegmentStatisticService(t *testing.T) {
//	conf.LoadConfig("config.yaml")
//
//	d := time.Now().Add(ctxTimeInMillisecond * time.Millisecond)
//	ctx, cancel := context.WithDeadline(context.Background(), d)
//	defer cancel()
//
//	mc := msgclient.ReaderMessageClient{}
//	pulsarAddr := "pulsar://"
//	pulsarAddr += conf.Config.Pulsar.Address
//	pulsarAddr += ":"
//	pulsarAddr += strconv.FormatInt(int64(conf.Config.Pulsar.Port), 10)
//
//	mc.InitClient(ctx, pulsarAddr)
//	mc.ReceiveMessage()
//
//	node := CreateQueryNode(ctx, 0, 0, &mc)
//
//	// Construct node, collection, partition and segment
//	var collection = node.newCollection(0, "collection0", "")
//	var partition = collection.newPartition("partition0")
//	var segment = partition.newSegment(0)
//	node.SegmentsMap[0] = segment
//
//	node.SegmentStatisticService()
//
//	node.Close()
//}
