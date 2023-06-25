// Code generated by mockery v2.16.0. DO NOT EDIT.

package allocator

import (
	internalallocator "github.com/milvus-io/milvus/internal/allocator"
	mock "github.com/stretchr/testify/mock"
)

// MockAllocator is an autogenerated mock type for the Allocator type
type MockAllocator struct {
	mock.Mock
}

type MockAllocator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAllocator) EXPECT() *MockAllocator_Expecter {
	return &MockAllocator_Expecter{mock: &_m.Mock}
}

// Alloc provides a mock function with given fields: count
func (_m *MockAllocator) Alloc(count uint32) (int64, int64, error) {
	ret := _m.Called(count)

	var r0 int64
	if rf, ok := ret.Get(0).(func(uint32) int64); ok {
		r0 = rf(count)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(uint32) int64); ok {
		r1 = rf(count)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint32) error); ok {
		r2 = rf(count)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAllocator_Alloc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Alloc'
type MockAllocator_Alloc_Call struct {
	*mock.Call
}

// Alloc is a helper method to define mock.On call
//  - count uint32
func (_e *MockAllocator_Expecter) Alloc(count interface{}) *MockAllocator_Alloc_Call {
	return &MockAllocator_Alloc_Call{Call: _e.mock.On("Alloc", count)}
}

func (_c *MockAllocator_Alloc_Call) Run(run func(count uint32)) *MockAllocator_Alloc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockAllocator_Alloc_Call) Return(_a0 int64, _a1 int64, _a2 error) *MockAllocator_Alloc_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

// AllocOne provides a mock function with given fields:
func (_m *MockAllocator) AllocOne() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAllocator_AllocOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllocOne'
type MockAllocator_AllocOne_Call struct {
	*mock.Call
}

// AllocOne is a helper method to define mock.On call
func (_e *MockAllocator_Expecter) AllocOne() *MockAllocator_AllocOne_Call {
	return &MockAllocator_AllocOne_Call{Call: _e.mock.On("AllocOne")}
}

func (_c *MockAllocator_AllocOne_Call) Run(run func()) *MockAllocator_AllocOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAllocator_AllocOne_Call) Return(_a0 int64, _a1 error) *MockAllocator_AllocOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockAllocator) Close() {
	_m.Called()
}

// MockAllocator_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockAllocator_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockAllocator_Expecter) Close() *MockAllocator_Close_Call {
	return &MockAllocator_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockAllocator_Close_Call) Run(run func()) *MockAllocator_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAllocator_Close_Call) Return() *MockAllocator_Close_Call {
	_c.Call.Return()
	return _c
}

// GetGenerator provides a mock function with given fields: count, done
func (_m *MockAllocator) GetGenerator(count int, done <-chan struct{}) (<-chan int64, error) {
	ret := _m.Called(count, done)

	var r0 <-chan int64
	if rf, ok := ret.Get(0).(func(int, <-chan struct{}) <-chan int64); ok {
		r0 = rf(count, done)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, <-chan struct{}) error); ok {
		r1 = rf(count, done)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAllocator_GetGenerator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGenerator'
type MockAllocator_GetGenerator_Call struct {
	*mock.Call
}

// GetGenerator is a helper method to define mock.On call
//  - count int
//  - done <-chan struct{}
func (_e *MockAllocator_Expecter) GetGenerator(count interface{}, done interface{}) *MockAllocator_GetGenerator_Call {
	return &MockAllocator_GetGenerator_Call{Call: _e.mock.On("GetGenerator", count, done)}
}

func (_c *MockAllocator_GetGenerator_Call) Run(run func(count int, done <-chan struct{})) *MockAllocator_GetGenerator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(<-chan struct{}))
	})
	return _c
}

func (_c *MockAllocator_GetGenerator_Call) Return(_a0 <-chan int64, _a1 error) *MockAllocator_GetGenerator_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetIDAlloactor provides a mock function with given fields:
func (_m *MockAllocator) GetIDAlloactor() *internalallocator.IDAllocator {
	ret := _m.Called()

	var r0 *internalallocator.IDAllocator
	if rf, ok := ret.Get(0).(func() *internalallocator.IDAllocator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalallocator.IDAllocator)
		}
	}

	return r0
}

// MockAllocator_GetIDAlloactor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIDAlloactor'
type MockAllocator_GetIDAlloactor_Call struct {
	*mock.Call
}

// GetIDAlloactor is a helper method to define mock.On call
func (_e *MockAllocator_Expecter) GetIDAlloactor() *MockAllocator_GetIDAlloactor_Call {
	return &MockAllocator_GetIDAlloactor_Call{Call: _e.mock.On("GetIDAlloactor")}
}

func (_c *MockAllocator_GetIDAlloactor_Call) Run(run func()) *MockAllocator_GetIDAlloactor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAllocator_GetIDAlloactor_Call) Return(_a0 *internalallocator.IDAllocator) *MockAllocator_GetIDAlloactor_Call {
	_c.Call.Return(_a0)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockAllocator) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAllocator_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockAllocator_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockAllocator_Expecter) Start() *MockAllocator_Start_Call {
	return &MockAllocator_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockAllocator_Start_Call) Run(run func()) *MockAllocator_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAllocator_Start_Call) Return(_a0 error) *MockAllocator_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockAllocator interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAllocator creates a new instance of MockAllocator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAllocator(t mockConstructorTestingTNewMockAllocator) *MockAllocator {
	mock := &MockAllocator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
