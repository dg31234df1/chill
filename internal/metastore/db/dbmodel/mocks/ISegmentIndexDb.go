// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dbmodel "github.com/milvus-io/milvus/internal/metastore/db/dbmodel"
	mock "github.com/stretchr/testify/mock"
)

// ISegmentIndexDb is an autogenerated mock type for the ISegmentIndexDb type
type ISegmentIndexDb struct {
	mock.Mock
}

// Get provides a mock function with given fields: tenantID, collectionID, buildID
func (_m *ISegmentIndexDb) Get(tenantID string, collectionID int64, buildID int64) ([]*dbmodel.SegmentIndexResult, error) {
	ret := _m.Called(tenantID, collectionID, buildID)

	var r0 []*dbmodel.SegmentIndexResult
	if rf, ok := ret.Get(0).(func(string, int64, int64) []*dbmodel.SegmentIndexResult); ok {
		r0 = rf(tenantID, collectionID, buildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.SegmentIndexResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64, int64) error); ok {
		r1 = rf(tenantID, collectionID, buildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: in
func (_m *ISegmentIndexDb) Insert(in []*dbmodel.SegmentIndex) error {
	ret := _m.Called(in)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*dbmodel.SegmentIndex) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: tenantID
func (_m *ISegmentIndexDb) List(tenantID string) ([]*dbmodel.SegmentIndexResult, error) {
	ret := _m.Called(tenantID)

	var r0 []*dbmodel.SegmentIndexResult
	if rf, ok := ret.Get(0).(func(string) []*dbmodel.SegmentIndexResult); ok {
		r0 = rf(tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.SegmentIndexResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkDeleted provides a mock function with given fields: tenantID, in
func (_m *ISegmentIndexDb) MarkDeleted(tenantID string, in []*dbmodel.SegmentIndex) error {
	ret := _m.Called(tenantID, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []*dbmodel.SegmentIndex) error); ok {
		r0 = rf(tenantID, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkDeletedByBuildID provides a mock function with given fields: tenantID, idxID
func (_m *ISegmentIndexDb) MarkDeletedByBuildID(tenantID string, idxID int64) error {
	ret := _m.Called(tenantID, idxID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(tenantID, idxID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkDeletedByCollectionID provides a mock function with given fields: tenantID, collID
func (_m *ISegmentIndexDb) MarkDeletedByCollectionID(tenantID string, collID int64) error {
	ret := _m.Called(tenantID, collID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(tenantID, collID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: in
func (_m *ISegmentIndexDb) Update(in *dbmodel.SegmentIndex) error {
	ret := _m.Called(in)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbmodel.SegmentIndex) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewISegmentIndexDb interface {
	mock.TestingT
	Cleanup(func())
}

// NewISegmentIndexDb creates a new instance of ISegmentIndexDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewISegmentIndexDb(t mockConstructorTestingTNewISegmentIndexDb) *ISegmentIndexDb {
	mock := &ISegmentIndexDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
