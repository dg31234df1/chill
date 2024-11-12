// Code generated by mockery v2.46.0. DO NOT EDIT.

package proxy

import (
	context "context"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"
	mock "github.com/stretchr/testify/mock"
)

// MockLBBalancer is an autogenerated mock type for the LBBalancer type
type MockLBBalancer struct {
	mock.Mock
}

type MockLBBalancer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLBBalancer) EXPECT() *MockLBBalancer_Expecter {
	return &MockLBBalancer_Expecter{mock: &_m.Mock}
}

// CancelWorkload provides a mock function with given fields: node, nq
func (_m *MockLBBalancer) CancelWorkload(node int64, nq int64) {
	_m.Called(node, nq)
}

// MockLBBalancer_CancelWorkload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelWorkload'
type MockLBBalancer_CancelWorkload_Call struct {
	*mock.Call
}

// CancelWorkload is a helper method to define mock.On call
//   - node int64
//   - nq int64
func (_e *MockLBBalancer_Expecter) CancelWorkload(node interface{}, nq interface{}) *MockLBBalancer_CancelWorkload_Call {
	return &MockLBBalancer_CancelWorkload_Call{Call: _e.mock.On("CancelWorkload", node, nq)}
}

func (_c *MockLBBalancer_CancelWorkload_Call) Run(run func(node int64, nq int64)) *MockLBBalancer_CancelWorkload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int64))
	})
	return _c
}

func (_c *MockLBBalancer_CancelWorkload_Call) Return() *MockLBBalancer_CancelWorkload_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLBBalancer_CancelWorkload_Call) RunAndReturn(run func(int64, int64)) *MockLBBalancer_CancelWorkload_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockLBBalancer) Close() {
	_m.Called()
}

// MockLBBalancer_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockLBBalancer_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockLBBalancer_Expecter) Close() *MockLBBalancer_Close_Call {
	return &MockLBBalancer_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockLBBalancer_Close_Call) Run(run func()) *MockLBBalancer_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLBBalancer_Close_Call) Return() *MockLBBalancer_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLBBalancer_Close_Call) RunAndReturn(run func()) *MockLBBalancer_Close_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterNodeInfo provides a mock function with given fields: nodeInfos
func (_m *MockLBBalancer) RegisterNodeInfo(nodeInfos []nodeInfo) {
	_m.Called(nodeInfos)
}

// MockLBBalancer_RegisterNodeInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterNodeInfo'
type MockLBBalancer_RegisterNodeInfo_Call struct {
	*mock.Call
}

// RegisterNodeInfo is a helper method to define mock.On call
//   - nodeInfos []nodeInfo
func (_e *MockLBBalancer_Expecter) RegisterNodeInfo(nodeInfos interface{}) *MockLBBalancer_RegisterNodeInfo_Call {
	return &MockLBBalancer_RegisterNodeInfo_Call{Call: _e.mock.On("RegisterNodeInfo", nodeInfos)}
}

func (_c *MockLBBalancer_RegisterNodeInfo_Call) Run(run func(nodeInfos []nodeInfo)) *MockLBBalancer_RegisterNodeInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]nodeInfo))
	})
	return _c
}

func (_c *MockLBBalancer_RegisterNodeInfo_Call) Return() *MockLBBalancer_RegisterNodeInfo_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLBBalancer_RegisterNodeInfo_Call) RunAndReturn(run func([]nodeInfo)) *MockLBBalancer_RegisterNodeInfo_Call {
	_c.Call.Return(run)
	return _c
}

// SelectNode provides a mock function with given fields: ctx, availableNodes, nq
func (_m *MockLBBalancer) SelectNode(ctx context.Context, availableNodes []int64, nq int64) (int64, error) {
	ret := _m.Called(ctx, availableNodes, nq)

	if len(ret) == 0 {
		panic("no return value specified for SelectNode")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64, int64) (int64, error)); ok {
		return rf(ctx, availableNodes, nq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int64, int64) int64); ok {
		r0 = rf(ctx, availableNodes, nq)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int64, int64) error); ok {
		r1 = rf(ctx, availableNodes, nq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLBBalancer_SelectNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectNode'
type MockLBBalancer_SelectNode_Call struct {
	*mock.Call
}

// SelectNode is a helper method to define mock.On call
//   - ctx context.Context
//   - availableNodes []int64
//   - nq int64
func (_e *MockLBBalancer_Expecter) SelectNode(ctx interface{}, availableNodes interface{}, nq interface{}) *MockLBBalancer_SelectNode_Call {
	return &MockLBBalancer_SelectNode_Call{Call: _e.mock.On("SelectNode", ctx, availableNodes, nq)}
}

func (_c *MockLBBalancer_SelectNode_Call) Run(run func(ctx context.Context, availableNodes []int64, nq int64)) *MockLBBalancer_SelectNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]int64), args[2].(int64))
	})
	return _c
}

func (_c *MockLBBalancer_SelectNode_Call) Return(_a0 int64, _a1 error) *MockLBBalancer_SelectNode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLBBalancer_SelectNode_Call) RunAndReturn(run func(context.Context, []int64, int64) (int64, error)) *MockLBBalancer_SelectNode_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *MockLBBalancer) Start(ctx context.Context) {
	_m.Called(ctx)
}

// MockLBBalancer_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockLBBalancer_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockLBBalancer_Expecter) Start(ctx interface{}) *MockLBBalancer_Start_Call {
	return &MockLBBalancer_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *MockLBBalancer_Start_Call) Run(run func(ctx context.Context)) *MockLBBalancer_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockLBBalancer_Start_Call) Return() *MockLBBalancer_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLBBalancer_Start_Call) RunAndReturn(run func(context.Context)) *MockLBBalancer_Start_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCostMetrics provides a mock function with given fields: node, cost
func (_m *MockLBBalancer) UpdateCostMetrics(node int64, cost *internalpb.CostAggregation) {
	_m.Called(node, cost)
}

// MockLBBalancer_UpdateCostMetrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCostMetrics'
type MockLBBalancer_UpdateCostMetrics_Call struct {
	*mock.Call
}

// UpdateCostMetrics is a helper method to define mock.On call
//   - node int64
//   - cost *internalpb.CostAggregation
func (_e *MockLBBalancer_Expecter) UpdateCostMetrics(node interface{}, cost interface{}) *MockLBBalancer_UpdateCostMetrics_Call {
	return &MockLBBalancer_UpdateCostMetrics_Call{Call: _e.mock.On("UpdateCostMetrics", node, cost)}
}

func (_c *MockLBBalancer_UpdateCostMetrics_Call) Run(run func(node int64, cost *internalpb.CostAggregation)) *MockLBBalancer_UpdateCostMetrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*internalpb.CostAggregation))
	})
	return _c
}

func (_c *MockLBBalancer_UpdateCostMetrics_Call) Return() *MockLBBalancer_UpdateCostMetrics_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLBBalancer_UpdateCostMetrics_Call) RunAndReturn(run func(int64, *internalpb.CostAggregation)) *MockLBBalancer_UpdateCostMetrics_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLBBalancer creates a new instance of MockLBBalancer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLBBalancer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLBBalancer {
	mock := &MockLBBalancer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
