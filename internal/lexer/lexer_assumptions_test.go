package lexer

import (
	"testing"
)

func TestNewLineWithStringComparison(test *testing.T) {
	input := `
      print This is a $test
    `
	newLineByte := '\n'
	if (input[0]) != byte(newLineByte) {
		test.Errorf("Failed to recognize new line!")
	}
}

func TestNonPrintableCharactersContributeToStringLength(test *testing.T) {
	input := `
123

`
	if len(input) != 6 {
		test.Errorf("Did not count every non printable character resulted in a count of %v instead of 5.", len(input))
	}
}
