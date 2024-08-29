// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_inspector

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	stats "github.com/milvus-io/milvus/internal/streamingnode/server/wal/interceptors/segment/stats"

	types "github.com/milvus-io/milvus/pkg/streaming/util/types"
)

// MockSealOperator is an autogenerated mock type for the SealOperator type
type MockSealOperator struct {
	mock.Mock
}

type MockSealOperator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSealOperator) EXPECT() *MockSealOperator_Expecter {
	return &MockSealOperator_Expecter{mock: &_m.Mock}
}

// Channel provides a mock function with given fields:
func (_m *MockSealOperator) Channel() types.PChannelInfo {
	ret := _m.Called()

	var r0 types.PChannelInfo
	if rf, ok := ret.Get(0).(func() types.PChannelInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.PChannelInfo)
	}

	return r0
}

// MockSealOperator_Channel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Channel'
type MockSealOperator_Channel_Call struct {
	*mock.Call
}

// Channel is a helper method to define mock.On call
func (_e *MockSealOperator_Expecter) Channel() *MockSealOperator_Channel_Call {
	return &MockSealOperator_Channel_Call{Call: _e.mock.On("Channel")}
}

func (_c *MockSealOperator_Channel_Call) Run(run func()) *MockSealOperator_Channel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSealOperator_Channel_Call) Return(_a0 types.PChannelInfo) *MockSealOperator_Channel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSealOperator_Channel_Call) RunAndReturn(run func() types.PChannelInfo) *MockSealOperator_Channel_Call {
	_c.Call.Return(run)
	return _c
}

// IsNoWaitSeal provides a mock function with given fields:
func (_m *MockSealOperator) IsNoWaitSeal() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSealOperator_IsNoWaitSeal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNoWaitSeal'
type MockSealOperator_IsNoWaitSeal_Call struct {
	*mock.Call
}

// IsNoWaitSeal is a helper method to define mock.On call
func (_e *MockSealOperator_Expecter) IsNoWaitSeal() *MockSealOperator_IsNoWaitSeal_Call {
	return &MockSealOperator_IsNoWaitSeal_Call{Call: _e.mock.On("IsNoWaitSeal")}
}

func (_c *MockSealOperator_IsNoWaitSeal_Call) Run(run func()) *MockSealOperator_IsNoWaitSeal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSealOperator_IsNoWaitSeal_Call) Return(_a0 bool) *MockSealOperator_IsNoWaitSeal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSealOperator_IsNoWaitSeal_Call) RunAndReturn(run func() bool) *MockSealOperator_IsNoWaitSeal_Call {
	_c.Call.Return(run)
	return _c
}

// MustSealSegments provides a mock function with given fields: ctx, infos
func (_m *MockSealOperator) MustSealSegments(ctx context.Context, infos ...stats.SegmentBelongs) {
	_va := make([]interface{}, len(infos))
	for _i := range infos {
		_va[_i] = infos[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockSealOperator_MustSealSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MustSealSegments'
type MockSealOperator_MustSealSegments_Call struct {
	*mock.Call
}

// MustSealSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - infos ...stats.SegmentBelongs
func (_e *MockSealOperator_Expecter) MustSealSegments(ctx interface{}, infos ...interface{}) *MockSealOperator_MustSealSegments_Call {
	return &MockSealOperator_MustSealSegments_Call{Call: _e.mock.On("MustSealSegments",
		append([]interface{}{ctx}, infos...)...)}
}

func (_c *MockSealOperator_MustSealSegments_Call) Run(run func(ctx context.Context, infos ...stats.SegmentBelongs)) *MockSealOperator_MustSealSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]stats.SegmentBelongs, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(stats.SegmentBelongs)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockSealOperator_MustSealSegments_Call) Return() *MockSealOperator_MustSealSegments_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSealOperator_MustSealSegments_Call) RunAndReturn(run func(context.Context, ...stats.SegmentBelongs)) *MockSealOperator_MustSealSegments_Call {
	_c.Call.Return(run)
	return _c
}

// TryToSealSegments provides a mock function with given fields: ctx, infos
func (_m *MockSealOperator) TryToSealSegments(ctx context.Context, infos ...stats.SegmentBelongs) {
	_va := make([]interface{}, len(infos))
	for _i := range infos {
		_va[_i] = infos[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockSealOperator_TryToSealSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TryToSealSegments'
type MockSealOperator_TryToSealSegments_Call struct {
	*mock.Call
}

// TryToSealSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - infos ...stats.SegmentBelongs
func (_e *MockSealOperator_Expecter) TryToSealSegments(ctx interface{}, infos ...interface{}) *MockSealOperator_TryToSealSegments_Call {
	return &MockSealOperator_TryToSealSegments_Call{Call: _e.mock.On("TryToSealSegments",
		append([]interface{}{ctx}, infos...)...)}
}

func (_c *MockSealOperator_TryToSealSegments_Call) Run(run func(ctx context.Context, infos ...stats.SegmentBelongs)) *MockSealOperator_TryToSealSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]stats.SegmentBelongs, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(stats.SegmentBelongs)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockSealOperator_TryToSealSegments_Call) Return() *MockSealOperator_TryToSealSegments_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSealOperator_TryToSealSegments_Call) RunAndReturn(run func(context.Context, ...stats.SegmentBelongs)) *MockSealOperator_TryToSealSegments_Call {
	_c.Call.Return(run)
	return _c
}

// TryToSealWaitedSegment provides a mock function with given fields: ctx
func (_m *MockSealOperator) TryToSealWaitedSegment(ctx context.Context) {
	_m.Called(ctx)
}

// MockSealOperator_TryToSealWaitedSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TryToSealWaitedSegment'
type MockSealOperator_TryToSealWaitedSegment_Call struct {
	*mock.Call
}

// TryToSealWaitedSegment is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSealOperator_Expecter) TryToSealWaitedSegment(ctx interface{}) *MockSealOperator_TryToSealWaitedSegment_Call {
	return &MockSealOperator_TryToSealWaitedSegment_Call{Call: _e.mock.On("TryToSealWaitedSegment", ctx)}
}

func (_c *MockSealOperator_TryToSealWaitedSegment_Call) Run(run func(ctx context.Context)) *MockSealOperator_TryToSealWaitedSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSealOperator_TryToSealWaitedSegment_Call) Return() *MockSealOperator_TryToSealWaitedSegment_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSealOperator_TryToSealWaitedSegment_Call) RunAndReturn(run func(context.Context)) *MockSealOperator_TryToSealWaitedSegment_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSealOperator creates a new instance of MockSealOperator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSealOperator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSealOperator {
	mock := &MockSealOperator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
