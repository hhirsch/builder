package controllers

import (
	_ "embed"
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/valyala/fasttemplate"
	"strings"
)

type ReferenceAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	controller  *Controller
	BaseAction
}

//go:embed referenceAction.md
var referenceMarkdown string

func NewReferenceAction(controller *Controller) *ReferenceAction {
	baseAction := BaseAction{
		controller:  controller,
		name:        "reference",
		description: "Show the blang reference.",
		brief:       "Show the blang reference.",
		help:        referenceMarkdown,
	}

	return &ReferenceAction{
		BaseAction:  baseAction,
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *ReferenceAction) rightPadString(string string, length int) string {
	if len(string) >= length {
		return string
	}
	padLength := length - len(string)
	pad := strings.Repeat(" ", padLength)
	return string + pad
}

func (this *ReferenceAction) Execute() {
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

func (this *ReferenceAction) GetDescription() string {
	return "Create builder directories in " + this.environment.GetProjectPath() + "."
}
