package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

type ScriptAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	controller  *Controller
	BaseAction
}

//go:embed scriptAction.md
var scriptMarkdown string

func NewScriptAction(controller *Controller) *ScriptAction {

	return &ScriptAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "script",
			description: "Run the script in <scriptpath>.",
			brief:       "Run a specific script.",
			help:        scriptMarkdown,
		},
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
	err := interpreter.Run(this.controller.Arguments[0])
	if err != nil {
		this.logger.Errorf("Interpreter failed: %v", err)
	}
}

func (this *ScriptAction) GetDescription() string {
	return "Run the script in <scriptpath>."
}
