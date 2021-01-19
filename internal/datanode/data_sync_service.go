package datanode

import (
	"context"
	"log"

	"github.com/zilliztech/milvus-distributed/internal/util/flowgraph"
)

type dataSyncService struct {
	ctx        context.Context
	fg         *flowgraph.TimeTickedFlowGraph
	ddChan     chan *ddlFlushSyncMsg
	insertChan chan *insertFlushSyncMsg
	replica    collectionReplica
}

func newDataSyncService(ctx context.Context,
	ddChan chan *ddlFlushSyncMsg, insertChan chan *insertFlushSyncMsg, replica collectionReplica) *dataSyncService {

	return &dataSyncService{
		ctx:        ctx,
		fg:         nil,
		ddChan:     ddChan,
		insertChan: insertChan,
		replica:    replica,
	}
}

func (dsService *dataSyncService) start() {
	dsService.initNodes()
	dsService.fg.Start()
}

func (dsService *dataSyncService) close() {
	if dsService.fg != nil {
		dsService.fg.Close()
	}
}

func (dsService *dataSyncService) initNodes() {
	// TODO: add delete pipeline support

	dsService.fg = flowgraph.NewTimeTickedFlowGraph(dsService.ctx)

	var dmStreamNode Node = newDmInputNode(dsService.ctx)
	var ddStreamNode Node = newDDInputNode(dsService.ctx)

	var filterDmNode Node = newFilteredDmNode()

	var ddNode Node = newDDNode(dsService.ctx, dsService.ddChan, dsService.replica)
	var insertBufferNode Node = newInsertBufferNode(dsService.ctx, dsService.insertChan, dsService.replica)
	var gcNode Node = newGCNode(dsService.replica)

	dsService.fg.AddNode(&dmStreamNode)
	dsService.fg.AddNode(&ddStreamNode)

	dsService.fg.AddNode(&filterDmNode)
	dsService.fg.AddNode(&ddNode)

	dsService.fg.AddNode(&insertBufferNode)
	dsService.fg.AddNode(&gcNode)

	// dmStreamNode
	var err = dsService.fg.SetEdges(dmStreamNode.Name(),
		[]string{},
		[]string{filterDmNode.Name()},
	)
	if err != nil {
		log.Fatal("set edges failed in node:", dmStreamNode.Name())
	}

	// ddStreamNode
	err = dsService.fg.SetEdges(ddStreamNode.Name(),
		[]string{},
		[]string{ddNode.Name()},
	)
	if err != nil {
		log.Fatal("set edges failed in node:", ddStreamNode.Name())
	}

	// filterDmNode
	err = dsService.fg.SetEdges(filterDmNode.Name(),
		[]string{dmStreamNode.Name(), ddNode.Name()},
		[]string{insertBufferNode.Name()},
	)
	if err != nil {
		log.Fatal("set edges failed in node:", filterDmNode.Name())
	}

	// ddNode
	err = dsService.fg.SetEdges(ddNode.Name(),
		[]string{ddStreamNode.Name()},
		[]string{filterDmNode.Name()},
	)
	if err != nil {
		log.Fatal("set edges failed in node:", ddNode.Name())
	}

	// insertBufferNode
	err = dsService.fg.SetEdges(insertBufferNode.Name(),
		[]string{filterDmNode.Name()},
		[]string{gcNode.Name()},
	)
	if err != nil {
		log.Fatal("set edges failed in node:", insertBufferNode.Name())
	}

	// gcNode
	err = dsService.fg.SetEdges(gcNode.Name(),
		[]string{insertBufferNode.Name()},
		[]string{})
	if err != nil {
		log.Fatal("set edges failed in node:", gcNode.Name())
	}
}
