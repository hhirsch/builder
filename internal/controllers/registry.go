package controllers

import (
	"github.com/hhirsch/builder/internal/models"
)

func NewRegistryController(environment *models.Environment) *Controller {
	var controller = NewController(environment)
	controller.AddAction(NewCreateAction(controller))
	return controller
}
