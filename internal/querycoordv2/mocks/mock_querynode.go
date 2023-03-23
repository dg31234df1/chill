// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	commonpb "github.com/milvus-io/milvus-proto/go-api/commonpb"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"

	milvuspb "github.com/milvus-io/milvus-proto/go-api/milvuspb"

	mock "github.com/stretchr/testify/mock"

	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
)

// MockQueryNodeServer is an autogenerated mock type for the QueryNodeServer type
type MockQueryNodeServer struct {
	mock.Mock
}

type MockQueryNodeServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQueryNodeServer) EXPECT() *MockQueryNodeServer_Expecter {
	return &MockQueryNodeServer_Expecter{mock: &_m.Mock}
}

// GetComponentStates provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) GetComponentStates(_a0 context.Context, _a1 *milvuspb.GetComponentStatesRequest) (*milvuspb.ComponentStates, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *milvuspb.ComponentStates
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetComponentStatesRequest) *milvuspb.ComponentStates); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.ComponentStates)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *milvuspb.GetComponentStatesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_GetComponentStates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComponentStates'
type MockQueryNodeServer_GetComponentStates_Call struct {
	*mock.Call
}

// GetComponentStates is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *milvuspb.GetComponentStatesRequest
func (_e *MockQueryNodeServer_Expecter) GetComponentStates(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_GetComponentStates_Call {
	return &MockQueryNodeServer_GetComponentStates_Call{Call: _e.mock.On("GetComponentStates", _a0, _a1)}
}

func (_c *MockQueryNodeServer_GetComponentStates_Call) Run(run func(_a0 context.Context, _a1 *milvuspb.GetComponentStatesRequest)) *MockQueryNodeServer_GetComponentStates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*milvuspb.GetComponentStatesRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_GetComponentStates_Call) Return(_a0 *milvuspb.ComponentStates, _a1 error) *MockQueryNodeServer_GetComponentStates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetDataDistribution provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) GetDataDistribution(_a0 context.Context, _a1 *querypb.GetDataDistributionRequest) (*querypb.GetDataDistributionResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *querypb.GetDataDistributionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.GetDataDistributionRequest) *querypb.GetDataDistributionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querypb.GetDataDistributionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.GetDataDistributionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_GetDataDistribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataDistribution'
type MockQueryNodeServer_GetDataDistribution_Call struct {
	*mock.Call
}

// GetDataDistribution is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.GetDataDistributionRequest
func (_e *MockQueryNodeServer_Expecter) GetDataDistribution(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_GetDataDistribution_Call {
	return &MockQueryNodeServer_GetDataDistribution_Call{Call: _e.mock.On("GetDataDistribution", _a0, _a1)}
}

func (_c *MockQueryNodeServer_GetDataDistribution_Call) Run(run func(_a0 context.Context, _a1 *querypb.GetDataDistributionRequest)) *MockQueryNodeServer_GetDataDistribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.GetDataDistributionRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_GetDataDistribution_Call) Return(_a0 *querypb.GetDataDistributionResponse, _a1 error) *MockQueryNodeServer_GetDataDistribution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetMetrics provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) GetMetrics(_a0 context.Context, _a1 *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *milvuspb.GetMetricsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetMetricsRequest) *milvuspb.GetMetricsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.GetMetricsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *milvuspb.GetMetricsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_GetMetrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetrics'
type MockQueryNodeServer_GetMetrics_Call struct {
	*mock.Call
}

// GetMetrics is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *milvuspb.GetMetricsRequest
func (_e *MockQueryNodeServer_Expecter) GetMetrics(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_GetMetrics_Call {
	return &MockQueryNodeServer_GetMetrics_Call{Call: _e.mock.On("GetMetrics", _a0, _a1)}
}

func (_c *MockQueryNodeServer_GetMetrics_Call) Run(run func(_a0 context.Context, _a1 *milvuspb.GetMetricsRequest)) *MockQueryNodeServer_GetMetrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*milvuspb.GetMetricsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_GetMetrics_Call) Return(_a0 *milvuspb.GetMetricsResponse, _a1 error) *MockQueryNodeServer_GetMetrics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSegmentInfo provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) GetSegmentInfo(_a0 context.Context, _a1 *querypb.GetSegmentInfoRequest) (*querypb.GetSegmentInfoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *querypb.GetSegmentInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.GetSegmentInfoRequest) *querypb.GetSegmentInfoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querypb.GetSegmentInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.GetSegmentInfoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_GetSegmentInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentInfo'
type MockQueryNodeServer_GetSegmentInfo_Call struct {
	*mock.Call
}

// GetSegmentInfo is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.GetSegmentInfoRequest
func (_e *MockQueryNodeServer_Expecter) GetSegmentInfo(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_GetSegmentInfo_Call {
	return &MockQueryNodeServer_GetSegmentInfo_Call{Call: _e.mock.On("GetSegmentInfo", _a0, _a1)}
}

func (_c *MockQueryNodeServer_GetSegmentInfo_Call) Run(run func(_a0 context.Context, _a1 *querypb.GetSegmentInfoRequest)) *MockQueryNodeServer_GetSegmentInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.GetSegmentInfoRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_GetSegmentInfo_Call) Return(_a0 *querypb.GetSegmentInfoResponse, _a1 error) *MockQueryNodeServer_GetSegmentInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetStatistics provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) GetStatistics(_a0 context.Context, _a1 *querypb.GetStatisticsRequest) (*internalpb.GetStatisticsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalpb.GetStatisticsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.GetStatisticsRequest) *internalpb.GetStatisticsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.GetStatisticsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.GetStatisticsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_GetStatistics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatistics'
type MockQueryNodeServer_GetStatistics_Call struct {
	*mock.Call
}

// GetStatistics is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.GetStatisticsRequest
func (_e *MockQueryNodeServer_Expecter) GetStatistics(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_GetStatistics_Call {
	return &MockQueryNodeServer_GetStatistics_Call{Call: _e.mock.On("GetStatistics", _a0, _a1)}
}

func (_c *MockQueryNodeServer_GetStatistics_Call) Run(run func(_a0 context.Context, _a1 *querypb.GetStatisticsRequest)) *MockQueryNodeServer_GetStatistics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.GetStatisticsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_GetStatistics_Call) Return(_a0 *internalpb.GetStatisticsResponse, _a1 error) *MockQueryNodeServer_GetStatistics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetStatisticsChannel provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) GetStatisticsChannel(_a0 context.Context, _a1 *internalpb.GetStatisticsChannelRequest) (*milvuspb.StringResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *milvuspb.StringResponse
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.GetStatisticsChannelRequest) *milvuspb.StringResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.StringResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalpb.GetStatisticsChannelRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_GetStatisticsChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatisticsChannel'
type MockQueryNodeServer_GetStatisticsChannel_Call struct {
	*mock.Call
}

// GetStatisticsChannel is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *internalpb.GetStatisticsChannelRequest
func (_e *MockQueryNodeServer_Expecter) GetStatisticsChannel(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_GetStatisticsChannel_Call {
	return &MockQueryNodeServer_GetStatisticsChannel_Call{Call: _e.mock.On("GetStatisticsChannel", _a0, _a1)}
}

func (_c *MockQueryNodeServer_GetStatisticsChannel_Call) Run(run func(_a0 context.Context, _a1 *internalpb.GetStatisticsChannelRequest)) *MockQueryNodeServer_GetStatisticsChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internalpb.GetStatisticsChannelRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_GetStatisticsChannel_Call) Return(_a0 *milvuspb.StringResponse, _a1 error) *MockQueryNodeServer_GetStatisticsChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTimeTickChannel provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) GetTimeTickChannel(_a0 context.Context, _a1 *internalpb.GetTimeTickChannelRequest) (*milvuspb.StringResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *milvuspb.StringResponse
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.GetTimeTickChannelRequest) *milvuspb.StringResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.StringResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalpb.GetTimeTickChannelRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_GetTimeTickChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTimeTickChannel'
type MockQueryNodeServer_GetTimeTickChannel_Call struct {
	*mock.Call
}

// GetTimeTickChannel is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *internalpb.GetTimeTickChannelRequest
func (_e *MockQueryNodeServer_Expecter) GetTimeTickChannel(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_GetTimeTickChannel_Call {
	return &MockQueryNodeServer_GetTimeTickChannel_Call{Call: _e.mock.On("GetTimeTickChannel", _a0, _a1)}
}

func (_c *MockQueryNodeServer_GetTimeTickChannel_Call) Run(run func(_a0 context.Context, _a1 *internalpb.GetTimeTickChannelRequest)) *MockQueryNodeServer_GetTimeTickChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internalpb.GetTimeTickChannelRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_GetTimeTickChannel_Call) Return(_a0 *milvuspb.StringResponse, _a1 error) *MockQueryNodeServer_GetTimeTickChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadPartitions provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) LoadPartitions(_a0 context.Context, _a1 *querypb.LoadPartitionsRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.LoadPartitionsRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.LoadPartitionsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_LoadPartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadPartitions'
type MockQueryNodeServer_LoadPartitions_Call struct {
	*mock.Call
}

// LoadPartitions is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.LoadPartitionsRequest
func (_e *MockQueryNodeServer_Expecter) LoadPartitions(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_LoadPartitions_Call {
	return &MockQueryNodeServer_LoadPartitions_Call{Call: _e.mock.On("LoadPartitions", _a0, _a1)}
}

func (_c *MockQueryNodeServer_LoadPartitions_Call) Run(run func(_a0 context.Context, _a1 *querypb.LoadPartitionsRequest)) *MockQueryNodeServer_LoadPartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.LoadPartitionsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_LoadPartitions_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_LoadPartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadSegments provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) LoadSegments(_a0 context.Context, _a1 *querypb.LoadSegmentsRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.LoadSegmentsRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.LoadSegmentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_LoadSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadSegments'
type MockQueryNodeServer_LoadSegments_Call struct {
	*mock.Call
}

// LoadSegments is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.LoadSegmentsRequest
func (_e *MockQueryNodeServer_Expecter) LoadSegments(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_LoadSegments_Call {
	return &MockQueryNodeServer_LoadSegments_Call{Call: _e.mock.On("LoadSegments", _a0, _a1)}
}

func (_c *MockQueryNodeServer_LoadSegments_Call) Run(run func(_a0 context.Context, _a1 *querypb.LoadSegmentsRequest)) *MockQueryNodeServer_LoadSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.LoadSegmentsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_LoadSegments_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_LoadSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Query provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) Query(_a0 context.Context, _a1 *querypb.QueryRequest) (*internalpb.RetrieveResults, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalpb.RetrieveResults
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.QueryRequest) *internalpb.RetrieveResults); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.RetrieveResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.QueryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockQueryNodeServer_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.QueryRequest
func (_e *MockQueryNodeServer_Expecter) Query(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_Query_Call {
	return &MockQueryNodeServer_Query_Call{Call: _e.mock.On("Query", _a0, _a1)}
}

func (_c *MockQueryNodeServer_Query_Call) Run(run func(_a0 context.Context, _a1 *querypb.QueryRequest)) *MockQueryNodeServer_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.QueryRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_Query_Call) Return(_a0 *internalpb.RetrieveResults, _a1 error) *MockQueryNodeServer_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ReleaseCollection provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) ReleaseCollection(_a0 context.Context, _a1 *querypb.ReleaseCollectionRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.ReleaseCollectionRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.ReleaseCollectionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_ReleaseCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseCollection'
type MockQueryNodeServer_ReleaseCollection_Call struct {
	*mock.Call
}

// ReleaseCollection is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.ReleaseCollectionRequest
func (_e *MockQueryNodeServer_Expecter) ReleaseCollection(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_ReleaseCollection_Call {
	return &MockQueryNodeServer_ReleaseCollection_Call{Call: _e.mock.On("ReleaseCollection", _a0, _a1)}
}

func (_c *MockQueryNodeServer_ReleaseCollection_Call) Run(run func(_a0 context.Context, _a1 *querypb.ReleaseCollectionRequest)) *MockQueryNodeServer_ReleaseCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.ReleaseCollectionRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_ReleaseCollection_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_ReleaseCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ReleasePartitions provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) ReleasePartitions(_a0 context.Context, _a1 *querypb.ReleasePartitionsRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.ReleasePartitionsRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.ReleasePartitionsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_ReleasePartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleasePartitions'
type MockQueryNodeServer_ReleasePartitions_Call struct {
	*mock.Call
}

// ReleasePartitions is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.ReleasePartitionsRequest
func (_e *MockQueryNodeServer_Expecter) ReleasePartitions(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_ReleasePartitions_Call {
	return &MockQueryNodeServer_ReleasePartitions_Call{Call: _e.mock.On("ReleasePartitions", _a0, _a1)}
}

func (_c *MockQueryNodeServer_ReleasePartitions_Call) Run(run func(_a0 context.Context, _a1 *querypb.ReleasePartitionsRequest)) *MockQueryNodeServer_ReleasePartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.ReleasePartitionsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_ReleasePartitions_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_ReleasePartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ReleaseSegments provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) ReleaseSegments(_a0 context.Context, _a1 *querypb.ReleaseSegmentsRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.ReleaseSegmentsRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.ReleaseSegmentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_ReleaseSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseSegments'
type MockQueryNodeServer_ReleaseSegments_Call struct {
	*mock.Call
}

// ReleaseSegments is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.ReleaseSegmentsRequest
func (_e *MockQueryNodeServer_Expecter) ReleaseSegments(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_ReleaseSegments_Call {
	return &MockQueryNodeServer_ReleaseSegments_Call{Call: _e.mock.On("ReleaseSegments", _a0, _a1)}
}

func (_c *MockQueryNodeServer_ReleaseSegments_Call) Run(run func(_a0 context.Context, _a1 *querypb.ReleaseSegmentsRequest)) *MockQueryNodeServer_ReleaseSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.ReleaseSegmentsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_ReleaseSegments_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_ReleaseSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) Search(_a0 context.Context, _a1 *querypb.SearchRequest) (*internalpb.SearchResults, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalpb.SearchResults
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.SearchRequest) *internalpb.SearchResults); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.SearchResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.SearchRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockQueryNodeServer_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.SearchRequest
func (_e *MockQueryNodeServer_Expecter) Search(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_Search_Call {
	return &MockQueryNodeServer_Search_Call{Call: _e.mock.On("Search", _a0, _a1)}
}

func (_c *MockQueryNodeServer_Search_Call) Run(run func(_a0 context.Context, _a1 *querypb.SearchRequest)) *MockQueryNodeServer_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.SearchRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_Search_Call) Return(_a0 *internalpb.SearchResults, _a1 error) *MockQueryNodeServer_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ShowConfigurations provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) ShowConfigurations(_a0 context.Context, _a1 *internalpb.ShowConfigurationsRequest) (*internalpb.ShowConfigurationsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *internalpb.ShowConfigurationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.ShowConfigurationsRequest) *internalpb.ShowConfigurationsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.ShowConfigurationsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalpb.ShowConfigurationsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_ShowConfigurations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShowConfigurations'
type MockQueryNodeServer_ShowConfigurations_Call struct {
	*mock.Call
}

// ShowConfigurations is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *internalpb.ShowConfigurationsRequest
func (_e *MockQueryNodeServer_Expecter) ShowConfigurations(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_ShowConfigurations_Call {
	return &MockQueryNodeServer_ShowConfigurations_Call{Call: _e.mock.On("ShowConfigurations", _a0, _a1)}
}

func (_c *MockQueryNodeServer_ShowConfigurations_Call) Run(run func(_a0 context.Context, _a1 *internalpb.ShowConfigurationsRequest)) *MockQueryNodeServer_ShowConfigurations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internalpb.ShowConfigurationsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_ShowConfigurations_Call) Return(_a0 *internalpb.ShowConfigurationsResponse, _a1 error) *MockQueryNodeServer_ShowConfigurations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SyncDistribution provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) SyncDistribution(_a0 context.Context, _a1 *querypb.SyncDistributionRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.SyncDistributionRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.SyncDistributionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_SyncDistribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncDistribution'
type MockQueryNodeServer_SyncDistribution_Call struct {
	*mock.Call
}

// SyncDistribution is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.SyncDistributionRequest
func (_e *MockQueryNodeServer_Expecter) SyncDistribution(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_SyncDistribution_Call {
	return &MockQueryNodeServer_SyncDistribution_Call{Call: _e.mock.On("SyncDistribution", _a0, _a1)}
}

func (_c *MockQueryNodeServer_SyncDistribution_Call) Run(run func(_a0 context.Context, _a1 *querypb.SyncDistributionRequest)) *MockQueryNodeServer_SyncDistribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.SyncDistributionRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_SyncDistribution_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_SyncDistribution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SyncReplicaSegments provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) SyncReplicaSegments(_a0 context.Context, _a1 *querypb.SyncReplicaSegmentsRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.SyncReplicaSegmentsRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.SyncReplicaSegmentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_SyncReplicaSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncReplicaSegments'
type MockQueryNodeServer_SyncReplicaSegments_Call struct {
	*mock.Call
}

// SyncReplicaSegments is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.SyncReplicaSegmentsRequest
func (_e *MockQueryNodeServer_Expecter) SyncReplicaSegments(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_SyncReplicaSegments_Call {
	return &MockQueryNodeServer_SyncReplicaSegments_Call{Call: _e.mock.On("SyncReplicaSegments", _a0, _a1)}
}

func (_c *MockQueryNodeServer_SyncReplicaSegments_Call) Run(run func(_a0 context.Context, _a1 *querypb.SyncReplicaSegmentsRequest)) *MockQueryNodeServer_SyncReplicaSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.SyncReplicaSegmentsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_SyncReplicaSegments_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_SyncReplicaSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UnsubDmChannel provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) UnsubDmChannel(_a0 context.Context, _a1 *querypb.UnsubDmChannelRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.UnsubDmChannelRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.UnsubDmChannelRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_UnsubDmChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnsubDmChannel'
type MockQueryNodeServer_UnsubDmChannel_Call struct {
	*mock.Call
}

// UnsubDmChannel is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.UnsubDmChannelRequest
func (_e *MockQueryNodeServer_Expecter) UnsubDmChannel(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_UnsubDmChannel_Call {
	return &MockQueryNodeServer_UnsubDmChannel_Call{Call: _e.mock.On("UnsubDmChannel", _a0, _a1)}
}

func (_c *MockQueryNodeServer_UnsubDmChannel_Call) Run(run func(_a0 context.Context, _a1 *querypb.UnsubDmChannelRequest)) *MockQueryNodeServer_UnsubDmChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.UnsubDmChannelRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_UnsubDmChannel_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_UnsubDmChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// WatchDmChannels provides a mock function with given fields: _a0, _a1
func (_m *MockQueryNodeServer) WatchDmChannels(_a0 context.Context, _a1 *querypb.WatchDmChannelsRequest) (*commonpb.Status, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.WatchDmChannelsRequest) *commonpb.Status); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.WatchDmChannelsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryNodeServer_WatchDmChannels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchDmChannels'
type MockQueryNodeServer_WatchDmChannels_Call struct {
	*mock.Call
}

// WatchDmChannels is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *querypb.WatchDmChannelsRequest
func (_e *MockQueryNodeServer_Expecter) WatchDmChannels(_a0 interface{}, _a1 interface{}) *MockQueryNodeServer_WatchDmChannels_Call {
	return &MockQueryNodeServer_WatchDmChannels_Call{Call: _e.mock.On("WatchDmChannels", _a0, _a1)}
}

func (_c *MockQueryNodeServer_WatchDmChannels_Call) Run(run func(_a0 context.Context, _a1 *querypb.WatchDmChannelsRequest)) *MockQueryNodeServer_WatchDmChannels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.WatchDmChannelsRequest))
	})
	return _c
}

func (_c *MockQueryNodeServer_WatchDmChannels_Call) Return(_a0 *commonpb.Status, _a1 error) *MockQueryNodeServer_WatchDmChannels_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockQueryNodeServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockQueryNodeServer creates a new instance of MockQueryNodeServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQueryNodeServer(t mockConstructorTestingTNewMockQueryNodeServer) *MockQueryNodeServer {
	mock := &MockQueryNodeServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
