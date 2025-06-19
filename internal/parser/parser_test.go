package parser

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/lexer"
)

func getStringFromTree(nodes []*ast.Node) string {
	syntaxTreeJson, _ := json.Marshal(nodes)
	return string(syntaxTreeJson)
}

func throwSyntaxTreeDoesNotMatchError(test *testing.T, tree *ast.Node, referenceTree *ast.Node) {
	if !cmp.Equal(tree, referenceTree) {
		test.Logf("Syntax tree:\n%s \n\ndoes not match reference:\n\n%s.", getStringFromTree(tree.Children), getStringFromTree(referenceTree.Children))
		test.Errorf("Syntax tree does not match reference.")
	}
}

func TestParserWithEmptyFileCreatesEmptyAst(test *testing.T) {
	referenceTree := ast.NewRoot(ast.NewEndOfFile())
	input := ""
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()

	if len(parser.errors) > 0 {
		test.Errorf("Parser should not have any errors.")
	}
	throwSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}

func TestSingleStatement(test *testing.T) {
	referenceTree := ast.NewRoot(ast.NewStatement("step", ast.NewLiteral("Test")), ast.NewEndOfFile())
	input := "step Test"
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()
	throwSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)

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
	throwSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
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

	throwSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
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

	throwSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}

func TestParserCreatesAstForFunctions(test *testing.T) {
	referenceTree := ast.NewRoot()
	referenceTree.Children = append(referenceTree.Children,
		ast.NewFunction("printString",
			[]*ast.Node{ast.NewIdentifierVariadic("string..."), ast.NewLineBreak()},
			[]*ast.Node{ast.NewStatement("print", ast.NewIdentifier("string"), ast.NewLineBreak())}),
	)
	referenceTree.Children = append(referenceTree.Children, ast.NewEndOfFile())
	input := `function printString $string...
print $string
done`
	lexer, _ := lexer.NewLexer(input)
	parser, _ := NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()
	throwSyntaxTreeDoesNotMatchError(test, syntaxTree, referenceTree)
}
