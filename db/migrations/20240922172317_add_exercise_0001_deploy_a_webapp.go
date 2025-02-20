package migrations

import (
	"os"
	"path/filepath"
)

func Init_20240922172317_add_exercise_0001_deploy_a_webapp() Migration {
	content, err := os.ReadFile(
		filepath.Join(
			MARKDOWN_PATH,
			"20240922172317_add_exercise_0001_deploy_a_webapp.md",
		),
	)
	if err != nil {
		panic(err)
	}

	return Migration{
		Id:        "20240922172317_add_exercise_0001_deploy_a_webapp",
		Statement: "INSERT INTO exercises(exercise_id, name, description, body) VALUES(?, ?, ?, ?)",
		Args: []any{
			"0001-deploy-a-webapp",
			"Deploy a Web Server with Nginx and AWS",
			"Learn how to put a website on the internet using Nginx and run it on an EC2 instance.",
			string(content),
		},
	}
}
