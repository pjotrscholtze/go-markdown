package entity

type boldMarkdownElement struct {
	WrappingSymbol rune
	Content        []MarkdownElement
}
type BoldMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	GetWrappingSymbolAsRune() rune
	SetWrappingSymbolAsRune(symb rune)
}

func (bme *boldMarkdownElement) Kind() string {
	return ElementKindBold
}
func (bme *boldMarkdownElement) GetWrappingSymbolAsRune() rune {
	return bme.WrappingSymbol
}
func (bme *boldMarkdownElement) SetWrappingSymbolAsRune(symb rune) {
	bme.WrappingSymbol = symb
}
func (bme *boldMarkdownElement) AsMarkdownString() string {
	wrappingSymbol := string(bme.WrappingSymbol) + string(bme.WrappingSymbol)
	return wrappingSymbol + GlueToString(bme.Content) + wrappingSymbol
}
func NewBoldMarkdownElement(input string, parserFn func(input string) []MarkdownElement) BoldMarkdownElement {
	inputAsRunes := []rune(input)
	symbol := inputAsRunes[0]
	content := string(inputAsRunes[2 : len(inputAsRunes)-2])
	return &boldMarkdownElement{
		WrappingSymbol: symbol,
		Content:        parserFn(content),
	}
}
