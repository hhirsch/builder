package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type PushFileCommand struct {
	environment *models.Environment
	description string
}

func NewPushFileCommand(environment *models.Environment) *PushFileCommand {
	controller := &PushFileCommand{
		environment: environment,
	}
	return controller
}

func (this *PushFileCommand) Execute(tokens []string) {
	if len(tokens) != 3 {
		this.environment.GetLogger().Fatal("pushFile needs 2 parameters")
	}
	this.environment.Client.PushFile(tokens[1], tokens[2])

}

func (this *PushFileCommand) Undo() {
	this.environment.GetLogger().Info("Nothing to undo.")
}

func (this *PushFileCommand) GetDescription(tokens []string) string {
	return "Ensure a binary is allowed to open ports."
}

func (this *PushFileCommand) GetHelp() string {
	return "[source <string>, target <string>]\tPush a file to the server."
}