package entity

import "strings"

type imageMarkdownElement struct {
	Alt   string
	Url   string
	Title *string
}
type ImageMarkdownElement interface {
	AsMarkdownString() string
	Kind() string

	GetAlt() string
	GetUrl() string
	GetTitle() *string
	SetTitle(title string)
	SetAlt(alt string)
	SetUrl(url string)
}

func (icme *imageMarkdownElement) GetAlt() string {
	return icme.Alt
}
func (icme *imageMarkdownElement) GetUrl() string {
	return icme.Url
}
func (icme *imageMarkdownElement) GetTitle() *string {
	return icme.Title
}
func (icme *imageMarkdownElement) SetTitle(title string) {
	if len(title) == 0 {
		icme.Title = nil
		return
	}
	icme.Title = &title
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
	title := ""
	if icme.Title != nil {
		title = " \"" + *icme.Title + "\""
	}
	return "![" + icme.Alt + "](" + icme.Url + title + ")"
}
func NewImageMarkdownElement(input string) ImageMarkdownElement {
	parts := strings.Split(input, "](")
	urlParts := strings.Split(parts[1][:len(parts[1])-1], " \"")
	url := urlParts[0]
	var title *string
	if len(urlParts) > 1 {
		protoTitle := urlParts[1][:len(urlParts[1])-1]
		title = &protoTitle
	}
	return &imageMarkdownElement{
		Alt:   parts[0][2:],
		Url:   url,
		Title: title,
	}
}
