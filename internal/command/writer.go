package command

type Writer interface {
	Write(message string) error
	GetHistory() []string
}
