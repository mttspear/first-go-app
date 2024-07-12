package logging

import (
	"os"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("example")

func SetupLogging() {
	// Create a file to write logs to
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatalf("Failed to open log file: %v", err)
	}

	// Example format string. Everything except the message has a custom color
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(logFile, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Set the backends to be used.
	logging.SetBackend(backend1Formatter, backend2Formatter)
}
