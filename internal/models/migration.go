package models

import ()

type Migration struct {
	data     map[string]string
	fileName string
}

func NewMigration(fileName string) *Migration {
	return &Migration{}
}
