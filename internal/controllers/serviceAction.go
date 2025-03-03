package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

type ServiceAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	controller  *Controller
	BaseAction
}

/*
 * Eg server list, server add, server alias, server require {serviceName} {serverName}
 * service list, service health, service health {serviceName}, service install {serviceName}
 */
func NewServiceAction(controller *Controller) *ServiceAction {

	return &ServiceAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "service",
			description: "Setup and monitor services.",
			help:        "Setup and monitor services.",
		},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *ServiceAction) Execute() {
	if this.ParameterValidationFailed(1, "script needs a file name as argument") {
		return
	}
	this.logger.Info("Builder started")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.Run(this.controller.Arguments[0])
}
