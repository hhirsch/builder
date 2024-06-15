package main

type Environment struct {
	configPath string
	logger     *Logger
}

func NewEnvironment() *Environment {
	logger := NewLogger()
	environment := &Environment{
		logger: logger,
	}
	return environment
}
