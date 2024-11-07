package commands

import (
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type ExecuteAndPrintCommand struct {
	environment *models.Environment
	description string
}

func NewExecuteAndPrintCommand(environment *models.Environment) *ExecuteAndPrintCommand {
	controller := &ExecuteAndPrintCommand{
		environment: environment,
	}
	return controller
}

func (this *ExecuteAndPrintCommand) Execute(tokens []string) {
	tokens = tokens[1:]
	parameters := strings.Join(tokens, " ")
	this.environment.Client.ExecuteAndPrint(parameters)
}

func (this *ExecuteAndPrintCommand) Undo() {
	this.environment.GetLogger().Info("Undoing execute and print.")
}

func (this *ExecuteAndPrintCommand) GetDescription(tokens []string) string {
	return "make sure a binary is executable"
}

func (this *ExecuteAndPrintCommand) GetHelp() string {
	return "[binaryPath <string>]\tEnsure a binary is executable."
}
