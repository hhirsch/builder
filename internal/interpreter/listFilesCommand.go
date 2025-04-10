package interpreter

import (
	"github.com/hhirsch/builder/internal/helpers"
)

type ListFilesCommand struct {
	BaseCommand
}

func NewListFilesCommand(interpreter *Interpreter, logger *helpers.Logger) *ListFilesCommand {
	return &ListFilesCommand{
		BaseCommand: BaseCommand{
			logger:             logger,
			command:            "ls",
			name:               "listFiles",
			requiresConnection: true,
			Interpreter:        interpreter,
		},
	}
}

func (listFilesCommand *ListFilesCommand) TestRequirements() bool {
	return listFilesCommand.FindBinary(listFilesCommand.command)
}

func (listFilesCommand *ListFilesCommand) GetDescription(tokens []string) string {
	return "List files in current directory."
}

func (listFilesCommand *ListFilesCommand) GetHelp() string {
	return "[]\tList files in current directory."
}
