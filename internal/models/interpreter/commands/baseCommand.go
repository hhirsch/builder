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

func (this *BaseCommand) GetResult() string {
	return this.result
}

func (this *BaseCommand) Execute(tokens []string) string {
	this.logger.Infof("Running %s", this.command)
	result := this.environment.Client.Execute(this.command)
	this.logger.Info(result)
	return result
}

func (this *BaseCommand) ExecuteOnLocalhost(tokens []string) string {
	parts := strings.Fields(this.command)
	if len(parts) == 0 {
		this.logger.Fatal("Command needs to be set.")
	}
	this.logger.Infof("Running %s on localhost.", parts[0])
	cmd := exec.Command(parts[0], parts[1:]...)

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	result := strings.TrimSpace(string(output))

	this.logger.Info(result)
	return result
}

func (this *BaseCommand) TestRequirements() bool {
	return true
}

func (this *BaseCommand) requireParameterAmount(tokens []string, requiredParameterAmount int) {
	if len(tokens) != requiredParameterAmount+1 {
		this.logger.Fatalf("%s needs %d parameters", this.commandName, requiredParameterAmount)
	}
}

func (this *BaseCommand) TrimResponseString(string string) string {
	return strings.TrimSpace(string)
}

func (this *BaseCommand) IsTrue(string string) bool {
	return strings.TrimSpace(string) == "true"
}

func (this *BaseCommand) FindBinary(binaryName string) bool {
	var command string = fmt.Sprintf("command -v %s >/dev/null 2>&1 && echo true || echo false", binaryName)
	if this.IsTrue(this.environment.Client.Execute(command)) {
		this.logger.Infof("Binary %s found.", binaryName)
		return true
	}

	this.logger.Fatalf("Unable to find required binary %s on target system.", binaryName)
	return false
}

func (this *BaseCommand) Undo() {
	this.logger.Info("Undo not possible.")
}

func (this *BaseCommand) GetName() string {
	return this.name
}

func (this *BaseCommand) GetDescription() string {
	return this.description
}

func (this *BaseCommand) GetBrief() string {
	return this.brief
}

func (this *BaseCommand) GetHelp() string {
	return this.help
}

func (this *BaseCommand) RequiresConnection() bool {
	return this.requiresConnection
}
