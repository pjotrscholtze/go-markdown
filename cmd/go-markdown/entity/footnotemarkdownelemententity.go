package entity

type footnoteMarkdownElement struct {
	Content string
}
type FootnoteMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (ime *footnoteMarkdownElement) Kind() string {
	return ElementKindFootnote
}
func (ime *footnoteMarkdownElement) AsMarkdownString() string {
	return "[^" + ime.Content + "]"
}
func NewFootnoteMarkdownElement(input string) FootnoteMarkdownElement {
	return &footnoteMarkdownElement{
		Content: input[2 : len(input)-1],
	}
}
