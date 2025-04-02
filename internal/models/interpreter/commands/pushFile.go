package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type PushFileCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewPushFileCommand(environment *models.Environment) *PushFileCommand {
	return &PushFileCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "pushFile",
			requiresConnection: true,
		},
	}
}

func (pushFileCommand *PushFileCommand) Execute(tokens []string) string {
	if len(tokens) != 3 {
		pushFileCommand.environment.GetLogger().Fatal("pushFile needs 2 parameters.")
	}
	pushFileCommand.environment.Client.PushFile(tokens[1], tokens[2])
	return ""
}

func (pushFileCommand *PushFileCommand) Undo() {
	pushFileCommand.environment.GetLogger().Info("Nothing to undo.")
}

func (pushFileCommand *PushFileCommand) GetDescription(tokens []string) string {
	return "Push a file to the server."
}

func (pushFileCommand *PushFileCommand) GetHelp() string {
	return "[source <string>, target <string>]\tPush a file to the server."
}
