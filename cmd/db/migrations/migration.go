package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

type Migration struct {
	Id        string
	Statement string
	Args      []any
}

func (m *Migration) Run(
	ctx context.Context,
	db *sql.DB,
) error {
	fmt.Fprintf(os.Stdout, "\t%s\n", m.Id)

	exec(ctx, db, m.Statement, m.Args...)

	add_migration_id_statement := fmt.Sprintf(
		"INSERT INTO migrations (id) values ('%s')",
		m.Id,
	)
	exec(ctx, db, add_migration_id_statement)

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
