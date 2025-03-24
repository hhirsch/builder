package models

import (
	_ "embed"
	"fmt"
	"github.com/valyala/fasttemplate"
)

//go:embed mysql-my-conf.txt
var myConfigTemplate string

type MySql struct {
}

func NewMySql() *MySql {
	return &MySql{}
}

func (this *MySql) GetMyConfigPath() string {
	return "~/.my.cfg"
}

func (this *MySql) GetCredentialsFileContent(user string, password string) string {
	template := fasttemplate.New(myConfigTemplate, "{{", "}}")
	myConfig := template.ExecuteString(map[string]interface{}{
		"user":     user,
		"password": password,
	})

	return myConfig

	//chmod 0600 .my.cnf
	//chmod 0400 .my.cnf read only

}

func (this *MySql) DeleteCredentialsFileCommand() string {
	return "rm " + this.GetMyConfigPath()
}

func (this *MySql) EnsureReadOnlyFilePermissionCommand() string {
	return "chmod 0600 " + this.GetMyConfigPath()
}

func (this *MySql) GetListDatabasesCommand() string {
	return "mysql -u root -e \"SHOW DATABASES;\""
}

func (this *MySql) GetListDatabasesQuery() string {
	return "SHOW DATABASES;"
}

func (this *MySql) DumpDatabaseCommand(databaseName string, fileName string) string {
	//return "mysqldump -u username -p database_name > database_dump.sql"
	//return fmt.Sprintf("mysqldump -u %s -p%s %s > %s", username, password, databaseName, outputFile)
	//mysqldump --socket=/tmp/mysqlsecond.sock --all-databases > $sqlfile
	//mysqldump --socket=/var/run/mysqld/mysqld.sock database_name > dump_file.sql
	return fmt.Sprintf("mysqldump --socket=/var/run/mysqld/mysqld.sock %s > %s", databaseName, fileName)
}

func (this *MySql) InstallDatabase(databaseName string, fileName string) {

}

func (this *MySql) CheckDatabaseExists(databaseName string) {

}

func (this *MySql) MoveDatabase(sourceHostHandle string, targetHostHandle string) {

}
