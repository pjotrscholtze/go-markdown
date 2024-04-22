package blockelements

import (
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func ParseLineTableElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)

	for _, entry := range input {
		if entry.Kind() != entity.ElementKindText {
			res = append(res, entry)
			continue
		}
		if entry.AsMarkdownString() == "" {
			continue
		}
		lines := strings.Split(entry.AsMarkdownString(), "\n")
		var table []string
		var preLines []string
		for _, line := range lines {
			if strings.HasPrefix(line, "|") {
				table = append(table, line)
			} else if len(table) > 0 {
				if len(preLines) > 0 {
					res = append(res, &entity.LineElement{
						Type:    entity.ElementKindText,
						Content: strings.Join(preLines, "\n"),
					})
				}
				res = append(res, entity.NewTableElementMarkdownElement(strings.Join(table, "\n"), parserFn))
				preLines = []string{line}
				table = nil
			} else {
				preLines = append(preLines, line)
			}
		}
		if len(table) > 0 {
			if len(preLines) > 0 {
				res = append(res, &entity.LineElement{
					Type:    entity.ElementKindText,
					Content: strings.Join(preLines, "\n"),
				})
				preLines = []string{}
			}
			res = append(res, entity.NewTableElementMarkdownElement(strings.Join(table, "\n"), parserFn))
		}
		if len(preLines) > 0 {
			res = append(res, &entity.LineElement{
				Type:    entity.ElementKindText,
				Content: strings.Join(preLines, "\n"),
			})
		}
	}
	return res
}
