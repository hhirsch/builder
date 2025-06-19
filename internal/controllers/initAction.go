package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/environment"
	"github.com/hhirsch/builder/internal/models"
)

//go:embed initAction.md
var initMarkdown string

type InitAction struct {
	model *models.BuilderModel
	BaseAction
}

func NewInitAction(controller *Controller) *InitAction {
	initAction := &InitAction{
		BaseAction: BaseAction{
			controller: controller,
			name:       "init",
			brief:      "Initialize builder in the current directory.",
			help:       initMarkdown,
		},
		model: models.NewBuilderModel(environment.GetProjectPath()),
	}

	return initAction
}

func (initAction *InitAction) Execute() (string, error) {
	if initAction.model.IsInitialized() {
		return "Already initialized, nothing to do.\n", nil
	}
	err := initAction.model.CreateDirectories()
	if err != nil {
		return "", err
	}
	return "Initializing\n", nil
}
