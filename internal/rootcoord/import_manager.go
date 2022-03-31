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

package rootcoord

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus/internal/kv"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/milvuspb"
	"github.com/milvus-io/milvus/internal/proto/rootcoordpb"
	"go.uber.org/zap"
)

const (
	Bucket          = "bucket"
	FailedReason    = "failed_reason"
	MaxPendingCount = 32
	delimiter       = "/"
)

// import task state
type importTaskState struct {
	stateCode    commonpb.ImportState // state code
	segments     []int64              // ID list of generated segments
	rowIDs       []int64              // ID list of auto-generated is for auto-id primary key
	rowCount     int64                // how many rows imported
	failedReason string               // failed reason
}

// importManager manager for import tasks
type importManager struct {
	ctx       context.Context    // reserved
	cancel    context.CancelFunc // reserved
	taskStore kv.MetaKv          // Persistent task info storage.

	// TODO: Make pendingTask a map to improve look up performance.
	pendingTasks []*datapb.ImportTaskInfo         // pending tasks
	workingTasks map[int64]*datapb.ImportTaskInfo // in-progress tasks
	pendingLock  sync.RWMutex                     // lock pending task list
	workingLock  sync.RWMutex                     // lock working task map
	nextTaskID   int64                            // for generating next import task ID
	lastReqID    int64                            // for generating a unique ID for import request

	callImportService func(ctx context.Context, req *datapb.ImportTask) *datapb.ImportTaskResponse
}

// newImportManager helper function to create a importManager
func newImportManager(ctx context.Context, client kv.MetaKv, importService func(ctx context.Context, req *datapb.ImportTask) *datapb.ImportTaskResponse) *importManager {
	ctx, cancel := context.WithCancel(ctx)
	mgr := &importManager{
		ctx:               ctx,
		cancel:            cancel,
		taskStore:         client,
		pendingTasks:      make([]*datapb.ImportTaskInfo, 0, MaxPendingCount), // currently task queue max size is 32
		workingTasks:      make(map[int64]*datapb.ImportTaskInfo),
		pendingLock:       sync.RWMutex{},
		workingLock:       sync.RWMutex{},
		nextTaskID:        0,
		lastReqID:         0,
		callImportService: importService,
	}

	return mgr
}

func (m *importManager) init() error {
	//  Read tasks from etcd and save them as pendingTasks or workingTasks.
	m.load()
	m.sendOutTasks()

	return nil
}

// sendOutTasks pushes all pending tasks to DataCoord, gets DataCoord response and re-add these tasks as working tasks.
func (m *importManager) sendOutTasks() error {
	m.pendingLock.Lock()
	defer m.pendingLock.Unlock()

	// Trigger Import() action to DataCoord.
	for len(m.pendingTasks) > 0 {
		task := m.pendingTasks[0]
		it := &datapb.ImportTask{
			CollectionId: task.GetCollectionId(),
			PartitionId:  task.GetPartitionId(),
			RowBased:     task.GetRowBased(),
			TaskId:       task.GetId(),
			Files:        task.GetFiles(),
			Infos: []*commonpb.KeyValuePair{
				{
					Key:   Bucket,
					Value: task.GetBucket(),
				},
			},
		}

		log.Debug("sending import task to DataCoord", zap.Int64("taskID", task.GetId()))
		// Call DataCoord.Import().
		resp := m.callImportService(m.ctx, it)
		if resp.Status.ErrorCode == commonpb.ErrorCode_UnexpectedError {
			log.Debug("import task is rejected", zap.Int64("task ID", it.GetTaskId()))
			break
		}
		task.DatanodeId = resp.GetDatanodeId()
		log.Debug("import task successfully assigned to DataNode",
			zap.Int64("task ID", it.GetTaskId()),
			zap.Int64("DataNode ID", task.GetDatanodeId()))

		// erase this task from head of pending list if the callImportService succeed
		m.pendingTasks = m.pendingTasks[1:]

		func() {
			m.workingLock.Lock()
			defer m.workingLock.Unlock()

			log.Debug("import task added as working task", zap.Int64("task ID", it.TaskId))
			task.State.StateCode = commonpb.ImportState_ImportPending
			m.workingTasks[task.GetId()] = task
			m.updateImportTask(task)
		}()
	}

	return nil
}

// genReqID generates a unique id for import request, this method has no lock, should only be called by importJob()
func (m *importManager) genReqID() int64 {
	if m.lastReqID == 0 {
		m.lastReqID = time.Now().Unix()

	} else {
		id := time.Now().Unix()
		if id == m.lastReqID {
			id++
		}
		m.lastReqID = id
	}

	return m.lastReqID
}

// importJob processes the import request, generates import tasks, sends these tasks to DataCoord, and returns
// immediately.
func (m *importManager) importJob(req *milvuspb.ImportRequest, cID int64) *milvuspb.ImportResponse {
	if req == nil || len(req.Files) == 0 {
		return &milvuspb.ImportResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UnexpectedError,
				Reason:    "import request is empty",
			},
		}
	}

	if m.callImportService == nil {
		return &milvuspb.ImportResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UnexpectedError,
				Reason:    "import service is not available",
			},
		}
	}

	resp := &milvuspb.ImportResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
		},
	}

	log.Debug("request received",
		zap.String("collection name", req.GetCollectionName()),
		zap.Int64("collection ID", cID))
	func() {
		m.pendingLock.Lock()
		defer m.pendingLock.Unlock()

		capacity := cap(m.pendingTasks)
		length := len(m.pendingTasks)

		taskCount := 1
		if req.RowBased {
			taskCount = len(req.Files)
		}

		// task queue size has a limit, return error if import request contains too many data files
		if capacity-length < taskCount {
			resp.Status = &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_IllegalArgument,
				Reason:    "Import task queue max size is " + strconv.Itoa(capacity) + ", currently there are " + strconv.Itoa(length) + " tasks is pending. Not able to execute this request with " + strconv.Itoa(taskCount) + " tasks.",
			}
			return
		}

		bucket := ""
		for _, kv := range req.Options {
			if kv.Key == Bucket {
				bucket = kv.Value
				break
			}
		}

		reqID := m.genReqID()
		// convert import request to import tasks
		if req.RowBased {
			// For row-based importing, each file makes a task.
			taskList := make([]int64, len(req.Files))
			for i := 0; i < len(req.Files); i++ {
				newTask := &datapb.ImportTaskInfo{
					Id:           m.nextTaskID,
					RequestId:    reqID,
					CollectionId: cID,
					Bucket:       bucket,
					RowBased:     req.GetRowBased(),
					Files:        []string{req.GetFiles()[i]},
					CreateTs:     time.Now().Unix(),
					State: &datapb.ImportTaskState{
						StateCode: commonpb.ImportState_ImportPending,
					},
				}
				taskList[i] = newTask.GetId()
				m.nextTaskID++
				log.Info("new task created as pending task", zap.Int64("task ID", newTask.GetId()))
				m.pendingTasks = append(m.pendingTasks, newTask)
				m.saveImportTask(newTask)
			}
			log.Info("row-based import request processed", zap.Int64("reqID", reqID), zap.Any("taskIDs", taskList))
		} else {
			// TODO: Merge duplicated code :(
			// for column-based, all files is a task
			newTask := &datapb.ImportTaskInfo{
				Id:           m.nextTaskID,
				RequestId:    reqID,
				CollectionId: cID,
				Bucket:       bucket,
				RowBased:     req.GetRowBased(),
				Files:        req.GetFiles(),
				CreateTs:     time.Now().Unix(),
				State: &datapb.ImportTaskState{
					StateCode: commonpb.ImportState_ImportPending,
				},
			}
			m.nextTaskID++
			log.Info("new task created as pending task", zap.Int64("task ID", newTask.GetId()))
			m.pendingTasks = append(m.pendingTasks, newTask)
			m.saveImportTask(newTask)
			log.Info("column-based import request processed", zap.Int64("reqID", reqID), zap.Int64("taskID", newTask.GetId()))
		}
	}()
	m.sendOutTasks()
	return resp
}

// updateTaskState updates the task's state in task store given ImportResult result, and returns the ImportTaskInfo of
// the given task.
func (m *importManager) updateTaskState(ir *rootcoordpb.ImportResult) (*datapb.ImportTaskInfo, error) {
	if ir == nil {
		return nil, errors.New("import result is nil")
	}
	log.Debug("import manager update task import result", zap.Int64("taskID", ir.GetTaskId()))

	found := false
	var v *datapb.ImportTaskInfo
	func() {
		m.workingLock.Lock()
		defer m.workingLock.Unlock()
		ok := false
		if v, ok = m.workingTasks[ir.TaskId]; ok {
			found = true
			v.State.StateCode = ir.GetState()
			v.State.Segments = ir.GetSegments()
			v.State.RowCount = ir.GetRowCount()
			for _, kv := range ir.GetInfos() {
				if kv.GetKey() == FailedReason {
					v.State.ErrorMessage = kv.GetValue()
					break
				}
			}
			// Update task in task store.
			m.updateImportTask(v)
		}
		m.updateImportTask(v)
	}()

	if !found {
		log.Debug("import manager update task import result failed", zap.Int64("taskID", ir.GetTaskId()))
		return nil, errors.New("failed to update import task, ID not found: " + strconv.FormatInt(ir.TaskId, 10))
	}
	return v, nil
}

// getTaskState looks for task with the given ID and returns its import state.
func (m *importManager) getTaskState(tID int64) *milvuspb.GetImportStateResponse {
	resp := &milvuspb.GetImportStateResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UnexpectedError,
			Reason:    "import task id doesn't exist",
		},
	}

	log.Debug("getting import task state", zap.Int64("taskID", tID))
	found := false
	func() {
		m.pendingLock.Lock()
		defer m.pendingLock.Unlock()
		for i := 0; i < len(m.pendingTasks); i++ {
			if tID == m.pendingTasks[i].Id {
				resp.Status = &commonpb.Status{
					ErrorCode: commonpb.ErrorCode_Success,
				}
				resp.State = commonpb.ImportState_ImportPending
				found = true
				break
			}
		}
	}()
	if found {
		return resp
	}

	func() {
		m.workingLock.Lock()
		defer m.workingLock.Unlock()
		if v, ok := m.workingTasks[tID]; ok {
			found = true
			resp.Status = &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_Success,
			}
			resp.State = v.GetState().GetStateCode()
			resp.RowCount = v.GetState().GetRowCount()
			resp.IdList = v.GetState().GetRowIds()
			resp.Infos = append(resp.Infos, &commonpb.KeyValuePair{
				Key:   FailedReason,
				Value: v.GetState().GetErrorMessage(),
			})
		}
	}()
	if found {
		return resp
	}
	log.Debug("get import task state failed", zap.Int64("taskID", tID))
	return resp
}

// load Loads task info from task store when RootCoord (re)starts.
func (m *importManager) load() error {
	log.Info("import manager starts loading from Etcd")
	_, v, err := m.taskStore.LoadWithPrefix(Params.RootCoordCfg.ImportTaskSubPath)
	if err != nil {
		log.Error("import manager failed to load from Etcd", zap.Error(err))
		return err
	}
	m.workingLock.Lock()
	defer m.workingLock.Unlock()
	m.pendingLock.Lock()
	defer m.pendingLock.Unlock()
	for i := range v {
		ti := &datapb.ImportTaskInfo{}
		if err := proto.Unmarshal([]byte(v[i]), ti); err != nil {
			log.Error("failed to unmarshal proto", zap.String("taskInfo", v[i]), zap.Error(err))
			// Ignore bad protos.
			continue
		}
		// Put tasks back to pending or working task list, given their import states.
		if ti.GetState().GetStateCode() == commonpb.ImportState_ImportPending {
			log.Info("task has been reloaded as a pending task", zap.Int64("TaskID", ti.GetId()))
			m.pendingTasks = append(m.pendingTasks, ti)
		} else {
			log.Info("task has been reloaded as a working tasks", zap.Int64("TaskID", ti.GetId()))
			m.workingTasks[ti.GetId()] = ti
		}
	}
	return nil
}

// saveImportTask signs a lease and saves import task info into Etcd with this lease.
func (m *importManager) saveImportTask(task *datapb.ImportTaskInfo) error {
	log.Debug("saving import task to Etcd", zap.Int64("Task ID", task.GetId()))
	// TODO: Change default lease time and read it into config, once we figure out a proper value.
	// Sign a lease.
	leaseID, err := m.taskStore.Grant(10800) /*3 hours*/
	if err != nil {
		log.Error("failed to grant lease from Etcd for data import",
			zap.Int64("Task ID", task.GetId()),
			zap.Error(err))
		return err
	}
	log.Debug("lease granted for task", zap.Int64("Task ID", task.GetId()))
	var taskInfo []byte
	if taskInfo, err = proto.Marshal(task); err != nil {
		log.Error("failed to marshall task proto", zap.Int64("Task ID", task.GetId()), zap.Error(err))
		return err
	} else if err = m.taskStore.SaveWithLease(BuildImportTaskKey(task.GetId()), string(taskInfo), leaseID); err != nil {
		log.Error("failed to save import task info into Etcd",
			zap.Int64("task ID", task.GetId()),
			zap.Error(err))
		return err
	}
	log.Debug("task info successfully saved", zap.Int64("Task ID", task.GetId()))
	return nil
}

// updateImportTask updates the task info in Etcd according to task ID. It won't change the lease on the key.
func (m *importManager) updateImportTask(ti *datapb.ImportTaskInfo) error {
	log.Debug("updating import task info in Etcd", zap.Int64("Task ID", ti.GetId()))
	if taskInfo, err := proto.Marshal(ti); err != nil {
		log.Error("failed to marshall task info proto", zap.Int64("Task ID", ti.GetId()), zap.Error(err))
		return err
	} else if err = m.taskStore.SaveWithIgnoreLease(BuildImportTaskKey(ti.GetId()), string(taskInfo)); err != nil {
		log.Error("failed to update import task info info in Etcd", zap.Int64("Task ID", ti.GetId()), zap.Error(err))
		return err
	}
	log.Debug("task info successfully updated in Etcd", zap.Int64("Task ID", ti.GetId()))
	return nil
}

// bringSegmentsOnline brings the segments online so that data in these segments become searchable.
func (m *importManager) bringSegmentsOnline(ti *datapb.ImportTaskInfo) {
	log.Info("Bringing import tasks segments online!", zap.Int64("Task ID", ti.GetId()))
	// TODO: Implement it.
}

// BuildImportTaskKey constructs and returns an Etcd key with given task ID.
func BuildImportTaskKey(taskID int64) string {
	return fmt.Sprintf("%s%s%d", Params.RootCoordCfg.ImportTaskSubPath, delimiter, taskID)
}
