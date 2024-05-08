// listelemententity.go
package entity

type ListTuple struct {
	ListItem ListItemElementMarkdownElement
	List     ListElementMarkdownElement
}

func (lt *ListTuple) GetWhitespaceInFront() string {
	if lt.List != nil {
		return lt.List.GetWhitespaceInFront()
	}
	return lt.ListItem.GetWhitespaceInFront()
}

type listElementMarkdownElement struct {
	Content []ListTuple
	// Content []ListItemElementMarkdownElement
}
type ListElementMarkdownElement interface {
	AsMarkdownString() string
	AppendListItem(listItem ListItemElementMarkdownElement)
	AppendList(list ListElementMarkdownElement)
	Kind() string
	GetSymbol() *string
	ItemCount() int
	GetContent() []ListTuple
	SymbolLength() int
	GetWhitespaceInFront() string
}

func (bqme *listElementMarkdownElement) GetWhitespaceInFront() string {
	if len(bqme.Content) == 0 {
		return ""
	}
	if bqme.Content[0].List != nil {
		return bqme.Content[0].GetWhitespaceInFront()
	}
	return bqme.Content[0].GetWhitespaceInFront()
}

func (bqme *listElementMarkdownElement) SymbolLength() int {
	if len(bqme.Content) == 0 {
		return 0
	}
	if bqme.Content[0].List != nil {
		return bqme.Content[0].List.SymbolLength()
	}
	return bqme.Content[0].ListItem.SymbolLength()
}
func (bqme *listElementMarkdownElement) GetContent() []ListTuple {
	return bqme.Content
}

func (bqme *listElementMarkdownElement) ItemCount() int {
	return len(bqme.Content)
}

func (bqme *listElementMarkdownElement) AppendList(list ListElementMarkdownElement) {
	bqme.Content = append(bqme.Content, ListTuple{
		List: list,
	})
}

func (bqme *listElementMarkdownElement) AppendListItem(listItem ListItemElementMarkdownElement) {
	bqme.Content = append(bqme.Content, ListTuple{
		ListItem: listItem,
	})
}
func (bqme *listElementMarkdownElement) GetSymbol() *string {
	if len(bqme.Content) == 0 {
		return nil
	}
	if bqme.Content[0].List != nil {
		return bqme.Content[0].List.GetSymbol()
	}
	symbol := bqme.Content[0].ListItem.GetSymbol()
	return &symbol
}
func (bqme *listElementMarkdownElement) Kind() string {
	return ElementKindList
}
func (bqme *listElementMarkdownElement) AsMarkdownString() string {
	content := make([]MarkdownElement, len(bqme.Content))
	for i, item := range bqme.Content {
		if item.List != nil {
			content[i] = item.List
		} else {
			content[i] = item.ListItem
		}
	}
	return GlueToString(content)
}
func NewListElementMarkdownElement() ListElementMarkdownElement {
	return &listElementMarkdownElement{
		Content: []ListTuple{},
	}
}
