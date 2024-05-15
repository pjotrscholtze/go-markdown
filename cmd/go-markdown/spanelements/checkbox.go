package spanelements

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

func ParseLineCheckboxElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)
	for _, entry := range input {
		if entry.Kind() != entity.ElementKindText {
			res = append(res, entry)
			continue
		}
		for _, entry := range util.FindPatternsAndNonPatterns(`\[[^\[]\]`, entry.AsMarkdownString(), entity.ElementKindCheckbox, entity.ElementKindText) {
			if entry.Type == entity.ElementKindCheckbox {
				res = append(res, entity.NewCheckboxMarkdownElement(entry.Content, parserFn))
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
