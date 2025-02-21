package migrations

func Init_20240828233901_create_exercises_table() Migration {
	return Migration{
		Id: "20240828233901_create_exercises_table",
		Statement: `
CREATE TABLE IF NOT EXISTS exercises(
  id TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  body TEXT NOT NULL
)
    `,
		Args: []any{},
	}
}
