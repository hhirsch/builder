package models

import (
	"fmt"
	"os"
)

type BuilderModel struct {
	projectPath string
}

func NewBuilderModel(projectPath string) *BuilderModel {

	controller := &BuilderModel{
		projectPath: projectPath,
	}
	return controller
}

func (builderModel *BuilderModel) IsInitialized() bool {
	if _, err := os.Stat(builderModel.projectPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func (builderModel *BuilderModel) CreateDirectories() error {
	err := os.Mkdir(builderModel.projectPath, 0755)
	if err != nil {
		return fmt.Errorf("#[file:line] can't create directory: %w", err)
	}
	return nil
}
