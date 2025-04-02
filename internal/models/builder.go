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

func (builderModel *BuilderModel) IsInitialized() bool {
	if _, err := os.Stat(builderModel.environment.GetProjectPath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func (builderModel *BuilderModel) CreateDirectories() {
	err := os.Mkdir(builderModel.environment.GetProjectPath(), 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}
