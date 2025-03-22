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

func (this *ListFilesCommand) TestRequirements() bool {
	return this.FindBinary(this.command)
}

func (this *ListFilesCommand) GetDescription(tokens []string) string {
	return "List files in current directory."
}

func (this *ListFilesCommand) GetHelp() string {
	return "[]\tList files in current directory."
}
