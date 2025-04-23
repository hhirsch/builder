package interpreter

import (
	"github.com/hhirsch/builder/internal/helpers"
	"testing"
)

func TestRequirementsPassWhenUnspecified(t *testing.T) {
	baseCommand := &BaseCommand{}
	if !baseCommand.TestRequirements() {
		t.Errorf("Empty test requirements did not pass")
	}
}

func TestReplaceVariablesInTokens(t *testing.T) {
	baseCommand := BaseCommand{logger: helpers.NewLogger("/tmp/testLogger.txt")}
	result := baseCommand.replaceVariablesInTokens([]string{"print", "hello", "$world"}, map[string]Variable{"world": Variable{stringContent: "test"}})
	if result[0] != "print" {
		t.Error("Command should be untouched by replace variables.")
	}

	if result[1] != "hello" {
		t.Error("Strings should be untouched by replace variables.")
	}

	if result[2] != "test" {
		t.Error("Variables should be replaced by the content of the variable.")
	}
}
