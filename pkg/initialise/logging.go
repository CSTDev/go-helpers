package initialise

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// SetupLogging initalises a logrus logger with the log level set by environment variable
func SetupLogging() {
	log.SetFormatter(&log.JSONFormatter{})

	logLevel := os.Getenv("LOG_LEVEL")
	log.SetFormatter(&log.JSONFormatter{})

	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
		break
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
		break
	default:
		log.SetLevel(log.InfoLevel)
	}
}
