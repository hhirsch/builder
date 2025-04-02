package commands

import (
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type EnsureServiceCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewEnsureServiceCommand(environment *models.Environment) *EnsureServiceCommand {
	controller := &EnsureServiceCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "ensureService",
			requiresConnection: true,
		},
	}
	return controller
}

func (ensureServiceCommand *EnsureServiceCommand) TestRequirement() bool {
	return true
}

func (ensureServiceCommand *EnsureServiceCommand) Execute(tokens []string) string {
	if len(tokens) < 4 {
		ensureServiceCommand.environment.GetLogger().Fatal("ensureService needs 2 parameters and a description string")
	}
	reducedTokens := tokens[3:]
	description := strings.Join(reducedTokens, " ")
	ensureServiceCommand.environment.GetLogger().Info("Creating service name: " + tokens[1] + "  binary: " + tokens[2] + "  description: " + description)
	ensureServiceCommand.environment.Client.EnsureService(tokens[1], tokens[2], description)

	return ""
}

func (ensureServiceCommand *EnsureServiceCommand) Undo() {
	ensureServiceCommand.environment.GetLogger().Info("Undoing ensureService.")
}

func (ensureServiceCommand *EnsureServiceCommand) GetDescription(tokens []string) string {
	return "Create a system service from a binary"
}

func (ensureServiceCommand *EnsureServiceCommand) GetHelp() string {
	return "[serviceName <string>, binary <string>, description <string>]\tCreates a system service from a binary."
}
