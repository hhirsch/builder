package command

import (
	format "fmt"
	"strings"
)

type PrintCommand struct {
	BaseCommand
}

func NewPrintCommand() *PrintCommand {
	controller := &PrintCommand{
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
	tokens = tokens[1:]
	parameters := strings.Join(tokens, " ")
	return parameters
}

func (printCommand *PrintCommand) Execute(tokens []string) (string, error) {
	format.Println(printCommand.getStringFromTokens(tokens))
	return "", nil // this is a token sink and has no return value
}
