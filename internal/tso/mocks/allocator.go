// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocktso

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Allocator is an autogenerated mock type for the Allocator type
type Allocator struct {
	mock.Mock
}

type Allocator_Expecter struct {
	mock *mock.Mock
}

func (_m *Allocator) EXPECT() *Allocator_Expecter {
	return &Allocator_Expecter{mock: &_m.Mock}
}

// GenerateTSO provides a mock function with given fields: count
func (_m *Allocator) GenerateTSO(count uint32) (uint64, error) {
	ret := _m.Called(count)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) (uint64, error)); ok {
		return rf(count)
	}
	if rf, ok := ret.Get(0).(func(uint32) uint64); ok {
		r0 = rf(count)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Allocator_GenerateTSO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateTSO'
type Allocator_GenerateTSO_Call struct {
	*mock.Call
}

// GenerateTSO is a helper method to define mock.On call
//  - count uint32
func (_e *Allocator_Expecter) GenerateTSO(count interface{}) *Allocator_GenerateTSO_Call {
	return &Allocator_GenerateTSO_Call{Call: _e.mock.On("GenerateTSO", count)}
}

func (_c *Allocator_GenerateTSO_Call) Run(run func(count uint32)) *Allocator_GenerateTSO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *Allocator_GenerateTSO_Call) Return(_a0 uint64, _a1 error) *Allocator_GenerateTSO_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Allocator_GenerateTSO_Call) RunAndReturn(run func(uint32) (uint64, error)) *Allocator_GenerateTSO_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastSavedTime provides a mock function with given fields:
func (_m *Allocator) GetLastSavedTime() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Allocator_GetLastSavedTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastSavedTime'
type Allocator_GetLastSavedTime_Call struct {
	*mock.Call
}

// GetLastSavedTime is a helper method to define mock.On call
func (_e *Allocator_Expecter) GetLastSavedTime() *Allocator_GetLastSavedTime_Call {
	return &Allocator_GetLastSavedTime_Call{Call: _e.mock.On("GetLastSavedTime")}
}

func (_c *Allocator_GetLastSavedTime_Call) Run(run func()) *Allocator_GetLastSavedTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Allocator_GetLastSavedTime_Call) Return(_a0 time.Time) *Allocator_GetLastSavedTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Allocator_GetLastSavedTime_Call) RunAndReturn(run func() time.Time) *Allocator_GetLastSavedTime_Call {
	_c.Call.Return(run)
	return _c
}

// Initialize provides a mock function with given fields:
func (_m *Allocator) Initialize() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Allocator_Initialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Initialize'
type Allocator_Initialize_Call struct {
	*mock.Call
}

// Initialize is a helper method to define mock.On call
func (_e *Allocator_Expecter) Initialize() *Allocator_Initialize_Call {
	return &Allocator_Initialize_Call{Call: _e.mock.On("Initialize")}
}

func (_c *Allocator_Initialize_Call) Run(run func()) *Allocator_Initialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Allocator_Initialize_Call) Return(_a0 error) *Allocator_Initialize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Allocator_Initialize_Call) RunAndReturn(run func() error) *Allocator_Initialize_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields:
func (_m *Allocator) Reset() {
	_m.Called()
}

// Allocator_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type Allocator_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *Allocator_Expecter) Reset() *Allocator_Reset_Call {
	return &Allocator_Reset_Call{Call: _e.mock.On("Reset")}
}

func (_c *Allocator_Reset_Call) Run(run func()) *Allocator_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Allocator_Reset_Call) Return() *Allocator_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *Allocator_Reset_Call) RunAndReturn(run func()) *Allocator_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// SetTSO provides a mock function with given fields: _a0
func (_m *Allocator) SetTSO(_a0 uint64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Allocator_SetTSO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTSO'
type Allocator_SetTSO_Call struct {
	*mock.Call
}

// SetTSO is a helper method to define mock.On call
//  - _a0 uint64
func (_e *Allocator_Expecter) SetTSO(_a0 interface{}) *Allocator_SetTSO_Call {
	return &Allocator_SetTSO_Call{Call: _e.mock.On("SetTSO", _a0)}
}

func (_c *Allocator_SetTSO_Call) Run(run func(_a0 uint64)) *Allocator_SetTSO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *Allocator_SetTSO_Call) Return(_a0 error) *Allocator_SetTSO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Allocator_SetTSO_Call) RunAndReturn(run func(uint64) error) *Allocator_SetTSO_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTSO provides a mock function with given fields:
func (_m *Allocator) UpdateTSO() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Allocator_UpdateTSO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTSO'
type Allocator_UpdateTSO_Call struct {
	*mock.Call
}

// UpdateTSO is a helper method to define mock.On call
func (_e *Allocator_Expecter) UpdateTSO() *Allocator_UpdateTSO_Call {
	return &Allocator_UpdateTSO_Call{Call: _e.mock.On("UpdateTSO")}
}

func (_c *Allocator_UpdateTSO_Call) Run(run func()) *Allocator_UpdateTSO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Allocator_UpdateTSO_Call) Return(_a0 error) *Allocator_UpdateTSO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Allocator_UpdateTSO_Call) RunAndReturn(run func() error) *Allocator_UpdateTSO_Call {
	_c.Call.Return(run)
	return _c
}

// NewAllocator creates a new instance of Allocator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAllocator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Allocator {
	mock := &Allocator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
