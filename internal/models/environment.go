package models

import (
	"errors"
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/melbahja/goph"
	"os"
	"os/user"
	"runtime"
)

type Environment struct {
	//Client    Client
	Client    *goph.Client
	logger    *helpers.Logger
	arguments []string
	registry  *Registry
}

func NewEnvironment() *Environment {
	environment := &Environment{
		arguments: os.Args,
	}
	logger := helpers.NewLogger(environment.GetLogFilePath())
	registry := NewRegistry(environment.GetBuilderHomePath() + "/builderGlobal.reg")
	environment.SetLogger(logger)
	encryption, err := NewEncryption(environment)
	if err != nil {
		logger.Error("No encryption possible: " + err.Error())
	} else {
		logger.Info("Encryption available.")
		registry.EnableRsa(*encryption)
	}
	err = registry.Load()
	if err != nil {
		logger.Fatalf("Registry loading failed: %v", err.Error())
	}

	environment.SetRegistry(registry)
	return environment
}

func (environment *Environment) getHomePath() string {
	currentUser, err := user.Current()
	if err != nil {
		environment.logger.Fatal(err.Error())
	}

	return currentUser.HomeDir
}

func (environment *Environment) GetLogFilePath() string {
	return "builder.log"
}

func (environment *Environment) GetProjectPath() string {
	return ".builder"
}

func (environment *Environment) GetProjectCommandsPath() string {
	return environment.GetProjectPath() + "/commands/"
}

func (environment *Environment) GetGlobalRegistryPath() string {
	return environment.GetBuilderHomePath() + "/builderGlobal.reg"
}

func (environment *Environment) GetKeyPath() string {
	return environment.getHomePath() + "/.ssh/id_rsa"
}

func (environment *Environment) IsColorEnabled() bool {
	value := os.Getenv("CLICOLOR")
	return value == "1"
}

func (environment *Environment) GetLogger() *helpers.Logger {
	if environment.logger == nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			fmt.Fprintf(os.Stderr, "Error: logger must be set before accessing logger at: %s:%d", file, line)
		} else {
			fmt.Fprintf(os.Stderr, "Error: Logger must be set before accessing logger")
		}
		os.Exit(1)
	}
	return environment.logger
}

func (environment *Environment) GetRegistry() *Registry {
	if environment.registry == nil {
		environment.logger.Fatal("No registry exists at environment point failing disgracefully.")
	}
	return environment.registry
}

func (environment *Environment) SetLogger(logger *helpers.Logger) {
	environment.logger = logger
}

func (environment *Environment) SetRegistry(registry *Registry) {
	environment.registry = registry
}

func (environment *Environment) GetArguments() []string {
	return environment.arguments
}

func (environment *Environment) GetBuilderHomePath() string {
	path := environment.getHomePath() + "/" + environment.GetProjectPath()

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			environment.logger.Fatal(err.Error())
		}
	}

	return path
}

func (environment *Environment) GetBuilderWorkingPath() {
}

func (environment *Environment) GetCommandsPath() {
}
