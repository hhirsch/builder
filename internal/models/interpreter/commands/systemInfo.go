package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type SystemInfoCommand struct {
	BaseCommand
	environment *models.Environment
}

func NewSystemInfoCommand(environment *models.Environment) *SystemInfoCommand {
	return &SystemInfoCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment,
			logger:             environment.GetLogger(),
			name:               "systemInfo",
			requiresConnection: true,
		},
	}
}

func (systemInfoCommand *SystemInfoCommand) TestRequirements() bool {
	return systemInfoCommand.FindBinary("lsb_release")
}

func (systemInfoCommand *SystemInfoCommand) Execute(tokens []string) string {
	systemInfoCommand.environment.GetLogger().Infof("System is %s.", systemInfoCommand.TrimResponseString(systemInfoCommand.environment.Client.Execute("lsb_release -ds")))
	return ""
}

func (systemInfoCommand *SystemInfoCommand) Undo() {
	systemInfoCommand.environment.GetLogger().Info("Nothing to undo for printing")
}

func (systemInfoCommand *SystemInfoCommand) GetDescription(tokens []string) string {
	return "SystemInfos text on screen."
}

func (systemInfoCommand *SystemInfoCommand) GetHelp() string {
	return "[print <string>]\tSystemInfos text on screen."
}
