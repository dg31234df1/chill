// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_message

import (
	message "github.com/milvus-io/milvus/internal/util/logserviceutil/message"
	mock "github.com/stretchr/testify/mock"
)

// MockMessageID is an autogenerated mock type for the MessageID type
type MockMessageID struct {
	mock.Mock
}

type MockMessageID_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMessageID) EXPECT() *MockMessageID_Expecter {
	return &MockMessageID_Expecter{mock: &_m.Mock}
}

// EQ provides a mock function with given fields: _a0
func (_m *MockMessageID) EQ(_a0 message.MessageID) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(message.MessageID) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockMessageID_EQ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EQ'
type MockMessageID_EQ_Call struct {
	*mock.Call
}

// EQ is a helper method to define mock.On call
//   - _a0 message.MessageID
func (_e *MockMessageID_Expecter) EQ(_a0 interface{}) *MockMessageID_EQ_Call {
	return &MockMessageID_EQ_Call{Call: _e.mock.On("EQ", _a0)}
}

func (_c *MockMessageID_EQ_Call) Run(run func(_a0 message.MessageID)) *MockMessageID_EQ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(message.MessageID))
	})
	return _c
}

func (_c *MockMessageID_EQ_Call) Return(_a0 bool) *MockMessageID_EQ_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageID_EQ_Call) RunAndReturn(run func(message.MessageID) bool) *MockMessageID_EQ_Call {
	_c.Call.Return(run)
	return _c
}

// LT provides a mock function with given fields: _a0
func (_m *MockMessageID) LT(_a0 message.MessageID) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(message.MessageID) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockMessageID_LT_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LT'
type MockMessageID_LT_Call struct {
	*mock.Call
}

// LT is a helper method to define mock.On call
//   - _a0 message.MessageID
func (_e *MockMessageID_Expecter) LT(_a0 interface{}) *MockMessageID_LT_Call {
	return &MockMessageID_LT_Call{Call: _e.mock.On("LT", _a0)}
}

func (_c *MockMessageID_LT_Call) Run(run func(_a0 message.MessageID)) *MockMessageID_LT_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(message.MessageID))
	})
	return _c
}

func (_c *MockMessageID_LT_Call) Return(_a0 bool) *MockMessageID_LT_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageID_LT_Call) RunAndReturn(run func(message.MessageID) bool) *MockMessageID_LT_Call {
	_c.Call.Return(run)
	return _c
}

// LTE provides a mock function with given fields: _a0
func (_m *MockMessageID) LTE(_a0 message.MessageID) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(message.MessageID) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockMessageID_LTE_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LTE'
type MockMessageID_LTE_Call struct {
	*mock.Call
}

// LTE is a helper method to define mock.On call
//   - _a0 message.MessageID
func (_e *MockMessageID_Expecter) LTE(_a0 interface{}) *MockMessageID_LTE_Call {
	return &MockMessageID_LTE_Call{Call: _e.mock.On("LTE", _a0)}
}

func (_c *MockMessageID_LTE_Call) Run(run func(_a0 message.MessageID)) *MockMessageID_LTE_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(message.MessageID))
	})
	return _c
}

func (_c *MockMessageID_LTE_Call) Return(_a0 bool) *MockMessageID_LTE_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageID_LTE_Call) RunAndReturn(run func(message.MessageID) bool) *MockMessageID_LTE_Call {
	_c.Call.Return(run)
	return _c
}

// Marshal provides a mock function with given fields:
func (_m *MockMessageID) Marshal() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockMessageID_Marshal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Marshal'
type MockMessageID_Marshal_Call struct {
	*mock.Call
}

// Marshal is a helper method to define mock.On call
func (_e *MockMessageID_Expecter) Marshal() *MockMessageID_Marshal_Call {
	return &MockMessageID_Marshal_Call{Call: _e.mock.On("Marshal")}
}

func (_c *MockMessageID_Marshal_Call) Run(run func()) *MockMessageID_Marshal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMessageID_Marshal_Call) Return(_a0 []byte) *MockMessageID_Marshal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageID_Marshal_Call) RunAndReturn(run func() []byte) *MockMessageID_Marshal_Call {
	_c.Call.Return(run)
	return _c
}

// WALName provides a mock function with given fields:
func (_m *MockMessageID) WALName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockMessageID_WALName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WALName'
type MockMessageID_WALName_Call struct {
	*mock.Call
}

// WALName is a helper method to define mock.On call
func (_e *MockMessageID_Expecter) WALName() *MockMessageID_WALName_Call {
	return &MockMessageID_WALName_Call{Call: _e.mock.On("WALName")}
}

func (_c *MockMessageID_WALName_Call) Run(run func()) *MockMessageID_WALName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMessageID_WALName_Call) Return(_a0 string) *MockMessageID_WALName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageID_WALName_Call) RunAndReturn(run func() string) *MockMessageID_WALName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMessageID creates a new instance of MockMessageID. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessageID(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessageID {
	mock := &MockMessageID{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
