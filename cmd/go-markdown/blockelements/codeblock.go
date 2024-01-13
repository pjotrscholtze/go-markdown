package blockelements

import (
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func parseLineCodeblockElement(input []entity.LineElement) []entity.LineElement {
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
		var chunkOfLines []string
		isCodeBlock := false
		for _, line := range lines {
			if strings.HasPrefix(strings.TrimLeft(line, " \t"), "```") {
				if isCodeBlock {
					chunkOfLines = append(chunkOfLines, line)
					res = append(res, entity.LineElement{
						Type:    entity.ElementKindCodeblock,
						Content: strings.Join(chunkOfLines, "\n"),
					})
					chunkOfLines = nil
					isCodeBlock = false
					continue
				} else {
					if len(chunkOfLines) > 0 {
						res = append(res, entity.LineElement{
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
				res = append(res, entity.LineElement{
					Type:    entity.ElementKindCodeblock,
					Content: strings.Join(chunkOfLines, "\n"),
				})
			} else {
				res = append(res, entity.LineElement{
					Type:    entity.ElementKindText,
					Content: strings.Join(chunkOfLines, "\n"),
				})
			}
		}
	}
	return res
}
