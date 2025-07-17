package command

import (
	"testing"
)

func TestBufferWriter(test *testing.T) {
	writer := NewBufferWriter()
	writer.Write("foo")
	if len(writer.GetHistory()) != 1 {
		test.Error("Expected one item to be in the history.")
	}
	if writer.GetHistory()[0] != "foo" {
		test.Error("Write did not show up in history.")
	}
}
