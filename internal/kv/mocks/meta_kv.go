// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	predicates "github.com/milvus-io/milvus/pkg/kv/predicates"
	mock "github.com/stretchr/testify/mock"
)

// MetaKv is an autogenerated mock type for the MetaKv type
type MetaKv struct {
	mock.Mock
}

type MetaKv_Expecter struct {
	mock *mock.Mock
}

func (_m *MetaKv) EXPECT() *MetaKv_Expecter {
	return &MetaKv_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MetaKv) Close() {
	_m.Called()
}

// MetaKv_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MetaKv_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MetaKv_Expecter) Close() *MetaKv_Close_Call {
	return &MetaKv_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MetaKv_Close_Call) Run(run func()) *MetaKv_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MetaKv_Close_Call) Return() *MetaKv_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MetaKv_Close_Call) RunAndReturn(run func()) *MetaKv_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CompareVersionAndSwap provides a mock function with given fields: key, version, target
func (_m *MetaKv) CompareVersionAndSwap(key string, version int64, target string) (bool, error) {
	ret := _m.Called(key, version, target)

	if len(ret) == 0 {
		panic("no return value specified for CompareVersionAndSwap")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64, string) (bool, error)); ok {
		return rf(key, version, target)
	}
	if rf, ok := ret.Get(0).(func(string, int64, string) bool); ok {
		r0 = rf(key, version, target)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, int64, string) error); ok {
		r1 = rf(key, version, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_CompareVersionAndSwap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompareVersionAndSwap'
type MetaKv_CompareVersionAndSwap_Call struct {
	*mock.Call
}

// CompareVersionAndSwap is a helper method to define mock.On call
//   - key string
//   - version int64
//   - target string
func (_e *MetaKv_Expecter) CompareVersionAndSwap(key interface{}, version interface{}, target interface{}) *MetaKv_CompareVersionAndSwap_Call {
	return &MetaKv_CompareVersionAndSwap_Call{Call: _e.mock.On("CompareVersionAndSwap", key, version, target)}
}

func (_c *MetaKv_CompareVersionAndSwap_Call) Run(run func(key string, version int64, target string)) *MetaKv_CompareVersionAndSwap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MetaKv_CompareVersionAndSwap_Call) Return(_a0 bool, _a1 error) *MetaKv_CompareVersionAndSwap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetaKv_CompareVersionAndSwap_Call) RunAndReturn(run func(string, int64, string) (bool, error)) *MetaKv_CompareVersionAndSwap_Call {
	_c.Call.Return(run)
	return _c
}

// GetPath provides a mock function with given fields: key
func (_m *MetaKv) GetPath(key string) string {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetPath")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MetaKv_GetPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPath'
type MetaKv_GetPath_Call struct {
	*mock.Call
}

// GetPath is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) GetPath(key interface{}) *MetaKv_GetPath_Call {
	return &MetaKv_GetPath_Call{Call: _e.mock.On("GetPath", key)}
}

func (_c *MetaKv_GetPath_Call) Run(run func(key string)) *MetaKv_GetPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_GetPath_Call) Return(_a0 string) *MetaKv_GetPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_GetPath_Call) RunAndReturn(run func(string) string) *MetaKv_GetPath_Call {
	_c.Call.Return(run)
	return _c
}

// Has provides a mock function with given fields: key
func (_m *MetaKv) Has(key string) (bool, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Has")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_Has_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Has'
type MetaKv_Has_Call struct {
	*mock.Call
}

// Has is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) Has(key interface{}) *MetaKv_Has_Call {
	return &MetaKv_Has_Call{Call: _e.mock.On("Has", key)}
}

func (_c *MetaKv_Has_Call) Run(run func(key string)) *MetaKv_Has_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_Has_Call) Return(_a0 bool, _a1 error) *MetaKv_Has_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetaKv_Has_Call) RunAndReturn(run func(string) (bool, error)) *MetaKv_Has_Call {
	_c.Call.Return(run)
	return _c
}

// HasPrefix provides a mock function with given fields: prefix
func (_m *MetaKv) HasPrefix(prefix string) (bool, error) {
	ret := _m.Called(prefix)

	if len(ret) == 0 {
		panic("no return value specified for HasPrefix")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(prefix)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(prefix)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_HasPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasPrefix'
type MetaKv_HasPrefix_Call struct {
	*mock.Call
}

// HasPrefix is a helper method to define mock.On call
//   - prefix string
func (_e *MetaKv_Expecter) HasPrefix(prefix interface{}) *MetaKv_HasPrefix_Call {
	return &MetaKv_HasPrefix_Call{Call: _e.mock.On("HasPrefix", prefix)}
}

func (_c *MetaKv_HasPrefix_Call) Run(run func(prefix string)) *MetaKv_HasPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_HasPrefix_Call) Return(_a0 bool, _a1 error) *MetaKv_HasPrefix_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetaKv_HasPrefix_Call) RunAndReturn(run func(string) (bool, error)) *MetaKv_HasPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Load provides a mock function with given fields: key
func (_m *MetaKv) Load(key string) (string, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type MetaKv_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) Load(key interface{}) *MetaKv_Load_Call {
	return &MetaKv_Load_Call{Call: _e.mock.On("Load", key)}
}

func (_c *MetaKv_Load_Call) Run(run func(key string)) *MetaKv_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_Load_Call) Return(_a0 string, _a1 error) *MetaKv_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetaKv_Load_Call) RunAndReturn(run func(string) (string, error)) *MetaKv_Load_Call {
	_c.Call.Return(run)
	return _c
}

// LoadWithPrefix provides a mock function with given fields: key
func (_m *MetaKv) LoadWithPrefix(key string) ([]string, []string, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for LoadWithPrefix")
	}

	var r0 []string
	var r1 []string
	var r2 error
	if rf, ok := ret.Get(0).(func(string) ([]string, []string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) []string); ok {
		r1 = rf(key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MetaKv_LoadWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadWithPrefix'
type MetaKv_LoadWithPrefix_Call struct {
	*mock.Call
}

// LoadWithPrefix is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) LoadWithPrefix(key interface{}) *MetaKv_LoadWithPrefix_Call {
	return &MetaKv_LoadWithPrefix_Call{Call: _e.mock.On("LoadWithPrefix", key)}
}

func (_c *MetaKv_LoadWithPrefix_Call) Run(run func(key string)) *MetaKv_LoadWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_LoadWithPrefix_Call) Return(_a0 []string, _a1 []string, _a2 error) *MetaKv_LoadWithPrefix_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MetaKv_LoadWithPrefix_Call) RunAndReturn(run func(string) ([]string, []string, error)) *MetaKv_LoadWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// MultiLoad provides a mock function with given fields: keys
func (_m *MetaKv) MultiLoad(keys []string) ([]string, error) {
	ret := _m.Called(keys)

	if len(ret) == 0 {
		panic("no return value specified for MultiLoad")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]string, error)); ok {
		return rf(keys)
	}
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_MultiLoad_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiLoad'
type MetaKv_MultiLoad_Call struct {
	*mock.Call
}

// MultiLoad is a helper method to define mock.On call
//   - keys []string
func (_e *MetaKv_Expecter) MultiLoad(keys interface{}) *MetaKv_MultiLoad_Call {
	return &MetaKv_MultiLoad_Call{Call: _e.mock.On("MultiLoad", keys)}
}

func (_c *MetaKv_MultiLoad_Call) Run(run func(keys []string)) *MetaKv_MultiLoad_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MetaKv_MultiLoad_Call) Return(_a0 []string, _a1 error) *MetaKv_MultiLoad_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetaKv_MultiLoad_Call) RunAndReturn(run func([]string) ([]string, error)) *MetaKv_MultiLoad_Call {
	_c.Call.Return(run)
	return _c
}

// MultiRemove provides a mock function with given fields: keys
func (_m *MetaKv) MultiRemove(keys []string) error {
	ret := _m.Called(keys)

	if len(ret) == 0 {
		panic("no return value specified for MultiRemove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiRemove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiRemove'
type MetaKv_MultiRemove_Call struct {
	*mock.Call
}

// MultiRemove is a helper method to define mock.On call
//   - keys []string
func (_e *MetaKv_Expecter) MultiRemove(keys interface{}) *MetaKv_MultiRemove_Call {
	return &MetaKv_MultiRemove_Call{Call: _e.mock.On("MultiRemove", keys)}
}

func (_c *MetaKv_MultiRemove_Call) Run(run func(keys []string)) *MetaKv_MultiRemove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MetaKv_MultiRemove_Call) Return(_a0 error) *MetaKv_MultiRemove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_MultiRemove_Call) RunAndReturn(run func([]string) error) *MetaKv_MultiRemove_Call {
	_c.Call.Return(run)
	return _c
}

// MultiSave provides a mock function with given fields: kvs
func (_m *MetaKv) MultiSave(kvs map[string]string) error {
	ret := _m.Called(kvs)

	if len(ret) == 0 {
		panic("no return value specified for MultiSave")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(kvs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiSave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSave'
type MetaKv_MultiSave_Call struct {
	*mock.Call
}

// MultiSave is a helper method to define mock.On call
//   - kvs map[string]string
func (_e *MetaKv_Expecter) MultiSave(kvs interface{}) *MetaKv_MultiSave_Call {
	return &MetaKv_MultiSave_Call{Call: _e.mock.On("MultiSave", kvs)}
}

func (_c *MetaKv_MultiSave_Call) Run(run func(kvs map[string]string)) *MetaKv_MultiSave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string))
	})
	return _c
}

func (_c *MetaKv_MultiSave_Call) Return(_a0 error) *MetaKv_MultiSave_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_MultiSave_Call) RunAndReturn(run func(map[string]string) error) *MetaKv_MultiSave_Call {
	_c.Call.Return(run)
	return _c
}

// MultiSaveAndRemove provides a mock function with given fields: saves, removals, preds
func (_m *MetaKv) MultiSaveAndRemove(saves map[string]string, removals []string, preds ...predicates.Predicate) error {
	_va := make([]interface{}, len(preds))
	for _i := range preds {
		_va[_i] = preds[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, saves, removals)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MultiSaveAndRemove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, []string, ...predicates.Predicate) error); ok {
		r0 = rf(saves, removals, preds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiSaveAndRemove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSaveAndRemove'
type MetaKv_MultiSaveAndRemove_Call struct {
	*mock.Call
}

// MultiSaveAndRemove is a helper method to define mock.On call
//   - saves map[string]string
//   - removals []string
//   - preds ...predicates.Predicate
func (_e *MetaKv_Expecter) MultiSaveAndRemove(saves interface{}, removals interface{}, preds ...interface{}) *MetaKv_MultiSaveAndRemove_Call {
	return &MetaKv_MultiSaveAndRemove_Call{Call: _e.mock.On("MultiSaveAndRemove",
		append([]interface{}{saves, removals}, preds...)...)}
}

func (_c *MetaKv_MultiSaveAndRemove_Call) Run(run func(saves map[string]string, removals []string, preds ...predicates.Predicate)) *MetaKv_MultiSaveAndRemove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]predicates.Predicate, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(predicates.Predicate)
			}
		}
		run(args[0].(map[string]string), args[1].([]string), variadicArgs...)
	})
	return _c
}

func (_c *MetaKv_MultiSaveAndRemove_Call) Return(_a0 error) *MetaKv_MultiSaveAndRemove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_MultiSaveAndRemove_Call) RunAndReturn(run func(map[string]string, []string, ...predicates.Predicate) error) *MetaKv_MultiSaveAndRemove_Call {
	_c.Call.Return(run)
	return _c
}

// MultiSaveAndRemoveWithPrefix provides a mock function with given fields: saves, removals, preds
func (_m *MetaKv) MultiSaveAndRemoveWithPrefix(saves map[string]string, removals []string, preds ...predicates.Predicate) error {
	_va := make([]interface{}, len(preds))
	for _i := range preds {
		_va[_i] = preds[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, saves, removals)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MultiSaveAndRemoveWithPrefix")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, []string, ...predicates.Predicate) error); ok {
		r0 = rf(saves, removals, preds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiSaveAndRemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSaveAndRemoveWithPrefix'
type MetaKv_MultiSaveAndRemoveWithPrefix_Call struct {
	*mock.Call
}

// MultiSaveAndRemoveWithPrefix is a helper method to define mock.On call
//   - saves map[string]string
//   - removals []string
//   - preds ...predicates.Predicate
func (_e *MetaKv_Expecter) MultiSaveAndRemoveWithPrefix(saves interface{}, removals interface{}, preds ...interface{}) *MetaKv_MultiSaveAndRemoveWithPrefix_Call {
	return &MetaKv_MultiSaveAndRemoveWithPrefix_Call{Call: _e.mock.On("MultiSaveAndRemoveWithPrefix",
		append([]interface{}{saves, removals}, preds...)...)}
}

func (_c *MetaKv_MultiSaveAndRemoveWithPrefix_Call) Run(run func(saves map[string]string, removals []string, preds ...predicates.Predicate)) *MetaKv_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]predicates.Predicate, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(predicates.Predicate)
			}
		}
		run(args[0].(map[string]string), args[1].([]string), variadicArgs...)
	})
	return _c
}

func (_c *MetaKv_MultiSaveAndRemoveWithPrefix_Call) Return(_a0 error) *MetaKv_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_MultiSaveAndRemoveWithPrefix_Call) RunAndReturn(run func(map[string]string, []string, ...predicates.Predicate) error) *MetaKv_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: key
func (_m *MetaKv) Remove(key string) error {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MetaKv_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) Remove(key interface{}) *MetaKv_Remove_Call {
	return &MetaKv_Remove_Call{Call: _e.mock.On("Remove", key)}
}

func (_c *MetaKv_Remove_Call) Run(run func(key string)) *MetaKv_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_Remove_Call) Return(_a0 error) *MetaKv_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_Remove_Call) RunAndReturn(run func(string) error) *MetaKv_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveWithPrefix provides a mock function with given fields: key
func (_m *MetaKv) RemoveWithPrefix(key string) error {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for RemoveWithPrefix")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_RemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveWithPrefix'
type MetaKv_RemoveWithPrefix_Call struct {
	*mock.Call
}

// RemoveWithPrefix is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) RemoveWithPrefix(key interface{}) *MetaKv_RemoveWithPrefix_Call {
	return &MetaKv_RemoveWithPrefix_Call{Call: _e.mock.On("RemoveWithPrefix", key)}
}

func (_c *MetaKv_RemoveWithPrefix_Call) Run(run func(key string)) *MetaKv_RemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_RemoveWithPrefix_Call) Return(_a0 error) *MetaKv_RemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_RemoveWithPrefix_Call) RunAndReturn(run func(string) error) *MetaKv_RemoveWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: key, value
func (_m *MetaKv) Save(key string, value string) error {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MetaKv_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *MetaKv_Expecter) Save(key interface{}, value interface{}) *MetaKv_Save_Call {
	return &MetaKv_Save_Call{Call: _e.mock.On("Save", key, value)}
}

func (_c *MetaKv_Save_Call) Run(run func(key string, value string)) *MetaKv_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MetaKv_Save_Call) Return(_a0 error) *MetaKv_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_Save_Call) RunAndReturn(run func(string, string) error) *MetaKv_Save_Call {
	_c.Call.Return(run)
	return _c
}

// WalkWithPrefix provides a mock function with given fields: prefix, paginationSize, fn
func (_m *MetaKv) WalkWithPrefix(prefix string, paginationSize int, fn func([]byte, []byte) error) error {
	ret := _m.Called(prefix, paginationSize, fn)

	if len(ret) == 0 {
		panic("no return value specified for WalkWithPrefix")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, func([]byte, []byte) error) error); ok {
		r0 = rf(prefix, paginationSize, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_WalkWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WalkWithPrefix'
type MetaKv_WalkWithPrefix_Call struct {
	*mock.Call
}

// WalkWithPrefix is a helper method to define mock.On call
//   - prefix string
//   - paginationSize int
//   - fn func([]byte , []byte) error
func (_e *MetaKv_Expecter) WalkWithPrefix(prefix interface{}, paginationSize interface{}, fn interface{}) *MetaKv_WalkWithPrefix_Call {
	return &MetaKv_WalkWithPrefix_Call{Call: _e.mock.On("WalkWithPrefix", prefix, paginationSize, fn)}
}

func (_c *MetaKv_WalkWithPrefix_Call) Run(run func(prefix string, paginationSize int, fn func([]byte, []byte) error)) *MetaKv_WalkWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int), args[2].(func([]byte, []byte) error))
	})
	return _c
}

func (_c *MetaKv_WalkWithPrefix_Call) Return(_a0 error) *MetaKv_WalkWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaKv_WalkWithPrefix_Call) RunAndReturn(run func(string, int, func([]byte, []byte) error) error) *MetaKv_WalkWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// NewMetaKv creates a new instance of MetaKv. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetaKv(t interface {
	mock.TestingT
	Cleanup(func())
}) *MetaKv {
	mock := &MetaKv{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
