package interpreter

import (
	"os/exec"
	"strings"
)

type Localhost struct {
}

func NewLocalhost() *Localhost {
	return &Localhost{}
}

func (localhost *Localhost) Execute(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(string(output))
	return result, nil
}
