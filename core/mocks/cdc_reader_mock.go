// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/zilliztech/milvus-cdc/core/model"
	"github.com/zilliztech/milvus-cdc/core/util"
)

// CDCReader is an autogenerated mock type for the CDCReader type
type CDCReader struct {
	util.CDCMark
	mock.Mock
}

// QuitRead provides a mock function with given fields: ctx
func (_m *CDCReader) QuitRead(ctx context.Context) {
	_m.Called(ctx)
}

// StartRead provides a mock function with given fields: ctx
func (_m *CDCReader) StartRead(ctx context.Context) <-chan *model.CDCData {
	ret := _m.Called(ctx)

	var r0 <-chan *model.CDCData
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *model.CDCData); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *model.CDCData)
		}
	}

	return r0
}

// cdc provides a mock function with given fields:
func (_m *CDCReader) cdc() {
	_m.Called()
}

type mockConstructorTestingTNewCDCReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewCDCReader creates a new instance of CDCReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCDCReader(t mockConstructorTestingTNewCDCReader) *CDCReader {
	mock := &CDCReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
