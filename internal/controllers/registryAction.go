package controllers

import (
	"github.com/hhirsch/builder/internal/environment"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/interpreter"
	"github.com/hhirsch/builder/internal/models"
)

type RegistryAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	controller  *Controller
	BaseAction
}

func NewRegistryAction(controller *Controller) *RegistryAction {

	return &RegistryAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "register",
			description: "Setup and monitor services.",
			help:        "View and edit registry.",
		},
		model:      models.NewBuilderModel(environment.GetProjectPath()),
		controller: controller,
	}

}

func (registryAction *RegistryAction) Execute() (string, error) {
	err := registryAction.RequireAmountOfParameters(1)
	if err != nil {
		return "", err
	}
	registryAction.logger.Info("Builder started")
	interpreter, err := interpreter.NewInterpreter(registryAction.environment)
	if err != nil {
		return "", err
	}
	err = interpreter.Run(registryAction.controller.Arguments[0])
	if err != nil {
		return "", err
	}
	return "", nil
}
