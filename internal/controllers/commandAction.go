package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

type CommandAction struct {
	environment *models.Environment
	logger      *helpers.Logger
}

func NewCommandAction(environment *models.Environment) *CommandAction {

	initAction := &CommandAction{
		environment: environment,
		logger:      environment.GetLogger(),
	}

	return initAction
}

func (this *CommandAction) Execute(controller *BuilderController) {
	if controller.ParameterValidationFailed(1, "command needs a command name as argument") {
		controller.HelpAction()
		return
	}
	this.logger.Info("Executing user defined command.")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.TestAndRun(this.environment.GetProjectCommandsPath() + controller.Arguments[0] + ".bld")
}

func (this *CommandAction) GetName() string {
	return "command"
}

func (this *CommandAction) GetDescription() string {
	return "execute command"
}

func (this *CommandAction) GetHelp() string {
	return "execute command"
}
