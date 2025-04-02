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
			parameters:  1,
		},
	}
	return controller
}

func (connectCommand *ConnectCommand) Execute(tokens []string) string {
	connectCommand.requireParameterAmount(tokens, 1)
	connectCommand.environment.Client = *models.NewClient(connectCommand.environment, "root", tokens[1])
	connectCommand.logger.Info("Connected to " + tokens[1])
	return "true"
}

func (connectCommand *ConnectCommand) GetDescription(tokens []string) string {
	return "Connect to a host. Only supports key auth."
}

func (connectCommand *ConnectCommand) GetHelp() string {
	return "(hostName <string>)\tConnect to host."
}
