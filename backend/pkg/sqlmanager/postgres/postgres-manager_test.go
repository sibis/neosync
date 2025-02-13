package sqlmanager_postgres

import (
	"testing"

	sqlmanager_shared "github.com/nucleuscloud/neosync/backend/pkg/sqlmanager/shared"
	"github.com/stretchr/testify/require"
)

func Test_EscapePgColumns(t *testing.T) {
	require.Empty(t, EscapePgColumns(nil))
	require.Equal(
		t,
		EscapePgColumns([]string{"foo", "bar", "baz"}),
		[]string{`"foo"`, `"bar"`, `"baz"`},
	)
}

func Test_BuildPgTruncateStatement(t *testing.T) {
	stmt, err := BuildPgTruncateStatement([]*sqlmanager_shared.SchemaTable{{Schema: "public", Table: "users"}, {Schema: "bad name", Table: "C$@111"}})
	require.NoError(t, err)
	require.Equal(
		t,
		"TRUNCATE \"public\".\"users\", \"bad name\".\"C$@111\" RESTART IDENTITY;",
		stmt,
	)
}

func Test_BuildPgTruncateCascadeStatement(t *testing.T) {
	actual, err := BuildPgTruncateCascadeStatement("public", "users")
	require.NoError(t, err)
	require.Equal(
		t,
		"TRUNCATE \"public\".\"users\" RESTART IDENTITY CASCADE;",
		actual,
	)
}

func Test_SequenceConfiguration(t *testing.T) {
	s := SequenceConfiguration{
		IncrementBy: 1,
		MinValue:    2,
		MaxValue:    5,
		StartValue:  2,
		CacheValue:  1,
		CycleOption: false,
	}
	require.Equal(
		t,
		"GENERATED BY DEFAULT AS IDENTITY ( INCREMENT BY 1 MINVALUE 2 MAXVALUE 5 START 2 CACHE 1 NO CYCLE )",
		s.ToGeneratedDefaultIdentity(),
	)
	require.Equal(
		t,
		"GENERATED ALWAYS AS IDENTITY ( INCREMENT BY 1 MINVALUE 2 MAXVALUE 5 START 2 CACHE 1 NO CYCLE )",
		s.ToGeneratedAlwaysIdentity(),
	)
}

func Test_BuildPgInsertIdentityAlwaysSql(t *testing.T) {
	t.Run("Basic insert query", func(t *testing.T) {
		input := "INSERT INTO users (id, name, email) VALUES (1, 'John', 'john@example.com')"
		expected := "INSERT INTO users (id, name, email) OVERRIDING SYSTEM VALUE VALUES(1, 'John', 'john@example.com')"
		result := BuildPgInsertIdentityAlwaysSql(input)
		require.Equal(t, expected, result)
	})

	t.Run("Insert query with multiple values", func(t *testing.T) {
		input := "INSERT INTO products (id, name, price) VALUES (1, 'Apple', 0.99), (2, 'Orange', 1.29)"
		expected := "INSERT INTO products (id, name, price) OVERRIDING SYSTEM VALUE VALUES(1, 'Apple', 0.99), (2, 'Orange', 1.29)"
		result := BuildPgInsertIdentityAlwaysSql(input)
		require.Equal(t, expected, result)
	})

	t.Run("Insert query with quoted values", func(t *testing.T) {
		input := "INSERT INTO quotes (id, text) VALUES (1, 'It''s a quote')"
		expected := "INSERT INTO quotes (id, text) OVERRIDING SYSTEM VALUE VALUES(1, 'It''s a quote')"
		result := BuildPgInsertIdentityAlwaysSql(input)
		require.Equal(t, expected, result)
	})

	t.Run("Insert query with parentheses in values", func(t *testing.T) {
		input := "INSERT INTO data (id, value) VALUES (1, '(nested) (parentheses)')"
		expected := "INSERT INTO data (id, value) OVERRIDING SYSTEM VALUE VALUES(1, '(nested) (parentheses)')"
		result := BuildPgInsertIdentityAlwaysSql(input)
		require.Equal(t, expected, result)
	})
}
