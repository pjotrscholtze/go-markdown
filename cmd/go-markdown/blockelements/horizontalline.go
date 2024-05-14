package blockelements

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

func ParseLineHorizontalLineElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)
	for _, entry := range input {
		if entry.Kind() != entity.ElementKindText {
			res = append(res, entry)
			continue
		}

		var tmp *entity.LineElement
		for _, entry := range util.FindPatternsAndNonPatterns(`((\*( )?)|(\-( )?)|(_( )?)){3,}`, entry.AsMarkdownString(), entity.ElementKindHorizontalLine, entity.ElementKindText) {
			if entry.Type == entity.ElementKindHorizontalLine {
				if tmp != nil {
					res = append(res, tmp)
					tmp = nil
				}

				res = append(res, entity.NewHorizontalLineMarkdownElement(entry.Content, parserFn))
			} else {
				if tmp == nil {
					tmp = &entity.LineElement{
						Type:    entry.Type,
						Content: "",
					}
				}

				tmp.Content += entry.Content
			}
		}
		if tmp != nil {
			res = append(res, tmp)
		}
	}
	return res
}
