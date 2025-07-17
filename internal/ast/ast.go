package ast

import (
	"github.com/hhirsch/builder/internal/token"
)

func NewRoot(nodes ...*Node) *Node {
	return &Node{
		Type:     token.ROOT,
		Value:    "root",
		Children: nodes,
	}
}

type Node struct {
	Type       token.Type `json:",omitempty"`
	Value      string     `json:",omitempty"`
	LeftNode   *Node      `json:",omitempty"`
	RightNode  *Node      `json:",omitempty"`
	Parameters []*Node    `json:",omitempty"`
	Children   []*Node    `json:",omitempty"`
}

func NewIllegal(value string, nodes ...*Node) *Node {
	return &Node{
		Type:     token.ILLEGAL,
		Value:    value,
		Children: nodes,
	}
}

// a statement is some action to be carried out e.g. assignment, call, return, exit (1)
func NewStatement(value string, nodes ...*Node) *Node {
	return &Node{
		Type:     token.STATEMENT,
		Value:    value,
		Children: nodes,
	}
}

// assign operation has the form of `$<identifier> = <statement>`
func NewOperation(identifier *Node, operator string, node *Node) *Node {
	return &Node{
		Type:      token.ASSIGN_OPERATION,
		Value:     operator,
		RightNode: node,
		LeftNode:  identifier,
	}
}

// an expression is a formula that can be evaluated to determine its value (2)
func NewExpression(value string) *Node {
	return &Node{
		Type:  token.EXPRESSION,
		Value: value,
	}
}

func NewLiteral(value string) *Node {
	return &Node{
		Type:  token.LITERAL,
		Value: value,
	}
}

// identifier is a variable name
func NewIdentifier(value string) *Node {
	return &Node{
		Type:  token.IDENTIFIER,
		Value: value,
	}
}

func NewReturn(node *Node) *Node {
	return &Node{
		Type:      token.RETURN,
		Value:     "return",
		RightNode: node,
	}
}

func NewIdentifierVariadic(value string) *Node {
	return &Node{
		Type:  token.IDENTIFIER_VARIADIC,
		Value: value,
	}
}

func NewFunction(value string, parameters []*Node, body []*Node) *Node {
	return &Node{
		Type:       token.FUNCTION,
		Value:      value,
		Parameters: parameters,
		Children:   body,
	}
}

func NewEndOfFile() *Node {
	return &Node{
		Type:  token.EOF,
		Value: "",
	}
}

func NewLineBreak() *Node {
	return &Node{
		Type:  token.LINE_BREAK,
		Value: "\n",
	}
}

func NewSpace() *Node {
	return &Node{
		Type:  token.SPACE,
		Value: " ",
	}
}

func NewNode(token token.Type, value string, child *Node) *Node {
	return &Node{
		Type:      token,
		Value:     value,
		RightNode: child,
	}
}

/*
 * References
 * (1) https://en.wikipedia.org/wiki/Statement_(computer_science)
 * (2) https://en.wikipedia.org/wiki/Expression_(computer_science)
 */
