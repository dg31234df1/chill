// Code generated by mockery v2.32.4. DO NOT EDIT.

package broker

import (
	context "context"

	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"

	msgpb "github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
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

// AssignSegmentID provides a mock function with given fields: ctx, reqs
func (_m *MockBroker) AssignSegmentID(ctx context.Context, reqs ...*datapb.SegmentIDRequest) ([]int64, error) {
	_va := make([]interface{}, len(reqs))
	for _i := range reqs {
		_va[_i] = reqs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*datapb.SegmentIDRequest) ([]int64, error)); ok {
		return rf(ctx, reqs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...*datapb.SegmentIDRequest) []int64); ok {
		r0 = rf(ctx, reqs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...*datapb.SegmentIDRequest) error); ok {
		r1 = rf(ctx, reqs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroker_AssignSegmentID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AssignSegmentID'
type MockBroker_AssignSegmentID_Call struct {
	*mock.Call
}

// AssignSegmentID is a helper method to define mock.On call
//   - ctx context.Context
//   - reqs ...*datapb.SegmentIDRequest
func (_e *MockBroker_Expecter) AssignSegmentID(ctx interface{}, reqs ...interface{}) *MockBroker_AssignSegmentID_Call {
	return &MockBroker_AssignSegmentID_Call{Call: _e.mock.On("AssignSegmentID",
		append([]interface{}{ctx}, reqs...)...)}
}

func (_c *MockBroker_AssignSegmentID_Call) Run(run func(ctx context.Context, reqs ...*datapb.SegmentIDRequest)) *MockBroker_AssignSegmentID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*datapb.SegmentIDRequest, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*datapb.SegmentIDRequest)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockBroker_AssignSegmentID_Call) Return(_a0 []int64, _a1 error) *MockBroker_AssignSegmentID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_AssignSegmentID_Call) RunAndReturn(run func(context.Context, ...*datapb.SegmentIDRequest) ([]int64, error)) *MockBroker_AssignSegmentID_Call {
	_c.Call.Return(run)
	return _c
}

// DropVirtualChannel provides a mock function with given fields: ctx, req
func (_m *MockBroker) DropVirtualChannel(ctx context.Context, req *datapb.DropVirtualChannelRequest) (*datapb.DropVirtualChannelResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *datapb.DropVirtualChannelResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.DropVirtualChannelRequest) (*datapb.DropVirtualChannelResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.DropVirtualChannelRequest) *datapb.DropVirtualChannelResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datapb.DropVirtualChannelResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *datapb.DropVirtualChannelRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBroker_DropVirtualChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropVirtualChannel'
type MockBroker_DropVirtualChannel_Call struct {
	*mock.Call
}

// DropVirtualChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - req *datapb.DropVirtualChannelRequest
func (_e *MockBroker_Expecter) DropVirtualChannel(ctx interface{}, req interface{}) *MockBroker_DropVirtualChannel_Call {
	return &MockBroker_DropVirtualChannel_Call{Call: _e.mock.On("DropVirtualChannel", ctx, req)}
}

func (_c *MockBroker_DropVirtualChannel_Call) Run(run func(ctx context.Context, req *datapb.DropVirtualChannelRequest)) *MockBroker_DropVirtualChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*datapb.DropVirtualChannelRequest))
	})
	return _c
}

func (_c *MockBroker_DropVirtualChannel_Call) Return(_a0 *datapb.DropVirtualChannelResponse, _a1 error) *MockBroker_DropVirtualChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_DropVirtualChannel_Call) RunAndReturn(run func(context.Context, *datapb.DropVirtualChannelRequest) (*datapb.DropVirtualChannelResponse, error)) *MockBroker_DropVirtualChannel_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegmentInfo provides a mock function with given fields: ctx, segmentIDs
func (_m *MockBroker) GetSegmentInfo(ctx context.Context, segmentIDs []int64) ([]*datapb.SegmentInfo, error) {
	ret := _m.Called(ctx, segmentIDs)

	var r0 []*datapb.SegmentInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) ([]*datapb.SegmentInfo, error)); ok {
		return rf(ctx, segmentIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int64) []*datapb.SegmentInfo); ok {
		r0 = rf(ctx, segmentIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datapb.SegmentInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(ctx, segmentIDs)
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
//   - segmentIDs []int64
func (_e *MockBroker_Expecter) GetSegmentInfo(ctx interface{}, segmentIDs interface{}) *MockBroker_GetSegmentInfo_Call {
	return &MockBroker_GetSegmentInfo_Call{Call: _e.mock.On("GetSegmentInfo", ctx, segmentIDs)}
}

func (_c *MockBroker_GetSegmentInfo_Call) Run(run func(ctx context.Context, segmentIDs []int64)) *MockBroker_GetSegmentInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]int64))
	})
	return _c
}

func (_c *MockBroker_GetSegmentInfo_Call) Return(_a0 []*datapb.SegmentInfo, _a1 error) *MockBroker_GetSegmentInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBroker_GetSegmentInfo_Call) RunAndReturn(run func(context.Context, []int64) ([]*datapb.SegmentInfo, error)) *MockBroker_GetSegmentInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ReportTimeTick provides a mock function with given fields: ctx, msgs
func (_m *MockBroker) ReportTimeTick(ctx context.Context, msgs []*msgpb.DataNodeTtMsg) error {
	ret := _m.Called(ctx, msgs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*msgpb.DataNodeTtMsg) error); ok {
		r0 = rf(ctx, msgs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBroker_ReportTimeTick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReportTimeTick'
type MockBroker_ReportTimeTick_Call struct {
	*mock.Call
}

// ReportTimeTick is a helper method to define mock.On call
//   - ctx context.Context
//   - msgs []*msgpb.DataNodeTtMsg
func (_e *MockBroker_Expecter) ReportTimeTick(ctx interface{}, msgs interface{}) *MockBroker_ReportTimeTick_Call {
	return &MockBroker_ReportTimeTick_Call{Call: _e.mock.On("ReportTimeTick", ctx, msgs)}
}

func (_c *MockBroker_ReportTimeTick_Call) Run(run func(ctx context.Context, msgs []*msgpb.DataNodeTtMsg)) *MockBroker_ReportTimeTick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*msgpb.DataNodeTtMsg))
	})
	return _c
}

func (_c *MockBroker_ReportTimeTick_Call) Return(_a0 error) *MockBroker_ReportTimeTick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBroker_ReportTimeTick_Call) RunAndReturn(run func(context.Context, []*msgpb.DataNodeTtMsg) error) *MockBroker_ReportTimeTick_Call {
	_c.Call.Return(run)
	return _c
}

// SaveBinlogPaths provides a mock function with given fields: ctx, req
func (_m *MockBroker) SaveBinlogPaths(ctx context.Context, req *datapb.SaveBinlogPathsRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.SaveBinlogPathsRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBroker_SaveBinlogPaths_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveBinlogPaths'
type MockBroker_SaveBinlogPaths_Call struct {
	*mock.Call
}

// SaveBinlogPaths is a helper method to define mock.On call
//   - ctx context.Context
//   - req *datapb.SaveBinlogPathsRequest
func (_e *MockBroker_Expecter) SaveBinlogPaths(ctx interface{}, req interface{}) *MockBroker_SaveBinlogPaths_Call {
	return &MockBroker_SaveBinlogPaths_Call{Call: _e.mock.On("SaveBinlogPaths", ctx, req)}
}

func (_c *MockBroker_SaveBinlogPaths_Call) Run(run func(ctx context.Context, req *datapb.SaveBinlogPathsRequest)) *MockBroker_SaveBinlogPaths_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*datapb.SaveBinlogPathsRequest))
	})
	return _c
}

func (_c *MockBroker_SaveBinlogPaths_Call) Return(_a0 error) *MockBroker_SaveBinlogPaths_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBroker_SaveBinlogPaths_Call) RunAndReturn(run func(context.Context, *datapb.SaveBinlogPathsRequest) error) *MockBroker_SaveBinlogPaths_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateChannelCheckpoint provides a mock function with given fields: ctx, channelCPs
func (_m *MockBroker) UpdateChannelCheckpoint(ctx context.Context, channelCPs []*msgpb.MsgPosition) error {
	ret := _m.Called(ctx, channelCPs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*msgpb.MsgPosition) error); ok {
		r0 = rf(ctx, channelCPs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBroker_UpdateChannelCheckpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateChannelCheckpoint'
type MockBroker_UpdateChannelCheckpoint_Call struct {
	*mock.Call
}

// UpdateChannelCheckpoint is a helper method to define mock.On call
//   - ctx context.Context
//   - channelCPs []*msgpb.MsgPosition
func (_e *MockBroker_Expecter) UpdateChannelCheckpoint(ctx interface{}, channelCPs interface{}) *MockBroker_UpdateChannelCheckpoint_Call {
	return &MockBroker_UpdateChannelCheckpoint_Call{Call: _e.mock.On("UpdateChannelCheckpoint", ctx, channelCPs)}
}

func (_c *MockBroker_UpdateChannelCheckpoint_Call) Run(run func(ctx context.Context, channelCPs []*msgpb.MsgPosition)) *MockBroker_UpdateChannelCheckpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*msgpb.MsgPosition))
	})
	return _c
}

func (_c *MockBroker_UpdateChannelCheckpoint_Call) Return(_a0 error) *MockBroker_UpdateChannelCheckpoint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBroker_UpdateChannelCheckpoint_Call) RunAndReturn(run func(context.Context, []*msgpb.MsgPosition) error) *MockBroker_UpdateChannelCheckpoint_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSegmentStatistics provides a mock function with given fields: ctx, req
func (_m *MockBroker) UpdateSegmentStatistics(ctx context.Context, req *datapb.UpdateSegmentStatisticsRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.UpdateSegmentStatisticsRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBroker_UpdateSegmentStatistics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSegmentStatistics'
type MockBroker_UpdateSegmentStatistics_Call struct {
	*mock.Call
}

// UpdateSegmentStatistics is a helper method to define mock.On call
//   - ctx context.Context
//   - req *datapb.UpdateSegmentStatisticsRequest
func (_e *MockBroker_Expecter) UpdateSegmentStatistics(ctx interface{}, req interface{}) *MockBroker_UpdateSegmentStatistics_Call {
	return &MockBroker_UpdateSegmentStatistics_Call{Call: _e.mock.On("UpdateSegmentStatistics", ctx, req)}
}

func (_c *MockBroker_UpdateSegmentStatistics_Call) Run(run func(ctx context.Context, req *datapb.UpdateSegmentStatisticsRequest)) *MockBroker_UpdateSegmentStatistics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*datapb.UpdateSegmentStatisticsRequest))
	})
	return _c
}

func (_c *MockBroker_UpdateSegmentStatistics_Call) Return(_a0 error) *MockBroker_UpdateSegmentStatistics_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBroker_UpdateSegmentStatistics_Call) RunAndReturn(run func(context.Context, *datapb.UpdateSegmentStatisticsRequest) error) *MockBroker_UpdateSegmentStatistics_Call {
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
