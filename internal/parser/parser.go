package parser

import (
	"fmt"
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/lexer"
	"github.com/hhirsch/builder/internal/token"
	"log/slog"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	errors       []string
}

func NewParser(lexer *lexer.Lexer) (*Parser, error) {
	token := lexer.NextToken()
	parser := &Parser{
		lexer:        lexer,
		currentToken: *token,
	}

	return parser, nil
}
func (parser *Parser) advance() {
	token := parser.lexer.NextToken()
	parser.currentToken = *token
}

func (parser *Parser) GetErrors() *[]string {
	return &parser.errors
}

func (parser *Parser) GetSyntaxTree() *ast.Node {
	var children []*ast.Node
	/*	if parser.currentToken.Type != token.LINE_BREAK {
		parser.advance()
	}*/
	for parser.currentToken.Type != token.EOF {
		//		if parser.currentToken.Type != token.LINE_BREAK {
		statement := parser.parse()
		children = append(children, statement)
		//}
		parser.advance()
	}
	//if len(children) == 0 {
	children = append(children, ast.NewEndOfFile())
	//}
	return ast.NewRoot(children...)
}

func (parser *Parser) parse() *ast.Node {
	switch parser.currentToken.Type {
	case token.LINE_BREAK:
		//panic("illegal line break that should have been skipped")
		//		parser.advance()
		return ast.NewLineBreak()

	case token.SPACE:
		//panic("illegal line break that should have been skipped")
		//		parser.advance()
		return ast.NewSpace()
	case token.STATEMENT:
		slog.Debug("Parsing statement.")
		return parser.parseStatement()
	case token.LITERAL:
		//panic("Literal result when only statements in the input! Literal: " + parser.currentToken.Literal)
		//return parser.parseStatement()
	case token.FUNCTION:
		slog.Debug("Parsing function.")
		return parser.parseFunction()
	case token.IDENTIFIER:
		slog.Debug("Parsing operation.")
		return parser.parseOperation()
	case token.RETURN:
		//return parser.parseReturnStatement()
	default:
		//return parser.parseStatement()
		//panic("Literal result when only statements in the input! Type: " + parser.currentToken.Type)
		return ast.NewLiteral(parser.currentToken.Literal)
		//return parser.parseExpressionStatement()
	}
	return ast.NewStatement("nil", ast.NewLiteral("nil"))
}

func (parser *Parser) parseStatement() *ast.Node {
	statementName := parser.currentToken.Literal
	parser.advance()
	var literals []*ast.Node
	for parser.currentToken.Type != token.LINE_BREAK && parser.currentToken.Type != token.EOF {
		switch parser.currentToken.Type {
		case token.LITERAL:
			literals = append(literals, ast.NewLiteral(parser.currentToken.Literal))
		case token.IDENTIFIER:
			literals = append(literals, ast.NewIdentifier(parser.currentToken.Literal))
		case token.ILLEGAL:
			parser.errors = append(parser.errors, "Illegal token type %v while parsing statement.", string(parser.currentToken.Type))
		default:
			parser.errors = append(parser.errors, "Unexpected token type %v while parsing statement.", string(parser.currentToken.Type))
		}
		parser.advance()
	}
	if parser.currentToken.Type == token.LINE_BREAK {
		literals = append(literals, ast.NewLineBreak())
	} // else if parser.currentToken.Type == token.EOF {
	//	literals = append(literals, ast.NewEndOfFile())
	//}

	//literals = append(literals, ast.NewLineBreak())
	return ast.NewStatement(statementName, literals...)
}

func (parser *Parser) parseFunction() *ast.Node {
	parser.advance()
	if parser.currentToken.Type != token.IDENTIFIER {
		parser.errors = append(parser.errors, "Expected type identifier after function keyword.")
	}

	functionName := parser.currentToken.Literal
	var parameters []*ast.Node
	var functionBody []*ast.Node
	// parameters
	for parser.currentToken.Type != token.LINE_BREAK && parser.currentToken.Type != token.EOF {
		switch parser.currentToken.Type {
		case token.IDENTIFIER:
			parameters = append(parameters, ast.NewIdentifier(parser.currentToken.Literal))
		case token.IDENTIFIER_VARIADIC:
			parameters = append(parameters, ast.NewIdentifierVariadic(parser.currentToken.Literal))
		case token.ILLEGAL:
			parser.errors = append(parser.errors, "Illegal token type %v while parsing statement.", string(parser.currentToken.Type))
		default:
			parser.errors = append(parser.errors, fmt.Sprintf("Unexpected token type %v when trying to parse function definition parameter.", parser.currentToken.Type))
		}
		parser.advance()
	}

	if parser.currentToken.Type == token.LINE_BREAK {
		parameters = append(parameters, ast.NewLineBreak())
		parser.advance()
	}
	// body
	for parser.currentToken.Type != token.DONE && parser.currentToken.Type != token.EOF {
		switch parser.currentToken.Type {
		case token.LINE_BREAK:
			functionBody = append(functionBody, ast.NewLineBreak())
		case token.SPACE:
			functionBody = append(functionBody, ast.NewSpace())
		case token.STATEMENT:
			functionBody = append(functionBody, parser.parseStatement())
		case token.DONE:
			parser.errors = append(parser.errors, "Illegal token type %v while parsing statement.", string(parser.currentToken.Type))
		case token.ILLEGAL:
			parser.errors = append(parser.errors, "Illegal token type %v while parsing statement.", string(parser.currentToken.Type))
		default:
			parser.errors = append(parser.errors, fmt.Sprintf("Unexpected token type %v when trying to parse function definition body.", parser.currentToken.Type))
		}
		parser.advance()
	}

	return ast.NewFunction(functionName, parameters, functionBody)
}

func (parser *Parser) parseBlockStatement() *ast.Node {
	return ast.NewOperation(ast.NewIdentifier("test"), token.ASSIGN, ast.NewStatement("listFiles"))
}

func (parser *Parser) parseOperation() *ast.Node {
	return ast.NewOperation(ast.NewIdentifier("test"), token.ASSIGN, ast.NewStatement("listFiles"))
}
