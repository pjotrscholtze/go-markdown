package entity

type emojiMarkdownElement struct {
	Content string
}
type EmojiMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	GetContent() string
	SetContent(content string)
}

func (ime *emojiMarkdownElement) GetContent() string {
	return ime.Content
}
func (ime *emojiMarkdownElement) SetContent(content string) {
	ime.Content = content
}
func (ime *emojiMarkdownElement) Kind() string {
	return ElementKindEmoji
}
func (ime *emojiMarkdownElement) AsMarkdownString() string {
	return ":" + ime.Content + ":"
}
func NewEmojiMarkdownElement(input string, parserFn func(input string) []MarkdownElement) EmojiMarkdownElement {
	// parserFn is not used, since in an emoji no sub elements can exist.
	// However, to keep the interface similiar between all entities, this
	// parameter has been created.
	return &emojiMarkdownElement{
		Content: input[1 : len(input)-1],
	}
}
