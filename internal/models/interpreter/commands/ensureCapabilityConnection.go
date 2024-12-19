package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type EnsureCapabilityConnectionCommand struct {
	environment *models.Environment
	description string
	BaseCommand
}

func NewEnsureCapabilityConnectionCommand(environment *models.Environment) *EnsureCapabilityConnectionCommand {
	controller := &EnsureCapabilityConnectionCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (this *EnsureCapabilityConnectionCommand) TestRequirement() bool {
	return true
}

func (this *EnsureCapabilityConnectionCommand) Execute(tokens []string) {
	if len(tokens) != 2 {
		this.environment.GetLogger().Fatal("ensureCapabilityConnection needs 1 parameters")
	}
	this.environment.Client.EnsureCapabilityConnection(tokens[1])
}

func (this *EnsureCapabilityConnectionCommand) Undo() {
	this.environment.GetLogger().Info("Nothing to undo.")
}

func (this *EnsureCapabilityConnectionCommand) GetDescription(tokens []string) string {
	return "Ensure a binary is allowed to open ports."
}

func (this *EnsureCapabilityConnectionCommand) GetHelp() string {
	return "[binaryPath <string>]\tEnsure a binary is allowed to open ports."
}
