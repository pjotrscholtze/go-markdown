package entity

import (
	"regexp"
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

type listItemElementMarkdownElement struct {
	Symbol  string
	Content []MarkdownElement
}
type ListItemElementMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	GetSymbol() string
	GetContentAsString() string
	GetContentAsStringMultiLine() string
	SymbolLength() int
	AddContent(content []MarkdownElement)
}

func (bqme *listItemElementMarkdownElement) AddContent(content []MarkdownElement) {
	bqme.Content = append(bqme.Content, content...)
}

func (bqme *listItemElementMarkdownElement) SymbolLength() int {
	return len(bqme.Symbol)
}

func (bqme *listItemElementMarkdownElement) GetContentAsStringMultiLine() string {
	out := ""
	for _, line := range util.SplitOnNewLine(GlueToString(bqme.Content)) {
		out += line[min(len(line), bqme.SymbolLength()):]
	}
	return out
}

func (bqme *listItemElementMarkdownElement) GetSymbol() string {
	return bqme.Symbol
}
func (bqme *listItemElementMarkdownElement) GetContentAsString() string {
	return bqme.AsMarkdownString()[len(bqme.Symbol):]
}
func (bqme *listItemElementMarkdownElement) Kind() string {
	return ElementKindListItem
}
func (bqme *listItemElementMarkdownElement) AsMarkdownString() string {
	content := GlueToString(bqme.Content)
	return bqme.Symbol + content
}
func NewListItemElementMarkdownElement(input string, parserFn func(input string) []MarkdownElement) MarkdownElement {
	symbol := ""
	findSymbol := regexp.MustCompile(`^\s*(([\-+\*]|(\d+\.)))+ `)
	lines := []string{}
	for i, line := range util.SplitOnNewLine(input) {
		if i == 0 {
			matches := findSymbol.FindAllStringSubmatchIndex(line, -1)
			idx := 0
			if len(matches) > 0 {
				idx = min(len(line), matches[0][1])
			}
			symbol = line[:idx]
			lines = append(lines, line[idx:])
			continue
		}
		lines = append(lines, line)
	}

	return &listItemElementMarkdownElement{
		Symbol:  symbol,
		Content: parserFn(strings.Join(lines, "")),
	}
}
