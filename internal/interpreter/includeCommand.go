package interpreter

import (
	"errors"
	"github.com/hhirsch/builder/internal/models"
	"slices"
)

type IncludeCommand struct {
	interpreter *Interpreter
	BaseCommand
}

func NewIncludeCommand(interpreter *Interpreter, environment *models.Environment) *IncludeCommand {
	controller := &IncludeCommand{
		interpreter: interpreter,
		BaseCommand: BaseCommand{
			environment: environment,
			name:        "include",
			commandName: "",
			logger:      environment.GetLogger(),
			parameters:  1,
		},
	}
	return controller
}

func (includeCommand *IncludeCommand) TestRequirement() bool {
	return true
}

func (includeCommand *IncludeCommand) Execute(tokens []string) (string, error) {
	includeCommand.requireParameterAmount(tokens, includeCommand.parameters)
	if slices.Contains(includeCommand.interpreter.includes, tokens[1]) {
		return "", errors.New("circular include " + tokens[1])
	}
	return "", includeCommand.interpreter.Run(tokens[1])
}

func (includeCommand *IncludeCommand) GetDescription(tokens []string) string {
	return "Include an external file."
}

func (includeCommand *IncludeCommand) GetHelp() string {
	return "[include <file name>]\tIncludes the file as if it was typed out at the position of the include."
}
