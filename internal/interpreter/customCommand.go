package interpreter

import (
	"strings"
)

type CustomCommand struct {
	interpreter        *Interpreter
	buffer             []string
	localVariables     map[string]Variable
	localVariableNames []string
	BaseCommand
}

func NewCustomCommand(interpreter *Interpreter, tokens []string) *CustomCommand {
	customCommand := &CustomCommand{
		localVariables: map[string]Variable{},
		interpreter:    interpreter,
		BaseCommand: BaseCommand{
			name:   tokens[1],
			logger: interpreter.logger,
		},
	}
	for _, variable := range tokens[2:] {
		customCommand.localVariableNames = append(customCommand.localVariableNames, strings.TrimPrefix(variable, "$"))
	}
	return customCommand
}

func (customCommand *CustomCommand) Execute(tokens []string) (string, error) {
	for index, variableName := range customCommand.localVariableNames {
		strippedTokens := tokens[1:]
		variableContent := strippedTokens[index]
		if strings.HasSuffix(variableName, "...") {
			variableName = strings.TrimSuffix(variableName, "...")
			variableContent = strings.Join(strippedTokens[index:], " ")
		}

		customCommand.logger.Debugf("variable name: %v", variableName)
		customCommand.localVariables[variableName] = *NewVariable(variableContent)
	}
	for _, line := range customCommand.buffer {
		tokens := strings.Fields(line)
		customCommand.logger.Debugf("replacing variables for line: %s", strings.Join(tokens, " "))
		processedTokens := customCommand.replaceVariablesInTokens(tokens, customCommand.localVariables)
		customCommand.logger.Debugf("line after replacing variables: %s", strings.Join(processedTokens, " "))
		err := customCommand.interpreter.handleLine(strings.Join(processedTokens, " "))
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func (customCommand *CustomCommand) AppendToBuffer(line string) {
	customCommand.buffer = append(customCommand.buffer, line)
}

func (customCommand *CustomCommand) TestRequirement() bool {
	return true
}
