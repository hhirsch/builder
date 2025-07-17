package command

import (
	"github.com/hhirsch/builder/internal/ast"
	"testing"
)

func TestOneParameter(test *testing.T) {
	baseCommand := NewBaseCommand()
	result := baseCommand.GetStringFromParameters([]*ast.Node{ast.NewLiteral("Test"), ast.NewLineBreak()})
	if result != "Test" {
		test.Errorf("Expected result to be Test got %v instead.", result)
	}
}

func TestMultipleParameters(test *testing.T) {
	baseCommand := NewBaseCommand()
	result := baseCommand.GetStringFromParameters([]*ast.Node{ast.NewLiteral("Test"), ast.NewLiteral("Message"), ast.NewLineBreak()})
	if result != "Test Message" {
		test.Errorf("Expected result to be TestMessage got %v instead.", result)
	}
}
