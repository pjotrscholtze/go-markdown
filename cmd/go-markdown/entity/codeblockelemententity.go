package entity

type codeBlockMarkdownElement struct {
	Content string
}
type CodeBlockMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *codeBlockMarkdownElement) Kind() string {
	return ElementKindCodeblock
}
func (bqme *codeBlockMarkdownElement) AsMarkdownString() string {
	return "```" + bqme.Content + "```"
}
func NewCodeBlockMarkdownElement(input string, parserFn func(input string) []MarkdownElement) CodeBlockMarkdownElement {
	// parserFn is not used, since in an code box no sub elements can exist.
	// However, to keep the interface similiar between all entities, this
	// parameter has been created.
	return &codeBlockMarkdownElement{
		Content: input[3 : len(input)-3],
	}
}
