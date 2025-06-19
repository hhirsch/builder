package interpreter

import (
	"github.com/hhirsch/builder/internal/models"
)

type ExecuteCommand struct {
	environment *models.Environment
	system      System
	BaseCommand
}

func NewExecuteCommand(system System) *ExecuteCommand {
	return &ExecuteCommand{
		system: system,
		BaseCommand: BaseCommand{
			name:               "execute",
			requiresConnection: true,
		},
	}
}

func (executeCommand *ExecuteCommand) Execute(tokens []string) (string, error) {
	if len(tokens) != 1 {
		executeCommand.environment.GetLogger().Fatal("upload needs 1 parameter.")
	}
	return executeCommand.system.Execute(tokens[1])
}

func (executeCommand *ExecuteCommand) GetDescription(tokens []string) string {
	return "Execute code on the target system."
}

func (executeCommand *ExecuteCommand) GetHelp() string {
	return "[commend <string>]\tExecute code on the target system."
}
