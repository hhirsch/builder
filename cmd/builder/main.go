package main

import (
	"github.com/hhirsch/builder/internal/controllers"
	"github.com/hhirsch/builder/internal/models"
)

var environment *models.Environment = models.NewEnvironment()

/**
 * This is the command line interface for the server maintenance tool builder
 **/
func main() {
	controller := controllers.NewBuilderController(environment)
	controller.ExecuteAction()
}
