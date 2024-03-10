package entity

type listElementMarkdownElement struct {
	Content string
}
type ListElementMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *listElementMarkdownElement) Kind() string {
	return ElementKindList
}
func (bqme *listElementMarkdownElement) AsMarkdownString() string {
	return bqme.Content
}
func NewListElementMarkdownElement(input string) ListElementMarkdownElement {
	return &listElementMarkdownElement{
		Content: input,
	}
}
