package entity

import "strings"

type imageMarkdownElement struct {
	Alt   []MarkdownElement
	Url   string
	Title []MarkdownElement
}
type ImageMarkdownElement interface {
	AsMarkdownString() string
	Kind() string

	GetAlt() string
	GetUrl() string
	GetTitle() *string
	SetTitle(title string, parserFn func(input string) []MarkdownElement)
	SetAlt(alt string, parserFn func(input string) []MarkdownElement)
	SetUrl(url string)
}

func (icme *imageMarkdownElement) GetAlt() string {
	return GlueToString(icme.Alt)
}
func (icme *imageMarkdownElement) GetUrl() string {
	return icme.Url
}
func (icme *imageMarkdownElement) GetTitle() *string {
	if len(icme.Title) == 0 {
		return nil
	}
	title := GlueToString(icme.Title)
	return &title
}
func (icme *imageMarkdownElement) SetTitle(title string, parserFn func(input string) []MarkdownElement) {
	if len(title) == 0 {
		icme.Title = nil
		return
	}
	icme.Title = parserFn(title)
}
func (icme *imageMarkdownElement) SetAlt(alt string, parserFn func(input string) []MarkdownElement) {
	icme.Alt = parserFn(alt)
}
func (icme *imageMarkdownElement) SetUrl(url string) {
	icme.Url = url
}

func (icme *imageMarkdownElement) Kind() string {
	return ElementKindImage
}
func (icme *imageMarkdownElement) AsMarkdownString() string {
	title := ""
	if len(icme.Title) > 0 {
		title = " \"" + GlueToString(icme.Title) + "\""
	}
	return "![" + GlueToString(icme.Alt) + "](" + icme.Url + title + ")"
}
func NewImageMarkdownElement(input string, parserFn func(input string) []MarkdownElement) ImageMarkdownElement {
	parts := strings.Split(input, "](")
	urlParts := strings.Split(parts[1][:len(parts[1])-1], " \"")
	url := urlParts[0]
	title := ""
	if len(urlParts) > 1 {
		title = urlParts[1][:len(urlParts[1])-1]
	}
	alt := []MarkdownElement{}
	if parts[0][2:] != "" {
		alt = parserFn(parts[0][2:])
	}
	titleArray := []MarkdownElement{}
	if title != "" {
		titleArray = parserFn(title)
	}
	return &imageMarkdownElement{
		Alt:   alt,
		Url:   url,
		Title: titleArray,
	}
}
