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

func (this *SqlCommand) uploadSqlCredentials() (err error) {
	userName, err := this.PromptEncryptedIfMissing("mysql.user")
	if err != nil {
		return
	}

	password, err := this.PromptEncryptedIfMissing("mysql.password")
	if err != nil {
		return
	}
	filePath, err := this.WriteStringToTempFile(this.mysql.GetCredentialsFileContent(userName, password))
	if err != nil {
		return
	}

	if password == "" || userName == "" {
		return fmt.Errorf("Empty input for credentials received aborting SQL command.")
	}
	this.environment.Client.Upload(filePath, this.mysql.GetMyConfigPath())
	return
}

func (this *SqlCommand) ExecuteSqlCommand(command string) string {
	return this.environment.Client.Execute(fmt.Sprintf("mysql -u root -e \"%v;\"", command))
}

func (this *SqlCommand) wipeSqlCredentialsFromServer() {
	this.environment.Client.Execute(fmt.Sprintf("rm %v", this.mysql.GetMyConfigPath()))
}
