package commands

import (
	"fmt"
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type BaseCommand struct {
	environment *models.Environment
}

func (this *BaseCommand) TestRequirements() bool {
	return true
}

func (this *BaseCommand) TrimResponseString(string string) string {
	return strings.TrimSpace(string)
}

func (this *BaseCommand) IsTrue(string string) bool {
	return strings.TrimSpace(string) == "true"
}

func (this *BaseCommand) FindBinary(binaryName string) bool {
	var command string = fmt.Sprintf("command -v %s >/dev/null 2>&1 && echo true || echo false", binaryName)
	if this.IsTrue(this.environment.Client.Execute(command)) {
		this.environment.GetLogger().Debug(fmt.Sprintf("Binary %s found.", binaryName))
		return true
	}
	var errorMessage string = fmt.Sprintf("Unable to find required binary %s on target system.", binaryName)
	this.environment.GetLogger().Info(errorMessage)
	return false
}
