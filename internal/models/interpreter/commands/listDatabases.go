package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type ListDatabasesCommand struct {
	environment *models.Environment
	description string
	BaseCommand
	SqlCommand
}

func NewListDatabasesCommand(environment *models.Environment) *ListDatabasesCommand {
	controller := &ListDatabasesCommand{
		environment: environment,
		BaseCommand: BaseCommand{environment: environment},
	}
	return controller
}

func (this *ListDatabasesCommand) TestRequirements() bool {
	return this.FindBinary("mysql")
}

func (this *ListDatabasesCommand) Execute(tokens []string) {
	//	this.environment.GetLogger().Info(this.environment.Client.Execute("pwd"))
	//this.environment.GetLogger().Info(this.environment.Client.Execute("mysql -u root -e \"SHOW DATABASES;\" --protocol=socket"))
	this.environment.GetLogger().Info(this.environment.Client.Execute("mysql --socket=/var/run/mysqld/mysqld.sock -u root -e \"SHOW DATABASES;\""))
}

func (this *ListDatabasesCommand) Undo() {
	this.environment.GetLogger().Info("Nothing to undo.")
}

func (this *ListDatabasesCommand) GetDescription(tokens []string) string {
	return "Ensure a binary is allowed to open ports."
}

func (this *ListDatabasesCommand) GetHelp() string {
	return "[source <string>, target <string>]\tPush a file to the server."
}
