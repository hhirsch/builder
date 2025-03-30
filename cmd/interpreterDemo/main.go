package main

import (
	"fmt"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/interpreter"
)

func main() {
	var environment = models.NewEnvironment()
	var interpreter = *interpreter.NewInterpreter(environment)
	err := interpreter.Run("./cmd/interpreterDemo/example.bld")
	if err != nil {
		fmt.Print(err.Error())
	}
}
