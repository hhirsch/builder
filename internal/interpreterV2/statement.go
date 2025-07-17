package interpreterV2

import (
	"fmt"
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/command"
)

type Statement struct {
	commands *command.Commands
}

func NewStatement(commands *command.Commands) *Statement {
	return &Statement{commands: commands}
}

func (statement *Statement) Execute(node *ast.Node) (string, error) {
	var error error
	var currentCommand *command.Command
	currentCommand, error = statement.commands.GetCommand(node.Value)
	if error != nil {
		return "", fmt.Errorf("Loading command: %w", error)
	}
	return (*currentCommand).Execute(node.Children)
}
