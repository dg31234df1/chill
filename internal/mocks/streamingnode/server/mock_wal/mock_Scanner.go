// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_wal

import (
	message "github.com/milvus-io/milvus/pkg/streaming/util/message"
	mock "github.com/stretchr/testify/mock"

	types "github.com/milvus-io/milvus/pkg/streaming/util/types"
)

// MockScanner is an autogenerated mock type for the Scanner type
type MockScanner struct {
	mock.Mock
}

type MockScanner_Expecter struct {
	mock *mock.Mock
}

func (_m *MockScanner) EXPECT() *MockScanner_Expecter {
	return &MockScanner_Expecter{mock: &_m.Mock}
}

// Chan provides a mock function with given fields:
func (_m *MockScanner) Chan() <-chan message.ImmutableMessage {
	ret := _m.Called()

	var r0 <-chan message.ImmutableMessage
	if rf, ok := ret.Get(0).(func() <-chan message.ImmutableMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan message.ImmutableMessage)
		}
	}

	return r0
}

// MockScanner_Chan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chan'
type MockScanner_Chan_Call struct {
	*mock.Call
}

// Chan is a helper method to define mock.On call
func (_e *MockScanner_Expecter) Chan() *MockScanner_Chan_Call {
	return &MockScanner_Chan_Call{Call: _e.mock.On("Chan")}
}

func (_c *MockScanner_Chan_Call) Run(run func()) *MockScanner_Chan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScanner_Chan_Call) Return(_a0 <-chan message.ImmutableMessage) *MockScanner_Chan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScanner_Chan_Call) RunAndReturn(run func() <-chan message.ImmutableMessage) *MockScanner_Chan_Call {
	_c.Call.Return(run)
	return _c
}

// Channel provides a mock function with given fields:
func (_m *MockScanner) Channel() types.PChannelInfo {
	ret := _m.Called()

	var r0 types.PChannelInfo
	if rf, ok := ret.Get(0).(func() types.PChannelInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.PChannelInfo)
	}

	return r0
}

// MockScanner_Channel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Channel'
type MockScanner_Channel_Call struct {
	*mock.Call
}

// Channel is a helper method to define mock.On call
func (_e *MockScanner_Expecter) Channel() *MockScanner_Channel_Call {
	return &MockScanner_Channel_Call{Call: _e.mock.On("Channel")}
}

func (_c *MockScanner_Channel_Call) Run(run func()) *MockScanner_Channel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScanner_Channel_Call) Return(_a0 types.PChannelInfo) *MockScanner_Channel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScanner_Channel_Call) RunAndReturn(run func() types.PChannelInfo) *MockScanner_Channel_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockScanner) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockScanner_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockScanner_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockScanner_Expecter) Close() *MockScanner_Close_Call {
	return &MockScanner_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockScanner_Close_Call) Run(run func()) *MockScanner_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScanner_Close_Call) Return(_a0 error) *MockScanner_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScanner_Close_Call) RunAndReturn(run func() error) *MockScanner_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Done provides a mock function with given fields:
func (_m *MockScanner) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// MockScanner_Done_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Done'
type MockScanner_Done_Call struct {
	*mock.Call
}

// Done is a helper method to define mock.On call
func (_e *MockScanner_Expecter) Done() *MockScanner_Done_Call {
	return &MockScanner_Done_Call{Call: _e.mock.On("Done")}
}

func (_c *MockScanner_Done_Call) Run(run func()) *MockScanner_Done_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScanner_Done_Call) Return(_a0 <-chan struct{}) *MockScanner_Done_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScanner_Done_Call) RunAndReturn(run func() <-chan struct{}) *MockScanner_Done_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields:
func (_m *MockScanner) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockScanner_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type MockScanner_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
func (_e *MockScanner_Expecter) Error() *MockScanner_Error_Call {
	return &MockScanner_Error_Call{Call: _e.mock.On("Error")}
}

func (_c *MockScanner_Error_Call) Run(run func()) *MockScanner_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockScanner_Error_Call) Return(_a0 error) *MockScanner_Error_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScanner_Error_Call) RunAndReturn(run func() error) *MockScanner_Error_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockScanner creates a new instance of MockScanner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockScanner(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockScanner {
	mock := &MockScanner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
