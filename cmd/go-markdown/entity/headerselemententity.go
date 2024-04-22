package entity

import "errors"

type headerMarkdownElement struct {
	Content      []MarkdownElement
	HeadingLevel uint8
}
type HeaderMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	GetHeadingLevel() uint8
	SetHeadingLevel(level uint8) error
	GetChildren() []MarkdownElement
}

func (bqme *headerMarkdownElement) GetChildren() []MarkdownElement {
	out := make([]MarkdownElement, len(bqme.Content))
	for i := range bqme.Content {
		out[i] = bqme.Content[i]
	}
	return out
}

func (bqme *headerMarkdownElement) GetHeadingLevel() uint8 {
	return bqme.HeadingLevel
}
func (bqme *headerMarkdownElement) SetHeadingLevel(level uint8) error {
	if level > 7 {
		return errors.New("Heading level cannot be higher than 6")
	}
	bqme.HeadingLevel = level
	return nil
}

func (bqme *headerMarkdownElement) Kind() string {
	return ElementKindHeader
}
func (bqme *headerMarkdownElement) AsMarkdownString() string {
	heading := ""
	for i := uint8(0); i < bqme.HeadingLevel; i++ {
		heading += "#"
	}
	return heading + GlueToString(bqme.Content)
}
func NewHeaderMarkdownElement(input string, parserFn func(input string) []MarkdownElement) HeaderMarkdownElement {
	i := 0
	for ; i < 7 && i < len(input) && input[i] == '#'; i++ {
	}

	return &headerMarkdownElement{
		Content:      parserFn(input[i:]),
		HeadingLevel: uint8(i),
	}
}
