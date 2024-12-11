package querybuilder2

import (
	sqlmanager_shared "github.com/nucleuscloud/neosync/backend/pkg/sqlmanager/shared"
	tabledependency "github.com/nucleuscloud/neosync/backend/pkg/table-dependency"
)

// QueryMapBuilderWrapper implements the SelectQueryMapBuilder interface
type QueryMapBuilderWrapper struct{}

type SelectQuery struct {
	Query string

	// If true, this query could return rows that violate foreign key constraints
	IsNotForeignKeySafeSubset bool
}

// BuildSelectQueryMap wraps the original BuildSelectQueryMap function
func (w *QueryMapBuilderWrapper) BuildSelectQueryMap(
	driver string,
	runConfigs []*tabledependency.RunConfig,
	subsetByForeignKeyConstraints bool,
	groupedColumnInfo map[string]map[string]*sqlmanager_shared.DatabaseSchemaRow,
) (map[string]map[tabledependency.RunType]*SelectQuery, error) {
	return BuildSelectQueryMap(
		driver,
		runConfigs,
		subsetByForeignKeyConstraints,
		groupedColumnInfo,
	)
}
