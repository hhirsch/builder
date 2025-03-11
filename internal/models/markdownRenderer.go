package models

import (
	_ "embed"
	"fmt"
	"github.com/charmbracelet/glamour"
)

type MarkdownRenderer struct {
	colour bool
}

func NewMarkdownRenderer() *MarkdownRenderer {
	return &MarkdownRenderer{
		colour: false,
	}

}

func (this *MarkdownRenderer) EnableColor() {
	this.colour = true
}

func (this *MarkdownRenderer) Render(markdown string) {
	var renderedMarkdown string
	if this.colour {
		renderedMarkdown, _ = glamour.Render(markdown, "dark")
	} else {
		renderedMarkdown, _ = glamour.Render(markdown, "notty")
	}

	fmt.Print(renderedMarkdown)
}
