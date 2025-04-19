package interpreter

import (
	"errors"
	format "fmt"
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type PrintCommand struct {
	environment *models.Environment
	interpreter *Interpreter
	BaseCommand
}

func NewPrintCommand(interpreter *Interpreter, environment *models.Environment) *PrintCommand {
	controller := &PrintCommand{
		environment: environment,
		interpreter: interpreter,
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

func (printCommand *PrintCommand) Execute(tokens []string) (string, error) {
	tokens = tokens[1:]
	if printCommand.interpreter.Variables == nil {
		return "", errors.New("variables uninitialized expected empty map got nil")
	}
	for index, variable := range tokens {
		if strings.HasPrefix(variable, "$") {
			variableName := strings.TrimPrefix(variable, "$")
			if foundVariable, isFoundVariable := printCommand.interpreter.Variables[variableName]; isFoundVariable {
				variable := foundVariable
				tokens[index], _ = variable.GetFlatString()
			}
		}
	}
	parameters := strings.Join(tokens, " ")
	format.Println(parameters)
	return "", nil
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
