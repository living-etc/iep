package db_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/google/go-cmp/cmp"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	"github.com/living-etc/iep/db"
	"github.com/living-etc/iep/ui"
)

func TestUnappliedMigrations(t *testing.T) {
	tempMigrationsDir := t.TempDir()

	test_cases := []struct {
		name                      string
		migration_files           []string
		unapplied_migrations_want []db.Migration
		completed_migration_ids   []string
	}{
		{
			name: "Add the first migration",
			migration_files: []string{
				"20240828233901_create_exercises_table.md",
			},
			unapplied_migrations_want: []db.Migration{
				{
					Id: "20240828233901_create_exercises_table",
					Filepath: filepath.Join(
						tempMigrationsDir,
						"20240828233901_create_exercises_table.sql",
					),
				},
			},
			completed_migration_ids: []string{},
		},
		{
			name: "No new migrations",
			migration_files: []string{
				"20240828233901_create_exercises_table.md",
			},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
			},
			unapplied_migrations_want: []db.Migration{},
		},
		{
			name: "Add the second migration",
			migration_files: []string{
				"20240828233901_create_exercises_table.md",
				"20240829233901_add_first_exercise.md",
			},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
			},
			unapplied_migrations_want: []db.Migration{
				{
					Id: "20240829233901_add_first_exercise",
					Filepath: filepath.Join(
						tempMigrationsDir,
						"20240829233901_add_first_exercise.sql",
					),
				},
			},
		},
		{
			name: "Add the third and fourth migrations",
			migration_files: []string{
				"20240828233901_create_exercises_table.md",
				"20240829233901_add_first_exercise.md",
				"20240830233901_modify_first_exercise.md",
				"20240831233901_add_second_exercise.md",
			},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
				"20240829233901_add_first_exercise",
			},
			unapplied_migrations_want: []db.Migration{
				{
					Id: "20240830233901_modify_first_exercise",
					Filepath: filepath.Join(
						tempMigrationsDir,
						"20240830233901_modify_first_exercise.sql",
					),
				},
				{
					Id: "20240831233901_add_second_exercise",
					Filepath: filepath.Join(
						tempMigrationsDir,
						"20240831233901_add_second_exercise.sql",
					),
				},
			},
		},
	}

	config, _ := ui.NewConfig([]byte{})
	config.ExerciseDatabase = ":memory:"

	logDirTemp := t.TempDir()
	logfiletemp, err := os.CreateTemp(logDirTemp, "iep.log")
	if err != nil {
		panic(err)
	}

	logfile, err := os.OpenFile(logfiletemp.Name(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	logger := ui.NewLogger(log.DebugLevel, logfile)

	ctx := context.Background()
	conn, err := db.InitDb(ctx, config.ExerciseDatabase)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	for _, tt := range test_cases {
		for _, file := range tt.migration_files {
			_, err := os.CreateTemp(tempMigrationsDir, file)
			if err != nil {
				panic(err)
			}
		}

		db.Exec(ctx, conn, "DELETE FROM migrations")

		for _, id := range tt.completed_migration_ids {
			db.Exec(ctx, conn, "INSERT INTO migrations (id) values (?)", id)
		}

		t.Run(tt.name, func(t *testing.T) {
			unapplied_migrations_got := db.UnappliedMigrations(
				tempMigrationsDir,
				ctx,
				conn,
				logger,
			)

			if diff := cmp.Diff(unapplied_migrations_got, tt.unapplied_migrations_want); diff != "" {
				t.Error(diff)
			}
		})

		os.RemoveAll(tempMigrationsDir)
		os.MkdirAll(tempMigrationsDir, 0777)
	}
}
