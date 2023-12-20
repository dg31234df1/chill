// Code generated by mockery v2.32.4. DO NOT EDIT.

package proxyutil

import (
	context "context"

	sessionutil "github.com/milvus-io/milvus/internal/util/sessionutil"
	mock "github.com/stretchr/testify/mock"
)

// MockProxyWatcher is an autogenerated mock type for the ProxyWatcherInterface type
type MockProxyWatcher struct {
	mock.Mock
}

type MockProxyWatcher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProxyWatcher) EXPECT() *MockProxyWatcher_Expecter {
	return &MockProxyWatcher_Expecter{mock: &_m.Mock}
}

// AddSessionFunc provides a mock function with given fields: fns
func (_m *MockProxyWatcher) AddSessionFunc(fns ...func(*sessionutil.Session)) {
	_va := make([]interface{}, len(fns))
	for _i := range fns {
		_va[_i] = fns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockProxyWatcher_AddSessionFunc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSessionFunc'
type MockProxyWatcher_AddSessionFunc_Call struct {
	*mock.Call
}

// AddSessionFunc is a helper method to define mock.On call
//   - fns ...func(*sessionutil.Session)
func (_e *MockProxyWatcher_Expecter) AddSessionFunc(fns ...interface{}) *MockProxyWatcher_AddSessionFunc_Call {
	return &MockProxyWatcher_AddSessionFunc_Call{Call: _e.mock.On("AddSessionFunc",
		append([]interface{}{}, fns...)...)}
}

func (_c *MockProxyWatcher_AddSessionFunc_Call) Run(run func(fns ...func(*sessionutil.Session))) *MockProxyWatcher_AddSessionFunc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*sessionutil.Session), len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(func(*sessionutil.Session))
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockProxyWatcher_AddSessionFunc_Call) Return() *MockProxyWatcher_AddSessionFunc_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockProxyWatcher_AddSessionFunc_Call) RunAndReturn(run func(...func(*sessionutil.Session))) *MockProxyWatcher_AddSessionFunc_Call {
	_c.Call.Return(run)
	return _c
}

// DelSessionFunc provides a mock function with given fields: fns
func (_m *MockProxyWatcher) DelSessionFunc(fns ...func(*sessionutil.Session)) {
	_va := make([]interface{}, len(fns))
	for _i := range fns {
		_va[_i] = fns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockProxyWatcher_DelSessionFunc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DelSessionFunc'
type MockProxyWatcher_DelSessionFunc_Call struct {
	*mock.Call
}

// DelSessionFunc is a helper method to define mock.On call
//   - fns ...func(*sessionutil.Session)
func (_e *MockProxyWatcher_Expecter) DelSessionFunc(fns ...interface{}) *MockProxyWatcher_DelSessionFunc_Call {
	return &MockProxyWatcher_DelSessionFunc_Call{Call: _e.mock.On("DelSessionFunc",
		append([]interface{}{}, fns...)...)}
}

func (_c *MockProxyWatcher_DelSessionFunc_Call) Run(run func(fns ...func(*sessionutil.Session))) *MockProxyWatcher_DelSessionFunc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*sessionutil.Session), len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(func(*sessionutil.Session))
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockProxyWatcher_DelSessionFunc_Call) Return() *MockProxyWatcher_DelSessionFunc_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockProxyWatcher_DelSessionFunc_Call) RunAndReturn(run func(...func(*sessionutil.Session))) *MockProxyWatcher_DelSessionFunc_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockProxyWatcher) Stop() {
	_m.Called()
}

// MockProxyWatcher_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockProxyWatcher_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockProxyWatcher_Expecter) Stop() *MockProxyWatcher_Stop_Call {
	return &MockProxyWatcher_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockProxyWatcher_Stop_Call) Run(run func()) *MockProxyWatcher_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProxyWatcher_Stop_Call) Return() *MockProxyWatcher_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockProxyWatcher_Stop_Call) RunAndReturn(run func()) *MockProxyWatcher_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// WatchProxy provides a mock function with given fields: ctx
func (_m *MockProxyWatcher) WatchProxy(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProxyWatcher_WatchProxy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchProxy'
type MockProxyWatcher_WatchProxy_Call struct {
	*mock.Call
}

// WatchProxy is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockProxyWatcher_Expecter) WatchProxy(ctx interface{}) *MockProxyWatcher_WatchProxy_Call {
	return &MockProxyWatcher_WatchProxy_Call{Call: _e.mock.On("WatchProxy", ctx)}
}

func (_c *MockProxyWatcher_WatchProxy_Call) Run(run func(ctx context.Context)) *MockProxyWatcher_WatchProxy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockProxyWatcher_WatchProxy_Call) Return(_a0 error) *MockProxyWatcher_WatchProxy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProxyWatcher_WatchProxy_Call) RunAndReturn(run func(context.Context) error) *MockProxyWatcher_WatchProxy_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProxyWatcher creates a new instance of MockProxyWatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProxyWatcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProxyWatcher {
	mock := &MockProxyWatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
