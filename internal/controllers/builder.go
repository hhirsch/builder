package controllers

import (
	"github.com/hhirsch/builder/internal/models"
)

func NewBuilderController(environment *models.Environment) *Controller {
	var controller = NewController(environment)
	var parameter = ""
	if len(controller.Arguments) > 0 {
		parameter = controller.Arguments[0]
	}
	var commandPath = environment.GetProjectCommandsPath() + parameter + ".bld"
	controller.AddAction(NewInitAction(controller))
	controller.AddAction(NewScriptAction(controller, parameter))
	controller.AddAction(NewCommandAction(controller, commandPath))
	controller.AddAction(NewServerAction(controller))
	controller.AddAction(NewServiceAction(controller))
	controller.AddAction(NewRegistryAction(controller))
	controller.AddAction(NewReferenceAction(controller))
	return controller
}
