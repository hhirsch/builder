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
}

func NewScriptAction(environment *models.Environment) *ScriptAction {

	scriptAction := &ScriptAction{
		environment: environment,
		logger:      environment.GetLogger(),
		model:       models.NewBuilderModel(environment),
	}

	return scriptAction
}

func (this *ScriptAction) Execute(controller *BuilderController) {
	if controller.ParameterValidationFailed(1, "script needs a file name as argument") {
		return
	}
	this.logger.Info("Builder started")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.Run(controller.Arguments[0])
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
