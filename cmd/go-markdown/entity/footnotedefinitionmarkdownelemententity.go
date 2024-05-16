package entity

type footnoteDefinitionMarkdownElement struct {
	Content    []MarkdownElement
	Definition []MarkdownElement
}
type FootnoteDefinitionMarkdownElement interface {
	GetContent() []MarkdownElement
	GetDefinition() []MarkdownElement
	SetDefinition(definition []MarkdownElement)
	AsMarkdownString() string
	Kind() string
}

func (ime *footnoteDefinitionMarkdownElement) GetDefinition() []MarkdownElement {
	return ime.Definition
}
func (ime *footnoteDefinitionMarkdownElement) SetDefinition(definition []MarkdownElement) {
	ime.Definition = definition
}

func (ime *footnoteDefinitionMarkdownElement) GetContent() []MarkdownElement {
	return ime.Content
}

func (ime *footnoteDefinitionMarkdownElement) Kind() string {
	return ElementKindFootnoteDefinition
}
func (ime *footnoteDefinitionMarkdownElement) AsMarkdownString() string {
	return "[^" + GlueToString(ime.Content) + "]: " + GlueToString(ime.Definition)
}
func NewFootnoteDefinitionMarkdownElement(input string, parserFn func(input string) []MarkdownElement) FootnoteDefinitionMarkdownElement {
	return &footnoteDefinitionMarkdownElement{
		Content: parserFn(input[2 : len(input)-3]),
	}
}
