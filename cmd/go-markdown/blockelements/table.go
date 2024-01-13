package blockelements

import (
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func parseLineTableElement(input []entity.LineElement) []entity.LineElement {
	res := make([]entity.LineElement, 0)

	for _, entry := range input {
		if entry.Type != entity.ElementKindText {
			res = append(res, entry)
			continue
		}
		if entry.Content == "" {
			continue
		}
		lines := strings.Split(entry.Content, "\n")
		var table []string
		var preLines []string
		for _, line := range lines {
			if strings.HasPrefix(line, "|") {
				table = append(table, line)
			} else if len(table) > 0 {
				if len(preLines) > 0 {
					res = append(res, entity.LineElement{
						Type:    entity.ElementKindText,
						Content: strings.Join(preLines, "\n"),
					})
				}
				res = append(res, entity.LineElement{
					Type:    entity.ElementKindTable,
					Content: strings.Join(table, "\n"),
				})
				preLines = []string{line}
				table = nil
			} else {
				preLines = append(preLines, line)
			}
		}
		if len(table) > 0 {
			if len(preLines) > 0 {
				res = append(res, entity.LineElement{
					Type:    entity.ElementKindText,
					Content: strings.Join(preLines, "\n"),
				})
				preLines = []string{}
			}
			res = append(res, entity.LineElement{
				Type:    entity.ElementKindTable,
				Content: strings.Join(table, "\n"),
			})
		}
		if len(preLines) > 0 {
			res = append(res, entity.LineElement{
				Type:    entity.ElementKindText,
				Content: strings.Join(preLines, "\n"),
			})
		}
	}
	return res
}
