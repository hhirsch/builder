package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type SystemInfoCommand struct {
	environment *models.Environment
	text        string
}

func NewSystemInfoCommand(environment *models.Environment) *SystemInfoCommand {
	controller := &SystemInfoCommand{
		environment: environment,
	}
	return controller
}

func (this *SystemInfoCommand) Execute(tokens []string) {
	this.environment.Client.ExecuteAndPrint("lsb_release -a")
	return
}

func (this *SystemInfoCommand) Undo() {
	this.environment.GetLogger().Info("Nothing to undo for printing")
}

func (this *SystemInfoCommand) GetDescription(tokens []string) string {
	return "SystemInfos text on screen."
}

func (this *SystemInfoCommand) GetHelp() string {
	return "[print <string>]\tSystemInfos text on screen."
}
