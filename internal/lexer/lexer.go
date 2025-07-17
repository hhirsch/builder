package lexer

import (
	"fmt"
	"github.com/hhirsch/builder/internal/token"
)

type Lexer struct {
	position       int
	character      byte
	input          string
	isAtFirstToken bool
	line           uint32
	column         uint16
}

func NewLexer(input string) (*Lexer, error) {
	if len(input) < 1 {
		return &Lexer{}, fmt.Errorf("Input has to be at least one character long.")
	}
	lexer := &Lexer{
		position:       0,
		character:      input[0],
		input:          input,
		isAtFirstToken: true,
		line:           1,
		column:         1,
	}
	return lexer, nil
}

func (lexer *Lexer) GetPosition() (line uint32, column uint16) {
	return lexer.line, lexer.column
}

func (lexer *Lexer) newLine() {
	lexer.line++
	lexer.column = 1
}

func (lexer *Lexer) NextToken() *token.Token {
	var nextToken *token.Token
	for lexer.character == token.WHITE_SPACE {
		lexer.nextCharacter()
		lexer.column++
		if lexer.isAtFirstToken {
			return token.NewToken(token.SPACE, string(token.WHITE_SPACE))
		}
	}

	nextToken = token.NewToken(token.EOF, "")
	// check the beginning of the line. Then handle the rest of the line with a function.
	switch lexer.character {
	case token.NEW_LINE:
		nextToken = token.NewToken(token.LINE_BREAK, string(token.NEW_LINE))
		lexer.isAtFirstToken = true
		lexer.newLine()
	case token.CARRIAGE_RETURN:
		nextToken = token.NewToken(token.LINE_BREAK, lexer.read(isNewLine))
		lexer.isAtFirstToken = true
		lexer.newLine()
	case token.TAB:
		if lexer.isAtFirstToken {
			nextToken = token.NewToken(token.TAB_INDENT, string(token.TAB))
			lexer.column++
		}
	case token.DOLLAR:
		lexer.isAtFirstToken = false
		lexer.nextCharacter()
		lexer.column++
		identifier := lexer.read(isLetter)
		lexer.nextCharacter()
		followUpDots := lexer.read(isDot)
		if len(followUpDots) == 3 {
			nextToken = token.NewToken(token.IDENTIFIER_VARIADIC, identifier)
		} else if len(followUpDots) == 0 {
			nextToken = token.NewToken(token.IDENTIFIER, identifier)
		} else {
			nextToken = token.NewToken(token.ILLEGAL, identifier+followUpDots)
		}
	case 0:
		nextToken = token.NewToken(token.EOF, "")
	default:
		if lexer.isAtFirstToken && isLetter(lexer.character) {
			literal := lexer.read(isLetter)
			nextToken = token.NewToken(token.GetStatement(literal), literal)
			lexer.isAtFirstToken = false
		} else if isLetter(lexer.character) {
			literal := lexer.read(isLetter)
			nextToken = token.NewToken(token.LITERAL, literal)
		} else if isDigit(lexer.character) {
			literal := lexer.read(isDigit)
			nextToken = token.NewToken(token.LITERAL, literal)
		} else {
			lexer.moveToEnd()
			nextToken = token.NewToken(token.ILLEGAL, string(lexer.character))
		}
	}
	lexer.nextCharacter()
	return nextToken
}

func (lexer *Lexer) read(limiterFunction func(byte) bool) string {
	startPosition := lexer.position
	for limiterFunction(lexer.character) {
		lexer.nextCharacter()
		lexer.column++
	}
	token := lexer.input[startPosition:lexer.position]
	/* We had to look ahead to find where our token ends.
	 * We go back one character so that we can handle
	 * characters between our tokens like white space and new lines.
	 */
	lexer.position--
	return token
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_' || character == '-'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func isDot(character byte) bool {
	return '.' == character
}

func isNewLine(character byte) bool {
	return '\n' == character || '\r' == character
}

func (lexer *Lexer) nextCharacter() {
	lexer.position++
	if lexer.position >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.position]
	}
}

func (lexer *Lexer) moveToEnd() {
	lexer.position = len(lexer.input) - 1
	lexer.character = 0
}
