package command

type Command interface {
	TestRequirements() bool
	Execute(tokens []string) (string, error)
	GetName() string
	RequiresConnection() bool
}
