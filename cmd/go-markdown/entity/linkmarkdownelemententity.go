package entity

import "strings"

type linkMarkdownElement struct {
	Content string
	Url     string
	Title   *string
}
type LinkMarkdownElement interface {
	AsMarkdownString() string
	Kind() string

	GetContent() string
	GetUrl() string
	GetTitle() *string
	SetTitle(title string)
	SetContent(Content string)
	SetUrl(url string)
}

func (icme *linkMarkdownElement) GetContent() string {
	return icme.Content
}
func (icme *linkMarkdownElement) GetUrl() string {
	return icme.Url
}
func (icme *linkMarkdownElement) GetTitle() *string {
	return icme.Title
}
func (icme *linkMarkdownElement) SetTitle(title string) {
	if len(title) == 0 {
		icme.Title = nil
		return
	}
	icme.Title = &title
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
	title := ""
	if icme.Title != nil {
		title = " \"" + *icme.Title + "\""
	}
	return "[" + icme.Content + "](" + icme.Url + title + ")"
}
func NewLinkMarkdownElement(input string) LinkMarkdownElement {
	parts := strings.Split(input, "](")
	urlParts := strings.Split(parts[1][:len(parts[1])-1], " \"")
	url := urlParts[0]
	var title *string
	if len(urlParts) > 1 {
		protoTitle := urlParts[1][:len(urlParts[1])-1]
		title = &protoTitle
	}
	return &linkMarkdownElement{
		Content: parts[0][1:],
		Url:     url,
		Title:   title,
	}
}
