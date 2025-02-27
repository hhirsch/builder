package main

import (
	"fmt"
	"github.com/hhirsch/builder/internal/models"
)

func main() {
	filePath := "/tmp/demo.reg"
	registry := models.NewRegistry(filePath)
	registry.Load()
	registry.Register("Registry.BEE.Bar", "dataxx")
	registry.Register("Registry.BEE.Baz", "more dataxx")
	error := registry.Save()
	if error != nil {
		fmt.Println("Error saving data:", error)
		return
	}

	error = registry.Load()
	if error != nil {
		fmt.Println("Error loading data:", error)
		return
	}
	loadedData := registry.GetRegistryDump()
	fmt.Println("Loaded data:", loadedData)
}
