package entity

type termDefinitionElementMarkdownElement struct {
	Content []MarkdownElement
}
type TermDefinitionElementMarkdownElement interface {
	GetContent() []MarkdownElement
	AsMarkdownString() string
	Kind() string
}

func (bqme *termDefinitionElementMarkdownElement) GetContent() []MarkdownElement {
	return bqme.Content
}

func (bqme *termDefinitionElementMarkdownElement) Kind() string {
	return ElementKindTermDefinitionLine
}
func (bqme *termDefinitionElementMarkdownElement) AsMarkdownString() string {
	return "^: " + GlueToString(bqme.Content)
}
func NewTermDefinitionElementMarkdownElement(input string, parserFn func(input string) []MarkdownElement) TermDefinitionElementMarkdownElement {
	return &termDefinitionElementMarkdownElement{
		Content: parserFn(input[3:]),
	}
}
