package db

import (
	"github.com/living-etc/iep/db/migrations"
)

var MigrationFunctionRegistry = migrationFunctionRegistry()

func migrationFunctionRegistry() map[string]func() migrations.Migration {
	return map[string]func() migrations.Migration{
		"20240828233901_create_exercises_table":                                    migrations.Init_20240828233901_create_exercises_table,
		"20240922172317_add_exercise_0001_depoy_a_webapp":                          migrations.Init_20240922172317_add_exercise_0001_depoy_a_webapp,
		"20240922183452_add_exercise_0002_create_a_subdomain":                      migrations.Init_20240922183452_add_exercise_0002_create_a_subdomain,
		"20241203113948_add_exercise_0003_create_your_first_aws_account":           migrations.Init_20241203113948_add_exercise_0003_create_your_first_aws_account,
		"20241206132441_add_exercise_0004_introduction_to_source_control_with_git": migrations.Init_20241206132441_add_exercise_0004_introduction_to_source_control_with_git,
		"20241206141041_add_exercise_0005_intermediate_git":                        migrations.Init_20241206141041_add_exercise_0005_intermediate_git,
		"20241206141112_add_exercise_0006_advanced_git":                            migrations.Init_20241206141112_add_exercise_0006_advanced_git,
	}
}
