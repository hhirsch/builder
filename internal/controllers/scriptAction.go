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

func (scriptAction *ScriptAction) Execute() (string, error) {
	err := scriptAction.RequireAmountOfParameters(1)
	if err != nil {
		return "", err
	}
	buffer := "Builder started\n"
	var interpreter = *interpreter.NewInterpreter(scriptAction.environment)
	err = interpreter.Run(scriptAction.controller.Arguments[0])
	if err != nil {
		return "", err
	}
	return buffer, nil
}

func (scriptAction *ScriptAction) GetDescription() string {
	return "Run the script in <scriptpath>."
}
