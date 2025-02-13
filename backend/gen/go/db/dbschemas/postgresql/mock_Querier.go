// Code generated by mockery. DO NOT EDIT.

package pg_queries

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

type MockQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerier) EXPECT() *MockQuerier_Expecter {
	return &MockQuerier_Expecter{mock: &_m.Mock}
}

// GetCustomFunctionsBySchemaAndTables provides a mock function with given fields: ctx, db, arg
func (_m *MockQuerier) GetCustomFunctionsBySchemaAndTables(ctx context.Context, db DBTX, arg *GetCustomFunctionsBySchemaAndTablesParams) ([]*GetCustomFunctionsBySchemaAndTablesRow, error) {
	ret := _m.Called(ctx, db, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomFunctionsBySchemaAndTables")
	}

	var r0 []*GetCustomFunctionsBySchemaAndTablesRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetCustomFunctionsBySchemaAndTablesParams) ([]*GetCustomFunctionsBySchemaAndTablesRow, error)); ok {
		return rf(ctx, db, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetCustomFunctionsBySchemaAndTablesParams) []*GetCustomFunctionsBySchemaAndTablesRow); ok {
		r0 = rf(ctx, db, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetCustomFunctionsBySchemaAndTablesRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, *GetCustomFunctionsBySchemaAndTablesParams) error); ok {
		r1 = rf(ctx, db, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetCustomFunctionsBySchemaAndTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomFunctionsBySchemaAndTables'
type MockQuerier_GetCustomFunctionsBySchemaAndTables_Call struct {
	*mock.Call
}

// GetCustomFunctionsBySchemaAndTables is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - arg *GetCustomFunctionsBySchemaAndTablesParams
func (_e *MockQuerier_Expecter) GetCustomFunctionsBySchemaAndTables(ctx interface{}, db interface{}, arg interface{}) *MockQuerier_GetCustomFunctionsBySchemaAndTables_Call {
	return &MockQuerier_GetCustomFunctionsBySchemaAndTables_Call{Call: _e.mock.On("GetCustomFunctionsBySchemaAndTables", ctx, db, arg)}
}

func (_c *MockQuerier_GetCustomFunctionsBySchemaAndTables_Call) Run(run func(ctx context.Context, db DBTX, arg *GetCustomFunctionsBySchemaAndTablesParams)) *MockQuerier_GetCustomFunctionsBySchemaAndTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].(*GetCustomFunctionsBySchemaAndTablesParams))
	})
	return _c
}

func (_c *MockQuerier_GetCustomFunctionsBySchemaAndTables_Call) Return(_a0 []*GetCustomFunctionsBySchemaAndTablesRow, _a1 error) *MockQuerier_GetCustomFunctionsBySchemaAndTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetCustomFunctionsBySchemaAndTables_Call) RunAndReturn(run func(context.Context, DBTX, *GetCustomFunctionsBySchemaAndTablesParams) ([]*GetCustomFunctionsBySchemaAndTablesRow, error)) *MockQuerier_GetCustomFunctionsBySchemaAndTables_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomSequencesBySchemaAndTables provides a mock function with given fields: ctx, db, arg
func (_m *MockQuerier) GetCustomSequencesBySchemaAndTables(ctx context.Context, db DBTX, arg *GetCustomSequencesBySchemaAndTablesParams) ([]*GetCustomSequencesBySchemaAndTablesRow, error) {
	ret := _m.Called(ctx, db, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomSequencesBySchemaAndTables")
	}

	var r0 []*GetCustomSequencesBySchemaAndTablesRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetCustomSequencesBySchemaAndTablesParams) ([]*GetCustomSequencesBySchemaAndTablesRow, error)); ok {
		return rf(ctx, db, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetCustomSequencesBySchemaAndTablesParams) []*GetCustomSequencesBySchemaAndTablesRow); ok {
		r0 = rf(ctx, db, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetCustomSequencesBySchemaAndTablesRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, *GetCustomSequencesBySchemaAndTablesParams) error); ok {
		r1 = rf(ctx, db, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetCustomSequencesBySchemaAndTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomSequencesBySchemaAndTables'
type MockQuerier_GetCustomSequencesBySchemaAndTables_Call struct {
	*mock.Call
}

// GetCustomSequencesBySchemaAndTables is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - arg *GetCustomSequencesBySchemaAndTablesParams
func (_e *MockQuerier_Expecter) GetCustomSequencesBySchemaAndTables(ctx interface{}, db interface{}, arg interface{}) *MockQuerier_GetCustomSequencesBySchemaAndTables_Call {
	return &MockQuerier_GetCustomSequencesBySchemaAndTables_Call{Call: _e.mock.On("GetCustomSequencesBySchemaAndTables", ctx, db, arg)}
}

func (_c *MockQuerier_GetCustomSequencesBySchemaAndTables_Call) Run(run func(ctx context.Context, db DBTX, arg *GetCustomSequencesBySchemaAndTablesParams)) *MockQuerier_GetCustomSequencesBySchemaAndTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].(*GetCustomSequencesBySchemaAndTablesParams))
	})
	return _c
}

func (_c *MockQuerier_GetCustomSequencesBySchemaAndTables_Call) Return(_a0 []*GetCustomSequencesBySchemaAndTablesRow, _a1 error) *MockQuerier_GetCustomSequencesBySchemaAndTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetCustomSequencesBySchemaAndTables_Call) RunAndReturn(run func(context.Context, DBTX, *GetCustomSequencesBySchemaAndTablesParams) ([]*GetCustomSequencesBySchemaAndTablesRow, error)) *MockQuerier_GetCustomSequencesBySchemaAndTables_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomTriggersBySchemaAndTables provides a mock function with given fields: ctx, db, schematables
func (_m *MockQuerier) GetCustomTriggersBySchemaAndTables(ctx context.Context, db DBTX, schematables []string) ([]*GetCustomTriggersBySchemaAndTablesRow, error) {
	ret := _m.Called(ctx, db, schematables)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomTriggersBySchemaAndTables")
	}

	var r0 []*GetCustomTriggersBySchemaAndTablesRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) ([]*GetCustomTriggersBySchemaAndTablesRow, error)); ok {
		return rf(ctx, db, schematables)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) []*GetCustomTriggersBySchemaAndTablesRow); ok {
		r0 = rf(ctx, db, schematables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetCustomTriggersBySchemaAndTablesRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, []string) error); ok {
		r1 = rf(ctx, db, schematables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetCustomTriggersBySchemaAndTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomTriggersBySchemaAndTables'
type MockQuerier_GetCustomTriggersBySchemaAndTables_Call struct {
	*mock.Call
}

// GetCustomTriggersBySchemaAndTables is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - schematables []string
func (_e *MockQuerier_Expecter) GetCustomTriggersBySchemaAndTables(ctx interface{}, db interface{}, schematables interface{}) *MockQuerier_GetCustomTriggersBySchemaAndTables_Call {
	return &MockQuerier_GetCustomTriggersBySchemaAndTables_Call{Call: _e.mock.On("GetCustomTriggersBySchemaAndTables", ctx, db, schematables)}
}

func (_c *MockQuerier_GetCustomTriggersBySchemaAndTables_Call) Run(run func(ctx context.Context, db DBTX, schematables []string)) *MockQuerier_GetCustomTriggersBySchemaAndTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].([]string))
	})
	return _c
}

func (_c *MockQuerier_GetCustomTriggersBySchemaAndTables_Call) Return(_a0 []*GetCustomTriggersBySchemaAndTablesRow, _a1 error) *MockQuerier_GetCustomTriggersBySchemaAndTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetCustomTriggersBySchemaAndTables_Call) RunAndReturn(run func(context.Context, DBTX, []string) ([]*GetCustomTriggersBySchemaAndTablesRow, error)) *MockQuerier_GetCustomTriggersBySchemaAndTables_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataTypesBySchemaAndTables provides a mock function with given fields: ctx, db, arg
func (_m *MockQuerier) GetDataTypesBySchemaAndTables(ctx context.Context, db DBTX, arg *GetDataTypesBySchemaAndTablesParams) ([]*GetDataTypesBySchemaAndTablesRow, error) {
	ret := _m.Called(ctx, db, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetDataTypesBySchemaAndTables")
	}

	var r0 []*GetDataTypesBySchemaAndTablesRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetDataTypesBySchemaAndTablesParams) ([]*GetDataTypesBySchemaAndTablesRow, error)); ok {
		return rf(ctx, db, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetDataTypesBySchemaAndTablesParams) []*GetDataTypesBySchemaAndTablesRow); ok {
		r0 = rf(ctx, db, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetDataTypesBySchemaAndTablesRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, *GetDataTypesBySchemaAndTablesParams) error); ok {
		r1 = rf(ctx, db, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetDataTypesBySchemaAndTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataTypesBySchemaAndTables'
type MockQuerier_GetDataTypesBySchemaAndTables_Call struct {
	*mock.Call
}

// GetDataTypesBySchemaAndTables is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - arg *GetDataTypesBySchemaAndTablesParams
func (_e *MockQuerier_Expecter) GetDataTypesBySchemaAndTables(ctx interface{}, db interface{}, arg interface{}) *MockQuerier_GetDataTypesBySchemaAndTables_Call {
	return &MockQuerier_GetDataTypesBySchemaAndTables_Call{Call: _e.mock.On("GetDataTypesBySchemaAndTables", ctx, db, arg)}
}

func (_c *MockQuerier_GetDataTypesBySchemaAndTables_Call) Run(run func(ctx context.Context, db DBTX, arg *GetDataTypesBySchemaAndTablesParams)) *MockQuerier_GetDataTypesBySchemaAndTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].(*GetDataTypesBySchemaAndTablesParams))
	})
	return _c
}

func (_c *MockQuerier_GetDataTypesBySchemaAndTables_Call) Return(_a0 []*GetDataTypesBySchemaAndTablesRow, _a1 error) *MockQuerier_GetDataTypesBySchemaAndTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetDataTypesBySchemaAndTables_Call) RunAndReturn(run func(context.Context, DBTX, *GetDataTypesBySchemaAndTablesParams) ([]*GetDataTypesBySchemaAndTablesRow, error)) *MockQuerier_GetDataTypesBySchemaAndTables_Call {
	_c.Call.Return(run)
	return _c
}

// GetDatabaseSchema provides a mock function with given fields: ctx, db
func (_m *MockQuerier) GetDatabaseSchema(ctx context.Context, db DBTX) ([]*GetDatabaseSchemaRow, error) {
	ret := _m.Called(ctx, db)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabaseSchema")
	}

	var r0 []*GetDatabaseSchemaRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX) ([]*GetDatabaseSchemaRow, error)); ok {
		return rf(ctx, db)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX) []*GetDatabaseSchemaRow); ok {
		r0 = rf(ctx, db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetDatabaseSchemaRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX) error); ok {
		r1 = rf(ctx, db)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetDatabaseSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseSchema'
type MockQuerier_GetDatabaseSchema_Call struct {
	*mock.Call
}

// GetDatabaseSchema is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
func (_e *MockQuerier_Expecter) GetDatabaseSchema(ctx interface{}, db interface{}) *MockQuerier_GetDatabaseSchema_Call {
	return &MockQuerier_GetDatabaseSchema_Call{Call: _e.mock.On("GetDatabaseSchema", ctx, db)}
}

func (_c *MockQuerier_GetDatabaseSchema_Call) Run(run func(ctx context.Context, db DBTX)) *MockQuerier_GetDatabaseSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX))
	})
	return _c
}

func (_c *MockQuerier_GetDatabaseSchema_Call) Return(_a0 []*GetDatabaseSchemaRow, _a1 error) *MockQuerier_GetDatabaseSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetDatabaseSchema_Call) RunAndReturn(run func(context.Context, DBTX) ([]*GetDatabaseSchemaRow, error)) *MockQuerier_GetDatabaseSchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetDatabaseTableSchemasBySchemasAndTables provides a mock function with given fields: ctx, db, schematables
func (_m *MockQuerier) GetDatabaseTableSchemasBySchemasAndTables(ctx context.Context, db DBTX, schematables []string) ([]*GetDatabaseTableSchemasBySchemasAndTablesRow, error) {
	ret := _m.Called(ctx, db, schematables)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabaseTableSchemasBySchemasAndTables")
	}

	var r0 []*GetDatabaseTableSchemasBySchemasAndTablesRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) ([]*GetDatabaseTableSchemasBySchemasAndTablesRow, error)); ok {
		return rf(ctx, db, schematables)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) []*GetDatabaseTableSchemasBySchemasAndTablesRow); ok {
		r0 = rf(ctx, db, schematables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetDatabaseTableSchemasBySchemasAndTablesRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, []string) error); ok {
		r1 = rf(ctx, db, schematables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseTableSchemasBySchemasAndTables'
type MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call struct {
	*mock.Call
}

// GetDatabaseTableSchemasBySchemasAndTables is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - schematables []string
func (_e *MockQuerier_Expecter) GetDatabaseTableSchemasBySchemasAndTables(ctx interface{}, db interface{}, schematables interface{}) *MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call {
	return &MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call{Call: _e.mock.On("GetDatabaseTableSchemasBySchemasAndTables", ctx, db, schematables)}
}

func (_c *MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call) Run(run func(ctx context.Context, db DBTX, schematables []string)) *MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].([]string))
	})
	return _c
}

func (_c *MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call) Return(_a0 []*GetDatabaseTableSchemasBySchemasAndTablesRow, _a1 error) *MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call) RunAndReturn(run func(context.Context, DBTX, []string) ([]*GetDatabaseTableSchemasBySchemasAndTablesRow, error)) *MockQuerier_GetDatabaseTableSchemasBySchemasAndTables_Call {
	_c.Call.Return(run)
	return _c
}

// GetExtensions provides a mock function with given fields: ctx, db
func (_m *MockQuerier) GetExtensions(ctx context.Context, db DBTX) ([]*GetExtensionsRow, error) {
	ret := _m.Called(ctx, db)

	if len(ret) == 0 {
		panic("no return value specified for GetExtensions")
	}

	var r0 []*GetExtensionsRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX) ([]*GetExtensionsRow, error)); ok {
		return rf(ctx, db)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX) []*GetExtensionsRow); ok {
		r0 = rf(ctx, db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetExtensionsRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX) error); ok {
		r1 = rf(ctx, db)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetExtensions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExtensions'
type MockQuerier_GetExtensions_Call struct {
	*mock.Call
}

// GetExtensions is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
func (_e *MockQuerier_Expecter) GetExtensions(ctx interface{}, db interface{}) *MockQuerier_GetExtensions_Call {
	return &MockQuerier_GetExtensions_Call{Call: _e.mock.On("GetExtensions", ctx, db)}
}

func (_c *MockQuerier_GetExtensions_Call) Run(run func(ctx context.Context, db DBTX)) *MockQuerier_GetExtensions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX))
	})
	return _c
}

func (_c *MockQuerier_GetExtensions_Call) Return(_a0 []*GetExtensionsRow, _a1 error) *MockQuerier_GetExtensions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetExtensions_Call) RunAndReturn(run func(context.Context, DBTX) ([]*GetExtensionsRow, error)) *MockQuerier_GetExtensions_Call {
	_c.Call.Return(run)
	return _c
}

// GetForeignKeyConstraintsBySchemas provides a mock function with given fields: ctx, db, schemas
func (_m *MockQuerier) GetForeignKeyConstraintsBySchemas(ctx context.Context, db DBTX, schemas []string) ([]*GetForeignKeyConstraintsBySchemasRow, error) {
	ret := _m.Called(ctx, db, schemas)

	if len(ret) == 0 {
		panic("no return value specified for GetForeignKeyConstraintsBySchemas")
	}

	var r0 []*GetForeignKeyConstraintsBySchemasRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) ([]*GetForeignKeyConstraintsBySchemasRow, error)); ok {
		return rf(ctx, db, schemas)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) []*GetForeignKeyConstraintsBySchemasRow); ok {
		r0 = rf(ctx, db, schemas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetForeignKeyConstraintsBySchemasRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, []string) error); ok {
		r1 = rf(ctx, db, schemas)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetForeignKeyConstraintsBySchemas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForeignKeyConstraintsBySchemas'
type MockQuerier_GetForeignKeyConstraintsBySchemas_Call struct {
	*mock.Call
}

// GetForeignKeyConstraintsBySchemas is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - schemas []string
func (_e *MockQuerier_Expecter) GetForeignKeyConstraintsBySchemas(ctx interface{}, db interface{}, schemas interface{}) *MockQuerier_GetForeignKeyConstraintsBySchemas_Call {
	return &MockQuerier_GetForeignKeyConstraintsBySchemas_Call{Call: _e.mock.On("GetForeignKeyConstraintsBySchemas", ctx, db, schemas)}
}

func (_c *MockQuerier_GetForeignKeyConstraintsBySchemas_Call) Run(run func(ctx context.Context, db DBTX, schemas []string)) *MockQuerier_GetForeignKeyConstraintsBySchemas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].([]string))
	})
	return _c
}

func (_c *MockQuerier_GetForeignKeyConstraintsBySchemas_Call) Return(_a0 []*GetForeignKeyConstraintsBySchemasRow, _a1 error) *MockQuerier_GetForeignKeyConstraintsBySchemas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetForeignKeyConstraintsBySchemas_Call) RunAndReturn(run func(context.Context, DBTX, []string) ([]*GetForeignKeyConstraintsBySchemasRow, error)) *MockQuerier_GetForeignKeyConstraintsBySchemas_Call {
	_c.Call.Return(run)
	return _c
}

// GetIndicesBySchemasAndTables provides a mock function with given fields: ctx, db, schematables
func (_m *MockQuerier) GetIndicesBySchemasAndTables(ctx context.Context, db DBTX, schematables []string) ([]*GetIndicesBySchemasAndTablesRow, error) {
	ret := _m.Called(ctx, db, schematables)

	if len(ret) == 0 {
		panic("no return value specified for GetIndicesBySchemasAndTables")
	}

	var r0 []*GetIndicesBySchemasAndTablesRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) ([]*GetIndicesBySchemasAndTablesRow, error)); ok {
		return rf(ctx, db, schematables)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) []*GetIndicesBySchemasAndTablesRow); ok {
		r0 = rf(ctx, db, schematables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetIndicesBySchemasAndTablesRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, []string) error); ok {
		r1 = rf(ctx, db, schematables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetIndicesBySchemasAndTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndicesBySchemasAndTables'
type MockQuerier_GetIndicesBySchemasAndTables_Call struct {
	*mock.Call
}

// GetIndicesBySchemasAndTables is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - schematables []string
func (_e *MockQuerier_Expecter) GetIndicesBySchemasAndTables(ctx interface{}, db interface{}, schematables interface{}) *MockQuerier_GetIndicesBySchemasAndTables_Call {
	return &MockQuerier_GetIndicesBySchemasAndTables_Call{Call: _e.mock.On("GetIndicesBySchemasAndTables", ctx, db, schematables)}
}

func (_c *MockQuerier_GetIndicesBySchemasAndTables_Call) Run(run func(ctx context.Context, db DBTX, schematables []string)) *MockQuerier_GetIndicesBySchemasAndTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].([]string))
	})
	return _c
}

func (_c *MockQuerier_GetIndicesBySchemasAndTables_Call) Return(_a0 []*GetIndicesBySchemasAndTablesRow, _a1 error) *MockQuerier_GetIndicesBySchemasAndTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetIndicesBySchemasAndTables_Call) RunAndReturn(run func(context.Context, DBTX, []string) ([]*GetIndicesBySchemasAndTablesRow, error)) *MockQuerier_GetIndicesBySchemasAndTables_Call {
	_c.Call.Return(run)
	return _c
}

// GetNonForeignKeyTableConstraintsBySchema provides a mock function with given fields: ctx, db, schema
func (_m *MockQuerier) GetNonForeignKeyTableConstraintsBySchema(ctx context.Context, db DBTX, schema []string) ([]*GetNonForeignKeyTableConstraintsBySchemaRow, error) {
	ret := _m.Called(ctx, db, schema)

	if len(ret) == 0 {
		panic("no return value specified for GetNonForeignKeyTableConstraintsBySchema")
	}

	var r0 []*GetNonForeignKeyTableConstraintsBySchemaRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) ([]*GetNonForeignKeyTableConstraintsBySchemaRow, error)); ok {
		return rf(ctx, db, schema)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, []string) []*GetNonForeignKeyTableConstraintsBySchemaRow); ok {
		r0 = rf(ctx, db, schema)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetNonForeignKeyTableConstraintsBySchemaRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, []string) error); ok {
		r1 = rf(ctx, db, schema)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNonForeignKeyTableConstraintsBySchema'
type MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call struct {
	*mock.Call
}

// GetNonForeignKeyTableConstraintsBySchema is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - schema []string
func (_e *MockQuerier_Expecter) GetNonForeignKeyTableConstraintsBySchema(ctx interface{}, db interface{}, schema interface{}) *MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call {
	return &MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call{Call: _e.mock.On("GetNonForeignKeyTableConstraintsBySchema", ctx, db, schema)}
}

func (_c *MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call) Run(run func(ctx context.Context, db DBTX, schema []string)) *MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].([]string))
	})
	return _c
}

func (_c *MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call) Return(_a0 []*GetNonForeignKeyTableConstraintsBySchemaRow, _a1 error) *MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call) RunAndReturn(run func(context.Context, DBTX, []string) ([]*GetNonForeignKeyTableConstraintsBySchemaRow, error)) *MockQuerier_GetNonForeignKeyTableConstraintsBySchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetPostgresRolePermissions provides a mock function with given fields: ctx, db
func (_m *MockQuerier) GetPostgresRolePermissions(ctx context.Context, db DBTX) ([]*GetPostgresRolePermissionsRow, error) {
	ret := _m.Called(ctx, db)

	if len(ret) == 0 {
		panic("no return value specified for GetPostgresRolePermissions")
	}

	var r0 []*GetPostgresRolePermissionsRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX) ([]*GetPostgresRolePermissionsRow, error)); ok {
		return rf(ctx, db)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX) []*GetPostgresRolePermissionsRow); ok {
		r0 = rf(ctx, db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetPostgresRolePermissionsRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX) error); ok {
		r1 = rf(ctx, db)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetPostgresRolePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPostgresRolePermissions'
type MockQuerier_GetPostgresRolePermissions_Call struct {
	*mock.Call
}

// GetPostgresRolePermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
func (_e *MockQuerier_Expecter) GetPostgresRolePermissions(ctx interface{}, db interface{}) *MockQuerier_GetPostgresRolePermissions_Call {
	return &MockQuerier_GetPostgresRolePermissions_Call{Call: _e.mock.On("GetPostgresRolePermissions", ctx, db)}
}

func (_c *MockQuerier_GetPostgresRolePermissions_Call) Run(run func(ctx context.Context, db DBTX)) *MockQuerier_GetPostgresRolePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX))
	})
	return _c
}

func (_c *MockQuerier_GetPostgresRolePermissions_Call) Return(_a0 []*GetPostgresRolePermissionsRow, _a1 error) *MockQuerier_GetPostgresRolePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetPostgresRolePermissions_Call) RunAndReturn(run func(context.Context, DBTX) ([]*GetPostgresRolePermissionsRow, error)) *MockQuerier_GetPostgresRolePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
