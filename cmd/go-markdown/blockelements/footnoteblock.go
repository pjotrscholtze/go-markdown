package blockelements

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

func ParseFootnoteDefinitionElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)
	justHadAFootnote := false
	for _, entry := range input {
		if entry.Kind() != entity.ElementKindText {
			res = append(res, entry)
			continue
		}
		for _, entry := range util.FindPatternsAndNonPatterns(`^\[\^[^\]]+\]: `, entry.AsMarkdownString(), entity.ElementKindFootnote, entity.ElementKindText) {
			if entry.Type == entity.ElementKindFootnote {
				res = append(res, entity.NewFootnoteDefinitionMarkdownElement(entry.Content, parserFn))
				justHadAFootnote = true
			} else {
				if justHadAFootnote {
					(res[len(res)-1].(entity.FootnoteDefinitionMarkdownElement)).SetDefinition(parserFn((entry.Content)))
					justHadAFootnote = false
					continue
				}
				res = append(res, &entity.LineElement{
					Type:    entry.Type,
					Content: entry.Content,
				})
			}
		}
	}
	return res
}
