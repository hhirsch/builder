package token

const (
	//token type
	ILLEGAL             = "ILLEGAL" // encountered unexpected token
	EOF                 = "EOF"     // parsing will stop when EOF is encountered
	LINE_BREAK          = "LINE_BREAK"
	ROOT                = "ROOT"
	ASSIGN_OPERATION    = "ASSIGN_OPERATION"
	STATEMENT           = "STATEMENT"
	LITERAL             = "LITERAL" // a string of characters
	FUNCTION_IDENTIFIER = "FUNCTION_IDENTIFIER"
	FUNCTION_PARAMETERS = "FUNCTION_PARAMETERS"
	IDENTIFIER          = "IDENTIFIER" // variables, data types, labels, subroutines, and modules. (1)
	IDENTIFIER_VARIADIC = "IDENTIFIER_VARIADIC"
	EXPRESSION          = "EXPRESSION" // expressions are evaluated to determine their value for example 1+1 (2)
	SPACE               = "SPACE"
	SINGLE_LINE_COMMENT = "SINGLE_LINE_COMMENT"
	TAB_INDENT          = "TAB_INDENT"
	//	FUNCTION_NAME       = "FUNCTION_NAME"

	// operators
	ASSIGN            = "="
	COMMA             = ","
	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	MORE              = "..."
	DOLLAR            = '$'
	NEW_LINE          = '\n'
	CARRIAGE_RETURN   = '\r'
	TAB               = '\t'
	WHITE_SPACE       = ' '
	SLASH             = "/"

	// keywords
	FUNCTION = "FUNCTION"
	DONE     = "DONE"
	IF       = "IF"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

type Type string

var keywords = map[string]Type{
	"function": FUNCTION,
	"done":     DONE,
	"if":       IF,
	"true":     TRUE,
	"false":    FALSE,
	"return":   RETURN,
}

type Token struct {
	Type         Type
	Literal      string
	LineNumber   uint
	ColumnNumber uint
	FileName     string
}

func NewToken(tokenType Type, literal string) *Token {
	return &Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func GetStatement(literal string) Type {
	if keyword, exists := keywords[literal]; exists {
		return keyword
	}
	return STATEMENT
}

/*
 * References
 * (1) https://en.wikipedia.org/wiki/Identifier_(computer_languages)
 * (2) https://en.wikipedia.org/wiki/Expression_(computer_science)
 */
