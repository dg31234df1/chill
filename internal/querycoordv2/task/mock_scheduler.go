// Code generated by mockery v2.32.4. DO NOT EDIT.

package task

import mock "github.com/stretchr/testify/mock"

// MockScheduler is an autogenerated mock type for the Scheduler type
type MockScheduler struct {
	mock.Mock
}

type MockScheduler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockScheduler) EXPECT() *MockScheduler_Expecter {
	return &MockScheduler_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: task
func (_m *MockScheduler) Add(task Task) error {
	ret := _m.Called(task)

	var r0 error
	if rf, ok := ret.Get(0).(func(Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockScheduler_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockScheduler_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//  - task Task
func (_e *MockScheduler_Expecter) Add(task interface{}) *MockScheduler_Add_Call {
	return &MockScheduler_Add_Call{Call: _e.mock.On("Add", task)}
}

func (_c *MockScheduler_Add_Call) Run(run func(task Task)) *MockScheduler_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Task))
	})
	return _c
}

func (_c *MockScheduler_Add_Call) Return(_a0 error) *MockScheduler_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScheduler_Add_Call) RunAndReturn(run func(Task) error) *MockScheduler_Add_Call {
	_c.Call.Return(run)
	return _c
}

// AddExecutor provides a mock function with given fields: nodeID
func (_m *MockScheduler) AddExecutor(nodeID int64) {
	_m.Called(nodeID)
}

// MockScheduler_AddExecutor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddExecutor'
type MockScheduler_AddExecutor_Call struct {
	*mock.Call
}

// AddExecutor is a helper method to define mock.On call
//  - nodeID int64
func (_e *MockScheduler_Expecter) AddExecutor(nodeID interface{}) *MockScheduler_AddExecutor_Call {
	return &MockScheduler_AddExecutor_Call{Call: _e.mock.On("AddExecutor", nodeID)}
}

func (_c *MockScheduler_AddExecutor_Call) Run(run func(nodeID int64)) *MockScheduler_AddExecutor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockScheduler_AddExecutor_Call) Return() *MockScheduler_AddExecutor_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockScheduler_AddExecutor_Call) RunAndReturn(run func(int64)) *MockScheduler_AddExecutor_Call {
	_c.Call.Return(run)
	return _c
}

// Dispatch provides a mock function with given fields: node
func (_m *MockScheduler) Dispatch(node int64) {
	_m.Called(node)
}

// MockScheduler_Dispatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dispatch'
type MockScheduler_Dispatch_Call struct {
	*mock.Call
}

// Dispatch is a helper method to define mock.On call
//  - node int64
func (_e *MockScheduler_Expecter) Dispatch(node interface{}) *MockScheduler_Dispatch_Call {
	return &MockScheduler_Dispatch_Call{Call: _e.mock.On("Dispatch", node)}
}

func (_c *MockScheduler_Dispatch_Call) Run(run func(node int64)) *MockScheduler_Dispatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockScheduler_Dispatch_Call) Return() *MockScheduler_Dispatch_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockScheduler_Dispatch_Call) RunAndReturn(run func(int64)) *MockScheduler_Dispatch_Call {
	_c.Call.Return(run)
	return _c
}

// GetChannelTaskNum provides a mock function with given fields:
func (_m *MockScheduler) GetChannelTaskNum() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockScheduler_GetChannelTaskNum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChannelTaskNum'
type MockScheduler_GetChannelTaskNum_Call struct {
	*mock.Call
}

// GetChannelTaskNum is a helper method to define mock.On call
func (_e *MockScheduler_Expecter) GetChannelTaskNum() *MockScheduler_GetChannelTaskNum_Call {
	return &MockScheduler_GetChannelTaskNum_Call{Call: _e.mock.On("GetChannelTaskNum")}
}

func (_c *MockScheduler_GetChannelTaskNum_Call) Run(run func()) *MockScheduler_GetChannelTaskNum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScheduler_GetChannelTaskNum_Call) Return(_a0 int) *MockScheduler_GetChannelTaskNum_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScheduler_GetChannelTaskNum_Call) RunAndReturn(run func() int) *MockScheduler_GetChannelTaskNum_Call {
	_c.Call.Return(run)
	return _c
}

// GetNodeChannelDelta provides a mock function with given fields: nodeID
func (_m *MockScheduler) GetNodeChannelDelta(nodeID int64) int {
	ret := _m.Called(nodeID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int64) int); ok {
		r0 = rf(nodeID)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockScheduler_GetNodeChannelDelta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNodeChannelDelta'
type MockScheduler_GetNodeChannelDelta_Call struct {
	*mock.Call
}

// GetNodeChannelDelta is a helper method to define mock.On call
//  - nodeID int64
func (_e *MockScheduler_Expecter) GetNodeChannelDelta(nodeID interface{}) *MockScheduler_GetNodeChannelDelta_Call {
	return &MockScheduler_GetNodeChannelDelta_Call{Call: _e.mock.On("GetNodeChannelDelta", nodeID)}
}

func (_c *MockScheduler_GetNodeChannelDelta_Call) Run(run func(nodeID int64)) *MockScheduler_GetNodeChannelDelta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockScheduler_GetNodeChannelDelta_Call) Return(_a0 int) *MockScheduler_GetNodeChannelDelta_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScheduler_GetNodeChannelDelta_Call) RunAndReturn(run func(int64) int) *MockScheduler_GetNodeChannelDelta_Call {
	_c.Call.Return(run)
	return _c
}

// GetNodeSegmentDelta provides a mock function with given fields: nodeID
func (_m *MockScheduler) GetNodeSegmentDelta(nodeID int64) int {
	ret := _m.Called(nodeID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int64) int); ok {
		r0 = rf(nodeID)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockScheduler_GetNodeSegmentDelta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNodeSegmentDelta'
type MockScheduler_GetNodeSegmentDelta_Call struct {
	*mock.Call
}

// GetNodeSegmentDelta is a helper method to define mock.On call
//  - nodeID int64
func (_e *MockScheduler_Expecter) GetNodeSegmentDelta(nodeID interface{}) *MockScheduler_GetNodeSegmentDelta_Call {
	return &MockScheduler_GetNodeSegmentDelta_Call{Call: _e.mock.On("GetNodeSegmentDelta", nodeID)}
}

func (_c *MockScheduler_GetNodeSegmentDelta_Call) Run(run func(nodeID int64)) *MockScheduler_GetNodeSegmentDelta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockScheduler_GetNodeSegmentDelta_Call) Return(_a0 int) *MockScheduler_GetNodeSegmentDelta_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScheduler_GetNodeSegmentDelta_Call) RunAndReturn(run func(int64) int) *MockScheduler_GetNodeSegmentDelta_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegmentTaskNum provides a mock function with given fields:
func (_m *MockScheduler) GetSegmentTaskNum() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockScheduler_GetSegmentTaskNum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentTaskNum'
type MockScheduler_GetSegmentTaskNum_Call struct {
	*mock.Call
}

// GetSegmentTaskNum is a helper method to define mock.On call
func (_e *MockScheduler_Expecter) GetSegmentTaskNum() *MockScheduler_GetSegmentTaskNum_Call {
	return &MockScheduler_GetSegmentTaskNum_Call{Call: _e.mock.On("GetSegmentTaskNum")}
}

func (_c *MockScheduler_GetSegmentTaskNum_Call) Run(run func()) *MockScheduler_GetSegmentTaskNum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScheduler_GetSegmentTaskNum_Call) Return(_a0 int) *MockScheduler_GetSegmentTaskNum_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScheduler_GetSegmentTaskNum_Call) RunAndReturn(run func() int) *MockScheduler_GetSegmentTaskNum_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveByNode provides a mock function with given fields: node
func (_m *MockScheduler) RemoveByNode(node int64) {
	_m.Called(node)
}

// MockScheduler_RemoveByNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveByNode'
type MockScheduler_RemoveByNode_Call struct {
	*mock.Call
}

// RemoveByNode is a helper method to define mock.On call
//  - node int64
func (_e *MockScheduler_Expecter) RemoveByNode(node interface{}) *MockScheduler_RemoveByNode_Call {
	return &MockScheduler_RemoveByNode_Call{Call: _e.mock.On("RemoveByNode", node)}
}

func (_c *MockScheduler_RemoveByNode_Call) Run(run func(node int64)) *MockScheduler_RemoveByNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockScheduler_RemoveByNode_Call) Return() *MockScheduler_RemoveByNode_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockScheduler_RemoveByNode_Call) RunAndReturn(run func(int64)) *MockScheduler_RemoveByNode_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveExecutor provides a mock function with given fields: nodeID
func (_m *MockScheduler) RemoveExecutor(nodeID int64) {
	_m.Called(nodeID)
}

// MockScheduler_RemoveExecutor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveExecutor'
type MockScheduler_RemoveExecutor_Call struct {
	*mock.Call
}

// RemoveExecutor is a helper method to define mock.On call
//  - nodeID int64
func (_e *MockScheduler_Expecter) RemoveExecutor(nodeID interface{}) *MockScheduler_RemoveExecutor_Call {
	return &MockScheduler_RemoveExecutor_Call{Call: _e.mock.On("RemoveExecutor", nodeID)}
}

func (_c *MockScheduler_RemoveExecutor_Call) Run(run func(nodeID int64)) *MockScheduler_RemoveExecutor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockScheduler_RemoveExecutor_Call) Return() *MockScheduler_RemoveExecutor_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockScheduler_RemoveExecutor_Call) RunAndReturn(run func(int64)) *MockScheduler_RemoveExecutor_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockScheduler) Start() {
	_m.Called()
}

// MockScheduler_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockScheduler_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockScheduler_Expecter) Start() *MockScheduler_Start_Call {
	return &MockScheduler_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockScheduler_Start_Call) Run(run func()) *MockScheduler_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScheduler_Start_Call) Return() *MockScheduler_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockScheduler_Start_Call) RunAndReturn(run func()) *MockScheduler_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockScheduler) Stop() {
	_m.Called()
}

// MockScheduler_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockScheduler_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockScheduler_Expecter) Stop() *MockScheduler_Stop_Call {
	return &MockScheduler_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockScheduler_Stop_Call) Run(run func()) *MockScheduler_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScheduler_Stop_Call) Return() *MockScheduler_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockScheduler_Stop_Call) RunAndReturn(run func()) *MockScheduler_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockScheduler creates a new instance of MockScheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockScheduler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockScheduler {
	mock := &MockScheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
