// Code generated by mockery v2.30.1. DO NOT EDIT.

package syncmgr

import (
	context "context"

	msgpb "github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
	mock "github.com/stretchr/testify/mock"
)

// MockTask is an autogenerated mock type for the Task type
type MockTask struct {
	mock.Mock
}

type MockTask_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTask) EXPECT() *MockTask_Expecter {
	return &MockTask_Expecter{mock: &_m.Mock}
}

// ChannelName provides a mock function with given fields:
func (_m *MockTask) ChannelName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTask_ChannelName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChannelName'
type MockTask_ChannelName_Call struct {
	*mock.Call
}

// ChannelName is a helper method to define mock.On call
func (_e *MockTask_Expecter) ChannelName() *MockTask_ChannelName_Call {
	return &MockTask_ChannelName_Call{Call: _e.mock.On("ChannelName")}
}

func (_c *MockTask_ChannelName_Call) Run(run func()) *MockTask_ChannelName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTask_ChannelName_Call) Return(_a0 string) *MockTask_ChannelName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTask_ChannelName_Call) RunAndReturn(run func() string) *MockTask_ChannelName_Call {
	_c.Call.Return(run)
	return _c
}

// Checkpoint provides a mock function with given fields:
func (_m *MockTask) Checkpoint() *msgpb.MsgPosition {
	ret := _m.Called()

	var r0 *msgpb.MsgPosition
	if rf, ok := ret.Get(0).(func() *msgpb.MsgPosition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msgpb.MsgPosition)
		}
	}

	return r0
}

// MockTask_Checkpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Checkpoint'
type MockTask_Checkpoint_Call struct {
	*mock.Call
}

// Checkpoint is a helper method to define mock.On call
func (_e *MockTask_Expecter) Checkpoint() *MockTask_Checkpoint_Call {
	return &MockTask_Checkpoint_Call{Call: _e.mock.On("Checkpoint")}
}

func (_c *MockTask_Checkpoint_Call) Run(run func()) *MockTask_Checkpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTask_Checkpoint_Call) Return(_a0 *msgpb.MsgPosition) *MockTask_Checkpoint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTask_Checkpoint_Call) RunAndReturn(run func() *msgpb.MsgPosition) *MockTask_Checkpoint_Call {
	_c.Call.Return(run)
	return _c
}

// HandleError provides a mock function with given fields: _a0
func (_m *MockTask) HandleError(_a0 error) {
	_m.Called(_a0)
}

// MockTask_HandleError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleError'
type MockTask_HandleError_Call struct {
	*mock.Call
}

// HandleError is a helper method to define mock.On call
//   - _a0 error
func (_e *MockTask_Expecter) HandleError(_a0 interface{}) *MockTask_HandleError_Call {
	return &MockTask_HandleError_Call{Call: _e.mock.On("HandleError", _a0)}
}

func (_c *MockTask_HandleError_Call) Run(run func(_a0 error)) *MockTask_HandleError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockTask_HandleError_Call) Return() *MockTask_HandleError_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTask_HandleError_Call) RunAndReturn(run func(error)) *MockTask_HandleError_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: _a0
func (_m *MockTask) Run(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTask_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockTask_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockTask_Expecter) Run(_a0 interface{}) *MockTask_Run_Call {
	return &MockTask_Run_Call{Call: _e.mock.On("Run", _a0)}
}

func (_c *MockTask_Run_Call) Run(run func(_a0 context.Context)) *MockTask_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockTask_Run_Call) Return(_a0 error) *MockTask_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTask_Run_Call) RunAndReturn(run func(context.Context) error) *MockTask_Run_Call {
	_c.Call.Return(run)
	return _c
}

// SegmentID provides a mock function with given fields:
func (_m *MockTask) SegmentID() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockTask_SegmentID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SegmentID'
type MockTask_SegmentID_Call struct {
	*mock.Call
}

// SegmentID is a helper method to define mock.On call
func (_e *MockTask_Expecter) SegmentID() *MockTask_SegmentID_Call {
	return &MockTask_SegmentID_Call{Call: _e.mock.On("SegmentID")}
}

func (_c *MockTask_SegmentID_Call) Run(run func()) *MockTask_SegmentID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTask_SegmentID_Call) Return(_a0 int64) *MockTask_SegmentID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTask_SegmentID_Call) RunAndReturn(run func() int64) *MockTask_SegmentID_Call {
	_c.Call.Return(run)
	return _c
}

// StartPosition provides a mock function with given fields:
func (_m *MockTask) StartPosition() *msgpb.MsgPosition {
	ret := _m.Called()

	var r0 *msgpb.MsgPosition
	if rf, ok := ret.Get(0).(func() *msgpb.MsgPosition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msgpb.MsgPosition)
		}
	}

	return r0
}

// MockTask_StartPosition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartPosition'
type MockTask_StartPosition_Call struct {
	*mock.Call
}

// StartPosition is a helper method to define mock.On call
func (_e *MockTask_Expecter) StartPosition() *MockTask_StartPosition_Call {
	return &MockTask_StartPosition_Call{Call: _e.mock.On("StartPosition")}
}

func (_c *MockTask_StartPosition_Call) Run(run func()) *MockTask_StartPosition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTask_StartPosition_Call) Return(_a0 *msgpb.MsgPosition) *MockTask_StartPosition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTask_StartPosition_Call) RunAndReturn(run func() *msgpb.MsgPosition) *MockTask_StartPosition_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTask creates a new instance of MockTask. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTask(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTask {
	mock := &MockTask{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
