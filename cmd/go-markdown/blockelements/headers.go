package blockelements

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

const (
	headerRegex        = `^(#{1,6}[ ]?[\d\w ]+[#]*)|([\w\d ]+\n[-=]{4,})`
	blockquoteRegex    = `^>`
	tableRegex         = `^(\|.+\|\n\|([ \-:]+\|)+\n)((\|.+\|\n)+)`
	listItemCandidate  = `^\s*(([\-+\*]|(\d+\.)).*)+$`
	codeblockStart     = `^\x60{3}`
	codeblockEnd       = `\x60{3}$`
	horizontalLine     = `((\*( )?)|(\-( )?)|(_( )?)){3,}$`
	termDefinitionLine = `^: `
)

// func parseLineHeaderElement(input []entity.LineElement) []entity.LineElement {
// 	res := make([]entity.LineElement, 0)
// 	for _, entry := range input {
// 		if entry.Type != "text" {
// 			res = append(res, entry)
// 			continue
// 		}
// 		for _, line := range strings.Split(entry.Content, "\n") {
// 			//
// 		}
// 		for _, entry := range util.FindPatternsAndNonPatterns(`^(#( )?){1,6}[^#]`, entry.Content, entity.ElementKindHeader, entity.ElementKindText) {
// 			res = append(res, entity.LineElement{
// 				Type:    entry.Type,
// 				Content: entry.Content,
// 			})
// 		}
// 	}
// 	return res
// }

func parseLineHeaderElement(input []entity.LineElement) []entity.LineElement {
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
		// var list []string
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
					res = append(res, entity.LineElement{
						Type:    entity.ElementKindText,
						Content: strings.Join(preLines, "\n"),
					})
					preLines = nil
				}
				res = append(res, entity.LineElement{
					Type:    entity.ElementKindHeader,
					Content: line,
				})

			} else {
				preLines = append(preLines, line)
			}
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
