package entity

type emojiMarkdownElement struct {
	Content string
}
type EmojiMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (ime *emojiMarkdownElement) Kind() string {
	return ElementKindEmoji
}
func (ime *emojiMarkdownElement) AsMarkdownString() string {
	return ":" + ime.Content + ":"
}
func NewEmojiMarkdownElement(input string) EmojiMarkdownElement {
	return &emojiMarkdownElement{
		Content: input[1 : len(input)-1],
	}
}
