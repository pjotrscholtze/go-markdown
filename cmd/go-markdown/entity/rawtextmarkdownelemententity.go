package entity

type RawTextMarkdownElement string

func (rtme *RawTextMarkdownElement) Kind() string {
	return ElementKindText
}
func (rtme *RawTextMarkdownElement) AsMarkdownString() string {
	return string(*rtme)
}

func NewRawTextMarkdownElement(input string, parserFn func(input string) []MarkdownElement) RawTextMarkdownElement {
	// parserFn is not used, since in a raw text element no sub elements can
	// exist. However, to keep the interface similiar between all entities, this
	// parameter has been created.
	return RawTextMarkdownElement(input)
}
