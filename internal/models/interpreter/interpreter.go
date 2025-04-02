package interpreter

import (
	"bufio"
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	com "github.com/hhirsch/builder/internal/models/interpreter/commands"
	"os"
	"slices"
	"strings"
)

type Interpreter struct {
	logger              *helpers.Logger
	environment         *models.Environment
	registry            *models.Registry
	commands            map[string]com.Command
	onlineCommands      map[string]com.Command
	variables           map[string]string
	checkedRequirements []string
}

func NewInterpreter(environment *models.Environment) *Interpreter {
	logger := environment.GetLogger()
	registry := models.NewRegistry(environment.GetGlobalRegistryPath())
	err := registry.Load()
	if err != nil {
		logger.Fatalf("unable to load registry: %s", err.Error())
	}
	variables := map[string]string{}
	commands := map[string]com.Command{}
	onlineCommands := map[string]com.Command{}
	interpreter := &Interpreter{
		logger:         logger,
		registry:       registry,
		environment:    environment,
		commands:       commands,
		variables:      variables,
		onlineCommands: onlineCommands,
	}
	interpreter.AddCommand(com.NewSetupHostCommand(environment))
	interpreter.AddCommand(com.NewConnectCommand(environment))
	interpreter.AddCommand(com.NewStepCommand(environment))
	interpreter.AddCommand(com.NewPrintCommand(environment))
	interpreter.AddCommand(com.NewListFilesCommand(environment))
	interpreter.AddCommand(com.NewSystemInfoCommand(environment))
	interpreter.AddCommand(com.NewEnsurePackageCommand(environment))
	interpreter.AddCommand(com.NewEnsureExecutableCommand(environment))
	interpreter.AddCommand(com.NewEnsureServiceCommand(environment))
	interpreter.AddCommand(com.NewListPackagesCommand(environment))
	interpreter.AddCommand(com.NewDumpPackagesCommand(environment))
	interpreter.AddCommand(com.NewExecuteAndPrintCommand(environment))
	interpreter.AddCommand(com.NewSetTargetUserCommand(environment))
	interpreter.AddCommand(com.NewPushFileCommand(environment))
	interpreter.AddCommand(com.NewListDatabasesCommand(environment))
	return interpreter
}

func (interpreter *Interpreter) AddCommand(command com.Command) {
	if command.RequiresConnection() {
		interpreter.onlineCommands[command.GetName()] = command
	}
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

func (interpreter *Interpreter) requireConnection() {
	if interpreter.environment.Client == (models.Client{}) { // if the client is not initialized we don't have a connection
		interpreter.logger.Fatal("Setup a host before using a command that requires a connection.")
	}
}

func (interpreter *Interpreter) handleCommandLine(tokens []string) string {
	commandName := tokens[0]
	if commandName == "connect" || commandName == "setupHost" {
		interpreter.checkedRequirements = []string{}
	}
	var command com.Command
	if offlineCommand, isOfflineCommand := interpreter.commands[commandName]; isOfflineCommand {
		command = offlineCommand
	}

	if onlineCommand, isOnlineCommand := interpreter.onlineCommands[commandName]; isOnlineCommand {
		interpreter.requireConnection()
		command = onlineCommand
	}

	if command == nil {
		interpreter.logger.Fatalf("Invalid command %s.", commandName)
	}

	interpreter.logger.Debugf("Testing requirements for %s.", commandName)
	if slices.Contains(interpreter.checkedRequirements, commandName) {
		interpreter.logger.Debugf("Passed requirenments for %s. (cached)", commandName)
	} else if command.TestRequirements() {
		interpreter.logger.Debugf("Passed requirements for %s.", commandName)
		interpreter.checkedRequirements = append(interpreter.checkedRequirements, commandName)
	} else {
		interpreter.logger.Fatalf("Failed requirenments for %s.", commandName)
	}

	return command.Execute(tokens)
}

func (interpreter *Interpreter) handleVariableLine(tokens []string) {
	variableName := strings.TrimPrefix(tokens[0], "$")
	interpreter.variables[variableName] = interpreter.handleCommandLine(tokens[2:])
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
