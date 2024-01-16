package entity

import "strings"

type linkMarkdownElement struct {
	Content string
	Url     string
	// @todo title support
}
type LinkMarkdownElement interface {
	AsMarkdownString() string
	Kind() string

	GetContent() string
	GetUrl() string
	SetContent(Content string)
	SetUrl(url string)
}

func (icme *linkMarkdownElement) GetContent() string {
	return icme.Content
}
func (icme *linkMarkdownElement) GetUrl() string {
	return icme.Url
}
func (icme *linkMarkdownElement) SetContent(Content string) {
	icme.Content = Content
}
func (icme *linkMarkdownElement) SetUrl(url string) {
	icme.Url = url
}

func (icme *linkMarkdownElement) Kind() string {
	return ElementKindLink
}
func (icme *linkMarkdownElement) AsMarkdownString() string {
	return "[" + icme.Content + "](" + icme.Url + ")"
}
func NewLinkMarkdownElement(input string) LinkMarkdownElement {
	parts := strings.Split(input, "](")
	return &linkMarkdownElement{
		Content: parts[0][1:],
		Url:     parts[1][:len(parts[1])-1],
	}
}
