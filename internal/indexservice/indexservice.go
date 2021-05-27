// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package indexservice

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/golang/protobuf/proto"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/allocator"
	"github.com/milvus-io/milvus/internal/kv"
	etcdkv "github.com/milvus-io/milvus/internal/kv/etcd"
	miniokv "github.com/milvus-io/milvus/internal/kv/minio"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/indexpb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/proto/milvuspb"
	"github.com/milvus-io/milvus/internal/tso"
	"github.com/milvus-io/milvus/internal/util/retry"
	"github.com/milvus-io/milvus/internal/util/sessionutil"
	"github.com/milvus-io/milvus/internal/util/tsoutil"
	"github.com/milvus-io/milvus/internal/util/typeutil"
)

const (
	reqTimeoutInterval = time.Second * 10
	durationInterval   = time.Second * 10
	recycleIndexLimit  = 20
)

type IndexService struct {
	nodeClients *PriorityQueue
	nodeStates  map[UniqueID]*internalpb.ComponentStates
	stateCode   internalpb.StateCode

	ID UniqueID

	loopCtx    context.Context
	loopCancel func()
	loopWg     sync.WaitGroup

	sched   *TaskScheduler
	session *sessionutil.Session

	eventChan <-chan *sessionutil.SessionEvent

	assignChan chan []UniqueID

	idAllocator *allocator.GlobalIDAllocator

	kv kv.BaseKV

	metaTable *metaTable

	nodeTasks *nodeTasks

	nodeLock sync.RWMutex

	// Add callback functions at different stages
	startCallbacks []func()
	closeCallbacks []func()
}

type UniqueID = typeutil.UniqueID
type Timestamp = typeutil.Timestamp

func NewIndexService(ctx context.Context) (*IndexService, error) {
	rand.Seed(time.Now().UnixNano())
	ctx1, cancel := context.WithCancel(ctx)
	i := &IndexService{
		loopCtx:     ctx1,
		loopCancel:  cancel,
		nodeClients: &PriorityQueue{},
		nodeTasks:   &nodeTasks{},
	}

	return i, nil
}

// Register register index service at etcd
func (i *IndexService) Register() error {
	i.session = sessionutil.NewSession(i.loopCtx, Params.MetaRootPath, []string{Params.EtcdAddress})
	i.session.Init(typeutil.IndexServiceRole, Params.Address, true)
	i.eventChan = i.session.WatchServices(typeutil.IndexNodeRole, 0)
	return nil
}

func (i *IndexService) Init() error {
	log.Debug("indexservice", zap.String("etcd address", Params.EtcdAddress))

	i.assignChan = make(chan []UniqueID, 1024)
	connectEtcdFn := func() error {
		etcdClient, err := clientv3.New(clientv3.Config{Endpoints: []string{Params.EtcdAddress}})
		if err != nil {
			return err
		}
		etcdKV := etcdkv.NewEtcdKV(etcdClient, Params.MetaRootPath)
		metakv, err := NewMetaTable(etcdKV)
		if err != nil {
			return err
		}
		i.metaTable = metakv
		return err
	}
	err := retry.Retry(100000, time.Millisecond*200, connectEtcdFn)
	if err != nil {
		return err
	}

	//init idAllocator
	kvRootPath := Params.KvRootPath
	i.idAllocator = allocator.NewGlobalIDAllocator("idTimestamp", tsoutil.NewTSOKVBase([]string{Params.EtcdAddress}, kvRootPath, "index_gid"))
	if err := i.idAllocator.Initialize(); err != nil {
		return err
	}

	i.ID, err = i.idAllocator.AllocOne()
	if err != nil {
		return err
	}

	option := &miniokv.Option{
		Address:           Params.MinIOAddress,
		AccessKeyID:       Params.MinIOAccessKeyID,
		SecretAccessKeyID: Params.MinIOSecretAccessKey,
		UseSSL:            Params.MinIOUseSSL,
		BucketName:        Params.MinioBucketName,
		CreateBucket:      true,
	}

	i.kv, err = miniokv.NewMinIOKV(i.loopCtx, option)
	if err != nil {
		return err
	}

	i.sched, err = NewTaskScheduler(i.loopCtx, i.idAllocator, i.kv, i.metaTable)
	if err != nil {
		return err
	}
	i.UpdateStateCode(internalpb.StateCode_Healthy)

	i.nodeTasks = NewNodeTasks()

	err = i.assignTasksServerStart()
	if err != nil {
		return err
	}
	return nil
}

func (i *IndexService) Start() error {
	i.loopWg.Add(1)
	go i.tsLoop()

	i.loopWg.Add(1)
	go i.recycleUnusedIndexFiles()

	i.loopWg.Add(1)
	go i.assignmentTasksLoop()

	i.loopWg.Add(1)
	go i.watchNodeLoop()

	i.loopWg.Add(1)
	go i.watchMetaLoop()

	i.sched.Start()
	// Start callbacks
	for _, cb := range i.startCallbacks {
		cb()
	}
	log.Debug("IndexService  start")

	return nil
}

func (i *IndexService) Stop() error {
	i.loopCancel()
	i.sched.Close()
	for _, cb := range i.closeCallbacks {
		cb()
	}
	return nil
}

func (i *IndexService) UpdateStateCode(code internalpb.StateCode) {
	i.stateCode = code
}

func (i *IndexService) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	log.Debug("get indexservice component states ...")
	stateInfo := &internalpb.ComponentInfo{
		NodeID:    i.ID,
		Role:      "IndexService",
		StateCode: i.stateCode,
	}

	ret := &internalpb.ComponentStates{
		State:              stateInfo,
		SubcomponentStates: nil, // todo add subcomponents states
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
	}
	return ret, nil
}

func (i *IndexService) GetTimeTickChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	log.Debug("get indexservice time tick channel ...")
	return &milvuspb.StringResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
			Reason:    "",
		},
		Value: "",
	}, nil
}

func (i *IndexService) GetStatisticsChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	log.Debug("get indexservice statistics channel ...")
	return &milvuspb.StringResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
			Reason:    "",
		},
		Value: "",
	}, nil
}

func (i *IndexService) BuildIndex(ctx context.Context, req *indexpb.BuildIndexRequest) (*indexpb.BuildIndexResponse, error) {
	log.Debug("IndexService building index ...",
		zap.Int64("IndexBuildID", req.IndexBuildID),
		zap.String("IndexName = ", req.IndexName),
		zap.Int64("IndexID = ", req.IndexID),
		zap.Strings("DataPath = ", req.DataPaths),
		zap.Any("TypeParams", req.TypeParams),
		zap.Any("IndexParams", req.IndexParams))
	hasIndex, indexBuildID := i.metaTable.HasSameReq(req)
	if hasIndex {
		log.Debug("IndexService", zap.Any("hasIndex true", indexBuildID))
		return &indexpb.BuildIndexResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_Success,
				Reason:    "already have same index",
			},
			IndexBuildID: indexBuildID,
		}, nil
	}
	ret := &indexpb.BuildIndexResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
		},
	}
	t := &IndexAddTask{
		BaseTask: BaseTask{
			ctx:   ctx,
			done:  make(chan error),
			table: i.metaTable,
		},
		req:         req,
		idAllocator: i.idAllocator,
		kv:          i.kv,
	}

	if i.nodeClients == nil || i.nodeClients.Len() <= 0 {
		ret.Status.Reason = "IndexBuilding Service not available"
		return ret, nil
	}
	t.nodeClients = i.nodeClients

	var cancel func()
	t.ctx, cancel = context.WithTimeout(ctx, reqTimeoutInterval)
	defer cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("IndexAddQueue enqueue timeout")
		default:
			return i.sched.IndexAddQueue.Enqueue(t)
		}
	}

	err := fn()
	if err != nil {
		ret.Status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		ret.Status.Reason = err.Error()
		return ret, nil
	}

	err = t.WaitToFinish()
	if err != nil {
		ret.Status.ErrorCode = commonpb.ErrorCode_UnexpectedError
		ret.Status.Reason = err.Error()
		return ret, nil
	}
	i.assignChan <- []UniqueID{t.indexBuildID}
	ret.Status.ErrorCode = commonpb.ErrorCode_Success
	ret.IndexBuildID = t.indexBuildID
	return ret, nil
}

func (i *IndexService) GetIndexStates(ctx context.Context, req *indexpb.GetIndexStatesRequest) (*indexpb.GetIndexStatesResponse, error) {
	log.Debug("get index states ...", zap.Int64s("IndexBuildIDs", req.IndexBuildIDs))
	var indexStates []*indexpb.IndexInfo
	for _, indexID := range req.IndexBuildIDs {
		indexState, err := i.metaTable.GetIndexState(indexID)
		if err != nil {
			indexState.Reason = err.Error()
		}
		indexStates = append(indexStates, indexState)
	}
	ret := &indexpb.GetIndexStatesResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		States: indexStates,
	}
	log.Debug("get index states success")
	log.Debug("get index states",
		zap.Any("index status", ret.Status),
		zap.Any("index states", ret.States))

	return ret, nil
}

func (i *IndexService) DropIndex(ctx context.Context, req *indexpb.DropIndexRequest) (*commonpb.Status, error) {
	log.Debug("IndexService", zap.Int64("Drop Index ID", req.IndexID))

	ret := &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_Success,
	}
	err := i.metaTable.MarkIndexAsDeleted(req.IndexID)
	if err != nil {
		ret.ErrorCode = commonpb.ErrorCode_UnexpectedError
		ret.Reason = err.Error()
		return ret, nil
	}

	defer func() {
		go func() {
			unissuedIndexBuildIDs := i.sched.IndexAddQueue.tryToRemoveUselessIndexAddTask(req.IndexID)
			for _, indexBuildID := range unissuedIndexBuildIDs {
				i.metaTable.DeleteIndex(indexBuildID)
			}
		}()
	}()

	log.Debug("IndexService", zap.Int64("DropIndex success by ID", req.IndexID))
	return ret, nil
}

func (i *IndexService) GetIndexFilePaths(ctx context.Context, req *indexpb.GetIndexFilePathsRequest) (*indexpb.GetIndexFilePathsResponse, error) {
	log.Debug("indexservice", zap.Int64s("get index file paths", req.IndexBuildIDs))
	var indexPaths []*indexpb.IndexFilePathInfo = nil

	for _, indexID := range req.IndexBuildIDs {
		indexPathInfo, err := i.metaTable.GetIndexFilePathInfo(indexID)
		if err != nil {
			return nil, err
		}
		indexPaths = append(indexPaths, indexPathInfo)
	}
	log.Debug("indexservice, get index file paths success")

	ret := &indexpb.GetIndexFilePathsResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
		FilePaths: indexPaths,
	}
	log.Debug("indexservice", zap.Any("index file paths", ret.FilePaths))

	return ret, nil
}

func (i *IndexService) tsLoop() {
	tsoTicker := time.NewTicker(tso.UpdateTimestampStep)
	defer tsoTicker.Stop()
	ctx, cancel := context.WithCancel(i.loopCtx)
	defer cancel()
	defer i.loopWg.Done()
	for {
		select {
		case <-tsoTicker.C:
			if err := i.idAllocator.UpdateID(); err != nil {
				log.Debug("indexservice", zap.String("failed to update id", err.Error()))
				return
			}
		case <-ctx.Done():
			// Server is closed and it should return nil.
			log.Debug("tsLoop is closed")
			return
		}
	}
}

func (i *IndexService) recycleUnusedIndexFiles() {
	ctx, cancel := context.WithCancel(i.loopCtx)

	defer cancel()
	defer i.loopWg.Done()

	timeTicker := time.NewTicker(durationInterval)
	log.Debug("IndexService start recycle unused index files loop")

	for {
		select {
		case <-ctx.Done():
			return
		case <-timeTicker.C:
			metas := i.metaTable.GetUnusedIndexFiles(recycleIndexLimit)
			for _, meta := range metas {
				if meta.indexMeta.MarkDeleted {
					unusedIndexFilePathPrefix := strconv.Itoa(int(meta.indexMeta.IndexBuildID))
					if err := i.kv.RemoveWithPrefix(unusedIndexFilePathPrefix); err != nil {
						log.Debug("IndexService", zap.String("Remove index files error", err.Error()))
					}
					i.metaTable.DeleteIndex(meta.indexMeta.IndexBuildID)
				} else {
					for j := 1; j < int(meta.indexMeta.Version); j++ {
						unusedIndexFilePathPrefix := strconv.Itoa(int(meta.indexMeta.IndexBuildID)) + "/" + strconv.Itoa(j)
						if err := i.kv.RemoveWithPrefix(unusedIndexFilePathPrefix); err != nil {
							log.Debug("IndexService", zap.String("Remove index files error", err.Error()))
						}
					}
					if err := i.metaTable.UpdateRecycleState(meta.indexMeta.IndexBuildID); err != nil {
						log.Debug("IndexService", zap.String("Remove index files error", err.Error()))
					}
				}
			}
		}
	}
}

func (i *IndexService) assignmentTasksLoop() {
	ctx, cancel := context.WithCancel(i.loopCtx)

	defer cancel()
	defer i.loopWg.Done()

	log.Debug("IndexService start assign tasks loop")

	for {
		select {
		case <-ctx.Done():
			return
		case indexBuildIDs := <-i.assignChan:
			for _, indexBuildID := range indexBuildIDs {
				meta := i.metaTable.GetIndexMeta(indexBuildID)
				log.Debug("IndexService", zap.Any("Meta", meta))
				if meta.indexMeta.State == commonpb.IndexState_Finished {
					continue
				}
				if err := i.metaTable.UpdateVersion(indexBuildID); err != nil {
					log.Debug("IndexService", zap.String("build index update version err", err.Error()))
				}
				nodeID, builderClient, nodeServerID := i.nodeClients.PeekClient()
				if builderClient == nil {
					log.Debug("IndexService has no available IndexNode")
					i.assignChan <- []UniqueID{indexBuildID}
					continue
				}
				req := &indexpb.CreateIndexRequest{
					IndexBuildID: indexBuildID,
					IndexName:    meta.indexMeta.Req.IndexName,
					IndexID:      meta.indexMeta.Req.IndexID,
					Version:      meta.indexMeta.Version + 1,
					MetaPath:     "/indexes/" + strconv.FormatInt(indexBuildID, 10),
					DataPaths:    meta.indexMeta.Req.DataPaths,
					TypeParams:   meta.indexMeta.Req.TypeParams,
					IndexParams:  meta.indexMeta.Req.IndexParams,
				}
				resp, err := builderClient.CreateIndex(ctx, req)
				if err != nil {
					log.Debug("IndexService", zap.String("build index err", err.Error()))
				}
				if err = i.metaTable.BuildIndex(indexBuildID, nodeServerID); err != nil {
					log.Debug("IndexService", zap.String("update meta table error", err.Error()))
				}
				if resp.ErrorCode != commonpb.ErrorCode_Success {
					log.Debug("IndexService", zap.String("build index err", resp.Reason))
				}
				i.nodeClients.IncPriority(nodeID, 1)
			}
		}
	}
}

func (i *IndexService) watchNodeLoop() {
	ctx, cancel := context.WithCancel(i.loopCtx)

	defer cancel()
	defer i.loopWg.Done()
	log.Debug("IndexService start watch node loop")

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-i.eventChan:
			switch event.EventType {
			case sessionutil.SessionAddEvent:
				serverID := event.Session.ServerID
				log.Debug("IndexService", zap.Any("Add IndexNode, session serverID", serverID))
			case sessionutil.SessionDelEvent:
				serverID := event.Session.ServerID
				log.Debug("IndexService", zap.Any("The IndexNode crashed with ID", serverID))
				indexBuildIDs := i.nodeTasks.getTasksByLeaseKey(serverID)
				i.assignChan <- indexBuildIDs
				i.nodeTasks.delete(serverID)
			}
		}
	}
}

func (i *IndexService) watchMetaLoop() {
	ctx, cancel := context.WithCancel(i.loopCtx)

	defer cancel()
	defer i.loopWg.Done()
	log.Debug("IndexService start watch meta loop")

	watchChan := i.metaTable.client.WatchWithPrefix("indexes")

	for {
		select {
		case <-ctx.Done():
			return
		case resp := <-watchChan:
			log.Debug("meta updated.")
			for _, event := range resp.Events {
				eventRevision := event.Kv.Version
				indexMeta := &indexpb.IndexMeta{}
				err := proto.UnmarshalText(string(event.Kv.Value), indexMeta)
				if err != nil {
					log.Debug("IndexService", zap.Any("Unmarshal error", err))
				}
				indexBuildID := indexMeta.IndexBuildID
				switch event.Type {
				case mvccpb.PUT:
					//TODO: get indexBuildID fast
					log.Debug("IndexService", zap.Any("Meta need load by IndexBuildID", indexBuildID))

					reload := i.metaTable.LoadMetaFromETCD(indexBuildID, eventRevision)
					if reload {
						i.nodeTasks.finishTask(indexBuildID)
					}
				case mvccpb.DELETE:
				}
			}
		}
	}
}

func (i *IndexService) assignTasksServerStart() error {
	sessions, _, err := i.session.GetSessions(typeutil.IndexNodeRole)
	if err != nil {
		return err
	}
	var serverIDs []int64
	for _, session := range sessions {
		serverIDs = append(serverIDs, session.ServerID)
	}
	tasks := i.metaTable.GetUnassignedTasks(serverIDs)
	for _, taskQueue := range tasks {
		i.assignChan <- taskQueue
	}

	return nil
}
