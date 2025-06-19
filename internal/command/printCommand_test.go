package command

import (
	"testing"
)

func TestPrint(test *testing.T) {
	expectedString := "foo bar baz"
	resultString := NewPrintCommand().getStringFromTokens([]string{"foo", "bar", "baz"})
	if resultString != expectedString {
		test.Errorf("Wrong token literal: %v. Expected: %v.", resultString, expectedString)
	}
}
