package parser

import (
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/lexer"
	"github.com/hhirsch/builder/internal/testUtils"
	"testing"
)

func TestParserWithEmptyFileCreatesEmptyAst(test *testing.T) {
	referenceTree := ast.NewRoot(ast.NewEndOfFile())
	input := ""
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()

	if len(parser.errors) > 0 {
		test.Errorf("Parser should not have any errors.")
	}
	testUtils.ThrowSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}

func TestSingleStatement(test *testing.T) {
	referenceTree := ast.NewRoot(ast.NewStatement("step", ast.NewLiteral("Test")), ast.NewEndOfFile())
	input := "step Test"
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()
	testUtils.ThrowSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}

func TestParserCreatesAst(test *testing.T) {
	referenceTree := ast.NewRoot()
	referenceTree.Children = append(referenceTree.Children,
		ast.NewLineBreak(),
		ast.NewStatement("step", ast.NewLiteral("Test"), ast.NewLineBreak()),
		ast.NewStatement("print",
			ast.NewLiteral("This"), ast.NewLiteral("is"), ast.NewLiteral("a"), ast.NewLiteral("test"),
			ast.NewLineBreak()),
		ast.NewSpace(),
		ast.NewSpace(),
		ast.NewSpace(),
		ast.NewSpace(),
		ast.NewEndOfFile(),
	)
	input := `
step Test
print This is a test
    `
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()
	testUtils.ThrowSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}

func TestParserCreatesAstVariables(test *testing.T) {
	referenceTree := ast.NewRoot()
	referenceTree.Children = append(referenceTree.Children,
		ast.NewLineBreak(),
		ast.NewSpace(), ast.NewSpace(), ast.NewSpace(), ast.NewSpace(), ast.NewSpace(), ast.NewSpace(), ast.NewStatement("print", ast.NewLiteral("This"), ast.NewLiteral("is"), ast.NewLiteral("a"), ast.NewIdentifier("test"), ast.NewLineBreak()),
		ast.NewSpace(), ast.NewSpace(), ast.NewSpace(), ast.NewSpace(), ast.NewEndOfFile(),
	)
	input := `
      print This is a $test
    `
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()

	testUtils.ThrowSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}

func TestParserSteps(test *testing.T) {
	referenceTree := ast.NewRoot()
	referenceTree.Children = append(referenceTree.Children,
		ast.NewStatement("step", ast.NewLiteral("Testing"), ast.NewLiteral("1"), ast.NewLineBreak()),
		ast.NewStatement("step", ast.NewLiteral("Testing"), ast.NewLiteral("2"), ast.NewLineBreak()),
		ast.NewStatement("step", ast.NewLiteral("Testing"), ast.NewLiteral("3")),
		ast.NewEndOfFile(),
	)
	input := `step Testing 1
step Testing 2
step Testing 3`
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()

	testUtils.ThrowSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}

func TestParserCreatesAstForFunctions(test *testing.T) {
	referenceTree := ast.NewRoot()
	referenceTree.Children = append(referenceTree.Children,
		ast.NewFunction("printString",
			[]*ast.Node{ast.NewIdentifierVariadic("string"), ast.NewLineBreak()},
			[]*ast.Node{ast.NewStatement("print", ast.NewIdentifier("string"), ast.NewLineBreak())}),
	)
	referenceTree.Children = append(referenceTree.Children, ast.NewEndOfFile())
	input := `function printString $string...
print $string
done`
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()
	testUtils.ThrowSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}

func TestParserCreatesAstForFunctionsWithReturn(test *testing.T) {
	referenceTree := ast.NewRoot()
	referenceTree.Children = append(referenceTree.Children,
		ast.NewFunction("printString",
			[]*ast.Node{ast.NewIdentifierVariadic("string"), ast.NewLineBreak()},
			[]*ast.Node{ast.NewReturn(ast.NewIdentifier("string")), ast.NewLineBreak()}),
	)
	referenceTree.Children = append(referenceTree.Children, ast.NewEndOfFile())
	input := `function printString $string...
return $string
done`
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	if len(parser.errors) > 0 {
		test.Errorf("%v parser errors.", len(parser.errors))
	}
	syntaxTree := parser.GetSyntaxTree()
	testUtils.ThrowSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}
