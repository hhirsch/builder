package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

// controller for the builder command line
type RegistryController struct {
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
	createAction  *CreateAction
}

func NewRegistryController(environment *models.Environment) *BuilderController {
	controller := &BuilderController{
		environment: environment,
		logger:      environment.GetLogger(),
		model:       models.NewBuilderModel(environment),
	}

	var initAction = NewInitAction(controller)
	var scriptAction = NewScriptAction(controller)
	var helpAction = NewHelpAction(controller)
	var actions = []Action{
		scriptAction,
		helpAction,
	}

	actionsMap := map[string]Action{
		scriptAction.GetName(): initAction,
		scriptAction.GetName(): scriptAction,
	}

	var arguments []string
	arguments = []string{}
	if len(environment.GetArguments()) > 2 {
		arguments = environment.GetArguments()[2:]
	}

	controller.Arguments = arguments
	controller.actions = actions
	controller.actionsMap = actionsMap
	controller.helpAction = helpAction

	return controller
}

func (this *RegistryController) GetEnvironment() *models.Environment {
	return this.environment
}

func (this *RegistryController) GetActionsMap() map[string]Action {
	return this.actionsMap
}

func (this *RegistryController) GetActions() []Action {
	return this.actions
}

func (this *RegistryController) ExecuteAction() {
	if len(this.Arguments) < 1 {
		this.logger.Info("You need to pass a command name as argument.")
		this.ShowHelp()
		return
	}

	var actionName = this.environment.GetArguments()[1]
	if action, exists := this.actionsMap[actionName]; exists {
		action.Execute()
		return
	}
	this.logger.Info("Builder called with unrecognized command " + actionName + ".")
	this.ShowHelp()
}

func (this *RegistryController) ShowHelp() {
	this.helpAction.Execute()
}
