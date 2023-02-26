// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SnapShotKV is an autogenerated mock type for the SnapShotKV type
type SnapShotKV struct {
	mock.Mock
}

type SnapShotKV_Expecter struct {
	mock *mock.Mock
}

func (_m *SnapShotKV) EXPECT() *SnapShotKV_Expecter {
	return &SnapShotKV_Expecter{mock: &_m.Mock}
}

// Load provides a mock function with given fields: key, ts
func (_m *SnapShotKV) Load(key string, ts uint64) (string, error) {
	ret := _m.Called(key, ts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, uint64) string); ok {
		r0 = rf(key, ts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint64) error); ok {
		r1 = rf(key, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SnapShotKV_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type SnapShotKV_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - key string
//   - ts uint64
func (_e *SnapShotKV_Expecter) Load(key interface{}, ts interface{}) *SnapShotKV_Load_Call {
	return &SnapShotKV_Load_Call{Call: _e.mock.On("Load", key, ts)}
}

func (_c *SnapShotKV_Load_Call) Run(run func(key string, ts uint64)) *SnapShotKV_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint64))
	})
	return _c
}

func (_c *SnapShotKV_Load_Call) Return(_a0 string, _a1 error) *SnapShotKV_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadWithPrefix provides a mock function with given fields: key, ts
func (_m *SnapShotKV) LoadWithPrefix(key string, ts uint64) ([]string, []string, error) {
	ret := _m.Called(key, ts)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, uint64) []string); ok {
		r0 = rf(key, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string, uint64) []string); ok {
		r1 = rf(key, ts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, uint64) error); ok {
		r2 = rf(key, ts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SnapShotKV_LoadWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadWithPrefix'
type SnapShotKV_LoadWithPrefix_Call struct {
	*mock.Call
}

// LoadWithPrefix is a helper method to define mock.On call
//   - key string
//   - ts uint64
func (_e *SnapShotKV_Expecter) LoadWithPrefix(key interface{}, ts interface{}) *SnapShotKV_LoadWithPrefix_Call {
	return &SnapShotKV_LoadWithPrefix_Call{Call: _e.mock.On("LoadWithPrefix", key, ts)}
}

func (_c *SnapShotKV_LoadWithPrefix_Call) Run(run func(key string, ts uint64)) *SnapShotKV_LoadWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint64))
	})
	return _c
}

func (_c *SnapShotKV_LoadWithPrefix_Call) Return(_a0 []string, _a1 []string, _a2 error) *SnapShotKV_LoadWithPrefix_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

// MultiSave provides a mock function with given fields: kvs, ts
func (_m *SnapShotKV) MultiSave(kvs map[string]string, ts uint64) error {
	ret := _m.Called(kvs, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, uint64) error); ok {
		r0 = rf(kvs, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SnapShotKV_MultiSave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSave'
type SnapShotKV_MultiSave_Call struct {
	*mock.Call
}

// MultiSave is a helper method to define mock.On call
//   - kvs map[string]string
//   - ts uint64
func (_e *SnapShotKV_Expecter) MultiSave(kvs interface{}, ts interface{}) *SnapShotKV_MultiSave_Call {
	return &SnapShotKV_MultiSave_Call{Call: _e.mock.On("MultiSave", kvs, ts)}
}

func (_c *SnapShotKV_MultiSave_Call) Run(run func(kvs map[string]string, ts uint64)) *SnapShotKV_MultiSave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string), args[1].(uint64))
	})
	return _c
}

func (_c *SnapShotKV_MultiSave_Call) Return(_a0 error) *SnapShotKV_MultiSave_Call {
	_c.Call.Return(_a0)
	return _c
}

// MultiSaveAndRemoveWithPrefix provides a mock function with given fields: saves, removals, ts
func (_m *SnapShotKV) MultiSaveAndRemoveWithPrefix(saves map[string]string, removals []string, ts uint64) error {
	ret := _m.Called(saves, removals, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, []string, uint64) error); ok {
		r0 = rf(saves, removals, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SnapShotKV_MultiSaveAndRemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSaveAndRemoveWithPrefix'
type SnapShotKV_MultiSaveAndRemoveWithPrefix_Call struct {
	*mock.Call
}

// MultiSaveAndRemoveWithPrefix is a helper method to define mock.On call
//   - saves map[string]string
//   - removals []string
//   - ts uint64
func (_e *SnapShotKV_Expecter) MultiSaveAndRemoveWithPrefix(saves interface{}, removals interface{}, ts interface{}) *SnapShotKV_MultiSaveAndRemoveWithPrefix_Call {
	return &SnapShotKV_MultiSaveAndRemoveWithPrefix_Call{Call: _e.mock.On("MultiSaveAndRemoveWithPrefix", saves, removals, ts)}
}

func (_c *SnapShotKV_MultiSaveAndRemoveWithPrefix_Call) Run(run func(saves map[string]string, removals []string, ts uint64)) *SnapShotKV_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string), args[1].([]string), args[2].(uint64))
	})
	return _c
}

func (_c *SnapShotKV_MultiSaveAndRemoveWithPrefix_Call) Return(_a0 error) *SnapShotKV_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

// Save provides a mock function with given fields: key, value, ts
func (_m *SnapShotKV) Save(key string, value string, ts uint64) error {
	ret := _m.Called(key, value, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, uint64) error); ok {
		r0 = rf(key, value, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SnapShotKV_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type SnapShotKV_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - key string
//   - value string
//   - ts uint64
func (_e *SnapShotKV_Expecter) Save(key interface{}, value interface{}, ts interface{}) *SnapShotKV_Save_Call {
	return &SnapShotKV_Save_Call{Call: _e.mock.On("Save", key, value, ts)}
}

func (_c *SnapShotKV_Save_Call) Run(run func(key string, value string, ts uint64)) *SnapShotKV_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *SnapShotKV_Save_Call) Return(_a0 error) *SnapShotKV_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewSnapShotKV interface {
	mock.TestingT
	Cleanup(func())
}

// NewSnapShotKV creates a new instance of SnapShotKV. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSnapShotKV(t mockConstructorTestingTNewSnapShotKV) *SnapShotKV {
	mock := &SnapShotKV{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
