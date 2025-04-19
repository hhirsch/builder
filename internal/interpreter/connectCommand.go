package interpreter

import (
	"errors"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/melbahja/goph"
	"os/user"
)

type ConnectCommand struct {
	logger      *helpers.Logger
	interpreter *Interpreter
	BaseCommand
}

func NewConnectCommand(interpreter *Interpreter) *ConnectCommand {
	controller := &ConnectCommand{
		logger:      interpreter.logger,
		interpreter: interpreter,
		BaseCommand: BaseCommand{
			name:        "connect",
			description: "Connect to a host. Only supports key auth.",
			brief:       "[binaryPath <string>]\tEnsure a binary is executable.",
			parameters:  1,
			Interpreter: interpreter,
		},
	}
	return controller
}

func (connectCommand *ConnectCommand) Execute(tokens []string) (string, error) {
	connectCommand.requireParameterAmount(tokens, connectCommand.parameters)
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	keyPath := currentUser.HomeDir + "/.ssh/id_rsa"
	auth, err := goph.Key(keyPath, "")
	if err != nil {
		return "", err
	}

	domain := tokens[1]
	if foundAlias, isFoundAlias := connectCommand.interpreter.Aliases[tokens[1]]; isFoundAlias {
		domain = foundAlias
	}
	if domain == "localhost" {
		connectCommand.interpreter.Client = nil
		connectCommand.interpreter.System = NewLocalhost()
		return "", nil
	}

	sshClient, err := goph.New("root", domain, auth)
	if err != nil {
		return "", err
	}
	connectCommand.interpreter.Client = sshClient
	if connectCommand.interpreter.Client == nil {
		return "", errors.New("client nil right after being set")
	}
	connectCommand.interpreter.System = NewServer(sshClient)
	connectCommand.logger.Info("Connected to " + tokens[1])
	return "true", nil
}

func (connectCommand *ConnectCommand) GetDescription(tokens []string) string {
	return "Connect to a host. Only supports key auth."
}

func (connectCommand *ConnectCommand) GetHelp() string {
	return "(hostName <string>)\tConnect to host."
}
