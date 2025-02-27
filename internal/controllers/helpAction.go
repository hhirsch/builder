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
	controller  *Controller
	BaseAction
}

func NewHelpAction(controller *Controller) *HelpAction {

	return &HelpAction{
		BaseAction:  BaseAction{controller: controller},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *HelpAction) Execute() {
	this.logger.Print(helpers.GetBannerText())

	if len(this.controller.Arguments) > 1 {
		this.logger.Print("Specific help for command " + this.controller.Arguments[0] + ".")
		return
	}
	for _, action := range this.controller.GetActions() {
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
	return "Show this help. Call help with an action name as parameter \n\t\tto get more details on the action."
}
