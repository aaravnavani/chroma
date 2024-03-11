// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	dbmodel "github.com/chroma-core/chroma/go/pkg/metastore/db/dbmodel"
	mock "github.com/stretchr/testify/mock"

	types "github.com/chroma-core/chroma/go/pkg/types"
)

// ISegmentDb is an autogenerated mock type for the ISegmentDb type
type ISegmentDb struct {
	mock.Mock
}

// DeleteAll provides a mock function with given fields:
func (_m *ISegmentDb) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSegmentByID provides a mock function with given fields: id
func (_m *ISegmentDb) DeleteSegmentByID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSegments provides a mock function with given fields: id, segmentType, scope, topic, collectionID
func (_m *ISegmentDb) GetSegments(id types.UniqueID, segmentType *string, scope *string, topic *string, collectionID types.UniqueID) ([]*dbmodel.SegmentAndMetadata, error) {
	ret := _m.Called(id, segmentType, scope, topic, collectionID)

	var r0 []*dbmodel.SegmentAndMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(types.UniqueID, *string, *string, *string, types.UniqueID) ([]*dbmodel.SegmentAndMetadata, error)); ok {
		return rf(id, segmentType, scope, topic, collectionID)
	}
	if rf, ok := ret.Get(0).(func(types.UniqueID, *string, *string, *string, types.UniqueID) []*dbmodel.SegmentAndMetadata); ok {
		r0 = rf(id, segmentType, scope, topic, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.SegmentAndMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(types.UniqueID, *string, *string, *string, types.UniqueID) error); ok {
		r1 = rf(id, segmentType, scope, topic, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *ISegmentDb) Insert(_a0 *dbmodel.Segment) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbmodel.Segment) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *ISegmentDb) Update(_a0 *dbmodel.UpdateSegment) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbmodel.UpdateSegment) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewISegmentDb creates a new instance of ISegmentDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewISegmentDb(t interface {
	mock.TestingT
	Cleanup(func())
}) *ISegmentDb {
	mock := &ISegmentDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}