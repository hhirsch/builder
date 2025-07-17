package command

import (
	"errors"
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/token"
	"github.com/hhirsch/builder/internal/variable"

	//	"github.com/hhirsch/builder/internal/token"
	"fmt"
	"log/slog"
)

type CustomCommand struct {
	buffer               []string
	body                 []*ast.Node
	parameterDefinitions []*ast.Node
	BaseCommand
}

func NewCustomCommand(functionNode *ast.Node) *CustomCommand {

	customCommand := &CustomCommand{
		body:                 functionNode.Children,
		parameterDefinitions: functionNode.Parameters,
		BaseCommand:          BaseCommand{},
	}
	slog.Debug("registering custom command.", slog.String("command name", customCommand.name))
	return customCommand
}

func (customCommand *CustomCommand) Validate(parameters []*ast.Node) error {
	if len(parameters) < len(customCommand.parameterDefinitions) {
		return errors.New("Less parameters than the parameter definition requires.")
	}
	for index, parameter := range customCommand.parameterDefinitions {
		if len(parameters) < (index + 1) {
			switch parameter.Type {
			case token.IDENTIFIER_VARIADIC:
				return errors.New("Variadic parameter definition requires at least one parameter.")
			case token.IDENTIFIER:
				return errors.New("Parameter definition requires an identifier counterpart to parameter.")
			case token.LITERAL:
				return errors.New("Parameter definition requires a literal counterpart to parameter.")
			}
		}
	}
	return nil
}

func (customCommand *CustomCommand) GetVariablesFromParameters(parameters []*ast.Node) (*variable.VariablePool, error) {
	variables := variable.NewVariablePool()
	error := customCommand.Validate(parameters)
	if error != nil {
		return nil, error
	}
	for index, parameterDefinition := range customCommand.parameterDefinitions {
		switch parameterDefinition.Type {
		case token.IDENTIFIER_VARIADIC:
			var concatenatedVariadicValue string = ""
			for _, variadicParameter := range parameters[index:] {
				if concatenatedVariadicValue != "" {
					concatenatedVariadicValue = concatenatedVariadicValue + " "
				}
				concatenatedVariadicValue += variadicParameter.Value
			}
			variables.SetVariable(parameterDefinition.Value, concatenatedVariadicValue)
		case token.IDENTIFIER:
			variables.SetVariable(parameterDefinition.Value, parameters[index].Value)
		default:
			return nil, fmt.Errorf("Unrecognized parameter type in definition: %v", parameterDefinition.Type)
		}
	}
	return variables, nil
}

// parameters are the children of a statement node
func (customCommand *CustomCommand) Execute(parameters []*ast.Node) (string, error) {
	//	return "", customCommand.PopulateLocalVariablesFromParameters(parameters)
	return "", nil
	// the custom command just wants to evaluate its code.
	// it needs to be able to run an interpreter with its own local variables as a parameter
	// the interpreter does not need to open a file
	// the interpreter needs an eval function that accepts local variables being passed
}

func (customCommand *CustomCommand) AppendToBuffer(line string) {
	customCommand.buffer = append(customCommand.buffer, line)
}

func (customCommand *CustomCommand) TestRequirement() bool {
	return true
}
