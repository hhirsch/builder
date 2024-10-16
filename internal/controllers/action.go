package controllers

type Action interface {
	Execute(controller *BuilderController)
	GetName() string
	GetDescription() string
	GetHelp() string
}
