// Code generated by mockery v2.46.0. DO NOT EDIT.

package datacoord

import (
	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"
)

// MockCompactionMeta is an autogenerated mock type for the CompactionMeta type
type MockCompactionMeta struct {
	mock.Mock
}

type MockCompactionMeta_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCompactionMeta) EXPECT() *MockCompactionMeta_Expecter {
	return &MockCompactionMeta_Expecter{mock: &_m.Mock}
}

// CheckAndSetSegmentsCompacting provides a mock function with given fields: segmentIDs
func (_m *MockCompactionMeta) CheckAndSetSegmentsCompacting(segmentIDs []int64) (bool, bool) {
	ret := _m.Called(segmentIDs)

	if len(ret) == 0 {
		panic("no return value specified for CheckAndSetSegmentsCompacting")
	}

	var r0 bool
	var r1 bool
	if rf, ok := ret.Get(0).(func([]int64) (bool, bool)); ok {
		return rf(segmentIDs)
	}
	if rf, ok := ret.Get(0).(func([]int64) bool); ok {
		r0 = rf(segmentIDs)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func([]int64) bool); ok {
		r1 = rf(segmentIDs)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockCompactionMeta_CheckAndSetSegmentsCompacting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckAndSetSegmentsCompacting'
type MockCompactionMeta_CheckAndSetSegmentsCompacting_Call struct {
	*mock.Call
}

// CheckAndSetSegmentsCompacting is a helper method to define mock.On call
//   - segmentIDs []int64
func (_e *MockCompactionMeta_Expecter) CheckAndSetSegmentsCompacting(segmentIDs interface{}) *MockCompactionMeta_CheckAndSetSegmentsCompacting_Call {
	return &MockCompactionMeta_CheckAndSetSegmentsCompacting_Call{Call: _e.mock.On("CheckAndSetSegmentsCompacting", segmentIDs)}
}

func (_c *MockCompactionMeta_CheckAndSetSegmentsCompacting_Call) Run(run func(segmentIDs []int64)) *MockCompactionMeta_CheckAndSetSegmentsCompacting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]int64))
	})
	return _c
}

func (_c *MockCompactionMeta_CheckAndSetSegmentsCompacting_Call) Return(_a0 bool, _a1 bool) *MockCompactionMeta_CheckAndSetSegmentsCompacting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCompactionMeta_CheckAndSetSegmentsCompacting_Call) RunAndReturn(run func([]int64) (bool, bool)) *MockCompactionMeta_CheckAndSetSegmentsCompacting_Call {
	_c.Call.Return(run)
	return _c
}

// CleanPartitionStatsInfo provides a mock function with given fields: info
func (_m *MockCompactionMeta) CleanPartitionStatsInfo(info *datapb.PartitionStatsInfo) error {
	ret := _m.Called(info)

	if len(ret) == 0 {
		panic("no return value specified for CleanPartitionStatsInfo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*datapb.PartitionStatsInfo) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCompactionMeta_CleanPartitionStatsInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CleanPartitionStatsInfo'
type MockCompactionMeta_CleanPartitionStatsInfo_Call struct {
	*mock.Call
}

// CleanPartitionStatsInfo is a helper method to define mock.On call
//   - info *datapb.PartitionStatsInfo
func (_e *MockCompactionMeta_Expecter) CleanPartitionStatsInfo(info interface{}) *MockCompactionMeta_CleanPartitionStatsInfo_Call {
	return &MockCompactionMeta_CleanPartitionStatsInfo_Call{Call: _e.mock.On("CleanPartitionStatsInfo", info)}
}

func (_c *MockCompactionMeta_CleanPartitionStatsInfo_Call) Run(run func(info *datapb.PartitionStatsInfo)) *MockCompactionMeta_CleanPartitionStatsInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*datapb.PartitionStatsInfo))
	})
	return _c
}

func (_c *MockCompactionMeta_CleanPartitionStatsInfo_Call) Return(_a0 error) *MockCompactionMeta_CleanPartitionStatsInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_CleanPartitionStatsInfo_Call) RunAndReturn(run func(*datapb.PartitionStatsInfo) error) *MockCompactionMeta_CleanPartitionStatsInfo_Call {
	_c.Call.Return(run)
	return _c
}

// CompleteCompactionMutation provides a mock function with given fields: t, result
func (_m *MockCompactionMeta) CompleteCompactionMutation(t *datapb.CompactionTask, result *datapb.CompactionPlanResult) ([]*SegmentInfo, *segMetricMutation, error) {
	ret := _m.Called(t, result)

	if len(ret) == 0 {
		panic("no return value specified for CompleteCompactionMutation")
	}

	var r0 []*SegmentInfo
	var r1 *segMetricMutation
	var r2 error
	if rf, ok := ret.Get(0).(func(*datapb.CompactionTask, *datapb.CompactionPlanResult) ([]*SegmentInfo, *segMetricMutation, error)); ok {
		return rf(t, result)
	}
	if rf, ok := ret.Get(0).(func(*datapb.CompactionTask, *datapb.CompactionPlanResult) []*SegmentInfo); ok {
		r0 = rf(t, result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*SegmentInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(*datapb.CompactionTask, *datapb.CompactionPlanResult) *segMetricMutation); ok {
		r1 = rf(t, result)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*segMetricMutation)
		}
	}

	if rf, ok := ret.Get(2).(func(*datapb.CompactionTask, *datapb.CompactionPlanResult) error); ok {
		r2 = rf(t, result)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCompactionMeta_CompleteCompactionMutation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteCompactionMutation'
type MockCompactionMeta_CompleteCompactionMutation_Call struct {
	*mock.Call
}

// CompleteCompactionMutation is a helper method to define mock.On call
//   - t *datapb.CompactionTask
//   - result *datapb.CompactionPlanResult
func (_e *MockCompactionMeta_Expecter) CompleteCompactionMutation(t interface{}, result interface{}) *MockCompactionMeta_CompleteCompactionMutation_Call {
	return &MockCompactionMeta_CompleteCompactionMutation_Call{Call: _e.mock.On("CompleteCompactionMutation", t, result)}
}

func (_c *MockCompactionMeta_CompleteCompactionMutation_Call) Run(run func(t *datapb.CompactionTask, result *datapb.CompactionPlanResult)) *MockCompactionMeta_CompleteCompactionMutation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*datapb.CompactionTask), args[1].(*datapb.CompactionPlanResult))
	})
	return _c
}

func (_c *MockCompactionMeta_CompleteCompactionMutation_Call) Return(_a0 []*SegmentInfo, _a1 *segMetricMutation, _a2 error) *MockCompactionMeta_CompleteCompactionMutation_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockCompactionMeta_CompleteCompactionMutation_Call) RunAndReturn(run func(*datapb.CompactionTask, *datapb.CompactionPlanResult) ([]*SegmentInfo, *segMetricMutation, error)) *MockCompactionMeta_CompleteCompactionMutation_Call {
	_c.Call.Return(run)
	return _c
}

// DropCompactionTask provides a mock function with given fields: task
func (_m *MockCompactionMeta) DropCompactionTask(task *datapb.CompactionTask) error {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for DropCompactionTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*datapb.CompactionTask) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCompactionMeta_DropCompactionTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropCompactionTask'
type MockCompactionMeta_DropCompactionTask_Call struct {
	*mock.Call
}

// DropCompactionTask is a helper method to define mock.On call
//   - task *datapb.CompactionTask
func (_e *MockCompactionMeta_Expecter) DropCompactionTask(task interface{}) *MockCompactionMeta_DropCompactionTask_Call {
	return &MockCompactionMeta_DropCompactionTask_Call{Call: _e.mock.On("DropCompactionTask", task)}
}

func (_c *MockCompactionMeta_DropCompactionTask_Call) Run(run func(task *datapb.CompactionTask)) *MockCompactionMeta_DropCompactionTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*datapb.CompactionTask))
	})
	return _c
}

func (_c *MockCompactionMeta_DropCompactionTask_Call) Return(_a0 error) *MockCompactionMeta_DropCompactionTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_DropCompactionTask_Call) RunAndReturn(run func(*datapb.CompactionTask) error) *MockCompactionMeta_DropCompactionTask_Call {
	_c.Call.Return(run)
	return _c
}

// GetAnalyzeMeta provides a mock function with given fields:
func (_m *MockCompactionMeta) GetAnalyzeMeta() *analyzeMeta {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAnalyzeMeta")
	}

	var r0 *analyzeMeta
	if rf, ok := ret.Get(0).(func() *analyzeMeta); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*analyzeMeta)
		}
	}

	return r0
}

// MockCompactionMeta_GetAnalyzeMeta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAnalyzeMeta'
type MockCompactionMeta_GetAnalyzeMeta_Call struct {
	*mock.Call
}

// GetAnalyzeMeta is a helper method to define mock.On call
func (_e *MockCompactionMeta_Expecter) GetAnalyzeMeta() *MockCompactionMeta_GetAnalyzeMeta_Call {
	return &MockCompactionMeta_GetAnalyzeMeta_Call{Call: _e.mock.On("GetAnalyzeMeta")}
}

func (_c *MockCompactionMeta_GetAnalyzeMeta_Call) Run(run func()) *MockCompactionMeta_GetAnalyzeMeta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionMeta_GetAnalyzeMeta_Call) Return(_a0 *analyzeMeta) *MockCompactionMeta_GetAnalyzeMeta_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_GetAnalyzeMeta_Call) RunAndReturn(run func() *analyzeMeta) *MockCompactionMeta_GetAnalyzeMeta_Call {
	_c.Call.Return(run)
	return _c
}

// GetCompactionTaskMeta provides a mock function with given fields:
func (_m *MockCompactionMeta) GetCompactionTaskMeta() *compactionTaskMeta {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCompactionTaskMeta")
	}

	var r0 *compactionTaskMeta
	if rf, ok := ret.Get(0).(func() *compactionTaskMeta); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compactionTaskMeta)
		}
	}

	return r0
}

// MockCompactionMeta_GetCompactionTaskMeta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCompactionTaskMeta'
type MockCompactionMeta_GetCompactionTaskMeta_Call struct {
	*mock.Call
}

// GetCompactionTaskMeta is a helper method to define mock.On call
func (_e *MockCompactionMeta_Expecter) GetCompactionTaskMeta() *MockCompactionMeta_GetCompactionTaskMeta_Call {
	return &MockCompactionMeta_GetCompactionTaskMeta_Call{Call: _e.mock.On("GetCompactionTaskMeta")}
}

func (_c *MockCompactionMeta_GetCompactionTaskMeta_Call) Run(run func()) *MockCompactionMeta_GetCompactionTaskMeta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionMeta_GetCompactionTaskMeta_Call) Return(_a0 *compactionTaskMeta) *MockCompactionMeta_GetCompactionTaskMeta_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_GetCompactionTaskMeta_Call) RunAndReturn(run func() *compactionTaskMeta) *MockCompactionMeta_GetCompactionTaskMeta_Call {
	_c.Call.Return(run)
	return _c
}

// GetCompactionTasks provides a mock function with given fields:
func (_m *MockCompactionMeta) GetCompactionTasks() map[int64][]*datapb.CompactionTask {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCompactionTasks")
	}

	var r0 map[int64][]*datapb.CompactionTask
	if rf, ok := ret.Get(0).(func() map[int64][]*datapb.CompactionTask); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64][]*datapb.CompactionTask)
		}
	}

	return r0
}

// MockCompactionMeta_GetCompactionTasks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCompactionTasks'
type MockCompactionMeta_GetCompactionTasks_Call struct {
	*mock.Call
}

// GetCompactionTasks is a helper method to define mock.On call
func (_e *MockCompactionMeta_Expecter) GetCompactionTasks() *MockCompactionMeta_GetCompactionTasks_Call {
	return &MockCompactionMeta_GetCompactionTasks_Call{Call: _e.mock.On("GetCompactionTasks")}
}

func (_c *MockCompactionMeta_GetCompactionTasks_Call) Run(run func()) *MockCompactionMeta_GetCompactionTasks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionMeta_GetCompactionTasks_Call) Return(_a0 map[int64][]*datapb.CompactionTask) *MockCompactionMeta_GetCompactionTasks_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_GetCompactionTasks_Call) RunAndReturn(run func() map[int64][]*datapb.CompactionTask) *MockCompactionMeta_GetCompactionTasks_Call {
	_c.Call.Return(run)
	return _c
}

// GetCompactionTasksByTriggerID provides a mock function with given fields: triggerID
func (_m *MockCompactionMeta) GetCompactionTasksByTriggerID(triggerID int64) []*datapb.CompactionTask {
	ret := _m.Called(triggerID)

	if len(ret) == 0 {
		panic("no return value specified for GetCompactionTasksByTriggerID")
	}

	var r0 []*datapb.CompactionTask
	if rf, ok := ret.Get(0).(func(int64) []*datapb.CompactionTask); ok {
		r0 = rf(triggerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datapb.CompactionTask)
		}
	}

	return r0
}

// MockCompactionMeta_GetCompactionTasksByTriggerID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCompactionTasksByTriggerID'
type MockCompactionMeta_GetCompactionTasksByTriggerID_Call struct {
	*mock.Call
}

// GetCompactionTasksByTriggerID is a helper method to define mock.On call
//   - triggerID int64
func (_e *MockCompactionMeta_Expecter) GetCompactionTasksByTriggerID(triggerID interface{}) *MockCompactionMeta_GetCompactionTasksByTriggerID_Call {
	return &MockCompactionMeta_GetCompactionTasksByTriggerID_Call{Call: _e.mock.On("GetCompactionTasksByTriggerID", triggerID)}
}

func (_c *MockCompactionMeta_GetCompactionTasksByTriggerID_Call) Run(run func(triggerID int64)) *MockCompactionMeta_GetCompactionTasksByTriggerID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCompactionMeta_GetCompactionTasksByTriggerID_Call) Return(_a0 []*datapb.CompactionTask) *MockCompactionMeta_GetCompactionTasksByTriggerID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_GetCompactionTasksByTriggerID_Call) RunAndReturn(run func(int64) []*datapb.CompactionTask) *MockCompactionMeta_GetCompactionTasksByTriggerID_Call {
	_c.Call.Return(run)
	return _c
}

// GetHealthySegment provides a mock function with given fields: segID
func (_m *MockCompactionMeta) GetHealthySegment(segID int64) *SegmentInfo {
	ret := _m.Called(segID)

	if len(ret) == 0 {
		panic("no return value specified for GetHealthySegment")
	}

	var r0 *SegmentInfo
	if rf, ok := ret.Get(0).(func(int64) *SegmentInfo); ok {
		r0 = rf(segID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SegmentInfo)
		}
	}

	return r0
}

// MockCompactionMeta_GetHealthySegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHealthySegment'
type MockCompactionMeta_GetHealthySegment_Call struct {
	*mock.Call
}

// GetHealthySegment is a helper method to define mock.On call
//   - segID int64
func (_e *MockCompactionMeta_Expecter) GetHealthySegment(segID interface{}) *MockCompactionMeta_GetHealthySegment_Call {
	return &MockCompactionMeta_GetHealthySegment_Call{Call: _e.mock.On("GetHealthySegment", segID)}
}

func (_c *MockCompactionMeta_GetHealthySegment_Call) Run(run func(segID int64)) *MockCompactionMeta_GetHealthySegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCompactionMeta_GetHealthySegment_Call) Return(_a0 *SegmentInfo) *MockCompactionMeta_GetHealthySegment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_GetHealthySegment_Call) RunAndReturn(run func(int64) *SegmentInfo) *MockCompactionMeta_GetHealthySegment_Call {
	_c.Call.Return(run)
	return _c
}

// GetIndexMeta provides a mock function with given fields:
func (_m *MockCompactionMeta) GetIndexMeta() *indexMeta {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetIndexMeta")
	}

	var r0 *indexMeta
	if rf, ok := ret.Get(0).(func() *indexMeta); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*indexMeta)
		}
	}

	return r0
}

// MockCompactionMeta_GetIndexMeta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndexMeta'
type MockCompactionMeta_GetIndexMeta_Call struct {
	*mock.Call
}

// GetIndexMeta is a helper method to define mock.On call
func (_e *MockCompactionMeta_Expecter) GetIndexMeta() *MockCompactionMeta_GetIndexMeta_Call {
	return &MockCompactionMeta_GetIndexMeta_Call{Call: _e.mock.On("GetIndexMeta")}
}

func (_c *MockCompactionMeta_GetIndexMeta_Call) Run(run func()) *MockCompactionMeta_GetIndexMeta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionMeta_GetIndexMeta_Call) Return(_a0 *indexMeta) *MockCompactionMeta_GetIndexMeta_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_GetIndexMeta_Call) RunAndReturn(run func() *indexMeta) *MockCompactionMeta_GetIndexMeta_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartitionStatsMeta provides a mock function with given fields:
func (_m *MockCompactionMeta) GetPartitionStatsMeta() *partitionStatsMeta {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPartitionStatsMeta")
	}

	var r0 *partitionStatsMeta
	if rf, ok := ret.Get(0).(func() *partitionStatsMeta); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*partitionStatsMeta)
		}
	}

	return r0
}

// MockCompactionMeta_GetPartitionStatsMeta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitionStatsMeta'
type MockCompactionMeta_GetPartitionStatsMeta_Call struct {
	*mock.Call
}

// GetPartitionStatsMeta is a helper method to define mock.On call
func (_e *MockCompactionMeta_Expecter) GetPartitionStatsMeta() *MockCompactionMeta_GetPartitionStatsMeta_Call {
	return &MockCompactionMeta_GetPartitionStatsMeta_Call{Call: _e.mock.On("GetPartitionStatsMeta")}
}

func (_c *MockCompactionMeta_GetPartitionStatsMeta_Call) Run(run func()) *MockCompactionMeta_GetPartitionStatsMeta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionMeta_GetPartitionStatsMeta_Call) Return(_a0 *partitionStatsMeta) *MockCompactionMeta_GetPartitionStatsMeta_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_GetPartitionStatsMeta_Call) RunAndReturn(run func() *partitionStatsMeta) *MockCompactionMeta_GetPartitionStatsMeta_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegment provides a mock function with given fields: segID
func (_m *MockCompactionMeta) GetSegment(segID int64) *SegmentInfo {
	ret := _m.Called(segID)

	if len(ret) == 0 {
		panic("no return value specified for GetSegment")
	}

	var r0 *SegmentInfo
	if rf, ok := ret.Get(0).(func(int64) *SegmentInfo); ok {
		r0 = rf(segID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SegmentInfo)
		}
	}

	return r0
}

// MockCompactionMeta_GetSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegment'
type MockCompactionMeta_GetSegment_Call struct {
	*mock.Call
}

// GetSegment is a helper method to define mock.On call
//   - segID int64
func (_e *MockCompactionMeta_Expecter) GetSegment(segID interface{}) *MockCompactionMeta_GetSegment_Call {
	return &MockCompactionMeta_GetSegment_Call{Call: _e.mock.On("GetSegment", segID)}
}

func (_c *MockCompactionMeta_GetSegment_Call) Run(run func(segID int64)) *MockCompactionMeta_GetSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCompactionMeta_GetSegment_Call) Return(_a0 *SegmentInfo) *MockCompactionMeta_GetSegment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_GetSegment_Call) RunAndReturn(run func(int64) *SegmentInfo) *MockCompactionMeta_GetSegment_Call {
	_c.Call.Return(run)
	return _c
}

// SaveCompactionTask provides a mock function with given fields: task
func (_m *MockCompactionMeta) SaveCompactionTask(task *datapb.CompactionTask) error {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for SaveCompactionTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*datapb.CompactionTask) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCompactionMeta_SaveCompactionTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveCompactionTask'
type MockCompactionMeta_SaveCompactionTask_Call struct {
	*mock.Call
}

// SaveCompactionTask is a helper method to define mock.On call
//   - task *datapb.CompactionTask
func (_e *MockCompactionMeta_Expecter) SaveCompactionTask(task interface{}) *MockCompactionMeta_SaveCompactionTask_Call {
	return &MockCompactionMeta_SaveCompactionTask_Call{Call: _e.mock.On("SaveCompactionTask", task)}
}

func (_c *MockCompactionMeta_SaveCompactionTask_Call) Run(run func(task *datapb.CompactionTask)) *MockCompactionMeta_SaveCompactionTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*datapb.CompactionTask))
	})
	return _c
}

func (_c *MockCompactionMeta_SaveCompactionTask_Call) Return(_a0 error) *MockCompactionMeta_SaveCompactionTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_SaveCompactionTask_Call) RunAndReturn(run func(*datapb.CompactionTask) error) *MockCompactionMeta_SaveCompactionTask_Call {
	_c.Call.Return(run)
	return _c
}

// SelectSegments provides a mock function with given fields: filters
func (_m *MockCompactionMeta) SelectSegments(filters ...SegmentFilter) []*SegmentInfo {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SelectSegments")
	}

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

// MockCompactionMeta_SelectSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectSegments'
type MockCompactionMeta_SelectSegments_Call struct {
	*mock.Call
}

// SelectSegments is a helper method to define mock.On call
//   - filters ...SegmentFilter
func (_e *MockCompactionMeta_Expecter) SelectSegments(filters ...interface{}) *MockCompactionMeta_SelectSegments_Call {
	return &MockCompactionMeta_SelectSegments_Call{Call: _e.mock.On("SelectSegments",
		append([]interface{}{}, filters...)...)}
}

func (_c *MockCompactionMeta_SelectSegments_Call) Run(run func(filters ...SegmentFilter)) *MockCompactionMeta_SelectSegments_Call {
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

func (_c *MockCompactionMeta_SelectSegments_Call) Return(_a0 []*SegmentInfo) *MockCompactionMeta_SelectSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_SelectSegments_Call) RunAndReturn(run func(...SegmentFilter) []*SegmentInfo) *MockCompactionMeta_SelectSegments_Call {
	_c.Call.Return(run)
	return _c
}

// SetSegmentsCompacting provides a mock function with given fields: segmentID, compacting
func (_m *MockCompactionMeta) SetSegmentsCompacting(segmentID []int64, compacting bool) {
	_m.Called(segmentID, compacting)
}

// MockCompactionMeta_SetSegmentsCompacting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSegmentsCompacting'
type MockCompactionMeta_SetSegmentsCompacting_Call struct {
	*mock.Call
}

// SetSegmentsCompacting is a helper method to define mock.On call
//   - segmentID []int64
//   - compacting bool
func (_e *MockCompactionMeta_Expecter) SetSegmentsCompacting(segmentID interface{}, compacting interface{}) *MockCompactionMeta_SetSegmentsCompacting_Call {
	return &MockCompactionMeta_SetSegmentsCompacting_Call{Call: _e.mock.On("SetSegmentsCompacting", segmentID, compacting)}
}

func (_c *MockCompactionMeta_SetSegmentsCompacting_Call) Run(run func(segmentID []int64, compacting bool)) *MockCompactionMeta_SetSegmentsCompacting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]int64), args[1].(bool))
	})
	return _c
}

func (_c *MockCompactionMeta_SetSegmentsCompacting_Call) Return() *MockCompactionMeta_SetSegmentsCompacting_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCompactionMeta_SetSegmentsCompacting_Call) RunAndReturn(run func([]int64, bool)) *MockCompactionMeta_SetSegmentsCompacting_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSegmentsInfo provides a mock function with given fields: operators
func (_m *MockCompactionMeta) UpdateSegmentsInfo(operators ...UpdateOperator) error {
	_va := make([]interface{}, len(operators))
	for _i := range operators {
		_va[_i] = operators[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSegmentsInfo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...UpdateOperator) error); ok {
		r0 = rf(operators...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCompactionMeta_UpdateSegmentsInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSegmentsInfo'
type MockCompactionMeta_UpdateSegmentsInfo_Call struct {
	*mock.Call
}

// UpdateSegmentsInfo is a helper method to define mock.On call
//   - operators ...UpdateOperator
func (_e *MockCompactionMeta_Expecter) UpdateSegmentsInfo(operators ...interface{}) *MockCompactionMeta_UpdateSegmentsInfo_Call {
	return &MockCompactionMeta_UpdateSegmentsInfo_Call{Call: _e.mock.On("UpdateSegmentsInfo",
		append([]interface{}{}, operators...)...)}
}

func (_c *MockCompactionMeta_UpdateSegmentsInfo_Call) Run(run func(operators ...UpdateOperator)) *MockCompactionMeta_UpdateSegmentsInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]UpdateOperator, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(UpdateOperator)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockCompactionMeta_UpdateSegmentsInfo_Call) Return(_a0 error) *MockCompactionMeta_UpdateSegmentsInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompactionMeta_UpdateSegmentsInfo_Call) RunAndReturn(run func(...UpdateOperator) error) *MockCompactionMeta_UpdateSegmentsInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCompactionMeta creates a new instance of MockCompactionMeta. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCompactionMeta(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCompactionMeta {
	mock := &MockCompactionMeta{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
