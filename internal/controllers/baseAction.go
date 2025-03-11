package controllers

import (
	"github.com/hhirsch/builder/internal/models"
)

type BaseAction struct {
	environment *models.Environment
	controller  *Controller
	name        string
	description string //describes what the program will do if run
	brief       string //information what the program is for
	help        string //detailed description of parameters with examples
}

func (this *BaseAction) ParameterValidationFailed(requiredAmountOfParameters int, errorMessage string) bool {
	if !this.HasEnoughParameters(requiredAmountOfParameters) {
		this.environment.GetLogger().Fatal(errorMessage)
	}
	return false
}

func (this *BaseAction) HasEnoughParameters(requiredAmountOfParameters int) bool {
	return len(this.controller.GetArguments()) >= requiredAmountOfParameters
}

func (this *BaseAction) GetName() string {
	return this.name
}

func (this *BaseAction) GetDescription() string {
	return this.description
}

func (this *BaseAction) GetBrief() string {
	return this.brief
}

func (this *BaseAction) GetHelp() string {
	return this.help
}
