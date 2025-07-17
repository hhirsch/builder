package command

import (
	"github.com/hhirsch/builder/internal/ast"
)

type StepCommand struct {
	printCommand *PrintCommand
	BaseCommand
}

func NewStepCommand(printCommand *PrintCommand) *StepCommand {
	return &StepCommand{printCommand: printCommand,
		BaseCommand: BaseCommand{
			name: "step",
		}}
}

func (stepCommand *StepCommand) TestRequirement() bool {
	return true
}

func (stepCommand *StepCommand) Validate(parameters []*ast.Node) error {
	return stepCommand.printCommand.Validate(parameters)
}

func (stepCommand *StepCommand) Execute(parameters []*ast.Node) (string, error) {
	error := stepCommand.Validate(parameters)
	if error != nil {
		return "", error
	}
	stepCommand.printCommand.Write("Step: " + stepCommand.GetStringFromParameters(parameters))
	return "", nil
}
