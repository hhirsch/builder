package commands

import (
	format "fmt"
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type PrintCommand struct {
	environment *models.Environment
	text        string
	BaseCommand
}

func NewPrintCommand(environment *models.Environment) *PrintCommand {
	controller := &PrintCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (this *PrintCommand) TestRequirement() bool {
	return true
}

func (this *PrintCommand) Execute(tokens []string) {
	tokens = tokens[1:]
	parameters := strings.Join(tokens, " ")
	format.Println(parameters)
	return
}

func (this *PrintCommand) Undo() {
	this.environment.GetLogger().Info("Nothing to undo for printing")
}

func (this *PrintCommand) GetDescription(tokens []string) string {
	return "Prints text on screen."
}

func (this *PrintCommand) GetHelp() string {
	return "[print <string>]\tPrints text on screen."
}
