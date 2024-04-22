package entity

import "strings"

type checkboxMarkdownElement struct {
	CheckContent string
	Content      []MarkdownElement
}
type CheckboxMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
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
