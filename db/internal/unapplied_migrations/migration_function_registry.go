package unapplied_migrations

import (
	"db/internal/migrator"
	"db/migrations"
)

var MigrationFunctionRegistry = migrationFunctionRegistry()

func migrationFunctionRegistry() map[string]func() migrator.Migration {
	return map[string]func() migrator.Migration{
		"20240828233901_create_exercises_table": migrations.Init_20240828233901_create_exercises_table,
	}
}
