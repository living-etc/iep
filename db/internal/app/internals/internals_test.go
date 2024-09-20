package internals_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"

	"db/internal/app/internals"
)

const (
	createMigrationsTableSQL = "CREATE TABLE IF NOT EXISTS migrations(id TEXT NOT NULL, PRIMARY KEY(id))"
)

func TestUnappliedMigrations(t *testing.T) {
	//	test_cases := []struct {
	//		name                      string
	//		migration_files           []string
	//		completed_migrations      []string
	//		unapplied_migrations_want []string
	//	}{
	//
	//		{
	//			name: "Add the first migration",
	//			migration_files: []string{
	//				"migrations/20240828233901_create_exercises_table.sql",
	//			},
	//			completed_migrations: []string{},
	//			unapplied_migrations_want: []string{
	//				"migrations/20240828233901_create_exercises_table.sql",
	//			},
	//		},
	//		{
	//			name: "No new migrations",
	//			migration_files: []string{
	//				"migrations/20240828233901_create_exercises_table.sql",
	//			},
	//			completed_migrations: []string{
	//				"20240828233901_create_exercises_table",
	//			},
	//			unapplied_migrations_want: []string{},
	//		},
	//		{
	//			name: "Add the second migration",
	//			migration_files: []string{
	//				"migrations/20240828233901_create_exercises_table.sql",
	//				"migrations/20240829233901_create_more_things.sql",
	//			},
	//			completed_migrations: []string{
	//				"20240828233901_create_exercises_table",
	//			},
	//			unapplied_migrations_want: []string{
	//				"migrations/20240829233901_create_more_things.sql",
	//			},
	//		},
	//		{
	//			name: "Add the third and fourth migrations",
	//			migration_files: []string{
	//				"migrations/20240828233901_create_exercises_table.sql",
	//				"migrations/20240829233901_create_more_things.sql",
	//				"migrations/20240830233901_create_more_things_2.sql",
	//				"migrations/20240831233901_create_more_things_3.sql",
	//			},
	//			completed_migrations: []string{
	//				"20240828233901_create_exercises_table",
	//				"20240829233901_create_more_things",
	//			},
	//			unapplied_migrations_want: []string{
	//				"migrations/20240830233901_create_more_things_2.sql",
	//				"migrations/20240831233901_create_more_things_3.sql",
	//			},
	//		},
	//	}
}

func openDb(dbName string) *sql.DB {
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db: %s", err)
		os.Exit(1)
	}

	return db
}

func Init_db(dbName string) {
	db := openDb(dbName)
	defer db.Close()

	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ping db: %s", err)
	}

	exec(ctx, db, createMigrationsTableSQL)
}

func exec(ctx context.Context, db *sql.DB, statement string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}

func TestUnappliedMigrationsNew(t *testing.T) {
	test_cases := []struct {
		name                      string
		migration_files           []string
		unapplied_migrations_want []internals.Migration
	}{
		{
			name: "Add the first migration",
			migration_files: []string{
				"migrations/20240828233901_create_exercises_table.sql",
			},
			unapplied_migrations_want: []internals.Migration{
				{
					Filepath: "migrations/20240828233901_create_exercises_table.sql",
				},
			},
		},
	}

	Init_db("file:exercises-test.db")
	db := openDb("file:exercises-test.db")

	for _, tt := range test_cases {
		ctx := context.Background()

		t.Run(tt.name, func(t *testing.T) {
			unapplied_migrations_got := internals.UnappliedMigrationsNew(
				ctx,
				db,
				tt.migration_files,
			)

			if !reflect.DeepEqual(unapplied_migrations_got, tt.unapplied_migrations_want) {
				t.Errorf("want %v, got %v", tt.unapplied_migrations_want, unapplied_migrations_got)
			}
		})
	}
}
