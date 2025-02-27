package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

type Controller struct {
	environment *models.Environment
	logger      *helpers.Logger
	Arguments   []string
	model       *models.BuilderModel
	actions     []Action
	actionsMap  map[string]Action
	helpAction  *HelpAction
}

func NewController(environment *models.Environment) *Controller {
	controller := &Controller{
		environment: environment,
		logger:      environment.GetLogger(),
		model:       models.NewBuilderModel(environment),
		actionsMap:  make(map[string]Action),
	}

	var helpAction = NewHelpAction(controller)
	controller.actions = []Action{
		helpAction,
	}
	controller.actionsMap[helpAction.GetName()] = helpAction

	var arguments []string
	arguments = []string{}
	if len(environment.GetArguments()) > 2 {
		arguments = environment.GetArguments()[2:]
	}

	controller.Arguments = arguments
	controller.helpAction = helpAction

	return controller
}

func (this *Controller) AddAction(action Action) {
	this.actions = append(this.actions, action)
	this.actionsMap[action.GetName()] = action
}

func (this *Controller) GetEnvironment() *models.Environment {
	return this.environment
}

func (this *Controller) GetActionsMap() map[string]Action {
	return this.actionsMap
}

func (this *Controller) GetActions() []Action {
	return this.actions
}

func (this *Controller) ExecuteAction() {
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

func (this *Controller) ShowHelp() {
	this.helpAction.Execute()
}

func (this *Controller) GetArguments() []string {
	return this.Arguments
}
