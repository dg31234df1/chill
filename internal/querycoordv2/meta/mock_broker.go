// Code generated by mockery v2.46.0. DO NOT EDIT.

package meta

import (
	context "context"

	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	indexpb "github.com/milvus-io/milvus/internal/proto/indexpb"

	milvuspb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"

	mock "github.com/stretchr/testify/mock"

	querypb "github.com/milvus-io/milvus/internal/proto/querypb"

	rootcoordpb "github.com/milvus-io/milvus/internal/proto/rootcoordpb"
)

// MockBroker is an autogenerated mock type for the Broker type
type MockBroker struct {
	mock.Mock
}

type MockBroker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBroker) EXPECT() *MockBroker_Expecter {
	return &MockBroker_Expecter{mock: &_m.Mock}
}

// DescribeCollection provides a mock function with given fields: ctx, collectionID
func (_m *MockBroker) DescribeCollection(ctx context.Context, collectionID int64) (*milvuspb.DescribeCollectionResponse, error) {
	ret := _m.Called(ctx, collectionID)

	if len(ret) == 0 {
		panic("no return value specified for DescribeCollection")
	}

	var r0 *milvuspb.DescribeCollectionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*milvuspb.DescribeCollectionResponse, error)); ok {
		return rf(ctx, collectionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *milvuspb.DescribeCollectionResponse); ok {
		r0 = rf(ctx, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.DescribeCollectionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroker_DescribeCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeCollection'
type MockBroker_DescribeCollection_Call struct {
	*mock.Call
}

// DescribeCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *MockBroker_Expecter) DescribeCollection(ctx interface{}, collectionID interface{}) *MockBroker_DescribeCollection_Call {
	return &MockBroker_DescribeCollection_Call{Call: _e.mock.On("DescribeCollection", ctx, collectionID)}
}

func (_c *MockBroker_DescribeCollection_Call) Run(run func(ctx context.Context, collectionID int64)) *MockBroker_DescribeCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockBroker_DescribeCollection_Call) Return(_a0 *milvuspb.DescribeCollectionResponse, _a1 error) *MockBroker_DescribeCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_DescribeCollection_Call) RunAndReturn(run func(context.Context, int64) (*milvuspb.DescribeCollectionResponse, error)) *MockBroker_DescribeCollection_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeDatabase provides a mock function with given fields: ctx, dbName
func (_m *MockBroker) DescribeDatabase(ctx context.Context, dbName string) (*rootcoordpb.DescribeDatabaseResponse, error) {
	ret := _m.Called(ctx, dbName)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDatabase")
	}

	var r0 *rootcoordpb.DescribeDatabaseResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*rootcoordpb.DescribeDatabaseResponse, error)); ok {
		return rf(ctx, dbName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *rootcoordpb.DescribeDatabaseResponse); ok {
		r0 = rf(ctx, dbName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rootcoordpb.DescribeDatabaseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, dbName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroker_DescribeDatabase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeDatabase'
type MockBroker_DescribeDatabase_Call struct {
	*mock.Call
}

// DescribeDatabase is a helper method to define mock.On call
//   - ctx context.Context
//   - dbName string
func (_e *MockBroker_Expecter) DescribeDatabase(ctx interface{}, dbName interface{}) *MockBroker_DescribeDatabase_Call {
	return &MockBroker_DescribeDatabase_Call{Call: _e.mock.On("DescribeDatabase", ctx, dbName)}
}

func (_c *MockBroker_DescribeDatabase_Call) Run(run func(ctx context.Context, dbName string)) *MockBroker_DescribeDatabase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBroker_DescribeDatabase_Call) Return(_a0 *rootcoordpb.DescribeDatabaseResponse, _a1 error) *MockBroker_DescribeDatabase_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_DescribeDatabase_Call) RunAndReturn(run func(context.Context, string) (*rootcoordpb.DescribeDatabaseResponse, error)) *MockBroker_DescribeDatabase_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionLoadInfo provides a mock function with given fields: ctx, collectionID
func (_m *MockBroker) GetCollectionLoadInfo(ctx context.Context, collectionID int64) ([]string, int64, error) {
	ret := _m.Called(ctx, collectionID)

	if len(ret) == 0 {
		panic("no return value specified for GetCollectionLoadInfo")
	}

	var r0 []string
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]string, int64, error)); ok {
		return rf(ctx, collectionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []string); ok {
		r0 = rf(ctx, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) int64); ok {
		r1 = rf(ctx, collectionID)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int64) error); ok {
		r2 = rf(ctx, collectionID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBroker_GetCollectionLoadInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionLoadInfo'
type MockBroker_GetCollectionLoadInfo_Call struct {
	*mock.Call
}

// GetCollectionLoadInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *MockBroker_Expecter) GetCollectionLoadInfo(ctx interface{}, collectionID interface{}) *MockBroker_GetCollectionLoadInfo_Call {
	return &MockBroker_GetCollectionLoadInfo_Call{Call: _e.mock.On("GetCollectionLoadInfo", ctx, collectionID)}
}

func (_c *MockBroker_GetCollectionLoadInfo_Call) Run(run func(ctx context.Context, collectionID int64)) *MockBroker_GetCollectionLoadInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockBroker_GetCollectionLoadInfo_Call) Return(_a0 []string, _a1 int64, _a2 error) *MockBroker_GetCollectionLoadInfo_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockBroker_GetCollectionLoadInfo_Call) RunAndReturn(run func(context.Context, int64) ([]string, int64, error)) *MockBroker_GetCollectionLoadInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetIndexInfo provides a mock function with given fields: ctx, collectionID, segmentIDs
func (_m *MockBroker) GetIndexInfo(ctx context.Context, collectionID int64, segmentIDs ...int64) (map[int64][]*querypb.FieldIndexInfo, error) {
	_va := make([]interface{}, len(segmentIDs))
	for _i := range segmentIDs {
		_va[_i] = segmentIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, collectionID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetIndexInfo")
	}

	var r0 map[int64][]*querypb.FieldIndexInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...int64) (map[int64][]*querypb.FieldIndexInfo, error)); ok {
		return rf(ctx, collectionID, segmentIDs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...int64) map[int64][]*querypb.FieldIndexInfo); ok {
		r0 = rf(ctx, collectionID, segmentIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64][]*querypb.FieldIndexInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...int64) error); ok {
		r1 = rf(ctx, collectionID, segmentIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroker_GetIndexInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndexInfo'
type MockBroker_GetIndexInfo_Call struct {
	*mock.Call
}

// GetIndexInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
//   - segmentIDs ...int64
func (_e *MockBroker_Expecter) GetIndexInfo(ctx interface{}, collectionID interface{}, segmentIDs ...interface{}) *MockBroker_GetIndexInfo_Call {
	return &MockBroker_GetIndexInfo_Call{Call: _e.mock.On("GetIndexInfo",
		append([]interface{}{ctx, collectionID}, segmentIDs...)...)}
}

func (_c *MockBroker_GetIndexInfo_Call) Run(run func(ctx context.Context, collectionID int64, segmentIDs ...int64)) *MockBroker_GetIndexInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockBroker_GetIndexInfo_Call) Return(_a0 map[int64][]*querypb.FieldIndexInfo, _a1 error) *MockBroker_GetIndexInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_GetIndexInfo_Call) RunAndReturn(run func(context.Context, int64, ...int64) (map[int64][]*querypb.FieldIndexInfo, error)) *MockBroker_GetIndexInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartitions provides a mock function with given fields: ctx, collectionID
func (_m *MockBroker) GetPartitions(ctx context.Context, collectionID int64) ([]int64, error) {
	ret := _m.Called(ctx, collectionID)

	if len(ret) == 0 {
		panic("no return value specified for GetPartitions")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]int64, error)); ok {
		return rf(ctx, collectionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []int64); ok {
		r0 = rf(ctx, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroker_GetPartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitions'
type MockBroker_GetPartitions_Call struct {
	*mock.Call
}

// GetPartitions is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *MockBroker_Expecter) GetPartitions(ctx interface{}, collectionID interface{}) *MockBroker_GetPartitions_Call {
	return &MockBroker_GetPartitions_Call{Call: _e.mock.On("GetPartitions", ctx, collectionID)}
}

func (_c *MockBroker_GetPartitions_Call) Run(run func(ctx context.Context, collectionID int64)) *MockBroker_GetPartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockBroker_GetPartitions_Call) Return(_a0 []int64, _a1 error) *MockBroker_GetPartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_GetPartitions_Call) RunAndReturn(run func(context.Context, int64) ([]int64, error)) *MockBroker_GetPartitions_Call {
	_c.Call.Return(run)
	return _c
}

// GetRecoveryInfo provides a mock function with given fields: ctx, collectionID, partitionID
func (_m *MockBroker) GetRecoveryInfo(ctx context.Context, collectionID int64, partitionID int64) ([]*datapb.VchannelInfo, []*datapb.SegmentBinlogs, error) {
	ret := _m.Called(ctx, collectionID, partitionID)

	if len(ret) == 0 {
		panic("no return value specified for GetRecoveryInfo")
	}

	var r0 []*datapb.VchannelInfo
	var r1 []*datapb.SegmentBinlogs
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) ([]*datapb.VchannelInfo, []*datapb.SegmentBinlogs, error)); ok {
		return rf(ctx, collectionID, partitionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []*datapb.VchannelInfo); ok {
		r0 = rf(ctx, collectionID, partitionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datapb.VchannelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) []*datapb.SegmentBinlogs); ok {
		r1 = rf(ctx, collectionID, partitionID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*datapb.SegmentBinlogs)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, int64, int64) error); ok {
		r2 = rf(ctx, collectionID, partitionID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBroker_GetRecoveryInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRecoveryInfo'
type MockBroker_GetRecoveryInfo_Call struct {
	*mock.Call
}

// GetRecoveryInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
//   - partitionID int64
func (_e *MockBroker_Expecter) GetRecoveryInfo(ctx interface{}, collectionID interface{}, partitionID interface{}) *MockBroker_GetRecoveryInfo_Call {
	return &MockBroker_GetRecoveryInfo_Call{Call: _e.mock.On("GetRecoveryInfo", ctx, collectionID, partitionID)}
}

func (_c *MockBroker_GetRecoveryInfo_Call) Run(run func(ctx context.Context, collectionID int64, partitionID int64)) *MockBroker_GetRecoveryInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *MockBroker_GetRecoveryInfo_Call) Return(_a0 []*datapb.VchannelInfo, _a1 []*datapb.SegmentBinlogs, _a2 error) *MockBroker_GetRecoveryInfo_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockBroker_GetRecoveryInfo_Call) RunAndReturn(run func(context.Context, int64, int64) ([]*datapb.VchannelInfo, []*datapb.SegmentBinlogs, error)) *MockBroker_GetRecoveryInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetRecoveryInfoV2 provides a mock function with given fields: ctx, collectionID, partitionIDs
func (_m *MockBroker) GetRecoveryInfoV2(ctx context.Context, collectionID int64, partitionIDs ...int64) ([]*datapb.VchannelInfo, []*datapb.SegmentInfo, error) {
	_va := make([]interface{}, len(partitionIDs))
	for _i := range partitionIDs {
		_va[_i] = partitionIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, collectionID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRecoveryInfoV2")
	}

	var r0 []*datapb.VchannelInfo
	var r1 []*datapb.SegmentInfo
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...int64) ([]*datapb.VchannelInfo, []*datapb.SegmentInfo, error)); ok {
		return rf(ctx, collectionID, partitionIDs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...int64) []*datapb.VchannelInfo); ok {
		r0 = rf(ctx, collectionID, partitionIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datapb.VchannelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...int64) []*datapb.SegmentInfo); ok {
		r1 = rf(ctx, collectionID, partitionIDs...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*datapb.SegmentInfo)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, int64, ...int64) error); ok {
		r2 = rf(ctx, collectionID, partitionIDs...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBroker_GetRecoveryInfoV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRecoveryInfoV2'
type MockBroker_GetRecoveryInfoV2_Call struct {
	*mock.Call
}

// GetRecoveryInfoV2 is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
//   - partitionIDs ...int64
func (_e *MockBroker_Expecter) GetRecoveryInfoV2(ctx interface{}, collectionID interface{}, partitionIDs ...interface{}) *MockBroker_GetRecoveryInfoV2_Call {
	return &MockBroker_GetRecoveryInfoV2_Call{Call: _e.mock.On("GetRecoveryInfoV2",
		append([]interface{}{ctx, collectionID}, partitionIDs...)...)}
}

func (_c *MockBroker_GetRecoveryInfoV2_Call) Run(run func(ctx context.Context, collectionID int64, partitionIDs ...int64)) *MockBroker_GetRecoveryInfoV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockBroker_GetRecoveryInfoV2_Call) Return(_a0 []*datapb.VchannelInfo, _a1 []*datapb.SegmentInfo, _a2 error) *MockBroker_GetRecoveryInfoV2_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockBroker_GetRecoveryInfoV2_Call) RunAndReturn(run func(context.Context, int64, ...int64) ([]*datapb.VchannelInfo, []*datapb.SegmentInfo, error)) *MockBroker_GetRecoveryInfoV2_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegmentInfo provides a mock function with given fields: ctx, segmentID
func (_m *MockBroker) GetSegmentInfo(ctx context.Context, segmentID ...int64) ([]*datapb.SegmentInfo, error) {
	_va := make([]interface{}, len(segmentID))
	for _i := range segmentID {
		_va[_i] = segmentID[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSegmentInfo")
	}

	var r0 []*datapb.SegmentInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...int64) ([]*datapb.SegmentInfo, error)); ok {
		return rf(ctx, segmentID...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...int64) []*datapb.SegmentInfo); ok {
		r0 = rf(ctx, segmentID...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datapb.SegmentInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...int64) error); ok {
		r1 = rf(ctx, segmentID...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroker_GetSegmentInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentInfo'
type MockBroker_GetSegmentInfo_Call struct {
	*mock.Call
}

// GetSegmentInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - segmentID ...int64
func (_e *MockBroker_Expecter) GetSegmentInfo(ctx interface{}, segmentID ...interface{}) *MockBroker_GetSegmentInfo_Call {
	return &MockBroker_GetSegmentInfo_Call{Call: _e.mock.On("GetSegmentInfo",
		append([]interface{}{ctx}, segmentID...)...)}
}

func (_c *MockBroker_GetSegmentInfo_Call) Run(run func(ctx context.Context, segmentID ...int64)) *MockBroker_GetSegmentInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockBroker_GetSegmentInfo_Call) Return(_a0 []*datapb.SegmentInfo, _a1 error) *MockBroker_GetSegmentInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_GetSegmentInfo_Call) RunAndReturn(run func(context.Context, ...int64) ([]*datapb.SegmentInfo, error)) *MockBroker_GetSegmentInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ListIndexes provides a mock function with given fields: ctx, collectionID
func (_m *MockBroker) ListIndexes(ctx context.Context, collectionID int64) ([]*indexpb.IndexInfo, error) {
	ret := _m.Called(ctx, collectionID)

	if len(ret) == 0 {
		panic("no return value specified for ListIndexes")
	}

	var r0 []*indexpb.IndexInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*indexpb.IndexInfo, error)); ok {
		return rf(ctx, collectionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*indexpb.IndexInfo); ok {
		r0 = rf(ctx, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*indexpb.IndexInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroker_ListIndexes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListIndexes'
type MockBroker_ListIndexes_Call struct {
	*mock.Call
}

// ListIndexes is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *MockBroker_Expecter) ListIndexes(ctx interface{}, collectionID interface{}) *MockBroker_ListIndexes_Call {
	return &MockBroker_ListIndexes_Call{Call: _e.mock.On("ListIndexes", ctx, collectionID)}
}

func (_c *MockBroker_ListIndexes_Call) Run(run func(ctx context.Context, collectionID int64)) *MockBroker_ListIndexes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockBroker_ListIndexes_Call) Return(_a0 []*indexpb.IndexInfo, _a1 error) *MockBroker_ListIndexes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_ListIndexes_Call) RunAndReturn(run func(context.Context, int64) ([]*indexpb.IndexInfo, error)) *MockBroker_ListIndexes_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBroker creates a new instance of MockBroker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBroker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBroker {
	mock := &MockBroker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
