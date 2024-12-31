package commands

import (
	"github.com/hhirsch/builder/internal/models"
)

type SqlCommand struct {
	environment *models.Environment
}

func (this *SqlCommand) uploadSqlCredentials() {
}

func (this *SqlCommand) executeSqlCommand() {
}

func (this *SqlCommand) wipeSqlCredentials() {
}
