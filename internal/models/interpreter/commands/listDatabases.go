package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type ListDatabasesCommand struct {
	environment *models.Environment
	BaseCommand
	SqlCommand
}

func NewListDatabasesCommand(environment *models.Environment) *ListDatabasesCommand {
	controller := &ListDatabasesCommand{
		environment: environment,
		SqlCommand:  *NewSqlCommand(environment),
		BaseCommand: BaseCommand{
			environment:        environment,
			name:               "listDatabases",
			requiresConnection: true,
		},
	}
	return controller
}

func (this *ListDatabasesCommand) TestRequirements() bool {
	return this.FindBinary("mysql")
}

func (this *ListDatabasesCommand) Execute(tokens []string) string {
	err := this.uploadSqlCredentials()
	if err != nil {
		this.environment.GetLogger().Fatalf("Error uploading SQL credentials: %v", err)
	}
	this.environment.GetLogger().Info(this.ExecuteSqlCommand(this.mysql.GetListDatabasesQuery()))
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
