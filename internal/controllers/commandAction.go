package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/interpreter"
)

//go:embed commandAction.md
var commandMarkdown string

type CommandAction struct {
	BaseAction
}

func NewCommandAction(controller *Controller, commandPath string) *CommandAction {

	return &CommandAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "command",
			brief:       "Execute command.",
			description: "Execute command.",
			help:        commandMarkdown,
		},
	}
}

func (commandAction *CommandAction) Execute() (string, error) {
	err := commandAction.RequireAmountOfParameters(1)
	if err != nil {
		return "", err
	}
	var interpreter = *interpreter.NewInterpreter(commandAction.environment)
	var commandName = commandAction.controller.Arguments[0]
	err = interpreter.Run(commandAction.environment.GetProjectCommandsPath() + commandName + ".bld")
	if err != nil {
		return "", err
	}
	return "", nil
}
