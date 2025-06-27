package command

import (
	"github.com/hhirsch/builder/internal/ast"
	"testing"
)

func TestCustomCommandPreservesBody(test *testing.T) {
	body := []*ast.Node{
		ast.NewStatement("print", ast.NewLiteral("hello"), ast.NewIdentifier("parameter")),
		ast.NewLineBreak(),
	}
	customCommand := NewCustomCommand(ast.NewFunction("testFunction",
		[]*ast.Node{
			ast.NewIdentifierVariadic("parameter"),
		},
		body,
	))
	ast.CompareAbstractSyntaxNodeSlices(body, customCommand.body)
}
