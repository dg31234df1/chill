// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dbmodel "github.com/milvus-io/milvus/internal/metastore/db/dbmodel"
	mock "github.com/stretchr/testify/mock"
)

// IPartitionDb is an autogenerated mock type for the IPartitionDb type
type IPartitionDb struct {
	mock.Mock
}

// GetByCollectionID provides a mock function with given fields: tenantID, collectionID, ts
func (_m *IPartitionDb) GetByCollectionID(tenantID string, collectionID int64, ts uint64) ([]*dbmodel.Partition, error) {
	ret := _m.Called(tenantID, collectionID, ts)

	var r0 []*dbmodel.Partition
	if rf, ok := ret.Get(0).(func(string, int64, uint64) []*dbmodel.Partition); ok {
		r0 = rf(tenantID, collectionID, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.Partition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64, uint64) error); ok {
		r1 = rf(tenantID, collectionID, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: in
func (_m *IPartitionDb) Insert(in []*dbmodel.Partition) error {
	ret := _m.Called(in)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*dbmodel.Partition) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIPartitionDb interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPartitionDb creates a new instance of IPartitionDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPartitionDb(t mockConstructorTestingTNewIPartitionDb) *IPartitionDb {
	mock := &IPartitionDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
