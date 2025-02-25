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
		SqlCommand:  *NewSqlCommand(environment),
		BaseCommand: *NewBaseCommand(environment),
	}
	return controller
}

func (this *ListDatabasesCommand) TestRequirements() bool {
	return this.FindBinary("mysql")
}

func (this *ListDatabasesCommand) Execute(tokens []string) string {
	this.uploadSqlCredentials()
	this.environment.GetLogger().Info(this.environment.Client.Execute(this.mysql.GetListDatabasesCommand()))
	this.wipeSqlCredentialsFromServer()
	return ""
}

func (this *ListDatabasesCommand) Undo() {
	this.environment.GetLogger().Info("Nothing to undo.")
}

func (this *ListDatabasesCommand) GetDescription(tokens []string) string {
	return "Lists all databases."
}

func (this *ListDatabasesCommand) GetHelp() string {
	return "[]\tLists all databases."
}
