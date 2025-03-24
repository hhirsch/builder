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

func (this *ListPackagesCommand) Execute(tokens []string) string {
	format.Println("list")
	this.environment.Client.ListPackages()
	return ""
}

func (this *ListPackagesCommand) Undo() {
	this.environment.GetLogger().Info("tbd")
}

func (this *ListPackagesCommand) GetDescription(tokens []string) string {
	return "tbd"
}

func (this *ListPackagesCommand) GetHelp() string {
	return "tbd"
}
