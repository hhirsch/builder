package interpreter

import (
	"github.com/hhirsch/builder/internal/models"
)

type UploadCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewUploadCommand(environment *models.Environment) *UploadCommand {
	return &UploadCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "upload",
			requiresConnection: true,
		},
	}
}

func (uploadCommand *UploadCommand) Execute(tokens []string) (string, error) {
	if len(tokens) != 3 {
		uploadCommand.environment.GetLogger().Fatal("upload needs 2 parameters.")
	}
	uploadCommand.environment.Client.Upload(tokens[1], tokens[2])
	return "", nil
}

func (uploadCommand *UploadCommand) GetDescription(tokens []string) string {
	return "Push a file to the server."
}

func (uploadCommand *UploadCommand) GetHelp() string {
	return "[source <string>, target <string>]\tPush a file to the server."
}
