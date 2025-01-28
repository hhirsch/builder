package commands

type Command interface {
	TestRequirements() bool
	Execute(tokens []string) string
	GetDescription(tokens []string) string
	GetHelp() string
}
