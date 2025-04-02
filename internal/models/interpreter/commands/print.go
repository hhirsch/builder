package commands

import (
	format "fmt"
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type PrintCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewPrintCommand(environment *models.Environment) *PrintCommand {
	controller := &PrintCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment: environment,
			name:        "print",
		},
	}
	return controller
}

func (printCommand *PrintCommand) TestRequirement() bool {
	return true
}

func (printCommand *PrintCommand) Execute(tokens []string) string {
	tokens = tokens[1:]
	parameters := strings.Join(tokens, " ")
	format.Println(parameters)
	return ""
}

func (printCommand *PrintCommand) Undo() {
	printCommand.environment.GetLogger().Info("Nothing to undo for printing")
}

func (printCommand *PrintCommand) GetDescription(tokens []string) string {
	return "Prints text on screen."
}

func (printCommand *PrintCommand) GetHelp() string {
	return "[print <string>]\tPrints text on screen."
}
