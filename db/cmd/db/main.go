package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	"db/internal/migrator"
	"db/internal/unapplied_migrations"
)

func InitDb() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ping db: %s", err)
	}

	exec(ctx, db, "CREATE TABLE IF NOT EXISTS migrations(id TEXT NOT NULL, PRIMARY KEY(id))")
}

func exec(ctx context.Context, db *sql.DB, statement string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}

func openDb() *sql.DB {
	db, err := sql.Open("libsql", "file:exercises.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db: %s", err)
		os.Exit(1)
	}

	return db
}

func main() {
	args := os.Args[1:]

	subcommand := args[0]

	ctx := context.Background()

	db := openDb()
	defer db.Close()

	switch subcommand {
	case "init":
		InitDb()
	case "migrate":
		unapplied_migrations := unapplied_migrations.Get(
			ctx,
			db,
			migrationFilePaths(),
		)

		if len(unapplied_migrations) == 0 {
			fmt.Fprintf(os.Stdout, "No migrations to run\n")
			os.Exit(0)
		}

		var migrationRunner migrator.MigrationRunner

		for _, migration := range unapplied_migrations {
			err := migrationRunner.Run(ctx, db, migration)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
				os.Exit(1)
			}
		}
	}
}

func migrationFilePaths() []string {
	migration_files, err := filepath.Glob("./migrations/*.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	return migration_files
}
