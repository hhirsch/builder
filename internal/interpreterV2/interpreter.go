package interpreterV2

import (
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/command"
	"github.com/hhirsch/builder/internal/token"
	"github.com/hhirsch/builder/internal/variable"
	"log/slog"
)

type Interpreter struct {
	commands  command.Commands
	variables variable.VariablePool
	function  Function
	statement Statement
}

func NewInterpreter() (*Interpreter, error) {
	commands := *command.NewCommands()
	interpreter := &Interpreter{
		commands:  commands,
		function:  *NewFunction(&commands),
		statement: *NewStatement(&commands),
	}
	interpreter.commands.AddCommand(command.NewPrintCommand())
	return interpreter, nil
}

func (interpreter *Interpreter) convertVariablesToLiterals(node *ast.Node) *ast.Node {
	return nil
}

func (interpreter *Interpreter) Run(rootNode *ast.Node) error {
	slog.Info("Interpreter is running.")
	for _, node := range rootNode.Children {
		slog.Debug("Processing child nodes.")
		switch node.Type {
		case token.STATEMENT:
			interpreter.statement.Execute(node)
		case token.FUNCTION:
			interpreter.function.Add(node)
		case token.EOF:
			slog.Debug("Encountered EOF. Stopping interpreter.")
			return nil
		default:
			slog.Error("Interpreter encountered unknown node.", slog.String("node type", string(node.Type)))
		}

	}
	return nil
}
