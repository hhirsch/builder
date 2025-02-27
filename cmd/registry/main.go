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
	controller := controllers.NewRegistryController(environment)
	controller.ExecuteAction()
}
