package controllers

import (
	"fmt"
	"os"
)

func isInitialized() bool {
	dirName := ".builder"

	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return false
	}
	return true
}

func createDirectories() {
	dirName := ".builder"
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}

func Init() string {
	if isInitialized() {
		return "Already initialized, nothing to do."
	}
	createDirectories()
	return "Initializing"
}
