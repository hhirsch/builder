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
		model:       models.NewBuilderModel(environment.GetProjectPath()),
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

func (controller *Controller) AddAction(action Action) {
	controller.actions = append(controller.actions, action)
	controller.actionsMap[action.GetName()] = action
}

func (controller *Controller) GetEnvironment() *models.Environment {
	return controller.environment
}

func (controller *Controller) GetActionsMap() map[string]Action {
	return controller.actionsMap
}

func (controller *Controller) GetActions() []Action {
	return controller.actions
}

func (controller *Controller) ExecuteAction() {
	if len(controller.Arguments) < 1 {
		controller.logger.Info("You need to pass a command name as argument.")
		controller.ShowHelp()
		return
	}

	var actionName = controller.environment.GetArguments()[1]
	if action, exists := controller.actionsMap[actionName]; exists {
		_, err := action.Execute()
		if err != nil {
			controller.logger.Errorf("action failed %s", err.Error())
		}
		return
	}
	controller.logger.Info("Builder called with unrecognized command " + actionName + ".")
	controller.ShowHelp()
}

func (controller *Controller) ShowHelp() {
	_, err := controller.helpAction.Execute()
	if err != nil {
		controller.logger.Errorf("show help failed %v", err)
	}
}

func (controller *Controller) GetArguments() []string {
	return controller.Arguments
}
