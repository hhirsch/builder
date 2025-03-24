package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type SetTargetUserCommand struct {
	BaseCommand
	environment *models.Environment
}

func NewSetTargetUserCommand(environment *models.Environment) *SetTargetUserCommand {
	controller := &SetTargetUserCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment,
			name: "setTargetUser",
		},
	}
	return controller
}

func (this *SetTargetUserCommand) Execute(tokens []string) string {
	this.environment.GetLogger().Info("Setting target user.")
	if len(tokens) != 2 {
		this.environment.GetLogger().Fatal("setTargetUser needs 1 parameter")
	}
	this.environment.Client.SetTargetUser(tokens[1])
	return ""
}

func (this *SetTargetUserCommand) Undo() {
	this.environment.GetLogger().Info("Nothing to undo.")
}

func (this *SetTargetUserCommand) GetDescription(tokens []string) string {
	return "Ensure a binary is allowed to open ports."
}

func (this *SetTargetUserCommand) GetHelp() string {
	return "[binaryPath <string>]\tEnsure a binary is allowed to open ports."
}
