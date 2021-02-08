package querynode

/*

#cgo CFLAGS: -I${SRCDIR}/../core/output/include

#cgo LDFLAGS: -L${SRCDIR}/../core/output/lib -lmilvus_segcore -Wl,-rpath=${SRCDIR}/../core/output/lib

#include "segcore/collection_c.h"
#include "segcore/segment_c.h"

*/
import "C"

import (
	"context"
	"errors"
	"fmt"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"io"
	"log"
	"sync/atomic"

	"github.com/zilliztech/milvus-distributed/internal/msgstream/pulsarms"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	queryPb "github.com/zilliztech/milvus-distributed/internal/proto/querypb"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
)

type Node interface {
	typeutil.Component

	AddQueryChannel(in *queryPb.AddQueryChannelsRequest) (*commonpb.Status, error)
	RemoveQueryChannel(in *queryPb.RemoveQueryChannelsRequest) (*commonpb.Status, error)
	WatchDmChannels(in *queryPb.WatchDmChannelsRequest) (*commonpb.Status, error)
	LoadSegments(in *queryPb.LoadSegmentRequest) (*commonpb.Status, error)
	ReleaseSegments(in *queryPb.ReleaseSegmentRequest) (*commonpb.Status, error)
	GetSegmentInfo(in *queryPb.SegmentInfoRequest) (*queryPb.SegmentInfoResponse, error)
}

type QueryService = typeutil.QueryServiceInterface

type QueryNode struct {
	typeutil.Service

	queryNodeLoopCtx    context.Context
	queryNodeLoopCancel context.CancelFunc

	QueryNodeID uint64
	stateCode   atomic.Value

	replica collectionReplica

	// internal services
	dataSyncService *dataSyncService
	metaService     *metaService
	searchService   *searchService
	loadService     *loadService
	statsService    *statsService

	//opentracing
	closer io.Closer

	// clients
	masterClient MasterServiceInterface
	queryClient  QueryServiceInterface
	indexClient  IndexServiceInterface
	dataClient   DataServiceInterface

	msFactory msgstream.Factory
}

func NewQueryNode(ctx context.Context, queryNodeID uint64, factory msgstream.Factory) *QueryNode {
	ctx1, cancel := context.WithCancel(ctx)
	node := &QueryNode{
		queryNodeLoopCtx:    ctx1,
		queryNodeLoopCancel: cancel,
		QueryNodeID:         queryNodeID,

		dataSyncService: nil,
		metaService:     nil,
		searchService:   nil,
		statsService:    nil,

		msFactory: factory,
	}

	node.replica = newCollectionReplicaImpl()
	node.stateCode.Store(internalpb2.StateCode_INITIALIZING)
	return node
}

func NewQueryNodeWithoutID(ctx context.Context, factory msgstream.Factory) *QueryNode {
	ctx1, cancel := context.WithCancel(ctx)
	node := &QueryNode{
		queryNodeLoopCtx:    ctx1,
		queryNodeLoopCancel: cancel,

		dataSyncService: nil,
		metaService:     nil,
		searchService:   nil,
		statsService:    nil,

		msFactory: factory,
	}

	node.replica = newCollectionReplicaImpl()
	node.stateCode.Store(internalpb2.StateCode_INITIALIZING)
	return node
}

// TODO: delete this and call node.Init()
func Init() {
	Params.Init()
}

func (node *QueryNode) Init() error {
	registerReq := &queryPb.RegisterNodeRequest{
		Address: &commonpb.Address{
			Ip:   Params.QueryNodeIP,
			Port: Params.QueryNodePort,
		},
	}

	response, err := node.queryClient.RegisterNode(registerReq)
	if err != nil {
		panic(err)
	}
	if response.Status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		panic(response.Status.Reason)
	}

	Params.QueryNodeID = response.InitParams.NodeID
	fmt.Println("QueryNodeID is", Params.QueryNodeID)

	if node.masterClient == nil {
		log.Println("WARN: null master service detected")
	}

	if node.indexClient == nil {
		log.Println("WARN: null index service detected")
	}

	if node.dataClient == nil {
		log.Println("WARN: null data service detected")
	}

	return nil
}

func (node *QueryNode) Start() error {
	var err error
	m := map[string]interface{}{
		"PulsarAddress":  Params.PulsarAddress,
		"ReceiveBufSize": 1024,
		"PulsarBufSize":  1024}
	err = node.msFactory.SetParams(m)
	if err != nil {
		return err
	}

	// init services and manager
	node.dataSyncService = newDataSyncService(node.queryNodeLoopCtx, node.replica, node.msFactory)
	node.searchService = newSearchService(node.queryNodeLoopCtx, node.replica, node.msFactory)
	//node.metaService = newMetaService(node.queryNodeLoopCtx, node.replica)

	node.loadService = newLoadService(node.queryNodeLoopCtx, node.masterClient, node.dataClient, node.indexClient, node.replica, node.dataSyncService.dmStream)
	node.statsService = newStatsService(node.queryNodeLoopCtx, node.replica, node.loadService.segLoader.indexLoader.fieldStatsChan, node.msFactory)

	// start services
	go node.dataSyncService.start()
	go node.searchService.start()
	//go node.metaService.start()
	go node.loadService.start()
	go node.statsService.start()

	node.stateCode.Store(internalpb2.StateCode_HEALTHY)
	<-node.queryNodeLoopCtx.Done()
	return nil
}

func (node *QueryNode) Stop() error {
	node.stateCode.Store(internalpb2.StateCode_ABNORMAL)
	node.queryNodeLoopCancel()

	// free collectionReplica
	node.replica.freeAll()

	// close services
	if node.dataSyncService != nil {
		node.dataSyncService.close()
	}
	if node.searchService != nil {
		node.searchService.close()
	}
	if node.loadService != nil {
		node.loadService.close()
	}
	if node.statsService != nil {
		node.statsService.close()
	}
	if node.closer != nil {
		node.closer.Close()
	}
	return nil
}

func (node *QueryNode) SetMasterService(master MasterServiceInterface) error {
	if master == nil {
		return errors.New("null master service interface")
	}
	node.masterClient = master
	return nil
}

func (node *QueryNode) SetQueryService(query QueryServiceInterface) error {
	if query == nil {
		return errors.New("null query service interface")
	}
	node.queryClient = query
	return nil
}

func (node *QueryNode) SetIndexService(index IndexServiceInterface) error {
	if index == nil {
		return errors.New("null index service interface")
	}
	node.indexClient = index
	return nil
}

func (node *QueryNode) SetDataService(data DataServiceInterface) error {
	if data == nil {
		return errors.New("null data service interface")
	}
	node.dataClient = data
	return nil
}

func (node *QueryNode) GetComponentStates() (*internalpb2.ComponentStates, error) {
	code, ok := node.stateCode.Load().(internalpb2.StateCode)
	if !ok {
		return nil, errors.New("unexpected error in type assertion")
	}
	info := &internalpb2.ComponentInfo{
		NodeID:    Params.QueryNodeID,
		Role:      typeutil.QueryNodeRole,
		StateCode: code,
	}
	stats := &internalpb2.ComponentStates{
		State: info,
	}
	return stats, nil
}

func (node *QueryNode) GetTimeTickChannel() (string, error) {
	return Params.QueryTimeTickChannelName, nil
}

func (node *QueryNode) GetStatisticsChannel() (string, error) {
	return Params.StatsChannelName, nil
}

func (node *QueryNode) AddQueryChannel(in *queryPb.AddQueryChannelsRequest) (*commonpb.Status, error) {
	if node.searchService == nil || node.searchService.searchMsgStream == nil {
		errMsg := "null search service or null search message stream"
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    errMsg,
		}

		return status, errors.New(errMsg)
	}

	searchStream, ok := node.searchService.searchMsgStream.(*pulsarms.PulsarMsgStream)
	if !ok {
		errMsg := "type assertion failed for search message stream"
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    errMsg,
		}

		return status, errors.New(errMsg)
	}

	resultStream, ok := node.searchService.searchResultMsgStream.(*pulsarms.PulsarMsgStream)
	if !ok {
		errMsg := "type assertion failed for search result message stream"
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    errMsg,
		}

		return status, errors.New(errMsg)
	}

	// add request channel
	consumeChannels := []string{in.RequestChannelID}
	consumeSubName := Params.MsgChannelSubName
	searchStream.AsConsumer(consumeChannels, consumeSubName)

	// add result channel
	producerChannels := []string{in.ResultChannelID}
	resultStream.AsProducer(producerChannels)

	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
	}
	return status, nil
}

func (node *QueryNode) RemoveQueryChannel(in *queryPb.RemoveQueryChannelsRequest) (*commonpb.Status, error) {
	if node.searchService == nil || node.searchService.searchMsgStream == nil {
		errMsg := "null search service or null search result message stream"
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    errMsg,
		}

		return status, errors.New(errMsg)
	}

	searchStream, ok := node.searchService.searchMsgStream.(*pulsarms.PulsarMsgStream)
	if !ok {
		errMsg := "type assertion failed for search message stream"
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    errMsg,
		}

		return status, errors.New(errMsg)
	}

	resultStream, ok := node.searchService.searchResultMsgStream.(*pulsarms.PulsarMsgStream)
	if !ok {
		errMsg := "type assertion failed for search result message stream"
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    errMsg,
		}

		return status, errors.New(errMsg)
	}

	// remove request channel
	consumeChannels := []string{in.RequestChannelID}
	consumeSubName := Params.MsgChannelSubName
	// TODO: searchStream.RemovePulsarConsumers(producerChannels)
	searchStream.AsConsumer(consumeChannels, consumeSubName)

	// remove result channel
	producerChannels := []string{in.ResultChannelID}
	// TODO: resultStream.RemovePulsarProducer(producerChannels)
	resultStream.AsProducer(producerChannels)

	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
	}
	return status, nil
}

func (node *QueryNode) WatchDmChannels(in *queryPb.WatchDmChannelsRequest) (*commonpb.Status, error) {
	if node.dataSyncService == nil || node.dataSyncService.dmStream == nil {
		errMsg := "null data sync service or null data manipulation stream"
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    errMsg,
		}

		return status, errors.New(errMsg)
	}

	fgDMMsgStream, ok := node.dataSyncService.dmStream.(*pulsarms.PulsarTtMsgStream)
	if !ok {
		errMsg := "type assertion failed for dm message stream"
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    errMsg,
		}

		return status, errors.New(errMsg)
	}

	// add request channel
	consumeChannels := in.ChannelIDs
	consumeSubName := Params.MsgChannelSubName
	fgDMMsgStream.AsConsumer(consumeChannels, consumeSubName)

	status := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
	}
	return status, nil
}

func (node *QueryNode) LoadSegments(in *queryPb.LoadSegmentRequest) (*commonpb.Status, error) {
	// TODO: support db
	collectionID := in.CollectionID
	partitionID := in.PartitionID
	segmentIDs := in.SegmentIDs
	fieldIDs := in.FieldIDs
	schema := in.Schema

	hasCollection := node.replica.hasCollection(collectionID)
	hasPartition := node.replica.hasPartition(partitionID)
	if !hasCollection {
		err := node.replica.addCollection(collectionID, schema)
		if err != nil {
			status := &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			}
			return status, err
		}
	}
	if !hasPartition {
		err := node.replica.addPartition(collectionID, partitionID)
		if err != nil {
			status := &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			}
			return status, err
		}
	}
	err := node.replica.enablePartition(partitionID)
	if err != nil {
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    err.Error(),
		}
		return status, err
	}

	// segments are ordered before LoadSegments calling
	for i, state := range in.SegmentStates {
		if state.State == commonpb.SegmentState_SegmentGrowing {
			position := state.StartPosition
			err := node.loadService.segLoader.seekSegment(position)
			if err != nil {
				status := &commonpb.Status{
					ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
					Reason:    err.Error(),
				}
				return status, err
			}
			segmentIDs = segmentIDs[:i]
			break
		}
	}

	err = node.loadService.loadSegment(collectionID, partitionID, segmentIDs, fieldIDs)
	if err != nil {
		status := &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    err.Error(),
		}
		return status, err
	}
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
	}, nil
}

func (node *QueryNode) ReleaseSegments(in *queryPb.ReleaseSegmentRequest) (*commonpb.Status, error) {
	for _, id := range in.PartitionIDs {
		err := node.replica.enablePartition(id)
		if err != nil {
			status := &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			}
			return status, err
		}
	}

	// release all fields in the segments
	for _, id := range in.SegmentIDs {
		err := node.loadService.segLoader.releaseSegment(id)
		if err != nil {
			status := &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			}
			return status, err
		}
	}
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
	}, nil
}

func (node *QueryNode) GetSegmentInfo(in *queryPb.SegmentInfoRequest) (*queryPb.SegmentInfoResponse, error) {
	infos := make([]*queryPb.SegmentInfo, 0)
	for _, id := range in.SegmentIDs {
		segment, err := node.replica.getSegmentByID(id)
		if err != nil {
			continue
		}
		info := &queryPb.SegmentInfo{
			SegmentID:    segment.ID(),
			CollectionID: segment.collectionID,
			PartitionID:  segment.partitionID,
			MemSize:      segment.getMemSize(),
			NumRows:      segment.getRowCount(),
			IndexName:    segment.getIndexName(),
			IndexID:      segment.getIndexID(),
		}
		infos = append(infos, info)
	}
	return &queryPb.SegmentInfoResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
		},
		Infos: infos,
	}, nil
}
