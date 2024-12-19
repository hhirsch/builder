package controllers

import (
	"fmt"
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
	initAction    *InitAction
	scriptAction  *ScriptAction
	commandAction *CommandAction
	actionsMap    map[string]Action
}

func NewBuilderController(environment *models.Environment) *BuilderController {
	var initAction = NewInitAction(environment)
	var scriptAction = NewScriptAction(environment)
	var commandAction = NewCommandAction(environment)
	var actions = []Action{
		initAction,
		scriptAction,
		commandAction,
	}

	actionsMap := map[string]Action{
		"init":    initAction,
		"script":  scriptAction,
		"command": commandAction,
	}

	var arguments []string
	if len(environment.GetArguments()) > 2 {
		arguments = environment.GetArguments()[2:]
	} else {
		arguments = []string{}
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
	}

	return controller
}
func (this *BuilderController) ExecuteAction() {
	this.Arguments = []string{}
	if len(this.environment.GetArguments()) < 2 {
		this.logger.Info("Please provide a command name as an argument.")
		this.HelpAction()
		return
	}

	this.Arguments = this.environment.GetArguments()[2:]

	if action, exists := this.actionsMap[this.environment.GetArguments()[1]]; exists {
		action.Execute(this)
	} else {
		this.logger.Info("Builder called with unrecognized parameter " + this.environment.GetArguments()[1] + ".")
		this.HelpAction()
	}
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

// Initialize builder in current directory
func (this *BuilderController) InitAction() {
	this.initAction.Execute(this)
}

// Execute builder code from file
func (this *BuilderController) ScriptAction() {
	this.scriptAction.Execute(this)
}

func (this *BuilderController) CommandAction() {
	this.commandAction.Execute(this)
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
	if this.HasEnoughParameters(1) {
		this.logger.Print("tbd show specific help for " + os.Args[2])
	}
	this.logger.Print(helpers.GetBannerText())
	this.logger.Print("help\t\tShow this help. Call help with an action name as parameter \n\t\tto get more details on the action.")
	for _, action := range this.actions {
		this.logger.Print(fmt.Sprintf("%s\t\t%+v", action.GetName(), action.GetHelp()))
	}
}
