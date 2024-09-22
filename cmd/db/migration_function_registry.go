package main

import (
	"iep/cmd/db/migrations"
)

var MigrationFunctionRegistry = migrationFunctionRegistry()

func migrationFunctionRegistry() map[string]func() migrations.Migration {
	return map[string]func() migrations.Migration{
		"20240828233901_create_exercises_table":               migrations.Init_20240828233901_create_exercises_table,
		"20240922172317_add_exercise_0001_depoy_a_webapp":     migrations.Init_20240922172317_add_exercise_0001_depoy_a_webapp,
		"20240922183452_add_exercise_0002_create_a_subdomain": migrations.Init_20240922183452_add_exercise_0002_create_a_subdomain,
	}
}
