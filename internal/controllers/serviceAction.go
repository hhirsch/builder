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

func (serviceAction *ServiceAction) getServiceDirectory(serviceName string) string {
	return "builder/services/" + serviceName
}

func (serviceAction *ServiceAction) runServiceScript(serviceName string, fileName string) {
	serviceAction.logger.Info("Builder started")
	interpreter, err := interpreter.NewInterpreter(serviceAction.environment)
	if err != nil {
		serviceAction.logger.Error(err.Error())
	}
	err = interpreter.Run(serviceAction.getServiceDirectory(serviceName) + fileName)
	if err != nil {
		serviceAction.logger.Error(err.Error())
	}
}

func (serviceAction *ServiceAction) install(serviceName string) {
	serviceAction.runServiceScript(serviceName, "install.bld")
}

func (serviceAction *ServiceAction) uninstall(serviceName string) {
	serviceAction.runServiceScript(serviceName, "uninstall.bld")
}

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
