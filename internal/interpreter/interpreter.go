package interpreter

import (
	"bufio"
	"errors"
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
	System              System
	Aliases             map[string]string
	ReadingBufferActive bool
	ReadingBuffer       string
	BufferName          string
	FunctionPool        map[string]string
	BufferParameters    []string
	ConnectToLocalhost  bool
	customCommand       *CustomCommand
}

func NewInterpreter(environment *models.Environment) (*Interpreter, error) {
	logger := environment.GetLogger()
	registry := models.NewRegistry(environment.GetGlobalRegistryPath())
	err := registry.Load()
	if err != nil {
		return nil, fmt.Errorf("load registry: %w", err)
	}

	interpreter := &Interpreter{
		logger:              logger,
		registry:            registry,
		environment:         environment,
		commands:            map[string]Command{},
		Variables:           map[string]Variable{},
		Aliases:             map[string]string{},
		ReadingBufferActive: false,
		BufferName:          "",
		FunctionPool:        map[string]string{},
		ConnectToLocalhost:  false,
	}
	interpreter.AddCommand(NewConnectCommand(interpreter))
	interpreter.AddCommand(NewStepCommand(logger))
	interpreter.AddCommand(NewPrintCommand(interpreter, environment))
	interpreter.AddCommand(NewListFilesCommand(interpreter, logger))
	interpreter.AddCommand(NewIncludeCommand(interpreter, environment))
	interpreter.AddCommand(NewAliasCommand(interpreter, environment))
	interpreter.AddCommand(NewEnsurePackageCommand(interpreter, logger))
	return interpreter, nil
}

func (interpreter *Interpreter) AddCommand(command Command) {
	interpreter.commands[command.GetName()] = command
}

func (interpreter *Interpreter) Run(fileName string) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		err = interpreter.handleLine(line)
		if err != nil {
			return fmt.Errorf("handling line: %w", err)
		}
	}
	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("scanning file: %w", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			err = fmt.Errorf("closing file: %w", err)
		}
	}()
	return err
}

func (interpreter *Interpreter) HandleStringSlices(input []string) error {
	for _, line := range input {
		err := interpreter.handleLine(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func (interpreter *Interpreter) handleLine(input string) error {
	if strings.TrimSpace(input) == "" {
		interpreter.logger.Debugf("Skipped empty line in builder file.")
		return nil
	}
	tokens := strings.Fields(input)
	if strings.HasPrefix(tokens[0], "//") {
		return nil
	}

	if interpreter.ReadingBufferActive {
		if tokens[0] == "done" {
			interpreter.ReadingBufferActive = false
			interpreter.FunctionPool[interpreter.BufferName] = interpreter.ReadingBuffer
			interpreter.ReadingBuffer = ""
			interpreter.AddCommand(interpreter.customCommand)
			return nil
		}
		interpreter.ReadingBuffer += strings.Join(tokens, " ") + "\n"
		interpreter.customCommand.AppendToBuffer(strings.Join(tokens, " "))
		return nil
	}

	if tokens[0] == "function" {
		if len(tokens) < 2 {
			return errors.New("function needs 2 parameters")
		}
		interpreter.ReadingBufferActive = true
		interpreter.BufferName = tokens[1]
		interpreter.customCommand = NewCustomCommand(interpreter, tokens)
		return nil
	}

	if strings.HasPrefix(tokens[0], "$") && tokens[1] == "=" {
		return interpreter.handleVariableLine(tokens)
	}
	_, err := interpreter.handleCommandLine(tokens)
	return err
}

func (interpreter *Interpreter) HasConnection() bool {
	return interpreter.System != nil
}

func (interpreter *Interpreter) requireConnection() {
	if !interpreter.HasConnection() { // if the client is not initialized we don't have a connection
		interpreter.logger.Fatal("Setup a host before using a command that requires a connection.")
	}
}

func (interpreter *Interpreter) handleCommandLine(tokens []string) (string, error) {
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
		return "", fmt.Errorf("invalid command %s", commandName)
	}

	if slices.Contains(interpreter.checkedRequirements, commandName) {
		interpreter.logger.Debugf("passed requirenments for %s (cached)", commandName)
	} else if command.TestRequirements() {
		interpreter.checkedRequirements = append(interpreter.checkedRequirements, commandName)
	} else {
		return "", fmt.Errorf("failed requirenments for %s", commandName)
	}
	output, err := command.Execute(tokens)
	if err != nil {
		return "", fmt.Errorf("execution failed: %w", err)
	}
	return output, nil
}

func (interpreter *Interpreter) handleVariableLine(tokens []string) error {
	variableName := strings.TrimPrefix(tokens[0], "$")
	result, error := interpreter.handleCommandLine(tokens[2:])
	if error != nil {
		return error
	}

	interpreter.Variables[variableName] = *NewVariable(result)
	return nil
}

func (interpreter *Interpreter) GetReferencePage() string {
	return ""
}
