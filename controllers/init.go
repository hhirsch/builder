package controllers

func isInitialized() bool {
	return true
}

func createDirectories() {}

func Init() string {
	if isInitialized() {
		return "Already initialized, nothing to do."
	}
	createDirectories()
	return "Initializing"
}
