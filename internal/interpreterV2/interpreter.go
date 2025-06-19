package interpreterV2

import (
	"fmt"
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/command"
	"github.com/hhirsch/builder/internal/token"
	"log/slog"
)

type Interpreter struct {
	commands command.Commands
}

func NewInterpreter() (*Interpreter, error) {
	interpreter := &Interpreter{commands: *command.NewCommands()}
	interpreter.commands.AddCommand(command.NewPrintCommand())
	return interpreter, nil
}

func (interpreter *Interpreter) Run(rootNode *ast.Node) error {
	slog.Debug("Interpreter is running.")
	var params []string
	for _, node := range rootNode.Children {
		slog.Debug("Processing child nodes.")
		switch node.Type {
		case token.STATEMENT:
			var error error
			var currentCommand *command.Command
			currentCommand, error = interpreter.commands.GetCommand(node.Value)
			if error != nil {
				return fmt.Errorf("Loading command: %w", error)
			}

			params = append(params, node.Value)
			for _, param := range node.Children {
				slog.Debug("Processing parameter node.")
				params = append(params, param.Value)
			}

			(*currentCommand).Execute(params)
		case token.FUNCTION:
			slog.Debug("Adding function.")
		case token.EOF:
			slog.Debug("Encountered EOF. Stopping interpreter.")
			return nil
		default:
			slog.Debug("Interpreter encountered unknown node.", slog.String("node type", string(node.Type)))
		}

	}
	return nil
}
