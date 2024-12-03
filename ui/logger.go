package ui

import (
	"io"

	"github.com/charmbracelet/log"
)

func NewLogger(logLevel log.Level, writer io.Writer) *log.Logger {
	return log.NewWithOptions(writer, log.Options{
		Level:           logLevel,
		ReportTimestamp: true,
		ReportCaller:    true,
	})
}
