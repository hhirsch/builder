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
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (this *EnsureServiceCommand) TestRequirement() bool {
	return true
}

func (this *EnsureServiceCommand) Execute(tokens []string) string {
	if len(tokens) < 4 {
		this.environment.GetLogger().Fatal("ensureService needs 2 parameters and a description string")
	}
	reducedTokens := tokens[3:]
	description := strings.Join(reducedTokens, " ")
	this.environment.GetLogger().Info("Creating service name: " + tokens[1] + "  binary: " + tokens[2] + "  description: " + description)
	this.environment.Client.EnsureService(tokens[1], tokens[2], description)

	return ""
}

func (this *EnsureServiceCommand) Undo() {
	this.environment.GetLogger().Info("Undoing ensureService.")
}

func (this *EnsureServiceCommand) GetDescription(tokens []string) string {
	return "Create a system service from a binary"
}

func (this *EnsureServiceCommand) GetHelp() string {
	return "[serviceName <string>, binary <string>, description <string>]\tCreates a system service from a binary."
}
