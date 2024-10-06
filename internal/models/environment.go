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
}

func NewEnvironment() *Environment {

	environment := &Environment{
		arguments: os.Args,
	}
	logger := helpers.NewLogger(environment.GetLogFilePath())
	environment.SetLogger(logger)
	return environment
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

func (this *Environment) SetLogger(logger *helpers.Logger) {
	this.logger = logger
}

func (this *Environment) GetArguments() []string {
	return this.arguments
}

func (this *Environment) GetBuilderHomePath() string {
	currentUser, err := user.Current()
	path := currentUser.HomeDir + "/" + this.GetProjectPath()

	if err != nil {
		this.logger.Fatal(err.Error())
	}

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
