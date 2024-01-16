package entity

import "strings"

type imageMarkdownElement struct {
	Alt string
	Url string
	// @todo title support
}
type ImageMarkdownElement interface {
	AsMarkdownString() string
	Kind() string

	GetAlt() string
	GetUrl() string
	SetAlt(alt string)
	SetUrl(url string)
}

func (icme *imageMarkdownElement) GetAlt() string {
	return icme.Alt
}
func (icme *imageMarkdownElement) GetUrl() string {
	return icme.Url
}
func (icme *imageMarkdownElement) SetAlt(alt string) {
	icme.Alt = alt
}
func (icme *imageMarkdownElement) SetUrl(url string) {
	icme.Url = url
}

func (icme *imageMarkdownElement) Kind() string {
	return ElementKindImage
}
func (icme *imageMarkdownElement) AsMarkdownString() string {
	return "![" + icme.Alt + "](" + icme.Url + ")"
}
func NewImageMarkdownElement(input string) ImageMarkdownElement {
	parts := strings.Split(input, "](")
	return &imageMarkdownElement{
		Alt: parts[0][2:],
		Url: parts[1][:len(parts[1])-1],
	}
}
