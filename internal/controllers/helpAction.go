package controllers

import (
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

type HelpAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
}

func NewHelpAction(environment *models.Environment) *HelpAction {

	initAction := &HelpAction{
		environment: environment,
		logger:      environment.GetLogger(),
		model:       models.NewBuilderModel(environment),
	}

	return initAction
}

func (this *HelpAction) Execute(controller *BuilderController) {
	this.logger.Print(helpers.GetBannerText())

	if len(controller.Arguments) > 1 {
		this.logger.Print("Specific help for command " + controller.Arguments[0] + ".")
		return
	}
	this.logger.Print("help\t\tShow this help. Call help with an action name as parameter \n\t\tto get more details on the action.")
	for _, action := range controller.GetActions() {
		this.logger.Print(fmt.Sprintf("%s\t\t%+v", action.GetName(), action.GetHelp()))
	}
}

func (this *HelpAction) GetName() string {
	return "help"
}

func (this *HelpAction) GetDescription() string {
	return "Create builder directories in " + this.environment.GetProjectPath() + "."
}

func (this *HelpAction) GetHelp() string {
	return "Get Help."
}
