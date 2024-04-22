package document

import (
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

type Document struct {
	elems []entity.MarkdownElement
}

func (d *Document) AsMarkdownString() string {
	out := []string{}
	for _, elm := range d.elems {
		out = append(out, elm.AsMarkdownString())
	}

	return strings.Join(out, "")
}

func NewDocumentFromString(inpt string) *Document {
	return &Document{}
}
