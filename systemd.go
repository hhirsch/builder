package main

import (
	"github.com/valyala/fasttemplate"
)

type Systemd struct {
}

func (this *Systemd) GetConfig(userName string, path string, description string) string {
	serviceConfig := `[Unit]
Description={{description}}
After=network.target

[Service]
User={{userName}}
ExecStart={{path}}
Restart=always

[Install]
WantedBy=default.target
`
	template := fasttemplate.New(serviceConfig, "{{", "}}")
	fileContent := template.ExecuteString(map[string]interface{}{
		"description":  description,
		"path": path,
		"userName": userName,
	})
	
	return fileContent
}
