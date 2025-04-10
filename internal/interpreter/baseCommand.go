package interpreter

import (
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/melbahja/goph"
	"os/exec"
	"strings"
)

type BaseCommand struct {
	environment        *models.Environment
	logger             *helpers.Logger
	commands           []string
	command            string
	commandName        string
	result             string
	name               string
	description        string //describes what the program will do if run
	brief              string //short description
	help               string //detailed description of parameters with examples
	parameters         int
	requiresConnection bool
	requirements       []string
	Interpreter        *Interpreter
}

func NewBaseCommand(environment *models.Environment) *BaseCommand {
	return &BaseCommand{
		environment:  environment,
		logger:       environment.GetLogger(),
		commands:     []string{},
		command:      "",
		commandName:  "undefined",
		result:       "",
		requirements: []string{}, // the binaries that need to be present on the target system
	}
}

func (baseCommand *BaseCommand) TestRequirements() bool {
	if baseCommand.requirements == nil {
		return true
	}
	for _, value := range baseCommand.requirements {
		if !baseCommand.FindBinary(value) {
			baseCommand.logger.Errorf("binary %s not present on the target system", value)
			return false
		}
	}
	baseCommand.logger.Info("all required binaries found on the target system")
	return true
}

func (baseCommand *BaseCommand) GetResult() string {
	return baseCommand.result
}

func (baseCommand *BaseCommand) Execute(tokens []string) (string, error) {
	baseCommand.logger.Infof("Running %s", baseCommand.command)
	result, _ := baseCommand.Interpreter.Client.Run(baseCommand.command)
	baseCommand.logger.Info(string(result))
	return string(result), nil
}

// this absolutely belongs into the client whenever the ssh client is nil
func (baseCommand *BaseCommand) ExecuteOnLocalhost(tokens []string) (string, error) {
	parts := strings.Fields(baseCommand.command)
	if len(parts) == 0 {
		baseCommand.logger.Fatal("Command needs to be set.")
	}
	baseCommand.logger.Infof("Running %s on localhost.", parts[0])
	cmd := exec.Command(parts[0], parts[1:]...)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(string(output))

	baseCommand.logger.Info(result)
	return result, nil
}

func (baseCommand *BaseCommand) GetClient() *goph.Client {
	if baseCommand.Interpreter.Client == nil {
		baseCommand.logger.Error("Client was nill when tried to get it")
	}
	return baseCommand.Interpreter.Client
}

func (baseCommand *BaseCommand) requireParameterAmount(tokens []string, requiredParameterAmount int) {
	if len(tokens) != requiredParameterAmount+1 {
		baseCommand.logger.Fatalf("%s needs %d parameters", baseCommand.commandName, requiredParameterAmount)
	}
}

func (baseCommand *BaseCommand) TrimResponseString(string string) string {
	return strings.TrimSpace(string)
}

func (baseCommand *BaseCommand) IsTrue(string string) bool {
	return strings.TrimSpace(string) == "true"
}

func (baseCommand *BaseCommand) FindBinary(binaryName string) bool {
	var command = fmt.Sprintf("command -v %s >/dev/null 2>&1 && echo true || echo false", binaryName)
	if baseCommand.Interpreter.Client == nil {
		baseCommand.logger.Error("Client is nil when trying to run FindBinary")
		return false
	}
	executionResult, _ := baseCommand.Interpreter.Client.Run(command)
	if baseCommand.IsTrue(string(executionResult)) {
		baseCommand.logger.Infof("Binary %s found.", binaryName)
		return true
	}

	baseCommand.logger.Fatalf("Unable to find required binary %s on target system.", binaryName)
	return false
}

func (baseCommand *BaseCommand) Undo() {
	baseCommand.logger.Infof("No undo available for %v", baseCommand.commandName)
}

func (baseCommand *BaseCommand) GetName() string {
	return baseCommand.name
}

func (baseCommand *BaseCommand) GetDescription() string {
	return baseCommand.description
}

func (baseCommand *BaseCommand) GetBrief() string {
	return baseCommand.brief
}

func (baseCommand *BaseCommand) GetHelp() string {
	return baseCommand.help
}

func (baseCommand *BaseCommand) RequiresConnection() bool {
	return baseCommand.requiresConnection
}
