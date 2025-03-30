package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

//go:embed commandAction.md
var commandMarkdown string

type CommandAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	BaseAction
}

func NewCommandAction(controller *Controller) *CommandAction {

	return &CommandAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "command",
			brief:       "Execute command.",
			description: "Execute command.",
			help:        commandMarkdown,
		},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
	}
}

func (this *CommandAction) Execute() {
	if this.ParameterValidationFailed(1, "command needs a command name as argument.") {
		this.controller.ShowHelp()
		return
	}
	this.logger.Info("Executing user defined command.")
	var interpreter = *interpreter.NewInterpreter(this.environment)
	var commandName = this.controller.Arguments[0]
	err := interpreter.Run(this.environment.GetProjectCommandsPath() + commandName + ".bld")
	if err != nil {
		this.logger.Error(err.Error())
	}
}
