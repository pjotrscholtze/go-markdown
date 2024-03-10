package entity

type termDefinitionElementMarkdownElement struct {
	Content string
}
type TermDefinitionElementMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *termDefinitionElementMarkdownElement) Kind() string {
	return ElementKindTermDefinitionLine
}
func (bqme *termDefinitionElementMarkdownElement) AsMarkdownString() string {
	return "^: " + bqme.Content
}
func NewTermDefinitionElementMarkdownElement(input string) TermDefinitionElementMarkdownElement {
	return &termDefinitionElementMarkdownElement{
		Content: input[3:],
	}
}
