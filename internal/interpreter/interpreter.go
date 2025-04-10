package interpreter

import (
	"bufio"
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/melbahja/goph"
	"os"
	"slices"
	"strings"
)

type Interpreter struct {
	logger              *helpers.Logger
	environment         *models.Environment
	registry            *models.Registry
	commands            map[string]Command
	Variables           map[string]Variable
	checkedRequirements []string
	includes            []string
	Client              *goph.Client
	Aliases             map[string]string
}

func NewInterpreter(environment *models.Environment) *Interpreter {
	logger := environment.GetLogger()
	registry := models.NewRegistry(environment.GetGlobalRegistryPath())
	err := registry.Load()
	if err != nil {
		logger.Fatalf("unable to load registry: %s", err.Error())
	}
	variables := map[string]Variable{}
	commands := map[string]Command{}
	aliases := map[string]string{}
	interpreter := &Interpreter{
		logger:      logger,
		registry:    registry,
		environment: environment,
		commands:    commands,
		Variables:   variables,
		Aliases:     aliases,
	}
	interpreter.AddCommand(NewConnectCommand(interpreter))
	interpreter.AddCommand(NewStepCommand(logger))
	interpreter.AddCommand(NewPrintCommand(interpreter, environment))
	interpreter.AddCommand(NewListFilesCommand(interpreter, environment.GetLogger()))
	interpreter.AddCommand(NewIncludeCommand(interpreter, environment))
	interpreter.AddCommand(NewAliasCommand(interpreter, environment))
	return interpreter
}

func (interpreter *Interpreter) AddCommand(command Command) {
	interpreter.commands[command.GetName()] = command
}

func (interpreter *Interpreter) Run(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("can't open file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		interpreter.handleLine(line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("[functionName] unable to scan file: %w", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file %v", err)
		}
	}()
	return nil
}

func (interpreter *Interpreter) HasConnection() bool {
	return interpreter.Client != nil
}

func (interpreter *Interpreter) requireConnection() {
	if interpreter.HasConnection() { // if the client is not initialized we don't have a connection
		interpreter.logger.Fatal("Setup a host before using a command that requires a connection.")
	}
}

func (interpreter *Interpreter) handleCommandLine(tokens []string) string {
	commandName := tokens[0]

	if commandName == "connect" || commandName == "setupHost" {
		interpreter.checkedRequirements = []string{}
	}
	var command Command

	if foundCommand, isFoundCommand := interpreter.commands[commandName]; isFoundCommand {
		command = foundCommand
		if command.RequiresConnection() {
			interpreter.requireConnection()
		}
	}

	if command == nil {
		interpreter.logger.Fatalf("Invalid command %s.", commandName)
	}

	if slices.Contains(interpreter.checkedRequirements, commandName) {
		interpreter.logger.Debugf("Passed requirenments for %s. (cached)", commandName)
	} else if command.TestRequirements() {
		interpreter.checkedRequirements = append(interpreter.checkedRequirements, commandName)
	} else {
		interpreter.logger.Fatalf("Failed requirenments for %s.", commandName)
	}
	commandOutput, err := command.Execute(tokens)
	if err != nil {
		interpreter.logger.Errorf("error while executing command %v", err)
	}
	return commandOutput
}

func (interpreter *Interpreter) handleVariableLine(tokens []string) {
	variableName := strings.TrimPrefix(tokens[0], "$")
	//interpreter.variables[variableName] = interpreter.handleCommandLine(tokens[2:])
	interpreter.Variables[variableName] = *NewVariable(interpreter.handleCommandLine(tokens[2:]))
}

func (interpreter *Interpreter) handleLine(input string) {
	tokens := strings.Fields(input)
	if strings.HasPrefix(tokens[0], "//") {
		return
	}

	if strings.HasPrefix(tokens[0], "$") && tokens[1] == "=" {
		interpreter.handleVariableLine(tokens)
		return
	}

	interpreter.handleCommandLine(tokens)
}

func (interpreter *Interpreter) GetReferencePage() string {
	return ""
}
