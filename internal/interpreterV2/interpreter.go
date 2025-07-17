package interpreterV2

import (
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/command"
	"github.com/hhirsch/builder/internal/token"
	"github.com/hhirsch/builder/internal/variable"
	"log/slog"
)

type Interpreter struct {
	commands        command.Commands
	variables       variable.VariablePool
	function        Function
	statement       Statement
	returnStatement ReturnStatement
}

func NewInterpreter() (*Interpreter, error) {
	commands := *command.NewCommands()
	interpreter := &Interpreter{
		commands:        commands,
		function:        *NewFunction(&commands),
		statement:       *NewStatement(&commands),
		returnStatement: *NewReturnStatement(),
	}
	var writer command.Writer = command.NewScreenWriter()
	printCommand := command.NewPrintCommand(&writer)
	interpreter.commands.AddCommand(printCommand)
	stepCommand := command.NewStepCommand(printCommand)
	interpreter.commands.AddCommand(stepCommand)
	return interpreter, nil
}

func (interpreter *Interpreter) Run(rootNode *ast.Node) error {
	return interpreter.Eval(rootNode, *variable.NewVariablePool())
}

func (interpreter *Interpreter) Eval(rootNode *ast.Node, variables variable.VariablePool) error {
	slog.Info("Interpreter is running.")
	for _, node := range rootNode.Children {
		slog.Debug("Processing child nodes.")
		switch node.Type {
		case token.STATEMENT:
			resolvedStatement, error := interpreter.variables.ResolveVariablesInStatement(node)
			if error != nil {
				slog.Error("Error resolving variables.")
				return error
			}
			_, error = interpreter.statement.Execute(resolvedStatement)
			if error != nil {
				slog.Error("Error executing statement.")
				return error
			}
		case token.FUNCTION:
			interpreter.function.Add(node)
		case token.RETURN:
			resolvedStatement, error := interpreter.variables.ResolveVariablesInStatement(node)
			if error != nil {
				slog.Error("Error resolving variables.")
				return error
			}
			interpreter.returnStatement.Execute(resolvedStatement)
		case token.EOF:
			slog.Debug("Encountered EOF. Stopping interpreter.")
			return nil
		default:
			slog.Error("Interpreter encountered unknown node.", slog.String("node type", string(node.Type)))
		}
	}
	return nil
}
