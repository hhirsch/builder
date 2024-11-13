package models

import (
	"fmt"
)

type MySQL struct {
	data     map[string]string
	fileName string
}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (this *MySQL) listDatabases() {

}

func (this *MySQL) dumpDatabase(databaseName string, fileName string) string {
	//return "mysqldump -u username -p database_name > database_dump.sql"
	//return fmt.Sprintf("mysqldump -u %s -p%s %s > %s", username, password, databaseName, outputFile)
	//mysqldump --socket=/tmp/mysqlsecond.sock --all-databases > $sqlfile
	//mysqldump --socket=/var/run/mysqld/mysqld.sock database_name > dump_file.sql
	return fmt.Sprintf("mysqldump --socket=/var/run/mysqld/mysqld.sock %s > %s", databaseName, fileName)
}

func (this *MySQL) loadDatabase(databaseName string, fileName string) {

}

func (this *MySQL) moveDatabase(sourceHostHandle string, targetHostHandle string) {

}
