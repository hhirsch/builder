package command

import (
	"github.com/hhirsch/builder/internal/ast"
)

type Command interface {
	TestRequirements() bool
	Execute(parameters []*ast.Node) (string, error)
	GetName() string
	RequiresConnection() bool
}
