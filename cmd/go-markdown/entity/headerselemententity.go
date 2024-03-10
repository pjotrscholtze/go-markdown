package entity

import "errors"

type headerMarkdownElement struct {
	Content      string
	HeadingLevel uint8
}
type HeaderMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	GetHeadingLevel() uint8
	SetHeadingLevel(level uint8) error
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
	return heading + bqme.Content
}
func NewHeaderMarkdownElement(input string) HeaderMarkdownElement {
	i := 0
	for ; i < 7 && i < len(input) && input[i] == '#'; i++ {
	}

	return &headerMarkdownElement{
		Content:      input[i:],
		HeadingLevel: uint8(i),
	}
}
