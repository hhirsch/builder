package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

//go:embed serviceAction.md
var serviceMarkdown string

type ServiceAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
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
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *ServiceAction) install(serviceName string) {
	this.logger.Info("Builder started")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.Run(this.controller.Arguments[0])
}
func (this *ServiceAction) uninstall(serviceName string) {}

func (this *ServiceAction) Execute() {
	if this.ParameterValidationFailed(2, "service needs a modifier and service name as an argument") {
		return
	}
	var modifier string = this.controller.Arguments[0]
	var serviceName string = this.controller.Arguments[1]
	if modifier == "install" {
		this.install(serviceName)
		return
	}
	if modifier == "uninstall" {
		this.uninstall(serviceName)
		return
	}
}
