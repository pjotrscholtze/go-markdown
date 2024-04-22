package entity

type footnoteMarkdownElement struct {
	Content []MarkdownElement
}
type FootnoteMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (ime *footnoteMarkdownElement) Kind() string {
	return ElementKindFootnote
}
func (ime *footnoteMarkdownElement) AsMarkdownString() string {
	return "[^" + GlueToString(ime.Content) + "]"
}
func NewFootnoteMarkdownElement(input string, parserFn func(input string) []MarkdownElement) FootnoteMarkdownElement {
	return &footnoteMarkdownElement{
		Content: parserFn(input[2 : len(input)-1]),
	}
}
