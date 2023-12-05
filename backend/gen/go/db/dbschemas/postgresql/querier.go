// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package pg_queries

import (
	"context"
)

type Querier interface {
	GetDatabaseSchema(ctx context.Context, db DBTX) ([]*GetDatabaseSchemaRow, error)
	GetDatabaseTableSchema(ctx context.Context, db DBTX, arg *GetDatabaseTableSchemaParams) ([]*GetDatabaseTableSchemaRow, error)
	GetForeignKeyConstraints(ctx context.Context, db DBTX, tableschema string) ([]*GetForeignKeyConstraintsRow, error)
	GetTableConstraints(ctx context.Context, db DBTX, arg *GetTableConstraintsParams) ([]*GetTableConstraintsRow, error)
}

var _ Querier = (*Queries)(nil)
