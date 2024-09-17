package internals_test

import (
	"context"
	"database/sql"
)

type MockDB struct{}

func (db *MockDB) QueryRow(statement string, args ...any) *sql.Row {
	return &sql.Row{}
}

func (db *MockDB) ExecContext(
	ctx context.Context,
	statement string,
	args ...any,
) (sql.Result, error) {
	return MockResult{}, nil
}

func (db *MockDB) QueryContext(context.Context, string, ...any) (*sql.Rows, error) {
	return &sql.Rows{}, nil
}
func (db *MockDB) PingContext(context.Context) error { return nil }
func (db *MockDB) Close() error                      { return nil }
