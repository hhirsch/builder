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
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (this *SetupHostCommand) Execute(tokens []string) string {
	if len(tokens) != 2 {
		this.logger.Fatal("setupHost needs 1 parameter")
	}
	var err error
	var userName string
	var userPath = "host." + tokens[1] + ".user"
	userName, err = this.environment.GetRegistry().GetValue(userPath)
	if err != nil {
		this.logger.Info("No user found in registry asking for user input.")
		userName = "root"
		nameInput := huh.NewInput().
			Title("Enter user name").
			Value(&userName)
		err = nameInput.Run()
		if err != nil {
			this.logger.Fatalf("Error reading input for user name: %s", err.Error())
		}
		this.logger.Info("Registering " + userPath + " as " + userName)
		this.environment.GetRegistry().Register(userPath, userName)
	}
	if len(userName) == 0 {
		this.logger.Fatal("User name must not be empty!")
	}
	var hostName string
	var hostPath = "host." + tokens[1] + ".host"
	hostName, err = this.environment.GetRegistry().GetValue(hostPath)
	if err != nil {
		this.logger.Info("No host found in registry asking for user input.")
		hostInput := huh.NewInput().
			Title("Enter host name or IP").
			Value(&hostName)
		err = hostInput.Run()
		if err != nil {
			this.logger.Fatalf("Error reading input for host name: %s", err.Error())
		}
		this.logger.Info("Registering " + hostPath + " as " + hostName)
		this.environment.GetRegistry().Register(hostPath, hostName)
	}

	if len(hostName) == 0 {
		this.logger.Fatal("Host name must not be empty!")
	}

	err = this.environment.GetRegistry().Save()
	if err != nil {
		this.logger.Fatalf("Error saving registry: %s", err.Error())
	}
	this.environment.Client = *models.NewClient(this.environment, userName, hostName)
	return ""
}

func (this *SetupHostCommand) Undo() {
	this.environment.GetLogger().Info("Undo not implemented")
}

func (this *SetupHostCommand) GetDescription(tokens []string) string {
	return "Interactively setup your host. Only supports key auth."
}

func (this *SetupHostCommand) GetHelp() string {
	return "(host name <string>)\tPrompt for some information of the host to connect to."
}
