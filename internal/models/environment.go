package models

import (
	"errors"
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"os"
	"os/user"
	"runtime"
)

type Environment struct {
	Client     Client
	configPath string
	logger     *helpers.Logger
	arguments  []string
	registry   *Registry
}

func NewEnvironment() *Environment {
	environment := &Environment{
		arguments: os.Args,
	}
	registry := NewRegistry(environment.GetBuilderHomePath() + "/builderGlobal.reg")
	registry.Load()

	logger := helpers.NewLogger(environment.GetLogFilePath())
	environment.SetLogger(logger)
	encryption, err := NewEncryption(environment)
	if err != nil {
		logger.Error("No encryption possible: " + err.Error())
	} else {
		logger.Info("Encryption available.")
		registry.EnableRsa(*encryption)
		if registry.EncryptionTest() == nil {
			logger.Info("Encryption test passed.")
		} else {
			logger.Fatal("Encryption test failed.")
		}
	}
	environment.SetRegistry(registry)
	return environment
}

func (this *Environment) getHomePath() string {
	currentUser, err := user.Current()
	if err != nil {
		this.logger.Fatal(err.Error())
	}

	return currentUser.HomeDir
}

func (this *Environment) GetLogFilePath() string {
	return "builder.log"
}

func (this *Environment) GetProjectPath() string {
	return ".builder"
}

func (this *Environment) GetProjectCommandsPath() string {
	return this.GetProjectPath() + "/commands/"
}

func (this *Environment) GetGlobalRegistryPath() string {
	return this.GetBuilderHomePath() + "/builderGlobal.reg"
}

func (this *Environment) GetKeyPath() string {
	return this.getHomePath() + "/.ssh/id_rsa"
}

func (this *Environment) GetLogger() *helpers.Logger {
	if this.logger == nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			fmt.Fprintf(os.Stderr, "Error: logger must be set before accessing logger at: %s:%d", file, line)
		} else {
			fmt.Fprintf(os.Stderr, "Error: Logger must be set before accessing logger")
		}
		os.Exit(1)
	}
	return this.logger
}

func (this *Environment) GetRegistry() *Registry {
	return this.registry
}

func (this *Environment) SetLogger(logger *helpers.Logger) {
	this.logger = logger
}

func (this *Environment) SetRegistry(registry *Registry) {
	this.registry = registry
}

func (this *Environment) GetArguments() []string {
	return this.arguments
}

func (this *Environment) GetBuilderHomePath() string {
	path := this.getHomePath() + "/" + this.GetProjectPath()

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			this.logger.Fatal(err.Error())
		}
	}

	return path
}

func (this *Environment) GetBuilderWorkingPath() {
}

func (this *Environment) GetCommandsPath() {
}
