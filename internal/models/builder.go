package models

import (
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"os"
)

type BuilderModel struct {
	environment *Environment
	logger      *helpers.Logger
	arguments   []string
}

func NewBuilderModel(environment *Environment) *BuilderModel {

	controller := &BuilderModel{
		environment: environment,
		logger:      environment.GetLogger(),
		arguments:   environment.GetArguments(),
	}
	return controller
}

func (this *BuilderModel) IsInitialized() bool {
	if _, err := os.Stat(this.environment.GetProjectPath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func (this *BuilderModel) CreateDirectories() {
	err := os.Mkdir(this.environment.GetProjectPath(), 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}
