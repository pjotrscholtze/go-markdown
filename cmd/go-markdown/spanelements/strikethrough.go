package spanelements

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

func ParseLineStrikethroughElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)
	for _, entry := range input {
		if entry.Kind() != entity.ElementKindText {
			res = append(res, entry)
			continue
		}
		for _, entry := range util.FindPatternsAndNonPatterns(`~~[^~]+~~`, entry.AsMarkdownString(), entity.ElementKindStrikethrough, entity.ElementKindText) {
			if entry.Type == entity.ElementKindStrikethrough {
				res = append(res, entity.NewStrikethroughMarkdownElement(entry.Content, parserFn))
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
