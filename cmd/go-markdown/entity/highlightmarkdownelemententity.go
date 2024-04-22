package entity

type highlightMarkdownElement struct {
	Content []MarkdownElement
}
type HighlightMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (icme *highlightMarkdownElement) Kind() string {
	return ElementKindHighlight
}
func (icme *highlightMarkdownElement) AsMarkdownString() string {
	return "==" + GlueToString(icme.Content) + "=="
}
func NewHighlightMarkdownElement(input string, parserFn func(input string) []MarkdownElement) HighlightMarkdownElement {
	return &highlightMarkdownElement{
		Content: parserFn(input[2 : len(input)-2]),
	}
}
