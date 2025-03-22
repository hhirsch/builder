package commands

type Command interface {
	TestRequirements() bool
	Execute(tokens []string) string
	GetDescription(tokens []string) string
	GetBrief() string
	GetHelp() string
	GetName() string
	RequiresConnection() bool
}
