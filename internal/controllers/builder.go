package controllers

import (
	"github.com/hhirsch/builder/internal/models"
)

func NewBuilderController(environment *models.Environment) *Controller {
	var controller = NewController(environment)
	var commandName = controller.Arguments[0]
	var commandPath = environment.GetProjectCommandsPath() + commandName + ".bld"

	controller.AddAction(NewInitAction(controller))
	controller.AddAction(NewScriptAction(controller))
	controller.AddAction(NewCommandAction(controller, commandPath))
	controller.AddAction(NewCreateAction(controller))
	controller.AddAction(NewServerAction(controller))
	controller.AddAction(NewServiceAction(controller))
	controller.AddAction(NewRegistryAction(controller))
	controller.AddAction(NewReferenceAction(controller))
	return controller
}
