package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type EnsureCapabilityConnectionCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewEnsureCapabilityConnectionCommand(environment *models.Environment) *EnsureCapabilityConnectionCommand {
	controller := &EnsureCapabilityConnectionCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (ensureCapabilityConnectionCommand *EnsureCapabilityConnectionCommand) TestRequirement() bool {
	return true
}

func (ensureCapabilityConnectionCommand *EnsureCapabilityConnectionCommand) Execute(tokens []string) string {
	if len(tokens) != 2 {
		ensureCapabilityConnectionCommand.environment.GetLogger().Fatal("ensureCapabilityConnection needs 1 parameters")
	}
	ensureCapabilityConnectionCommand.environment.Client.EnsureCapabilityConnection(tokens[1])
	return ""
}

func (ensureCapabilityConnectionCommand *EnsureCapabilityConnectionCommand) Undo() {
	ensureCapabilityConnectionCommand.environment.GetLogger().Info("Nothing to undo.")
}

func (ensureCapabilityConnectionCommand *EnsureCapabilityConnectionCommand) GetDescription(tokens []string) string {
	return "Ensure a binary is allowed to open ports."
}

func (ensureCapabilityConnectionCommand *EnsureCapabilityConnectionCommand) GetHelp() string {
	return "[binaryPath <string>]\tEnsure a binary is allowed to open ports."
}
