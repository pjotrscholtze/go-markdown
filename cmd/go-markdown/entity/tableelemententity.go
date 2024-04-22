package entity

type tableElementMarkdownElement struct {
	Content string
}
type TableElementMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *tableElementMarkdownElement) Kind() string {
	return ElementKindTable
}
func (bqme *tableElementMarkdownElement) AsMarkdownString() string {
	return bqme.Content
}
func NewTableElementMarkdownElement(input string, parserFn func(input string) []MarkdownElement) TableElementMarkdownElement {
	// @todo properly parse the table here, then parseFn can be used...
	return &tableElementMarkdownElement{
		Content: input,
	}
}
