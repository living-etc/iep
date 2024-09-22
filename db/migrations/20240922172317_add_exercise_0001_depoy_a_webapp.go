package migrations

import "db/internal/migrator"

func Init_20240922172317_add_exercise_0001_depoy_a_webapp() migrator.Migration {
	return migrator.Migration{
		Id:        "20240922172317_add_exercise_0001_depoy_a_webapp",
		Statement: "INSERT INTO exercises(exercise_id, name, description, body) VALUES(?, ?, ?, ?)",
		Args: []any{
			"0001-deploy-a-webapp",
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
