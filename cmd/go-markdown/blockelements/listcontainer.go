package blockelements

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func ParseListContainerDefinition(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement {
	res := make([]entity.MarkdownElement, 0)
	listContainer := entity.NewListElementMarkdownElement()
	for _, inputEntry := range input {
		for _, entry := range ParseLineListItemElement(inputEntry.AsMarkdownString(), parserFn) {
			if entry.Kind() != entity.ElementKindListItem {
				if listContainer.ItemCount() > 0 {
					res = append(res, listContainer)
				}
				res = append(res, entry)
				listContainer = entity.NewListElementMarkdownElement()
				continue
			}
			listItem := entry.(entity.ListItemElementMarkdownElement)
			listContainer.AppendListItem(listItem)
		}
		if listContainer.ItemCount() > 0 {
			res = append(res, listContainer)
			listContainer = entity.NewListElementMarkdownElement()
		}
	}
	if listContainer.ItemCount() > 0 {
		res = append(res, listContainer)
		listContainer = entity.NewListElementMarkdownElement()
	}
	return res
}
