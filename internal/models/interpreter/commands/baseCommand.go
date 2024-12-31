package commands

import (
	"fmt"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/models"
	"strings"
)

type BaseCommand struct {
	environment *models.Environment
	logger      *helpers.Logger
}

func NewBaseCommand(environment *models.Environment) *BaseCommand {
	return &BaseCommand{
		environment: environment,
		logger:      environment.GetLogger(),
	}
}

func (this *BaseCommand) TestRequirements() bool {
	return true
}

func (this *BaseCommand) requireParameterAmount(tokens []string, requiredParameterAmount int) {
	if len(tokens) != requiredParameterAmount+1 {
		this.logger.Fatalf("This command needs %d parameters", requiredParameterAmount)
	}
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
		this.logger.Debugf("Binary %s found.", binaryName)
		return true
	}

	this.logger.Infof("Unable to find required binary %s on target system.", binaryName)
	return false
}
