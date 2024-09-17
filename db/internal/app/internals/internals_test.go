package internals_test

import (
	"context"
	"reflect"
	"testing"

	"db/internal/app/internals"
)

func TestUnappliedMigrations(t *testing.T) {
	test_cases := []struct {
		name                      string
		migration_files           []string
		completed_migrations      []string
		unapplied_migrations_want []string
	}{
		{
			name: "Add the first migration",
			migration_files: []string{
				"migrations/20240828233901_create_exercises_table.sql",
			},
			completed_migrations: []string{},
			unapplied_migrations_want: []string{
				"migrations/20240828233901_create_exercises_table.sql",
			},
		},
		{
			name: "No new migrations",
			migration_files: []string{
				"migrations/20240828233901_create_exercises_table.sql",
			},
			completed_migrations: []string{
				"20240828233901_create_exercises_table",
			},
			unapplied_migrations_want: []string{},
		},
		{
			name: "Add the second migration",
			migration_files: []string{
				"migrations/20240828233901_create_exercises_table.sql",
				"migrations/20240829233901_create_more_things.sql",
			},
			completed_migrations: []string{
				"20240828233901_create_exercises_table",
			},
			unapplied_migrations_want: []string{
				"migrations/20240829233901_create_more_things.sql",
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
			completed_migrations: []string{
				"20240828233901_create_exercises_table",
				"20240829233901_create_more_things",
			},
			unapplied_migrations_want: []string{
				"migrations/20240830233901_create_more_things_2.sql",
				"migrations/20240831233901_create_more_things_3.sql",
			},
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.name, func(t *testing.T) {
			unapplied_migrations_got := internals.UnappliedMigrations(
				tt.migration_files,
				tt.completed_migrations,
			)

			if !reflect.DeepEqual(unapplied_migrations_got, tt.unapplied_migrations_want) {
				t.Errorf("want %v, got %v", tt.unapplied_migrations_want, unapplied_migrations_got)
			}
		})
	}
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

	for _, tt := range test_cases {
		ctx := context.Background()

		db := &MockDB{}

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
