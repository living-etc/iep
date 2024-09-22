package unapplied_migrations

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"db/internal/migrator"
)

const (
	getMigrationIdsSQL = "SELECT * FROM migrations"
)

func Get(
	ctx context.Context,
	db *sql.DB,
	migrationFilePaths []string,
) []migrator.Migration {
	unapplied_migrations := []migrator.Migration{}

	for _, migrationFilePath := range migrationFilePaths {
		migration_id := strings.TrimSuffix(
			filepath.Base(migrationFilePath),
			filepath.Ext(migrationFilePath),
		)

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
