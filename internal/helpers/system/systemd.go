package system

import (
	_ "embed"
	"github.com/valyala/fasttemplate"
)

//go:embed template.txt
var serviceConfig string

type Systemd struct {
}

func (this *Systemd) GetConfig(userName string, path string, description string) string {
	template := fasttemplate.New(serviceConfig, "{{", "}}")
	fileContent := template.ExecuteString(map[string]interface{}{
		"description": description,
		"path":        path,
		"userName":    userName,
	})

	return fileContent
}
