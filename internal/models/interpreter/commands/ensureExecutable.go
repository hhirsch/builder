package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type EnsureExecutableCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewEnsureExecutableCommand(environment *models.Environment) *EnsureExecutableCommand {
	return &EnsureExecutableCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "ensureExecutable",
			description:        "Ensures a binary is executable.",
			brief:              "[binaryPath <string>]\tEnsure a binary is executable.",
			requiresConnection: true,
		},
	}
}

func (ensureExecutableCommand *EnsureExecutableCommand) Execute(tokens []string) string {
	ensureExecutableCommand.environment.GetLogger().Info("Ensuring target is executable.")
	if len(tokens) != 2 {
		ensureExecutableCommand.environment.GetLogger().Fatal("ensureExecutable needs 1 parameter")
	}
	ensureExecutableCommand.environment.Client.EnsureExecutable(tokens[1])
	return ""
}

func (ensureExecutableCommand *EnsureExecutableCommand) Undo() {
	ensureExecutableCommand.environment.GetLogger().Info("Undoing ensureExecutable.")
}

func (ensureExecutableCommand *EnsureExecutableCommand) GetDescription(tokens []string) string {
	return "make sure a binary is executable"
}

func (ensureExecutableCommand *EnsureExecutableCommand) GetHelp() string {
	return "[binaryPath <string>]\tEnsure a binary is executable."
}
