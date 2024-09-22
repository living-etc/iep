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
	"db/internal/unapplied_migrations"
	"db/migrations"
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
		unapplied_migrations := unapplied_migrations.Get(
			ctx,
			db,
			migrationFilePaths(),
		)

		if len(unapplied_migrations) == 0 {
			fmt.Fprintf(os.Stdout, "No migrations to run\n")
			os.Exit(0)
		}

		var migrationRunner migrations.MigrationRunner

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
