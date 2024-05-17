package blockelements

import (
	"regexp"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func ParseParagraphElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)
	r := regexp.MustCompile(`(?m)\n{2,}`)
	paragraph := ""
	for _, entry := range input {
		if entry.Kind() != entity.ElementKindText {
			if len(paragraph) > 0 {
				res = append(res, entity.NewParagraphMarkdownElement(paragraph, parserFn))
				paragraph = ""
			}
			res = append(res, entry)
			continue
		}
		s := entry.AsMarkdownString()
		if s == "" {
			continue
		}
		matches := r.FindAllStringSubmatchIndex(s, -1)
		prev := 0
		for _, m := range matches {
			res = append(res,
				entity.NewParagraphMarkdownElement(s[prev:m[0]], parserFn),
			)
			prev = m[1]
		}
		res = append(res,
			entity.NewParagraphMarkdownElement(s[prev:], parserFn),
		)

	}

	return res
}
