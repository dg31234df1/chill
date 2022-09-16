// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	commonpb "github.com/milvus-io/milvus/api/commonpb"

	indexpb "github.com/milvus-io/milvus/internal/proto/indexpb"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"

	milvuspb "github.com/milvus-io/milvus/api/milvuspb"

	mock "github.com/stretchr/testify/mock"
)

// MockIndexCoord is an autogenerated mock type for the IndexCoord type
type MockIndexCoord struct {
	mock.Mock
}

type MockIndexCoord_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIndexCoord) EXPECT() *MockIndexCoord_Expecter {
	return &MockIndexCoord_Expecter{mock: &_m.Mock}
}

// CreateIndex provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) CreateIndex(ctx context.Context, req *indexpb.CreateIndexRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.CreateIndexRequest) *commonpb.Status); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.CreateIndexRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_CreateIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateIndex'
type MockIndexCoord_CreateIndex_Call struct {
	*mock.Call
}

// CreateIndex is a helper method to define mock.On call
//  - ctx context.Context
//  - req *indexpb.CreateIndexRequest
func (_e *MockIndexCoord_Expecter) CreateIndex(ctx interface{}, req interface{}) *MockIndexCoord_CreateIndex_Call {
	return &MockIndexCoord_CreateIndex_Call{Call: _e.mock.On("CreateIndex", ctx, req)}
}

func (_c *MockIndexCoord_CreateIndex_Call) Run(run func(ctx context.Context, req *indexpb.CreateIndexRequest)) *MockIndexCoord_CreateIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.CreateIndexRequest))
	})
	return _c
}

func (_c *MockIndexCoord_CreateIndex_Call) Return(_a0 *commonpb.Status, _a1 error) *MockIndexCoord_CreateIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DescribeIndex provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) DescribeIndex(ctx context.Context, req *indexpb.DescribeIndexRequest) (*indexpb.DescribeIndexResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *indexpb.DescribeIndexResponse
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.DescribeIndexRequest) *indexpb.DescribeIndexResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*indexpb.DescribeIndexResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.DescribeIndexRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_DescribeIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeIndex'
type MockIndexCoord_DescribeIndex_Call struct {
	*mock.Call
}

// DescribeIndex is a helper method to define mock.On call
//  - ctx context.Context
//  - req *indexpb.DescribeIndexRequest
func (_e *MockIndexCoord_Expecter) DescribeIndex(ctx interface{}, req interface{}) *MockIndexCoord_DescribeIndex_Call {
	return &MockIndexCoord_DescribeIndex_Call{Call: _e.mock.On("DescribeIndex", ctx, req)}
}

func (_c *MockIndexCoord_DescribeIndex_Call) Run(run func(ctx context.Context, req *indexpb.DescribeIndexRequest)) *MockIndexCoord_DescribeIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.DescribeIndexRequest))
	})
	return _c
}

func (_c *MockIndexCoord_DescribeIndex_Call) Return(_a0 *indexpb.DescribeIndexResponse, _a1 error) *MockIndexCoord_DescribeIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DropIndex provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) DropIndex(ctx context.Context, req *indexpb.DropIndexRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.DropIndexRequest) *commonpb.Status); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.DropIndexRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_DropIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropIndex'
type MockIndexCoord_DropIndex_Call struct {
	*mock.Call
}

// DropIndex is a helper method to define mock.On call
//  - ctx context.Context
//  - req *indexpb.DropIndexRequest
func (_e *MockIndexCoord_Expecter) DropIndex(ctx interface{}, req interface{}) *MockIndexCoord_DropIndex_Call {
	return &MockIndexCoord_DropIndex_Call{Call: _e.mock.On("DropIndex", ctx, req)}
}

func (_c *MockIndexCoord_DropIndex_Call) Run(run func(ctx context.Context, req *indexpb.DropIndexRequest)) *MockIndexCoord_DropIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.DropIndexRequest))
	})
	return _c
}

func (_c *MockIndexCoord_DropIndex_Call) Return(_a0 *commonpb.Status, _a1 error) *MockIndexCoord_DropIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetComponentStates provides a mock function with given fields: ctx
func (_m *MockIndexCoord) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	ret := _m.Called(ctx)

	var r0 *internalpb.ComponentStates
	if rf, ok := ret.Get(0).(func(context.Context) *internalpb.ComponentStates); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.ComponentStates)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_GetComponentStates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComponentStates'
type MockIndexCoord_GetComponentStates_Call struct {
	*mock.Call
}

// GetComponentStates is a helper method to define mock.On call
//  - ctx context.Context
func (_e *MockIndexCoord_Expecter) GetComponentStates(ctx interface{}) *MockIndexCoord_GetComponentStates_Call {
	return &MockIndexCoord_GetComponentStates_Call{Call: _e.mock.On("GetComponentStates", ctx)}
}

func (_c *MockIndexCoord_GetComponentStates_Call) Run(run func(ctx context.Context)) *MockIndexCoord_GetComponentStates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIndexCoord_GetComponentStates_Call) Return(_a0 *internalpb.ComponentStates, _a1 error) *MockIndexCoord_GetComponentStates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetIndexBuildProgress provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) GetIndexBuildProgress(ctx context.Context, req *indexpb.GetIndexBuildProgressRequest) (*indexpb.GetIndexBuildProgressResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *indexpb.GetIndexBuildProgressResponse
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.GetIndexBuildProgressRequest) *indexpb.GetIndexBuildProgressResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*indexpb.GetIndexBuildProgressResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.GetIndexBuildProgressRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_GetIndexBuildProgress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndexBuildProgress'
type MockIndexCoord_GetIndexBuildProgress_Call struct {
	*mock.Call
}

// GetIndexBuildProgress is a helper method to define mock.On call
//  - ctx context.Context
//  - req *indexpb.GetIndexBuildProgressRequest
func (_e *MockIndexCoord_Expecter) GetIndexBuildProgress(ctx interface{}, req interface{}) *MockIndexCoord_GetIndexBuildProgress_Call {
	return &MockIndexCoord_GetIndexBuildProgress_Call{Call: _e.mock.On("GetIndexBuildProgress", ctx, req)}
}

func (_c *MockIndexCoord_GetIndexBuildProgress_Call) Run(run func(ctx context.Context, req *indexpb.GetIndexBuildProgressRequest)) *MockIndexCoord_GetIndexBuildProgress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.GetIndexBuildProgressRequest))
	})
	return _c
}

func (_c *MockIndexCoord_GetIndexBuildProgress_Call) Return(_a0 *indexpb.GetIndexBuildProgressResponse, _a1 error) *MockIndexCoord_GetIndexBuildProgress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetIndexInfos provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) GetIndexInfos(ctx context.Context, req *indexpb.GetIndexInfoRequest) (*indexpb.GetIndexInfoResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *indexpb.GetIndexInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.GetIndexInfoRequest) *indexpb.GetIndexInfoResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*indexpb.GetIndexInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.GetIndexInfoRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_GetIndexInfos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndexInfos'
type MockIndexCoord_GetIndexInfos_Call struct {
	*mock.Call
}

// GetIndexInfos is a helper method to define mock.On call
//  - ctx context.Context
//  - req *indexpb.GetIndexInfoRequest
func (_e *MockIndexCoord_Expecter) GetIndexInfos(ctx interface{}, req interface{}) *MockIndexCoord_GetIndexInfos_Call {
	return &MockIndexCoord_GetIndexInfos_Call{Call: _e.mock.On("GetIndexInfos", ctx, req)}
}

func (_c *MockIndexCoord_GetIndexInfos_Call) Run(run func(ctx context.Context, req *indexpb.GetIndexInfoRequest)) *MockIndexCoord_GetIndexInfos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.GetIndexInfoRequest))
	})
	return _c
}

func (_c *MockIndexCoord_GetIndexInfos_Call) Return(_a0 *indexpb.GetIndexInfoResponse, _a1 error) *MockIndexCoord_GetIndexInfos_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetIndexState provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) GetIndexState(ctx context.Context, req *indexpb.GetIndexStateRequest) (*indexpb.GetIndexStateResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *indexpb.GetIndexStateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.GetIndexStateRequest) *indexpb.GetIndexStateResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*indexpb.GetIndexStateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.GetIndexStateRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_GetIndexState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndexState'
type MockIndexCoord_GetIndexState_Call struct {
	*mock.Call
}

// GetIndexState is a helper method to define mock.On call
//  - ctx context.Context
//  - req *indexpb.GetIndexStateRequest
func (_e *MockIndexCoord_Expecter) GetIndexState(ctx interface{}, req interface{}) *MockIndexCoord_GetIndexState_Call {
	return &MockIndexCoord_GetIndexState_Call{Call: _e.mock.On("GetIndexState", ctx, req)}
}

func (_c *MockIndexCoord_GetIndexState_Call) Run(run func(ctx context.Context, req *indexpb.GetIndexStateRequest)) *MockIndexCoord_GetIndexState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.GetIndexStateRequest))
	})
	return _c
}

func (_c *MockIndexCoord_GetIndexState_Call) Return(_a0 *indexpb.GetIndexStateResponse, _a1 error) *MockIndexCoord_GetIndexState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetMetrics provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) GetMetrics(ctx context.Context, req *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *milvuspb.GetMetricsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetMetricsRequest) *milvuspb.GetMetricsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.GetMetricsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *milvuspb.GetMetricsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_GetMetrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetrics'
type MockIndexCoord_GetMetrics_Call struct {
	*mock.Call
}

// GetMetrics is a helper method to define mock.On call
//  - ctx context.Context
//  - req *milvuspb.GetMetricsRequest
func (_e *MockIndexCoord_Expecter) GetMetrics(ctx interface{}, req interface{}) *MockIndexCoord_GetMetrics_Call {
	return &MockIndexCoord_GetMetrics_Call{Call: _e.mock.On("GetMetrics", ctx, req)}
}

func (_c *MockIndexCoord_GetMetrics_Call) Run(run func(ctx context.Context, req *milvuspb.GetMetricsRequest)) *MockIndexCoord_GetMetrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*milvuspb.GetMetricsRequest))
	})
	return _c
}

func (_c *MockIndexCoord_GetMetrics_Call) Return(_a0 *milvuspb.GetMetricsResponse, _a1 error) *MockIndexCoord_GetMetrics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSegmentIndexState provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) GetSegmentIndexState(ctx context.Context, req *indexpb.GetSegmentIndexStateRequest) (*indexpb.GetSegmentIndexStateResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *indexpb.GetSegmentIndexStateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *indexpb.GetSegmentIndexStateRequest) *indexpb.GetSegmentIndexStateResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*indexpb.GetSegmentIndexStateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *indexpb.GetSegmentIndexStateRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_GetSegmentIndexState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentIndexState'
type MockIndexCoord_GetSegmentIndexState_Call struct {
	*mock.Call
}

// GetSegmentIndexState is a helper method to define mock.On call
//  - ctx context.Context
//  - req *indexpb.GetSegmentIndexStateRequest
func (_e *MockIndexCoord_Expecter) GetSegmentIndexState(ctx interface{}, req interface{}) *MockIndexCoord_GetSegmentIndexState_Call {
	return &MockIndexCoord_GetSegmentIndexState_Call{Call: _e.mock.On("GetSegmentIndexState", ctx, req)}
}

func (_c *MockIndexCoord_GetSegmentIndexState_Call) Run(run func(ctx context.Context, req *indexpb.GetSegmentIndexStateRequest)) *MockIndexCoord_GetSegmentIndexState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*indexpb.GetSegmentIndexStateRequest))
	})
	return _c
}

func (_c *MockIndexCoord_GetSegmentIndexState_Call) Return(_a0 *indexpb.GetSegmentIndexStateResponse, _a1 error) *MockIndexCoord_GetSegmentIndexState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetStatisticsChannel provides a mock function with given fields: ctx
func (_m *MockIndexCoord) GetStatisticsChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	ret := _m.Called(ctx)

	var r0 *milvuspb.StringResponse
	if rf, ok := ret.Get(0).(func(context.Context) *milvuspb.StringResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.StringResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_GetStatisticsChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatisticsChannel'
type MockIndexCoord_GetStatisticsChannel_Call struct {
	*mock.Call
}

// GetStatisticsChannel is a helper method to define mock.On call
//  - ctx context.Context
func (_e *MockIndexCoord_Expecter) GetStatisticsChannel(ctx interface{}) *MockIndexCoord_GetStatisticsChannel_Call {
	return &MockIndexCoord_GetStatisticsChannel_Call{Call: _e.mock.On("GetStatisticsChannel", ctx)}
}

func (_c *MockIndexCoord_GetStatisticsChannel_Call) Run(run func(ctx context.Context)) *MockIndexCoord_GetStatisticsChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIndexCoord_GetStatisticsChannel_Call) Return(_a0 *milvuspb.StringResponse, _a1 error) *MockIndexCoord_GetStatisticsChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Init provides a mock function with given fields:
func (_m *MockIndexCoord) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIndexCoord_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockIndexCoord_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *MockIndexCoord_Expecter) Init() *MockIndexCoord_Init_Call {
	return &MockIndexCoord_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *MockIndexCoord_Init_Call) Run(run func()) *MockIndexCoord_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexCoord_Init_Call) Return(_a0 error) *MockIndexCoord_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

// Register provides a mock function with given fields:
func (_m *MockIndexCoord) Register() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIndexCoord_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockIndexCoord_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
func (_e *MockIndexCoord_Expecter) Register() *MockIndexCoord_Register_Call {
	return &MockIndexCoord_Register_Call{Call: _e.mock.On("Register")}
}

func (_c *MockIndexCoord_Register_Call) Run(run func()) *MockIndexCoord_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexCoord_Register_Call) Return(_a0 error) *MockIndexCoord_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

// ShowConfigurations provides a mock function with given fields: ctx, req
func (_m *MockIndexCoord) ShowConfigurations(ctx context.Context, req *internalpb.ShowConfigurationsRequest) (*internalpb.ShowConfigurationsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *internalpb.ShowConfigurationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.ShowConfigurationsRequest) *internalpb.ShowConfigurationsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.ShowConfigurationsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *internalpb.ShowConfigurationsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIndexCoord_ShowConfigurations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShowConfigurations'
type MockIndexCoord_ShowConfigurations_Call struct {
	*mock.Call
}

// ShowConfigurations is a helper method to define mock.On call
//  - ctx context.Context
//  - req *internalpb.ShowConfigurationsRequest
func (_e *MockIndexCoord_Expecter) ShowConfigurations(ctx interface{}, req interface{}) *MockIndexCoord_ShowConfigurations_Call {
	return &MockIndexCoord_ShowConfigurations_Call{Call: _e.mock.On("ShowConfigurations", ctx, req)}
}

func (_c *MockIndexCoord_ShowConfigurations_Call) Run(run func(ctx context.Context, req *internalpb.ShowConfigurationsRequest)) *MockIndexCoord_ShowConfigurations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internalpb.ShowConfigurationsRequest))
	})
	return _c
}

func (_c *MockIndexCoord_ShowConfigurations_Call) Return(_a0 *internalpb.ShowConfigurationsResponse, _a1 error) *MockIndexCoord_ShowConfigurations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockIndexCoord) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIndexCoord_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockIndexCoord_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockIndexCoord_Expecter) Start() *MockIndexCoord_Start_Call {
	return &MockIndexCoord_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockIndexCoord_Start_Call) Run(run func()) *MockIndexCoord_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexCoord_Start_Call) Return(_a0 error) *MockIndexCoord_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockIndexCoord) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIndexCoord_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockIndexCoord_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockIndexCoord_Expecter) Stop() *MockIndexCoord_Stop_Call {
	return &MockIndexCoord_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockIndexCoord_Stop_Call) Run(run func()) *MockIndexCoord_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIndexCoord_Stop_Call) Return(_a0 error) *MockIndexCoord_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockIndexCoord interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIndexCoord creates a new instance of MockIndexCoord. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIndexCoord(t mockConstructorTestingTNewMockIndexCoord) *MockIndexCoord {
	mock := &MockIndexCoord{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
