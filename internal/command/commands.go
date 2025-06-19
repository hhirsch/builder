package command

import (
	"fmt"
)

type Commands struct {
	commands map[string]Command
}

func NewCommands() *Commands {
	return &Commands{
		commands: map[string]Command{},
	}
}

func (commands *Commands) AddCommand(command Command) {
	commands.commands[command.GetName()] = command
}

func (commands *Commands) GetCommand(commandName string) (*Command, error) {
	command, exists := commands.commands[commandName]

	if !exists {
		return nil, fmt.Errorf("No such command: %v", commandName)
	}
	return &command, nil
}
