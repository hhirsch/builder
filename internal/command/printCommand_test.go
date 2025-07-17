package command

import (
	"github.com/hhirsch/builder/internal/ast"
	"testing"
)

func TestPrintBuffer(test *testing.T) {
	expectedString := "foo"
	var writer Writer = NewBufferWriter()
	printCommand := NewPrintCommand(&writer)
	printCommand.Execute([]*ast.Node{ast.NewLiteral(expectedString)})
	if writer.GetHistory()[0] != expectedString {
		test.Errorf("Wrong token literal: %v. Expected: %v.", writer.GetHistory()[0], expectedString)
	}
}
