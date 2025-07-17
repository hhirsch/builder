package interpreterV2

import (
	"errors"
	"github.com/hhirsch/builder/internal/ast"
)

type ReturnStatement struct {
}

func NewReturnStatement() *ReturnStatement {
	return &ReturnStatement{}
}

func (returnStatement *ReturnStatement) Validate(parameters []*ast.Node) error {
	if len(parameters) != 1 {
		return errors.New("Return statement needs exactly one parameter.")
	}
	return nil
}

func (returnStatement *ReturnStatement) Execute(node *ast.Node) (string, error) {
	error := returnStatement.Validate(node.Children)
	if error != nil {
		return "", error
	}

	return node.Children[0].Value, nil
}
