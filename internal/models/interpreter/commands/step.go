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

func (stepCommand *StepCommand) Execute(tokens []string) string {
	stepCommand.description = strings.Join(tokens, " ")
	stepCommand.environment.GetLogger().Info(stepCommand.description)
	return ""
}

func (stepCommand *StepCommand) Undo() {
	stepCommand.environment.GetLogger().Info("Undoing step \"" + stepCommand.description + "\"")
}

func (stepCommand *StepCommand) GetDescription(tokens []string) string {
	return "Prints " + stepCommand.description + " on screen and logs it to file."
}

func (stepCommand *StepCommand) GetHelp() string {
	return "[step <string>]\tPrints description on screen and logs it to file."
}
