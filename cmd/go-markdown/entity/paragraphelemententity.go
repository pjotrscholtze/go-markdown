package entity

type paragraphMarkdownElement struct {
	Content []MarkdownElement
}
type ParagraphMarkdownElement interface {
	GetContent() []MarkdownElement
	AsMarkdownString() string
	Kind() string
}

func (ime *paragraphMarkdownElement) GetContent() []MarkdownElement {
	return ime.Content
}

func (ime *paragraphMarkdownElement) Kind() string {
	return ElementKindParagraph
}
func (ime *paragraphMarkdownElement) AsMarkdownString() string {
	return GlueToString(ime.Content)
}
func NewParagraphMarkdownElement(input string, parserFn func(input string) []MarkdownElement) ParagraphMarkdownElement {
	return &paragraphMarkdownElement{
		Content: parserFn(input),
	}
}
