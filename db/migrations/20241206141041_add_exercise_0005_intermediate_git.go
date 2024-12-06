package migrations

import (
	"os"
	"path/filepath"
)

func Init_20241206141041_add_exercise_0005_intermediate_git() Migration {
	content, err := os.ReadFile(
		filepath.Join(
			MARKDOWN_PATH,
			"20241206141041_add_exercise_0005_intermediate_git.md",
		),
	)
	if err != nil {
		panic(err)
	}

	return Migration{
		Id:        "20241206141041_add_exercise_0005_intermediate_git",
		Statement: "INSERT INTO exercises(exercise_id, name, description, body) VALUES(?, ?, ?, ?)",
		Args: []any{
			"0005-intermediate-git",
			"Intermediate Git",
			"Intermediate Git",
			string(content),
		},
	}
}
