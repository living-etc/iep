package internals

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	createMigrationsTableSQL = "CREATE TABLE IF NOT EXISTS migrations(id TEXT NOT NULL, PRIMARY KEY(id))"
	addMigrationIdSQL        = "INSERT INTO migrations (id) values ('%s')"
	getMigrationIdsSQL       = "SELECT * FROM migrations"
	dbName                   = "file:exercises.db"
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

func Init_db() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ping db: %s", err)
	}

	exec(ctx, db, createMigrationsTableSQL)
}

func Migrate() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	migration_files := migration_files()
	completed_migrations := completed_migrations(ctx, db)
	unapplied_migrations := Unapplied_migrations(migration_files, completed_migrations)

	if len(unapplied_migrations) == 0 {
		fmt.Fprintf(os.Stdout, "No migrations to run\n")
		os.Exit(0)
	}

	for _, file := range unapplied_migrations {
		statement, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Running migrations...\n")
		migration_id := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		fmt.Fprintf(os.Stdout, "\t%s\n", migration_id)
		fmt.Fprintf(os.Stdout, "Finished migrations\n")

		exec(ctx, db, string(statement))

		add_migration_id_statement := fmt.Sprintf(addMigrationIdSQL, migration_id)
		exec(ctx, db, add_migration_id_statement)
	}
}

func migration_files() []string {
	migration_files, err := filepath.Glob("./migrations/*.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	return migration_files
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

func Unapplied_migrations(migration_files []string, completed_migrations []string) []string {
	unapplied_migrations := []string{}

	for _, migration_file := range migration_files {
		if !migration_completed(
			strings.TrimSuffix(filepath.Base(migration_file), filepath.Ext(migration_file)),
			completed_migrations,
		) {
			unapplied_migrations = append(unapplied_migrations, migration_file)
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
