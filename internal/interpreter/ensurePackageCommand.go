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

func (ensurePackageCommand *EnsurePackageCommand) EnsurePackage(packageName string) (err error) {
	ensurePackageCommand.logger.Info("Checking status of package " + packageName)
	ensurePackageCommand.logger.Info("Status of " + packageName + " is not installed")
	_, err = ensurePackageCommand.Interpreter.System.Execute("dpkg --status " + packageName)
	if err != nil {
		return
	}
	ensurePackageCommand.logger.Info("Installing " + packageName)
	_, err = ensurePackageCommand.Interpreter.System.Execute("apt-get update")
	if err != nil {
		return
	}
	_, err = ensurePackageCommand.Interpreter.System.Execute("apt-get install " + packageName)
	if err != nil {
		return
	}
	return
}

func (ensurePackageCommand *EnsurePackageCommand) Execute(tokens []string) (string, error) {
	tokens = tokens[1:]
	parameters := strings.Join(tokens, " ")
	err := ensurePackageCommand.EnsurePackage(parameters)
	return "", err
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
