// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_balancer

import (
	mock "github.com/stretchr/testify/mock"
	balancer "google.golang.org/grpc/balancer"

	resolver "google.golang.org/grpc/resolver"
)

// MockSubConn is an autogenerated mock type for the SubConn type
type MockSubConn struct {
	mock.Mock
}

type MockSubConn_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSubConn) EXPECT() *MockSubConn_Expecter {
	return &MockSubConn_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields:
func (_m *MockSubConn) Connect() {
	_m.Called()
}

// MockSubConn_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockSubConn_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockSubConn_Expecter) Connect() *MockSubConn_Connect_Call {
	return &MockSubConn_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockSubConn_Connect_Call) Run(run func()) *MockSubConn_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSubConn_Connect_Call) Return() *MockSubConn_Connect_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSubConn_Connect_Call) RunAndReturn(run func()) *MockSubConn_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrBuildProducer provides a mock function with given fields: _a0
func (_m *MockSubConn) GetOrBuildProducer(_a0 balancer.ProducerBuilder) (balancer.Producer, func()) {
	ret := _m.Called(_a0)

	var r0 balancer.Producer
	var r1 func()
	if rf, ok := ret.Get(0).(func(balancer.ProducerBuilder) (balancer.Producer, func())); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(balancer.ProducerBuilder) balancer.Producer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(balancer.Producer)
		}
	}

	if rf, ok := ret.Get(1).(func(balancer.ProducerBuilder) func()); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	return r0, r1
}

// MockSubConn_GetOrBuildProducer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrBuildProducer'
type MockSubConn_GetOrBuildProducer_Call struct {
	*mock.Call
}

// GetOrBuildProducer is a helper method to define mock.On call
//   - _a0 balancer.ProducerBuilder
func (_e *MockSubConn_Expecter) GetOrBuildProducer(_a0 interface{}) *MockSubConn_GetOrBuildProducer_Call {
	return &MockSubConn_GetOrBuildProducer_Call{Call: _e.mock.On("GetOrBuildProducer", _a0)}
}

func (_c *MockSubConn_GetOrBuildProducer_Call) Run(run func(_a0 balancer.ProducerBuilder)) *MockSubConn_GetOrBuildProducer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(balancer.ProducerBuilder))
	})
	return _c
}

func (_c *MockSubConn_GetOrBuildProducer_Call) Return(p balancer.Producer, close func()) *MockSubConn_GetOrBuildProducer_Call {
	_c.Call.Return(p, close)
	return _c
}

func (_c *MockSubConn_GetOrBuildProducer_Call) RunAndReturn(run func(balancer.ProducerBuilder) (balancer.Producer, func())) *MockSubConn_GetOrBuildProducer_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAddresses provides a mock function with given fields: _a0
func (_m *MockSubConn) UpdateAddresses(_a0 []resolver.Address) {
	_m.Called(_a0)
}

// MockSubConn_UpdateAddresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAddresses'
type MockSubConn_UpdateAddresses_Call struct {
	*mock.Call
}

// UpdateAddresses is a helper method to define mock.On call
//   - _a0 []resolver.Address
func (_e *MockSubConn_Expecter) UpdateAddresses(_a0 interface{}) *MockSubConn_UpdateAddresses_Call {
	return &MockSubConn_UpdateAddresses_Call{Call: _e.mock.On("UpdateAddresses", _a0)}
}

func (_c *MockSubConn_UpdateAddresses_Call) Run(run func(_a0 []resolver.Address)) *MockSubConn_UpdateAddresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]resolver.Address))
	})
	return _c
}

func (_c *MockSubConn_UpdateAddresses_Call) Return() *MockSubConn_UpdateAddresses_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSubConn_UpdateAddresses_Call) RunAndReturn(run func([]resolver.Address)) *MockSubConn_UpdateAddresses_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSubConn creates a new instance of MockSubConn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSubConn(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSubConn {
	mock := &MockSubConn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
