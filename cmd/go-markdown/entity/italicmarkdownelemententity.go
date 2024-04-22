package entity

type italicMarkdownElement struct {
	WrappingSymbol rune
	Content        []MarkdownElement
}
type ItalicMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	GetWrappingSymbolAsRune() rune
	SetWrappingSymbolAsRune(symb rune)
}

func (ime *italicMarkdownElement) Kind() string {
	return ElementKindItalic
}
func (ime *italicMarkdownElement) GetWrappingSymbolAsRune() rune {
	return ime.WrappingSymbol
}
func (ime *italicMarkdownElement) SetWrappingSymbolAsRune(symb rune) {
	ime.WrappingSymbol = symb
}
func (ime *italicMarkdownElement) AsMarkdownString() string {
	wrappingSymbol := string(ime.WrappingSymbol)
	return wrappingSymbol + GlueToString(ime.Content) + wrappingSymbol
}
func NewItalicMarkdownElement(input string, parserFn func(input string) []MarkdownElement) ItalicMarkdownElement {
	inputAsRunes := []rune(input)
	symbol := inputAsRunes[0]
	content := string(inputAsRunes[1 : len(inputAsRunes)-1])
	return &italicMarkdownElement{
		WrappingSymbol: symbol,
		Content:        parserFn(content),
	}
}
