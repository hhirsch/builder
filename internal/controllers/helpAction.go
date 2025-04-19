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
		controller:  controller,
	}

}

func (helpAction *HelpAction) rightPadString(string string, length int) string {
	if len(string) >= length {
		return string
	}
	padLength := length - len(string)
	pad := strings.Repeat(" ", padLength)
	return string + pad
}

func (helpAction *HelpAction) Execute() (string, error) {
	var markdownRenderer = models.NewMarkdownRenderer()
	buffer := helpers.GetBannerText()
	if helpAction.HasEnoughParameters(1) {
		var actionName = helpAction.controller.Arguments[0]
		if helpAction.environment.IsColorEnabled() {
			markdownRenderer.EnableColor()
		}
		action := helpAction.controller.actionsMap[actionName]
		template := fasttemplate.New(helpHeader+action.GetHelp(), "{{", "}}")
		markdownContent := template.ExecuteString(map[string]interface{}{
			"actionName": actionName,
			"binaryName": helpAction.environment.GetArguments()[0],
			"brief":      action.GetBrief(),
		})
		markdownRenderer.Render(markdownContent)
		return "", nil
	}

	buffer += fmt.Sprintf("  %s <command> [<arguments>]\n\n", helpAction.environment.GetArguments()[0])
	for _, action := range helpAction.controller.GetActions() {
		buffer += fmt.Sprintf("  %s\t%+v\n", helpAction.rightPadString(action.GetName(), 10), action.GetBrief())
	}
	buffer += "\n  Set the environment variable CLICOLOR to 1 to enable colors.\n"
	fmt.Print(buffer)
	return buffer, nil
}

func (helpAction *HelpAction) GetDescription() string {
	return "Create builder directories in " + helpAction.environment.GetProjectPath() + "."
}
