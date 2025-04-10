package interpreter

type Command interface {
	TestRequirements() bool
	Execute(tokens []string) (string, error)
	GetDescription(tokens []string) string
	GetBrief() string
	GetHelp() string
	GetName() string
	RequiresConnection() bool
}
