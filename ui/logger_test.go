package ui_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/charmbracelet/log"

	"github.com/living-etc/iep/ui"
)

func TestNewLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := ui.NewLogger(log.DebugLevel, &buf)

	t.Run("some test", func(t *testing.T) {
		level := logger.GetLevel().String()
		if level != "debug" {
			t.Errorf("want %v, got %v", "debug", level)
		}

		testLog := "this is only a test"
		logger.Debug(testLog)

		reader := bufio.NewReader(&buf)
		logline, _, err := reader.ReadLine()
		if err != nil {
			t.Fatal("unable to read from bugger:", err)
		}

		if !strings.Contains(string(logline), "this is only a test") {
			t.Errorf("log %v does not contain %v", "this is only a test", string(logline))
		}
	})
}
