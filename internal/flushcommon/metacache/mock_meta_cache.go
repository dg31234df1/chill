// Code generated by mockery v2.32.4. DO NOT EDIT.

package metacache

import (
	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"

	pkoracle "github.com/milvus-io/milvus/internal/flushcommon/metacache/pkoracle"

	schemapb "github.com/milvus-io/milvus-proto/go-api/v2/schemapb"

	storage "github.com/milvus-io/milvus/internal/storage"
)

// MockMetaCache is an autogenerated mock type for the MetaCache type
type MockMetaCache struct {
	mock.Mock
}

type MockMetaCache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetaCache) EXPECT() *MockMetaCache_Expecter {
	return &MockMetaCache_Expecter{mock: &_m.Mock}
}

// AddSegment provides a mock function with given fields: segInfo, pkFactory, bmFactory, actions
func (_m *MockMetaCache) AddSegment(segInfo *datapb.SegmentInfo, pkFactory PkStatsFactory, bmFactory BM25StatsFactory, actions ...SegmentAction) {
	_va := make([]interface{}, len(actions))
	for _i := range actions {
		_va[_i] = actions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, segInfo, pkFactory, bmFactory)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockMetaCache_AddSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSegment'
type MockMetaCache_AddSegment_Call struct {
	*mock.Call
}

// AddSegment is a helper method to define mock.On call
//   - segInfo *datapb.SegmentInfo
//   - pkFactory PkStatsFactory
//   - bmFactory BM25StatsFactory
//   - actions ...SegmentAction
func (_e *MockMetaCache_Expecter) AddSegment(segInfo interface{}, pkFactory interface{}, bmFactory interface{}, actions ...interface{}) *MockMetaCache_AddSegment_Call {
	return &MockMetaCache_AddSegment_Call{Call: _e.mock.On("AddSegment",
		append([]interface{}{segInfo, pkFactory, bmFactory}, actions...)...)}
}

func (_c *MockMetaCache_AddSegment_Call) Run(run func(segInfo *datapb.SegmentInfo, pkFactory PkStatsFactory, bmFactory BM25StatsFactory, actions ...SegmentAction)) *MockMetaCache_AddSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentAction, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentAction)
			}
		}
		run(args[0].(*datapb.SegmentInfo), args[1].(PkStatsFactory), args[2].(BM25StatsFactory), variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_AddSegment_Call) Return() *MockMetaCache_AddSegment_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMetaCache_AddSegment_Call) RunAndReturn(run func(*datapb.SegmentInfo, PkStatsFactory, BM25StatsFactory, ...SegmentAction)) *MockMetaCache_AddSegment_Call {
	_c.Call.Return(run)
	return _c
}

// Collection provides a mock function with given fields:
func (_m *MockMetaCache) Collection() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockMetaCache_Collection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Collection'
type MockMetaCache_Collection_Call struct {
	*mock.Call
}

// Collection is a helper method to define mock.On call
func (_e *MockMetaCache_Expecter) Collection() *MockMetaCache_Collection_Call {
	return &MockMetaCache_Collection_Call{Call: _e.mock.On("Collection")}
}

func (_c *MockMetaCache_Collection_Call) Run(run func()) *MockMetaCache_Collection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMetaCache_Collection_Call) Return(_a0 int64) *MockMetaCache_Collection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaCache_Collection_Call) RunAndReturn(run func() int64) *MockMetaCache_Collection_Call {
	_c.Call.Return(run)
	return _c
}

// DetectMissingSegments provides a mock function with given fields: segments
func (_m *MockMetaCache) DetectMissingSegments(segments map[int64]struct{}) []int64 {
	ret := _m.Called(segments)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(map[int64]struct{}) []int64); ok {
		r0 = rf(segments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	return r0
}

// MockMetaCache_DetectMissingSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DetectMissingSegments'
type MockMetaCache_DetectMissingSegments_Call struct {
	*mock.Call
}

// DetectMissingSegments is a helper method to define mock.On call
//   - segments map[int64]struct{}
func (_e *MockMetaCache_Expecter) DetectMissingSegments(segments interface{}) *MockMetaCache_DetectMissingSegments_Call {
	return &MockMetaCache_DetectMissingSegments_Call{Call: _e.mock.On("DetectMissingSegments", segments)}
}

func (_c *MockMetaCache_DetectMissingSegments_Call) Run(run func(segments map[int64]struct{})) *MockMetaCache_DetectMissingSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[int64]struct{}))
	})
	return _c
}

func (_c *MockMetaCache_DetectMissingSegments_Call) Return(_a0 []int64) *MockMetaCache_DetectMissingSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaCache_DetectMissingSegments_Call) RunAndReturn(run func(map[int64]struct{}) []int64) *MockMetaCache_DetectMissingSegments_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegmentByID provides a mock function with given fields: id, filters
func (_m *MockMetaCache) GetSegmentByID(id int64, filters ...SegmentFilter) (*SegmentInfo, bool) {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *SegmentInfo
	var r1 bool
	if rf, ok := ret.Get(0).(func(int64, ...SegmentFilter) (*SegmentInfo, bool)); ok {
		return rf(id, filters...)
	}
	if rf, ok := ret.Get(0).(func(int64, ...SegmentFilter) *SegmentInfo); ok {
		r0 = rf(id, filters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SegmentInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, ...SegmentFilter) bool); ok {
		r1 = rf(id, filters...)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockMetaCache_GetSegmentByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentByID'
type MockMetaCache_GetSegmentByID_Call struct {
	*mock.Call
}

// GetSegmentByID is a helper method to define mock.On call
//   - id int64
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) GetSegmentByID(id interface{}, filters ...interface{}) *MockMetaCache_GetSegmentByID_Call {
	return &MockMetaCache_GetSegmentByID_Call{Call: _e.mock.On("GetSegmentByID",
		append([]interface{}{id}, filters...)...)}
}

func (_c *MockMetaCache_GetSegmentByID_Call) Run(run func(id int64, filters ...SegmentFilter)) *MockMetaCache_GetSegmentByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(args[0].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_GetSegmentByID_Call) Return(_a0 *SegmentInfo, _a1 bool) *MockMetaCache_GetSegmentByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetaCache_GetSegmentByID_Call) RunAndReturn(run func(int64, ...SegmentFilter) (*SegmentInfo, bool)) *MockMetaCache_GetSegmentByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegmentIDsBy provides a mock function with given fields: filters
func (_m *MockMetaCache) GetSegmentIDsBy(filters ...SegmentFilter) []int64 {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(...SegmentFilter) []int64); ok {
		r0 = rf(filters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	return r0
}

// MockMetaCache_GetSegmentIDsBy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentIDsBy'
type MockMetaCache_GetSegmentIDsBy_Call struct {
	*mock.Call
}

// GetSegmentIDsBy is a helper method to define mock.On call
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) GetSegmentIDsBy(filters ...interface{}) *MockMetaCache_GetSegmentIDsBy_Call {
	return &MockMetaCache_GetSegmentIDsBy_Call{Call: _e.mock.On("GetSegmentIDsBy",
		append([]interface{}{}, filters...)...)}
}

func (_c *MockMetaCache_GetSegmentIDsBy_Call) Run(run func(filters ...SegmentFilter)) *MockMetaCache_GetSegmentIDsBy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_GetSegmentIDsBy_Call) Return(_a0 []int64) *MockMetaCache_GetSegmentIDsBy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaCache_GetSegmentIDsBy_Call) RunAndReturn(run func(...SegmentFilter) []int64) *MockMetaCache_GetSegmentIDsBy_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegmentsBy provides a mock function with given fields: filters
func (_m *MockMetaCache) GetSegmentsBy(filters ...SegmentFilter) []*SegmentInfo {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*SegmentInfo
	if rf, ok := ret.Get(0).(func(...SegmentFilter) []*SegmentInfo); ok {
		r0 = rf(filters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*SegmentInfo)
		}
	}

	return r0
}

// MockMetaCache_GetSegmentsBy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentsBy'
type MockMetaCache_GetSegmentsBy_Call struct {
	*mock.Call
}

// GetSegmentsBy is a helper method to define mock.On call
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) GetSegmentsBy(filters ...interface{}) *MockMetaCache_GetSegmentsBy_Call {
	return &MockMetaCache_GetSegmentsBy_Call{Call: _e.mock.On("GetSegmentsBy",
		append([]interface{}{}, filters...)...)}
}

func (_c *MockMetaCache_GetSegmentsBy_Call) Run(run func(filters ...SegmentFilter)) *MockMetaCache_GetSegmentsBy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_GetSegmentsBy_Call) Return(_a0 []*SegmentInfo) *MockMetaCache_GetSegmentsBy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaCache_GetSegmentsBy_Call) RunAndReturn(run func(...SegmentFilter) []*SegmentInfo) *MockMetaCache_GetSegmentsBy_Call {
	_c.Call.Return(run)
	return _c
}

// PredictSegments provides a mock function with given fields: pk, filters
func (_m *MockMetaCache) PredictSegments(pk storage.PrimaryKey, filters ...SegmentFilter) ([]int64, bool) {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pk)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []int64
	var r1 bool
	if rf, ok := ret.Get(0).(func(storage.PrimaryKey, ...SegmentFilter) ([]int64, bool)); ok {
		return rf(pk, filters...)
	}
	if rf, ok := ret.Get(0).(func(storage.PrimaryKey, ...SegmentFilter) []int64); ok {
		r0 = rf(pk, filters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(storage.PrimaryKey, ...SegmentFilter) bool); ok {
		r1 = rf(pk, filters...)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockMetaCache_PredictSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PredictSegments'
type MockMetaCache_PredictSegments_Call struct {
	*mock.Call
}

// PredictSegments is a helper method to define mock.On call
//   - pk storage.PrimaryKey
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) PredictSegments(pk interface{}, filters ...interface{}) *MockMetaCache_PredictSegments_Call {
	return &MockMetaCache_PredictSegments_Call{Call: _e.mock.On("PredictSegments",
		append([]interface{}{pk}, filters...)...)}
}

func (_c *MockMetaCache_PredictSegments_Call) Run(run func(pk storage.PrimaryKey, filters ...SegmentFilter)) *MockMetaCache_PredictSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(args[0].(storage.PrimaryKey), variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_PredictSegments_Call) Return(_a0 []int64, _a1 bool) *MockMetaCache_PredictSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetaCache_PredictSegments_Call) RunAndReturn(run func(storage.PrimaryKey, ...SegmentFilter) ([]int64, bool)) *MockMetaCache_PredictSegments_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveSegments provides a mock function with given fields: filters
func (_m *MockMetaCache) RemoveSegments(filters ...SegmentFilter) []int64 {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(...SegmentFilter) []int64); ok {
		r0 = rf(filters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	return r0
}

// MockMetaCache_RemoveSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveSegments'
type MockMetaCache_RemoveSegments_Call struct {
	*mock.Call
}

// RemoveSegments is a helper method to define mock.On call
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) RemoveSegments(filters ...interface{}) *MockMetaCache_RemoveSegments_Call {
	return &MockMetaCache_RemoveSegments_Call{Call: _e.mock.On("RemoveSegments",
		append([]interface{}{}, filters...)...)}
}

func (_c *MockMetaCache_RemoveSegments_Call) Run(run func(filters ...SegmentFilter)) *MockMetaCache_RemoveSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_RemoveSegments_Call) Return(_a0 []int64) *MockMetaCache_RemoveSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaCache_RemoveSegments_Call) RunAndReturn(run func(...SegmentFilter) []int64) *MockMetaCache_RemoveSegments_Call {
	_c.Call.Return(run)
	return _c
}

// Schema provides a mock function with given fields:
func (_m *MockMetaCache) Schema() *schemapb.CollectionSchema {
	ret := _m.Called()

	var r0 *schemapb.CollectionSchema
	if rf, ok := ret.Get(0).(func() *schemapb.CollectionSchema); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schemapb.CollectionSchema)
		}
	}

	return r0
}

// MockMetaCache_Schema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Schema'
type MockMetaCache_Schema_Call struct {
	*mock.Call
}

// Schema is a helper method to define mock.On call
func (_e *MockMetaCache_Expecter) Schema() *MockMetaCache_Schema_Call {
	return &MockMetaCache_Schema_Call{Call: _e.mock.On("Schema")}
}

func (_c *MockMetaCache_Schema_Call) Run(run func()) *MockMetaCache_Schema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMetaCache_Schema_Call) Return(_a0 *schemapb.CollectionSchema) *MockMetaCache_Schema_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaCache_Schema_Call) RunAndReturn(run func() *schemapb.CollectionSchema) *MockMetaCache_Schema_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSegmentView provides a mock function with given fields: partitionID, newSegments, newSegmentsBF, allSegments
func (_m *MockMetaCache) UpdateSegmentView(partitionID int64, newSegments []*datapb.SyncSegmentInfo, newSegmentsBF []*pkoracle.BloomFilterSet, allSegments map[int64]struct{}) {
	_m.Called(partitionID, newSegments, newSegmentsBF, allSegments)
}

// MockMetaCache_UpdateSegmentView_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSegmentView'
type MockMetaCache_UpdateSegmentView_Call struct {
	*mock.Call
}

// UpdateSegmentView is a helper method to define mock.On call
//   - partitionID int64
//   - newSegments []*datapb.SyncSegmentInfo
//   - newSegmentsBF []*pkoracle.BloomFilterSet
//   - allSegments map[int64]struct{}
func (_e *MockMetaCache_Expecter) UpdateSegmentView(partitionID interface{}, newSegments interface{}, newSegmentsBF interface{}, allSegments interface{}) *MockMetaCache_UpdateSegmentView_Call {
	return &MockMetaCache_UpdateSegmentView_Call{Call: _e.mock.On("UpdateSegmentView", partitionID, newSegments, newSegmentsBF, allSegments)}
}

func (_c *MockMetaCache_UpdateSegmentView_Call) Run(run func(partitionID int64, newSegments []*datapb.SyncSegmentInfo, newSegmentsBF []*pkoracle.BloomFilterSet, allSegments map[int64]struct{})) *MockMetaCache_UpdateSegmentView_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].([]*datapb.SyncSegmentInfo), args[2].([]*pkoracle.BloomFilterSet), args[3].(map[int64]struct{}))
	})
	return _c
}

func (_c *MockMetaCache_UpdateSegmentView_Call) Return() *MockMetaCache_UpdateSegmentView_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMetaCache_UpdateSegmentView_Call) RunAndReturn(run func(int64, []*datapb.SyncSegmentInfo, []*pkoracle.BloomFilterSet, map[int64]struct{})) *MockMetaCache_UpdateSegmentView_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSegments provides a mock function with given fields: action, filters
func (_m *MockMetaCache) UpdateSegments(action SegmentAction, filters ...SegmentFilter) {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, action)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockMetaCache_UpdateSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSegments'
type MockMetaCache_UpdateSegments_Call struct {
	*mock.Call
}

// UpdateSegments is a helper method to define mock.On call
//   - action SegmentAction
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) UpdateSegments(action interface{}, filters ...interface{}) *MockMetaCache_UpdateSegments_Call {
	return &MockMetaCache_UpdateSegments_Call{Call: _e.mock.On("UpdateSegments",
		append([]interface{}{action}, filters...)...)}
}

func (_c *MockMetaCache_UpdateSegments_Call) Run(run func(action SegmentAction, filters ...SegmentFilter)) *MockMetaCache_UpdateSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(args[0].(SegmentAction), variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_UpdateSegments_Call) Return() *MockMetaCache_UpdateSegments_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMetaCache_UpdateSegments_Call) RunAndReturn(run func(SegmentAction, ...SegmentFilter)) *MockMetaCache_UpdateSegments_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetaCache creates a new instance of MockMetaCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetaCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetaCache {
	mock := &MockMetaCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
