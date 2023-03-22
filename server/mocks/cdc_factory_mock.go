// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	reader "github.com/zilliztech/milvus-cdc/core/reader"
	"github.com/zilliztech/milvus-cdc/core/util"
	writer "github.com/zilliztech/milvus-cdc/core/writer"
)

// CDCFactory is an autogenerated mock type for the CDCFactory type
type CDCFactory struct {
	util.CDCMark
	mock.Mock
}

// NewReader provides a mock function with given fields:
func (_m *CDCFactory) NewReader() (reader.CDCReader, error) {
	ret := _m.Called()

	var r0 reader.CDCReader
	var r1 error
	if rf, ok := ret.Get(0).(func() (reader.CDCReader, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() reader.CDCReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reader.CDCReader)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWriter provides a mock function with given fields:
func (_m *CDCFactory) NewWriter() (writer.CDCWriter, error) {
	ret := _m.Called()

	var r0 writer.CDCWriter
	var r1 error
	if rf, ok := ret.Get(0).(func() (writer.CDCWriter, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() writer.CDCWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(writer.CDCWriter)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// cdc provides a mock function with given fields:
func (_m *CDCFactory) cdc() {
	_m.Called()
}

type mockConstructorTestingTNewCDCFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewCDCFactory creates a new instance of CDCFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCDCFactory(t mockConstructorTestingTNewCDCFactory) *CDCFactory {
	mock := &CDCFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
