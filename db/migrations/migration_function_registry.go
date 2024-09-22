package migrations

var MigrationFunctionRegistry = migrationFunctionRegistry()

func migrationFunctionRegistry() map[string]func() Migration {
	return map[string]func() Migration{
		"20240828233901_create_exercises_table": Init_20240828233901_create_exercises_table,
	}
}
