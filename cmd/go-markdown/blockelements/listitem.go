package blockelements

import (
	"regexp"
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

func ParseLineListItemElement(input string, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)

	prelines := []string{}
	lineItemFound := false
	for _, line := range util.SplitOnNewLine(input) {
		match, _ := regexp.MatchString(`^\s*(([\-+\*]|(\d+\.?)))+(\t| )`, line)
		if match {
			if len(prelines) > 0 {
				if lineItemFound {
					res = append(res, entity.NewListItemElementMarkdownElement(strings.Join(prelines, ""), parserFn))
				} else {
					res = append(res, &entity.LineElement{
						Type:    entity.ElementKindText,
						Content: strings.Join(prelines, ""),
					})
				}
			}
			// New line item found
			prelines = []string{line}
			lineItemFound = true
		} else if lineItemFound {
			idx := min(strings.Index(prelines[0], " "), len(line))
			prefixMatch, _ := regexp.MatchString(`^\s+$`, line[:idx])
			whiteSpaceLine, _ := regexp.MatchString(`^\s+$`, line)
			if prefixMatch && !whiteSpaceLine {
				// Continuation found.
				prelines = append(prelines, line)
			} else {
				// End found
				res = append(res, entity.NewListItemElementMarkdownElement(strings.Join(prelines, ""), parserFn))
				prelines = []string{line}
				lineItemFound = false
			}
		} else {
			// Just plain text
			prelines = append(prelines, line)
		}
	}
	if len(prelines) > 0 {
		if lineItemFound {
			res = append(res, entity.NewListItemElementMarkdownElement(strings.Join(prelines, ""), parserFn))
		} else {
			res = append(res, &entity.LineElement{
				Type:    entity.ElementKindText,
				Content: strings.Join(prelines, ""),
			})
		}
	}

	return res
}
