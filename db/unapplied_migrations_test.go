package db_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/charmbracelet/log"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	"github.com/living-etc/iep/db"
	"github.com/living-etc/iep/db/migrations"
	"github.com/living-etc/iep/db/test_migrations"
	"github.com/living-etc/iep/ui"
)

func TestGet(t *testing.T) {
	test_cases := []struct {
		name                       string
		migration_funtion_registry map[string]func() migrations.Migration
		unapplied_migrations_want  []migrations.Migration
		completed_migration_ids    []string
	}{
		{
			name: "Add the first migration",
			migration_funtion_registry: map[string]func() migrations.Migration{
				"20240828233901_create_exercises_table": test_migrations.Init_20240828233901_create_exercises_table,
			},
			unapplied_migrations_want: []migrations.Migration{
				{
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
				},
			},
			completed_migration_ids: []string{},
		},
		{
			name: "No new migrations",
			migration_funtion_registry: map[string]func() migrations.Migration{
				"20240828233901_create_exercises_table": test_migrations.Init_20240828233901_create_exercises_table,
			},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
			},
			unapplied_migrations_want: []migrations.Migration{},
		},
		{
			name: "Add the second migration",
			migration_funtion_registry: map[string]func() migrations.Migration{
				"20240828233901_create_exercises_table": test_migrations.Init_20240828233901_create_exercises_table,
				"20240829233901_add_first_exercise":     test_migrations.Init_20240829233901_add_first_exercise,
			},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
			},
			unapplied_migrations_want: []migrations.Migration{
				{
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
				},
			},
		},
		{
			name: "Add the third and fourth migrations",
			migration_funtion_registry: map[string]func() migrations.Migration{
				"20240828233901_create_exercises_table": test_migrations.Init_20240828233901_create_exercises_table,
				"20240829233901_add_first_exercise":     test_migrations.Init_20240829233901_add_first_exercise,
				"20240830233901_modify_first_exercise":  test_migrations.Init_20240830233901_modify_first_exercise,
				"20240831233901_add_second_exercise":    test_migrations.Init_20240831233901_add_second_exercise,
			},
			completed_migration_ids: []string{
				"20240828233901_create_exercises_table",
				"20240829233901_add_first_exercise",
			},
			unapplied_migrations_want: []migrations.Migration{
				{
					Id:        "20240830233901_modify_first_exercise",
					Statement: "UPDATE exercises SET description = '?' WHERE exercise_id = '?'",
					Args: []any{
						"Deploy a Web Server with Nginx on AWS",
						"0001-deploy-a-web-server",
					},
				},
				{
					Id:        "20240831233901_add_second_exercise",
					Statement: "INSERT INTO exercises(exercise_id, name, description, body) VALUES(?, ?, ?, ?)",
					Args: []any{
						"0002-set-up-a-subdomain",
						"Set up a Subdomain",
						"Learn how to put a website on the internet using Nginx and run it on an EC2 instance.",
						`# Deploy A Web App

In this exercise you will set up a DNS subdomain`,
					},
				},
			},
		},
	}

	dbName := "file::memory:"

	config := ui.Config{
		ExerciseDatabase: ":memory:",
		LogFile:          "/Users/chris/Code/personal/infrastructure-exercism-prototype/log/test.log",
	}

	logfile, err := os.OpenFile(config.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	logger := ui.NewLogger(log.DebugLevel, logfile)

	ctx := context.Background()
	conn, err := db.InitDb(ctx, dbName)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	originalMigrationFunctionRegistry := db.MigrationFunctionRegistry
	defer func() { db.MigrationFunctionRegistry = originalMigrationFunctionRegistry }()

	for _, tt := range test_cases {
		db.MigrationFunctionRegistry = tt.migration_funtion_registry

		db.Exec(ctx, conn, "DELETE FROM migrations")

		for _, id := range tt.completed_migration_ids {
			db.Exec(ctx, conn, "INSERT INTO migrations (id) values (?)", id)
		}

		t.Run(tt.name, func(t *testing.T) {
			unapplied_migrations_got := db.Get(
				ctx,
				conn,
			)

			if !reflect.DeepEqual(unapplied_migrations_got, tt.unapplied_migrations_want) {
				t.Errorf("want %v, got %v", tt.unapplied_migrations_want, unapplied_migrations_got)
			}
		})
	}
}
