package migrations

func Init_20250220155451_create_tests_table() Migration {
	return Migration{
		Id: "20250220155451_create_tests_table",
		Statement: `
CREATE TABLE IF NOT EXISTS tests(
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  exercise_id TEXT NOT NULL,
  resource_type TEXT NOT NULL,
  resource_name TEXT NOT NULL,
  resource_attribute TEXT NOT NULL,
  resource_attribute_value TEXT NOT NULL,
  negation INTEGER NOT NULL
)
    `,
		Args: []any{},
	}
}
