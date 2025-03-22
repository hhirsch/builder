package commands

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

type ConnectCommand struct {
	environment *models.Environment
	logger      *helpers.Logger
	BaseCommand
}

func NewConnectCommand(environment *models.Environment) *ConnectCommand {
	controller := &ConnectCommand{
		environment: environment,
		logger:      environment.GetLogger(),
		BaseCommand: BaseCommand{
			environment: environment,
			name:        "connect",
			description: "Connect to a host. Only supports key auth.",
			brief:       "[binaryPath <string>]\tEnsure a binary is executable.",
		},
	}
	return controller
}

func (this *ConnectCommand) Execute(tokens []string) string {
	this.requireParameterAmount(tokens, 1)
	this.environment.Client = *models.NewClient(this.environment, "root", tokens[1])
	this.logger.Info("Connected to " + tokens[1])
	return "true"
}

func (this *ConnectCommand) GetDescription(tokens []string) string {
	return "Connect to a host. Only supports key auth."
}

func (this *ConnectCommand) GetHelp() string {
	return "(hostName <string>)\tConnect to host."
}
