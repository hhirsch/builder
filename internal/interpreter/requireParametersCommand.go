package interpreter

import (
	"errors"
	"github.com/hhirsch/builder/internal/models"
)

type RequireParametersCommand struct {
	environment *models.Environment
	interpreter *Interpreter
	BaseCommand
}

func NewRequireParametersCommand(interpreter *Interpreter, environment *models.Environment) *RequireParametersCommand {
	controller := &RequireParametersCommand{
		environment: environment,
		interpreter: interpreter,
		BaseCommand: BaseCommand{
			name: "requireParameters",
		},
	}
	return controller
}

func (requireParametersCommand *RequireParametersCommand) TestRequirement() bool {
	return true
}

func (requireParametersCommand *RequireParametersCommand) Execute(tokens []string) (string, error) {
	if string(len(requireParametersCommand.interpreter.BufferParameters)) == tokens[1] {
		return "", nil
	}
	return "", errors.New("invalid amount of buffer parameters")
}

func (requireParametersCommand *RequireParametersCommand) GetDescription(tokens []string) string {
	return "Include an external file."
}

func (requireParametersCommand *RequireParametersCommand) GetHelp() string {
	return "[include <file name>]\tIncludes the file as if it was typed out at the position of the include."
}
