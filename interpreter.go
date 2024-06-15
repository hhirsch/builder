package main

import (
	"bufio"
	"errors"
	format "fmt"
	"github.com/charmbracelet/huh"
	"log"
	"os"
	"os/user"
	"strings"
)

type Interpreter struct {
	client   Client
	logger   *Logger
	registry *Registry
	step     string
}

func NewInterpreter() *Interpreter {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	path := currentUser.HomeDir + "/.builder"
	logger := NewLogger()
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			logger.Warn(err.Error())
		}
	}
	registry := NewRegistry(path + "/builderGlobal.reg")
	registry.Load()
	return &Interpreter{
		logger:   logger,
		registry: registry,
	}
}

func (this *Interpreter) run(fileName string) {
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
	if this.client == (Client{}) {
		this.logger.Fatal("Setup a host before using a command that requires a connection.")
	}
}

func (this *Interpreter) Step(step string) {
	this.step = step
	this.logger.Info(step)
}

func (this *Interpreter) setupHost(tokens []string) {
	if len(tokens) != 2 {
		this.logger.Fatal("setupHost needs 1 parameter")
	}
	var err error
	var userName string
	var userPath = "host." + tokens[1] + ".user"
	userName, err = this.registry.GetValue(userPath)
	if err != nil {
		this.logger.Info("No user found in registry asking for user input.")
		userName = "root"
		huh.NewInput().
			Title("Enter user name").
			Value(&userName).
			Run()

		this.logger.Info("Registering " + userPath + " as " + userName)
		this.registry.Register(userPath, userName)
	}
	if len(userName) == 0 {
		this.logger.Fatal("User name must not be empty!")
	}
	var hostName string
	var hostPath = "host." + tokens[1] + ".host"
	hostName, err = this.registry.GetValue(hostPath)
	if err != nil {
		this.logger.Info("No host found in registry asking for user input.")
		huh.NewInput().
			Title("Enter host name or IP").
			Value(&hostName).
			Run()
		this.logger.Info("Registering " + hostPath + " as " + hostName)
		this.registry.Register(hostPath, hostName)
	}

	if len(hostName) == 0 {
		this.logger.Fatal("Host name must not be empty!")
	}

	this.registry.Save()
	this.client = *NewClient(userName, hostName)
}

func (this *Interpreter) handleLine(input string) {
	tokens := strings.Fields(input)
	if strings.HasPrefix(tokens[0], "//") {
		return
	}

	switch tokens[0] {
	case "setupHost":
		this.setupHost(tokens)
		return
	case "step":
		tokens = tokens[1:]
		parameters := strings.Join(tokens, " ")
		this.Step(parameters)
		return
	case "ensureService":
		if len(tokens) < 4 {
			this.logger.Fatal("ensureService needs 2 parameters and a description string")
		}
		reducedTokens := tokens[3:]
		description := strings.Join(reducedTokens, " ")
		this.logger.Info("Creating service name: " + tokens[1] + "  binary: " + tokens[2] + "  description: " + description)
		this.client.EnsureService(tokens[1], tokens[2], description)
		return
	}
	this.requireConnection()
	switch tokens[0] {
	case "listPackages":
		format.Println("list")
		this.client.ListPackages()
	case "dumpPackages":
		format.Println("list")
		this.client.DumpPackages()
	case "executeAndPrint":
		tokens = tokens[1:]
		parameters := strings.Join(tokens, " ")
		this.client.ExecuteAndPrint(parameters)
	case "ensureCapabilityConnection":
		if len(tokens) != 2 {
			this.logger.Fatal("ensureCapabilityConnection needs 1 parameters")
		}
		this.client.EnsureCapabilityConnection(tokens[1])
	case "ensurePackage":
		tokens = tokens[1:]
		parameters := strings.Join(tokens, " ")
		this.client.EnsurePackage(parameters)
	case "ensureExecutable":
		this.logger.Info("Ensuring target is executable.")
		if len(tokens) != 2 {
			this.logger.Fatal("ensureExecutable needs 1 parameter")
		}
		this.client.EnsureExecutable(tokens[1])
	case "print":
		tokens = tokens[1:]
		parameters := strings.Join(tokens, " ")
		format.Println(parameters)
	case "setTargetUser":
		this.logger.Info("Setting target user.")
		if len(tokens) != 2 {
			this.logger.Fatal("setTargetUser needs 1 parameter")
		}
		this.client.SetTargetUser(tokens[1])
	case "pushFile":
		if len(tokens) != 3 {
			this.logger.Fatal("pushFile needs 2 parameters")
		}
		this.client.PushFile(tokens[1], tokens[2])
	default:
		format.Println("Invalid command " + tokens[0])
	}
}
