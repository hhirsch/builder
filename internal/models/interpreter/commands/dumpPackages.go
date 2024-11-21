package commands

import (
	format "fmt"
	"github.com/hhirsch/builder/internal/models"
)

type DumpPackagesCommand struct {
	environment *models.Environment
	description string
	BaseCommand
}

func NewDumpPackagesCommand(environment *models.Environment) *DumpPackagesCommand {
	controller := &DumpPackagesCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (this *DumpPackagesCommand) Execute(tokens []string) {
	format.Println("dump")
	this.environment.Client.DumpPackages()
}

func (this *DumpPackagesCommand) Undo() {
	this.environment.GetLogger().Info("Undoing ensureExecutable.")
}

func (this *DumpPackagesCommand) GetDescription(tokens []string) string {
	return "dump packages deployed on the target system"
}

func (this *DumpPackagesCommand) GetHelp() string {
	return "[]\tDump a list of packages deployed on the target system."
}
