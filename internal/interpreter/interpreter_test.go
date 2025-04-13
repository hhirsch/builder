package interpreter

import (
	"testing"
)

func TestRunOnNonExistingFileCausesError(t *testing.T) {
	interpreter := Interpreter{}
	error := interpreter.Run("nonExistingFileName")
	if error == nil {
		t.Errorf("Expected non nil error.")
	}
	expectedErrorMessage := "open file: open nonExistingFileName: no such file or directory"

	if error.Error() != expectedErrorMessage {
		t.Errorf("Expected: %s", expectedErrorMessage)
	}
}
