package migrations

import (
	"os"
	"path/filepath"
)

func Init_20241206132441_add_exercise_0004_introduction_to_source_control_with_git() Migration {
	content, err := os.ReadFile(
		filepath.Join(
			MARKDOWN_PATH,
			"20241206132441_add_exercise_0004_introduction_to_source_control_with_git.md",
		),
	)
	if err != nil {
		panic(err)
	}

	return Migration{
		Id:        "20241206132441_add_exercise_0004_introduction_to_source_control_with_git",
		Statement: "INSERT INTO exercises(id, name, description, body) VALUES(?, ?, ?, ?)",
		Args: []any{
			"0004-intro-to-git",
			"Introduction to source control with Git",
			"Learn to manage changes to your code using Git",
			string(content),
		},
	}
}
