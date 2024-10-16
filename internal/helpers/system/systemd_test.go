package system

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var exampleConfig string

func TestConfigOutput(test *testing.T) {
	systemd := Systemd{}

	result := systemd.GetConfig("User", "/usr/bin/foo", "Test description")
	if exampleConfig != result {
		test.Errorf("Config does not match example %s", result)
	}
}
