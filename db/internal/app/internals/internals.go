package internals

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

const (
	createMigrationsTableSQL = "CREATE TABLE IF NOT EXISTS migrations(id TEXT NOT NULL, PRIMARY KEY(id))"
	addMigrationIdSQL        = "INSERT INTO migrations (id) values ('%s')"
	getMigrationIdsSQL       = "SELECT * FROM migrations"
	dbName                   = "file:exercises.db"
	addExerciseSQL           = "INSERT INTO exercises(id, name, description, body) VALUES(?, ?, ?, ?)"
	getExerciseByIdSQL       = "SELECT COUNT(id) FROM exercises WHERE id = ? ORDER BY id desc LIMIT 1"
)

func exec(ctx context.Context, db *sql.DB, statement string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}

func query(ctx context.Context, db *sql.DB, statement string, args ...any) *sql.Rows {
	res, err := db.QueryContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to query statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}

func openDb() *sql.DB {
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db: %s", err)
		os.Exit(1)
	}

	return db
}

func InitDb() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ping db: %s", err)
	}

	exec(ctx, db, createMigrationsTableSQL)
}

func UnappliedMigrations(
	ctx context.Context,
	db *sql.DB,
	migrationFilePaths []string,
) []Migration {
	unapplied_migrations := []Migration{}

	for _, migrationFilePath := range migrationFilePaths {
		if !migration_completed(
			strings.TrimSuffix(filepath.Base(migrationFilePath), filepath.Ext(migrationFilePath)),
			completed_migrations(ctx, db),
		) {
			unapplied_migrations = append(
				unapplied_migrations,
				Migration{Filepath: migrationFilePath},
			)
		}
	}

	return unapplied_migrations
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

func migration_completed(migration_file string, completed_migrations []string) bool {
	for _, completed_migration := range completed_migrations {
		if completed_migration == migration_file {
			return true
		}
	}

	return false
}
