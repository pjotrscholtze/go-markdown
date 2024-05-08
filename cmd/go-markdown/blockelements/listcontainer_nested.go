package blockelements

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func parseListContainerNestedElement(entry entity.ListElementMarkdownElement) entity.ListElementMarkdownElement {
	whiteSpace := entry.GetWhitespaceInFront()
	hasNesting := false
	for _, element := range entry.GetContent() {
		if element.GetWhitespaceInFront() != whiteSpace {
			hasNesting = true
			break
		}
	}

	if !hasNesting {
		return entry
	}

	// We're dealing with a nested list here.

	newList := entity.NewListElementMarkdownElement()
	var nestedList entity.ListElementMarkdownElement
	for _, element := range entry.GetContent() {
		if element.GetWhitespaceInFront() == whiteSpace {
			if nestedList != nil {
				// Recursive call here
				newList.AppendList(parseListContainerNestedElement(nestedList))
			}

			if element.List != nil {
				newList.AppendList(element.List)
			} else {
				newList.AppendListItem(element.ListItem)
			}

			nestedList = nil
			continue
		}
		if nestedList == nil {
			nestedList = entity.NewListElementMarkdownElement()
		}
		if element.List != nil {
			nestedList.AppendList(element.List)
		} else {
			nestedList.AppendListItem(element.ListItem)
		}
	}

	if nestedList != nil {
		// Recursive call here
		newList.AppendList(parseListContainerNestedElement(nestedList))
	}

	return newList
}
func ParseListContainerNestedElement(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)
	for _, inputEntry := range input {
		if inputEntry.Kind() != entity.ElementKindList {
			res = append(res, inputEntry)
			continue
		}

		entry := inputEntry.(entity.ListElementMarkdownElement)
		res = append(res, parseListContainerNestedElement(entry))
	}
	return res
}
