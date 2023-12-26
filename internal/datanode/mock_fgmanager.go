// Code generated by mockery v2.32.4. DO NOT EDIT.

package datanode

import (
	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"

	schemapb "github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
)

// MockFlowgraphManager is an autogenerated mock type for the FlowgraphManager type
type MockFlowgraphManager struct {
	mock.Mock
}

type MockFlowgraphManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFlowgraphManager) EXPECT() *MockFlowgraphManager_Expecter {
	return &MockFlowgraphManager_Expecter{mock: &_m.Mock}
}

// AddFlowgraph provides a mock function with given fields: ds
func (_m *MockFlowgraphManager) AddFlowgraph(ds *dataSyncService) {
	_m.Called(ds)
}

// MockFlowgraphManager_AddFlowgraph_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFlowgraph'
type MockFlowgraphManager_AddFlowgraph_Call struct {
	*mock.Call
}

// AddFlowgraph is a helper method to define mock.On call
//   - ds *dataSyncService
func (_e *MockFlowgraphManager_Expecter) AddFlowgraph(ds interface{}) *MockFlowgraphManager_AddFlowgraph_Call {
	return &MockFlowgraphManager_AddFlowgraph_Call{Call: _e.mock.On("AddFlowgraph", ds)}
}

func (_c *MockFlowgraphManager_AddFlowgraph_Call) Run(run func(ds *dataSyncService)) *MockFlowgraphManager_AddFlowgraph_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*dataSyncService))
	})
	return _c
}

func (_c *MockFlowgraphManager_AddFlowgraph_Call) Return() *MockFlowgraphManager_AddFlowgraph_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFlowgraphManager_AddFlowgraph_Call) RunAndReturn(run func(*dataSyncService)) *MockFlowgraphManager_AddFlowgraph_Call {
	_c.Call.Return(run)
	return _c
}

// AddandStartWithEtcdTickler provides a mock function with given fields: dn, vchan, schema, tickler
func (_m *MockFlowgraphManager) AddandStartWithEtcdTickler(dn *DataNode, vchan *datapb.VchannelInfo, schema *schemapb.CollectionSchema, tickler *etcdTickler) error {
	ret := _m.Called(dn, vchan, schema, tickler)

	var r0 error
	if rf, ok := ret.Get(0).(func(*DataNode, *datapb.VchannelInfo, *schemapb.CollectionSchema, *etcdTickler) error); ok {
		r0 = rf(dn, vchan, schema, tickler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFlowgraphManager_AddandStartWithEtcdTickler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddandStartWithEtcdTickler'
type MockFlowgraphManager_AddandStartWithEtcdTickler_Call struct {
	*mock.Call
}

// AddandStartWithEtcdTickler is a helper method to define mock.On call
//   - dn *DataNode
//   - vchan *datapb.VchannelInfo
//   - schema *schemapb.CollectionSchema
//   - tickler *etcdTickler
func (_e *MockFlowgraphManager_Expecter) AddandStartWithEtcdTickler(dn interface{}, vchan interface{}, schema interface{}, tickler interface{}) *MockFlowgraphManager_AddandStartWithEtcdTickler_Call {
	return &MockFlowgraphManager_AddandStartWithEtcdTickler_Call{Call: _e.mock.On("AddandStartWithEtcdTickler", dn, vchan, schema, tickler)}
}

func (_c *MockFlowgraphManager_AddandStartWithEtcdTickler_Call) Run(run func(dn *DataNode, vchan *datapb.VchannelInfo, schema *schemapb.CollectionSchema, tickler *etcdTickler)) *MockFlowgraphManager_AddandStartWithEtcdTickler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*DataNode), args[1].(*datapb.VchannelInfo), args[2].(*schemapb.CollectionSchema), args[3].(*etcdTickler))
	})
	return _c
}

func (_c *MockFlowgraphManager_AddandStartWithEtcdTickler_Call) Return(_a0 error) *MockFlowgraphManager_AddandStartWithEtcdTickler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFlowgraphManager_AddandStartWithEtcdTickler_Call) RunAndReturn(run func(*DataNode, *datapb.VchannelInfo, *schemapb.CollectionSchema, *etcdTickler) error) *MockFlowgraphManager_AddandStartWithEtcdTickler_Call {
	_c.Call.Return(run)
	return _c
}

// ClearFlowgraphs provides a mock function with given fields:
func (_m *MockFlowgraphManager) ClearFlowgraphs() {
	_m.Called()
}

// MockFlowgraphManager_ClearFlowgraphs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearFlowgraphs'
type MockFlowgraphManager_ClearFlowgraphs_Call struct {
	*mock.Call
}

// ClearFlowgraphs is a helper method to define mock.On call
func (_e *MockFlowgraphManager_Expecter) ClearFlowgraphs() *MockFlowgraphManager_ClearFlowgraphs_Call {
	return &MockFlowgraphManager_ClearFlowgraphs_Call{Call: _e.mock.On("ClearFlowgraphs")}
}

func (_c *MockFlowgraphManager_ClearFlowgraphs_Call) Run(run func()) *MockFlowgraphManager_ClearFlowgraphs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFlowgraphManager_ClearFlowgraphs_Call) Return() *MockFlowgraphManager_ClearFlowgraphs_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFlowgraphManager_ClearFlowgraphs_Call) RunAndReturn(run func()) *MockFlowgraphManager_ClearFlowgraphs_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionIDs provides a mock function with given fields:
func (_m *MockFlowgraphManager) GetCollectionIDs() []int64 {
	ret := _m.Called()

	var r0 []int64
	if rf, ok := ret.Get(0).(func() []int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	return r0
}

// MockFlowgraphManager_GetCollectionIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionIDs'
type MockFlowgraphManager_GetCollectionIDs_Call struct {
	*mock.Call
}

// GetCollectionIDs is a helper method to define mock.On call
func (_e *MockFlowgraphManager_Expecter) GetCollectionIDs() *MockFlowgraphManager_GetCollectionIDs_Call {
	return &MockFlowgraphManager_GetCollectionIDs_Call{Call: _e.mock.On("GetCollectionIDs")}
}

func (_c *MockFlowgraphManager_GetCollectionIDs_Call) Run(run func()) *MockFlowgraphManager_GetCollectionIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFlowgraphManager_GetCollectionIDs_Call) Return(_a0 []int64) *MockFlowgraphManager_GetCollectionIDs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFlowgraphManager_GetCollectionIDs_Call) RunAndReturn(run func() []int64) *MockFlowgraphManager_GetCollectionIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetFlowgraphCount provides a mock function with given fields:
func (_m *MockFlowgraphManager) GetFlowgraphCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockFlowgraphManager_GetFlowgraphCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlowgraphCount'
type MockFlowgraphManager_GetFlowgraphCount_Call struct {
	*mock.Call
}

// GetFlowgraphCount is a helper method to define mock.On call
func (_e *MockFlowgraphManager_Expecter) GetFlowgraphCount() *MockFlowgraphManager_GetFlowgraphCount_Call {
	return &MockFlowgraphManager_GetFlowgraphCount_Call{Call: _e.mock.On("GetFlowgraphCount")}
}

func (_c *MockFlowgraphManager_GetFlowgraphCount_Call) Run(run func()) *MockFlowgraphManager_GetFlowgraphCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFlowgraphManager_GetFlowgraphCount_Call) Return(_a0 int) *MockFlowgraphManager_GetFlowgraphCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFlowgraphManager_GetFlowgraphCount_Call) RunAndReturn(run func() int) *MockFlowgraphManager_GetFlowgraphCount_Call {
	_c.Call.Return(run)
	return _c
}

// GetFlowgraphService provides a mock function with given fields: channel
func (_m *MockFlowgraphManager) GetFlowgraphService(channel string) (*dataSyncService, bool) {
	ret := _m.Called(channel)

	var r0 *dataSyncService
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (*dataSyncService, bool)); ok {
		return rf(channel)
	}
	if rf, ok := ret.Get(0).(func(string) *dataSyncService); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dataSyncService)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(channel)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockFlowgraphManager_GetFlowgraphService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlowgraphService'
type MockFlowgraphManager_GetFlowgraphService_Call struct {
	*mock.Call
}

// GetFlowgraphService is a helper method to define mock.On call
//   - channel string
func (_e *MockFlowgraphManager_Expecter) GetFlowgraphService(channel interface{}) *MockFlowgraphManager_GetFlowgraphService_Call {
	return &MockFlowgraphManager_GetFlowgraphService_Call{Call: _e.mock.On("GetFlowgraphService", channel)}
}

func (_c *MockFlowgraphManager_GetFlowgraphService_Call) Run(run func(channel string)) *MockFlowgraphManager_GetFlowgraphService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFlowgraphManager_GetFlowgraphService_Call) Return(_a0 *dataSyncService, _a1 bool) *MockFlowgraphManager_GetFlowgraphService_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFlowgraphManager_GetFlowgraphService_Call) RunAndReturn(run func(string) (*dataSyncService, bool)) *MockFlowgraphManager_GetFlowgraphService_Call {
	_c.Call.Return(run)
	return _c
}

// HasFlowgraph provides a mock function with given fields: channel
func (_m *MockFlowgraphManager) HasFlowgraph(channel string) bool {
	ret := _m.Called(channel)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(channel)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockFlowgraphManager_HasFlowgraph_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasFlowgraph'
type MockFlowgraphManager_HasFlowgraph_Call struct {
	*mock.Call
}

// HasFlowgraph is a helper method to define mock.On call
//   - channel string
func (_e *MockFlowgraphManager_Expecter) HasFlowgraph(channel interface{}) *MockFlowgraphManager_HasFlowgraph_Call {
	return &MockFlowgraphManager_HasFlowgraph_Call{Call: _e.mock.On("HasFlowgraph", channel)}
}

func (_c *MockFlowgraphManager_HasFlowgraph_Call) Run(run func(channel string)) *MockFlowgraphManager_HasFlowgraph_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFlowgraphManager_HasFlowgraph_Call) Return(_a0 bool) *MockFlowgraphManager_HasFlowgraph_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFlowgraphManager_HasFlowgraph_Call) RunAndReturn(run func(string) bool) *MockFlowgraphManager_HasFlowgraph_Call {
	_c.Call.Return(run)
	return _c
}

// HasFlowgraphWithOpID provides a mock function with given fields: channel, opID
func (_m *MockFlowgraphManager) HasFlowgraphWithOpID(channel string, opID int64) bool {
	ret := _m.Called(channel, opID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, int64) bool); ok {
		r0 = rf(channel, opID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockFlowgraphManager_HasFlowgraphWithOpID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasFlowgraphWithOpID'
type MockFlowgraphManager_HasFlowgraphWithOpID_Call struct {
	*mock.Call
}

// HasFlowgraphWithOpID is a helper method to define mock.On call
//   - channel string
//   - opID int64
func (_e *MockFlowgraphManager_Expecter) HasFlowgraphWithOpID(channel interface{}, opID interface{}) *MockFlowgraphManager_HasFlowgraphWithOpID_Call {
	return &MockFlowgraphManager_HasFlowgraphWithOpID_Call{Call: _e.mock.On("HasFlowgraphWithOpID", channel, opID)}
}

func (_c *MockFlowgraphManager_HasFlowgraphWithOpID_Call) Run(run func(channel string, opID int64)) *MockFlowgraphManager_HasFlowgraphWithOpID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64))
	})
	return _c
}

func (_c *MockFlowgraphManager_HasFlowgraphWithOpID_Call) Return(_a0 bool) *MockFlowgraphManager_HasFlowgraphWithOpID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFlowgraphManager_HasFlowgraphWithOpID_Call) RunAndReturn(run func(string, int64) bool) *MockFlowgraphManager_HasFlowgraphWithOpID_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveFlowgraph provides a mock function with given fields: channel
func (_m *MockFlowgraphManager) RemoveFlowgraph(channel string) {
	_m.Called(channel)
}

// MockFlowgraphManager_RemoveFlowgraph_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveFlowgraph'
type MockFlowgraphManager_RemoveFlowgraph_Call struct {
	*mock.Call
}

// RemoveFlowgraph is a helper method to define mock.On call
//   - channel string
func (_e *MockFlowgraphManager_Expecter) RemoveFlowgraph(channel interface{}) *MockFlowgraphManager_RemoveFlowgraph_Call {
	return &MockFlowgraphManager_RemoveFlowgraph_Call{Call: _e.mock.On("RemoveFlowgraph", channel)}
}

func (_c *MockFlowgraphManager_RemoveFlowgraph_Call) Run(run func(channel string)) *MockFlowgraphManager_RemoveFlowgraph_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFlowgraphManager_RemoveFlowgraph_Call) Return() *MockFlowgraphManager_RemoveFlowgraph_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFlowgraphManager_RemoveFlowgraph_Call) RunAndReturn(run func(string)) *MockFlowgraphManager_RemoveFlowgraph_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFlowgraphManager creates a new instance of MockFlowgraphManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFlowgraphManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFlowgraphManager {
	mock := &MockFlowgraphManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
