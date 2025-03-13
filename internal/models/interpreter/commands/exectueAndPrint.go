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
		BaseCommand: BaseCommand{environment: environment},
	}
}

func (this *ExecuteAndPrintCommand) getCommandFromTokens(tokens []string) string {
	tokens = tokens[1:]
	return strings.Join(tokens, " ")
}

func (this *ExecuteAndPrintCommand) Execute(tokens []string) string {
	this.environment.Client.ExecuteAndPrint(this.getCommandFromTokens(tokens))
	return ""
}

func (this *ExecuteAndPrintCommand) Undo() {
	this.environment.GetLogger().Info("Undo unavailable for execute and print.")
}

func (this *ExecuteAndPrintCommand) GetDescription(tokens []string) string {
	return "Execute " + this.getCommandFromTokens(tokens) + " and print the output."
}

func (this *ExecuteAndPrintCommand) GetHelp() string {
	return "[command <string>]\tExecute a command and print the output."
}
