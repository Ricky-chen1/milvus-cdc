// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	clientv3 "go.etcd.io/etcd/client/v3"

	mock "github.com/stretchr/testify/mock"
)

// KVApi is an autogenerated mock type for the KVApi type
type KVApi struct {
	mock.Mock
}

type KVApi_Expecter struct {
	mock *mock.Mock
}

func (_m *KVApi) EXPECT() *KVApi_Expecter {
	return &KVApi_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *KVApi) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KVApi_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type KVApi_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *KVApi_Expecter) Close() *KVApi_Close_Call {
	return &KVApi_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *KVApi_Close_Call) Run(run func()) *KVApi_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *KVApi_Close_Call) Return(_a0 error) *KVApi_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KVApi_Close_Call) RunAndReturn(run func() error) *KVApi_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Compact provides a mock function with given fields: ctx, rev, opts
func (_m *KVApi) Compact(ctx context.Context, rev int64, opts ...clientv3.CompactOption) (*clientv3.CompactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, rev)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientv3.CompactResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...clientv3.CompactOption) (*clientv3.CompactResponse, error)); ok {
		return rf(ctx, rev, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...clientv3.CompactOption) *clientv3.CompactResponse); ok {
		r0 = rf(ctx, rev, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.CompactResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...clientv3.CompactOption) error); ok {
		r1 = rf(ctx, rev, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KVApi_Compact_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compact'
type KVApi_Compact_Call struct {
	*mock.Call
}

// Compact is a helper method to define mock.On call
//   - ctx context.Context
//   - rev int64
//   - opts ...clientv3.CompactOption
func (_e *KVApi_Expecter) Compact(ctx interface{}, rev interface{}, opts ...interface{}) *KVApi_Compact_Call {
	return &KVApi_Compact_Call{Call: _e.mock.On("Compact",
		append([]interface{}{ctx, rev}, opts...)...)}
}

func (_c *KVApi_Compact_Call) Run(run func(ctx context.Context, rev int64, opts ...clientv3.CompactOption)) *KVApi_Compact_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.CompactOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.CompactOption)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *KVApi_Compact_Call) Return(_a0 *clientv3.CompactResponse, _a1 error) *KVApi_Compact_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KVApi_Compact_Call) RunAndReturn(run func(context.Context, int64, ...clientv3.CompactOption) (*clientv3.CompactResponse, error)) *KVApi_Compact_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, key, opts
func (_m *KVApi) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientv3.DeleteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) (*clientv3.DeleteResponse, error)); ok {
		return rf(ctx, key, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) *clientv3.DeleteResponse); ok {
		r0 = rf(ctx, key, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.DeleteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...clientv3.OpOption) error); ok {
		r1 = rf(ctx, key, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KVApi_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type KVApi_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - opts ...clientv3.OpOption
func (_e *KVApi_Expecter) Delete(ctx interface{}, key interface{}, opts ...interface{}) *KVApi_Delete_Call {
	return &KVApi_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{ctx, key}, opts...)...)}
}

func (_c *KVApi_Delete_Call) Run(run func(ctx context.Context, key string, opts ...clientv3.OpOption)) *KVApi_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *KVApi_Delete_Call) Return(_a0 *clientv3.DeleteResponse, _a1 error) *KVApi_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KVApi_Delete_Call) RunAndReturn(run func(context.Context, string, ...clientv3.OpOption) (*clientv3.DeleteResponse, error)) *KVApi_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Do provides a mock function with given fields: ctx, op
func (_m *KVApi) Do(ctx context.Context, op clientv3.Op) (clientv3.OpResponse, error) {
	ret := _m.Called(ctx, op)

	var r0 clientv3.OpResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.Op) (clientv3.OpResponse, error)); ok {
		return rf(ctx, op)
	}
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.Op) clientv3.OpResponse); ok {
		r0 = rf(ctx, op)
	} else {
		r0 = ret.Get(0).(clientv3.OpResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, clientv3.Op) error); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KVApi_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type KVApi_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//   - ctx context.Context
//   - op clientv3.Op
func (_e *KVApi_Expecter) Do(ctx interface{}, op interface{}) *KVApi_Do_Call {
	return &KVApi_Do_Call{Call: _e.mock.On("Do", ctx, op)}
}

func (_c *KVApi_Do_Call) Run(run func(ctx context.Context, op clientv3.Op)) *KVApi_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(clientv3.Op))
	})
	return _c
}

func (_c *KVApi_Do_Call) Return(_a0 clientv3.OpResponse, _a1 error) *KVApi_Do_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KVApi_Do_Call) RunAndReturn(run func(context.Context, clientv3.Op) (clientv3.OpResponse, error)) *KVApi_Do_Call {
	_c.Call.Return(run)
	return _c
}

// Endpoints provides a mock function with given fields:
func (_m *KVApi) Endpoints() []string {
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

// KVApi_Endpoints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Endpoints'
type KVApi_Endpoints_Call struct {
	*mock.Call
}

// Endpoints is a helper method to define mock.On call
func (_e *KVApi_Expecter) Endpoints() *KVApi_Endpoints_Call {
	return &KVApi_Endpoints_Call{Call: _e.mock.On("Endpoints")}
}

func (_c *KVApi_Endpoints_Call) Run(run func()) *KVApi_Endpoints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *KVApi_Endpoints_Call) Return(_a0 []string) *KVApi_Endpoints_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KVApi_Endpoints_Call) RunAndReturn(run func() []string) *KVApi_Endpoints_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key, opts
func (_m *KVApi) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientv3.GetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) (*clientv3.GetResponse, error)); ok {
		return rf(ctx, key, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) *clientv3.GetResponse); ok {
		r0 = rf(ctx, key, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.GetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...clientv3.OpOption) error); ok {
		r1 = rf(ctx, key, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KVApi_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type KVApi_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - opts ...clientv3.OpOption
func (_e *KVApi_Expecter) Get(ctx interface{}, key interface{}, opts ...interface{}) *KVApi_Get_Call {
	return &KVApi_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, key}, opts...)...)}
}

func (_c *KVApi_Get_Call) Run(run func(ctx context.Context, key string, opts ...clientv3.OpOption)) *KVApi_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *KVApi_Get_Call) Return(_a0 *clientv3.GetResponse, _a1 error) *KVApi_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KVApi_Get_Call) RunAndReturn(run func(context.Context, string, ...clientv3.OpOption) (*clientv3.GetResponse, error)) *KVApi_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: ctx, key, val, opts
func (_m *KVApi) Put(ctx context.Context, key string, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key, val)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientv3.PutResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...clientv3.OpOption) (*clientv3.PutResponse, error)); ok {
		return rf(ctx, key, val, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...clientv3.OpOption) *clientv3.PutResponse); ok {
		r0 = rf(ctx, key, val, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.PutResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...clientv3.OpOption) error); ok {
		r1 = rf(ctx, key, val, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KVApi_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type KVApi_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - val string
//   - opts ...clientv3.OpOption
func (_e *KVApi_Expecter) Put(ctx interface{}, key interface{}, val interface{}, opts ...interface{}) *KVApi_Put_Call {
	return &KVApi_Put_Call{Call: _e.mock.On("Put",
		append([]interface{}{ctx, key, val}, opts...)...)}
}

func (_c *KVApi_Put_Call) Run(run func(ctx context.Context, key string, val string, opts ...clientv3.OpOption)) *KVApi_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *KVApi_Put_Call) Return(_a0 *clientv3.PutResponse, _a1 error) *KVApi_Put_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KVApi_Put_Call) RunAndReturn(run func(context.Context, string, string, ...clientv3.OpOption) (*clientv3.PutResponse, error)) *KVApi_Put_Call {
	_c.Call.Return(run)
	return _c
}

// RequestProgress provides a mock function with given fields: ctx
func (_m *KVApi) RequestProgress(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KVApi_RequestProgress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestProgress'
type KVApi_RequestProgress_Call struct {
	*mock.Call
}

// RequestProgress is a helper method to define mock.On call
//   - ctx context.Context
func (_e *KVApi_Expecter) RequestProgress(ctx interface{}) *KVApi_RequestProgress_Call {
	return &KVApi_RequestProgress_Call{Call: _e.mock.On("RequestProgress", ctx)}
}

func (_c *KVApi_RequestProgress_Call) Run(run func(ctx context.Context)) *KVApi_RequestProgress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *KVApi_RequestProgress_Call) Return(_a0 error) *KVApi_RequestProgress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KVApi_RequestProgress_Call) RunAndReturn(run func(context.Context) error) *KVApi_RequestProgress_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with given fields: ctx, endpoint
func (_m *KVApi) Status(ctx context.Context, endpoint string) (*clientv3.StatusResponse, error) {
	ret := _m.Called(ctx, endpoint)

	var r0 *clientv3.StatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*clientv3.StatusResponse, error)); ok {
		return rf(ctx, endpoint)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *clientv3.StatusResponse); ok {
		r0 = rf(ctx, endpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.StatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, endpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KVApi_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type KVApi_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
//   - ctx context.Context
//   - endpoint string
func (_e *KVApi_Expecter) Status(ctx interface{}, endpoint interface{}) *KVApi_Status_Call {
	return &KVApi_Status_Call{Call: _e.mock.On("Status", ctx, endpoint)}
}

func (_c *KVApi_Status_Call) Run(run func(ctx context.Context, endpoint string)) *KVApi_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *KVApi_Status_Call) Return(_a0 *clientv3.StatusResponse, _a1 error) *KVApi_Status_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KVApi_Status_Call) RunAndReturn(run func(context.Context, string) (*clientv3.StatusResponse, error)) *KVApi_Status_Call {
	_c.Call.Return(run)
	return _c
}

// Txn provides a mock function with given fields: ctx
func (_m *KVApi) Txn(ctx context.Context) clientv3.Txn {
	ret := _m.Called(ctx)

	var r0 clientv3.Txn
	if rf, ok := ret.Get(0).(func(context.Context) clientv3.Txn); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientv3.Txn)
		}
	}

	return r0
}

// KVApi_Txn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Txn'
type KVApi_Txn_Call struct {
	*mock.Call
}

// Txn is a helper method to define mock.On call
//   - ctx context.Context
func (_e *KVApi_Expecter) Txn(ctx interface{}) *KVApi_Txn_Call {
	return &KVApi_Txn_Call{Call: _e.mock.On("Txn", ctx)}
}

func (_c *KVApi_Txn_Call) Run(run func(ctx context.Context)) *KVApi_Txn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *KVApi_Txn_Call) Return(_a0 clientv3.Txn) *KVApi_Txn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KVApi_Txn_Call) RunAndReturn(run func(context.Context) clientv3.Txn) *KVApi_Txn_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, key, opts
func (_m *KVApi) Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 clientv3.WatchChan
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) clientv3.WatchChan); ok {
		r0 = rf(ctx, key, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientv3.WatchChan)
		}
	}

	return r0
}

// KVApi_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type KVApi_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - opts ...clientv3.OpOption
func (_e *KVApi_Expecter) Watch(ctx interface{}, key interface{}, opts ...interface{}) *KVApi_Watch_Call {
	return &KVApi_Watch_Call{Call: _e.mock.On("Watch",
		append([]interface{}{ctx, key}, opts...)...)}
}

func (_c *KVApi_Watch_Call) Run(run func(ctx context.Context, key string, opts ...clientv3.OpOption)) *KVApi_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *KVApi_Watch_Call) Return(_a0 clientv3.WatchChan) *KVApi_Watch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KVApi_Watch_Call) RunAndReturn(run func(context.Context, string, ...clientv3.OpOption) clientv3.WatchChan) *KVApi_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewKVApi interface {
	mock.TestingT
	Cleanup(func())
}

// NewKVApi creates a new instance of KVApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKVApi(t mockConstructorTestingTNewKVApi) *KVApi {
	mock := &KVApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
