package initialise_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/cstdev/go-helpers/pkg/initialise"
	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetEnvironmentVariableKillsWhenNotFound(t *testing.T) {
	if os.Getenv("NO_ENV") == "1" {
		initialise.GetEnvironmentVariable("DOES_NOT_EXIST")
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
	res := initialise.GetEnvironmentVariable("TEST_VAL")
	if res != testValue {
		t.Fatalf("Result: %s \n Actual %s", res, testValue)
	}
}

func TestOptionalEnvironmentVariableReturnsError(t *testing.T) {
	os.Clearenv()
	_, err := initialise.OptionalEnvironmentVariable("TEST_VAL")
	if err == nil {
		t.Fatal("expected error did not return")
	}
}

func TestGetEnvironmentVariableIntKillsWhenString(t *testing.T) {
	if os.Getenv("NO_ENV") == "1" {
		os.Setenv("TEST_VAL", "STRING")
		_ = initialise.GetEnvironmentVariableInt("TEST_VAL")
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestGetEnvironmentVariableIntKillsWhenString")
	cmd.Env = append(os.Environ(), "NO_ENV=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with error %v, want exist status 1", err)

}
