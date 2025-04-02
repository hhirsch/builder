package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

type ScriptAction struct {
	environment *models.Environment
	model       *models.BuilderModel
	controller  *Controller
	fileName    string
	BaseAction
}

//go:embed scriptAction.md
var scriptMarkdown string

func NewScriptAction(controller *Controller, fileName string) *ScriptAction {

	return &ScriptAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "script",
			description: "Run the script in <scriptpath>.",
			brief:       "Run a specific script.",
			help:        scriptMarkdown,
		},
		environment: controller.GetEnvironment(),
		model:       models.NewBuilderModel(controller.GetEnvironment().GetProjectPath()),
		controller:  controller,
		fileName:    fileName,
	}
}

func (scriptAction *ScriptAction) Execute() (string, error) {
	err := scriptAction.RequireAmountOfParameters(1)
	if err != nil {
		return "", err
	}
	buffer := "Builder started\n"
	var interpreter = *interpreter.NewInterpreter(scriptAction.environment)
	err = interpreter.Run(scriptAction.fileName)
	if err != nil {
		return "", err
	}
	return buffer, nil
}

func (scriptAction *ScriptAction) GetDescription() string {
	return "Run the script in <scriptpath>."
}
