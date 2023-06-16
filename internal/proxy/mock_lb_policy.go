// Code generated by mockery v2.21.1. DO NOT EDIT.

package proxy

import (
	context "context"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"
	mock "github.com/stretchr/testify/mock"
)

// MockLBPolicy is an autogenerated mock type for the LBPolicy type
type MockLBPolicy struct {
	mock.Mock
}

type MockLBPolicy_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLBPolicy) EXPECT() *MockLBPolicy_Expecter {
	return &MockLBPolicy_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, workload
func (_m *MockLBPolicy) Execute(ctx context.Context, workload CollectionWorkLoad) error {
	ret := _m.Called(ctx, workload)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CollectionWorkLoad) error); ok {
		r0 = rf(ctx, workload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLBPolicy_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockLBPolicy_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - workload CollectionWorkLoad
func (_e *MockLBPolicy_Expecter) Execute(ctx interface{}, workload interface{}) *MockLBPolicy_Execute_Call {
	return &MockLBPolicy_Execute_Call{Call: _e.mock.On("Execute", ctx, workload)}
}

func (_c *MockLBPolicy_Execute_Call) Run(run func(ctx context.Context, workload CollectionWorkLoad)) *MockLBPolicy_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(CollectionWorkLoad))
	})
	return _c
}

func (_c *MockLBPolicy_Execute_Call) Return(_a0 error) *MockLBPolicy_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLBPolicy_Execute_Call) RunAndReturn(run func(context.Context, CollectionWorkLoad) error) *MockLBPolicy_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteWithRetry provides a mock function with given fields: ctx, workload
func (_m *MockLBPolicy) ExecuteWithRetry(ctx context.Context, workload ChannelWorkload) error {
	ret := _m.Called(ctx, workload)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ChannelWorkload) error); ok {
		r0 = rf(ctx, workload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLBPolicy_ExecuteWithRetry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteWithRetry'
type MockLBPolicy_ExecuteWithRetry_Call struct {
	*mock.Call
}

// ExecuteWithRetry is a helper method to define mock.On call
//   - ctx context.Context
//   - workload ChannelWorkload
func (_e *MockLBPolicy_Expecter) ExecuteWithRetry(ctx interface{}, workload interface{}) *MockLBPolicy_ExecuteWithRetry_Call {
	return &MockLBPolicy_ExecuteWithRetry_Call{Call: _e.mock.On("ExecuteWithRetry", ctx, workload)}
}

func (_c *MockLBPolicy_ExecuteWithRetry_Call) Run(run func(ctx context.Context, workload ChannelWorkload)) *MockLBPolicy_ExecuteWithRetry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ChannelWorkload))
	})
	return _c
}

func (_c *MockLBPolicy_ExecuteWithRetry_Call) Return(_a0 error) *MockLBPolicy_ExecuteWithRetry_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLBPolicy_ExecuteWithRetry_Call) RunAndReturn(run func(context.Context, ChannelWorkload) error) *MockLBPolicy_ExecuteWithRetry_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCostMetrics provides a mock function with given fields: node, cost
func (_m *MockLBPolicy) UpdateCostMetrics(node int64, cost *internalpb.CostAggregation) {
	_m.Called(node, cost)
}

// MockLBPolicy_UpdateCostMetrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCostMetrics'
type MockLBPolicy_UpdateCostMetrics_Call struct {
	*mock.Call
}

// UpdateCostMetrics is a helper method to define mock.On call
//   - node int64
//   - cost *internalpb.CostAggregation
func (_e *MockLBPolicy_Expecter) UpdateCostMetrics(node interface{}, cost interface{}) *MockLBPolicy_UpdateCostMetrics_Call {
	return &MockLBPolicy_UpdateCostMetrics_Call{Call: _e.mock.On("UpdateCostMetrics", node, cost)}
}

func (_c *MockLBPolicy_UpdateCostMetrics_Call) Run(run func(node int64, cost *internalpb.CostAggregation)) *MockLBPolicy_UpdateCostMetrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*internalpb.CostAggregation))
	})
	return _c
}

func (_c *MockLBPolicy_UpdateCostMetrics_Call) Return() *MockLBPolicy_UpdateCostMetrics_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLBPolicy_UpdateCostMetrics_Call) RunAndReturn(run func(int64, *internalpb.CostAggregation)) *MockLBPolicy_UpdateCostMetrics_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockLBPolicy interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockLBPolicy creates a new instance of MockLBPolicy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockLBPolicy(t mockConstructorTestingTNewMockLBPolicy) *MockLBPolicy {
	mock := &MockLBPolicy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
