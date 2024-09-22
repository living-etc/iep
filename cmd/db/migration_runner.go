package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"iep/cmd/db/migrations"
)

type MigrationRunner struct{}

func (r *MigrationRunner) Run(
	ctx context.Context,
	db *sql.DB,
	migration migrations.Migration,
) error {
	fmt.Fprintf(os.Stdout, "\t%s\n", migration.Id)

	exec(ctx, db, migration.Statement, migration.Args...)

	add_migration_id_statement := fmt.Sprintf(
		"INSERT INTO migrations (id) values ('%s')",
		migration.Id,
	)
	exec(ctx, db, add_migration_id_statement)

	return nil
}
