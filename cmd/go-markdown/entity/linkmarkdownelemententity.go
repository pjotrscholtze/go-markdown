package entity

import "strings"

type linkMarkdownElement struct {
	Content []MarkdownElement
	Url     string
	Title   []MarkdownElement
}
type LinkMarkdownElement interface {
	AsMarkdownString() string
	Kind() string

	GetContent() string
	GetUrl() string
	GetTitle() *string
	SetTitle(title string, parserFn func(input string) []MarkdownElement)
	SetContent(Content string, parserFn func(input string) []MarkdownElement)
	SetUrl(url string)
}

func (icme *linkMarkdownElement) GetContent() string {
	return GlueToString(icme.Content)
}
func (icme *linkMarkdownElement) GetUrl() string {
	return icme.Url
}
func (icme *linkMarkdownElement) GetTitle() *string {
	if len(icme.Title) == 0 {
		return nil
	}
	title := GlueToString(icme.Title)
	return &title
}
func (icme *linkMarkdownElement) SetTitle(title string, parserFn func(input string) []MarkdownElement) {
	if len(title) == 0 {
		icme.Title = nil
		return
	}
	icme.Title = parserFn(title)
}
func (icme *linkMarkdownElement) SetContent(Content string, parserFn func(input string) []MarkdownElement) {
	icme.Content = parserFn(Content)
}
func (icme *linkMarkdownElement) SetUrl(url string) {
	icme.Url = url
}

func (icme *linkMarkdownElement) Kind() string {
	return ElementKindLink
}
func (icme *linkMarkdownElement) AsMarkdownString() string {
	title := ""
	if len(icme.Title) > 0 {
		title = " \"" + GlueToString(icme.Title) + "\""
		if title == ` ""` {
			title = ""
		}
	}
	return "[" + GlueToString(icme.Content) + "](" + icme.Url + title + ")"
}
func NewLinkMarkdownElement(input string, parserFn func(input string) []MarkdownElement) LinkMarkdownElement {
	parts := strings.Split(input, "](")
	urlParts := strings.Split(parts[1][:len(parts[1])-1], " \"")
	url := urlParts[0]
	title := ""
	if len(urlParts) > 1 {
		title = urlParts[1][:len(urlParts[1])-1]
	}
	titleArray := parserFn(title)
	if GlueToString(titleArray) == "" {
		titleArray = []MarkdownElement{}
	}
	return &linkMarkdownElement{
		Content: parserFn(parts[0][1:]),
		Url:     url,
		Title:   titleArray,
	}
}
