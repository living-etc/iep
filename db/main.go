package main

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

func main() {
	args := os.Args[1:]

	subcommand := args[0]

	switch subcommand {
	case "init":
		init_db()
	case "migrate":
		migrate()
	}
}

func openDb() *sql.DB {
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db: %s", err)
		os.Exit(1)
	}

	return db
}

func init_db() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ping db: %s", err)
	}

	exec(ctx, db, createMigrationsTableSQL)
}

func migrate() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	migration_files, err := filepath.Glob("./migrations/*.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	existing_migrations := exec(ctx, db, getMigrationIdsSQL)

	migrations := unapplied_migrations(migration_files, existing_migrations)

	for _, file := range migrations {
		statement, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Running migrations...\n")
		migration_id := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		fmt.Fprintf(os.Stdout, "\t%s\n", migration_id)
		fmt.Fprintf(os.Stdout, "Finished migrations\n")

		_, err = db.ExecContext(ctx, string(statement))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
			os.Exit(1)
		}

		add_migration_id_statement := fmt.Sprintf(addMigrationIdSQL, migration_id)
		exec(ctx, db, add_migration_id_statement)
	}
}

func unapplied_migrations(migration_files []string, existing_migrations sql.Result) []string {
	return []string{}
}
