package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	"db/internal/app/internals"
)

const (
	dbName = "file:exercises.db"
)

func openDb() *sql.DB {
	db, err := sql.Open("libsql", dbName)
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
		internals.InitDb()
	case "migrate":
		unapplied_migrations := internals.UnappliedMigrations(ctx, db, migrationFilePaths())

		if len(unapplied_migrations) == 0 {
			fmt.Fprintf(os.Stdout, "No migrations to run\n")
			os.Exit(0)
		}

		for _, migration := range unapplied_migrations {
			migration.Run(ctx, db)
		}
	}
}

func migrationFilePaths() []string {
	migration_files, err := filepath.Glob("./migrations/*.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	return migration_files
}
