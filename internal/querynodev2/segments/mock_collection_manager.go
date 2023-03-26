// Code generated by mockery v2.15.0. DO NOT EDIT.

package segments

import (
	schemapb "github.com/milvus-io/milvus-proto/go-api/schemapb"
	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
	mock "github.com/stretchr/testify/mock"
)

// MockCollectionManager is an autogenerated mock type for the CollectionManager type
type MockCollectionManager struct {
	mock.Mock
}

type MockCollectionManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCollectionManager) EXPECT() *MockCollectionManager_Expecter {
	return &MockCollectionManager_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: collectionID
func (_m *MockCollectionManager) Get(collectionID int64) *Collection {
	ret := _m.Called(collectionID)

	var r0 *Collection
	if rf, ok := ret.Get(0).(func(int64) *Collection); ok {
		r0 = rf(collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Collection)
		}
	}

	return r0
}

// MockCollectionManager_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCollectionManager_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - collectionID int64
func (_e *MockCollectionManager_Expecter) Get(collectionID interface{}) *MockCollectionManager_Get_Call {
	return &MockCollectionManager_Get_Call{Call: _e.mock.On("Get", collectionID)}
}

func (_c *MockCollectionManager_Get_Call) Run(run func(collectionID int64)) *MockCollectionManager_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCollectionManager_Get_Call) Return(_a0 *Collection) *MockCollectionManager_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

// Put provides a mock function with given fields: collectionID, schema, loadMeta
func (_m *MockCollectionManager) Put(collectionID int64, schema *schemapb.CollectionSchema, loadMeta *querypb.LoadMetaInfo) {
	_m.Called(collectionID, schema, loadMeta)
}

// MockCollectionManager_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockCollectionManager_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - collectionID int64
//   - schema *schemapb.CollectionSchema
//   - loadMeta *querypb.LoadMetaInfo
func (_e *MockCollectionManager_Expecter) Put(collectionID interface{}, schema interface{}, loadMeta interface{}) *MockCollectionManager_Put_Call {
	return &MockCollectionManager_Put_Call{Call: _e.mock.On("Put", collectionID, schema, loadMeta)}
}

func (_c *MockCollectionManager_Put_Call) Run(run func(collectionID int64, schema *schemapb.CollectionSchema, loadMeta *querypb.LoadMetaInfo)) *MockCollectionManager_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*schemapb.CollectionSchema), args[2].(*querypb.LoadMetaInfo))
	})
	return _c
}

func (_c *MockCollectionManager_Put_Call) Return() *MockCollectionManager_Put_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockCollectionManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCollectionManager creates a new instance of MockCollectionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCollectionManager(t mockConstructorTestingTNewMockCollectionManager) *MockCollectionManager {
	mock := &MockCollectionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
