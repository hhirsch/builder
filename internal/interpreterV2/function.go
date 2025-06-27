package interpreterV2

import (
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/command"
	"log/slog"
)

type Function struct {
	commands *command.Commands
}

func NewFunction(commands *command.Commands) *Function {
	return &Function{commands: commands}
}

func (function *Function) Add(node *ast.Node) error {
	slog.Debug("Adding function.")
	function.commands.AddCommand(command.NewCustomCommand(node))
	return nil
}
