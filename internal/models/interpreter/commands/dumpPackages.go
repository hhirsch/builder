package commands

import (
	format "fmt"
	"github.com/hhirsch/builder/internal/models"
	"os"
	"time"
)

type DumpPackagesCommand struct {
	environment *models.Environment
	BaseCommand
}

func NewDumpPackagesCommand(environment *models.Environment) *DumpPackagesCommand {
	return &DumpPackagesCommand{
		environment: environment,
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "dumpPackages",
			requiresConnection: true,
		},
	}
}

func (dumpPackagesCommand *DumpPackagesCommand) Execute(tokens []string) string {
	dumpPackagesCommand.environment.GetLogger().Info("Dumping Packages")
	currentTime := time.Now()
	fileName := "snapshots/" + currentTime.Format("02-01-2006_15-04-05") + ".dmp" // File name format: DD-MM-YYYY_HH-MM-SS

	err := os.WriteFile(fileName, []byte(dumpPackagesCommand.environment.Client.Execute("dpkg --get-selections")), 0644)
	if err != nil {
		dumpPackagesCommand.environment.GetLogger().Fatal(format.Printf("Error writing file: %v", err))
	}

	dumpPackagesCommand.environment.GetLogger().Info("File " + fileName + " created and string written successfully!\n")
	return ""
}

func (dumpPackagesCommand *DumpPackagesCommand) Undo() {
	dumpPackagesCommand.environment.GetLogger().Info("Undoing ensureExecutable.")
}

func (dumpPackagesCommand *DumpPackagesCommand) GetDescription(tokens []string) string {
	return "dump packages deployed on the target system"
}

func (dumpPackagesCommand *DumpPackagesCommand) GetHelp() string {
	return "[]\tDump a list of packages deployed on the target system."
}
