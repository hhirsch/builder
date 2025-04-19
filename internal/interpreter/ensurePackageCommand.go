package interpreter

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type EnsurePackageCommand struct {
	BaseCommand
	environment *models.Environment
}

func NewEnsurePackageCommand(interpreter *Interpreter, logger *helpers.Logger) *EnsurePackageCommand {
	controller := &EnsurePackageCommand{
		BaseCommand: BaseCommand{
			logger:             logger,
			name:               "ensurePackage",
			requiresConnection: true,
			Interpreter:        interpreter,
		},
	}
	return controller
}

func (ensurePackageCommand *EnsurePackageCommand) EnsurePackage(packageName string) {
	ensurePackageCommand.logger.Info("Checking status of package " + packageName)
	ensurePackageCommand.logger.Info("Status of " + packageName + " is not installed")
	ensurePackageCommand.Interpreter.System.Execute("dpkg --status " + packageName)
	ensurePackageCommand.logger.Info("Installing " + packageName)
	ensurePackageCommand.Interpreter.System.Execute("apt-get update")
	ensurePackageCommand.Interpreter.System.Execute("apt-get install " + packageName)
}

func (ensurePackageCommand *EnsurePackageCommand) Execute(tokens []string) (string, error) {
	tokens = tokens[1:]
	parameters := strings.Join(tokens, " ")
	ensurePackageCommand.EnsurePackage(parameters)
	return "", nil
}

func (ensurePackageCommand *EnsurePackageCommand) Undo() {
	ensurePackageCommand.environment.GetLogger().Info("Undoing ensurePackage.")
}

func (ensurePackageCommand *EnsurePackageCommand) GetDescription(tokens []string) string {
	return "Create a system service from a binary"
}

func (ensurePackageCommand *EnsurePackageCommand) GetHelp() string {
	return "[packageName <string>]\tInstalls a package"
}
