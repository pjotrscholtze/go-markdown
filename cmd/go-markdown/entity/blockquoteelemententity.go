package entity

type blockQuoteMarkdownElement struct {
	Content []MarkdownElement
}
type BlockQuoteMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	GetContent() []MarkdownElement
}

func (bqme *blockQuoteMarkdownElement) Kind() string {
	return ElementKindBlockquote
}
func (bqme *blockQuoteMarkdownElement) AsMarkdownString() string {
	return ">" + GlueToString(bqme.Content)
}
func (bme *blockQuoteMarkdownElement) GetContent() []MarkdownElement {
	return bme.Content
}
func NewBlockQuoteMarkdownElement(input string, parserFn func(input string) []MarkdownElement) BlockQuoteMarkdownElement {
	return &blockQuoteMarkdownElement{
		Content: parserFn(input[1:]),
	}
}
