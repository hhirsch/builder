package commands

import (
	"github.com/charmbracelet/huh"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

type SetupHostCommand struct {
	environment *models.Environment
	logger      *helpers.Logger
	BaseCommand
}

func NewSetupHostCommand(environment *models.Environment) *SetupHostCommand {
	controller := &SetupHostCommand{
		environment: environment,
		logger:      environment.GetLogger(),
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "setupHost",
			requiresConnection: false,
			parameters:         2,
		},
	}
	return controller
}

func (setupHostCommand *SetupHostCommand) Execute(tokens []string) string {
	if len(tokens) != 2 {
		setupHostCommand.logger.Fatal("setupHost needs 1 parameter")
	}
	var err error
	var userName string
	var userPath = "host." + tokens[1] + ".user"
	userName, err = setupHostCommand.environment.GetRegistry().GetValue(userPath)
	if err != nil {
		setupHostCommand.logger.Info("No user found in registry asking for user input.")
		userName = "root"
		nameInput := huh.NewInput().
			Title("Enter user name").
			Value(&userName)
		err = nameInput.Run()
		if err != nil {
			setupHostCommand.logger.Fatalf("Error reading input for user name: %s", err.Error())
		}
		setupHostCommand.logger.Info("Registering " + userPath + " as " + userName)
		setupHostCommand.environment.GetRegistry().Register(userPath, userName)
	}
	if len(userName) == 0 {
		setupHostCommand.logger.Fatal("User name must not be empty!")
	}
	var hostName string
	var hostPath = "host." + tokens[1] + ".host"
	hostName, err = setupHostCommand.environment.GetRegistry().GetValue(hostPath)
	if err != nil {
		setupHostCommand.logger.Info("No host found in registry asking for user input.")
		hostInput := huh.NewInput().
			Title("Enter host name or IP").
			Value(&hostName)
		err = hostInput.Run()
		if err != nil {
			setupHostCommand.logger.Fatalf("Error reading input for host name: %s", err.Error())
		}
		setupHostCommand.logger.Info("Registering " + hostPath + " as " + hostName)
		setupHostCommand.environment.GetRegistry().Register(hostPath, hostName)
	}

	if len(hostName) == 0 {
		setupHostCommand.logger.Fatal("Host name must not be empty!")
	}

	err = setupHostCommand.environment.GetRegistry().Save()
	if err != nil {
		setupHostCommand.logger.Fatalf("Error saving registry: %s", err.Error())
	}
	setupHostCommand.environment.Client = *models.NewClient(setupHostCommand.environment, userName, hostName)
	return ""
}

func (setupHostCommand *SetupHostCommand) Undo() {
	setupHostCommand.environment.GetLogger().Info("Undo not implemented")
}

func (setupHostCommand *SetupHostCommand) GetDescription(tokens []string) string {
	return "Interactively setup your host. Only supports key auth."
}

func (setupHostCommand *SetupHostCommand) GetHelp() string {
	return "(host name <string>)\tPrompt for some information of the host to connect to."
}
