package main

import (
	"os"
	"server/controllers"
	"server/helpers"
)

/**
 * This is the command line interface for the server maintenance tool called builder
 **/
func main() {
	environment := NewEnvironment()
	logger := environment.logger
	if len(os.Args) < 2 {
		logger.Print("Please provide a command name as an argument")
		logger.Print(helpers.GetHelpText())
		return
	}
	commandName := os.Args[1]
	switch commandName {
	case "init":
		logger.Info(controllers.Init())
	case "update":
		logger.Print(helpers.GetBannerText())
		if len(os.Args) < 3 {
			logger.Print("update needs a file name as argument")
			logger.Print(helpers.GetHelpText())
			return
		}
		logger.Info("Builder started")
		var interpreter Interpreter = *NewInterpreter()
		interpreter.run(os.Args[2])
	case "help":
		logger.Print(helpers.GetHelpText())
	default:
		logger.Info("Unrecognized parameter")
	}
}
