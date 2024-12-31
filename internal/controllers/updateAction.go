package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

type UpdateAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	BaseAction
}

func NewUpdateAction(controller *BuilderController) *UpdateAction {

	return &UpdateAction{
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
	}
}

func (this *UpdateAction) Execute(controller *BuilderController) {
	if this.ParameterValidationFailed(1, "command needs a command name as argument") {
		controller.ShowHelp()
		return
	}
	this.logger.Print("executing user defined command")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.Run("./.builder/commands/" + controller.Arguments[0] + ".bld")
}

func (this *UpdateAction) GetName() string {
	return "update"
}

func (this *UpdateAction) GetDescription() string {
	return "Run migrations."
}

func (this *UpdateAction) GetHelp() string {
	return "Run migrations."
}
