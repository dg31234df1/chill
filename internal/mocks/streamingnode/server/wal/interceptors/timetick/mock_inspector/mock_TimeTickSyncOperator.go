// Code generated by mockery v2.46.0. DO NOT EDIT.

package mock_inspector

import (
	context "context"

	inspector "github.com/milvus-io/milvus/internal/streamingnode/server/wal/interceptors/timetick/inspector"
	mock "github.com/stretchr/testify/mock"

	types "github.com/milvus-io/milvus/pkg/streaming/util/types"
)

// MockTimeTickSyncOperator is an autogenerated mock type for the TimeTickSyncOperator type
type MockTimeTickSyncOperator struct {
	mock.Mock
}

type MockTimeTickSyncOperator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTimeTickSyncOperator) EXPECT() *MockTimeTickSyncOperator_Expecter {
	return &MockTimeTickSyncOperator_Expecter{mock: &_m.Mock}
}

// Channel provides a mock function with given fields:
func (_m *MockTimeTickSyncOperator) Channel() types.PChannelInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Channel")
	}

	var r0 types.PChannelInfo
	if rf, ok := ret.Get(0).(func() types.PChannelInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.PChannelInfo)
	}

	return r0
}

// MockTimeTickSyncOperator_Channel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Channel'
type MockTimeTickSyncOperator_Channel_Call struct {
	*mock.Call
}

// Channel is a helper method to define mock.On call
func (_e *MockTimeTickSyncOperator_Expecter) Channel() *MockTimeTickSyncOperator_Channel_Call {
	return &MockTimeTickSyncOperator_Channel_Call{Call: _e.mock.On("Channel")}
}

func (_c *MockTimeTickSyncOperator_Channel_Call) Run(run func()) *MockTimeTickSyncOperator_Channel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTimeTickSyncOperator_Channel_Call) Return(_a0 types.PChannelInfo) *MockTimeTickSyncOperator_Channel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTimeTickSyncOperator_Channel_Call) RunAndReturn(run func() types.PChannelInfo) *MockTimeTickSyncOperator_Channel_Call {
	_c.Call.Return(run)
	return _c
}

// Sync provides a mock function with given fields: ctx
func (_m *MockTimeTickSyncOperator) Sync(ctx context.Context) {
	_m.Called(ctx)
}

// MockTimeTickSyncOperator_Sync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sync'
type MockTimeTickSyncOperator_Sync_Call struct {
	*mock.Call
}

// Sync is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockTimeTickSyncOperator_Expecter) Sync(ctx interface{}) *MockTimeTickSyncOperator_Sync_Call {
	return &MockTimeTickSyncOperator_Sync_Call{Call: _e.mock.On("Sync", ctx)}
}

func (_c *MockTimeTickSyncOperator_Sync_Call) Run(run func(ctx context.Context)) *MockTimeTickSyncOperator_Sync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockTimeTickSyncOperator_Sync_Call) Return() *MockTimeTickSyncOperator_Sync_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTimeTickSyncOperator_Sync_Call) RunAndReturn(run func(context.Context)) *MockTimeTickSyncOperator_Sync_Call {
	_c.Call.Return(run)
	return _c
}

// TimeTickNotifier provides a mock function with given fields:
func (_m *MockTimeTickSyncOperator) TimeTickNotifier() *inspector.TimeTickNotifier {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TimeTickNotifier")
	}

	var r0 *inspector.TimeTickNotifier
	if rf, ok := ret.Get(0).(func() *inspector.TimeTickNotifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*inspector.TimeTickNotifier)
		}
	}

	return r0
}

// MockTimeTickSyncOperator_TimeTickNotifier_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TimeTickNotifier'
type MockTimeTickSyncOperator_TimeTickNotifier_Call struct {
	*mock.Call
}

// TimeTickNotifier is a helper method to define mock.On call
func (_e *MockTimeTickSyncOperator_Expecter) TimeTickNotifier() *MockTimeTickSyncOperator_TimeTickNotifier_Call {
	return &MockTimeTickSyncOperator_TimeTickNotifier_Call{Call: _e.mock.On("TimeTickNotifier")}
}

func (_c *MockTimeTickSyncOperator_TimeTickNotifier_Call) Run(run func()) *MockTimeTickSyncOperator_TimeTickNotifier_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTimeTickSyncOperator_TimeTickNotifier_Call) Return(_a0 *inspector.TimeTickNotifier) *MockTimeTickSyncOperator_TimeTickNotifier_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTimeTickSyncOperator_TimeTickNotifier_Call) RunAndReturn(run func() *inspector.TimeTickNotifier) *MockTimeTickSyncOperator_TimeTickNotifier_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTimeTickSyncOperator creates a new instance of MockTimeTickSyncOperator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTimeTickSyncOperator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTimeTickSyncOperator {
	mock := &MockTimeTickSyncOperator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
