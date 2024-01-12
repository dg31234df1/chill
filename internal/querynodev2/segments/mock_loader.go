// Code generated by mockery v2.32.4. DO NOT EDIT.

package segments

import (
	context "context"

	commonpb "github.com/milvus-io/milvus-proto/go-api/v2/commonpb"

	datapb "github.com/milvus-io/milvus/internal/proto/datapb"

	mock "github.com/stretchr/testify/mock"

	pkoracle "github.com/milvus-io/milvus/internal/querynodev2/pkoracle"

	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
)

// MockLoader is an autogenerated mock type for the Loader type
type MockLoader struct {
	mock.Mock
}

type MockLoader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLoader) EXPECT() *MockLoader_Expecter {
	return &MockLoader_Expecter{mock: &_m.Mock}
}

// Load provides a mock function with given fields: ctx, collectionID, segmentType, version, segments
func (_m *MockLoader) Load(ctx context.Context, collectionID int64, segmentType commonpb.SegmentState, version int64, segments ...*querypb.SegmentLoadInfo) ([]Segment, error) {
	_va := make([]interface{}, len(segments))
	for _i := range segments {
		_va[_i] = segments[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, collectionID, segmentType, version)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []Segment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, commonpb.SegmentState, int64, ...*querypb.SegmentLoadInfo) ([]Segment, error)); ok {
		return rf(ctx, collectionID, segmentType, version, segments...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, commonpb.SegmentState, int64, ...*querypb.SegmentLoadInfo) []Segment); ok {
		r0 = rf(ctx, collectionID, segmentType, version, segments...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Segment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, commonpb.SegmentState, int64, ...*querypb.SegmentLoadInfo) error); ok {
		r1 = rf(ctx, collectionID, segmentType, version, segments...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLoader_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type MockLoader_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
//   - segmentType commonpb.SegmentState
//   - version int64
//   - segments ...*querypb.SegmentLoadInfo
func (_e *MockLoader_Expecter) Load(ctx interface{}, collectionID interface{}, segmentType interface{}, version interface{}, segments ...interface{}) *MockLoader_Load_Call {
	return &MockLoader_Load_Call{Call: _e.mock.On("Load",
		append([]interface{}{ctx, collectionID, segmentType, version}, segments...)...)}
}

func (_c *MockLoader_Load_Call) Run(run func(ctx context.Context, collectionID int64, segmentType commonpb.SegmentState, version int64, segments ...*querypb.SegmentLoadInfo)) *MockLoader_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.SegmentLoadInfo, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.SegmentLoadInfo)
			}
		}
		run(args[0].(context.Context), args[1].(int64), args[2].(commonpb.SegmentState), args[3].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockLoader_Load_Call) Return(_a0 []Segment, _a1 error) *MockLoader_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLoader_Load_Call) RunAndReturn(run func(context.Context, int64, commonpb.SegmentState, int64, ...*querypb.SegmentLoadInfo) ([]Segment, error)) *MockLoader_Load_Call {
	_c.Call.Return(run)
	return _c
}

// LoadBloomFilterSet provides a mock function with given fields: ctx, collectionID, version, infos
func (_m *MockLoader) LoadBloomFilterSet(ctx context.Context, collectionID int64, version int64, infos ...*querypb.SegmentLoadInfo) ([]*pkoracle.BloomFilterSet, error) {
	_va := make([]interface{}, len(infos))
	for _i := range infos {
		_va[_i] = infos[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, collectionID, version)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*pkoracle.BloomFilterSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, ...*querypb.SegmentLoadInfo) ([]*pkoracle.BloomFilterSet, error)); ok {
		return rf(ctx, collectionID, version, infos...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, ...*querypb.SegmentLoadInfo) []*pkoracle.BloomFilterSet); ok {
		r0 = rf(ctx, collectionID, version, infos...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pkoracle.BloomFilterSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, ...*querypb.SegmentLoadInfo) error); ok {
		r1 = rf(ctx, collectionID, version, infos...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLoader_LoadBloomFilterSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadBloomFilterSet'
type MockLoader_LoadBloomFilterSet_Call struct {
	*mock.Call
}

// LoadBloomFilterSet is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
//   - version int64
//   - infos ...*querypb.SegmentLoadInfo
func (_e *MockLoader_Expecter) LoadBloomFilterSet(ctx interface{}, collectionID interface{}, version interface{}, infos ...interface{}) *MockLoader_LoadBloomFilterSet_Call {
	return &MockLoader_LoadBloomFilterSet_Call{Call: _e.mock.On("LoadBloomFilterSet",
		append([]interface{}{ctx, collectionID, version}, infos...)...)}
}

func (_c *MockLoader_LoadBloomFilterSet_Call) Run(run func(ctx context.Context, collectionID int64, version int64, infos ...*querypb.SegmentLoadInfo)) *MockLoader_LoadBloomFilterSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.SegmentLoadInfo, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.SegmentLoadInfo)
			}
		}
		run(args[0].(context.Context), args[1].(int64), args[2].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockLoader_LoadBloomFilterSet_Call) Return(_a0 []*pkoracle.BloomFilterSet, _a1 error) *MockLoader_LoadBloomFilterSet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLoader_LoadBloomFilterSet_Call) RunAndReturn(run func(context.Context, int64, int64, ...*querypb.SegmentLoadInfo) ([]*pkoracle.BloomFilterSet, error)) *MockLoader_LoadBloomFilterSet_Call {
	_c.Call.Return(run)
	return _c
}

// LoadDelta provides a mock function with given fields: ctx, collectionID, segment
func (_m *MockLoader) LoadDelta(ctx context.Context, collectionID int64, segment *LocalSegment) error {
	ret := _m.Called(ctx, collectionID, segment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *LocalSegment) error); ok {
		r0 = rf(ctx, collectionID, segment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLoader_LoadDelta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadDelta'
type MockLoader_LoadDelta_Call struct {
	*mock.Call
}

// LoadDelta is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
//   - segment *LocalSegment
func (_e *MockLoader_Expecter) LoadDelta(ctx interface{}, collectionID interface{}, segment interface{}) *MockLoader_LoadDelta_Call {
	return &MockLoader_LoadDelta_Call{Call: _e.mock.On("LoadDelta", ctx, collectionID, segment)}
}

func (_c *MockLoader_LoadDelta_Call) Run(run func(ctx context.Context, collectionID int64, segment *LocalSegment)) *MockLoader_LoadDelta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*LocalSegment))
	})
	return _c
}

func (_c *MockLoader_LoadDelta_Call) Return(_a0 error) *MockLoader_LoadDelta_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLoader_LoadDelta_Call) RunAndReturn(run func(context.Context, int64, *LocalSegment) error) *MockLoader_LoadDelta_Call {
	_c.Call.Return(run)
	return _c
}

// LoadDeltaLogs provides a mock function with given fields: ctx, segment, deltaLogs
func (_m *MockLoader) LoadDeltaLogs(ctx context.Context, segment Segment, deltaLogs []*datapb.FieldBinlog) error {
	ret := _m.Called(ctx, segment, deltaLogs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Segment, []*datapb.FieldBinlog) error); ok {
		r0 = rf(ctx, segment, deltaLogs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLoader_LoadDeltaLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadDeltaLogs'
type MockLoader_LoadDeltaLogs_Call struct {
	*mock.Call
}

// LoadDeltaLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - segment Segment
//   - deltaLogs []*datapb.FieldBinlog
func (_e *MockLoader_Expecter) LoadDeltaLogs(ctx interface{}, segment interface{}, deltaLogs interface{}) *MockLoader_LoadDeltaLogs_Call {
	return &MockLoader_LoadDeltaLogs_Call{Call: _e.mock.On("LoadDeltaLogs", ctx, segment, deltaLogs)}
}

func (_c *MockLoader_LoadDeltaLogs_Call) Run(run func(ctx context.Context, segment Segment, deltaLogs []*datapb.FieldBinlog)) *MockLoader_LoadDeltaLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Segment), args[2].([]*datapb.FieldBinlog))
	})
	return _c
}

func (_c *MockLoader_LoadDeltaLogs_Call) Return(_a0 error) *MockLoader_LoadDeltaLogs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLoader_LoadDeltaLogs_Call) RunAndReturn(run func(context.Context, Segment, []*datapb.FieldBinlog) error) *MockLoader_LoadDeltaLogs_Call {
	_c.Call.Return(run)
	return _c
}

// LoadIndex provides a mock function with given fields: ctx, segment, info, version
func (_m *MockLoader) LoadIndex(ctx context.Context, segment *LocalSegment, info *querypb.SegmentLoadInfo, version int64) error {
	ret := _m.Called(ctx, segment, info, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *LocalSegment, *querypb.SegmentLoadInfo, int64) error); ok {
		r0 = rf(ctx, segment, info, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLoader_LoadIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadIndex'
type MockLoader_LoadIndex_Call struct {
	*mock.Call
}

// LoadIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - segment *LocalSegment
//   - info *querypb.SegmentLoadInfo
//   - version int64
func (_e *MockLoader_Expecter) LoadIndex(ctx interface{}, segment interface{}, info interface{}, version interface{}) *MockLoader_LoadIndex_Call {
	return &MockLoader_LoadIndex_Call{Call: _e.mock.On("LoadIndex", ctx, segment, info, version)}
}

func (_c *MockLoader_LoadIndex_Call) Run(run func(ctx context.Context, segment *LocalSegment, info *querypb.SegmentLoadInfo, version int64)) *MockLoader_LoadIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*LocalSegment), args[2].(*querypb.SegmentLoadInfo), args[3].(int64))
	})
	return _c
}

func (_c *MockLoader_LoadIndex_Call) Return(_a0 error) *MockLoader_LoadIndex_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLoader_LoadIndex_Call) RunAndReturn(run func(context.Context, *LocalSegment, *querypb.SegmentLoadInfo, int64) error) *MockLoader_LoadIndex_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLoader creates a new instance of MockLoader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLoader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLoader {
	mock := &MockLoader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
