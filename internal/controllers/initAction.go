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
			brief:       "Initialize builder in the current directory.\n\t\tAdditional information.",
			help:        initMarkdown,
		},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
	}

	return initAction
}

func (this *InitAction) Execute() {
	if this.model.IsInitialized() {
		this.logger.Info("Already initialized, nothing to do.")
	}
	this.model.CreateDirectories()
	this.logger.Info("Initializing")
}

func (this *InitAction) GetDescription() string {
	return "Create builder directories in " + this.environment.GetProjectPath() + "."
}
