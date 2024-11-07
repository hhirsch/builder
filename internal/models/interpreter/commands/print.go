package commands

import (
	format "fmt"
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type PrintCommand struct {
	environment *models.Environment
	text        string
}

func NewPrintCommand(environment *models.Environment) *PrintCommand {
	controller := &PrintCommand{
		environment: environment,
	}
	return controller
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
