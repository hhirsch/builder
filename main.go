package main

import (
	"github.com/hhirsch/builder/internal/controllers"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"os"
	"reflect"
	"strings"
)

var arguments []string = os.Args
var environment *models.Environment = models.NewEnvironment()
var logger *helpers.Logger = environment.GetLogger()

/**
 * This is the command line interface for the server maintenance tool builder
 **/
func main() {
	controller := controllers.NewBuilderController(environment)
	if len(arguments) < 2 {
		logger.Info(helpers.GetCommandNameRequiredText())
		controller.HelpAction()
		return
	}

	actionName := strings.ToUpper(os.Args[1][:1]) + strings.ToLower(os.Args[1][1:])
	interfaceValue := reflect.ValueOf(controller)
	method := interfaceValue.MethodByName(actionName + "Action")

	if method.IsValid() && method.Kind() == reflect.Func {
		method.Call(nil)
	} else {
		logger.Info("Builder called with unrecognized parameter " + os.Args[1] + ".")
		controller.HelpAction()

	}
}
