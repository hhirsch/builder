package interpreter

import (
	"bufio"
	format "fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	com "github.com/hhirsch/builder/internal/models/interpreter/commands"
	"os"
	"strings"
)

type Interpreter struct {
	logger            *helpers.Logger
	environment       *models.Environment
	registry          *models.Registry
	step              string
	commands          map[string]com.Command
	onlineCommands    map[string]com.Command
	testRequirenments bool
}

func NewInterpreter(environment *models.Environment) *Interpreter {
	logger := environment.GetLogger()
	registry := models.NewRegistry(environment.GetGlobalRegistryPath())
	registry.Load()
	commands := map[string]com.Command{
		"step":      com.NewStepCommand(environment),
		"print":     com.NewPrintCommand(environment),
		"setupHost": com.NewSetupHostCommand(environment),
		"connect":   com.NewConnectCommand(environment),
	}
	onlineCommands := map[string]com.Command{
		"systemInfo":       com.NewSystemInfoCommand(environment),
		"ensurePackage":    com.NewEnsurePackageCommand(environment),
		"ensureExecutable": com.NewEnsureExecutableCommand(environment),
		"ensureService":    com.NewEnsureServiceCommand(environment),
		"listPackages":     com.NewListPackagesCommand(environment),
		"dumpPackages":     com.NewDumpPackagesCommand(environment),
		"executeAndPrint":  com.NewExecuteAndPrintCommand(environment),
		"setTargetUser":    com.NewSetTargetUserCommand(environment),
		"pushFile":         com.NewPushFileCommand(environment),
		"saveDatabase":     com.NewPushFileCommand(environment),      //host, database, localFileName
		"listDatabases":    com.NewListDatabasesCommand(environment), //host, database, localFileName
	}
	return &Interpreter{
		logger:            logger,
		registry:          registry,
		environment:       environment,
		commands:          commands,
		onlineCommands:    onlineCommands,
		testRequirenments: false,
	}
}

func (this *Interpreter) TestAndRun(fileName string) {
	this.logger.Info("Testing requirements of file " + fileName)
	this.Test(fileName)
	this.logger.Info("All requirements passed for file " + fileName)
	this.logger.Info("Executing file " + fileName)
	this.Run(fileName)
}

func (this *Interpreter) Test(fileName string) {
	this.testRequirenments = true
	this.Run(fileName)
	this.testRequirenments = false
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

func (this *Interpreter) handleLine(input string) {
	tokens := strings.Fields(input)
	if strings.HasPrefix(tokens[0], "//") {
		return
	}

	if command, exists := this.commands[tokens[0]]; exists {
		command.Execute(tokens)
		return
	}
	this.requireConnection()

	if command, exists := this.onlineCommands[tokens[0]]; exists {
		if this.testRequirenments {
			this.logger.Debug("Testing requirements for " + tokens[0])
			if command.TestRequirements() {
				this.logger.Debug("Passed requirenments for " + tokens[0])
			} else {
				this.logger.Error("Failed requirenments for " + tokens[0])
			}
		} else {
			command.Execute(tokens)
		}
	} else {
		this.logger.Error("Invalid command " + tokens[0])
	}
}

func (this *Interpreter) GetReferencePage() string {
	return ""
}
