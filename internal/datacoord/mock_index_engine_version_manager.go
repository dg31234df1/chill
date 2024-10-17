// Code generated by mockery v2.46.0. DO NOT EDIT.

package datacoord

import (
	sessionutil "github.com/milvus-io/milvus/internal/util/sessionutil"
	mock "github.com/stretchr/testify/mock"
)

// MockVersionManager is an autogenerated mock type for the IndexEngineVersionManager type
type MockVersionManager struct {
	mock.Mock
}

type MockVersionManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVersionManager) EXPECT() *MockVersionManager_Expecter {
	return &MockVersionManager_Expecter{mock: &_m.Mock}
}

// AddNode provides a mock function with given fields: session
func (_m *MockVersionManager) AddNode(session *sessionutil.Session) {
	_m.Called(session)
}

// MockVersionManager_AddNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNode'
type MockVersionManager_AddNode_Call struct {
	*mock.Call
}

// AddNode is a helper method to define mock.On call
//   - session *sessionutil.Session
func (_e *MockVersionManager_Expecter) AddNode(session interface{}) *MockVersionManager_AddNode_Call {
	return &MockVersionManager_AddNode_Call{Call: _e.mock.On("AddNode", session)}
}

func (_c *MockVersionManager_AddNode_Call) Run(run func(session *sessionutil.Session)) *MockVersionManager_AddNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*sessionutil.Session))
	})
	return _c
}

func (_c *MockVersionManager_AddNode_Call) Return() *MockVersionManager_AddNode_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockVersionManager_AddNode_Call) RunAndReturn(run func(*sessionutil.Session)) *MockVersionManager_AddNode_Call {
	_c.Call.Return(run)
	return _c
}

// GetCurrentIndexEngineVersion provides a mock function with given fields:
func (_m *MockVersionManager) GetCurrentIndexEngineVersion() int32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentIndexEngineVersion")
	}

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// MockVersionManager_GetCurrentIndexEngineVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentIndexEngineVersion'
type MockVersionManager_GetCurrentIndexEngineVersion_Call struct {
	*mock.Call
}

// GetCurrentIndexEngineVersion is a helper method to define mock.On call
func (_e *MockVersionManager_Expecter) GetCurrentIndexEngineVersion() *MockVersionManager_GetCurrentIndexEngineVersion_Call {
	return &MockVersionManager_GetCurrentIndexEngineVersion_Call{Call: _e.mock.On("GetCurrentIndexEngineVersion")}
}

func (_c *MockVersionManager_GetCurrentIndexEngineVersion_Call) Run(run func()) *MockVersionManager_GetCurrentIndexEngineVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVersionManager_GetCurrentIndexEngineVersion_Call) Return(_a0 int32) *MockVersionManager_GetCurrentIndexEngineVersion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVersionManager_GetCurrentIndexEngineVersion_Call) RunAndReturn(run func() int32) *MockVersionManager_GetCurrentIndexEngineVersion_Call {
	_c.Call.Return(run)
	return _c
}

// GetMinimalIndexEngineVersion provides a mock function with given fields:
func (_m *MockVersionManager) GetMinimalIndexEngineVersion() int32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMinimalIndexEngineVersion")
	}

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// MockVersionManager_GetMinimalIndexEngineVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMinimalIndexEngineVersion'
type MockVersionManager_GetMinimalIndexEngineVersion_Call struct {
	*mock.Call
}

// GetMinimalIndexEngineVersion is a helper method to define mock.On call
func (_e *MockVersionManager_Expecter) GetMinimalIndexEngineVersion() *MockVersionManager_GetMinimalIndexEngineVersion_Call {
	return &MockVersionManager_GetMinimalIndexEngineVersion_Call{Call: _e.mock.On("GetMinimalIndexEngineVersion")}
}

func (_c *MockVersionManager_GetMinimalIndexEngineVersion_Call) Run(run func()) *MockVersionManager_GetMinimalIndexEngineVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVersionManager_GetMinimalIndexEngineVersion_Call) Return(_a0 int32) *MockVersionManager_GetMinimalIndexEngineVersion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVersionManager_GetMinimalIndexEngineVersion_Call) RunAndReturn(run func() int32) *MockVersionManager_GetMinimalIndexEngineVersion_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveNode provides a mock function with given fields: session
func (_m *MockVersionManager) RemoveNode(session *sessionutil.Session) {
	_m.Called(session)
}

// MockVersionManager_RemoveNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveNode'
type MockVersionManager_RemoveNode_Call struct {
	*mock.Call
}

// RemoveNode is a helper method to define mock.On call
//   - session *sessionutil.Session
func (_e *MockVersionManager_Expecter) RemoveNode(session interface{}) *MockVersionManager_RemoveNode_Call {
	return &MockVersionManager_RemoveNode_Call{Call: _e.mock.On("RemoveNode", session)}
}

func (_c *MockVersionManager_RemoveNode_Call) Run(run func(session *sessionutil.Session)) *MockVersionManager_RemoveNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*sessionutil.Session))
	})
	return _c
}

func (_c *MockVersionManager_RemoveNode_Call) Return() *MockVersionManager_RemoveNode_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockVersionManager_RemoveNode_Call) RunAndReturn(run func(*sessionutil.Session)) *MockVersionManager_RemoveNode_Call {
	_c.Call.Return(run)
	return _c
}

// Startup provides a mock function with given fields: sessions
func (_m *MockVersionManager) Startup(sessions map[string]*sessionutil.Session) {
	_m.Called(sessions)
}

// MockVersionManager_Startup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Startup'
type MockVersionManager_Startup_Call struct {
	*mock.Call
}

// Startup is a helper method to define mock.On call
//   - sessions map[string]*sessionutil.Session
func (_e *MockVersionManager_Expecter) Startup(sessions interface{}) *MockVersionManager_Startup_Call {
	return &MockVersionManager_Startup_Call{Call: _e.mock.On("Startup", sessions)}
}

func (_c *MockVersionManager_Startup_Call) Run(run func(sessions map[string]*sessionutil.Session)) *MockVersionManager_Startup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]*sessionutil.Session))
	})
	return _c
}

func (_c *MockVersionManager_Startup_Call) Return() *MockVersionManager_Startup_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockVersionManager_Startup_Call) RunAndReturn(run func(map[string]*sessionutil.Session)) *MockVersionManager_Startup_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: session
func (_m *MockVersionManager) Update(session *sessionutil.Session) {
	_m.Called(session)
}

// MockVersionManager_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockVersionManager_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - session *sessionutil.Session
func (_e *MockVersionManager_Expecter) Update(session interface{}) *MockVersionManager_Update_Call {
	return &MockVersionManager_Update_Call{Call: _e.mock.On("Update", session)}
}

func (_c *MockVersionManager_Update_Call) Run(run func(session *sessionutil.Session)) *MockVersionManager_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*sessionutil.Session))
	})
	return _c
}

func (_c *MockVersionManager_Update_Call) Return() *MockVersionManager_Update_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockVersionManager_Update_Call) RunAndReturn(run func(*sessionutil.Session)) *MockVersionManager_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVersionManager creates a new instance of MockVersionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVersionManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVersionManager {
	mock := &MockVersionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
