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
			ast.NewIdentifier("parameter"),
		},
		body,
	))
	ast.CompareAbstractSyntaxNodeSlices(body, customCommand.body)
}

func TestCustomCommandVariadicParameter(test *testing.T) {
	body := []*ast.Node{
		ast.NewStatement("return", ast.NewIdentifier("text")),
		ast.NewLineBreak(),
	}
	customCommand := NewCustomCommand(ast.NewFunction("printText",
		[]*ast.Node{
			ast.NewIdentifierVariadic("text"),
		},
		body,
	))
	result, _ := customCommand.Execute([]*ast.Node{ast.NewLiteral("foo"), ast.NewLiteral("bar"), ast.NewLiteral("baz")})
	expectation := "foo bar baz"
	if result != expectation {
		test.Errorf("Variadic parameter lead to wrong result. Expected \"%v\" got \"%v\" instead.", expectation, result)
	}
}

func TestCustomCommandFails(test *testing.T) {
	body := []*ast.Node{
		ast.NewStatement("return", ast.NewIdentifier("text")),
		ast.NewLineBreak(),
	}
	customCommand := NewCustomCommand(ast.NewFunction("printText",
		[]*ast.Node{
			ast.NewIdentifier("text"),
			ast.NewIdentifierVariadic("text"),
		},
		body,
	))
	error := customCommand.Validate([]*ast.Node{ast.NewLiteral("foo")})
	if error == nil {
		test.Errorf("Missing parameters should trigger an error.")
	}
}

func TestCustomCommandValidates(test *testing.T) {
	body := []*ast.Node{
		ast.NewStatement("return", ast.NewIdentifier("text")),
		ast.NewLineBreak(),
	}
	customCommand := NewCustomCommand(ast.NewFunction("printText",
		[]*ast.Node{
			ast.NewIdentifierVariadic("text"),
		},
		body,
	))
	error := customCommand.Validate([]*ast.Node{ast.NewLiteral("foo"), ast.NewLiteral("bar"), ast.NewLiteral("baz")})
	if error != nil {
		test.Errorf("Variadic parameter not accepted error \"%v\".", error.Error())
	}
}

func TestCustomCommandVariadicCorrectVariableInPool(test *testing.T) {
	body := []*ast.Node{
		ast.NewStatement("return", ast.NewIdentifier("text")),
		ast.NewLineBreak(),
	}
	customCommand := NewCustomCommand(ast.NewFunction("printText",
		[]*ast.Node{
			ast.NewIdentifierVariadic("text"),
		},
		body,
	))
	parameters := []*ast.Node{ast.NewLiteral("foo"), ast.NewLiteral("bar"), ast.NewLiteral("baz")}
	error := customCommand.Validate(parameters)
	customCommand.Execute(parameters)
	if error != nil {
		test.Errorf("Variadic parameter not accepted error \"%v\".", error.Error())
	}
	/*
		result, error := customCommand.variables.GetVariable("text")
		if error != nil {
			test.Errorf("Unable to get variable \"%v\".", error.Error())
		}
		if result.GetStringContent() != "foo bar baz" {
			test.Errorf("Unexpected string content \"%v\".", result.GetStringContent())
		}*/
}
