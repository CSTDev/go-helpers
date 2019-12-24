package initialise

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
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

// GetEnvironmentVariableInt will try to find the value for the provided key
// If it can't be found it logs and exits the program
func GetEnvironmentVariableInt(key string) int {
	variable := os.Getenv(key)
	if variable == "" {
		log.Error("Missing environment variable: " + key)
		os.Exit(1)
	}

	value, err := strconv.Atoi(variable)
	if err != nil {
		log.Error("error converting value to string")
		os.Exit(1)
	}
	return value
}

// OptionalEnvironmentVariableInt will try to find the value for the provided key
// If it can't be found it logs and returns and error that can be ignored
func OptionalEnvironmentVariableInt(key string) (int, error) {
	variable := os.Getenv(key)
	if variable == "" {
		log.Warn("Missing environment variable: " + key)
		return -1, errors.New("Missing environment variable: " + key)
	}
	return strconv.Atoi(variable)
}
