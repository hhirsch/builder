package command

import (
	"testing"
)

func TestPrint(test *testing.T) {
	expectedString := "bar baz"
	resultString := NewPrintCommand().getStringFromTokens([]string{"print", "bar", "baz"})
	if resultString != expectedString {
		test.Errorf("Wrong token literal: %v. Expected: %v.", resultString, expectedString)
	}
}
