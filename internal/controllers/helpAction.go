package controllers

import (
	_ "embed"
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/valyala/fasttemplate"
	"strings"
)

type HelpAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	controller  *Controller
	BaseAction
}

//go:embed helpHeader.md
var helpHeader string

//go:embed helpAction.md
var helpMarkdown string

func NewHelpAction(controller *Controller) *HelpAction {
	baseAction := BaseAction{
		controller:  controller,
		name:        "help",
		description: "Show the help for [parameter].",
		brief:       "Show this help. Call help with an action name as parameter \n\t\tto get more details on the action.",
		help:        helpMarkdown,
	}

	return &HelpAction{
		BaseAction:  baseAction,
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *HelpAction) rightPadString(string string, length int) string {
	if len(string) >= length {
		return string
	}
	padLength := length - len(string)
	pad := strings.Repeat(" ", padLength)
	return string + pad
}

func (this *HelpAction) Execute() {
	var markdownRenderer = models.NewMarkdownRenderer()

	if this.HasEnoughParameters(1) {
		var actionName = this.controller.Arguments[0]
		this.logger.Print(helpers.GetBannerText())
		if this.environment.IsColorEnabled() {
			markdownRenderer.EnableColor()
		}
		action := this.controller.actionsMap[actionName]
		template := fasttemplate.New(helpHeader+action.GetHelp(), "{{", "}}")
		markdownContent := template.ExecuteString(map[string]interface{}{
			"actionName": actionName,
			"binaryName": this.environment.GetArguments()[0],
			"brief":      action.GetBrief(),
		})
		markdownRenderer.Render(markdownContent)
		return
	} else {
		this.logger.Print(helpers.GetBannerText())
	}

	fmt.Printf("  %s <command> [<arguments>]\n\n", this.environment.GetArguments()[0])
	for _, action := range this.controller.GetActions() {
		this.logger.Print(fmt.Sprintf("  %s\t%+v", this.rightPadString(action.GetName(), 10), action.GetBrief()))
	}
	fmt.Print("\n  Set the environment variable CLICOLOR to 1 to enable colors.\n")
}

func (this *HelpAction) GetDescription() string {
	return "Create builder directories in " + this.environment.GetProjectPath() + "."
}
