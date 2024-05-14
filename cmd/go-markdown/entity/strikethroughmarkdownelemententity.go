package entity

type strikethroughMarkdownElement struct {
	Content []MarkdownElement
}
type StrikethroughMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	GetContent() []MarkdownElement
}

func (ime *strikethroughMarkdownElement) Kind() string {
	return ElementKindStrikethrough
}
func (ime *strikethroughMarkdownElement) AsMarkdownString() string {
	return "~~" + GlueToString(ime.Content) + "~~"
}
func (bme *strikethroughMarkdownElement) GetContent() []MarkdownElement {
	return bme.Content
}
func NewStrikethroughMarkdownElement(input string, parserFn func(input string) []MarkdownElement) StrikethroughMarkdownElement {
	return &strikethroughMarkdownElement{
		Content: parserFn(input[2 : len(input)-2]),
	}
}
