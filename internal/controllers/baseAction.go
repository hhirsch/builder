package controllers

import (
	"github.com/hhirsch/builder/internal/models"
)

type BaseAction struct {
	environment *models.Environment
	controller  *BuilderController
}

func (this *BaseAction) ParameterValidationFailed(requiredAmountOfParameters int, errorMessage string) bool {
	if !this.HasEnoughParameters(requiredAmountOfParameters) {
		this.environment.GetLogger().Fatal(errorMessage)
	}
	return false
}

func (this *BaseAction) HasEnoughParameters(requiredAmountOfParameters int) bool {
	return len(this.controller.Arguments) >= requiredAmountOfParameters
}
