package testUtils

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/hhirsch/builder/internal/ast"
	"testing"
)

func getStringFromTree(nodes []*ast.Node) string {
	syntaxTreeJson, _ := json.Marshal(nodes)
	return string(syntaxTreeJson)
}

func ThrowSyntaxTreeDoesNotMatchErrorForSlices(test *testing.T, tree []*ast.Node, referenceTree []*ast.Node) {
	if !cmp.Equal(tree, referenceTree) {
		test.Logf("Syntax tree:\n%s \n\ndoes not match reference:\n\n%s.", getStringFromTree(tree), getStringFromTree(referenceTree))
		test.Errorf("Syntax tree does not match reference.")
	}
}

func ThrowSyntaxTreeDoesNotMatchError(test *testing.T, tree *ast.Node, referenceTree *ast.Node) {
	ThrowSyntaxTreeDoesNotMatchErrorForSlices(test, tree.Children, referenceTree.Children)
}
