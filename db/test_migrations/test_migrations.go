package test_migrations

import (
	"db/internal/migrator"
)

var MigrationFunctionRegistry = migrationFunctionRegistry()

func migrationFunctionRegistry() map[string]func() migrator.Migration {
	return map[string]func() migrator.Migration{
		"20240828233901_create_exercises_table": Init_20240828233901_create_exercises_table,
		"20240829233901_add_first_exercise":     Init_20240829233901_add_first_exercise,
		"20240830233901_modify_first_exercise":  Init_20240830233901_modify_first_exercise,
		"20240831233901_add_second_exercise":    Init_20240831233901_add_second_exercise,
	}
}

func Init_20240828233901_create_exercises_table() migrator.Migration {
	return migrator.Migration{
		Id: "20240828233901_create_exercises_table",
		Statement: `
CREATE TABLE IF NOT EXISTS exercises(
  id INTEGER PRIMARY KEY,
  exercise_id TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  body TEXT NOT NULL,
)
    `,
		Args: []any{},
	}
}

func Init_20240829233901_add_first_exercise() migrator.Migration {
	return migrator.Migration{
		Id:        "20240829233901_add_first_exercise",
		Statement: "INSERT INTO exercises(exercise_id, name, description, body) VALUES(?, ?, ?, ?)",
		Args: []any{
			"0001-deploy-a-web-server",
			"Deploy a Web Server with Nginx and AWS",
			"Learn how to put a website on the internet using Nginx and run it on an EC2 instance.",
			`# Deploy A Web App

In this exercise you will deploy a web app to a Linux virtual machine running on AWS. In doing so, you will learn how to

- start a web app and keep it running using Systemd
- install and configure nginx to send traffic to the web app
- configure the security group to allow inbound traffic from the internet

The final setup will look like this:

` + "```" + `
                 ┌──────────────────────────────────────┐
                 │                                      │
                 │  ┌────────────────────────────────┐  │
                 │  │                                │  │
┌─────────┐      │  │  ┌─────────┐      ┌─────────┐  │  │
│         │      │  │  │         │      │         │  │  │
│  Users  ├──────┼──┼──►  Nginx  ├──────►   App   │  │  │
│         │      │  │  │         │      │         │  │  │
└─────────┘      │  │  └─────────┘      └─────────┘  │  │
                 │  │                                │  │
                 │  │    Virtual Machine (Ubuntu)    │  │
                 │  └────────────────────────────────┘  │
                 │                                      │
                 │       Security Group (Firewall)      │
                 └──────────────────────────────────────┘
` + "```",
		},
	}
}

func Init_20240830233901_modify_first_exercise() migrator.Migration {
	return migrator.Migration{
		Id:        "20240830233901_modify_first_exercise",
		Statement: "UPDATE exercises SET description = '?' WHERE exercise_id = '?'",
		Args: []any{
			"Deploy a Web Server with Nginx on AWS",
			"0001-deploy-a-web-server",
		},
	}
}

func Init_20240831233901_add_second_exercise() migrator.Migration {
	return migrator.Migration{
		Id:        "20240831233901_add_second_exercise",
		Statement: "INSERT INTO exercises(exercise_id, name, description, body) VALUES(?, ?, ?, ?)",
		Args: []any{
			"0002-set-up-a-subdomain",
			"Set up a Subdomain",
			"Learn how to put a website on the internet using Nginx and run it on an EC2 instance.",
			`# Deploy A Web App

In this exercise you will set up a DNS subdomain`,
		},
	}
}
