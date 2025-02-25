package controllers

import (
	_ "embed"
	"github.com/charmbracelet/huh"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/traits"
	"github.com/valyala/fasttemplate"
	"strings"
)

type CreateAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	model       *models.BuilderModel
	controller  *BuilderController
	BaseAction
	traits.FileSystem
}

//go:embed structTemplate.txt
var structTemplate string

func NewCreateAction(controller *BuilderController) *CreateAction {

	return &CreateAction{
		BaseAction:  BaseAction{controller: controller},
		FileSystem:  traits.FileSystem{},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment()),
		controller:  controller,
	}

}

func (this *CreateAction) Execute() {
	this.logger.Print(helpers.GetBannerText())

	if this.ParameterValidationFailed(1, "create needs a file path as parameter.") {
		return
	}
	var structName string
	huh.NewInput().
		Title("Name the struct").
		Value(&structName).
		Run()

	var packageName string
	huh.NewInput().
		Title("Name the package").
		Value(&packageName).
		Run()

	if !strings.HasSuffix(this.controller.Arguments[0], ".go") {
		this.logger.Fatal("File ending .go not found in parameter.")
	}

	template := fasttemplate.New(structTemplate, "{{", "}}")
	fileContent := template.ExecuteString(map[string]interface{}{
		"packageName": packageName,
		"structName":  structName,
	})
	this.WriteStringToFile(this.controller.Arguments[0], fileContent)
}

func (this *CreateAction) GetName() string {
	return "create"
}

func (this *CreateAction) GetDescription() string {
	return "Create a struct in the project."
}

func (this *CreateAction) GetHelp() string {
	return "Create a struct in the project."
}
