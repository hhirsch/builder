package commands

import (
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

func Step(environment *models.Environment, tokens []string) {
	tokens = tokens[1:]
	parameters := strings.Join(tokens, " ")
	environment.GetLogger().Info(parameters)
}

type StepCommand struct {
	environment *models.Environment
	description string
}

func NewStep(environment *models.Environment, tokens []string) *StepCommand {
	tokens = tokens[1:]
	controller := &StepCommand{
		environment: environment,
		description: strings.Join(tokens, " "),
	}
	return controller
}

func (this *StepCommand) Execute() {
	this.environment.GetLogger().Info(this.description)
}

func (this *StepCommand) Undo() {
	this.environment.GetLogger().Info("Undoing step \"" + this.description + "\"")
}

func (this *StepCommand) Describe() {
	this.environment.GetLogger().Info("Prints " + this.description + " on screen and logs it to file.")
}

func (this *StepCommand) Help() {
	this.environment.GetLogger().Info("Prints description on screen and logs it to file.")
}
