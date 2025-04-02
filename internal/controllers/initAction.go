package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

//go:embed initAction.md
var initMarkdown string

type InitAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	BaseAction
}

func NewInitAction(controller *Controller) *InitAction {
	initAction := &InitAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "init",
			description: "Initialize builder in the current directory.",
			brief:       "Initialize builder in the current directory.",
			help:        initMarkdown,
		},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
	}

	return initAction
}

func (initAction *InitAction) Execute() {
	if initAction.model.IsInitialized() {
		initAction.logger.Info("Already initialized, nothing to do.")
	}
	initAction.model.CreateDirectories()
	initAction.logger.Info("Initializing")
}

func (initAction *InitAction) GetDescription() string {
	return "Create builder directories in " + initAction.environment.GetProjectPath() + "."
}
