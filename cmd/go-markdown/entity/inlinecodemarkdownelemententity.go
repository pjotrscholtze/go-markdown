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
func NewInlineCodeMarkdownElement(input string, parserFn func(input string) []MarkdownElement) InlineCodeMarkdownElement {
	// parserFn is not used, since in inline code no sub elements can exist.
	// However, to keep the interface similiar between all entities, this
	// parameter has been created.
	return &inlineCodeMarkdownElement{
		Content: input[1 : len(input)-1],
	}
}
