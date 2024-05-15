package entity

import "strings"

type checkboxMarkdownElement struct {
	CheckContent string
	Content      []MarkdownElement
}
type CheckboxMarkdownElement interface {
	GetCheckContent() string
	AsMarkdownString() string
	GetContent() []MarkdownElement
	Kind() string
}

func (ime *checkboxMarkdownElement) GetCheckContent() string {
	return ime.CheckContent
}

func (ime *checkboxMarkdownElement) GetContent() []MarkdownElement {
	return ime.Content
}

func (ime *checkboxMarkdownElement) Kind() string {
	return ElementKindCheckbox
}
func (ime *checkboxMarkdownElement) AsMarkdownString() string {
	return "[" + ime.CheckContent + "]" + GlueToString(ime.Content)
}
func NewCheckboxMarkdownElement(input string, parserFn func(input string) []MarkdownElement) CheckboxMarkdownElement {
	checkContent := ""
	content := input[2:]
	if idx := strings.Index(input, "]"); idx != -1 {
		checkContent = input[1:idx]
		content = input[idx+1:]
	}
	return &checkboxMarkdownElement{
		CheckContent: checkContent,
		Content:      parserFn(content),
	}
}
