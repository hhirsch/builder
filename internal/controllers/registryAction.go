package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

type RegistryAction struct {
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
func NewRegistryAction(controller *Controller) *RegistryAction {

	return &RegistryAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "register",
			description: "Setup and monitor services.",
			help:        "View and edit registry.",
		},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *RegistryAction) Execute() {
	if this.ParameterValidationFailed(1, "registry needs an argument") {
		return
	}
	this.logger.Info("Builder started")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.Run(this.controller.Arguments[0])
}
