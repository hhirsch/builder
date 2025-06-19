package controllers

import (
	_ "embed"
	"github.com/hhirsch/builder/internal/models"
)

type ListCommandsAction struct {
	environment *models.Environment
	BaseAction
}

func NewListAction(controller *Controller, commandPath string) *ListCommandsAction {
	return &ListCommandsAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "command",
			brief:       "Execute command.",
			description: "Execute command.",
			help:        commandMarkdown,
		},
	}
}

func (listCommandsAction *ListCommandsAction) Execute() (string, error) {
	return "", nil
}
