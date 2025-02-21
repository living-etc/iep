package migrations

import (
	"os"
	"path/filepath"
)

func Init_20241203113948_add_exercise_0003_create_your_first_aws_account() Migration {
	content, err := os.ReadFile(
		filepath.Join(
			MARKDOWN_PATH,
			"20241203113948_add_exercise_0003_create_your_first_aws_account.md",
		),
	)
	if err != nil {
		panic(err)
	}

	return Migration{
		Id:        "20241203113948_add_exercise_0003_create_your_first_aws_account",
		Statement: "INSERT INTO exercises(id, name, description, body) VALUES(?, ?, ?, ?)",
		Args: []any{
			"0003-create-an-aws-account",
			"Create your first AWS account",
			"Learn to bring together all the moving parts involved in setting up your first AWS account",
			string(content),
		},
	}
}
