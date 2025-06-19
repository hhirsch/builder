package controllers

import (
	_ "embed"
	"fmt"
	"github.com/hhirsch/builder/internal/environment"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/valyala/fasttemplate"
	"strings"
)

type ReferenceAction struct {
	logger     *helpers.Logger
	model      *models.BuilderModel
	controller *Controller
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
		BaseAction: baseAction,
		model:      models.NewBuilderModel(environment.GetProjectPath()),
		controller: controller,
	}

}

func (referenceAction *ReferenceAction) rightPadString(string string, length int) string {
	if len(string) >= length {
		return string
	}
	padLength := length - len(string)
	pad := strings.Repeat(" ", padLength)
	return string + pad
}

func (referenceAction *ReferenceAction) Execute() (string, error) {
	var markdownRenderer = models.NewMarkdownRenderer()
	err := referenceAction.RequireAmountOfParameters(1)
	if err != nil {
		return "", err
	}
	var actionName = referenceAction.controller.Arguments[0]
	referenceAction.logger.Print(helpers.GetBannerText())
	if referenceAction.environment.IsColorEnabled() {
		markdownRenderer.EnableColor()
	}
	action := referenceAction.controller.actionsMap[actionName]
	template := fasttemplate.New(helpHeader+action.GetHelp(), "{{", "}}")
	markdownContent := template.ExecuteString(map[string]interface{}{
		"actionName": actionName,
		"binaryName": referenceAction.environment.GetArguments()[0],
		"brief":      action.GetBrief(),
	})
	markdownRenderer.Render(markdownContent)

	fmt.Printf("  %s <command> [<arguments>]\n\n", referenceAction.environment.GetArguments()[0])
	for _, action := range referenceAction.controller.GetActions() {
		referenceAction.logger.Print(fmt.Sprintf("  %s\t%+v", referenceAction.rightPadString(action.GetName(), 10), action.GetBrief()))
	}
	fmt.Print("\n  Set the environment variable CLICOLOR to 1 to enable colors.\n")
	return "", nil
}
