package entity

type listElementMarkdownElement struct {
	Content []MarkdownElement
}
type ListElementMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *listElementMarkdownElement) Kind() string {
	return ElementKindList
}
func (bqme *listElementMarkdownElement) AsMarkdownString() string {
	return GlueToString(bqme.Content)
}
func NewListElementMarkdownElement(input string, parserFn func(input string) []MarkdownElement) ListElementMarkdownElement {
	return &listElementMarkdownElement{
		Content: parserFn(input),
	}
}
