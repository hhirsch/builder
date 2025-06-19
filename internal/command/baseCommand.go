package command

import ()

type BaseCommand struct {
	commandName        string
	result             string
	name               string
	parameters         int
	requiresConnection bool
	requirements       []string
}

func NewBaseCommand() *BaseCommand {
	return &BaseCommand{
		commandName:  "undefined",
		result:       "",
		requirements: []string{}, // the binaries that need to be present on the target system
	}
}

func (baseCommand *BaseCommand) TestRequirements() bool {
	return true
}

func (baseCommand *BaseCommand) GetName() string {
	return baseCommand.name
}

func (baseCommand *BaseCommand) RequiresConnection() bool {
	return baseCommand.requiresConnection
}
