package initialise

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// GetEnvironmentVariable will try to find the value for the provided key
// If it can't be found it logs and exits the program
func GetEnvironmentVariable(key string) string {
	variable := os.Getenv(key)
	if variable == "" {
		log.Error("Missing environment variable: " + key)
		os.Exit(1)
	}
	return variable
}

// OptionalEnvironmentVariable will try to find the value for the provided key
// If it can't be found it logs and returns and error that can be ignored
func OptionalEnvironmentVariable(key string) (string, error) {
	variable := os.Getenv(key)
	if variable == "" {
		log.Warn("Missing environment variable: " + key)
		return "", errors.New("Missing environment variable: " + key)
	}
	return variable, nil
}
