package controllers

import (
	_ "embed"
	"fmt"
	"github.com/hhirsch/builder/internal/interpreter"
	"github.com/hhirsch/builder/internal/models"
)

//go:embed commandAction.md
var commandMarkdown string

type CommandAction struct {
	environment *models.Environment
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

func (commandAction *CommandAction) getCommandPath(commandName string) string {
	return "builder/commands/" + commandName + ".bld"
}

func (commandAction *CommandAction) Execute() (string, error) {
	err := commandAction.RequireAmountOfParameters(1)
	if err != nil {
		return "", err
	}

	interpreterObject, err := interpreter.NewInterpreter(commandAction.environment)
	if err != nil {
		return "", fmt.Errorf("new interpreter: %w", err)
	}
	interpreter := *interpreterObject
	if err != nil {
		return "", err
	}
	var commandName = commandAction.controller.Arguments[0]
	err = interpreter.Run(commandAction.getCommandPath(commandName))
	if err != nil {
		return "", err
	}
	return "", nil
}
