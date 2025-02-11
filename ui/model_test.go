package ui_test

import (
	"context"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/teatest"

	"github.com/living-etc/iep/db"
	"github.com/living-etc/iep/ui"
)

func TestNewModel(t *testing.T) {
	testcases := []struct {
		name       string
		keystrokes string
		modelWant  ui.Model
	}{
		{
			name:       "j",
			modelWant:  ui.Model{Cursor: 1},
			keystrokes: "j",
		},
		{
			name:       "j-k",
			modelWant:  ui.Model{Cursor: 0},
			keystrokes: "jk",
		},
		{
			name:       "j-j-j-j",
			modelWant:  ui.Model{Cursor: 4},
			keystrokes: "jjjj",
		},
	}

	config := ui.NewConfig()
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
			testModel := teatest.NewTestModel(t, ui.NewModel(config, logger, conn))

			logger.Debug("Sending", "keystroke", testcase.keystrokes)
			testModel.Type(testcase.keystrokes)

			testModel.Send(tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("q"),
			})

			modelGot := testModel.FinalModel(t).(ui.Model)
			modelWant := testcase.modelWant

			if modelGot.Cursor != modelWant.Cursor {
				t.Fatalf("want: %v, got: %v", modelWant.Cursor, modelGot.Cursor)
			}
		})
	}
}
