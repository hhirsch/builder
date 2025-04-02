package commands

import (
	format "fmt"
	"github.com/hhirsch/builder/internal/models"
)

type ListPackagesCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewListPackagesCommand(environment *models.Environment) *ListPackagesCommand {
	controller := &ListPackagesCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "listPackages",
			requiresConnection: true,
		},
	}
	return controller
}

func (listPackagesCommand *ListPackagesCommand) Execute(tokens []string) string {
	format.Println("list")
	listPackagesCommand.environment.Client.ListPackages()
	return ""
}

func (listPackagesCommand *ListPackagesCommand) Undo() {
	listPackagesCommand.environment.GetLogger().Info("tbd")
}

func (listPackagesCommand *ListPackagesCommand) GetDescription(tokens []string) string {
	return "tbd"
}

func (listPackagesCommand *ListPackagesCommand) GetHelp() string {
	return "tbd"
}
