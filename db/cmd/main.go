package main

import (
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	"db/internal/app/internals"
)

func main() {
	args := os.Args[1:]

	subcommand := args[0]

	switch subcommand {
	case "init":
		internals.Init_db()
	case "migrate":
		internals.Migrate()
	case "migrate_data":
		internals.MigrateData()
	}
}
