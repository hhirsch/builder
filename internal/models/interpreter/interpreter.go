package interpreter

import (
	"bufio"
	format "fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter/commands"
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
		logger.Fatalf("Unable to load registry: %s", err.Error())
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

func (this *Interpreter) AddCommand(command commands.Command) {
	if command.RequiresConnection() {
		this.onlineCommands[command.GetName()] = command
	}
	this.commands[command.GetName()] = command
}

func (this *Interpreter) Run(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		this.logger.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		this.handleLine(line)
	}

	if err := scanner.Err(); err != nil {
		this.logger.Fatal(format.Printf("Error scanning file: %s", err.Error()))
	}
	file.Close()
}

func (this *Interpreter) requireConnection() {
	if this.environment.Client == (models.Client{}) { // if the client is not initialized we don't have a connection
		this.logger.Fatal("Setup a host before using a command that requires a connection.")
	}
}

func (this *Interpreter) handleCommandLine(tokens []string) string {
	var commandName string = tokens[0]
	if commandName == "connect" || commandName == "setupHost" {
		//this.checkedRequirements = this.checkedRequirements[:0]
		this.checkedRequirements = []string{}
	}
	var command com.Command
	if offlineCommand, isOfflineCommand := this.commands[commandName]; isOfflineCommand {
		command = offlineCommand
	}

	if onlineCommand, isOnlineCommand := this.onlineCommands[commandName]; isOnlineCommand {
		this.requireConnection()
		command = onlineCommand
	}

	if command == nil {
		this.logger.Fatalf("Invalid command %s.", commandName)
	}

	this.logger.Debugf("Testing requirements for %s.", commandName)
	if slices.Contains(this.checkedRequirements, commandName) {
		this.logger.Debugf("Passed requirenments for %s. (cached)", commandName)
	} else if command.TestRequirements() {
		this.logger.Debugf("Passed requirements for %s.", commandName)
		this.checkedRequirements = append(this.checkedRequirements, commandName)
	} else {
		this.logger.Fatalf("Failed requirenments for %s.", commandName)
	}

	return command.Execute(tokens)
}

func (this *Interpreter) handleVariableLine(tokens []string) {
	variableName := strings.TrimPrefix(tokens[0], "$")
	this.variables[variableName] = this.handleCommandLine(tokens[2:])
}

func (this *Interpreter) handleLine(input string) {
	tokens := strings.Fields(input)
	if strings.HasPrefix(tokens[0], "//") {
		return
	}

	if strings.HasPrefix(tokens[0], "$") && tokens[1] == "=" {
		this.handleVariableLine(tokens)
		return
	}

	this.handleCommandLine(tokens)
}

func (this *Interpreter) GetReferencePage() string {
	return ""
}
