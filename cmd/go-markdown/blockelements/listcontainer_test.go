package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestListContainerDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{
		// {name: "No content", input: "", expect: []entity.MarkdownElement{}},
		{name: "Text before the list", input: `Some text before the list.
- List item`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "Some text before the list.\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: "- List item",
			}}},
		{name: "List item with no content before", input: `- One
- Two
- Two dot one
- Three`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: "- One\n- Two\n- Two dot one\n- Three",
			},
		}},

		{name: "List item with bold text", input: `- Item with **bold text**`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: `- Item with **bold text**`,
			}}},
		{name: "List item with an ordered marker", input: `1. One
2. Two
2.1 Two dot one
3. Three`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: "1. One\n2. Two\n2.1 Two dot one\n3. Three",
			}}},
		{name: "List item with continuous indentation", input: `- One
	- One dot one
- Two
	- Two dot one
- Three
	- Three dot one`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: "- One\n	- One dot one\n- Two\n	- Two dot one\n- Three\n	- Three dot one",
			}}},
		{name: "List item with nested list starting on the first line", input: `* * Sublist item 1.1
* List item 2
	* Sublist item 2.1
	* Sublist item 2.2
* List item 3`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: "* * Sublist item 1.1\n* List item 2\n	* Sublist item 2.1\n	* Sublist item 2.2\n* List item 3"}}},
		{name: "List item with a blank line between the items", input: `+ list item 1

+ list item 2`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: "+ list item 1\n",
			},
			&entity.LineElement{Type: entity.ElementKindText, Content: "\n"},
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: `+ list item 2`}}},
		{name: "Text, a heading, and a paragraph before and after the list", input: `# Heading
Some text before the list.
This is a paragraph before the list.
- List item
Some text after the list.
This is a paragraph after the list.
# Heading`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `# Heading
Some text before the list.
This is a paragraph before the list.` + "\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindList,
				Content: "- List item\n",
			},
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `Some text after the list.
This is a paragraph after the list.
# Heading`,
			}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseListContainerElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}},
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
