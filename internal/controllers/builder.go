package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
	"os"
)

// controller for the builder command line
type BuilderController struct {
	environment   *models.Environment
	logger        *helpers.Logger
	Arguments     []string
	model         *models.BuilderModel
	actions       []Action
	actionsMap    map[string]Action
	initAction    *InitAction
	scriptAction  *ScriptAction
	commandAction *CommandAction
	helpAction    *HelpAction
}

func NewBuilderController(environment *models.Environment) *BuilderController {
	var initAction = NewInitAction(environment)
	var scriptAction = NewScriptAction(environment)
	var commandAction = NewCommandAction(environment)
	var helpAction = NewHelpAction(environment)
	var actions = []Action{
		initAction,
		scriptAction,
		commandAction,
		helpAction,
	}

	actionsMap := map[string]Action{
		initAction.GetName():    initAction,
		scriptAction.GetName():  scriptAction,
		commandAction.GetName(): commandAction,
		helpAction.GetName():    helpAction,
	}

	var arguments []string
	arguments = []string{}
	if len(environment.GetArguments()) > 2 {
		arguments = environment.GetArguments()[2:]
	}

	controller := &BuilderController{
		environment:   environment,
		logger:        environment.GetLogger(),
		model:         models.NewBuilderModel(environment),
		Arguments:     arguments,
		actions:       actions,
		actionsMap:    actionsMap,
		initAction:    initAction,
		scriptAction:  scriptAction,
		commandAction: commandAction,
		helpAction:    helpAction,
	}

	return controller
}

func (this *BuilderController) GetActionsMap() map[string]Action {
	return this.actionsMap
}

func (this *BuilderController) GetActions() []Action {
	return this.actions
}

func (this *BuilderController) ExecuteAction() {
	if len(this.Arguments) < 1 {
		this.logger.Info("Please provide a command name as an argument.")
		this.HelpAction()
		return
	}

	if action, exists := this.actionsMap[this.environment.GetArguments()[1]]; exists {
		action.Execute(this)
		return
	}
	this.logger.Info("Builder called with unrecognized parameter " + this.Arguments[0] + ".")
	this.HelpAction()
}

func (this *BuilderController) ParameterValidationFailed(requiredAmountOfParameters int, errorMessage string) bool {
	if !this.HasEnoughParameters(requiredAmountOfParameters) {
		this.logger.Fatal(errorMessage)
	}
	return false
}

func (this *BuilderController) HasEnoughParameters(requiredAmountOfParameters int) bool {
	return len(this.Arguments) >= requiredAmountOfParameters
}

func (this *BuilderController) UpdateAction() {
	if this.ParameterValidationFailed(1, "command needs a command name as argument") {
		this.HelpAction()
		return
	}
	this.logger.Print("executing user defined command")
	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(this.environment)
	interpreter.Run("./.builder/commands/" + os.Args[2] + ".bld")
}

// show help
func (this *BuilderController) HelpAction() {
	this.helpAction.Execute(this)
}
