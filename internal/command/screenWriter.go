package command

import (
	"fmt"
)

type ScreenWriter struct {
}

func NewScreenWriter() *ScreenWriter {
	return &ScreenWriter{}
}

func (screenWriter *ScreenWriter) Write(message string) error {
	fmt.Println(message)
	return nil
}

func (screenWriter *ScreenWriter) GetHistory() []string {
	return []string{}
}
