package interpreter

type System interface {
	Execute(command string) (string, error)
}
