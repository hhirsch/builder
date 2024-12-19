package models

import (
	_ "embed"
	"fmt"
	"github.com/valyala/fasttemplate"
)

//go:embed mysql-my-conf.txt
var myConfigTemplate string

type MySQL struct {
	data     map[string]string
	fileName string
}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (this *MySQL) GetMyConfigPath() string {
	return "~/.my.cfg"
}

func (this *MySQL) GetCredentialsFileContent(user string, password string) string {
	template := fasttemplate.New(myConfigTemplate, "{{", "}}")
	myConfig := template.ExecuteString(map[string]interface{}{
		"user":     user,
		"password": password,
	})

	return myConfig

	//chmod 0600 .my.cnf
	//chmod 0400 .my.cnf read only

}

func (this *MySQL) DeleteCredentialsFileCommand() string {
	return "rm " + this.GetMyConfigPath()
}

func (this *MySQL) EnsureReadOnlyFilePermissionCommand() string {
	return "chmod 0400 " + this.GetMyConfigPath()
}

func (this *MySQL) GetListDatabasesCommand() string {
	return "mysql -u root -e \"SHOW DATABASES;\""
}

func (this *MySQL) DumpDatabaseCommand(databaseName string, fileName string) string {
	//return "mysqldump -u username -p database_name > database_dump.sql"
	//return fmt.Sprintf("mysqldump -u %s -p%s %s > %s", username, password, databaseName, outputFile)
	//mysqldump --socket=/tmp/mysqlsecond.sock --all-databases > $sqlfile
	//mysqldump --socket=/var/run/mysqld/mysqld.sock database_name > dump_file.sql
	return fmt.Sprintf("mysqldump --socket=/var/run/mysqld/mysqld.sock %s > %s", databaseName, fileName)
}

func (this *MySQL) InstallDatabase(databaseName string, fileName string) {

}

func (this *MySQL) CheckDatabaseExists(databaseName string) {

}

func (this *MySQL) MoveDatabase(sourceHostHandle string, targetHostHandle string) {

}
