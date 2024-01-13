package blockelements

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func parseLineBlockquoteElement(input []entity.LineElement) []entity.LineElement {
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
		var list []string
		var preLines []string
		for _, line := range lines {
			match, err := regexp.MatchString(`^>`, line)
			if err != nil {
				fmt.Println("An error occurred:", err, " therfore, ignoring it as a list")
				res = append(res, entry)
				continue

			}
			if match {
				list = append(list, line)
			} else if len(list) > 0 {
				if len(preLines) > 0 {
					res = append(res, entity.LineElement{
						Type:    entity.ElementKindText,
						Content: strings.Join(preLines, "\n"),
					})
				}
				res = append(res, entity.LineElement{
					Type:    entity.ElementKindBlockquote,
					Content: strings.Join(list, "\n"),
				})
				preLines = []string{line}
				list = nil
			} else {
				preLines = append(preLines, line)
			}
		}
		if len(list) > 0 {
			if len(preLines) > 0 {
				res = append(res, entity.LineElement{
					Type:    entity.ElementKindText,
					Content: strings.Join(preLines, "\n"),
				})
				preLines = []string{}
			}
			res = append(res, entity.LineElement{
				Type:    entity.ElementKindBlockquote,
				Content: strings.Join(list, "\n"),
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
