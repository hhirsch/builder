package main

import (
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

func main() {
	var environment *models.Environment = models.NewEnvironment()

	var interpreter interpreter.Interpreter = *interpreter.NewInterpreter(environment)
	interpreter.Run("./cmd/interpreterDemo/example.bld")
}
