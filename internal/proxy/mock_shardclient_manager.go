// Code generated by mockery v2.46.0. DO NOT EDIT.

package proxy

import (
	context "context"

	types "github.com/milvus-io/milvus/internal/types"
	mock "github.com/stretchr/testify/mock"
)

// MockShardClientManager is an autogenerated mock type for the shardClientMgr type
type MockShardClientManager struct {
	mock.Mock
}

type MockShardClientManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockShardClientManager) EXPECT() *MockShardClientManager_Expecter {
	return &MockShardClientManager_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockShardClientManager) Close() {
	_m.Called()
}

// MockShardClientManager_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockShardClientManager_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockShardClientManager_Expecter) Close() *MockShardClientManager_Close_Call {
	return &MockShardClientManager_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockShardClientManager_Close_Call) Run(run func()) *MockShardClientManager_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockShardClientManager_Close_Call) Return() *MockShardClientManager_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockShardClientManager_Close_Call) RunAndReturn(run func()) *MockShardClientManager_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetClient provides a mock function with given fields: ctx, nodeInfo1
func (_m *MockShardClientManager) GetClient(ctx context.Context, nodeInfo1 nodeInfo) (types.QueryNodeClient, error) {
	ret := _m.Called(ctx, nodeInfo1)

	if len(ret) == 0 {
		panic("no return value specified for GetClient")
	}

	var r0 types.QueryNodeClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, nodeInfo) (types.QueryNodeClient, error)); ok {
		return rf(ctx, nodeInfo1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, nodeInfo) types.QueryNodeClient); ok {
		r0 = rf(ctx, nodeInfo1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.QueryNodeClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, nodeInfo) error); ok {
		r1 = rf(ctx, nodeInfo1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShardClientManager_GetClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClient'
type MockShardClientManager_GetClient_Call struct {
	*mock.Call
}

// GetClient is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeInfo1 nodeInfo
func (_e *MockShardClientManager_Expecter) GetClient(ctx interface{}, nodeInfo1 interface{}) *MockShardClientManager_GetClient_Call {
	return &MockShardClientManager_GetClient_Call{Call: _e.mock.On("GetClient", ctx, nodeInfo1)}
}

func (_c *MockShardClientManager_GetClient_Call) Run(run func(ctx context.Context, nodeInfo1 nodeInfo)) *MockShardClientManager_GetClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(nodeInfo))
	})
	return _c
}

func (_c *MockShardClientManager_GetClient_Call) Return(_a0 types.QueryNodeClient, _a1 error) *MockShardClientManager_GetClient_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShardClientManager_GetClient_Call) RunAndReturn(run func(context.Context, nodeInfo) (types.QueryNodeClient, error)) *MockShardClientManager_GetClient_Call {
	_c.Call.Return(run)
	return _c
}

// SetClientCreatorFunc provides a mock function with given fields: creator
func (_m *MockShardClientManager) SetClientCreatorFunc(creator queryNodeCreatorFunc) {
	_m.Called(creator)
}

// MockShardClientManager_SetClientCreatorFunc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetClientCreatorFunc'
type MockShardClientManager_SetClientCreatorFunc_Call struct {
	*mock.Call
}

// SetClientCreatorFunc is a helper method to define mock.On call
//   - creator queryNodeCreatorFunc
func (_e *MockShardClientManager_Expecter) SetClientCreatorFunc(creator interface{}) *MockShardClientManager_SetClientCreatorFunc_Call {
	return &MockShardClientManager_SetClientCreatorFunc_Call{Call: _e.mock.On("SetClientCreatorFunc", creator)}
}

func (_c *MockShardClientManager_SetClientCreatorFunc_Call) Run(run func(creator queryNodeCreatorFunc)) *MockShardClientManager_SetClientCreatorFunc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(queryNodeCreatorFunc))
	})
	return _c
}

func (_c *MockShardClientManager_SetClientCreatorFunc_Call) Return() *MockShardClientManager_SetClientCreatorFunc_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockShardClientManager_SetClientCreatorFunc_Call) RunAndReturn(run func(queryNodeCreatorFunc)) *MockShardClientManager_SetClientCreatorFunc_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockShardClientManager creates a new instance of MockShardClientManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockShardClientManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockShardClientManager {
	mock := &MockShardClientManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
