// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_wal

import (
	context "context"

	message "github.com/milvus-io/milvus/pkg/streaming/util/message"
	mock "github.com/stretchr/testify/mock"

	types "github.com/milvus-io/milvus/pkg/streaming/util/types"

	wal "github.com/milvus-io/milvus/internal/streamingnode/server/wal"
)

// MockWAL is an autogenerated mock type for the WAL type
type MockWAL struct {
	mock.Mock
}

type MockWAL_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWAL) EXPECT() *MockWAL_Expecter {
	return &MockWAL_Expecter{mock: &_m.Mock}
}

// Append provides a mock function with given fields: ctx, msg
func (_m *MockWAL) Append(ctx context.Context, msg message.MutableMessage) (message.MessageID, error) {
	ret := _m.Called(ctx, msg)

	var r0 message.MessageID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, message.MutableMessage) (message.MessageID, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, message.MutableMessage) message.MessageID); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(message.MessageID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, message.MutableMessage) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWAL_Append_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Append'
type MockWAL_Append_Call struct {
	*mock.Call
}

// Append is a helper method to define mock.On call
//   - ctx context.Context
//   - msg message.MutableMessage
func (_e *MockWAL_Expecter) Append(ctx interface{}, msg interface{}) *MockWAL_Append_Call {
	return &MockWAL_Append_Call{Call: _e.mock.On("Append", ctx, msg)}
}

func (_c *MockWAL_Append_Call) Run(run func(ctx context.Context, msg message.MutableMessage)) *MockWAL_Append_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(message.MutableMessage))
	})
	return _c
}

func (_c *MockWAL_Append_Call) Return(_a0 message.MessageID, _a1 error) *MockWAL_Append_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWAL_Append_Call) RunAndReturn(run func(context.Context, message.MutableMessage) (message.MessageID, error)) *MockWAL_Append_Call {
	_c.Call.Return(run)
	return _c
}

// AppendAsync provides a mock function with given fields: ctx, msg, cb
func (_m *MockWAL) AppendAsync(ctx context.Context, msg message.MutableMessage, cb func(message.MessageID, error)) {
	_m.Called(ctx, msg, cb)
}

// MockWAL_AppendAsync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AppendAsync'
type MockWAL_AppendAsync_Call struct {
	*mock.Call
}

// AppendAsync is a helper method to define mock.On call
//   - ctx context.Context
//   - msg message.MutableMessage
//   - cb func(message.MessageID , error)
func (_e *MockWAL_Expecter) AppendAsync(ctx interface{}, msg interface{}, cb interface{}) *MockWAL_AppendAsync_Call {
	return &MockWAL_AppendAsync_Call{Call: _e.mock.On("AppendAsync", ctx, msg, cb)}
}

func (_c *MockWAL_AppendAsync_Call) Run(run func(ctx context.Context, msg message.MutableMessage, cb func(message.MessageID, error))) *MockWAL_AppendAsync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(message.MutableMessage), args[2].(func(message.MessageID, error)))
	})
	return _c
}

func (_c *MockWAL_AppendAsync_Call) Return() *MockWAL_AppendAsync_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWAL_AppendAsync_Call) RunAndReturn(run func(context.Context, message.MutableMessage, func(message.MessageID, error))) *MockWAL_AppendAsync_Call {
	_c.Call.Return(run)
	return _c
}

// Channel provides a mock function with given fields:
func (_m *MockWAL) Channel() types.PChannelInfo {
	ret := _m.Called()

	var r0 types.PChannelInfo
	if rf, ok := ret.Get(0).(func() types.PChannelInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.PChannelInfo)
	}

	return r0
}

// MockWAL_Channel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Channel'
type MockWAL_Channel_Call struct {
	*mock.Call
}

// Channel is a helper method to define mock.On call
func (_e *MockWAL_Expecter) Channel() *MockWAL_Channel_Call {
	return &MockWAL_Channel_Call{Call: _e.mock.On("Channel")}
}

func (_c *MockWAL_Channel_Call) Run(run func()) *MockWAL_Channel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWAL_Channel_Call) Return(_a0 types.PChannelInfo) *MockWAL_Channel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWAL_Channel_Call) RunAndReturn(run func() types.PChannelInfo) *MockWAL_Channel_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockWAL) Close() {
	_m.Called()
}

// MockWAL_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockWAL_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockWAL_Expecter) Close() *MockWAL_Close_Call {
	return &MockWAL_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockWAL_Close_Call) Run(run func()) *MockWAL_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWAL_Close_Call) Return() *MockWAL_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWAL_Close_Call) RunAndReturn(run func()) *MockWAL_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, deliverPolicy
func (_m *MockWAL) Read(ctx context.Context, deliverPolicy wal.ReadOption) (wal.Scanner, error) {
	ret := _m.Called(ctx, deliverPolicy)

	var r0 wal.Scanner
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, wal.ReadOption) (wal.Scanner, error)); ok {
		return rf(ctx, deliverPolicy)
	}
	if rf, ok := ret.Get(0).(func(context.Context, wal.ReadOption) wal.Scanner); ok {
		r0 = rf(ctx, deliverPolicy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wal.Scanner)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, wal.ReadOption) error); ok {
		r1 = rf(ctx, deliverPolicy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWAL_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockWAL_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - deliverPolicy wal.ReadOption
func (_e *MockWAL_Expecter) Read(ctx interface{}, deliverPolicy interface{}) *MockWAL_Read_Call {
	return &MockWAL_Read_Call{Call: _e.mock.On("Read", ctx, deliverPolicy)}
}

func (_c *MockWAL_Read_Call) Run(run func(ctx context.Context, deliverPolicy wal.ReadOption)) *MockWAL_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wal.ReadOption))
	})
	return _c
}

func (_c *MockWAL_Read_Call) Return(_a0 wal.Scanner, _a1 error) *MockWAL_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWAL_Read_Call) RunAndReturn(run func(context.Context, wal.ReadOption) (wal.Scanner, error)) *MockWAL_Read_Call {
	_c.Call.Return(run)
	return _c
}

// WALName provides a mock function with given fields:
func (_m *MockWAL) WALName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockWAL_WALName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WALName'
type MockWAL_WALName_Call struct {
	*mock.Call
}

// WALName is a helper method to define mock.On call
func (_e *MockWAL_Expecter) WALName() *MockWAL_WALName_Call {
	return &MockWAL_WALName_Call{Call: _e.mock.On("WALName")}
}

func (_c *MockWAL_WALName_Call) Run(run func()) *MockWAL_WALName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWAL_WALName_Call) Return(_a0 string) *MockWAL_WALName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWAL_WALName_Call) RunAndReturn(run func() string) *MockWAL_WALName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWAL creates a new instance of MockWAL. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWAL(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWAL {
	mock := &MockWAL{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
