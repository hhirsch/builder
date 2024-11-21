package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type SaveDatabaseCommand struct {
	environment *models.Environment
	description string
	BaseCommand
}

func NewSaveDatabaseCommand(environment *models.Environment) *SaveDatabaseCommand {
	controller := &SaveDatabaseCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (this *SaveDatabaseCommand) Execute(tokens []string) {
}

func (this *SaveDatabaseCommand) Undo() {
	this.environment.GetLogger().Info("Nothing to undo.")
}

func (this *SaveDatabaseCommand) GetDescription(tokens []string) string {
	return "Ensure a binary is allowed to open ports."
}

func (this *SaveDatabaseCommand) GetHelp() string {
	return "[source <string>, target <string>]\tPush a file to the server."
}
