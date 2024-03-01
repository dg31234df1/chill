// Code generated by mockery v2.30.1. DO NOT EDIT.

package datacoord

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockManager is an autogenerated mock type for the Manager type
type MockManager struct {
	mock.Mock
}

type MockManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockManager) EXPECT() *MockManager_Expecter {
	return &MockManager_Expecter{mock: &_m.Mock}
}

// AllocImportSegment provides a mock function with given fields: ctx, taskID, collectionID, partitionID, channelName
func (_m *MockManager) AllocImportSegment(ctx context.Context, taskID int64, collectionID int64, partitionID int64, channelName string) (*SegmentInfo, error) {
	ret := _m.Called(ctx, taskID, collectionID, partitionID, channelName)

	var r0 *SegmentInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64, string) (*SegmentInfo, error)); ok {
		return rf(ctx, taskID, collectionID, partitionID, channelName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64, string) *SegmentInfo); ok {
		r0 = rf(ctx, taskID, collectionID, partitionID, channelName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SegmentInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, int64, string) error); ok {
		r1 = rf(ctx, taskID, collectionID, partitionID, channelName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_AllocImportSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllocImportSegment'
type MockManager_AllocImportSegment_Call struct {
	*mock.Call
}

// AllocImportSegment is a helper method to define mock.On call
//  - ctx context.Context
//  - taskID int64
//  - collectionID int64
//  - partitionID int64
//  - channelName string
func (_e *MockManager_Expecter) AllocImportSegment(ctx interface{}, taskID interface{}, collectionID interface{}, partitionID interface{}, channelName interface{}) *MockManager_AllocImportSegment_Call {
	return &MockManager_AllocImportSegment_Call{Call: _e.mock.On("AllocImportSegment", ctx, taskID, collectionID, partitionID, channelName)}
}

func (_c *MockManager_AllocImportSegment_Call) Run(run func(ctx context.Context, taskID int64, collectionID int64, partitionID int64, channelName string)) *MockManager_AllocImportSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(string))
	})
	return _c
}

func (_c *MockManager_AllocImportSegment_Call) Return(_a0 *SegmentInfo, _a1 error) *MockManager_AllocImportSegment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_AllocImportSegment_Call) RunAndReturn(run func(context.Context, int64, int64, int64, string) (*SegmentInfo, error)) *MockManager_AllocImportSegment_Call {
	_c.Call.Return(run)
	return _c
}

// AllocSegment provides a mock function with given fields: ctx, collectionID, partitionID, channelName, requestRows
func (_m *MockManager) AllocSegment(ctx context.Context, collectionID int64, partitionID int64, channelName string, requestRows int64) ([]*Allocation, error) {
	ret := _m.Called(ctx, collectionID, partitionID, channelName, requestRows)

	var r0 []*Allocation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, int64) ([]*Allocation, error)); ok {
		return rf(ctx, collectionID, partitionID, channelName, requestRows)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, int64) []*Allocation); ok {
		r0 = rf(ctx, collectionID, partitionID, channelName, requestRows)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Allocation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, string, int64) error); ok {
		r1 = rf(ctx, collectionID, partitionID, channelName, requestRows)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_AllocSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllocSegment'
type MockManager_AllocSegment_Call struct {
	*mock.Call
}

// AllocSegment is a helper method to define mock.On call
//  - ctx context.Context
//  - collectionID int64
//  - partitionID int64
//  - channelName string
//  - requestRows int64
func (_e *MockManager_Expecter) AllocSegment(ctx interface{}, collectionID interface{}, partitionID interface{}, channelName interface{}, requestRows interface{}) *MockManager_AllocSegment_Call {
	return &MockManager_AllocSegment_Call{Call: _e.mock.On("AllocSegment", ctx, collectionID, partitionID, channelName, requestRows)}
}

func (_c *MockManager_AllocSegment_Call) Run(run func(ctx context.Context, collectionID int64, partitionID int64, channelName string, requestRows int64)) *MockManager_AllocSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64), args[3].(string), args[4].(int64))
	})
	return _c
}

func (_c *MockManager_AllocSegment_Call) Return(_a0 []*Allocation, _a1 error) *MockManager_AllocSegment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_AllocSegment_Call) RunAndReturn(run func(context.Context, int64, int64, string, int64) ([]*Allocation, error)) *MockManager_AllocSegment_Call {
	_c.Call.Return(run)
	return _c
}

// DropSegment provides a mock function with given fields: ctx, segmentID
func (_m *MockManager) DropSegment(ctx context.Context, segmentID int64) {
	_m.Called(ctx, segmentID)
}

// MockManager_DropSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropSegment'
type MockManager_DropSegment_Call struct {
	*mock.Call
}

// DropSegment is a helper method to define mock.On call
//  - ctx context.Context
//  - segmentID int64
func (_e *MockManager_Expecter) DropSegment(ctx interface{}, segmentID interface{}) *MockManager_DropSegment_Call {
	return &MockManager_DropSegment_Call{Call: _e.mock.On("DropSegment", ctx, segmentID)}
}

func (_c *MockManager_DropSegment_Call) Run(run func(ctx context.Context, segmentID int64)) *MockManager_DropSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockManager_DropSegment_Call) Return() *MockManager_DropSegment_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_DropSegment_Call) RunAndReturn(run func(context.Context, int64)) *MockManager_DropSegment_Call {
	_c.Call.Return(run)
	return _c
}

// DropSegmentsOfChannel provides a mock function with given fields: ctx, channel
func (_m *MockManager) DropSegmentsOfChannel(ctx context.Context, channel string) {
	_m.Called(ctx, channel)
}

// MockManager_DropSegmentsOfChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropSegmentsOfChannel'
type MockManager_DropSegmentsOfChannel_Call struct {
	*mock.Call
}

// DropSegmentsOfChannel is a helper method to define mock.On call
//  - ctx context.Context
//  - channel string
func (_e *MockManager_Expecter) DropSegmentsOfChannel(ctx interface{}, channel interface{}) *MockManager_DropSegmentsOfChannel_Call {
	return &MockManager_DropSegmentsOfChannel_Call{Call: _e.mock.On("DropSegmentsOfChannel", ctx, channel)}
}

func (_c *MockManager_DropSegmentsOfChannel_Call) Run(run func(ctx context.Context, channel string)) *MockManager_DropSegmentsOfChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockManager_DropSegmentsOfChannel_Call) Return() *MockManager_DropSegmentsOfChannel_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_DropSegmentsOfChannel_Call) RunAndReturn(run func(context.Context, string)) *MockManager_DropSegmentsOfChannel_Call {
	_c.Call.Return(run)
	return _c
}

// ExpireAllocations provides a mock function with given fields: channel, ts
func (_m *MockManager) ExpireAllocations(channel string, ts uint64) error {
	ret := _m.Called(channel, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint64) error); ok {
		r0 = rf(channel, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManager_ExpireAllocations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpireAllocations'
type MockManager_ExpireAllocations_Call struct {
	*mock.Call
}

// ExpireAllocations is a helper method to define mock.On call
//  - channel string
//  - ts uint64
func (_e *MockManager_Expecter) ExpireAllocations(channel interface{}, ts interface{}) *MockManager_ExpireAllocations_Call {
	return &MockManager_ExpireAllocations_Call{Call: _e.mock.On("ExpireAllocations", channel, ts)}
}

func (_c *MockManager_ExpireAllocations_Call) Run(run func(channel string, ts uint64)) *MockManager_ExpireAllocations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint64))
	})
	return _c
}

func (_c *MockManager_ExpireAllocations_Call) Return(_a0 error) *MockManager_ExpireAllocations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_ExpireAllocations_Call) RunAndReturn(run func(string, uint64) error) *MockManager_ExpireAllocations_Call {
	_c.Call.Return(run)
	return _c
}

// FlushImportSegments provides a mock function with given fields: ctx, collectionID, segmentIDs
func (_m *MockManager) FlushImportSegments(ctx context.Context, collectionID int64, segmentIDs []int64) error {
	ret := _m.Called(ctx, collectionID, segmentIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []int64) error); ok {
		r0 = rf(ctx, collectionID, segmentIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManager_FlushImportSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FlushImportSegments'
type MockManager_FlushImportSegments_Call struct {
	*mock.Call
}

// FlushImportSegments is a helper method to define mock.On call
//  - ctx context.Context
//  - collectionID int64
//  - segmentIDs []int64
func (_e *MockManager_Expecter) FlushImportSegments(ctx interface{}, collectionID interface{}, segmentIDs interface{}) *MockManager_FlushImportSegments_Call {
	return &MockManager_FlushImportSegments_Call{Call: _e.mock.On("FlushImportSegments", ctx, collectionID, segmentIDs)}
}

func (_c *MockManager_FlushImportSegments_Call) Run(run func(ctx context.Context, collectionID int64, segmentIDs []int64)) *MockManager_FlushImportSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].([]int64))
	})
	return _c
}

func (_c *MockManager_FlushImportSegments_Call) Return(_a0 error) *MockManager_FlushImportSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_FlushImportSegments_Call) RunAndReturn(run func(context.Context, int64, []int64) error) *MockManager_FlushImportSegments_Call {
	_c.Call.Return(run)
	return _c
}

// GetFlushableSegments provides a mock function with given fields: ctx, channel, ts
func (_m *MockManager) GetFlushableSegments(ctx context.Context, channel string, ts uint64) ([]int64, error) {
	ret := _m.Called(ctx, channel, ts)

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) ([]int64, error)); ok {
		return rf(ctx, channel, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) []int64); ok {
		r0 = rf(ctx, channel, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint64) error); ok {
		r1 = rf(ctx, channel, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_GetFlushableSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlushableSegments'
type MockManager_GetFlushableSegments_Call struct {
	*mock.Call
}

// GetFlushableSegments is a helper method to define mock.On call
//  - ctx context.Context
//  - channel string
//  - ts uint64
func (_e *MockManager_Expecter) GetFlushableSegments(ctx interface{}, channel interface{}, ts interface{}) *MockManager_GetFlushableSegments_Call {
	return &MockManager_GetFlushableSegments_Call{Call: _e.mock.On("GetFlushableSegments", ctx, channel, ts)}
}

func (_c *MockManager_GetFlushableSegments_Call) Run(run func(ctx context.Context, channel string, ts uint64)) *MockManager_GetFlushableSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *MockManager_GetFlushableSegments_Call) Return(_a0 []int64, _a1 error) *MockManager_GetFlushableSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_GetFlushableSegments_Call) RunAndReturn(run func(context.Context, string, uint64) ([]int64, error)) *MockManager_GetFlushableSegments_Call {
	_c.Call.Return(run)
	return _c
}

// SealAllSegments provides a mock function with given fields: ctx, collectionID, segIDs
func (_m *MockManager) SealAllSegments(ctx context.Context, collectionID int64, segIDs []int64) ([]int64, error) {
	ret := _m.Called(ctx, collectionID, segIDs)

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []int64) ([]int64, error)); ok {
		return rf(ctx, collectionID, segIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, []int64) []int64); ok {
		r0 = rf(ctx, collectionID, segIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, []int64) error); ok {
		r1 = rf(ctx, collectionID, segIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_SealAllSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SealAllSegments'
type MockManager_SealAllSegments_Call struct {
	*mock.Call
}

// SealAllSegments is a helper method to define mock.On call
//  - ctx context.Context
//  - collectionID int64
//  - segIDs []int64
func (_e *MockManager_Expecter) SealAllSegments(ctx interface{}, collectionID interface{}, segIDs interface{}) *MockManager_SealAllSegments_Call {
	return &MockManager_SealAllSegments_Call{Call: _e.mock.On("SealAllSegments", ctx, collectionID, segIDs)}
}

func (_c *MockManager_SealAllSegments_Call) Run(run func(ctx context.Context, collectionID int64, segIDs []int64)) *MockManager_SealAllSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].([]int64))
	})
	return _c
}

func (_c *MockManager_SealAllSegments_Call) Return(_a0 []int64, _a1 error) *MockManager_SealAllSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_SealAllSegments_Call) RunAndReturn(run func(context.Context, int64, []int64) ([]int64, error)) *MockManager_SealAllSegments_Call {
	_c.Call.Return(run)
	return _c
}

// allocSegmentForImport provides a mock function with given fields: ctx, collectionID, partitionID, channelName, requestRows, taskID
func (_m *MockManager) allocSegmentForImport(ctx context.Context, collectionID int64, partitionID int64, channelName string, requestRows int64, taskID int64) (*Allocation, error) {
	ret := _m.Called(ctx, collectionID, partitionID, channelName, requestRows, taskID)

	var r0 *Allocation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, int64, int64) (*Allocation, error)); ok {
		return rf(ctx, collectionID, partitionID, channelName, requestRows, taskID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, int64, int64) *Allocation); ok {
		r0 = rf(ctx, collectionID, partitionID, channelName, requestRows, taskID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Allocation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, string, int64, int64) error); ok {
		r1 = rf(ctx, collectionID, partitionID, channelName, requestRows, taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManager_allocSegmentForImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'allocSegmentForImport'
type MockManager_allocSegmentForImport_Call struct {
	*mock.Call
}

// allocSegmentForImport is a helper method to define mock.On call
//  - ctx context.Context
//  - collectionID int64
//  - partitionID int64
//  - channelName string
//  - requestRows int64
//  - taskID int64
func (_e *MockManager_Expecter) allocSegmentForImport(ctx interface{}, collectionID interface{}, partitionID interface{}, channelName interface{}, requestRows interface{}, taskID interface{}) *MockManager_allocSegmentForImport_Call {
	return &MockManager_allocSegmentForImport_Call{Call: _e.mock.On("allocSegmentForImport", ctx, collectionID, partitionID, channelName, requestRows, taskID)}
}

func (_c *MockManager_allocSegmentForImport_Call) Run(run func(ctx context.Context, collectionID int64, partitionID int64, channelName string, requestRows int64, taskID int64)) *MockManager_allocSegmentForImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64), args[3].(string), args[4].(int64), args[5].(int64))
	})
	return _c
}

func (_c *MockManager_allocSegmentForImport_Call) Return(_a0 *Allocation, _a1 error) *MockManager_allocSegmentForImport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_allocSegmentForImport_Call) RunAndReturn(run func(context.Context, int64, int64, string, int64, int64) (*Allocation, error)) *MockManager_allocSegmentForImport_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockManager creates a new instance of MockManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockManager {
	mock := &MockManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
