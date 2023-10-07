// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datanode

import (
	"context"
	"sync"

	"github.com/cockroachdb/errors"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus/internal/datanode/allocator"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/internal/types"
	"github.com/milvus-io/milvus/internal/util/flowgraph"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/mq/msgdispatcher"
	"github.com/milvus-io/milvus/pkg/mq/msgstream"
	"github.com/milvus-io/milvus/pkg/util/commonpbutil"
	"github.com/milvus-io/milvus/pkg/util/conc"
	"github.com/milvus-io/milvus/pkg/util/merr"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
	"github.com/milvus-io/milvus/pkg/util/retry"
)

// dataSyncService controls a flowgraph for a specific collection
type dataSyncService struct {
	ctx          context.Context
	cancelFn     context.CancelFunc
	channel      Channel  // channel stores meta of channel
	collectionID UniqueID // collection id of vchan for which this data sync service serves
	vchannelName string

	// TODO: should be equal to paramtable.GetNodeID(), but intergrationtest has 1 paramtable for a minicluster, the NodeID
	// varies, will cause savebinglogpath check fail. So we pass ServerID into dataSyncService to aviod it failure.
	serverID UniqueID

	fg *flowgraph.TimeTickedFlowGraph // internal flowgraph processes insert/delta messages

	delBufferManager *DeltaBufferManager
	flushManager     flushManager // flush manager handles flush process

	flushCh          chan flushMsg
	resendTTCh       chan resendTTMsg    // chan to ask for resending DataNode time tick message.
	timetickSender   *timeTickSender     // reference to timeTickSender
	compactor        *compactionExecutor // reference to compaction executor
	flushingSegCache *Cache              // a guarding cache stores currently flushing segment ids

	clearSignal  chan<- string       // signal channel to notify flowgraph close for collection/partition drop msg consumed
	idAllocator  allocator.Allocator // id/timestamp allocator
	msFactory    msgstream.Factory
	dispClient   msgdispatcher.Client
	dataCoord    types.DataCoordClient // DataCoord instance to interact with
	chunkManager storage.ChunkManager

	// test only
	flushListener chan *segmentFlushPack // chan to listen flush event

	stopOnce sync.Once
}

type nodeConfig struct {
	msFactory    msgstream.Factory // msgStream factory
	collectionID UniqueID
	vChannelName string
	channel      Channel // Channel info
	allocator    allocator.Allocator
	serverID     UniqueID
}

// start the flow graph in datasyncservice
func (dsService *dataSyncService) start() {
	if dsService.fg != nil {
		log.Info("dataSyncService starting flow graph", zap.Int64("collectionID", dsService.collectionID),
			zap.String("vChanName", dsService.vchannelName))
		dsService.fg.Start()
	} else {
		log.Warn("dataSyncService starting flow graph is nil", zap.Int64("collectionID", dsService.collectionID),
			zap.String("vChanName", dsService.vchannelName))
	}
}

func (dsService *dataSyncService) GracefullyClose() {
	if dsService.fg != nil {
		log.Info("dataSyncService gracefully closing flowgraph")
		dsService.fg.SetCloseMethod(flowgraph.CloseGracefully)
		dsService.close()
	}
}

func (dsService *dataSyncService) close() {
	dsService.stopOnce.Do(func() {
		log := log.Ctx(dsService.ctx).With(
			zap.Int64("collectionID", dsService.collectionID),
			zap.String("vChanName", dsService.vchannelName),
		)
		if dsService.fg != nil {
			log.Info("dataSyncService closing flowgraph")
			dsService.dispClient.Deregister(dsService.vchannelName)
			dsService.fg.Close()
			log.Info("dataSyncService flowgraph closed")
		}

		dsService.clearGlobalFlushingCache()
		close(dsService.flushCh)

		if dsService.flushManager != nil {
			dsService.flushManager.close()
			log.Info("dataSyncService flush manager closed")
		}

		dsService.cancelFn()
		dsService.channel.close()

		log.Info("dataSyncService closed")
	})
}

func (dsService *dataSyncService) clearGlobalFlushingCache() {
	segments := dsService.channel.listAllSegmentIDs()
	dsService.flushingSegCache.Remove(segments...)
}

// getSegmentInfos return the SegmentInfo details according to the given ids through RPC to datacoord
// TODO: add a broker for the rpc
func getSegmentInfos(ctx context.Context, datacoord types.DataCoordClient, segmentIDs []int64) ([]*datapb.SegmentInfo, error) {
	infoResp, err := datacoord.GetSegmentInfo(ctx, &datapb.GetSegmentInfoRequest{
		Base: commonpbutil.NewMsgBase(
			commonpbutil.WithMsgType(commonpb.MsgType_SegmentInfo),
			commonpbutil.WithMsgID(0),
			commonpbutil.WithSourceID(paramtable.GetNodeID()),
		),
		SegmentIDs:       segmentIDs,
		IncludeUnHealthy: true,
	})
	if err := merr.CheckRPCCall(infoResp, err); err != nil {
		log.Error("Fail to get SegmentInfo by ids from datacoord", zap.Error(err))
		return nil, err
	}

	return infoResp.Infos, nil
}

// getChannelWithEtcdTickler updates progress into etcd when a new segment is added into channel.
func getChannelWithEtcdTickler(initCtx context.Context, node *DataNode, info *datapb.ChannelWatchInfo, tickler *etcdTickler, unflushed, flushed []*datapb.SegmentInfo) (Channel, error) {
	var (
		channelName  = info.GetVchan().GetChannelName()
		collectionID = info.GetVchan().GetCollectionID()
		recoverTs    = info.GetVchan().GetSeekPosition().GetTimestamp()
	)

	// init channel meta
	channel := newChannel(channelName, collectionID, info.GetSchema(), node.rootCoord, node.chunkManager)

	// tickler will update addSegment progress to watchInfo
	tickler.watch()
	defer tickler.stop()
	futures := make([]*conc.Future[any], 0, len(unflushed)+len(flushed))

	for _, us := range unflushed {
		log.Info("recover growing segments from checkpoints",
			zap.String("vChannelName", us.GetInsertChannel()),
			zap.Int64("segmentID", us.GetID()),
			zap.Int64("numRows", us.GetNumOfRows()),
		)

		// avoid closure capture iteration variable
		segment := us
		future := getOrCreateIOPool().Submit(func() (interface{}, error) {
			if err := channel.addSegment(initCtx, addSegmentReq{
				segType:      datapb.SegmentType_Normal,
				segID:        segment.GetID(),
				collID:       segment.CollectionID,
				partitionID:  segment.PartitionID,
				numOfRows:    segment.GetNumOfRows(),
				statsBinLogs: segment.Statslogs,
				binLogs:      segment.GetBinlogs(),
				endPos:       segment.GetDmlPosition(),
				recoverTs:    recoverTs,
			}); err != nil {
				return nil, err
			}
			tickler.inc()
			return nil, nil
		})
		futures = append(futures, future)
	}

	for _, fs := range flushed {
		log.Info("recover sealed segments form checkpoints",
			zap.String("vChannelName", fs.GetInsertChannel()),
			zap.Int64("segmentID", fs.GetID()),
			zap.Int64("numRows", fs.GetNumOfRows()),
		)
		// avoid closure capture iteration variable
		segment := fs
		future := getOrCreateIOPool().Submit(func() (interface{}, error) {
			if err := channel.addSegment(initCtx, addSegmentReq{
				segType:      datapb.SegmentType_Flushed,
				segID:        segment.GetID(),
				collID:       segment.GetCollectionID(),
				partitionID:  segment.GetPartitionID(),
				numOfRows:    segment.GetNumOfRows(),
				statsBinLogs: segment.GetStatslogs(),
				binLogs:      segment.GetBinlogs(),
				recoverTs:    recoverTs,
			}); err != nil {
				return nil, err
			}
			tickler.inc()
			return nil, nil
		})
		futures = append(futures, future)
	}

	if err := conc.AwaitAll(futures...); err != nil {
		return nil, err
	}

	if tickler.isWatchFailed.Load() {
		return nil, errors.Errorf("tickler watch failed")
	}
	return channel, nil
}

func getServiceWithChannel(initCtx context.Context, node *DataNode, info *datapb.ChannelWatchInfo, channel Channel, unflushed, flushed []*datapb.SegmentInfo) (*dataSyncService, error) {
	var (
		channelName  = info.GetVchan().GetChannelName()
		collectionID = info.GetVchan().GetCollectionID()
	)

	config := &nodeConfig{
		msFactory: node.factory,
		allocator: node.allocator,

		collectionID: collectionID,
		vChannelName: channelName,
		channel:      channel,
		serverID:     node.session.ServerID,
	}

	var (
		flushCh          = make(chan flushMsg, 100)
		resendTTCh       = make(chan resendTTMsg, 100)
		delBufferManager = &DeltaBufferManager{
			channel:    channel,
			delBufHeap: &PriorityQueue{},
		}
	)

	ctx, cancel := context.WithCancel(node.ctx)
	ds := &dataSyncService{
		ctx:              ctx,
		cancelFn:         cancel,
		flushCh:          flushCh,
		resendTTCh:       resendTTCh,
		delBufferManager: delBufferManager,

		dispClient: node.dispClient,
		msFactory:  node.factory,
		dataCoord:  node.dataCoord,

		idAllocator:  config.allocator,
		channel:      config.channel,
		collectionID: config.collectionID,
		vchannelName: config.vChannelName,
		serverID:     config.serverID,

		flushingSegCache: node.segmentCache,
		clearSignal:      node.clearSignal,
		chunkManager:     node.chunkManager,
		compactor:        node.compactionExecutor,
		timetickSender:   node.timeTickSender,

		fg:           nil,
		flushManager: nil,
	}

	// init flushManager
	flushManager := NewRendezvousFlushManager(
		node.allocator,
		node.chunkManager,
		channel,
		flushNotifyFunc(ds, retry.Attempts(50)), dropVirtualChannelFunc(ds),
	)
	ds.flushManager = flushManager

	// init flowgraph
	fg := flowgraph.NewTimeTickedFlowGraph(node.ctx)
	dmStreamNode, err := newDmInputNode(initCtx, node.dispClient, info.GetVchan().GetSeekPosition(), config)
	if err != nil {
		return nil, err
	}

	ddNode, err := newDDNode(
		node.ctx,
		collectionID,
		channelName,
		info.GetVchan().GetDroppedSegmentIds(),
		flushed,
		unflushed,
		node.compactionExecutor,
	)
	if err != nil {
		return nil, err
	}

	insertBufferNode, err := newInsertBufferNode(
		node.ctx,
		flushCh,
		resendTTCh,
		delBufferManager,
		flushManager,
		node.segmentCache,
		node.timeTickSender,
		config,
	)
	if err != nil {
		return nil, err
	}

	deleteNode, err := newDeleteNode(node.ctx, flushManager, delBufferManager, node.clearSignal, config)
	if err != nil {
		return nil, err
	}

	ttNode, err := newTTNode(config, node.dataCoord)
	if err != nil {
		return nil, err
	}

	if err := fg.AssembleNodes(dmStreamNode, ddNode, insertBufferNode, deleteNode, ttNode); err != nil {
		return nil, err
	}
	ds.fg = fg

	return ds, nil
}

// newServiceWithEtcdTickler gets a dataSyncService, but flowgraphs are not running
// initCtx is used to init the dataSyncService only, if initCtx.Canceled or initCtx.Timeout
// newServiceWithEtcdTickler stops and returns the initCtx.Err()
func newServiceWithEtcdTickler(initCtx context.Context, node *DataNode, info *datapb.ChannelWatchInfo, tickler *etcdTickler) (*dataSyncService, error) {
	// recover segment checkpoints
	unflushedSegmentInfos, err := getSegmentInfos(initCtx, node.dataCoord, info.GetVchan().GetUnflushedSegmentIds())
	if err != nil {
		return nil, err
	}
	flushedSegmentInfos, err := getSegmentInfos(initCtx, node.dataCoord, info.GetVchan().GetFlushedSegmentIds())
	if err != nil {
		return nil, err
	}

	// init channel meta
	channel, err := getChannelWithEtcdTickler(initCtx, node, info, tickler, unflushedSegmentInfos, flushedSegmentInfos)
	if err != nil {
		return nil, err
	}

	return getServiceWithChannel(initCtx, node, info, channel, unflushedSegmentInfos, flushedSegmentInfos)
}
