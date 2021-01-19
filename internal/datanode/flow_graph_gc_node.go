package datanode

import (
	"log"
)

type gcNode struct {
	BaseNode
	replica collectionReplica
}

func (gcNode *gcNode) Name() string {
	return "gcNode"
}

func (gcNode *gcNode) Operate(in []*Msg) []*Msg {
	//fmt.Println("Do gcNode operation")

	if len(in) != 1 {
		log.Println("Invalid operate message input in gcNode, input length = ", len(in))
		// TODO: add error handling
	}

	gcMsg, ok := (*in[0]).(*gcMsg)
	if !ok {
		log.Println("type assertion failed for gcMsg")
		// TODO: add error handling
	}

	// drop collections
	for _, collectionID := range gcMsg.gcRecord.collections {
		err := gcNode.replica.removeCollection(collectionID)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func newGCNode(replica collectionReplica) *gcNode {
	maxQueueLength := Params.FlowGraphMaxQueueLength
	maxParallelism := Params.FlowGraphMaxParallelism

	baseNode := BaseNode{}
	baseNode.SetMaxQueueLength(maxQueueLength)
	baseNode.SetMaxParallelism(maxParallelism)

	return &gcNode{
		BaseNode: baseNode,
		replica:  replica,
	}
}
