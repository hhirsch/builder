package interpreterV2

import (
	"testing"

	"github.com/hhirsch/builder/internal/command"
	"github.com/hhirsch/builder/internal/lexer"
	"github.com/hhirsch/builder/internal/parser"
)

func TestRun(test *testing.T) {
	input := `print Testing 1
print Testing 2
print Testing 3`
	lexer, _ := lexer.NewLexer(input)
	parser, _ := parser.NewParser(lexer)
	syntaxTree := parser.GetSyntaxTree()

	interpreter, error := NewInterpreter()
	if error != nil {
		test.Error(error.Error())
	}
	var writer command.Writer = command.NewBufferWriter()
	interpreter.commands.AddCommand(command.NewPrintCommand(&writer))
	error = interpreter.Run(syntaxTree)
	if error != nil {
		test.Error(error.Error())
	}
	if len(writer.GetHistory()) != 3 {
		test.Fatal("Buffer size should be exactly 3.")
	}

	if writer.GetHistory()[0] != "Testing 1" {
		test.Errorf("Unexpected string in first position: %v", writer.GetHistory()[0])
	}
	if writer.GetHistory()[1] != "Testing 2" {
		test.Errorf("Unexpected string in second position: %v", writer.GetHistory()[1])
	}
	if writer.GetHistory()[2] != "Testing 3" {
		test.Errorf("Unexpected string in third position: %v", writer.GetHistory()[2])
	}
}
