package commands

type Command interface {
	TestRequirements() bool
	Execute(tokens []string)
	GetDescription(tokens []string) string
	GetHelp() string
}
