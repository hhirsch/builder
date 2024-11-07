package interpreter

import (
	"bufio"
	format "fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	com "github.com/hhirsch/builder/internal/models/interpreter/commands"
	"log"
	"os"
	"strings"
)

type Interpreter struct {
	logger         *helpers.Logger
	environment    *models.Environment
	registry       *models.Registry
	step           string
	commands       map[string]com.Command
	onlineCommands map[string]com.Command
}

func NewInterpreter(environment *models.Environment) *Interpreter {
	logger := environment.GetLogger()
	registry := models.NewRegistry(environment.GetGlobalRegistryPath())
	registry.Load()
	commands := map[string]com.Command{
		"step":      com.NewStepCommand(environment),
		"print":     com.NewPrintCommand(environment),
		"setupHost": com.NewSetupHostCommand(environment),
	}
	onlineCommands := map[string]com.Command{
		"ensurePackage":    com.NewEnsurePackageCommand(environment),
		"ensureExecutable": com.NewEnsureExecutableCommand(environment),
		"ensureService":    com.NewEnsureServiceCommand(environment),
		"listPackages":     com.NewListPackagesCommand(environment),
		"dumpPackages":     com.NewDumpPackagesCommand(environment),
		"executeAndPrint":  com.NewExecuteAndPrintCommand(environment),
		"setTargetUser":    com.NewSetTargetUserCommand(environment),
		"pushFile":         com.NewPushFileCommand(environment),
	}
	return &Interpreter{
		logger:         logger,
		registry:       registry,
		environment:    environment,
		commands:       commands,
		onlineCommands: onlineCommands,
	}
}

func (this *Interpreter) Run(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		this.handleLine(line)
	}

	if err := scanner.Err(); err != nil {
		format.Println("Error scanning file:", err)
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
		command.Execute(tokens)
	} else {
		format.Println("Invalid command " + tokens[0])
	}

}
