package parser

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

// var DefaultMarkdownOrder = []func(input []entity.MarkdownElement, parser func(input []entity.MarkdownElement) []entity.MarkdownElement) []entity.MarkdownElement{
// 	blockelements.ParseLineHeaderElement,
// 	blockelements.ParseLineTableElement,
// 	blockelements.ParseLineHorizontalLineElement,
// 	blockelements.ParseLineCodeblockElement,
// 	blockelements.ParseLineBlockquoteElement,
// 	blockelements.ParseLineListElement,
// 	blockelements.ParseLineTermDefinitionLineElement,

// 	spanelements.ParseInlineCodeElement,
// 	spanelements.ParseLineLinkElement,
// 	spanelements.ParseLineImageElement,

// 	spanelements.ParseLineCheckboxElement,

// 	spanelements.ParseLineBoldElement,
// 	spanelements.ParseLineBoldAltElement,
// 	spanelements.ParseLineEmojiElement,
// 	spanelements.ParseLineFootnoteElement,
// 	spanelements.ParseLineHighlightElement,
// 	spanelements.ParseLineItalicElement,
// 	spanelements.ParseLineItalicAltElement,
// 	spanelements.ParseLineStrikethroughElement,
// }

func ParseString(
	inpt string,
	parseOrder []func(input []entity.MarkdownElement, parser func(input string) []entity.MarkdownElement) []entity.MarkdownElement,
) []entity.MarkdownElement {

	out := []entity.MarkdownElement{
		&entity.LineElement{
			Type:    entity.ElementKindText,
			Content: inpt,
		},
	}

	for i, parseFn := range parseOrder {
		_ = i
		out = parseFn(out, func(input string) []entity.MarkdownElement {
			if i+1 == len(parseOrder) {
				return []entity.MarkdownElement{
					&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					},
				}
			}

			return ParseString(input, parseOrder[i+1:])
		})
	}

	return out
}
