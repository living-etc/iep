package ui

import (
	"os"

	"github.com/charmbracelet/log"
)

var (
	logfile, _ = os.OpenFile("./log/iep", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	Logger     = log.New(logfile)
)
