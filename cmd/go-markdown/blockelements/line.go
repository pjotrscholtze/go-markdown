package blockelements

import (
	"regexp"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func ParseLineLineElement(input []entity.MarkdownElement, parserFn func(input []entity.MarkdownElement) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)

	for _, entry := range input {
		if entry.Kind() != entity.ElementKindText {
			res = append(res, entry)
			continue
		}
		if entry.AsMarkdownString() == "" {
			continue
		}
		r := regexp.MustCompile("\r?\n") //.Split(inputString, -1)
		content := entry.AsMarkdownString()
		matches := r.FindAllIndex([]byte(content), -1)

		// lines := strings.Split(entry.AsMarkdownString(), "\n")
		before := 0
		for _, match := range matches {
			line := content[before:match[1]]
			before = match[1]

			res = append(res, &entity.LineElement{
				Type:    entity.ElementKindText,
				Content: line,
			})
		}
		res = append(res, &entity.LineElement{
			Type:    entity.ElementKindText,
			Content: content[:before],
		})
	}
	return res
}
