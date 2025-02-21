package ui_test

import (
	"context"
	"os"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/google/go-cmp/cmp"

	"github.com/living-etc/iep/db"
	"github.com/living-etc/iep/ui"
)

func TestTests(t *testing.T) {
	testcases := []struct {
		name               string
		exerciseAttributes map[string]string
		testsWant          []ui.Test
	}{
		{
			name: "read_two_tests",
			exerciseAttributes: map[string]string{
				"Id":          "0001-deploy-a-webapp",
				"title":       "Deploy a webapp",
				"description": "Learn to deploy a web app with nginx",
				"content":     "exercise content",
			},
			testsWant: []ui.Test{
				{
					Id:                     1,
					Name:                   "Nginx is installed",
					ExerciseId:             "0001-deploy-a-webapp",
					ResourceType:           "Package",
					ResourceName:           "nginx",
					ResourceAttribute:      "Status",
					ResourceAttributeValue: "install ok installed",
					Negation:               true,
				},
				{
					Id:                     2,
					Name:                   "Nginx service is running",
					ExerciseId:             "0001-deploy-a-webapp",
					ResourceType:           "Service",
					ResourceName:           "nginx",
					ResourceAttribute:      "ActiveState",
					ResourceAttributeValue: "active",
					Negation:               true,
				},
			},
		},
	}

	config, _ := ui.NewConfig([]byte{})
	config.ExerciseDatabase = ":memory:"

	logfile, err := os.OpenFile(config.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	logger := ui.NewLogger(log.DebugLevel, logfile)

	ctx := context.Background()
	conn, err := db.InitDb(ctx, config.ExerciseDatabase)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()
	db.RunMigrations(config, logger, conn)

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			exercise := ui.NewExercise(testcase.exerciseAttributes)
			testsGot := exercise.Tests(conn, logger)

			if diff := cmp.Diff(testcase.testsWant, testsGot); diff != "" {
				t.Error(diff)
			}
		})
	}
}
