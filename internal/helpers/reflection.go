package helpers

import (
	"fmt"
	"reflect"
)

type ReflectionHelper struct {
}

func NewReflectionHelper(fileName string) *ReflectionHelper {
	return &ReflectionHelper{}
}

// Function to call a method based on method name as a string
func (this *ReflectionHelper) CallMethodByName(instance interface{}, methodName string) {
	// Get the reflect.Value of the instance
	v := reflect.ValueOf(instance)

	// Get the method by name
	method := v.MethodByName(methodName)

	// Check if method exists and is callable
	if method.IsValid() && method.Kind() == reflect.Func {
		// Call the method
		method.Call(nil)
	} else {
		fmt.Printf("Method %s not found or is not callable", methodName)
	}
}
