package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

type InitAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
}

func NewInitAction(controller *Controller) *InitAction {
	initAction := &InitAction{
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

func (this *InitAction) GetName() string {
	return "init"
}

func (this *InitAction) GetDescription() string {
	return "Create builder directories in " + this.environment.GetProjectPath() + "."
}

func (this *InitAction) GetHelp() string {
	return "Initialize builder in the current directory."
}
