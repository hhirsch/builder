package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type EnsureExecutableCommand struct {
	environment *models.Environment
	description string
	BaseCommand
}

func NewEnsureExecutableCommand(environment *models.Environment) *EnsureExecutableCommand {
	return &EnsureExecutableCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment},
	}
}

func (this *EnsureExecutableCommand) Execute(tokens []string) {
	this.environment.GetLogger().Info("Ensuring target is executable.")
	if len(tokens) != 2 {
		this.environment.GetLogger().Fatal("ensureExecutable needs 1 parameter")
	}
	this.environment.Client.EnsureExecutable(tokens[1])
}

func (this *EnsureExecutableCommand) Undo() {
	this.environment.GetLogger().Info("Undoing ensureExecutable.")
}

func (this *EnsureExecutableCommand) GetDescription(tokens []string) string {
	return "make sure a binary is executable"
}

func (this *EnsureExecutableCommand) GetHelp() string {
	return "[binaryPath <string>]\tEnsure a binary is executable."
}
