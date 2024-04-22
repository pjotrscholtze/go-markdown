package blockelements

import (
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func ParseLineCodeblockElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
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
		var chunkOfLines []string
		isCodeBlock := false
		for _, line := range lines {
			if strings.HasPrefix(strings.TrimLeft(line, " \t"), "```") {
				if isCodeBlock {
					chunkOfLines = append(chunkOfLines, line)
					res = append(res, entity.NewCodeBlockMarkdownElement(strings.Join(chunkOfLines, "\n"), parserFn))
					chunkOfLines = nil
					isCodeBlock = false
					continue
				} else {
					if len(chunkOfLines) > 0 {
						res = append(res, &entity.LineElement{
							Type:    entity.ElementKindText,
							Content: strings.Join(chunkOfLines, "\n"),
						})
					}
					isCodeBlock = true
				}
				chunkOfLines = nil
			}
			chunkOfLines = append(chunkOfLines, line)
		}
		if len(chunkOfLines) > 0 {
			if isCodeBlock {
				res = append(res, entity.NewCodeBlockMarkdownElement(strings.Join(chunkOfLines, "\n"), parserFn))
			} else {
				res = append(res, &entity.LineElement{
					Type:    entity.ElementKindText,
					Content: strings.Join(chunkOfLines, "\n"),
				})
			}
		}
	}
	return res
}
