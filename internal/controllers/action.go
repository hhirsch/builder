package controllers

type Action interface {
	Execute() (string, error)
	GetName() string
	GetDescription() string
	GetHelp() string
	GetBrief() string
}
