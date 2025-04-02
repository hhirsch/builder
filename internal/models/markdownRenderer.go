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

func (markdownRenderer *MarkdownRenderer) EnableColor() {
	markdownRenderer.colour = true
}

func (markdownRenderer *MarkdownRenderer) Render(markdown string) {
	var renderedMarkdown string
	if markdownRenderer.colour {
		renderedMarkdown, _ = glamour.Render(markdown, "dark")
	} else {
		renderedMarkdown, _ = glamour.Render(markdown, "notty")
	}

	fmt.Print(renderedMarkdown)
}
