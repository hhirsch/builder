package controllers

import (
	"github.com/hhirsch/builder/internal/models"
)

func NewBuilderController(environment *models.Environment) *Controller {
	var controller = NewController(environment)
	controller.AddAction(NewInitAction(controller))
	controller.AddAction(NewScriptAction(controller))
	controller.AddAction(NewCommandAction(controller))
	controller.AddAction(NewCreateAction(controller))
	controller.AddAction(NewServerAction(controller))
	controller.AddAction(NewServiceAction(controller))
	controller.AddAction(NewRegistryAction(controller))
	return controller
}
