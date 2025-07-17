package main

import (
	"github.com/hhirsch/builder/internal/controllers"
	"github.com/hhirsch/builder/internal/environment"
	"github.com/hhirsch/builder/internal/models"
	"log/slog"
)

var appEnvironment *models.Environment = models.NewEnvironment()
var registry *models.Registry = models.NewRegistry(appEnvironment.GetBuilderHomePath() + "/builderGlobal.reg")
var testing bool = true

func runControllerAction() {
	slog.Info("Running controller action.")
	controller := controllers.NewBuilderController(appEnvironment)
	controller.ExecuteAction()
}

func setupLogging() {
	logFilePath, error := environment.GetLogFilePath()
	if testing || error == nil {
		environment.SetupLoggingForTesting()
		slog.Info("Logging set up for testing.")
		return
	}
	environment.SetupLoggingForProduction(logFilePath)
	slog.Info("Logging set up for production.")
}

/**
 * This is the command line interface for the server maintenance tool builder
 **/
func main() {
	setupLogging()
	runControllerAction()
	slog.Info("Done")
}
