package commands

type Command interface {
	Execute(tokens []string)
	GetDescription(tokens []string) string
	GetHelp() string
}
