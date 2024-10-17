// Code generated by mockery v2.46.0. DO NOT EDIT.

package mock_types

import (
	context "context"

	types "github.com/milvus-io/milvus/pkg/streaming/util/types"
	mock "github.com/stretchr/testify/mock"
)

// MockAssignmentDiscoverWatcher is an autogenerated mock type for the AssignmentDiscoverWatcher type
type MockAssignmentDiscoverWatcher struct {
	mock.Mock
}

type MockAssignmentDiscoverWatcher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAssignmentDiscoverWatcher) EXPECT() *MockAssignmentDiscoverWatcher_Expecter {
	return &MockAssignmentDiscoverWatcher_Expecter{mock: &_m.Mock}
}

// AssignmentDiscover provides a mock function with given fields: ctx, cb
func (_m *MockAssignmentDiscoverWatcher) AssignmentDiscover(ctx context.Context, cb func(*types.VersionedStreamingNodeAssignments) error) error {
	ret := _m.Called(ctx, cb)

	if len(ret) == 0 {
		panic("no return value specified for AssignmentDiscover")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(*types.VersionedStreamingNodeAssignments) error) error); ok {
		r0 = rf(ctx, cb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAssignmentDiscoverWatcher_AssignmentDiscover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AssignmentDiscover'
type MockAssignmentDiscoverWatcher_AssignmentDiscover_Call struct {
	*mock.Call
}

// AssignmentDiscover is a helper method to define mock.On call
//   - ctx context.Context
//   - cb func(*types.VersionedStreamingNodeAssignments) error
func (_e *MockAssignmentDiscoverWatcher_Expecter) AssignmentDiscover(ctx interface{}, cb interface{}) *MockAssignmentDiscoverWatcher_AssignmentDiscover_Call {
	return &MockAssignmentDiscoverWatcher_AssignmentDiscover_Call{Call: _e.mock.On("AssignmentDiscover", ctx, cb)}
}

func (_c *MockAssignmentDiscoverWatcher_AssignmentDiscover_Call) Run(run func(ctx context.Context, cb func(*types.VersionedStreamingNodeAssignments) error)) *MockAssignmentDiscoverWatcher_AssignmentDiscover_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(*types.VersionedStreamingNodeAssignments) error))
	})
	return _c
}

func (_c *MockAssignmentDiscoverWatcher_AssignmentDiscover_Call) Return(_a0 error) *MockAssignmentDiscoverWatcher_AssignmentDiscover_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAssignmentDiscoverWatcher_AssignmentDiscover_Call) RunAndReturn(run func(context.Context, func(*types.VersionedStreamingNodeAssignments) error) error) *MockAssignmentDiscoverWatcher_AssignmentDiscover_Call {
	_c.Call.Return(run)
	return _c
}

// ReportAssignmentError provides a mock function with given fields: ctx, pchannel, err
func (_m *MockAssignmentDiscoverWatcher) ReportAssignmentError(ctx context.Context, pchannel types.PChannelInfo, err error) error {
	ret := _m.Called(ctx, pchannel, err)

	if len(ret) == 0 {
		panic("no return value specified for ReportAssignmentError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.PChannelInfo, error) error); ok {
		r0 = rf(ctx, pchannel, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAssignmentDiscoverWatcher_ReportAssignmentError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReportAssignmentError'
type MockAssignmentDiscoverWatcher_ReportAssignmentError_Call struct {
	*mock.Call
}

// ReportAssignmentError is a helper method to define mock.On call
//   - ctx context.Context
//   - pchannel types.PChannelInfo
//   - err error
func (_e *MockAssignmentDiscoverWatcher_Expecter) ReportAssignmentError(ctx interface{}, pchannel interface{}, err interface{}) *MockAssignmentDiscoverWatcher_ReportAssignmentError_Call {
	return &MockAssignmentDiscoverWatcher_ReportAssignmentError_Call{Call: _e.mock.On("ReportAssignmentError", ctx, pchannel, err)}
}

func (_c *MockAssignmentDiscoverWatcher_ReportAssignmentError_Call) Run(run func(ctx context.Context, pchannel types.PChannelInfo, err error)) *MockAssignmentDiscoverWatcher_ReportAssignmentError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.PChannelInfo), args[2].(error))
	})
	return _c
}

func (_c *MockAssignmentDiscoverWatcher_ReportAssignmentError_Call) Return(_a0 error) *MockAssignmentDiscoverWatcher_ReportAssignmentError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAssignmentDiscoverWatcher_ReportAssignmentError_Call) RunAndReturn(run func(context.Context, types.PChannelInfo, error) error) *MockAssignmentDiscoverWatcher_ReportAssignmentError_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAssignmentDiscoverWatcher creates a new instance of MockAssignmentDiscoverWatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAssignmentDiscoverWatcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAssignmentDiscoverWatcher {
	mock := &MockAssignmentDiscoverWatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
