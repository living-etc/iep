package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

type MigrationRunner struct{}

func (r *MigrationRunner) Run(
	ctx context.Context,
	db *sql.DB,
	migration Migration,
) error {
	fmt.Fprintf(os.Stdout, "Running migrations...\n")
	fmt.Fprintf(os.Stdout, "\t%s\n", migration.Id)
	fmt.Fprintf(os.Stdout, "Finished migrations\n")

	exec(ctx, db, string(migration.Statement))

	add_migration_id_statement := fmt.Sprintf(
		"INSERT INTO migrations (id) values ('%s')",
		migration.Id,
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
