package main

import (
	"fmt"
	"github.com/hhirsch/builder/internal/interpreter"
	"github.com/hhirsch/builder/internal/models"
)

func main() {
	var environment = models.NewEnvironment()
	interpreterObject, err := interpreter.NewInterpreter(environment)
	if err != nil {
		fmt.Print(err.Error())
	}
	interpreter := *interpreterObject
	err = interpreter.Run("./cmd/interpreterDemo/example.bld")
	if err != nil {
		fmt.Print(err.Error())
	}
}
