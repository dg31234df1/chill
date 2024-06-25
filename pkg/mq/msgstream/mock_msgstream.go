// Code generated by mockery v2.32.4. DO NOT EDIT.

package msgstream

import (
	"context"

	"github.com/milvus-io/milvus/pkg/mq/common"
	"github.com/stretchr/testify/mock"

	"github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
)

// MockMsgStream is an autogenerated mock type for the MsgStream type
type MockMsgStream struct {
	mock.Mock
}

type MockMsgStream_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMsgStream) EXPECT() *MockMsgStream_Expecter {
	return &MockMsgStream_Expecter{mock: &_m.Mock}
}

// AsConsumer provides a mock function with given fields: ctx, channels, subName, position
func (_m *MockMsgStream) AsConsumer(ctx context.Context, channels []string, subName string, position common.SubscriptionInitialPosition) error {
	ret := _m.Called(ctx, channels, subName, position)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, common.SubscriptionInitialPosition) error); ok {
		r0 = rf(ctx, channels, subName, position)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMsgStream_AsConsumer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AsConsumer'
type MockMsgStream_AsConsumer_Call struct {
	*mock.Call
}

// AsConsumer is a helper method to define mock.On call
//   - ctx context.Context
//   - channels []string
//   - subName string
//   - position mqwrapper.SubscriptionInitialPosition
func (_e *MockMsgStream_Expecter) AsConsumer(ctx interface{}, channels interface{}, subName interface{}, position interface{}) *MockMsgStream_AsConsumer_Call {
	return &MockMsgStream_AsConsumer_Call{Call: _e.mock.On("AsConsumer", ctx, channels, subName, position)}
}

func (_c *MockMsgStream_AsConsumer_Call) Run(run func(ctx context.Context, channels []string, subName string, position common.SubscriptionInitialPosition)) *MockMsgStream_AsConsumer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(string), args[3].(common.SubscriptionInitialPosition))
	})
	return _c
}

func (_c *MockMsgStream_AsConsumer_Call) Return(_a0 error) *MockMsgStream_AsConsumer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMsgStream_AsConsumer_Call) RunAndReturn(run func(context.Context, []string, string, common.SubscriptionInitialPosition) error) *MockMsgStream_AsConsumer_Call {
	_c.Call.Return(run)
	return _c
}

// AsProducer provides a mock function with given fields: channels
func (_m *MockMsgStream) AsProducer(channels []string) {
	_m.Called(channels)
}

// MockMsgStream_AsProducer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AsProducer'
type MockMsgStream_AsProducer_Call struct {
	*mock.Call
}

// AsProducer is a helper method to define mock.On call
//   - channels []string
func (_e *MockMsgStream_Expecter) AsProducer(channels interface{}) *MockMsgStream_AsProducer_Call {
	return &MockMsgStream_AsProducer_Call{Call: _e.mock.On("AsProducer", channels)}
}

func (_c *MockMsgStream_AsProducer_Call) Run(run func(channels []string)) *MockMsgStream_AsProducer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MockMsgStream_AsProducer_Call) Return() *MockMsgStream_AsProducer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMsgStream_AsProducer_Call) RunAndReturn(run func([]string)) *MockMsgStream_AsProducer_Call {
	_c.Call.Return(run)
	return _c
}

// Broadcast provides a mock function with given fields: _a0
func (_m *MockMsgStream) Broadcast(_a0 *MsgPack) (map[string][]common.MessageID, error) {
	ret := _m.Called(_a0)

	var r0 map[string][]common.MessageID
	var r1 error
	if rf, ok := ret.Get(0).(func(*MsgPack) (map[string][]common.MessageID, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*MsgPack) map[string][]common.MessageID); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]common.MessageID)
		}
	}

	if rf, ok := ret.Get(1).(func(*MsgPack) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMsgStream_Broadcast_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Broadcast'
type MockMsgStream_Broadcast_Call struct {
	*mock.Call
}

// Broadcast is a helper method to define mock.On call
//   - _a0 *MsgPack
func (_e *MockMsgStream_Expecter) Broadcast(_a0 interface{}) *MockMsgStream_Broadcast_Call {
	return &MockMsgStream_Broadcast_Call{Call: _e.mock.On("Broadcast", _a0)}
}

func (_c *MockMsgStream_Broadcast_Call) Run(run func(_a0 *MsgPack)) *MockMsgStream_Broadcast_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*MsgPack))
	})
	return _c
}

func (_c *MockMsgStream_Broadcast_Call) Return(_a0 map[string][]common.MessageID, _a1 error) *MockMsgStream_Broadcast_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMsgStream_Broadcast_Call) RunAndReturn(run func(*MsgPack) (map[string][]common.MessageID, error)) *MockMsgStream_Broadcast_Call {
	_c.Call.Return(run)
	return _c
}

// Chan provides a mock function with given fields:
func (_m *MockMsgStream) Chan() <-chan *MsgPack {
	ret := _m.Called()

	var r0 <-chan *MsgPack
	if rf, ok := ret.Get(0).(func() <-chan *MsgPack); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *MsgPack)
		}
	}

	return r0
}

// MockMsgStream_Chan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chan'
type MockMsgStream_Chan_Call struct {
	*mock.Call
}

// Chan is a helper method to define mock.On call
func (_e *MockMsgStream_Expecter) Chan() *MockMsgStream_Chan_Call {
	return &MockMsgStream_Chan_Call{Call: _e.mock.On("Chan")}
}

func (_c *MockMsgStream_Chan_Call) Run(run func()) *MockMsgStream_Chan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMsgStream_Chan_Call) Return(_a0 <-chan *MsgPack) *MockMsgStream_Chan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMsgStream_Chan_Call) RunAndReturn(run func() <-chan *MsgPack) *MockMsgStream_Chan_Call {
	_c.Call.Return(run)
	return _c
}

// CheckTopicValid provides a mock function with given fields: channel
func (_m *MockMsgStream) CheckTopicValid(channel string) error {
	ret := _m.Called(channel)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(channel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMsgStream_CheckTopicValid_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckTopicValid'
type MockMsgStream_CheckTopicValid_Call struct {
	*mock.Call
}

// CheckTopicValid is a helper method to define mock.On call
//   - channel string
func (_e *MockMsgStream_Expecter) CheckTopicValid(channel interface{}) *MockMsgStream_CheckTopicValid_Call {
	return &MockMsgStream_CheckTopicValid_Call{Call: _e.mock.On("CheckTopicValid", channel)}
}

func (_c *MockMsgStream_CheckTopicValid_Call) Run(run func(channel string)) *MockMsgStream_CheckTopicValid_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockMsgStream_CheckTopicValid_Call) Return(_a0 error) *MockMsgStream_CheckTopicValid_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMsgStream_CheckTopicValid_Call) RunAndReturn(run func(string) error) *MockMsgStream_CheckTopicValid_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockMsgStream) Close() {
	_m.Called()
}

// MockMsgStream_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockMsgStream_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockMsgStream_Expecter) Close() *MockMsgStream_Close_Call {
	return &MockMsgStream_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockMsgStream_Close_Call) Run(run func()) *MockMsgStream_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMsgStream_Close_Call) Return() *MockMsgStream_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMsgStream_Close_Call) RunAndReturn(run func()) *MockMsgStream_Close_Call {
	_c.Call.Return(run)
	return _c
}

// EnableProduce provides a mock function with given fields: can
func (_m *MockMsgStream) EnableProduce(can bool) {
	_m.Called(can)
}

// MockMsgStream_EnableProduce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnableProduce'
type MockMsgStream_EnableProduce_Call struct {
	*mock.Call
}

// EnableProduce is a helper method to define mock.On call
//   - can bool
func (_e *MockMsgStream_Expecter) EnableProduce(can interface{}) *MockMsgStream_EnableProduce_Call {
	return &MockMsgStream_EnableProduce_Call{Call: _e.mock.On("EnableProduce", can)}
}

func (_c *MockMsgStream_EnableProduce_Call) Run(run func(can bool)) *MockMsgStream_EnableProduce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockMsgStream_EnableProduce_Call) Return() *MockMsgStream_EnableProduce_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMsgStream_EnableProduce_Call) RunAndReturn(run func(bool)) *MockMsgStream_EnableProduce_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestMsgID provides a mock function with given fields: channel
func (_m *MockMsgStream) GetLatestMsgID(channel string) (common.MessageID, error) {
	ret := _m.Called(channel)

	var r0 common.MessageID
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (common.MessageID, error)); ok {
		return rf(channel)
	}
	if rf, ok := ret.Get(0).(func(string) common.MessageID); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.MessageID)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMsgStream_GetLatestMsgID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestMsgID'
type MockMsgStream_GetLatestMsgID_Call struct {
	*mock.Call
}

// GetLatestMsgID is a helper method to define mock.On call
//   - channel string
func (_e *MockMsgStream_Expecter) GetLatestMsgID(channel interface{}) *MockMsgStream_GetLatestMsgID_Call {
	return &MockMsgStream_GetLatestMsgID_Call{Call: _e.mock.On("GetLatestMsgID", channel)}
}

func (_c *MockMsgStream_GetLatestMsgID_Call) Run(run func(channel string)) *MockMsgStream_GetLatestMsgID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockMsgStream_GetLatestMsgID_Call) Return(_a0 common.MessageID, _a1 error) *MockMsgStream_GetLatestMsgID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMsgStream_GetLatestMsgID_Call) RunAndReturn(run func(string) (common.MessageID, error)) *MockMsgStream_GetLatestMsgID_Call {
	_c.Call.Return(run)
	return _c
}

// GetProduceChannels provides a mock function with given fields:
func (_m *MockMsgStream) GetProduceChannels() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockMsgStream_GetProduceChannels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProduceChannels'
type MockMsgStream_GetProduceChannels_Call struct {
	*mock.Call
}

// GetProduceChannels is a helper method to define mock.On call
func (_e *MockMsgStream_Expecter) GetProduceChannels() *MockMsgStream_GetProduceChannels_Call {
	return &MockMsgStream_GetProduceChannels_Call{Call: _e.mock.On("GetProduceChannels")}
}

func (_c *MockMsgStream_GetProduceChannels_Call) Run(run func()) *MockMsgStream_GetProduceChannels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMsgStream_GetProduceChannels_Call) Return(_a0 []string) *MockMsgStream_GetProduceChannels_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMsgStream_GetProduceChannels_Call) RunAndReturn(run func() []string) *MockMsgStream_GetProduceChannels_Call {
	_c.Call.Return(run)
	return _c
}

// Produce provides a mock function with given fields: _a0
func (_m *MockMsgStream) Produce(_a0 *MsgPack) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*MsgPack) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMsgStream_Produce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Produce'
type MockMsgStream_Produce_Call struct {
	*mock.Call
}

// Produce is a helper method to define mock.On call
//   - _a0 *MsgPack
func (_e *MockMsgStream_Expecter) Produce(_a0 interface{}) *MockMsgStream_Produce_Call {
	return &MockMsgStream_Produce_Call{Call: _e.mock.On("Produce", _a0)}
}

func (_c *MockMsgStream_Produce_Call) Run(run func(_a0 *MsgPack)) *MockMsgStream_Produce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*MsgPack))
	})
	return _c
}

func (_c *MockMsgStream_Produce_Call) Return(_a0 error) *MockMsgStream_Produce_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMsgStream_Produce_Call) RunAndReturn(run func(*MsgPack) error) *MockMsgStream_Produce_Call {
	_c.Call.Return(run)
	return _c
}

// Seek provides a mock function with given fields: ctx, msgPositions, includeCurrentMsg
func (_m *MockMsgStream) Seek(ctx context.Context, msgPositions []*msgpb.MsgPosition, includeCurrentMsg bool) error {
	ret := _m.Called(ctx, msgPositions, includeCurrentMsg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*msgpb.MsgPosition, bool) error); ok {
		r0 = rf(ctx, msgPositions, includeCurrentMsg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMsgStream_Seek_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Seek'
type MockMsgStream_Seek_Call struct {
	*mock.Call
}

// Seek is a helper method to define mock.On call
//   - ctx context.Context
//   - msgPositions []*msgpb.MsgPosition
//   - includeCurrentMsg bool
func (_e *MockMsgStream_Expecter) Seek(ctx interface{}, msgPositions interface{}, includeCurrentMsg interface{}) *MockMsgStream_Seek_Call {
	return &MockMsgStream_Seek_Call{Call: _e.mock.On("Seek", ctx, msgPositions, includeCurrentMsg)}
}

func (_c *MockMsgStream_Seek_Call) Run(run func(ctx context.Context, msgPositions []*msgpb.MsgPosition, includeCurrentMsg bool)) *MockMsgStream_Seek_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*msgpb.MsgPosition), args[2].(bool))
	})
	return _c
}

func (_c *MockMsgStream_Seek_Call) Return(_a0 error) *MockMsgStream_Seek_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMsgStream_Seek_Call) RunAndReturn(run func(context.Context, []*msgpb.MsgPosition, bool) error) *MockMsgStream_Seek_Call {
	_c.Call.Return(run)
	return _c
}

// SetRepackFunc provides a mock function with given fields: repackFunc
func (_m *MockMsgStream) SetRepackFunc(repackFunc RepackFunc) {
	_m.Called(repackFunc)
}

// MockMsgStream_SetRepackFunc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetRepackFunc'
type MockMsgStream_SetRepackFunc_Call struct {
	*mock.Call
}

// SetRepackFunc is a helper method to define mock.On call
//   - repackFunc RepackFunc
func (_e *MockMsgStream_Expecter) SetRepackFunc(repackFunc interface{}) *MockMsgStream_SetRepackFunc_Call {
	return &MockMsgStream_SetRepackFunc_Call{Call: _e.mock.On("SetRepackFunc", repackFunc)}
}

func (_c *MockMsgStream_SetRepackFunc_Call) Run(run func(repackFunc RepackFunc)) *MockMsgStream_SetRepackFunc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(RepackFunc))
	})
	return _c
}

func (_c *MockMsgStream_SetRepackFunc_Call) Return() *MockMsgStream_SetRepackFunc_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMsgStream_SetRepackFunc_Call) RunAndReturn(run func(RepackFunc)) *MockMsgStream_SetRepackFunc_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMsgStream creates a new instance of MockMsgStream. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMsgStream(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMsgStream {
	mock := &MockMsgStream{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
