package internals

import (
	"context"
	"database/sql"
)

type DB interface {
	QueryRow(string, ...any) *sql.Row
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	PingContext(context.Context) error
	Close() error
}
