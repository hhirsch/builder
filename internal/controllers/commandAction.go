package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

type CommandAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	BaseAction
}

func NewCommandAction(controller *BuilderController) *CommandAction {

	return &CommandAction{
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
	}
}

func (this *CommandAction) Execute() {
	if this.ParameterValidationFailed(1, "command needs a command name as argument") {
		this.controller.ShowHelp()
		return
	}
	this.logger.Info("Executing user defined command.")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.TestAndRun(this.environment.GetProjectCommandsPath() + this.controller.Arguments[0] + ".bld")
}

func (this *CommandAction) GetName() string {
	return "command"
}

func (this *CommandAction) GetDescription() string {
	return "Execute command"
}

func (this *CommandAction) GetHelp() string {
	return "Execute command."
}
