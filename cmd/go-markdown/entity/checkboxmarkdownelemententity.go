package entity

import "strings"

type checkboxMarkdownElement struct {
	CheckContent string
	Content      string
}
type CheckboxMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (ime *checkboxMarkdownElement) Kind() string {
	return ElementKindCheckbox
}
func (ime *checkboxMarkdownElement) AsMarkdownString() string {
	return "[" + ime.CheckContent + "]" + ime.Content
}
func NewCheckboxMarkdownElement(input string) CheckboxMarkdownElement {
	checkContent := ""
	content := input[2:]
	if idx := strings.Index(input, "]"); idx != -1 {
		checkContent = input[1:idx]
		content = input[idx+1:]
	}
	return &checkboxMarkdownElement{
		CheckContent: checkContent,
		Content:      content,
	}
}
