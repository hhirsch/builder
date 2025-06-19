package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
)

type UpdateAction struct {
	logger *helpers.Logger
	BaseAction
}

func NewUpdateAction(controller *Controller) *UpdateAction {

	return &UpdateAction{}
}

func (updateAction *UpdateAction) Execute(controller *Controller) {
}

func (updateAction *UpdateAction) GetName() string {
	return "update"
}

func (updateAction *UpdateAction) GetHelp() string {
	return "Run migrations."
}
