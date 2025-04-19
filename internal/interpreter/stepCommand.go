package interpreter

import (
	"fmt"
	"strings"

	"github.com/hhirsch/builder/internal/helpers"
)

type StepCommand struct {
	description string
	BaseCommand
}

func NewStepCommand(logger *helpers.Logger) *StepCommand {
	controller := &StepCommand{
		BaseCommand: BaseCommand{
			logger:             logger,
			name:               "step",
			requiresConnection: false,
		},
	}
	return controller
}

func (stepCommand *StepCommand) Execute(tokens []string) (string, error) {
	stepCommand.description = strings.Join(tokens, " ")
	stepCommand.logger.Info(stepCommand.description)
	fmt.Println("Step: " + strings.Join(tokens[1:], " "))
	return "", nil
}

func (stepCommand *StepCommand) GetDescription(tokens []string) string {
	return "Prints " + stepCommand.description + " on screen and logs it to file."
}

func (stepCommand *StepCommand) GetHelp() string {
	return "[step <string>]\tPrints description on screen and logs it to file."
}
