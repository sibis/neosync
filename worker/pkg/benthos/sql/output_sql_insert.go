package neosync_benthos_sql

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	mysql_queries "github.com/nucleuscloud/neosync/backend/gen/go/db/dbschemas/mysql"
	neosync_benthos "github.com/nucleuscloud/neosync/worker/pkg/benthos"
	querybuilder "github.com/nucleuscloud/neosync/worker/pkg/query-builder"
	"github.com/redpanda-data/benthos/v4/public/service"
)

func sqlInsertOutputSpec() *service.ConfigSpec {
	return service.NewConfigSpec().
		Field(service.NewStringField("connection_id")).
		Field(service.NewStringField("schema")).
		Field(service.NewStringField("table")).
		Field(service.NewStringListField("primary_key_columns")).
		Field(service.NewBoolField("on_conflict_do_nothing").Optional().Default(false)).
		Field(service.NewBoolField("on_conflict_do_update").Optional().Default(false)).
		Field(service.NewBoolField("skip_foreign_key_violations").Optional().Default(false)).
		Field(service.NewBoolField("truncate_on_retry").Optional().Default(false).Deprecated()).
		Field(service.NewBoolField("should_override_column_default").Optional().Default(false)).
		Field(service.NewIntField("max_in_flight").Default(64)).
		Field(service.NewBatchPolicyField("batching")).
		Field(service.NewStringField("prefix").Optional()).
		Field(service.NewStringField("suffix").Optional())
}

// Registers an output on a benthos environment called pooled_sql_raw
func RegisterPooledSqlInsertOutput(env *service.Environment, dbprovider ConnectionProvider, isRetry bool, logger *slog.Logger) error {
	return env.RegisterBatchOutput(
		"pooled_sql_insert", sqlInsertOutputSpec(),
		func(conf *service.ParsedConfig, mgr *service.Resources) (service.BatchOutput, service.BatchPolicy, int, error) {
			batchPolicy, err := conf.FieldBatchPolicy("batching")
			if err != nil {
				return nil, batchPolicy, -1, err
			}

			maxInFlight, err := conf.FieldInt("max_in_flight")
			if err != nil {
				return nil, service.BatchPolicy{}, -1, err
			}
			out, err := newInsertOutput(conf, mgr, dbprovider, isRetry, logger)
			if err != nil {
				return nil, service.BatchPolicy{}, -1, err
			}
			return out, batchPolicy, maxInFlight, nil
		},
	)
}

var _ service.BatchOutput = &pooledInsertOutput{}

type pooledInsertOutput struct {
	driver       string
	connectionId string
	provider     ConnectionProvider
	dbMut        sync.RWMutex
	db           mysql_queries.DBTX
	logger       *service.Logger
	slogger      *slog.Logger

	schema                   string
	table                    string
	onConflictDoNothing      bool
	skipForeignKeyViolations bool
	queryBuilder             querybuilder.InsertQueryBuilder
	isRetry                  bool
}

func newInsertOutput(conf *service.ParsedConfig, mgr *service.Resources, provider ConnectionProvider, isRetry bool, logger *slog.Logger) (*pooledInsertOutput, error) {
	connectionId, err := conf.FieldString("connection_id")
	if err != nil {
		return nil, err
	}

	schema, err := conf.FieldString("schema")
	if err != nil {
		return nil, err
	}

	table, err := conf.FieldString("table")
	if err != nil {
		return nil, err
	}

	primaryKeyColumns, err := conf.FieldStringList("primary_key_columns")
	if err != nil {
		return nil, err
	}

	onConflictDoNothing, err := conf.FieldBool("on_conflict_do_nothing")
	if err != nil {
		return nil, err
	}

	onConflictDoUpdate, err := conf.FieldBool("on_conflict_do_update")
	if err != nil {
		return nil, err
	}

	skipForeignKeyViolations, err := conf.FieldBool("skip_foreign_key_violations")
	if err != nil {
		return nil, err
	}

	shouldOverrideColumnDefault, err := conf.FieldBool("should_override_column_default")
	if err != nil {
		return nil, err
	}

	var prefix *string
	if conf.Contains("prefix") {
		prefixStr, err := conf.FieldString("prefix")
		if err != nil {
			return nil, err
		}
		prefix = &prefixStr
	}

	var suffix *string
	if conf.Contains("suffix") {
		suffixStr, err := conf.FieldString("suffix")
		if err != nil {
			return nil, err
		}
		suffix = &suffixStr
	}

	driver, err := provider.GetDriver(connectionId)
	if err != nil {
		return nil, err
	}

	options := []querybuilder.InsertOption{
		querybuilder.WithPrefix(prefix),
		querybuilder.WithSuffix(suffix),
	}

	if onConflictDoUpdate {
		options = append(options, querybuilder.WithOnConflictDoUpdate(primaryKeyColumns))
	} else if onConflictDoNothing || isRetry {
		options = append(options, querybuilder.WithOnConflictDoNothing())
	}
	if shouldOverrideColumnDefault {
		options = append(options, querybuilder.WithShouldOverrideColumnDefault())
	}
	builder, err := querybuilder.GetInsertBuilder(
		logger,
		driver,
		schema,
		table,
		options...,
	)
	if err != nil {
		return nil, err
	}

	output := &pooledInsertOutput{
		connectionId:             connectionId,
		driver:                   driver,
		logger:                   mgr.Logger(),
		slogger:                  logger,
		provider:                 provider,
		schema:                   schema,
		table:                    table,
		onConflictDoNothing:      onConflictDoNothing,
		queryBuilder:             builder,
		skipForeignKeyViolations: skipForeignKeyViolations,
		isRetry:                  isRetry,
	}
	return output, nil
}

func (s *pooledInsertOutput) Connect(ctx context.Context) error {
	s.dbMut.Lock()
	defer s.dbMut.Unlock()

	if s.db != nil {
		return nil
	}

	db, err := s.provider.GetDb(ctx, s.connectionId)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *pooledInsertOutput) WriteBatch(ctx context.Context, batch service.MessageBatch) error {
	s.dbMut.RLock()
	if s.db == nil {
		s.dbMut.RUnlock()
		return service.ErrNotConnected
	}
	db := s.db
	defer s.dbMut.RUnlock()

	batchLen := len(batch)
	if batchLen == 0 {
		return nil
	}

	rows := []map[string]any{}
	for _, msg := range batch {
		m, _ := msg.AsStructured()
		msgMap, ok := m.(map[string]any)
		if !ok {
			return fmt.Errorf("message returned non-map result: %T", msgMap)
		}

		rows = append(rows, msgMap)
	}
	if len(rows) == 0 {
		s.logger.Debug("no rows to insert")
		return nil
	}

	insertQuery, args, err := s.queryBuilder.BuildInsertQuery(rows)
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}
	if _, err := db.ExecContext(ctx, insertQuery, args...); err != nil {
		shouldRetry := neosync_benthos.ShouldRetryInsert(err.Error(), s.skipForeignKeyViolations)
		if !shouldRetry {
			return fmt.Errorf("failed to execute insert query: %w", err)
		}
		s.logger.Infof("received error during batch write that is retryable, proceeding with row by row insert: %s", err.Error())

		err = retryInsertRowByRow(ctx, db, s.queryBuilder, rows, s.skipForeignKeyViolations, s.logger)
		if err != nil {
			return fmt.Errorf("failed to retry insert query: %w", err)
		}
	}
	return nil
}

func retryInsertRowByRow(
	ctx context.Context,
	db mysql_queries.DBTX,
	builder querybuilder.InsertQueryBuilder,
	rows []map[string]any,
	skipForeignKeyViolations bool,
	logger *service.Logger,
) error {
	fkErrorCount := 0
	otherErrorCount := 0
	insertCount := 0
	for _, row := range rows {
		insertQuery, args, err := builder.BuildInsertQuery([]map[string]any{row})
		if err != nil {
			return err
		}
		_, err = db.ExecContext(ctx, insertQuery, args...)
		if err != nil {
			if !neosync_benthos.ShouldRetryInsert(err.Error(), skipForeignKeyViolations) {
				return fmt.Errorf("failed to retry insert query: %w", err)
			} else if neosync_benthos.IsForeignKeyViolationError(err.Error()) {
				fkErrorCount++
			} else {
				otherErrorCount++
				logger.Warnf("received retryable error during row by row insert. skipping row: %s", err.Error())
			}
		}
		if err == nil {
			insertCount++
		}
	}
	logger.Infof("Completed row-by-row insert with %d foreign key violations. Total Skipped rows: %d, Successfully inserted: %d", fkErrorCount, otherErrorCount, insertCount)
	return nil
}

func (s *pooledInsertOutput) Close(ctx context.Context) error {
	s.dbMut.RLock()
	if s.db == nil {
		s.dbMut.RUnlock()
		return nil
	}
	s.dbMut.RUnlock()

	// Take write lock to null out the connection
	s.dbMut.Lock()
	s.db = nil
	s.dbMut.Unlock()
	return nil
}
