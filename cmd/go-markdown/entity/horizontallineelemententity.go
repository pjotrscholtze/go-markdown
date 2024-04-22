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
func NewHorizontalLineMarkdownElement(input string, parserFn func(input string) []MarkdownElement) HorizontalLineMarkdownElement {
	// parserFn is not used, since in an emoji no sub elements can exist.
	// However, to keep the interface similiar between all entities, this
	// parameter has been created.
	return &horizontalLineMarkdownElement{
		Content: input,
	}
}
