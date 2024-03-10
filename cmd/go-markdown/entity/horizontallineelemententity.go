package entity

type horizontalLineMarkdownElement struct {
	Content string
}
type HorizontalLineMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *horizontalLineMarkdownElement) Kind() string {
	return ElementKindHorizontalLine
}
func (bqme *horizontalLineMarkdownElement) AsMarkdownString() string {
	return bqme.Content
}
func NewHorizontalLineMarkdownElement(input string) HorizontalLineMarkdownElement {
	return &horizontalLineMarkdownElement{
		Content: input,
	}
}
