package lexer

import (
	"github.com/hhirsch/builder/internal/token"
	"testing"
)

func TestIsLetter(test *testing.T) {
	testData := []byte{
		'A',
		'B',
		'z',
		'-',
		'_',
	}
	for _, value := range testData {
		if !isLetter(value) {
			test.Errorf("Did not recognize valid letter %v.", string(value))
		}
	}
}

func TestNextCharacterWithNewLineCharacters(test *testing.T) {

	input := `print new line test
a`
	lexer, _ := NewLexer(input)
	testData := []string{
		"print",
		"new",
		"line",
		"test",
		"\n",
	}
	for key, value := range testData {
		nextToken := lexer.NextToken()
		if nextToken.Literal != value {
			test.Errorf("Token %v is not %v, got \"%v\":\"%v\" instead.", key, value, nextToken.Type, nextToken.Literal)
		}
	}

	if lexer.position != 20 {
		test.Errorf("Unexpected position: %v.", lexer.position)
	}
	if lexer.character != 'a' {
		test.Errorf("Unexpected character: %q.", string(lexer.character))
	}
}

func TestLexerConstructorWithNewLine(test *testing.T) {
	input := "\n"
	lexer, error := NewLexer(input)

	if error != nil {
		test.Errorf("Error should be nil if lexer input is a new line.")
		return
	}
	nextToken := lexer.NextToken()
	if nextToken.Type != token.LINE_BREAK {
		test.Errorf("Token type should be line break, not %v", nextToken.Type)
	}
	if nextToken.Literal != "\n" {
		test.Errorf("Token literal should be line break, not %v", nextToken.Literal)
	}
}

func TestLexerConstructorWithEmptyStringCreatesError(test *testing.T) {
	input := ""
	_, error := NewLexer(input)
	if error == nil {
		test.Errorf("Error shouldn't be nil if lexer input is empty.")
		return
	}
	if error.Error() != "Input has to be at least one character long." {
		test.Errorf("Expected empty input to lead to an error.")
	}
}

func TestLexerConstructorWithWhiteSpaceReturnsEof(test *testing.T) {
	input := " "
	lexer, error := NewLexer(input)
	if error != nil {
		test.Errorf("Error should be nil if lexer input is a space.")
		return
	}
	lexer.nextCharacter()
	if lexer.character != 0 {
		test.Errorf("Expected character to be 0, got %v instead.", string(lexer.character))
	}
}

func TestNextCharacter(test *testing.T) {
	lexer, _ := NewLexer("test")
	if lexer.character != 't' {
		test.Errorf("Next character gave unexpected result: %v. Expected: t", string(lexer.character))
	}
}

func TestNextToken(test *testing.T) {
	input := "$variable"

	lexer, _ := NewLexer(input)
	if lexer.character != input[0] {
		test.Errorf("First character is not set as the character during lexer creation.")
	}
	lexer.nextCharacter()
	if lexer.character != input[1] {
		test.Errorf("Second character is not set as the character after nextCharacter call. Instead it is: %v", string(lexer.character))
	}
}

func TestNextTokenLinuxLineBreak(test *testing.T) {
	input := "\n"

	lexer, _ := NewLexer(input)
	nextToken := lexer.NextToken()
	if nextToken.Type != token.LINE_BREAK {
		test.Errorf("Wrong token type: %v. Expected type: LINE_BREAK.", nextToken.Type)
	}

	if nextToken.Literal != input {
		test.Errorf("Wrong literal content expected %q got: %q.", nextToken.Literal, input)
	}
}

func TestNextTokenWindowsLineBreak(test *testing.T) {
	input := "\r\n"

	lexer, _ := NewLexer(input)
	nextToken := lexer.NextToken()
	if nextToken.Type != token.LINE_BREAK {
		test.Errorf("Wrong token type: %v. Expected type: LINE_BREAK.", nextToken.Type)
	}
	if nextToken.Literal != input {
		test.Errorf("Wrong literal content expected %q got: %q.", input, nextToken.Literal)
	}
}

func TestNextTokenMacLineBreak(test *testing.T) {
	input := "\r"

	lexer, _ := NewLexer(input)
	nextToken := lexer.NextToken()
	if nextToken.Type != token.LINE_BREAK {
		test.Errorf("Wrong token type: %v. Expected type: LINE_BREAK.", nextToken.Type)
	}
	if nextToken.Literal != input {
		test.Errorf("Wrong literal content expected %q got: %q.", nextToken.Literal, input)
	}
}

func TestStatementWithIdentifier(test *testing.T) {
	input := "print $testVariable myLiteral"

	lexer, _ := NewLexer(input)
	if lexer.character == ' ' {
		test.Errorf("Lexer should have the first character set when initialized.")
	}
	if lexer.character != 'p' {
		test.Errorf("Expected the character to be initialized to p got \"%s\" instead", string(lexer.character))
	}

	if !lexer.isAtFirstToken {
		test.Errorf("Lexer is not at first token when initialized.")
	}
	nextToken := lexer.NextToken()

	if nextToken.Literal != "print" {
		test.Errorf("Wrong token literal: %v. Expected: print.", nextToken.Type)
	}
	if nextToken.Type != token.STATEMENT {
		test.Errorf("Wrong token type: %v. Expected type: STATEMENT.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.IDENTIFIER {
		test.Errorf("Wrong token type: %v. Expected type: IDENTIFIER.", nextToken.Type)
	}
	if nextToken.Literal != "testVariable" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "testvariable")
	}

	nextToken = lexer.NextToken()
	if nextToken.Type != token.LITERAL {
		test.Errorf("Wrong token type: %v. Expected type: LITERAL.", nextToken.Type)
	}
}

func TestNextTokenVariable(test *testing.T) {
	input := "$testVariable"

	lexer, _ := NewLexer(input)
	nextToken := lexer.NextToken()
	if nextToken.Type != token.IDENTIFIER {
		test.Errorf("Wrong token type: %v. Expected type: IDENTIFIER.", nextToken.Type)
	}
	if nextToken.Literal != "testVariable" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "testvariable")
	}
}

func TestMultipleStatements(test *testing.T) {
	input := `step Hello World
print Hello World
    `
	//	input := "print Hello World\nprint Hello World"

	lexer, _ := NewLexer(input)
	nextToken := lexer.NextToken()
	if nextToken.Type != token.STATEMENT {
		test.Errorf("Wrong token type: %v. Expected type: STATEMENT.", nextToken.Type)
	}
	if nextToken.Literal != "step" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "step")
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.LITERAL {
		test.Errorf("Wrong token type: %v. Expected type: LITERAL.", nextToken.Type)
	}
	if nextToken.Literal != "Hello" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "Hello")
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.LITERAL {
		test.Errorf("Wrong token type: %v. Expected type: LITERAL.", nextToken.Type)
	}
	if nextToken.Literal != "World" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "World")
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.LINE_BREAK {
		test.Errorf("Wrong token type: %v. Expected type: LINE_BREAK.", nextToken.Type)
	}
	if nextToken.Literal != "\n" {
		test.Errorf("Wrong token literal: %q. Expected: %q.", nextToken.Literal, "\n")
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.STATEMENT {
		test.Errorf("Wrong token type: %v. Expected type: STATEMENT.", nextToken.Type)
	}
	if nextToken.Literal != "print" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "print")
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.LITERAL {
		test.Errorf("Wrong token type: %v. Expected type: LITERAL.", nextToken.Type)
	}
	if nextToken.Literal != "Hello" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "Hello")
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.LITERAL {
		test.Errorf("Wrong token type: %v. Expected type: LITERAL.", nextToken.Type)
	}
	if nextToken.Literal != "World" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "World")
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.LINE_BREAK {
		test.Errorf("Wrong token type: %v. Expected type: LINE_BREAK.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.SPACE {
		test.Errorf("Wrong token type: %v. Expected type: SPACE.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.SPACE {
		test.Errorf("Wrong token type: %v. Expected type: SPACE.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.SPACE {
		test.Errorf("Wrong token type: %v. Expected type: SPACE.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.SPACE {
		test.Errorf("Wrong token type: %v. Expected type: SPACE.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.EOF {
		test.Errorf("Wrong token type: %v. Expected type: EOF.", nextToken.Type)
	}
}
func TestMultipleStatementsWithVariables(test *testing.T) {
	input := `
step Test
print This is a test
    `
	lexer, _ := NewLexer(input)
	nextToken := lexer.NextToken()
	if nextToken.Type != token.LINE_BREAK {
		test.Errorf("Wrong token type: %v. Expected type: LINE_BREAK.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.STATEMENT {
		test.Errorf("Wrong token type: %v. Expected type: STATEMENT.", nextToken.Type)
	}
	if nextToken.Literal != "step" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "step")
	}
}

func TestMultipleStatementsWithWhiteSpace(test *testing.T) {
	input := `
    step Test
    print This is a test
    `
	lexer, _ := NewLexer(input)
	nextToken := lexer.NextToken()

	if nextToken.Type != token.LINE_BREAK {
		test.Errorf("Wrong token type: %v. Expected type: LINE_BREAK.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	nextToken = lexer.NextToken()
	nextToken = lexer.NextToken()
	nextToken = lexer.NextToken()
	nextToken = lexer.NextToken()
	if nextToken.Type != token.STATEMENT {
		test.Errorf("Wrong token type: %v. Expected type: STATEMENT.", nextToken.Type)
	}
	if nextToken.Literal != "step" {
		test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, "step")
	}
}

func TestFunctionDefinition(test *testing.T) {
	reference := []token.Token{
		token.Token{Type: token.FUNCTION, Literal: "function"},
		token.Token{Type: token.LITERAL, Literal: "printString"},
		token.Token{Type: token.IDENTIFIER_VARIADIC, Literal: "string"},
		token.Token{Type: token.LINE_BREAK, Literal: "\n"},
		token.Token{Type: token.STATEMENT, Literal: "print"},
		token.Token{Type: token.IDENTIFIER, Literal: "string"},
		token.Token{Type: token.LINE_BREAK, Literal: "\n"},
		token.Token{Type: token.DONE, Literal: "done"},
	}

	input := `function printString $string...
print $string
done`
	lexer, _ := NewLexer(input)

	for _, referenceToken := range reference {
		nextToken := lexer.NextToken()
		if referenceToken.Type != nextToken.Type {
			test.Errorf("Wrong token type: %v. Expected type: %v. Literal: %v", nextToken.Type, referenceToken.Type, nextToken.Literal)
		}
		if referenceToken.Literal != nextToken.Literal {
			test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, referenceToken.Literal)
		}
	}
}

func TestReturnStatement(test *testing.T) {
	reference := []token.Token{
		token.Token{Type: token.FUNCTION, Literal: "function"},
		token.Token{Type: token.LITERAL, Literal: "getString"},
		token.Token{Type: token.IDENTIFIER_VARIADIC, Literal: "string"},
		token.Token{Type: token.LINE_BREAK, Literal: "\n"},
		token.Token{Type: token.RETURN, Literal: "return"},
		token.Token{Type: token.IDENTIFIER, Literal: "string"},
		token.Token{Type: token.LINE_BREAK, Literal: "\n"},
		token.Token{Type: token.DONE, Literal: "done"},
	}

	input := `function getString $string...
return $string
done`
	lexer, _ := NewLexer(input)

	for _, referenceToken := range reference {
		nextToken := lexer.NextToken()
		if referenceToken.Type != nextToken.Type {
			test.Errorf("Wrong token type: %v. Expected type: %v. Literal: %v", nextToken.Type, referenceToken.Type, nextToken.Literal)
		}
		if referenceToken.Literal != nextToken.Literal {
			test.Errorf("Wrong token literal: %v. Expected: %v.", nextToken.Literal, referenceToken.Literal)
		}
	}
}

func TestInvalidInputCreatesOneIllegalTokenAndThenEof(test *testing.T) {
	input := `..uiaeiautrdn iuaetrdiuae`
	lexer, _ := NewLexer(input)
	nextToken := lexer.NextToken()
	if nextToken.Type != token.ILLEGAL {
		test.Errorf("Wrong token type: %v. Expected Illegal.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.EOF {
		test.Errorf("Wrong token type: %v. Expected EOF.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.EOF {
		test.Errorf("Wrong token type: %v. Expected EOF.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.EOF {
		test.Errorf("Wrong token type: %v. Expected EOF.", nextToken.Type)
	}
	nextToken = lexer.NextToken()
	if nextToken.Type != token.EOF {
		test.Errorf("Wrong token type: %v. Expected EOF.", nextToken.Type)
	}
}
