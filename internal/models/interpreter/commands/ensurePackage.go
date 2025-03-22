package commands

import (
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type EnsurePackageCommand struct {
	BaseCommand
	environment *models.Environment
}

func NewEnsurePackageCommand(environment *models.Environment) *EnsurePackageCommand {
	controller := &EnsurePackageCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "ensurePackage",
			requiresConnection: true,
		},
	}
	return controller
}

func (this *EnsurePackageCommand) Execute(tokens []string) string {
	tokens = tokens[1:]
	parameters := strings.Join(tokens, " ")
	this.environment.Client.EnsurePackage(parameters)
	return ""
}

func (this *EnsurePackageCommand) Undo() {
	this.environment.GetLogger().Info("Undoing ensurePackage.")
}

func (this *EnsurePackageCommand) GetDescription(tokens []string) string {
	return "Create a system service from a binary"
}

func (this *EnsurePackageCommand) GetHelp() string {
	return "[packageName <string>]\tInstalls a package"
}
