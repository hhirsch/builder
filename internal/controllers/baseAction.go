package controllers

import (
	"fmt"
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

func (baseAction *BaseAction) ParameterValidationFailed(requiredAmountOfParameters int, errorMessage string) bool {
	if !baseAction.HasEnoughParameters(requiredAmountOfParameters) {
		panic(errorMessage)
	}
	return false
}

func (baseAction *BaseAction) RequireAmountOfParameters(requiredAmountOfParameters int) error {
	if !baseAction.HasEnoughParameters(requiredAmountOfParameters) {
		return fmt.Errorf("required amount of parameters is %d but got %d", requiredAmountOfParameters, len(baseAction.controller.GetArguments()))
	}
	return nil
}

func (baseAction *BaseAction) HasEnoughParameters(requiredAmountOfParameters int) bool {
	return len(baseAction.controller.GetArguments()) >= requiredAmountOfParameters
}

func (baseAction *BaseAction) GetName() string {
	return baseAction.name
}

func (baseAction *BaseAction) GetDescription() string {
	return baseAction.description
}

func (baseAction *BaseAction) GetBrief() string {
	return baseAction.brief
}

func (baseAction *BaseAction) GetHelp() string {
	return baseAction.help
}
