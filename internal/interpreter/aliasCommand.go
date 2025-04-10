package interpreter

import (
	"errors"
	"github.com/hhirsch/builder/internal/models"
)

type AliasCommand struct {
	environment *models.Environment
	interpreter *Interpreter
	BaseCommand
}

func NewAliasCommand(interpreter *Interpreter, environment *models.Environment) *AliasCommand {
	controller := &AliasCommand{
		environment: environment,
		interpreter: interpreter,
		BaseCommand: BaseCommand{
			name: "alias",
		},
	}
	return controller
}

func (aliasCommand *AliasCommand) TestRequirement() bool {
	return true
}

func (aliasCommand *AliasCommand) Execute(tokens []string) (string, error) {
	if aliasCommand.interpreter.Aliases == nil {
		return "", errors.New("aliases uninitialized expected empty map got nil")
	}
	aliasCommand.interpreter.Aliases[tokens[1]] = tokens[2]
	return "", nil
}

func (aliasCommand *AliasCommand) GetDescription(tokens []string) string {
	return "Include an external file."
}

func (aliasCommand *AliasCommand) GetHelp() string {
	return "[include <file name>]\tIncludes the file as if it was typed out at the position of the include."
}
