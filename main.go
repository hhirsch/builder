package main

import (
	"github.com/hhirsch/builder/internal/controllers"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"os"
)

var arguments []string = os.Args
var environment *models.Environment = models.NewEnvironment()
var logger *helpers.Logger = environment.GetLogger()

/**
 * This is the command line interface for the server maintenance tool builder
 **/
func main() {
	controller := controllers.NewBuilderController(environment)
	if controller.ParameterValidationFailed(2, helpers.GetCommandNameRequiredText()) {
		controller.Help()
		return
	}
	commands := map[string]func(){
		"init":    controller.Init,
		"update":  controller.Update,
		"help":    controller.Help,
		"command": controller.Command,
	}

	commandName := os.Args[1]
	if command, exists := commands[commandName]; exists {
		command()
	} else {
		logger.Info("Unrecognized parameter " + commandName)
		controller.Help()
	}
}
