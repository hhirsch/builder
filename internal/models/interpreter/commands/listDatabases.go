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

func (listDatabasesCommand *ListDatabasesCommand) TestRequirements() bool {
	return listDatabasesCommand.FindBinary("mysql")
}

func (listDatabasesCommand *ListDatabasesCommand) Execute(tokens []string) string {
	err := listDatabasesCommand.uploadSqlCredentials()
	if err != nil {
		listDatabasesCommand.environment.GetLogger().Fatalf("Error uploading SQL credentials: %v", err)
	}
	listDatabasesCommand.environment.GetLogger().Info(listDatabasesCommand.ExecuteSqlCommand(listDatabasesCommand.mysql.GetListDatabasesQuery()))
	listDatabasesCommand.wipeSqlCredentialsFromServer()
	return ""
}

func (listDatabasesCommand *ListDatabasesCommand) Undo() {
	listDatabasesCommand.environment.GetLogger().Info("Nothing to undo.")
}

func (listDatabasesCommand *ListDatabasesCommand) GetDescription(tokens []string) string {
	return "Lists all databases."
}

func (listDatabasesCommand *ListDatabasesCommand) GetHelp() string {
	return "[]\tLists all databases."
}
