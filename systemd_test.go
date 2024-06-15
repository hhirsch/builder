package main

import "testing"

func TestConfigOutput(test *testing.T) {
	exampleConfig := `[Unit]
Description=Test description
After=network.target

[Service]
User=User
ExecStart=/usr/bin/foo
Restart=always

[Install]
WantedBy=default.target
`
	systemd := Systemd{}
	
	result := systemd.GetConfig("User", "/usr/bin/foo", "Test description")
	if exampleConfig != result {
		test.Errorf("Config does not match example %s", result)
	}
}
