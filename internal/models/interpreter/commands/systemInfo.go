package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type SystemInfoCommand struct {
	BaseCommand
	environment *models.Environment
	text        string
}

func NewSystemInfoCommand(environment *models.Environment) *SystemInfoCommand {
	return &SystemInfoCommand{
		environment: environment,
		BaseCommand: *NewBaseCommand(environment),
	}
}

func (this *SystemInfoCommand) TestRequirements() bool {
	return this.FindBinary("lsb_release")
}

func (this *SystemInfoCommand) Execute(tokens []string) {
	this.environment.GetLogger().Info("System Info Command", "System is "+this.TrimResponseString(this.environment.Client.Execute("lsb_release -ds")))
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
