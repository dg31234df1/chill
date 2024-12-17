// Code generated by mockery v2.46.0. DO NOT EDIT.

package syncmgr

import (
	context "context"

	conc "github.com/milvus-io/milvus/pkg/util/conc"

	mock "github.com/stretchr/testify/mock"
)

// MockSyncManager is an autogenerated mock type for the SyncManager type
type MockSyncManager struct {
	mock.Mock
}

type MockSyncManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSyncManager) EXPECT() *MockSyncManager_Expecter {
	return &MockSyncManager_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockSyncManager) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSyncManager_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSyncManager_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockSyncManager_Expecter) Close() *MockSyncManager_Close_Call {
	return &MockSyncManager_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockSyncManager_Close_Call) Run(run func()) *MockSyncManager_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSyncManager_Close_Call) Return(_a0 error) *MockSyncManager_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSyncManager_Close_Call) RunAndReturn(run func() error) *MockSyncManager_Close_Call {
	_c.Call.Return(run)
	return _c
}

// SyncData provides a mock function with given fields: ctx, task, callbacks
func (_m *MockSyncManager) SyncData(ctx context.Context, task Task, callbacks ...func(error) error) (*conc.Future[struct{}], error) {
	_va := make([]interface{}, len(callbacks))
	for _i := range callbacks {
		_va[_i] = callbacks[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, task)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SyncData")
	}

	var r0 *conc.Future[struct{}]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, Task, ...func(error) error) (*conc.Future[struct{}], error)); ok {
		return rf(ctx, task, callbacks...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, Task, ...func(error) error) *conc.Future[struct{}]); ok {
		r0 = rf(ctx, task, callbacks...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*conc.Future[struct{}])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, Task, ...func(error) error) error); ok {
		r1 = rf(ctx, task, callbacks...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSyncManager_SyncData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncData'
type MockSyncManager_SyncData_Call struct {
	*mock.Call
}

// SyncData is a helper method to define mock.On call
//   - ctx context.Context
//   - task Task
//   - callbacks ...func(error) error
func (_e *MockSyncManager_Expecter) SyncData(ctx interface{}, task interface{}, callbacks ...interface{}) *MockSyncManager_SyncData_Call {
	return &MockSyncManager_SyncData_Call{Call: _e.mock.On("SyncData",
		append([]interface{}{ctx, task}, callbacks...)...)}
}

func (_c *MockSyncManager_SyncData_Call) Run(run func(ctx context.Context, task Task, callbacks ...func(error) error)) *MockSyncManager_SyncData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(error) error, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(error) error)
			}
		}
		run(args[0].(context.Context), args[1].(Task), variadicArgs...)
	})
	return _c
}

func (_c *MockSyncManager_SyncData_Call) Return(_a0 *conc.Future[struct{}], _a1 error) *MockSyncManager_SyncData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSyncManager_SyncData_Call) RunAndReturn(run func(context.Context, Task, ...func(error) error) (*conc.Future[struct{}], error)) *MockSyncManager_SyncData_Call {
	_c.Call.Return(run)
	return _c
}

// TaskStatsJSON provides a mock function with given fields:
func (_m *MockSyncManager) TaskStatsJSON() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TaskStatsJSON")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSyncManager_TaskStatsJSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TaskStatsJSON'
type MockSyncManager_TaskStatsJSON_Call struct {
	*mock.Call
}

// TaskStatsJSON is a helper method to define mock.On call
func (_e *MockSyncManager_Expecter) TaskStatsJSON() *MockSyncManager_TaskStatsJSON_Call {
	return &MockSyncManager_TaskStatsJSON_Call{Call: _e.mock.On("TaskStatsJSON")}
}

func (_c *MockSyncManager_TaskStatsJSON_Call) Run(run func()) *MockSyncManager_TaskStatsJSON_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSyncManager_TaskStatsJSON_Call) Return(_a0 string) *MockSyncManager_TaskStatsJSON_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSyncManager_TaskStatsJSON_Call) RunAndReturn(run func() string) *MockSyncManager_TaskStatsJSON_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSyncManager creates a new instance of MockSyncManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSyncManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSyncManager {
	mock := &MockSyncManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
