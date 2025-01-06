// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/zilliztech/milvus-cdc/core/api"

	mock "github.com/stretchr/testify/mock"
)

// ReplicateMeta is an autogenerated mock type for the ReplicateMeta type
type ReplicateMeta struct {
	mock.Mock
}

type ReplicateMeta_Expecter struct {
	mock *mock.Mock
}

func (_m *ReplicateMeta) EXPECT() *ReplicateMeta_Expecter {
	return &ReplicateMeta_Expecter{mock: &_m.Mock}
}

// GetTaskDropCollectionMsg provides a mock function with given fields: ctx, taskID, msgID
func (_m *ReplicateMeta) GetTaskDropCollectionMsg(ctx context.Context, taskID string, msgID string) ([]api.TaskDropCollectionMsg, error) {
	ret := _m.Called(ctx, taskID, msgID)

	var r0 []api.TaskDropCollectionMsg
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]api.TaskDropCollectionMsg, error)); ok {
		return rf(ctx, taskID, msgID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []api.TaskDropCollectionMsg); ok {
		r0 = rf(ctx, taskID, msgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.TaskDropCollectionMsg)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, taskID, msgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicateMeta_GetTaskDropCollectionMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTaskDropCollectionMsg'
type ReplicateMeta_GetTaskDropCollectionMsg_Call struct {
	*mock.Call
}

// GetTaskDropCollectionMsg is a helper method to define mock.On call
//  - ctx context.Context
//  - taskID string
//  - msgID string
func (_e *ReplicateMeta_Expecter) GetTaskDropCollectionMsg(ctx interface{}, taskID interface{}, msgID interface{}) *ReplicateMeta_GetTaskDropCollectionMsg_Call {
	return &ReplicateMeta_GetTaskDropCollectionMsg_Call{Call: _e.mock.On("GetTaskDropCollectionMsg", ctx, taskID, msgID)}
}

func (_c *ReplicateMeta_GetTaskDropCollectionMsg_Call) Run(run func(ctx context.Context, taskID string, msgID string)) *ReplicateMeta_GetTaskDropCollectionMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ReplicateMeta_GetTaskDropCollectionMsg_Call) Return(_a0 []api.TaskDropCollectionMsg, _a1 error) *ReplicateMeta_GetTaskDropCollectionMsg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicateMeta_GetTaskDropCollectionMsg_Call) RunAndReturn(run func(context.Context, string, string) ([]api.TaskDropCollectionMsg, error)) *ReplicateMeta_GetTaskDropCollectionMsg_Call {
	_c.Call.Return(run)
	return _c
}

// GetTaskDropPartitionMsg provides a mock function with given fields: ctx, taskID, msgID
func (_m *ReplicateMeta) GetTaskDropPartitionMsg(ctx context.Context, taskID string, msgID string) ([]api.TaskDropPartitionMsg, error) {
	ret := _m.Called(ctx, taskID, msgID)

	var r0 []api.TaskDropPartitionMsg
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]api.TaskDropPartitionMsg, error)); ok {
		return rf(ctx, taskID, msgID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []api.TaskDropPartitionMsg); ok {
		r0 = rf(ctx, taskID, msgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.TaskDropPartitionMsg)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, taskID, msgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicateMeta_GetTaskDropPartitionMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTaskDropPartitionMsg'
type ReplicateMeta_GetTaskDropPartitionMsg_Call struct {
	*mock.Call
}

// GetTaskDropPartitionMsg is a helper method to define mock.On call
//  - ctx context.Context
//  - taskID string
//  - msgID string
func (_e *ReplicateMeta_Expecter) GetTaskDropPartitionMsg(ctx interface{}, taskID interface{}, msgID interface{}) *ReplicateMeta_GetTaskDropPartitionMsg_Call {
	return &ReplicateMeta_GetTaskDropPartitionMsg_Call{Call: _e.mock.On("GetTaskDropPartitionMsg", ctx, taskID, msgID)}
}

func (_c *ReplicateMeta_GetTaskDropPartitionMsg_Call) Run(run func(ctx context.Context, taskID string, msgID string)) *ReplicateMeta_GetTaskDropPartitionMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ReplicateMeta_GetTaskDropPartitionMsg_Call) Return(_a0 []api.TaskDropPartitionMsg, _a1 error) *ReplicateMeta_GetTaskDropPartitionMsg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicateMeta_GetTaskDropPartitionMsg_Call) RunAndReturn(run func(context.Context, string, string) ([]api.TaskDropPartitionMsg, error)) *ReplicateMeta_GetTaskDropPartitionMsg_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveTaskMsg provides a mock function with given fields: ctx, taskID, msgID
func (_m *ReplicateMeta) RemoveTaskMsg(ctx context.Context, taskID string, msgID string) error {
	ret := _m.Called(ctx, taskID, msgID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, taskID, msgID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReplicateMeta_RemoveTaskMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveTaskMsg'
type ReplicateMeta_RemoveTaskMsg_Call struct {
	*mock.Call
}

// RemoveTaskMsg is a helper method to define mock.On call
//  - ctx context.Context
//  - taskID string
//  - msgID string
func (_e *ReplicateMeta_Expecter) RemoveTaskMsg(ctx interface{}, taskID interface{}, msgID interface{}) *ReplicateMeta_RemoveTaskMsg_Call {
	return &ReplicateMeta_RemoveTaskMsg_Call{Call: _e.mock.On("RemoveTaskMsg", ctx, taskID, msgID)}
}

func (_c *ReplicateMeta_RemoveTaskMsg_Call) Run(run func(ctx context.Context, taskID string, msgID string)) *ReplicateMeta_RemoveTaskMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ReplicateMeta_RemoveTaskMsg_Call) Return(_a0 error) *ReplicateMeta_RemoveTaskMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReplicateMeta_RemoveTaskMsg_Call) RunAndReturn(run func(context.Context, string, string) error) *ReplicateMeta_RemoveTaskMsg_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTaskDropCollectionMsg provides a mock function with given fields: ctx, msg
func (_m *ReplicateMeta) UpdateTaskDropCollectionMsg(ctx context.Context, msg api.TaskDropCollectionMsg) (bool, error) {
	ret := _m.Called(ctx, msg)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, api.TaskDropCollectionMsg) (bool, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.TaskDropCollectionMsg) bool); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.TaskDropCollectionMsg) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicateMeta_UpdateTaskDropCollectionMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTaskDropCollectionMsg'
type ReplicateMeta_UpdateTaskDropCollectionMsg_Call struct {
	*mock.Call
}

// UpdateTaskDropCollectionMsg is a helper method to define mock.On call
//  - ctx context.Context
//  - msg api.TaskDropCollectionMsg
func (_e *ReplicateMeta_Expecter) UpdateTaskDropCollectionMsg(ctx interface{}, msg interface{}) *ReplicateMeta_UpdateTaskDropCollectionMsg_Call {
	return &ReplicateMeta_UpdateTaskDropCollectionMsg_Call{Call: _e.mock.On("UpdateTaskDropCollectionMsg", ctx, msg)}
}

func (_c *ReplicateMeta_UpdateTaskDropCollectionMsg_Call) Run(run func(ctx context.Context, msg api.TaskDropCollectionMsg)) *ReplicateMeta_UpdateTaskDropCollectionMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(api.TaskDropCollectionMsg))
	})
	return _c
}

func (_c *ReplicateMeta_UpdateTaskDropCollectionMsg_Call) Return(_a0 bool, _a1 error) *ReplicateMeta_UpdateTaskDropCollectionMsg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicateMeta_UpdateTaskDropCollectionMsg_Call) RunAndReturn(run func(context.Context, api.TaskDropCollectionMsg) (bool, error)) *ReplicateMeta_UpdateTaskDropCollectionMsg_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTaskDropPartitionMsg provides a mock function with given fields: ctx, msg
func (_m *ReplicateMeta) UpdateTaskDropPartitionMsg(ctx context.Context, msg api.TaskDropPartitionMsg) (bool, error) {
	ret := _m.Called(ctx, msg)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, api.TaskDropPartitionMsg) (bool, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.TaskDropPartitionMsg) bool); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.TaskDropPartitionMsg) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicateMeta_UpdateTaskDropPartitionMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTaskDropPartitionMsg'
type ReplicateMeta_UpdateTaskDropPartitionMsg_Call struct {
	*mock.Call
}

// UpdateTaskDropPartitionMsg is a helper method to define mock.On call
//  - ctx context.Context
//  - msg api.TaskDropPartitionMsg
func (_e *ReplicateMeta_Expecter) UpdateTaskDropPartitionMsg(ctx interface{}, msg interface{}) *ReplicateMeta_UpdateTaskDropPartitionMsg_Call {
	return &ReplicateMeta_UpdateTaskDropPartitionMsg_Call{Call: _e.mock.On("UpdateTaskDropPartitionMsg", ctx, msg)}
}

func (_c *ReplicateMeta_UpdateTaskDropPartitionMsg_Call) Run(run func(ctx context.Context, msg api.TaskDropPartitionMsg)) *ReplicateMeta_UpdateTaskDropPartitionMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(api.TaskDropPartitionMsg))
	})
	return _c
}

func (_c *ReplicateMeta_UpdateTaskDropPartitionMsg_Call) Return(_a0 bool, _a1 error) *ReplicateMeta_UpdateTaskDropPartitionMsg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicateMeta_UpdateTaskDropPartitionMsg_Call) RunAndReturn(run func(context.Context, api.TaskDropPartitionMsg) (bool, error)) *ReplicateMeta_UpdateTaskDropPartitionMsg_Call {
	_c.Call.Return(run)
	return _c
}

// NewReplicateMeta creates a new instance of ReplicateMeta. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReplicateMeta(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReplicateMeta {
	mock := &ReplicateMeta{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
