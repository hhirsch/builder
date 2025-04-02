package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type SaveDatabaseCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewSaveDatabaseCommand(environment *models.Environment) *SaveDatabaseCommand {
	controller := &SaveDatabaseCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (saveDatabaseCommand *SaveDatabaseCommand) Execute(tokens []string) {
}

func (saveDatabaseCommand *SaveDatabaseCommand) Undo() {
	saveDatabaseCommand.environment.GetLogger().Info("Nothing to undo.")
}

func (saveDatabaseCommand *SaveDatabaseCommand) GetDescription(tokens []string) string {
	return "Ensure a binary is allowed to open ports."
}

func (saveDatabaseCommand *SaveDatabaseCommand) GetHelp() string {
	return "[source <string>, target <string>]\tPush a file to the server."
}
