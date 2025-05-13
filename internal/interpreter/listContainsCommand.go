package interpreter

import (
	"github.com/hhirsch/builder/internal/models"
	"slices"
)

type ListContainsCommand struct {
	environment *models.Environment
	interpreter *Interpreter
	BaseCommand
}

func NewListContainsCommand(interpreter *Interpreter, environment *models.Environment) *ListContainsCommand {
	controller := &ListContainsCommand{
		environment: environment,
		interpreter: interpreter,
		BaseCommand: BaseCommand{
			environment: environment,
			name:        "listContains",
		},
	}
	return controller
}

func (listContainsCommand *ListContainsCommand) TestRequirement() bool {
	return true
}

func (listContainsCommand *ListContainsCommand) Execute(tokens []string) (string, error) {
	tokens = tokens[1:]
	tokens = listContainsCommand.replaceVariablesInTokens(tokens, listContainsCommand.interpreter.Variables)
	needle := tokens[0]
	haystack := tokens[1:]
	if slices.Contains(haystack, needle) {
		return "true", nil
	}
	return "false", nil
}

func (listContainsCommand *ListContainsCommand) Undo() {
	listContainsCommand.environment.GetLogger().Info("Nothing to undo for printing")
}

func (listContainsCommand *ListContainsCommand) GetDescription(tokens []string) string {
	return "Check if a list of strings contains a string."
}

func (listContainsCommand *ListContainsCommand) GetHelp() string {
	return "[listContains <needle> <haystack>]\tReturns true if the string is in the list."
}
