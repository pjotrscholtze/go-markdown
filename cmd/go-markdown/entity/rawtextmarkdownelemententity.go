package entity

type RawTextMarkdownElement string

func (rtme *RawTextMarkdownElement) Kind() string {
	return ElementKindText
}
func (rtme *RawTextMarkdownElement) AsMarkdownString() string {
	return string(*rtme)
}

func NewRawTextMarkdownElement(input string) RawTextMarkdownElement {
	return RawTextMarkdownElement(input)
}
