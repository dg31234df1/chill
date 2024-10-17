// Code generated by mockery v2.46.0. DO NOT EDIT.

package mock_resolver

import (
	mock "github.com/stretchr/testify/mock"
	resolver "google.golang.org/grpc/resolver"

	serviceconfig "google.golang.org/grpc/serviceconfig"
)

// MockClientConn is an autogenerated mock type for the ClientConn type
type MockClientConn struct {
	mock.Mock
}

type MockClientConn_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientConn) EXPECT() *MockClientConn_Expecter {
	return &MockClientConn_Expecter{mock: &_m.Mock}
}

// NewAddress provides a mock function with given fields: addresses
func (_m *MockClientConn) NewAddress(addresses []resolver.Address) {
	_m.Called(addresses)
}

// MockClientConn_NewAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewAddress'
type MockClientConn_NewAddress_Call struct {
	*mock.Call
}

// NewAddress is a helper method to define mock.On call
//   - addresses []resolver.Address
func (_e *MockClientConn_Expecter) NewAddress(addresses interface{}) *MockClientConn_NewAddress_Call {
	return &MockClientConn_NewAddress_Call{Call: _e.mock.On("NewAddress", addresses)}
}

func (_c *MockClientConn_NewAddress_Call) Run(run func(addresses []resolver.Address)) *MockClientConn_NewAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]resolver.Address))
	})
	return _c
}

func (_c *MockClientConn_NewAddress_Call) Return() *MockClientConn_NewAddress_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClientConn_NewAddress_Call) RunAndReturn(run func([]resolver.Address)) *MockClientConn_NewAddress_Call {
	_c.Call.Return(run)
	return _c
}

// ParseServiceConfig provides a mock function with given fields: serviceConfigJSON
func (_m *MockClientConn) ParseServiceConfig(serviceConfigJSON string) *serviceconfig.ParseResult {
	ret := _m.Called(serviceConfigJSON)

	if len(ret) == 0 {
		panic("no return value specified for ParseServiceConfig")
	}

	var r0 *serviceconfig.ParseResult
	if rf, ok := ret.Get(0).(func(string) *serviceconfig.ParseResult); ok {
		r0 = rf(serviceConfigJSON)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceconfig.ParseResult)
		}
	}

	return r0
}

// MockClientConn_ParseServiceConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseServiceConfig'
type MockClientConn_ParseServiceConfig_Call struct {
	*mock.Call
}

// ParseServiceConfig is a helper method to define mock.On call
//   - serviceConfigJSON string
func (_e *MockClientConn_Expecter) ParseServiceConfig(serviceConfigJSON interface{}) *MockClientConn_ParseServiceConfig_Call {
	return &MockClientConn_ParseServiceConfig_Call{Call: _e.mock.On("ParseServiceConfig", serviceConfigJSON)}
}

func (_c *MockClientConn_ParseServiceConfig_Call) Run(run func(serviceConfigJSON string)) *MockClientConn_ParseServiceConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockClientConn_ParseServiceConfig_Call) Return(_a0 *serviceconfig.ParseResult) *MockClientConn_ParseServiceConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientConn_ParseServiceConfig_Call) RunAndReturn(run func(string) *serviceconfig.ParseResult) *MockClientConn_ParseServiceConfig_Call {
	_c.Call.Return(run)
	return _c
}

// ReportError provides a mock function with given fields: _a0
func (_m *MockClientConn) ReportError(_a0 error) {
	_m.Called(_a0)
}

// MockClientConn_ReportError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReportError'
type MockClientConn_ReportError_Call struct {
	*mock.Call
}

// ReportError is a helper method to define mock.On call
//   - _a0 error
func (_e *MockClientConn_Expecter) ReportError(_a0 interface{}) *MockClientConn_ReportError_Call {
	return &MockClientConn_ReportError_Call{Call: _e.mock.On("ReportError", _a0)}
}

func (_c *MockClientConn_ReportError_Call) Run(run func(_a0 error)) *MockClientConn_ReportError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockClientConn_ReportError_Call) Return() *MockClientConn_ReportError_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClientConn_ReportError_Call) RunAndReturn(run func(error)) *MockClientConn_ReportError_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateState provides a mock function with given fields: _a0
func (_m *MockClientConn) UpdateState(_a0 resolver.State) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(resolver.State) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClientConn_UpdateState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateState'
type MockClientConn_UpdateState_Call struct {
	*mock.Call
}

// UpdateState is a helper method to define mock.On call
//   - _a0 resolver.State
func (_e *MockClientConn_Expecter) UpdateState(_a0 interface{}) *MockClientConn_UpdateState_Call {
	return &MockClientConn_UpdateState_Call{Call: _e.mock.On("UpdateState", _a0)}
}

func (_c *MockClientConn_UpdateState_Call) Run(run func(_a0 resolver.State)) *MockClientConn_UpdateState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(resolver.State))
	})
	return _c
}

func (_c *MockClientConn_UpdateState_Call) Return(_a0 error) *MockClientConn_UpdateState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientConn_UpdateState_Call) RunAndReturn(run func(resolver.State) error) *MockClientConn_UpdateState_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientConn creates a new instance of MockClientConn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientConn(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientConn {
	mock := &MockClientConn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
