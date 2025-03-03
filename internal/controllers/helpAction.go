package controllers

import (
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"strings"
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
		BaseAction: BaseAction{
			controller: controller,
			name:       "help",
			help:       "Show this help. Call help with an action name as parameter \n\t\tto get more details on the action.",
		},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *HelpAction) rightPadString(s string, length int) string {
	if len(s) >= length {
		return s
	}
	padLength := length - len(s)
	pad := strings.Repeat(" ", padLength)
	return s + pad
}

func (this *HelpAction) Execute() {
	this.logger.Print(helpers.GetBannerText())
	if !this.HasEnoughParameters(2) {
		this.logger.Print("Specific help for command " + this.controller.Arguments[0] + ".")
		return
	}

	for _, action := range this.controller.GetActions() {
		this.logger.Print(fmt.Sprintf("%s\t%+v", this.rightPadString(action.GetName(), 10), action.GetHelp()))
	}
}

func (this *HelpAction) GetDescription() string {
	return "Create builder directories in " + this.environment.GetProjectPath() + "."
}
