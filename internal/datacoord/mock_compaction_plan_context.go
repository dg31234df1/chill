// Code generated by mockery v2.32.4. DO NOT EDIT.

package datacoord

import (
	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"
)

// MockCompactionPlanContext is an autogenerated mock type for the compactionPlanContext type
type MockCompactionPlanContext struct {
	mock.Mock
}

type MockCompactionPlanContext_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCompactionPlanContext) EXPECT() *MockCompactionPlanContext_Expecter {
	return &MockCompactionPlanContext_Expecter{mock: &_m.Mock}
}

// execCompactionPlan provides a mock function with given fields: signal, plan
func (_m *MockCompactionPlanContext) execCompactionPlan(signal *compactionSignal, plan *datapb.CompactionPlan) error {
	ret := _m.Called(signal, plan)

	var r0 error
	if rf, ok := ret.Get(0).(func(*compactionSignal, *datapb.CompactionPlan) error); ok {
		r0 = rf(signal, plan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCompactionPlanContext_execCompactionPlan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'execCompactionPlan'
type MockCompactionPlanContext_execCompactionPlan_Call struct {
	*mock.Call
}

// execCompactionPlan is a helper method to define mock.On call
//   - signal *compactionSignal
//   - plan *datapb.CompactionPlan
func (_e *MockCompactionPlanContext_Expecter) execCompactionPlan(signal interface{}, plan interface{}) *MockCompactionPlanContext_execCompactionPlan_Call {
	return &MockCompactionPlanContext_execCompactionPlan_Call{Call: _e.mock.On("execCompactionPlan", signal, plan)}
}

func (_c *MockCompactionPlanContext_execCompactionPlan_Call) Run(run func(signal *compactionSignal, plan *datapb.CompactionPlan)) *MockCompactionPlanContext_execCompactionPlan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*compactionSignal), args[1].(*datapb.CompactionPlan))
	})
	return _c
}

func (_c *MockCompactionPlanContext_execCompactionPlan_Call) Return(_a0 error) *MockCompactionPlanContext_execCompactionPlan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_execCompactionPlan_Call) RunAndReturn(run func(*compactionSignal, *datapb.CompactionPlan) error) *MockCompactionPlanContext_execCompactionPlan_Call {
	_c.Call.Return(run)
	return _c
}

// getCompaction provides a mock function with given fields: planID
func (_m *MockCompactionPlanContext) getCompaction(planID int64) *compactionTask {
	ret := _m.Called(planID)

	var r0 *compactionTask
	if rf, ok := ret.Get(0).(func(int64) *compactionTask); ok {
		r0 = rf(planID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compactionTask)
		}
	}

	return r0
}

// MockCompactionPlanContext_getCompaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getCompaction'
type MockCompactionPlanContext_getCompaction_Call struct {
	*mock.Call
}

// getCompaction is a helper method to define mock.On call
//   - planID int64
func (_e *MockCompactionPlanContext_Expecter) getCompaction(planID interface{}) *MockCompactionPlanContext_getCompaction_Call {
	return &MockCompactionPlanContext_getCompaction_Call{Call: _e.mock.On("getCompaction", planID)}
}

func (_c *MockCompactionPlanContext_getCompaction_Call) Run(run func(planID int64)) *MockCompactionPlanContext_getCompaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCompactionPlanContext_getCompaction_Call) Return(_a0 *compactionTask) *MockCompactionPlanContext_getCompaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_getCompaction_Call) RunAndReturn(run func(int64) *compactionTask) *MockCompactionPlanContext_getCompaction_Call {
	_c.Call.Return(run)
	return _c
}

// getCompactionTasksBySignalID provides a mock function with given fields: signalID
func (_m *MockCompactionPlanContext) getCompactionTasksBySignalID(signalID int64) []*compactionTask {
	ret := _m.Called(signalID)

	var r0 []*compactionTask
	if rf, ok := ret.Get(0).(func(int64) []*compactionTask); ok {
		r0 = rf(signalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*compactionTask)
		}
	}

	return r0
}

// MockCompactionPlanContext_getCompactionTasksBySignalID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getCompactionTasksBySignalID'
type MockCompactionPlanContext_getCompactionTasksBySignalID_Call struct {
	*mock.Call
}

// getCompactionTasksBySignalID is a helper method to define mock.On call
//   - signalID int64
func (_e *MockCompactionPlanContext_Expecter) getCompactionTasksBySignalID(signalID interface{}) *MockCompactionPlanContext_getCompactionTasksBySignalID_Call {
	return &MockCompactionPlanContext_getCompactionTasksBySignalID_Call{Call: _e.mock.On("getCompactionTasksBySignalID", signalID)}
}

func (_c *MockCompactionPlanContext_getCompactionTasksBySignalID_Call) Run(run func(signalID int64)) *MockCompactionPlanContext_getCompactionTasksBySignalID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCompactionPlanContext_getCompactionTasksBySignalID_Call) Return(_a0 []*compactionTask) *MockCompactionPlanContext_getCompactionTasksBySignalID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_getCompactionTasksBySignalID_Call) RunAndReturn(run func(int64) []*compactionTask) *MockCompactionPlanContext_getCompactionTasksBySignalID_Call {
	_c.Call.Return(run)
	return _c
}

// isFull provides a mock function with given fields:
func (_m *MockCompactionPlanContext) isFull() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCompactionPlanContext_isFull_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isFull'
type MockCompactionPlanContext_isFull_Call struct {
	*mock.Call
}

// isFull is a helper method to define mock.On call
func (_e *MockCompactionPlanContext_Expecter) isFull() *MockCompactionPlanContext_isFull_Call {
	return &MockCompactionPlanContext_isFull_Call{Call: _e.mock.On("isFull")}
}

func (_c *MockCompactionPlanContext_isFull_Call) Run(run func()) *MockCompactionPlanContext_isFull_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionPlanContext_isFull_Call) Return(_a0 bool) *MockCompactionPlanContext_isFull_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_isFull_Call) RunAndReturn(run func() bool) *MockCompactionPlanContext_isFull_Call {
	_c.Call.Return(run)
	return _c
}

// start provides a mock function with given fields:
func (_m *MockCompactionPlanContext) start() {
	_m.Called()
}

// MockCompactionPlanContext_start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'start'
type MockCompactionPlanContext_start_Call struct {
	*mock.Call
}

// start is a helper method to define mock.On call
func (_e *MockCompactionPlanContext_Expecter) start() *MockCompactionPlanContext_start_Call {
	return &MockCompactionPlanContext_start_Call{Call: _e.mock.On("start")}
}

func (_c *MockCompactionPlanContext_start_Call) Run(run func()) *MockCompactionPlanContext_start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionPlanContext_start_Call) Return() *MockCompactionPlanContext_start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCompactionPlanContext_start_Call) RunAndReturn(run func()) *MockCompactionPlanContext_start_Call {
	_c.Call.Return(run)
	return _c
}

// stop provides a mock function with given fields:
func (_m *MockCompactionPlanContext) stop() {
	_m.Called()
}

// MockCompactionPlanContext_stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'stop'
type MockCompactionPlanContext_stop_Call struct {
	*mock.Call
}

// stop is a helper method to define mock.On call
func (_e *MockCompactionPlanContext_Expecter) stop() *MockCompactionPlanContext_stop_Call {
	return &MockCompactionPlanContext_stop_Call{Call: _e.mock.On("stop")}
}

func (_c *MockCompactionPlanContext_stop_Call) Run(run func()) *MockCompactionPlanContext_stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionPlanContext_stop_Call) Return() *MockCompactionPlanContext_stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCompactionPlanContext_stop_Call) RunAndReturn(run func()) *MockCompactionPlanContext_stop_Call {
	_c.Call.Return(run)
	return _c
}

// updateCompaction provides a mock function with given fields: ts
func (_m *MockCompactionPlanContext) updateCompaction(ts uint64) error {
	ret := _m.Called(ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCompactionPlanContext_updateCompaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'updateCompaction'
type MockCompactionPlanContext_updateCompaction_Call struct {
	*mock.Call
}

// updateCompaction is a helper method to define mock.On call
//   - ts uint64
func (_e *MockCompactionPlanContext_Expecter) updateCompaction(ts interface{}) *MockCompactionPlanContext_updateCompaction_Call {
	return &MockCompactionPlanContext_updateCompaction_Call{Call: _e.mock.On("updateCompaction", ts)}
}

func (_c *MockCompactionPlanContext_updateCompaction_Call) Run(run func(ts uint64)) *MockCompactionPlanContext_updateCompaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockCompactionPlanContext_updateCompaction_Call) Return(_a0 error) *MockCompactionPlanContext_updateCompaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionPlanContext_updateCompaction_Call) RunAndReturn(run func(uint64) error) *MockCompactionPlanContext_updateCompaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCompactionPlanContext creates a new instance of MockCompactionPlanContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCompactionPlanContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCompactionPlanContext {
	mock := &MockCompactionPlanContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
