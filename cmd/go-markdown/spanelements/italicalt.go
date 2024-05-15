package spanelements

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

func ParseLineItalicAltElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)
	for _, entry := range input {
		if entry.Kind() != entity.ElementKindText {
			res = append(res, entry)
			continue
		}
		for _, entry := range util.FindPatternsAndNonPatterns(`_[^_]+_`, entry.AsMarkdownString(), entity.ElementKindItalic, entity.ElementKindText) {
			if entry.Type == entity.ElementKindItalic {
				res = append(res, entity.NewItalicMarkdownElement(entry.Content, parserFn))
			} else {
				res = append(res, &entity.LineElement{
					Type:    entry.Type,
					Content: entry.Content,
				})
			}
		}
	}
	return res
}
