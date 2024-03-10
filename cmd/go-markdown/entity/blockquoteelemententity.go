package entity

type blockQuoteMarkdownElement struct {
	Content string
}
type BlockQuoteMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *blockQuoteMarkdownElement) Kind() string {
	return ElementKindBlockquote
}
func (bqme *blockQuoteMarkdownElement) AsMarkdownString() string {
	return ">" + bqme.Content
}
func NewBlockQuoteMarkdownElement(input string) BlockQuoteMarkdownElement {
	return &blockQuoteMarkdownElement{
		Content: input[1:],
	}
}
