package commands

import (
	"fmt"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/models/traits"
)

type SqlCommand struct {
	environment *models.Environment
	mysql       *models.MySql
	traits.FileSystem
	traits.HostRegistry
}

func NewSqlCommand(environment *models.Environment) *SqlCommand {
	return &SqlCommand{
		environment:  environment,
		mysql:        models.NewMySql(),
		FileSystem:   traits.FileSystem{},
		HostRegistry: *traits.NewHostRegistry(environment),
	}
}

func (sqlCommand *SqlCommand) uploadSqlCredentials() (err error) {
	userName, err := sqlCommand.PromptEncryptedIfMissing("mysql.user")
	if err != nil {
		return
	}

	password, err := sqlCommand.PromptEncryptedIfMissing("mysql.password")
	if err != nil {
		return
	}
	filePath, err := sqlCommand.WriteStringToTempFile(sqlCommand.mysql.GetCredentialsFileContent(userName, password))
	if err != nil {
		return
	}

	if password == "" || userName == "" {
		return fmt.Errorf("empty input for credentials received aborting SQL command")
	}
	sqlCommand.environment.Client.Upload(filePath, sqlCommand.mysql.GetMyConfigPath())
	return
}

func (sqlCommand *SqlCommand) ExecuteSqlCommand(command string) string {
	return sqlCommand.environment.Client.Execute(fmt.Sprintf("mysql -u root -e \"%v;\"", command))
}

func (sqlCommand *SqlCommand) wipeSqlCredentialsFromServer() {
	sqlCommand.environment.Client.Execute(fmt.Sprintf("rm %v", sqlCommand.mysql.GetMyConfigPath()))
}
