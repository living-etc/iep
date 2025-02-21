package migrations

import (
	"os"
	"path/filepath"
)

func Init_20241206141112_add_exercise_0006_advanced_git() Migration {
	content, err := os.ReadFile(
		filepath.Join(
			MARKDOWN_PATH,
			"20241206141112_add_exercise_0006_advanced_git.md",
		),
	)
	if err != nil {
		panic(err)
	}

	return Migration{
		Id:        "20241206141112_add_exercise_0006_advanced_git",
		Statement: "INSERT INTO exercises(id, name, description, body) VALUES(?, ?, ?, ?)",
		Args: []any{
			"0005-advanced-git",
			"Advanced Git",
			"Advanced Git",
			string(content),
		},
	}
}
