package interpreter

type System interface {
	Execute(command string) (string, error)
	Upload(source string, target string) error
	Download(source string, target string) error
	Delete(target string) error
}
