package command

import (
	"github.com/hhirsch/builder/internal/ast"
)

type ReturnCommand struct {
	BaseCommand
}

func NewReturnCommand(functionNode *ast.Node) *ReturnCommand {
	returnCommand := &ReturnCommand{
		BaseCommand: BaseCommand{},
	}
	return returnCommand
}

func (returnCommand *ReturnCommand) Execute(parameters []*ast.Node) (string, error) {

	return "", nil
}
