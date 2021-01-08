package writenode

import (
	"context"
)

type WriteNode struct {
	ctx              context.Context
	WriteNodeID      uint64
	dataSyncService  *dataSyncService
	flushSyncService *flushSyncService
	metaService      *metaService
	replica          collectionReplica
}

func NewWriteNode(ctx context.Context, writeNodeID uint64) *WriteNode {

	collections := make([]*Collection, 0)

	var replica collectionReplica = &collectionReplicaImpl{
		collections: collections,
	}

	node := &WriteNode{
		ctx:              ctx,
		WriteNodeID:      writeNodeID,
		dataSyncService:  nil,
		flushSyncService: nil,
		metaService:      nil,
		replica:          replica,
	}

	return node
}

func Init() {
	Params.Init()
}

func (node *WriteNode) Start() error {

	// TODO GOOSE Init Size??
	chanSize := 100
	ddChan := make(chan *ddlFlushSyncMsg, chanSize)
	insertChan := make(chan *insertFlushSyncMsg, chanSize)
	node.flushSyncService = newFlushSyncService(node.ctx, ddChan, insertChan)

	node.dataSyncService = newDataSyncService(node.ctx, ddChan, insertChan, node.replica)
	node.metaService = newMetaService(node.ctx, node.replica)

	go node.dataSyncService.start()
	go node.flushSyncService.start()
	node.metaService.start()
	return nil
}

func (node *WriteNode) Close() {
	<-node.ctx.Done()

	// close services
	if node.dataSyncService != nil {
		(*node.dataSyncService).close()
	}
}
