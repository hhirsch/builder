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

func (commandAction *CommandAction) Execute() {
	if commandAction.ParameterValidationFailed(1, "command needs a command name as argument.") {
		commandAction.controller.ShowHelp()
		return
	}
	commandAction.logger.Info("Executing user defined command.")
	var interpreter = *interpreter.NewInterpreter(commandAction.environment)
	var commandName = commandAction.controller.Arguments[0]
	err := interpreter.Run(commandAction.environment.GetProjectCommandsPath() + commandName + ".bld")
	if err != nil {
		commandAction.logger.Error(err.Error())
	}
}
