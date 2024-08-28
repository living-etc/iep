package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

const (
	createMigrationsTableSQL = "CREATE TABLE IF NOT EXISTS migrations(id INTEGER NOT NULL, PRIMARY KEY(id))"
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

func main() {
	args := os.Args[1:]

	subcommand := args[0]

	switch subcommand {
	case "init":
		init_db()
	case "up":
		migrate_up()
	case "down":
		migrate_down()
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

func migrate_up() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	var migration Migration_20240828233901
	migration.up(ctx, db)
}

func migrate_down() {
	db := openDb()
	defer db.Close()

	ctx := context.Background()

	var migration Migration_20240828233901
	migration.down(ctx, db)
}
