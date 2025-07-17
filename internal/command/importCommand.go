package command

import (
	"errors"
	"slices"

	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/interpreter"
	"github.com/hhirsch/builder/internal/token"
)

type ImportCommand struct {
	importedFiles []string
	interpreter   *interpreter.Interpreter
	BaseCommand
}

func NewImportCommand(interpreter *interpreter.Interpreter) *ImportCommand {
	controller := &ImportCommand{
		importedFiles: []string{},
		interpreter:   interpreter,
		BaseCommand: BaseCommand{
			name:       "import",
			parameters: 1,
		},
	}
	return controller
}

func (importCommand *ImportCommand) TestRequirement() bool {
	return true
}

func (importCommand *ImportCommand) Validate(parameters []*ast.Node) error {
	if len(parameters) > 0 {
		return errors.New("Import takes at least one parameter")
	}
	for _, parameter := range parameters {
		if parameter.Type != token.LITERAL {
			return errors.New("Import only takes literals as a parameter.")
		}
	}

	if slices.Contains(importCommand.importedFiles, parameters[0].Value) {
		return errors.New("circular import " + parameters[0].Value)
	}

	return nil
}

func (importCommand *ImportCommand) Execute(parameters []*ast.Node) (string, error) {
	error := importCommand.Validate(parameters)
	if error != nil {
		return "", error
	}
	//	return "", importCommand.interpreter.Run(tokens[1])
	return "", nil
}
