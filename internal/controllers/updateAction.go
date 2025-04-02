package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

type UpdateAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	BaseAction
}

func NewUpdateAction(controller *Controller) *UpdateAction {

	return &UpdateAction{
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
	}
}

func (updateAction *UpdateAction) Execute(controller *Controller) {
}

func (updateAction *UpdateAction) GetName() string {
	return "update"
}

func (updateAction *UpdateAction) GetDescription() string {
	return "Run migrations."
}

func (updateAction *UpdateAction) GetHelp() string {
	return "Run migrations."
}
