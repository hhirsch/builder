package controllers

type Action interface {
	Execute()
	GetName() string
	GetDescription() string
	GetHelp() string
	GetBrief() string
}
