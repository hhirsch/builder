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
	controller  *Controller
	BaseAction
	traits.FileSystem
}

//go:embed structTemplate.txt
var structTemplate string

//go:embed createAction.md
var createMarkdown string

func NewCreateAction(controller *Controller) *CreateAction {
	baseAction := BaseAction{
		controller:  controller,
		name:        "create",
		description: "Create a struct from template.",
		brief:       "Create a struct from template.",
		help:        createMarkdown,
	}
	return &CreateAction{
		BaseAction:  baseAction,
		FileSystem:  traits.FileSystem{},
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		model:       models.NewBuilderModel(controller.GetEnvironment().GetProjectPath()),
		controller:  controller,
	}

}

func (createAction *CreateAction) Execute() (string, error) {
	createAction.logger.Print(helpers.GetBannerText())
	err := createAction.RequireAmountOfParameters(1)
	if err != nil {
		return "", err
	}

	var structName string
	structNameInput := huh.NewInput().
		Title("Name the struct").
		Value(&structName)
	err = structNameInput.Run()
	if err != nil {
		return "", err
	}

	var packageName string
	packageNameInput := huh.NewInput().
		Title("Name the package").
		Value(&packageName)
	err = packageNameInput.Run()
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(createAction.controller.Arguments[0], ".go") {
		createAction.logger.Fatal("File ending .go not found in parameter.")
	}

	template := fasttemplate.New(structTemplate, "{{", "}}")
	fileContent := template.ExecuteString(map[string]interface{}{
		"packageName": packageName,
		"structName":  structName,
	})
	err = createAction.WriteStringToFile(createAction.controller.Arguments[0], fileContent)
	if err != nil {
		createAction.logger.Fatalf("Error writing to file: %s", err.Error())
	}
	return "", nil
}

func (createAction *CreateAction) GetName() string {
	return "create"
}
