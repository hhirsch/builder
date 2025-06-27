package command

import "strings"

type Variable struct {
	stringContent string
}

func NewVariable(stringContent string) *Variable {
	variable := &Variable{
		stringContent: stringContent,
	}
	return variable
}

func (variable *Variable) GetSlice() ([]string, error) {
	return strings.Split(variable.stringContent, "\n"), nil
}

func (variable *Variable) SetString(content string) {
	variable.stringContent = content
}

func (variable *Variable) GetString() string {
	return variable.stringContent
}

func (variable *Variable) GetStringOrError() (string, error) {
	return variable.stringContent, nil
}

func (variable *Variable) GetFlatString() (string, error) {
	return strings.ReplaceAll(variable.stringContent, "\n", " "), nil
}

func (variable *Variable) GetCommaSeparatedString() (string, error) {
	return strings.ReplaceAll(variable.stringContent, "\n", ","), nil
}

func (variable *Variable) GetSeparatedString(seperator string) (string, error) {
	return strings.ReplaceAll(variable.stringContent, "\n", "seperator"), nil
}
