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

func TestUnappliedMigrations(t *testing.T) {
	test_cases := []struct {
		name                      string
		migration_files           []string
		unapplied_migrations_want []internals.Migration
		completed_migration_ids   []string
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
			completed_migration_ids: []string{},
		},
		{
			name: "No new migrations",
			migration_files: []string{
				"migrations/20240828233901_create_exercises_table.sql",
			},
			unapplied_migrations_want: []internals.Migration{},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
			},
		},
		{
			name: "Add the second migration",
			migration_files: []string{
				"migrations/20240828233901_create_exercises_table.sql",
				"migrations/20240829233901_create_more_things.sql",
			},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
			},
			unapplied_migrations_want: []internals.Migration{
				{
					Filepath: "migrations/20240829233901_create_more_things.sql",
				},
			},
		},
		{
			name: "Add the third and fourth migrations",
			migration_files: []string{
				"migrations/20240828233901_create_exercises_table.sql",
				"migrations/20240829233901_create_more_things.sql",
				"migrations/20240830233901_create_more_things_2.sql",
				"migrations/20240831233901_create_more_things_3.sql",
			},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
				"20240829233901_create_more_things",
			},
			unapplied_migrations_want: []internals.Migration{
				{
					Filepath: "migrations/20240830233901_create_more_things_2.sql",
				},
				{
					Filepath: "migrations/20240831233901_create_more_things_3.sql",
				},
			},
		},
	}

	dbName := "file:exercises-test.db"

	Init_db(dbName)
	db := openDb(dbName)

	for _, tt := range test_cases {
		ctx := context.Background()

		exec(ctx, db, "DELETE FROM migrations")

		for _, id := range tt.completed_migration_ids {
			exec(ctx, db, "INSERT INTO migrations (id) values (?)", id)
		}

		t.Run(tt.name, func(t *testing.T) {
			unapplied_migrations_got := internals.UnappliedMigrations(
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
