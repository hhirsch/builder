package commands

import (
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type StepCommand struct {
	environment *models.Environment
	description string
	BaseCommand
}

func NewStepCommand(environment *models.Environment) *StepCommand {
	controller := &StepCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment: environment,
			name:        "step",
		},
	}
	return controller
}

func (this *StepCommand) Execute(tokens []string) string {
	this.description = strings.Join(tokens, " ")
	this.environment.GetLogger().Info(this.description)
	return ""
}

func (this *StepCommand) Undo() {
	this.environment.GetLogger().Info("Undoing step \"" + this.description + "\"")
}

func (this *StepCommand) GetDescription(tokens []string) string {
	return "Prints " + this.description + " on screen and logs it to file."
}

func (this *StepCommand) GetHelp() string {
	return "[step <string>]\tPrints description on screen and logs it to file."
}
