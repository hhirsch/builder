package command

import (
	"github.com/hhirsch/builder/internal/ast"
)

type BaseCommand struct {
	commandName        string
	result             string
	name               string
	parameters         int
	requiresConnection bool
	requirements       []string
}

func NewBaseCommand() *BaseCommand {
	return &BaseCommand{
		commandName:  "undefined",
		result:       "",
		requirements: []string{}, // the binaries that need to be present on the target system
	}
}

func (baseCommand *BaseCommand) TestRequirements() bool {
	return true
}

func (baseCommand *BaseCommand) GetName() string {
	return baseCommand.name
}

func (baseCommand *BaseCommand) RequiresConnection() bool {
	return baseCommand.requiresConnection
}

func (baseCommand *BaseCommand) GetStringFromParameters(parameters []*ast.Node) string {
	var parameterString string
	var parameterLength = len(parameters)
	for index, node := range parameters {
		parameterString += node.Value
		if parameterLength != index+1 {
			parameterString += " "
		}
	}
	return parameterString
}
