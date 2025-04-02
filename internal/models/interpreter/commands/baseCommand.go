package commands

import (
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
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
	brief              string //information what the program is for
	help               string //detailed description of parameters with examples
	parameters         int
	requiresConnection bool
}

func NewBaseCommand(environment *models.Environment) *BaseCommand {
	return &BaseCommand{
		environment: environment,
		logger:      environment.GetLogger(),
		commands:    []string{},
		command:     "",
		commandName: "undefined",
		result:      "",
	}
}

func (baseCommand *BaseCommand) GetResult() string {
	return baseCommand.result
}

func (baseCommand *BaseCommand) Execute(tokens []string) string {
	baseCommand.logger.Infof("Running %s", baseCommand.command)
	result := baseCommand.environment.Client.Execute(baseCommand.command)
	baseCommand.logger.Info(result)
	return result
}

func (baseCommand *BaseCommand) ExecuteOnLocalhost(tokens []string) string {
	parts := strings.Fields(baseCommand.command)
	if len(parts) == 0 {
		baseCommand.logger.Fatal("Command needs to be set.")
	}
	baseCommand.logger.Infof("Running %s on localhost.", parts[0])
	cmd := exec.Command(parts[0], parts[1:]...)

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	result := strings.TrimSpace(string(output))

	baseCommand.logger.Info(result)
	return result
}

func (baseCommand *BaseCommand) TestRequirements() bool {
	return true
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
	if baseCommand.IsTrue(baseCommand.environment.Client.Execute(command)) {
		baseCommand.logger.Infof("Binary %s found.", binaryName)
		return true
	}

	baseCommand.logger.Fatalf("Unable to find required binary %s on target system.", binaryName)
	return false
}

func (baseCommand *BaseCommand) Undo() {
	baseCommand.logger.Info("Undo not possible.")
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
