package entity

import (
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

type codeBlockMarkdownElement struct {
	Content string
}
type CodeBlockMarkdownElement interface {
	WithoutFirstLine() string
	FirstLine() string
	AsMarkdownString() string
	Kind() string
}

func (bqme *codeBlockMarkdownElement) WithoutFirstLine() string {
	return strings.Join(util.SplitOnNewLine(bqme.Content)[1:], "")
}

func (bqme *codeBlockMarkdownElement) FirstLine() string {
	return util.SplitOnNewLine(bqme.Content)[0]
}

func (bqme *codeBlockMarkdownElement) Kind() string {
	return ElementKindCodeblock
}
func (bqme *codeBlockMarkdownElement) AsMarkdownString() string {
	return "```" + bqme.Content + "```"
}
func (bqme *codeBlockMarkdownElement) GetContent() string {
	return bqme.Content
}
func NewCodeBlockMarkdownElement(input string, parserFn func(input string) []MarkdownElement) CodeBlockMarkdownElement {
	// parserFn is not used, since in an code box no sub elements can exist.
	// However, to keep the interface similiar between all entities, this
	// parameter has been created.
	return &codeBlockMarkdownElement{
		Content: input[3 : len(input)-3],
	}
}
