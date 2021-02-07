package querynode

import (
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
)

type ddNode struct {
	baseNode
	ddMsg   *ddMsg
	replica collectionReplica
}

func (ddNode *ddNode) Name() string {
	return "ddNode"
}

func (ddNode *ddNode) Operate(in []*Msg) []*Msg {
	//fmt.Println("Do filterDmNode operation")

	if len(in) != 1 {
		log.Println("Invalid operate message input in ddNode, input length = ", len(in))
		// TODO: add error handling
	}

	msMsg, ok := (*in[0]).(*MsgStreamMsg)
	if !ok {
		log.Println("type assertion failed for MsgStreamMsg")
		// TODO: add error handling
	}

	var ddMsg = ddMsg{
		collectionRecords: make(map[UniqueID][]metaOperateRecord),
		partitionRecords:  make(map[UniqueID][]metaOperateRecord),
		timeRange: TimeRange{
			timestampMin: msMsg.TimestampMin(),
			timestampMax: msMsg.TimestampMax(),
		},
	}
	ddNode.ddMsg = &ddMsg
	gcRecord := gcRecord{
		collections: make([]UniqueID, 0),
		partitions:  make([]partitionWithID, 0),
	}
	ddNode.ddMsg.gcRecord = &gcRecord

	// sort tsMessages
	//tsMessages := msMsg.TsMessages()
	//sort.Slice(tsMessages,
	//	func(i, j int) bool {
	//		return tsMessages[i].BeginTs() < tsMessages[j].BeginTs()
	//	})

	// do dd tasks
	//for _, msg := range tsMessages {
	//	switch msg.Type() {
	//	case commonpb.MsgType_kCreateCollection:
	//		ddNode.createCollection(msg.(*msgstream.CreateCollectionMsg))
	//	case commonpb.MsgType_kDropCollection:
	//		ddNode.dropCollection(msg.(*msgstream.DropCollectionMsg))
	//	case commonpb.MsgType_kCreatePartition:
	//		ddNode.createPartition(msg.(*msgstream.CreatePartitionMsg))
	//	case commonpb.MsgType_kDropPartition:
	//		ddNode.dropPartition(msg.(*msgstream.DropPartitionMsg))
	//	default:
	//		log.Println("Non supporting message type:", msg.Type())
	//	}
	//}

	var res Msg = ddNode.ddMsg
	return []*Msg{&res}
}

func (ddNode *ddNode) createCollection(msg *msgstream.CreateCollectionMsg) {
	collectionID := msg.CollectionID

	hasCollection := ddNode.replica.hasCollection(collectionID)
	if hasCollection {
		log.Println("collection already exists, id = ", collectionID)
		return
	}

	var schema schemapb.CollectionSchema
	err := proto.Unmarshal(msg.Schema, &schema)
	if err != nil {
		log.Println(err)
		return
	}

	// add collection
	err = ddNode.replica.addCollection(collectionID, &schema)
	if err != nil {
		log.Println(err)
		return
	}

	// add default partition
	// TODO: allocate default partition id in master
	err = ddNode.replica.addPartition(collectionID, UniqueID(2021))
	if err != nil {
		log.Println(err)
		return
	}

	ddNode.ddMsg.collectionRecords[collectionID] = append(ddNode.ddMsg.collectionRecords[collectionID],
		metaOperateRecord{
			createOrDrop: true,
			timestamp:    msg.Base.Timestamp,
		})
}

func (ddNode *ddNode) dropCollection(msg *msgstream.DropCollectionMsg) {
	collectionID := msg.CollectionID

	ddNode.ddMsg.collectionRecords[collectionID] = append(ddNode.ddMsg.collectionRecords[collectionID],
		metaOperateRecord{
			createOrDrop: false,
			timestamp:    msg.Base.Timestamp,
		})

	ddNode.ddMsg.gcRecord.collections = append(ddNode.ddMsg.gcRecord.collections, collectionID)
}

func (ddNode *ddNode) createPartition(msg *msgstream.CreatePartitionMsg) {
	collectionID := msg.CollectionID
	partitionID := msg.PartitionID

	err := ddNode.replica.addPartition(collectionID, partitionID)
	if err != nil {
		log.Println(err)
		return
	}

	ddNode.ddMsg.partitionRecords[partitionID] = append(ddNode.ddMsg.partitionRecords[partitionID],
		metaOperateRecord{
			createOrDrop: true,
			timestamp:    msg.Base.Timestamp,
		})
}

func (ddNode *ddNode) dropPartition(msg *msgstream.DropPartitionMsg) {
	collectionID := msg.CollectionID
	partitionID := msg.PartitionID

	ddNode.ddMsg.partitionRecords[partitionID] = append(ddNode.ddMsg.partitionRecords[partitionID],
		metaOperateRecord{
			createOrDrop: false,
			timestamp:    msg.Base.Timestamp,
		})

	ddNode.ddMsg.gcRecord.partitions = append(ddNode.ddMsg.gcRecord.partitions, partitionWithID{
		partitionID:  partitionID,
		collectionID: collectionID,
	})
}

func newDDNode(replica collectionReplica) *ddNode {
	maxQueueLength := Params.FlowGraphMaxQueueLength
	maxParallelism := Params.FlowGraphMaxParallelism

	baseNode := baseNode{}
	baseNode.SetMaxQueueLength(maxQueueLength)
	baseNode.SetMaxParallelism(maxParallelism)

	return &ddNode{
		baseNode: baseNode,
		replica:  replica,
	}
}
