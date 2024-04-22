package entity

type strikethroughMarkdownElement struct {
	Content []MarkdownElement
}
type StrikethroughMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (ime *strikethroughMarkdownElement) Kind() string {
	return ElementKindStrikethrough
}
func (ime *strikethroughMarkdownElement) AsMarkdownString() string {
	return "~~" + GlueToString(ime.Content) + "~~"
}
func NewStrikethroughMarkdownElement(input string, parserFn func(input string) []MarkdownElement) StrikethroughMarkdownElement {
	return &strikethroughMarkdownElement{
		Content: parserFn(input[2 : len(input)-2]),
	}
}
