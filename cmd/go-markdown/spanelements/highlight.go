package spanelements

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

func parseLineHighlightElement(input []entity.LineElement) []entity.LineElement {
	res := make([]entity.LineElement, 0)
	for _, entry := range input {
		if entry.Type != entity.ElementKindText {
			res = append(res, entry)
			continue
		}
		for _, entry := range util.FindPatternsAndNonPatterns(`==[^=]+==`, entry.Content, entity.ElementKindHighlight, entity.ElementKindText) {
			res = append(res, entity.LineElement{
				Type:    entry.Type,
				Content: entry.Content,
			})
		}
	}
	return res
}
