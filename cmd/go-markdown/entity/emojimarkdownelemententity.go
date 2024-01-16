package entity

type emojithroughMarkdownElement struct {
	Content string
}
type EmojiMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (ime *emojithroughMarkdownElement) Kind() string {
	return ElementKindEmoji
}
func (ime *emojithroughMarkdownElement) AsMarkdownString() string {
	return ":" + ime.Content + ":"
}
func NewEmojiMarkdownElement(input string) EmojiMarkdownElement {
	return &emojithroughMarkdownElement{
		Content: input[1 : len(input)-1],
	}
}
