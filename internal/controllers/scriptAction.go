package controllers

import (
	_ "embed"
	"fmt"
	"github.com/hhirsch/builder/internal/environment"
	"github.com/hhirsch/builder/internal/interpreterV2"
	"github.com/hhirsch/builder/internal/lexer"
	"github.com/hhirsch/builder/internal/models"
	"github.com/hhirsch/builder/internal/parser"
	"io"
	"log/slog"
	"os"
)

type ScriptAction struct {
	model      *models.BuilderModel
	controller *Controller
	fileName   string
	BaseAction
}

//go:embed scriptAction.md
var scriptMarkdown string

func NewScriptAction(controller *Controller, fileName string) *ScriptAction {

	return &ScriptAction{
		BaseAction: BaseAction{
			controller:  controller,
			name:        "script",
			description: "Run the script in <scriptpath>.",
			brief:       "Run a specific script.",
			help:        scriptMarkdown,
		},
		model:    models.NewBuilderModel(environment.GetProjectPath()),
		fileName: fileName,
	}
}

func (scriptAction *ScriptAction) Execute() (string, error) {
	error := scriptAction.RequireAmountOfParameters(1)
	if error != nil {
		return "", error
	}
	buffer := "Builder started\n"
	slog.Info("Builder started")
	interpreterObject, error := interpreterV2.NewInterpreter()
	if error != nil {
		return "", fmt.Errorf("interpreter: %w", error)
	}
	interpreter := *interpreterObject
	file, error := os.Open(scriptAction.fileName)
	if error != nil {
		return "", fmt.Errorf("interpreter: %w", error)
	}
	defer file.Close()

	data, error := io.ReadAll(file)
	if error != nil {
		return "", error
	}
	slog.Info("File loaded")
	lexer, error := lexer.NewLexer(string(data))
	if error != nil {
		return "", error
	}
	slog.Info("Parsing...")
	parser, error := parser.NewParser(lexer)
	if error != nil {
		return "", error
	}
	slog.Info("Running...")
	interpreter.Run(parser.GetSyntaxTree())
	if len(*parser.GetErrors()) > 0 {
		fmt.Println("Error parsing file: ")
	}
	for _, error := range *parser.GetErrors() {
		fmt.Println(error)
	}
	/*	if error != nil {
		return "", fmt.Errorf("interpreter run: %w", err)
	}*/
	slog.Info("Interpreter has finnished.")
	return buffer, nil
}

func (scriptAction *ScriptAction) GetDescription() string {
	return "Run the script in <scriptpath>."
}
