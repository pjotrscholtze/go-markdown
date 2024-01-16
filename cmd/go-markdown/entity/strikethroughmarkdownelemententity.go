package entity

type strikethroughMarkdownElement struct {
	Content string
}
type StrikethroughMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (ime *strikethroughMarkdownElement) Kind() string {
	return ElementKindStrikethrough
}
func (ime *strikethroughMarkdownElement) AsMarkdownString() string {
	return "~~" + ime.Content + "~~"
}
func NewStrikethroughMarkdownElement(input string) StrikethroughMarkdownElement {
	return &strikethroughMarkdownElement{
		Content: input[2 : len(input)-2],
	}
}
