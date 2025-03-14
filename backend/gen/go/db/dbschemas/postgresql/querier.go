// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package pg_queries

import (
	"context"
)

type Querier interface {
	GetAllSchemas(ctx context.Context, db DBTX) ([]string, error)
	GetAllTables(ctx context.Context, db DBTX) ([]*GetAllTablesRow, error)
	GetCustomFunctionsBySchemaAndTables(ctx context.Context, db DBTX, arg *GetCustomFunctionsBySchemaAndTablesParams) ([]*GetCustomFunctionsBySchemaAndTablesRow, error)
	GetCustomSequencesBySchemaAndTables(ctx context.Context, db DBTX, arg *GetCustomSequencesBySchemaAndTablesParams) ([]*GetCustomSequencesBySchemaAndTablesRow, error)
	GetCustomTriggersBySchemaAndTables(ctx context.Context, db DBTX, schematables []string) ([]*GetCustomTriggersBySchemaAndTablesRow, error)
	GetDataTypesBySchemaAndTables(ctx context.Context, db DBTX, arg *GetDataTypesBySchemaAndTablesParams) ([]*GetDataTypesBySchemaAndTablesRow, error)
	GetDatabaseSchema(ctx context.Context, db DBTX) ([]*GetDatabaseSchemaRow, error)
	GetDatabaseTableSchemasBySchemasAndTables(ctx context.Context, db DBTX, schematables []string) ([]*GetDatabaseTableSchemasBySchemasAndTablesRow, error)
	GetExtensionsBySchemas(ctx context.Context, db DBTX, schema []string) ([]*GetExtensionsBySchemasRow, error)
	GetForeignKeyConstraintsBySchemas(ctx context.Context, db DBTX, schemas []string) ([]*GetForeignKeyConstraintsBySchemasRow, error)
	GetIndicesBySchemasAndTables(ctx context.Context, db DBTX, schematables []string) ([]*GetIndicesBySchemasAndTablesRow, error)
	GetNonForeignKeyTableConstraintsBySchema(ctx context.Context, db DBTX, schemas []string) ([]*GetNonForeignKeyTableConstraintsBySchemaRow, error)
	GetPartitionHierarchyByTable(ctx context.Context, db DBTX, table string) ([]*GetPartitionHierarchyByTableRow, error)
	GetPartitionedTablesBySchema(ctx context.Context, db DBTX, schema []string) ([]*GetPartitionedTablesBySchemaRow, error)
	GetPostgresRolePermissions(ctx context.Context, db DBTX) ([]*GetPostgresRolePermissionsRow, error)
	GetUniqueIndexesBySchema(ctx context.Context, db DBTX, schema []string) ([]*GetUniqueIndexesBySchemaRow, error)
}

var _ Querier = (*Queries)(nil)
