package interpreter

import (
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
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

func (baseCommand *BaseCommand) Execute(tokens []string) (string, error) {
	baseCommand.logger.Infof("Running %s", baseCommand.command)
	result, err := baseCommand.Interpreter.System.Execute(baseCommand.command)
	if err != nil {
		return "", err
	}
	baseCommand.logger.Debugf("Command %s gave us the following result:\n%s", baseCommand.command, string(result))
	return string(result), nil
}

func (baseCommand *BaseCommand) requireParameterAmount(tokens []string, requiredParameterAmount int) {
	if len(tokens) != requiredParameterAmount+1 {
		baseCommand.logger.Fatalf("%s needs %d parameters", baseCommand.commandName, requiredParameterAmount)
	}
}

// skips the first token
func (baseCommand *BaseCommand) replaceVariablesInTokens(tokens []string, variables map[string]Variable) []string {
	for index := 1; index < len(tokens); index++ {
		variable := tokens[index]
		if strings.HasPrefix(variable, "$") {
			baseCommand.logger.Debugf("Detected $ prefix for: %s", variable)
			variableName := strings.TrimPrefix(variable, "$")
			if foundVariable, isFoundVariable := variables[variableName]; isFoundVariable {
				variable := foundVariable
				tokens[index], _ = variable.GetFlatString()
			}
		}
	}
	return tokens
}

func (baseCommand *BaseCommand) TrimResponseString(string string) string {
	return strings.TrimSpace(string)
}

func (baseCommand *BaseCommand) IsTrue(string string) bool {
	return strings.TrimSpace(string) == "true"
}

func (baseCommand *BaseCommand) FindBinary(binaryName string) bool {
	var command = fmt.Sprintf("command -v %s >/dev/null 2>&1 && echo true || echo false", binaryName)
	executionResult, err := baseCommand.Interpreter.System.Execute(command)
	if err != nil {
		baseCommand.logger.Infof("Error during execution:  %s", err.Error())
	}

	if baseCommand.IsTrue(string(executionResult)) {
		baseCommand.logger.Infof("Binary %s found.", binaryName)
		return true
	}

	return false
}

func (baseCommand *BaseCommand) Undo() {
	baseCommand.logger.Infof("No undo available for %v", baseCommand.commandName)
}

func (baseCommand *BaseCommand) GetName() string {
	return baseCommand.name
}

func (baseCommand *BaseCommand) GetDescription(tokens []string) string {
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
