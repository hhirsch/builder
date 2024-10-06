package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
	"os"
)

// controller for the builder command line
type BuilderController struct {
	environment *models.Environment
	logger      *helpers.Logger
	arguments   []string
	model       *models.BuilderModel
}

func NewBuilderController(environment *models.Environment) *BuilderController {

	controller := &BuilderController{
		environment: environment,
		logger:      environment.GetLogger(),
		arguments:   environment.GetArguments(),
		model:       models.NewBuilderModel(environment),
	}
	return controller
}

func (this *BuilderController) ParameterValidationFailed(requiredAmountOfParameters int, errorMessage string) bool {
	return len(this.arguments) < requiredAmountOfParameters
}

// Initialize builder in current directory
func (this *BuilderController) Init() {
	if this.model.IsInitialized() {
		this.logger.Info("Already initialized, nothing to do.")
	}
	this.model.CreateDirectories()
	this.logger.Info("Initializing")
}

// Execute builder code from file
func (this *BuilderController) Update() {
	this.logger.Print(helpers.GetBannerText())
	if this.ParameterValidationFailed(3, "update needs a file name as argument") {
		return
	}
	this.logger.Info("Builder started")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.Run(os.Args[2])
}

// run custom builder command
func (this *BuilderController) Command() {
	if this.ParameterValidationFailed(3, "command needs a command name as argument") {
		this.Help()
		return
	}
	this.logger.Print("executing user defined command")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.Run("./.builder/commands/" + os.Args[2] + ".bld")
}

// show help
func (this *BuilderController) Help() {
	this.logger.Print(helpers.GetBannerText() + "\n" + helpers.GetHelpText())
}
