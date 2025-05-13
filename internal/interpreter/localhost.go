package interpreter

import (
	"fmt"
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

func (localhost *Localhost) Upload(source string, target string) error {
	_, err := localhost.Execute(fmt.Sprintf("cp %s %s", source, target))
	return err
}

func (localhost *Localhost) Download(source string, target string) error {
	_, err := localhost.Execute(fmt.Sprintf("cp %s %s", source, target))
	return err
}

func (localhost *Localhost) Delete(target string) error {
	return nil
}
