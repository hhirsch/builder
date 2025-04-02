package commands

import (
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type ExecuteAndPrintCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewExecuteAndPrintCommand(environment *models.Environment) *ExecuteAndPrintCommand {
	return &ExecuteAndPrintCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "executeAndPrint",
			requiresConnection: true,
		},
	}
}

func (executeAndPrintCommand *ExecuteAndPrintCommand) getCommandFromTokens(tokens []string) string {
	tokens = tokens[1:]
	return strings.Join(tokens, " ")
}

func (executeAndPrintCommand *ExecuteAndPrintCommand) Execute(tokens []string) string {
	executeAndPrintCommand.environment.Client.ExecuteAndPrint(executeAndPrintCommand.getCommandFromTokens(tokens))
	return ""
}

func (executeAndPrintCommand *ExecuteAndPrintCommand) Undo() {
	executeAndPrintCommand.environment.GetLogger().Info("Undo unavailable for execute and print.")
}

func (executeAndPrintCommand *ExecuteAndPrintCommand) GetDescription(tokens []string) string {
	return "Execute " + executeAndPrintCommand.getCommandFromTokens(tokens) + " and print the output."
}

func (executeAndPrintCommand *ExecuteAndPrintCommand) GetHelp() string {
	return "[command <string>]\tExecute a command and print the output."
}
