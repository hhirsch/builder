package main

import (
	"github.com/hhirsch/builder/internal/controllers"
	"github.com/hhirsch/builder/internal/models"
)

var environment *models.Environment = models.NewEnvironment()

func main() {
	controller := controllers.NewRegistryController(environment)
	controller.ExecuteAction()
}
