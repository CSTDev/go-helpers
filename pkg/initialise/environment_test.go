package initialise

import (
	"os"
	"os/exec"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetEnvironmentVariableKillsWhenNotFound(t *testing.T) {
	if os.Getenv("NO_ENV") == "1" {
		GetEnvironmentVariable("DOES_NOT_EXIST")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestGetEnvironmentVariableKillsWhenNotFound")
	cmd.Env = append(os.Environ(), "NO_ENV=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with error %v, want exist status 1", err)
}

func TestGetEnvironmentVariableReturnsValue(t *testing.T) {
	testValue := "OK"
	os.Setenv("TEST_VAL", testValue)
	res := GetEnvironmentVariable("TEST_VAL")
	if res != testValue {
		t.Fatalf("Result: %s \n Actual %s", res, testValue)
	}
}

func TestOptionalEnvironmentVariableReturnsError(t *testing.T) {
	os.Clearenv()
	_, err := OptionalEnvironmentVariable("TEST_VAL")
	if err == nil {
		t.Fatal("expected error did not return")
	}
}
