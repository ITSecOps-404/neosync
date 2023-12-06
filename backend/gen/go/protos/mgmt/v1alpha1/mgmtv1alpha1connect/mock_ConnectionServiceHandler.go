// Code generated by mockery. DO NOT EDIT.

package mgmtv1alpha1connect

import (
	context "context"

	connect "connectrpc.com/connect"

	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MockConnectionServiceHandler is an autogenerated mock type for the ConnectionServiceHandler type
type MockConnectionServiceHandler struct {
	mock.Mock
}

type MockConnectionServiceHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnectionServiceHandler) EXPECT() *MockConnectionServiceHandler_Expecter {
	return &MockConnectionServiceHandler_Expecter{mock: &_m.Mock}
}

// CheckConnectionConfig provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) CheckConnectionConfig(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) (*connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckConnectionConfig")
	}

	var r0 *connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) (*connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) *connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_CheckConnectionConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckConnectionConfig'
type MockConnectionServiceHandler_CheckConnectionConfig_Call struct {
	*mock.Call
}

// CheckConnectionConfig is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]
func (_e *MockConnectionServiceHandler_Expecter) CheckConnectionConfig(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_CheckConnectionConfig_Call {
	return &MockConnectionServiceHandler_CheckConnectionConfig_Call{Call: _e.mock.On("CheckConnectionConfig", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_CheckConnectionConfig_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest])) *MockConnectionServiceHandler_CheckConnectionConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_CheckConnectionConfig_Call) Return(_a0 *connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse], _a1 error) *MockConnectionServiceHandler_CheckConnectionConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_CheckConnectionConfig_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) (*connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse], error)) *MockConnectionServiceHandler_CheckConnectionConfig_Call {
	_c.Call.Return(run)
	return _c
}

// CheckSqlQuery provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) CheckSqlQuery(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) (*connect.Response[mgmtv1alpha1.CheckSqlQueryResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckSqlQuery")
	}

	var r0 *connect.Response[mgmtv1alpha1.CheckSqlQueryResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) (*connect.Response[mgmtv1alpha1.CheckSqlQueryResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) *connect.Response[mgmtv1alpha1.CheckSqlQueryResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.CheckSqlQueryResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_CheckSqlQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckSqlQuery'
type MockConnectionServiceHandler_CheckSqlQuery_Call struct {
	*mock.Call
}

// CheckSqlQuery is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]
func (_e *MockConnectionServiceHandler_Expecter) CheckSqlQuery(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_CheckSqlQuery_Call {
	return &MockConnectionServiceHandler_CheckSqlQuery_Call{Call: _e.mock.On("CheckSqlQuery", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_CheckSqlQuery_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest])) *MockConnectionServiceHandler_CheckSqlQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_CheckSqlQuery_Call) Return(_a0 *connect.Response[mgmtv1alpha1.CheckSqlQueryResponse], _a1 error) *MockConnectionServiceHandler_CheckSqlQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_CheckSqlQuery_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) (*connect.Response[mgmtv1alpha1.CheckSqlQueryResponse], error)) *MockConnectionServiceHandler_CheckSqlQuery_Call {
	_c.Call.Return(run)
	return _c
}

// CreateConnection provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) CreateConnection(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) (*connect.Response[mgmtv1alpha1.CreateConnectionResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateConnection")
	}

	var r0 *connect.Response[mgmtv1alpha1.CreateConnectionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) (*connect.Response[mgmtv1alpha1.CreateConnectionResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) *connect.Response[mgmtv1alpha1.CreateConnectionResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.CreateConnectionResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_CreateConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateConnection'
type MockConnectionServiceHandler_CreateConnection_Call struct {
	*mock.Call
}

// CreateConnection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.CreateConnectionRequest]
func (_e *MockConnectionServiceHandler_Expecter) CreateConnection(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_CreateConnection_Call {
	return &MockConnectionServiceHandler_CreateConnection_Call{Call: _e.mock.On("CreateConnection", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_CreateConnection_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CreateConnectionRequest])) *MockConnectionServiceHandler_CreateConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.CreateConnectionRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_CreateConnection_Call) Return(_a0 *connect.Response[mgmtv1alpha1.CreateConnectionResponse], _a1 error) *MockConnectionServiceHandler_CreateConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_CreateConnection_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) (*connect.Response[mgmtv1alpha1.CreateConnectionResponse], error)) *MockConnectionServiceHandler_CreateConnection_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteConnection provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) DeleteConnection(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) (*connect.Response[mgmtv1alpha1.DeleteConnectionResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteConnection")
	}

	var r0 *connect.Response[mgmtv1alpha1.DeleteConnectionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) (*connect.Response[mgmtv1alpha1.DeleteConnectionResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) *connect.Response[mgmtv1alpha1.DeleteConnectionResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.DeleteConnectionResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_DeleteConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteConnection'
type MockConnectionServiceHandler_DeleteConnection_Call struct {
	*mock.Call
}

// DeleteConnection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]
func (_e *MockConnectionServiceHandler_Expecter) DeleteConnection(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_DeleteConnection_Call {
	return &MockConnectionServiceHandler_DeleteConnection_Call{Call: _e.mock.On("DeleteConnection", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_DeleteConnection_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.DeleteConnectionRequest])) *MockConnectionServiceHandler_DeleteConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.DeleteConnectionRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_DeleteConnection_Call) Return(_a0 *connect.Response[mgmtv1alpha1.DeleteConnectionResponse], _a1 error) *MockConnectionServiceHandler_DeleteConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_DeleteConnection_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) (*connect.Response[mgmtv1alpha1.DeleteConnectionResponse], error)) *MockConnectionServiceHandler_DeleteConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnection provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) GetConnection(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetConnection")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetConnectionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionRequest]) *connect.Response[mgmtv1alpha1.GetConnectionResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetConnectionResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type MockConnectionServiceHandler_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionRequest]
func (_e *MockConnectionServiceHandler_Expecter) GetConnection(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_GetConnection_Call {
	return &MockConnectionServiceHandler_GetConnection_Call{Call: _e.mock.On("GetConnection", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_GetConnection_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionRequest])) *MockConnectionServiceHandler_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnection_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetConnectionResponse], _a1 error) *MockConnectionServiceHandler_GetConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnection_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionResponse], error)) *MockConnectionServiceHandler_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionDataStream provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockConnectionServiceHandler) GetConnectionDataStream(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest], _a2 *connect.ServerStream[mgmtv1alpha1.GetConnectionDataStreamResponse]) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionDataStream")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest], *connect.ServerStream[mgmtv1alpha1.GetConnectionDataStreamResponse]) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnectionServiceHandler_GetConnectionDataStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionDataStream'
type MockConnectionServiceHandler_GetConnectionDataStream_Call struct {
	*mock.Call
}

// GetConnectionDataStream is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]
//   - _a2 *connect.ServerStream[mgmtv1alpha1.GetConnectionDataStreamResponse]
func (_e *MockConnectionServiceHandler_Expecter) GetConnectionDataStream(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockConnectionServiceHandler_GetConnectionDataStream_Call {
	return &MockConnectionServiceHandler_GetConnectionDataStream_Call{Call: _e.mock.On("GetConnectionDataStream", _a0, _a1, _a2)}
}

func (_c *MockConnectionServiceHandler_GetConnectionDataStream_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest], _a2 *connect.ServerStream[mgmtv1alpha1.GetConnectionDataStreamResponse])) *MockConnectionServiceHandler_GetConnectionDataStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]), args[2].(*connect.ServerStream[mgmtv1alpha1.GetConnectionDataStreamResponse]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnectionDataStream_Call) Return(_a0 error) *MockConnectionServiceHandler_GetConnectionDataStream_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnectionDataStream_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest], *connect.ServerStream[mgmtv1alpha1.GetConnectionDataStreamResponse]) error) *MockConnectionServiceHandler_GetConnectionDataStream_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionForeignConstraints provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) GetConnectionForeignConstraints(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionForeignConstraints")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) *connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_GetConnectionForeignConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionForeignConstraints'
type MockConnectionServiceHandler_GetConnectionForeignConstraints_Call struct {
	*mock.Call
}

// GetConnectionForeignConstraints is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]
func (_e *MockConnectionServiceHandler_Expecter) GetConnectionForeignConstraints(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_GetConnectionForeignConstraints_Call {
	return &MockConnectionServiceHandler_GetConnectionForeignConstraints_Call{Call: _e.mock.On("GetConnectionForeignConstraints", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_GetConnectionForeignConstraints_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest])) *MockConnectionServiceHandler_GetConnectionForeignConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnectionForeignConstraints_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse], _a1 error) *MockConnectionServiceHandler_GetConnectionForeignConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnectionForeignConstraints_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse], error)) *MockConnectionServiceHandler_GetConnectionForeignConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionSchema provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) GetConnectionSchema(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionSchema")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) *connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_GetConnectionSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionSchema'
type MockConnectionServiceHandler_GetConnectionSchema_Call struct {
	*mock.Call
}

// GetConnectionSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]
func (_e *MockConnectionServiceHandler_Expecter) GetConnectionSchema(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_GetConnectionSchema_Call {
	return &MockConnectionServiceHandler_GetConnectionSchema_Call{Call: _e.mock.On("GetConnectionSchema", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_GetConnectionSchema_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest])) *MockConnectionServiceHandler_GetConnectionSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnectionSchema_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse], _a1 error) *MockConnectionServiceHandler_GetConnectionSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnectionSchema_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse], error)) *MockConnectionServiceHandler_GetConnectionSchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnections provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) GetConnections(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionsResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetConnections")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetConnectionsResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionsResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) *connect.Response[mgmtv1alpha1.GetConnectionsResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetConnectionsResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_GetConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnections'
type MockConnectionServiceHandler_GetConnections_Call struct {
	*mock.Call
}

// GetConnections is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionsRequest]
func (_e *MockConnectionServiceHandler_Expecter) GetConnections(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_GetConnections_Call {
	return &MockConnectionServiceHandler_GetConnections_Call{Call: _e.mock.On("GetConnections", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_GetConnections_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionsRequest])) *MockConnectionServiceHandler_GetConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionsRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnections_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetConnectionsResponse], _a1 error) *MockConnectionServiceHandler_GetConnections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_GetConnections_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionsResponse], error)) *MockConnectionServiceHandler_GetConnections_Call {
	_c.Call.Return(run)
	return _c
}

// IsConnectionNameAvailable provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) IsConnectionNameAvailable(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsConnectionNameAvailable")
	}

	var r0 *connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) *connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_IsConnectionNameAvailable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsConnectionNameAvailable'
type MockConnectionServiceHandler_IsConnectionNameAvailable_Call struct {
	*mock.Call
}

// IsConnectionNameAvailable is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]
func (_e *MockConnectionServiceHandler_Expecter) IsConnectionNameAvailable(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_IsConnectionNameAvailable_Call {
	return &MockConnectionServiceHandler_IsConnectionNameAvailable_Call{Call: _e.mock.On("IsConnectionNameAvailable", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_IsConnectionNameAvailable_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest])) *MockConnectionServiceHandler_IsConnectionNameAvailable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_IsConnectionNameAvailable_Call) Return(_a0 *connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse], _a1 error) *MockConnectionServiceHandler_IsConnectionNameAvailable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_IsConnectionNameAvailable_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse], error)) *MockConnectionServiceHandler_IsConnectionNameAvailable_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateConnection provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceHandler) UpdateConnection(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) (*connect.Response[mgmtv1alpha1.UpdateConnectionResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConnection")
	}

	var r0 *connect.Response[mgmtv1alpha1.UpdateConnectionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) (*connect.Response[mgmtv1alpha1.UpdateConnectionResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) *connect.Response[mgmtv1alpha1.UpdateConnectionResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.UpdateConnectionResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceHandler_UpdateConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateConnection'
type MockConnectionServiceHandler_UpdateConnection_Call struct {
	*mock.Call
}

// UpdateConnection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]
func (_e *MockConnectionServiceHandler_Expecter) UpdateConnection(_a0 interface{}, _a1 interface{}) *MockConnectionServiceHandler_UpdateConnection_Call {
	return &MockConnectionServiceHandler_UpdateConnection_Call{Call: _e.mock.On("UpdateConnection", _a0, _a1)}
}

func (_c *MockConnectionServiceHandler_UpdateConnection_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.UpdateConnectionRequest])) *MockConnectionServiceHandler_UpdateConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.UpdateConnectionRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceHandler_UpdateConnection_Call) Return(_a0 *connect.Response[mgmtv1alpha1.UpdateConnectionResponse], _a1 error) *MockConnectionServiceHandler_UpdateConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceHandler_UpdateConnection_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) (*connect.Response[mgmtv1alpha1.UpdateConnectionResponse], error)) *MockConnectionServiceHandler_UpdateConnection_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnectionServiceHandler creates a new instance of MockConnectionServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnectionServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnectionServiceHandler {
	mock := &MockConnectionServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
