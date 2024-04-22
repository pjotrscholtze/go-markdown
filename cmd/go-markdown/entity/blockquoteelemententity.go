package entity

type blockQuoteMarkdownElement struct {
	Content []MarkdownElement
}
type BlockQuoteMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *blockQuoteMarkdownElement) Kind() string {
	return ElementKindBlockquote
}
func (bqme *blockQuoteMarkdownElement) AsMarkdownString() string {
	return ">" + GlueToString(bqme.Content)
}
func NewBlockQuoteMarkdownElement(input string, parserFn func(input string) []MarkdownElement) BlockQuoteMarkdownElement {
	return &blockQuoteMarkdownElement{
		Content: parserFn(input[1:]),
	}
}
