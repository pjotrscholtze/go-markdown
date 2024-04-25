package entity

import (
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
	out := []string{}
	for i, line := range util.SplitOnNewLine(content) {
		if i == 0 {
			out = append(out, bqme.Symbol+line)
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "")
}
func NewListItemElementMarkdownElement(input string, parserFn func(input string) []MarkdownElement) MarkdownElement {
	idx := 0
	symbol := ""
	lines := []string{}
	for i, line := range util.SplitOnNewLine(input) {
		if i == 0 {
			idx = strings.Index(line, " ")
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
