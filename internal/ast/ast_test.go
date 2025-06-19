package ast

import (
	"github.com/hhirsch/builder/internal/token"
	"testing"
)

func TestIsNestable(test *testing.T) {
	rootNode := NewRoot()
	rootNode.Children = append(rootNode.Children,
		NewStatement("step", NewLiteral("Hello World")),
		NewStatement("print", NewLiteral("Hello"), NewIdentifier("world"), NewLiteral("!")),
	)
	if rootNode.Type != token.ROOT {
		test.Errorf("Wrong type for root node.")
	}

	if rootNode.Children[0].Value != "step" {
		test.Errorf("First statement is not step.")
	}

	if rootNode.Children[1].Value != "print" {
		test.Errorf("Second statement is not print.")
	}
}
