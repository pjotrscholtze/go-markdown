package blockelements

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func parseLineListElement(input []entity.MarkdownElement) []entity.MarkdownElement {
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
		var list []string
		var preLines []string
		for _, line := range lines {
			match, err := regexp.MatchString(`^\s*(([\-+\*]|(\d+\.)).*)+$`, line)
			if err != nil {
				fmt.Println("An error occurred:", err, " therfore, ignoring it as a list")
				res = append(res, entry)
				continue

			}
			if match {
				list = append(list, line)
			} else if len(list) > 0 {
				if len(preLines) > 0 {
					res = append(res, &entity.LineElement{
						Type:    entity.ElementKindText,
						Content: strings.Join(preLines, "\n"),
					})
				}
				res = append(res, entity.NewListElementMarkdownElement(strings.Join(list, "\n")))
				preLines = []string{line}
				list = nil
			} else {
				preLines = append(preLines, line)
			}
		}
		if len(list) > 0 {
			if len(preLines) > 0 {
				res = append(res, &entity.LineElement{
					Type:    entity.ElementKindText,
					Content: strings.Join(preLines, "\n"),
				})
				preLines = []string{}
			}
			res = append(res, entity.NewListElementMarkdownElement(strings.Join(list, "\n")))
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
