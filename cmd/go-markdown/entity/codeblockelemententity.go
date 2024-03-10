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
func NewCodeBlockMarkdownElement(input string) CodeBlockMarkdownElement {
	return &codeBlockMarkdownElement{
		Content: input[3 : len(input)-3],
	}
}
