package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

type ScriptAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	controller  *BuilderController
	BaseAction
}

func NewScriptAction(controller *BuilderController) *ScriptAction {

	return &ScriptAction{
		BaseAction:  BaseAction{controller: controller},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *ScriptAction) Execute() {
	if this.ParameterValidationFailed(1, "script needs a file name as argument") {
		return
	}
	this.logger.Info("Builder started")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.TestAndRun(this.controller.Arguments[0])
}

func (this *ScriptAction) GetName() string {
	return "script"
}

func (this *ScriptAction) GetDescription() string {
	return "Run the script in <scriptpath>."
}

func (this *ScriptAction) GetHelp() string {
	return "Run a specific script."
}
