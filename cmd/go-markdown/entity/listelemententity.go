// listelemententity.go
package entity

type listElementMarkdownElement struct {
	Content []ListItemElementMarkdownElement
}
type ListElementMarkdownElement interface {
	AsMarkdownString() string
	AppendListItem(listItem ListItemElementMarkdownElement)
	Kind() string
	GetSymbol() *string
	ItemCount() int
	GetContent() []ListItemElementMarkdownElement
	SymbolLength() int
}

func (bqme *listElementMarkdownElement) SymbolLength() int {
	if len(bqme.Content) == 0 {
		return 0
	}
	symbol := bqme.Content[0].SymbolLength()
	return symbol
}
func (bqme *listElementMarkdownElement) GetContent() []ListItemElementMarkdownElement {
	return bqme.Content
}

func (bqme *listElementMarkdownElement) ItemCount() int {
	return len(bqme.Content)
}

func (bqme *listElementMarkdownElement) AppendListItem(listItem ListItemElementMarkdownElement) {
	bqme.Content = append(bqme.Content, listItem)
}
func (bqme *listElementMarkdownElement) GetSymbol() *string {
	if len(bqme.Content) == 0 {
		return nil
	}
	symbol := bqme.Content[0].GetSymbol()
	return &symbol
}
func (bqme *listElementMarkdownElement) Kind() string {
	return ElementKindList
}
func (bqme *listElementMarkdownElement) AsMarkdownString() string {
	content := make([]MarkdownElement, len(bqme.Content))
	for i, item := range bqme.Content {
		content[i] = item
	}
	return GlueToString(content)
}
func NewListElementMarkdownElement() ListElementMarkdownElement {
	return &listElementMarkdownElement{
		Content: []ListItemElementMarkdownElement{},
	}
}
