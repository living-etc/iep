package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sort"

	"iep/cmd/db/migrations"
)

const (
	getMigrationIdsSQL = "SELECT * FROM migrations"
)

func Get(
	ctx context.Context,
	db *sql.DB,
) []migrations.Migration {
	unapplied_migrations := []migrations.Migration{}

	migration_ids := []string{}
	for k := range MigrationFunctionRegistry {
		migration_ids = append(migration_ids, k)
	}
	sort.Strings(migration_ids)

	for _, migration_id := range migration_ids {
		if !migration_completed(
			migration_id,
			completed_migrations(ctx, db),
		) {
			unapplied_migrations = append(
				unapplied_migrations,
				MigrationFunctionRegistry[migration_id](),
			)
		}
	}

	return unapplied_migrations
}

func migration_completed(migration_file string, completed_migrations []string) bool {
	for _, completed_migration := range completed_migrations {
		if completed_migration == migration_file {
			return true
		}
	}

	return false
}

func completed_migrations(ctx context.Context, db *sql.DB) []string {
	completed_migration_rows := query(ctx, db, getMigrationIdsSQL)
	defer completed_migration_rows.Close()

	completed_migrations := []string{}
	for completed_migration_rows.Next() {
		var migration string

		if err := completed_migration_rows.Scan(&migration); err != nil {
			fmt.Fprintf(os.Stderr, "Scan err: %s", err)
			os.Exit(1)
		}

		completed_migrations = append(completed_migrations, migration)
	}
	if err := completed_migration_rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Row err: %s", err)
		os.Exit(1)
	}

	return completed_migrations
}

func query(ctx context.Context, db *sql.DB, statement string, args ...any) *sql.Rows {
	res, err := db.QueryContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to query statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}
