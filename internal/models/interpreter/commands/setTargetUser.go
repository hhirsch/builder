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

func (setTargetUserCommand *SetTargetUserCommand) Execute(tokens []string) string {
	setTargetUserCommand.environment.GetLogger().Info("Setting target user.")
	if len(tokens) != 2 {
		setTargetUserCommand.environment.GetLogger().Fatal("setTargetUser needs 1 parameter")
	}
	setTargetUserCommand.environment.Client.SetTargetUser(tokens[1])
	return ""
}

func (setTargetUserCommand *SetTargetUserCommand) Undo() {
	setTargetUserCommand.environment.GetLogger().Info("Nothing to undo.")
}

func (setTargetUserCommand *SetTargetUserCommand) GetDescription(tokens []string) string {
	return "Ensure a binary is allowed to open ports."
}

func (setTargetUserCommand *SetTargetUserCommand) GetHelp() string {
	return "[binaryPath <string>]\tEnsure a binary is allowed to open ports."
}
