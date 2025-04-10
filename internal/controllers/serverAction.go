package controllers

import (
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
)

type ServerAction struct {
	environment *models.Environment
	logger      *helpers.Logger
	controller  *Controller
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
		environment: controller.GetEnvironment(),
		logger:      controller.GetEnvironment().GetLogger(),
		controller:  controller,
	}

}

func (serverAction *ServerAction) Execute() (string, error) {
	return "", nil
}

func (serverAction *ServerAction) GetDescription() string {
	return "Manage the list of servers."
}
