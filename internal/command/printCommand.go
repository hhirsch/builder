package command

import (
	"errors"
	"github.com/hhirsch/builder/internal/ast"
	"strings"
)

type PrintCommand struct {
	writer *Writer
	BaseCommand
}

func NewPrintCommand(writer *Writer) *PrintCommand {
	controller := &PrintCommand{
		writer: writer,
		BaseCommand: BaseCommand{
			name: "print",
		},
	}
	return controller
}

func (printCommand *PrintCommand) TestRequirement() bool {
	return true
}

func (printCommand *PrintCommand) getStringFromTokens(tokens []string) string {
	tokens = tokens[1 : len(tokens)-2]
	parameters := strings.Join(tokens, " ")
	return parameters
}

func (printCommand *PrintCommand) Validate(parameters []*ast.Node) error {
	if len(parameters) == 0 {
		return errors.New("Print command needs at least one parameter.")
	}
	return nil
}

func (printCommand *PrintCommand) Write(message string) {
	writer := *printCommand.writer
	writer.Write(message)

}

func (printCommand *PrintCommand) Execute(parameters []*ast.Node) (string, error) {
	error := printCommand.Validate(parameters)
	if error != nil {
		return "", error
	}
	//writer := *printCommand.writer
	printCommand.Write(printCommand.GetStringFromParameters(parameters))
	//	writer.Write(printCommand.GetStringFromParameters(parameters))
	return "", nil // this is a token sink and has no return value
}
