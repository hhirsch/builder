package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
)

type ServerAction struct {
	logger     *helpers.Logger
	controller *Controller
	BaseAction
}

/*
 * Eg server list, server add, server alias, server require {serviceName} {serverName}
 * service list, service health, service health {serviceName}, service install {serviceName}
 */
func NewServerAction(controller *Controller) *ServerAction {

	return &ServerAction{
		BaseAction: BaseAction{
			controller: controller,
			name:       "server",
			help:       "Manage the list of servers.",
		},
		controller: controller,
	}

}

func (serverAction *ServerAction) Execute() (string, error) {
	return "", nil
}
