package command

import (
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/token"
	//	"github.com/hhirsch/builder/internal/token"
	//"fmt"
	"log/slog"
)

type CustomCommand struct {
	buffer     []string
	body       []*ast.Node
	parameters []*ast.Node
	BaseCommand
}

func NewCustomCommand(functionNode *ast.Node) *CustomCommand {
	customCommand := &CustomCommand{
		body:        functionNode.Children,
		parameters:  functionNode.Parameters,
		BaseCommand: BaseCommand{},
	}
	slog.Debug("registering custom command.", slog.String("command name", customCommand.name))
	return customCommand
}

// parameters are the children of a statement node
func (customCommand *CustomCommand) Execute(parameters []*ast.Node) (string, error) {
	for _, parameter := range customCommand.parameters {
		switch parameter.Type {
		case token.IDENTIFIER:

		}
	}
	return "", nil
}

func (customCommand *CustomCommand) AppendToBuffer(line string) {
	customCommand.buffer = append(customCommand.buffer, line)
}

func (customCommand *CustomCommand) TestRequirement() bool {
	return true
}
