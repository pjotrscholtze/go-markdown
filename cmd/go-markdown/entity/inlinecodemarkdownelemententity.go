package entity

type inlineCodeMarkdownElement struct {
	Content string
}
type InlineCodeMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (icme *inlineCodeMarkdownElement) Kind() string {
	return ElementKindInlineCode
}
func (icme *inlineCodeMarkdownElement) AsMarkdownString() string {
	return "`" + icme.Content + "`"
}
func NewInlineCodeMarkdownElement(input string) InlineCodeMarkdownElement {
	return &inlineCodeMarkdownElement{
		Content: input[1 : len(input)-1],
	}
}
