package blockelements

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func parseLineHeaderElement(input []entity.MarkdownElement) []entity.MarkdownElement {
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
		var preLines []string
		for _, line := range lines {
			match, err := regexp.MatchString(`^(#( )?){1,6}[^#]`, line)
			if err != nil {
				fmt.Println("An error occurred:", err, " therfore, ignoring it as a heading")
				res = append(res, entry)
				continue

			}
			if match {
				if len(preLines) > 0 {
					res = append(res, &entity.LineElement{
						Type:    entity.ElementKindText,
						Content: strings.Join(preLines, "\n"),
					})
					preLines = nil
				}
				res = append(res, entity.NewHeaderMarkdownElement(line))

			} else {
				preLines = append(preLines, line)
			}
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
