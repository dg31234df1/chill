// Code generated by mockery v2.32.4. DO NOT EDIT.

package proxy

import (
	context "context"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"
	mock "github.com/stretchr/testify/mock"

	typeutil "github.com/milvus-io/milvus/pkg/util/typeutil"
)

// MockCache is an autogenerated mock type for the Cache type
type MockCache struct {
	mock.Mock
}

type MockCache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCache) EXPECT() *MockCache_Expecter {
	return &MockCache_Expecter{mock: &_m.Mock}
}

// DeprecateShardCache provides a mock function with given fields: database, collectionName
func (_m *MockCache) DeprecateShardCache(database string, collectionName string) {
	_m.Called(database, collectionName)
}

// MockCache_DeprecateShardCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeprecateShardCache'
type MockCache_DeprecateShardCache_Call struct {
	*mock.Call
}

// DeprecateShardCache is a helper method to define mock.On call
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) DeprecateShardCache(database interface{}, collectionName interface{}) *MockCache_DeprecateShardCache_Call {
	return &MockCache_DeprecateShardCache_Call{Call: _e.mock.On("DeprecateShardCache", database, collectionName)}
}

func (_c *MockCache_DeprecateShardCache_Call) Run(run func(database string, collectionName string)) *MockCache_DeprecateShardCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockCache_DeprecateShardCache_Call) Return() *MockCache_DeprecateShardCache_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_DeprecateShardCache_Call) RunAndReturn(run func(string, string)) *MockCache_DeprecateShardCache_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionID provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) GetCollectionID(ctx context.Context, database string, collectionName string) (int64, error) {
	ret := _m.Called(ctx, database, collectionName)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (int64, error)); ok {
		return rf(ctx, database, collectionName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, database, collectionName)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCollectionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionID'
type MockCache_GetCollectionID_Call struct {
	*mock.Call
}

// GetCollectionID is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetCollectionID(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_GetCollectionID_Call {
	return &MockCache_GetCollectionID_Call{Call: _e.mock.On("GetCollectionID", ctx, database, collectionName)}
}

func (_c *MockCache_GetCollectionID_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_GetCollectionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_GetCollectionID_Call) Return(_a0 int64, _a1 error) *MockCache_GetCollectionID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetCollectionID_Call) RunAndReturn(run func(context.Context, string, string) (int64, error)) *MockCache_GetCollectionID_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionInfo provides a mock function with given fields: ctx, database, collectionName, collectionID
func (_m *MockCache) GetCollectionInfo(ctx context.Context, database string, collectionName string, collectionID int64) (*collectionBasicInfo, error) {
	ret := _m.Called(ctx, database, collectionName, collectionID)

	var r0 *collectionBasicInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) (*collectionBasicInfo, error)); ok {
		return rf(ctx, database, collectionName, collectionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *collectionBasicInfo); ok {
		r0 = rf(ctx, database, collectionName, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collectionBasicInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, database, collectionName, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCollectionInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionInfo'
type MockCache_GetCollectionInfo_Call struct {
	*mock.Call
}

// GetCollectionInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
//   - collectionID int64
func (_e *MockCache_Expecter) GetCollectionInfo(ctx interface{}, database interface{}, collectionName interface{}, collectionID interface{}) *MockCache_GetCollectionInfo_Call {
	return &MockCache_GetCollectionInfo_Call{Call: _e.mock.On("GetCollectionInfo", ctx, database, collectionName, collectionID)}
}

func (_c *MockCache_GetCollectionInfo_Call) Run(run func(ctx context.Context, database string, collectionName string, collectionID int64)) *MockCache_GetCollectionInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(int64))
	})
	return _c
}

func (_c *MockCache_GetCollectionInfo_Call) Return(_a0 *collectionBasicInfo, _a1 error) *MockCache_GetCollectionInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetCollectionInfo_Call) RunAndReturn(run func(context.Context, string, string, int64) (*collectionBasicInfo, error)) *MockCache_GetCollectionInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionName provides a mock function with given fields: ctx, database, collectionID
func (_m *MockCache) GetCollectionName(ctx context.Context, database string, collectionID int64) (string, error) {
	ret := _m.Called(ctx, database, collectionID)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (string, error)); ok {
		return rf(ctx, database, collectionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) string); ok {
		r0 = rf(ctx, database, collectionID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, database, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCollectionName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionName'
type MockCache_GetCollectionName_Call struct {
	*mock.Call
}

// GetCollectionName is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionID int64
func (_e *MockCache_Expecter) GetCollectionName(ctx interface{}, database interface{}, collectionID interface{}) *MockCache_GetCollectionName_Call {
	return &MockCache_GetCollectionName_Call{Call: _e.mock.On("GetCollectionName", ctx, database, collectionID)}
}

func (_c *MockCache_GetCollectionName_Call) Run(run func(ctx context.Context, database string, collectionID int64)) *MockCache_GetCollectionName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockCache_GetCollectionName_Call) Return(_a0 string, _a1 error) *MockCache_GetCollectionName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetCollectionName_Call) RunAndReturn(run func(context.Context, string, int64) (string, error)) *MockCache_GetCollectionName_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionSchema provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) GetCollectionSchema(ctx context.Context, database string, collectionName string) (*schemaInfo, error) {
	ret := _m.Called(ctx, database, collectionName)

	var r0 *schemaInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*schemaInfo, error)); ok {
		return rf(ctx, database, collectionName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *schemaInfo); ok {
		r0 = rf(ctx, database, collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schemaInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCollectionSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionSchema'
type MockCache_GetCollectionSchema_Call struct {
	*mock.Call
}

// GetCollectionSchema is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetCollectionSchema(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_GetCollectionSchema_Call {
	return &MockCache_GetCollectionSchema_Call{Call: _e.mock.On("GetCollectionSchema", ctx, database, collectionName)}
}

func (_c *MockCache_GetCollectionSchema_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_GetCollectionSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_GetCollectionSchema_Call) Return(_a0 *schemaInfo, _a1 error) *MockCache_GetCollectionSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetCollectionSchema_Call) RunAndReturn(run func(context.Context, string, string) (*schemaInfo, error)) *MockCache_GetCollectionSchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetCredentialInfo provides a mock function with given fields: ctx, username
func (_m *MockCache) GetCredentialInfo(ctx context.Context, username string) (*internalpb.CredentialInfo, error) {
	ret := _m.Called(ctx, username)

	var r0 *internalpb.CredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*internalpb.CredentialInfo, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *internalpb.CredentialInfo); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.CredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetCredentialInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCredentialInfo'
type MockCache_GetCredentialInfo_Call struct {
	*mock.Call
}

// GetCredentialInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *MockCache_Expecter) GetCredentialInfo(ctx interface{}, username interface{}) *MockCache_GetCredentialInfo_Call {
	return &MockCache_GetCredentialInfo_Call{Call: _e.mock.On("GetCredentialInfo", ctx, username)}
}

func (_c *MockCache_GetCredentialInfo_Call) Run(run func(ctx context.Context, username string)) *MockCache_GetCredentialInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_GetCredentialInfo_Call) Return(_a0 *internalpb.CredentialInfo, _a1 error) *MockCache_GetCredentialInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetCredentialInfo_Call) RunAndReturn(run func(context.Context, string) (*internalpb.CredentialInfo, error)) *MockCache_GetCredentialInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartitionID provides a mock function with given fields: ctx, database, collectionName, partitionName
func (_m *MockCache) GetPartitionID(ctx context.Context, database string, collectionName string, partitionName string) (int64, error) {
	ret := _m.Called(ctx, database, collectionName, partitionName)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (int64, error)); ok {
		return rf(ctx, database, collectionName, partitionName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) int64); ok {
		r0 = rf(ctx, database, collectionName, partitionName)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, database, collectionName, partitionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetPartitionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitionID'
type MockCache_GetPartitionID_Call struct {
	*mock.Call
}

// GetPartitionID is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
//   - partitionName string
func (_e *MockCache_Expecter) GetPartitionID(ctx interface{}, database interface{}, collectionName interface{}, partitionName interface{}) *MockCache_GetPartitionID_Call {
	return &MockCache_GetPartitionID_Call{Call: _e.mock.On("GetPartitionID", ctx, database, collectionName, partitionName)}
}

func (_c *MockCache_GetPartitionID_Call) Run(run func(ctx context.Context, database string, collectionName string, partitionName string)) *MockCache_GetPartitionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockCache_GetPartitionID_Call) Return(_a0 int64, _a1 error) *MockCache_GetPartitionID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetPartitionID_Call) RunAndReturn(run func(context.Context, string, string, string) (int64, error)) *MockCache_GetPartitionID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartitionInfo provides a mock function with given fields: ctx, database, collectionName, partitionName
func (_m *MockCache) GetPartitionInfo(ctx context.Context, database string, collectionName string, partitionName string) (*partitionInfo, error) {
	ret := _m.Called(ctx, database, collectionName, partitionName)

	var r0 *partitionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*partitionInfo, error)); ok {
		return rf(ctx, database, collectionName, partitionName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *partitionInfo); ok {
		r0 = rf(ctx, database, collectionName, partitionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*partitionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, database, collectionName, partitionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetPartitionInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitionInfo'
type MockCache_GetPartitionInfo_Call struct {
	*mock.Call
}

// GetPartitionInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
//   - partitionName string
func (_e *MockCache_Expecter) GetPartitionInfo(ctx interface{}, database interface{}, collectionName interface{}, partitionName interface{}) *MockCache_GetPartitionInfo_Call {
	return &MockCache_GetPartitionInfo_Call{Call: _e.mock.On("GetPartitionInfo", ctx, database, collectionName, partitionName)}
}

func (_c *MockCache_GetPartitionInfo_Call) Run(run func(ctx context.Context, database string, collectionName string, partitionName string)) *MockCache_GetPartitionInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockCache_GetPartitionInfo_Call) Return(_a0 *partitionInfo, _a1 error) *MockCache_GetPartitionInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetPartitionInfo_Call) RunAndReturn(run func(context.Context, string, string, string) (*partitionInfo, error)) *MockCache_GetPartitionInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartitions provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) GetPartitions(ctx context.Context, database string, collectionName string) (map[string]int64, error) {
	ret := _m.Called(ctx, database, collectionName)

	var r0 map[string]int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (map[string]int64, error)); ok {
		return rf(ctx, database, collectionName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[string]int64); ok {
		r0 = rf(ctx, database, collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetPartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitions'
type MockCache_GetPartitions_Call struct {
	*mock.Call
}

// GetPartitions is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetPartitions(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_GetPartitions_Call {
	return &MockCache_GetPartitions_Call{Call: _e.mock.On("GetPartitions", ctx, database, collectionName)}
}

func (_c *MockCache_GetPartitions_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_GetPartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_GetPartitions_Call) Return(_a0 map[string]int64, _a1 error) *MockCache_GetPartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetPartitions_Call) RunAndReturn(run func(context.Context, string, string) (map[string]int64, error)) *MockCache_GetPartitions_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartitionsIndex provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) GetPartitionsIndex(ctx context.Context, database string, collectionName string) ([]string, error) {
	ret := _m.Called(ctx, database, collectionName)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]string, error)); ok {
		return rf(ctx, database, collectionName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []string); ok {
		r0 = rf(ctx, database, collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, database, collectionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetPartitionsIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitionsIndex'
type MockCache_GetPartitionsIndex_Call struct {
	*mock.Call
}

// GetPartitionsIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) GetPartitionsIndex(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_GetPartitionsIndex_Call {
	return &MockCache_GetPartitionsIndex_Call{Call: _e.mock.On("GetPartitionsIndex", ctx, database, collectionName)}
}

func (_c *MockCache_GetPartitionsIndex_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_GetPartitionsIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_GetPartitionsIndex_Call) Return(_a0 []string, _a1 error) *MockCache_GetPartitionsIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetPartitionsIndex_Call) RunAndReturn(run func(context.Context, string, string) ([]string, error)) *MockCache_GetPartitionsIndex_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrivilegeInfo provides a mock function with given fields: ctx
func (_m *MockCache) GetPrivilegeInfo(ctx context.Context) []string {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockCache_GetPrivilegeInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrivilegeInfo'
type MockCache_GetPrivilegeInfo_Call struct {
	*mock.Call
}

// GetPrivilegeInfo is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCache_Expecter) GetPrivilegeInfo(ctx interface{}) *MockCache_GetPrivilegeInfo_Call {
	return &MockCache_GetPrivilegeInfo_Call{Call: _e.mock.On("GetPrivilegeInfo", ctx)}
}

func (_c *MockCache_GetPrivilegeInfo_Call) Run(run func(ctx context.Context)) *MockCache_GetPrivilegeInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCache_GetPrivilegeInfo_Call) Return(_a0 []string) *MockCache_GetPrivilegeInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_GetPrivilegeInfo_Call) RunAndReturn(run func(context.Context) []string) *MockCache_GetPrivilegeInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetShards provides a mock function with given fields: ctx, withCache, database, collectionName, collectionID
func (_m *MockCache) GetShards(ctx context.Context, withCache bool, database string, collectionName string, collectionID int64) (map[string][]nodeInfo, error) {
	ret := _m.Called(ctx, withCache, database, collectionName, collectionID)

	var r0 map[string][]nodeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, string, string, int64) (map[string][]nodeInfo, error)); ok {
		return rf(ctx, withCache, database, collectionName, collectionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, string, string, int64) map[string][]nodeInfo); ok {
		r0 = rf(ctx, withCache, database, collectionName, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]nodeInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, string, string, int64) error); ok {
		r1 = rf(ctx, withCache, database, collectionName, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetShards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShards'
type MockCache_GetShards_Call struct {
	*mock.Call
}

// GetShards is a helper method to define mock.On call
//   - ctx context.Context
//   - withCache bool
//   - database string
//   - collectionName string
//   - collectionID int64
func (_e *MockCache_Expecter) GetShards(ctx interface{}, withCache interface{}, database interface{}, collectionName interface{}, collectionID interface{}) *MockCache_GetShards_Call {
	return &MockCache_GetShards_Call{Call: _e.mock.On("GetShards", ctx, withCache, database, collectionName, collectionID)}
}

func (_c *MockCache_GetShards_Call) Run(run func(ctx context.Context, withCache bool, database string, collectionName string, collectionID int64)) *MockCache_GetShards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(string), args[3].(string), args[4].(int64))
	})
	return _c
}

func (_c *MockCache_GetShards_Call) Return(_a0 map[string][]nodeInfo, _a1 error) *MockCache_GetShards_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetShards_Call) RunAndReturn(run func(context.Context, bool, string, string, int64) (map[string][]nodeInfo, error)) *MockCache_GetShards_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserRole provides a mock function with given fields: username
func (_m *MockCache) GetUserRole(username string) []string {
	ret := _m.Called(username)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockCache_GetUserRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserRole'
type MockCache_GetUserRole_Call struct {
	*mock.Call
}

// GetUserRole is a helper method to define mock.On call
//   - username string
func (_e *MockCache_Expecter) GetUserRole(username interface{}) *MockCache_GetUserRole_Call {
	return &MockCache_GetUserRole_Call{Call: _e.mock.On("GetUserRole", username)}
}

func (_c *MockCache_GetUserRole_Call) Run(run func(username string)) *MockCache_GetUserRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCache_GetUserRole_Call) Return(_a0 []string) *MockCache_GetUserRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_GetUserRole_Call) RunAndReturn(run func(string) []string) *MockCache_GetUserRole_Call {
	_c.Call.Return(run)
	return _c
}

// HasDatabase provides a mock function with given fields: ctx, database
func (_m *MockCache) HasDatabase(ctx context.Context, database string) bool {
	ret := _m.Called(ctx, database)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, database)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCache_HasDatabase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasDatabase'
type MockCache_HasDatabase_Call struct {
	*mock.Call
}

// HasDatabase is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
func (_e *MockCache_Expecter) HasDatabase(ctx interface{}, database interface{}) *MockCache_HasDatabase_Call {
	return &MockCache_HasDatabase_Call{Call: _e.mock.On("HasDatabase", ctx, database)}
}

func (_c *MockCache_HasDatabase_Call) Run(run func(ctx context.Context, database string)) *MockCache_HasDatabase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_HasDatabase_Call) Return(_a0 bool) *MockCache_HasDatabase_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_HasDatabase_Call) RunAndReturn(run func(context.Context, string) bool) *MockCache_HasDatabase_Call {
	_c.Call.Return(run)
	return _c
}

// InitPolicyInfo provides a mock function with given fields: info, userRoles
func (_m *MockCache) InitPolicyInfo(info []string, userRoles []string) {
	_m.Called(info, userRoles)
}

// MockCache_InitPolicyInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitPolicyInfo'
type MockCache_InitPolicyInfo_Call struct {
	*mock.Call
}

// InitPolicyInfo is a helper method to define mock.On call
//   - info []string
//   - userRoles []string
func (_e *MockCache_Expecter) InitPolicyInfo(info interface{}, userRoles interface{}) *MockCache_InitPolicyInfo_Call {
	return &MockCache_InitPolicyInfo_Call{Call: _e.mock.On("InitPolicyInfo", info, userRoles)}
}

func (_c *MockCache_InitPolicyInfo_Call) Run(run func(info []string, userRoles []string)) *MockCache_InitPolicyInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string), args[1].([]string))
	})
	return _c
}

func (_c *MockCache_InitPolicyInfo_Call) Return() *MockCache_InitPolicyInfo_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_InitPolicyInfo_Call) RunAndReturn(run func([]string, []string)) *MockCache_InitPolicyInfo_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshPolicyInfo provides a mock function with given fields: op
func (_m *MockCache) RefreshPolicyInfo(op typeutil.CacheOp) error {
	ret := _m.Called(op)

	var r0 error
	if rf, ok := ret.Get(0).(func(typeutil.CacheOp) error); ok {
		r0 = rf(op)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCache_RefreshPolicyInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshPolicyInfo'
type MockCache_RefreshPolicyInfo_Call struct {
	*mock.Call
}

// RefreshPolicyInfo is a helper method to define mock.On call
//   - op typeutil.CacheOp
func (_e *MockCache_Expecter) RefreshPolicyInfo(op interface{}) *MockCache_RefreshPolicyInfo_Call {
	return &MockCache_RefreshPolicyInfo_Call{Call: _e.mock.On("RefreshPolicyInfo", op)}
}

func (_c *MockCache_RefreshPolicyInfo_Call) Run(run func(op typeutil.CacheOp)) *MockCache_RefreshPolicyInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(typeutil.CacheOp))
	})
	return _c
}

func (_c *MockCache_RefreshPolicyInfo_Call) Return(_a0 error) *MockCache_RefreshPolicyInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_RefreshPolicyInfo_Call) RunAndReturn(run func(typeutil.CacheOp) error) *MockCache_RefreshPolicyInfo_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveAlias provides a mock function with given fields: ctx, database, alias
func (_m *MockCache) RemoveAlias(ctx context.Context, database string, alias string) {
	_m.Called(ctx, database, alias)
}

// MockCache_RemoveAlias_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveAlias'
type MockCache_RemoveAlias_Call struct {
	*mock.Call
}

// RemoveAlias is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - alias string
func (_e *MockCache_Expecter) RemoveAlias(ctx interface{}, database interface{}, alias interface{}) *MockCache_RemoveAlias_Call {
	return &MockCache_RemoveAlias_Call{Call: _e.mock.On("RemoveAlias", ctx, database, alias)}
}

func (_c *MockCache_RemoveAlias_Call) Run(run func(ctx context.Context, database string, alias string)) *MockCache_RemoveAlias_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_RemoveAlias_Call) Return() *MockCache_RemoveAlias_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_RemoveAlias_Call) RunAndReturn(run func(context.Context, string, string)) *MockCache_RemoveAlias_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveCollection provides a mock function with given fields: ctx, database, collectionName
func (_m *MockCache) RemoveCollection(ctx context.Context, database string, collectionName string) {
	_m.Called(ctx, database, collectionName)
}

// MockCache_RemoveCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCollection'
type MockCache_RemoveCollection_Call struct {
	*mock.Call
}

// RemoveCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
func (_e *MockCache_Expecter) RemoveCollection(ctx interface{}, database interface{}, collectionName interface{}) *MockCache_RemoveCollection_Call {
	return &MockCache_RemoveCollection_Call{Call: _e.mock.On("RemoveCollection", ctx, database, collectionName)}
}

func (_c *MockCache_RemoveCollection_Call) Run(run func(ctx context.Context, database string, collectionName string)) *MockCache_RemoveCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockCache_RemoveCollection_Call) Return() *MockCache_RemoveCollection_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_RemoveCollection_Call) RunAndReturn(run func(context.Context, string, string)) *MockCache_RemoveCollection_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveCollectionsByID provides a mock function with given fields: ctx, collectionID
func (_m *MockCache) RemoveCollectionsByID(ctx context.Context, collectionID int64) []string {
	ret := _m.Called(ctx, collectionID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, int64) []string); ok {
		r0 = rf(ctx, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockCache_RemoveCollectionsByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCollectionsByID'
type MockCache_RemoveCollectionsByID_Call struct {
	*mock.Call
}

// RemoveCollectionsByID is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *MockCache_Expecter) RemoveCollectionsByID(ctx interface{}, collectionID interface{}) *MockCache_RemoveCollectionsByID_Call {
	return &MockCache_RemoveCollectionsByID_Call{Call: _e.mock.On("RemoveCollectionsByID", ctx, collectionID)}
}

func (_c *MockCache_RemoveCollectionsByID_Call) Run(run func(ctx context.Context, collectionID int64)) *MockCache_RemoveCollectionsByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockCache_RemoveCollectionsByID_Call) Return(_a0 []string) *MockCache_RemoveCollectionsByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_RemoveCollectionsByID_Call) RunAndReturn(run func(context.Context, int64) []string) *MockCache_RemoveCollectionsByID_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveCredential provides a mock function with given fields: username
func (_m *MockCache) RemoveCredential(username string) {
	_m.Called(username)
}

// MockCache_RemoveCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCredential'
type MockCache_RemoveCredential_Call struct {
	*mock.Call
}

// RemoveCredential is a helper method to define mock.On call
//   - username string
func (_e *MockCache_Expecter) RemoveCredential(username interface{}) *MockCache_RemoveCredential_Call {
	return &MockCache_RemoveCredential_Call{Call: _e.mock.On("RemoveCredential", username)}
}

func (_c *MockCache_RemoveCredential_Call) Run(run func(username string)) *MockCache_RemoveCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCache_RemoveCredential_Call) Return() *MockCache_RemoveCredential_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_RemoveCredential_Call) RunAndReturn(run func(string)) *MockCache_RemoveCredential_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveDatabase provides a mock function with given fields: ctx, database
func (_m *MockCache) RemoveDatabase(ctx context.Context, database string) {
	_m.Called(ctx, database)
}

// MockCache_RemoveDatabase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveDatabase'
type MockCache_RemoveDatabase_Call struct {
	*mock.Call
}

// RemoveDatabase is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
func (_e *MockCache_Expecter) RemoveDatabase(ctx interface{}, database interface{}) *MockCache_RemoveDatabase_Call {
	return &MockCache_RemoveDatabase_Call{Call: _e.mock.On("RemoveDatabase", ctx, database)}
}

func (_c *MockCache_RemoveDatabase_Call) Run(run func(ctx context.Context, database string)) *MockCache_RemoveDatabase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_RemoveDatabase_Call) Return() *MockCache_RemoveDatabase_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_RemoveDatabase_Call) RunAndReturn(run func(context.Context, string)) *MockCache_RemoveDatabase_Call {
	_c.Call.Return(run)
	return _c
}

// RemovePartition provides a mock function with given fields: ctx, database, collectionName, partitionName
func (_m *MockCache) RemovePartition(ctx context.Context, database string, collectionName string, partitionName string) {
	_m.Called(ctx, database, collectionName, partitionName)
}

// MockCache_RemovePartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemovePartition'
type MockCache_RemovePartition_Call struct {
	*mock.Call
}

// RemovePartition is a helper method to define mock.On call
//   - ctx context.Context
//   - database string
//   - collectionName string
//   - partitionName string
func (_e *MockCache_Expecter) RemovePartition(ctx interface{}, database interface{}, collectionName interface{}, partitionName interface{}) *MockCache_RemovePartition_Call {
	return &MockCache_RemovePartition_Call{Call: _e.mock.On("RemovePartition", ctx, database, collectionName, partitionName)}
}

func (_c *MockCache_RemovePartition_Call) Run(run func(ctx context.Context, database string, collectionName string, partitionName string)) *MockCache_RemovePartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockCache_RemovePartition_Call) Return() *MockCache_RemovePartition_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_RemovePartition_Call) RunAndReturn(run func(context.Context, string, string, string)) *MockCache_RemovePartition_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCredential provides a mock function with given fields: credInfo
func (_m *MockCache) UpdateCredential(credInfo *internalpb.CredentialInfo) {
	_m.Called(credInfo)
}

// MockCache_UpdateCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCredential'
type MockCache_UpdateCredential_Call struct {
	*mock.Call
}

// UpdateCredential is a helper method to define mock.On call
//   - credInfo *internalpb.CredentialInfo
func (_e *MockCache_Expecter) UpdateCredential(credInfo interface{}) *MockCache_UpdateCredential_Call {
	return &MockCache_UpdateCredential_Call{Call: _e.mock.On("UpdateCredential", credInfo)}
}

func (_c *MockCache_UpdateCredential_Call) Run(run func(credInfo *internalpb.CredentialInfo)) *MockCache_UpdateCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*internalpb.CredentialInfo))
	})
	return _c
}

func (_c *MockCache_UpdateCredential_Call) Return() *MockCache_UpdateCredential_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_UpdateCredential_Call) RunAndReturn(run func(*internalpb.CredentialInfo)) *MockCache_UpdateCredential_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCache creates a new instance of MockCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCache {
	mock := &MockCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
