package migrations

import "db/internal/migrator"

func Init_20240828233901_create_exercises_table() migrator.Migration {
	return migrator.Migration{
		Id: "20240828233901_create_exercises_table",
		Statement: `
CREATE TABLE IF NOT EXISTS exercises(
  id INTEGER PRIMARY KEY,
  exercise_id TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  body TEXT NOT NULL
)
    `,
		Args: []any{},
	}
}
