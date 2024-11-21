package commands

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

type ConnectCommand struct {
	environment *models.Environment
	description string
	logger      *helpers.Logger
	BaseCommand
}

func NewConnectCommand(environment *models.Environment) *ConnectCommand {
	controller := &ConnectCommand{
		environment: environment,
		logger:      environment.GetLogger(),
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (this *ConnectCommand) Execute(tokens []string) {
	if len(tokens) != 2 {
		this.logger.Fatal("connect needs 1 parameter")
	}
	this.environment.Client = *models.NewClient(this.environment, "root", tokens[1])
	this.logger.Info("Connected to " + tokens[1])
}

func (this *ConnectCommand) Undo() {
	this.environment.GetLogger().Info("Undo not implemented")
}

func (this *ConnectCommand) GetDescription(tokens []string) string {
	return "Connect to a host. Only supports key auth."
}

func (this *ConnectCommand) GetHelp() string {
	return "(hostName <string>)\tConnect to host."
}
