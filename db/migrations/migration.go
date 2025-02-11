package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

const MARKDOWN_PATH = "/Users/chris/Code/personal/infrastructure-exercism-prototype/db/migrations/markdown"

type Migration struct {
	Id        string
	Statement string
	Args      []any
}

func (m *Migration) Run(
	ctx context.Context,
	db *sql.DB,
	logger *log.Logger,
) error {
	exec(ctx, db, m.Statement, m.Args...)
	exec(ctx, db, "INSERT INTO migrations (id) values (?)", m.Id)

	return nil
}

func exec(ctx context.Context, db *sql.DB, statement string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}
