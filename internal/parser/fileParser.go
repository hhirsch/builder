package parser

import (
	"fmt"
	"github.com/hhirsch/builder/internal/ast"
	"github.com/hhirsch/builder/internal/lexer"
	"io"
	"log/slog"
	"os"
)

func GetSyntaxTreeFromDirectory(fileName string) (*ast.Node, error) {
	return nil, nil
}

func GetSyntaxTreeFromFile(fileName string) (*ast.Node, error) {
	file, error := os.Open(fileName)
	if error != nil {
		return nil, fmt.Errorf("File Parser: %w", error)
	}
	defer file.Close()

	data, error := io.ReadAll(file)
	if error != nil {
		return nil, error
	}
	slog.Info("File loaded")
	lexer, error := lexer.NewLexer(string(data))
	if error != nil {
		return nil, error
	}
	slog.Info("Parsing...")
	parser, error := NewParser(lexer)
	if error != nil {
		return nil, error
	}
	if len(*parser.GetErrors()) > 0 {
		fmt.Printf("Error parsing file %v:", fileName)
		for _, error := range *parser.GetErrors() {
			fmt.Println(error)
		}
		return nil, fmt.Errorf("%v parse errors.", len(*parser.GetErrors()))
	}
	slog.Info("Running...")
	return parser.GetSyntaxTree(), nil
}
