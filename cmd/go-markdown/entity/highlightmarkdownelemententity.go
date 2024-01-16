package entity

type highlightMarkdownElement struct {
	Content string
}
type HighlightMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (icme *highlightMarkdownElement) Kind() string {
	return ElementKindHighlight
}
func (icme *highlightMarkdownElement) AsMarkdownString() string {
	return "==" + icme.Content + "=="
}
func NewHighlightMarkdownElement(input string) HighlightMarkdownElement {
	return &highlightMarkdownElement{
		Content: input[2 : len(input)-2],
	}
}
