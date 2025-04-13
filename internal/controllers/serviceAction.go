package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/interpreter"
	"github.com/hhirsch/builder/internal/models"
)

//go:embed serviceAction.md
var serviceMarkdown string

type ServiceAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	controller  *Controller
	BaseAction
}

func NewServiceAction(controller *Controller) *ServiceAction {

	return &ServiceAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "service",
			description: "Setup and monitor services.",
			brief:       "Setup and monitor services.",
			help:        serviceMarkdown,
		},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		controller:  controller,
	}

}

func (serviceAction *ServiceAction) install(serviceName string) {
	serviceAction.logger.Info("Builder started")
	interpreter, err := interpreter.NewInterpreter(serviceAction.environment)
	err = interpreter.Run(serviceAction.controller.Arguments[0])
	if err != nil {
		serviceAction.logger.Error(err.Error())
	}
}
func (serviceAction *ServiceAction) uninstall(serviceName string) {}

func (serviceAction *ServiceAction) Execute() (string, error) {
	err := serviceAction.RequireAmountOfParameters(1)
	if err != nil {
		return "", err
	}
	var modifier = serviceAction.controller.Arguments[0]
	var serviceName = serviceAction.controller.Arguments[1]
	if modifier == "install" {
		serviceAction.install(serviceName)
	}
	if modifier == "uninstall" {
		serviceAction.uninstall(serviceName)
	}
	return "", nil
}
