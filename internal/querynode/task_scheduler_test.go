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

package querynode

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"

	"github.com/milvus-io/milvus/internal/util/paramtable"
	"github.com/stretchr/testify/assert"
)

type mockTask struct {
	baseTask
	preExecuteError bool
	executeError    bool
	timestamp       Timestamp
}

func (m *mockTask) Timestamp() Timestamp {
	return m.timestamp
}

func (m *mockTask) OnEnqueue() error {
	return nil
}

func (m *mockTask) PreExecute(ctx context.Context) error {
	if m.preExecuteError {
		return errors.New("test error")
	}
	return nil
}

func (m *mockTask) Execute(ctx context.Context) error {
	if m.executeError {
		return errors.New("test error")
	}
	return nil
}

func (m *mockTask) PostExecute(ctx context.Context) error {
	return nil
}

var _ readTask = (*mockReadTask)(nil)

type mockReadTask struct {
	mockTask
	cpuUsage     int32
	maxCPU       int32
	collectionID UniqueID
	ready        bool
	canMerge     bool
	timeout      bool
	timeoutError error
	step         TaskStep
	readyError   error
}

func (m *mockReadTask) GetCollectionID() UniqueID {
	return m.collectionID
}

func (m *mockReadTask) Ready() (bool, error) {
	return m.ready, m.readyError
}

func (m *mockReadTask) Merge(o readTask) {

}

func (m *mockReadTask) CPUUsage() int32 {
	return m.cpuUsage
}

func (m *mockReadTask) Timeout() bool {
	return m.timeout
}

func (m *mockReadTask) TimeoutError() error {
	return m.timeoutError
}

func (m *mockReadTask) SetMaxCPUUsage(cpu int32) {
	m.maxCPU = cpu
}

func (m *mockReadTask) SetStep(step TaskStep) {
	m.step = step
}

func (m *mockReadTask) CanMergeWith(o readTask) bool {
	return m.canMerge
}

func TestTaskScheduler(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tSafe := newTSafeReplica()

	ts := newTaskScheduler(ctx, tSafe)
	ts.Start()

	task := &mockTask{
		baseTask: baseTask{
			ctx:  ctx,
			done: make(chan error, 1024),
		},
		preExecuteError: true,
		executeError:    false,
	}
	ts.processTask(task, ts.queue)

	task.preExecuteError = false
	task.executeError = true
	ts.processTask(task, ts.queue)

	ts.Close()
}

func TestTaskScheduler_tryEvictUnsolvedReadTask(t *testing.T) {
	t.Run("evict canceled task", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		tSafe := newTSafeReplica()

		ts := newTaskScheduler(ctx, tSafe)

		taskCanceled := &mockReadTask{
			mockTask: mockTask{
				baseTask: baseTask{
					ctx:  ctx,
					done: make(chan error, 1024),
				},
			},
			timeout:      true,
			timeoutError: context.Canceled,
		}
		taskNormal := &mockReadTask{
			mockTask: mockTask{
				baseTask: baseTask{
					ctx:  context.Background(),
					done: make(chan error, 1024),
				},
			},
		}

		ts.unsolvedReadTasks.PushBack(taskNormal)
		ts.unsolvedReadTasks.PushBack(taskCanceled)

		// set max len to 2
		paramtable.Get().Save(Params.QueryNodeCfg.MaxUnsolvedQueueSize.Key, "2")
		ts.tryEvictUnsolvedReadTask(1)
		paramtable.Get().Reset(Params.QueryNodeCfg.MaxUnsolvedQueueSize.Key)

		err := <-taskCanceled.done
		assert.ErrorIs(t, err, context.Canceled)

		select {
		case <-taskNormal.done:
			t.Fail()
		default:
		}

		assert.Equal(t, 1, ts.unsolvedReadTasks.Len())
	})
}

func TestTaskScheduler_executeReadTasks(t *testing.T) {
	t.Run("execute canceled task", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		tSafe := newTSafeReplica()

		ts := newTaskScheduler(ctx, tSafe)
		ts.Start()
		defer ts.Close()

		taskCanceled := &mockReadTask{
			mockTask: mockTask{
				baseTask: baseTask{
					ctx:  ctx,
					done: make(chan error, 1024),
				},
			},
			timeout:      true,
			timeoutError: context.Canceled,
		}

		ts.executeReadTaskChan <- taskCanceled

		err := <-taskCanceled.done
		assert.ErrorIs(t, err, context.Canceled)
	})
}
