package controllers

import (
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"log/slog"
	"os"
	"strings"
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

func (controller *Controller) GetActionsMap() map[string]Action {
	return controller.actionsMap
}

func (controller *Controller) GetActions() []Action {
	return controller.actions
}

func (controller *Controller) ExecuteAction() {
	if len(os.Args) < 2 {
		slog.Error("Controller called without command name argument.", slog.String("arguments", strings.Join(os.Args, " ")))
		fmt.Println("You need to pass a command name as argument.")
		controller.ShowHelp()
		return
	}

	var actionName = os.Args[1]
	if action, exists := controller.actionsMap[actionName]; exists {
		slog.Info("Builder called with recognized command: " + actionName + ".")
		_, err := action.Execute()
		if err != nil {
			slog.Error("Executing action.", slog.String("error message", err.Error()))
			fmt.Printf("Error executing action: %v.\n", err.Error())
			controller.ShowHelp()
		}
		return
	}
	slog.Info("Builder called with unrecognized command: " + actionName + ".")
	controller.ShowHelp()
}

func (controller *Controller) ShowHelp() {
	_, err := controller.helpAction.Execute()
	if err != nil {
		slog.Error("Unable to show help.", slog.String("error message", err.Error()))
	}
}

func (controller *Controller) GetArguments() []string {
	return controller.Arguments
}
