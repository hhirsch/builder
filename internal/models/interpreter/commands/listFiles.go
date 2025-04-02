package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type ListFilesCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewListFilesCommand(environment *models.Environment) *ListFilesCommand {
	return &ListFilesCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment,
			logger:             environment.GetLogger(),
			command:            "ls",
			name:               "listFiles",
			requiresConnection: true,
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
