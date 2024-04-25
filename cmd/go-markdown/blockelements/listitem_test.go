// list_test.go
package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestListItemDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No content", input: "", expect: []entity.MarkdownElement{}},

		{name: "Empty list item", input: `- `, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- `,
			}}},
		{name: "List item with leading spaces", input: `-   Item with leading spaces`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `-   Item with leading spaces`,
			}}},
		{name: "List item with trailing spaces", input: `- Item with trailing spaces  `, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with trailing spaces  `,
			}}},
		{name: "List item with leading and trailing spaces", input: `-   Item with leading and trailing spaces  `, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `-   Item with leading and trailing spaces  `,
			}}},
		{name: "List item with nested list", input: `- Item with nested list
1. Sublist item`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Item with nested list\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "1. Sublist item",
			}}},
		{name: "List item with inline code", input: `- Item with ` + "`inline code`", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with ` + "`inline code`",
			}}},
		{name: "List item with bold text", input: `- Item with **bold text**`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with **bold text**`,
			}}},
		{name: "List item with italic text", input: `- Item with *italic text*`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with *italic text*`,
			}}},
		{name: "List item with link", input: `- Item with [link](#)`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with [link](#)`,
			}}},
		{name: "List item with image", input: `- Item with ![image](url)`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with ![image](url)`,
			}}},
		{name: "List item with strikethrough text", input: `- Item with ~~strikethrough text~~`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with ~~strikethrough text~~`,
			}}},
		{name: "List item with task list", input: `- Item with task list
- [ ] Task 1
- [x] Task 2`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Item with task list\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- [ ] Task 1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- [x] Task 2",
			}}},
		{name: "List item with quote", input: `- Item with > quote`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with > quote`,
			}}},
		{name: "List item with horizontal rule", input: `- Item with ---`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with ---`,
			}}},
		{name: "List item with a footnote", input: `- Item with a footnote[^1]`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a footnote[^1]`,
			}}},
		{name: "List item with a heading ID", input: `- Item with a heading ID {#custom-id}`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a heading ID {#custom-id}`,
			}}},
		{name: "List item with a strikethrough", input: `- Item with a strikethrough ~~strikethrough text~~`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a strikethrough ~~strikethrough text~~`,
			}}},
		{name: "List item with a task list", input: `- Item with a task list
- [x] Task 1
- [ ] Task 2`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Item with a task list\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- [x] Task 1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- [ ] Task 2",
			}}},
		{name: "List item with an emoji", input: `- Item with an emoji :joy:`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with an emoji :joy:`,
			}}},
		{name: "List item with a highlight", input: `- Item with a highlight ==very important words==`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a highlight ==very important words==`,
			}}},
		{name: "List item with a subscript", input: `- Item with a subscript H~2~O`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a subscript H~2~O`,
			}}},
		{name: "List item with a superscript", input: `- Item with a superscript X^2^`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a superscript X^2^`,
			}}},
		{name: "List item with a non-breaking space character after it", input: `* &nbsp;
* List item`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* &nbsp;\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* List item",
			}}},
		{name: "List item with an ordered marker", input: `1. One
2. Two
2.1 Two dot one
3. Three`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "1. One\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "2. Two\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "2.1 Two dot one\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "3. Three",
			}}},
		{name: "List item with no content before", input: `- One
- Two
- Two dot one
- Three`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- One\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Two\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Two dot one\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Three",
			},
		}},
		{name: "List item with continuous indentation", input: `- One
	- One dot one
- Two
	- Two dot one
- Three
	- Three dot one`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- One\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	- One dot one\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Two\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	- Two dot one\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Three\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	- Three dot one",
			}}},
		{name: "List item with empty lines before and after", input: `... a line with text.

- One
- Two
	- Two dot one
- Three

Another line with text...`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "... a line with text.\n\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- One\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Two\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	- Two dot one\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Three\n"},
			&entity.LineElement{Type: entity.ElementKindText,
				Content: "\nAnother line with text..."},
		}},
		{name: "List item with punctuation after items", input: `- One.
- Two.
- Three.`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- One.\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Two.\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Three."}}},
		{name: "List item with nested list starting on the first line", input: `* * Sublist item 1.1
* List item 2
	* Sublist item 2.1
	* Sublist item 2.2
* List item 3`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* * Sublist item 1.1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* List item 2\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	* Sublist item 2.1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	* Sublist item 2.2\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* List item 3"}}},
		{name: "List item with a blank line between the items", input: `+ list item 1

+ list item 2`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "+ list item 1\n",
			},
			&entity.LineElement{Type: entity.ElementKindText, Content: "\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `+ list item 2`}}},
		{name: "List item with a heading ID", input: `- Item with a heading ID {#custom-id}`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a heading ID {#custom-id}`}}},
		{name: "List item with a strikethrough", input: `- Item with a strikethrough ~~strikethrough text~~`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a strikethrough ~~strikethrough text~~`}}},
		{name: "List item with a task list", input: `- Item with a task list
- [x] Task 1
- [ ] Task 2`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Item with a task list\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- [x] Task 1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- [ ] Task 2"}}},
		{name: "List item with an emoji", input: `- Item with an emoji :joy:`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with an emoji :joy:`}}},
		{name: "List item with a highlight", input: `- Item with a highlight ==very important words==`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a highlight ==very important words==`}}},
		{name: "List item with a subscript", input: `- Item with a subscript H~2~O`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a subscript H~2~O`}}},
		{name: "List item with a superscript", input: `- Item with a superscript X^2^`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- Item with a superscript X^2^`}}},
		{name: "List item with a non-breaking space character after it", input: `* &nbsp;
* List item`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* &nbsp;\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* List item"}}},
		{name: "List item with an ordered marker", input: `1. One
2. Two
2.1 Two dot one
3. Three`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "1. One\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "2. Two\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "2.1 Two dot one\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "3. Three"}}},
		{name: "List item with continuous indentation", input: `- One
	- One dot one
- Two
	- Two dot one
- Three
	- Three dot one`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- One\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	- One dot one\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Two\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	- Two dot one\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Three\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	- Three dot one",
			}}},
		{name: "List item with a blank line between the items", input: `+ list item 1

+ list item 2`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "+ list item 1\n"},
			&entity.LineElement{Type: entity.ElementKindText, Content: "\n"},
			&entity.LineElement{Type: entity.ElementKindListItem, Content: `+ list item 2`}}},
		{name: "No blank lines between Markdown list items", input: `* list item 1
* list item 2`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* list item 1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* list item 2",
			}}},
		{name: "List item with a mix of different list markers", input: `- Item 1
+ Item 2
* Item 3`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Item 1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "+ Item 2\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "* Item 3"},
		}},
		{name: "List item with a mix of different list markers in a nested list", input: `- Item 1
+ Item 1.1
	* Item 1.1.1`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Item 1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "+ Item 1.1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	* Item 1.1.1"},
		}},
		{name: "List item with a mix of different list markers in a nested list with multiple levels", input: `- Item 1
+ Item 1.1
	* Item 1.1.1
	+ Item 1.1.2
	- Item 1.1.2.1`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- Item 1\n"},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "+ Item 1.1\n"}, &entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	* Item 1.1.1\n"}, &entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	+ Item 1.1.2\n"}, &entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "	- Item 1.1.2.1",
			}}},
		{name: "Text before the list", input: `Some text before the list.
- List item`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "Some text before the list.\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- List item",
			}}},
		{name: "Text after the list", input: `- List item
Some text after the list.`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- List item\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `Some text after the list.`,
			}}},
		{name: "Text before and after the list", input: `Some text before the list.
- List item
Some text after the list.`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "Some text before the list.\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- List item\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `Some text after the list.`,
			}}},
		{name: "Text and a heading before the list", input: `# Heading
Some text before the list.
- List item`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "# Heading\nSome text before the list.\n",
			},
			&entity.LineElement{Type: entity.ElementKindListItem,
				Content: `- List item`}}},
		{name: "Text and a heading after the list", input: `- List item
Some text after the list.
# Heading`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- List item\n",
			},
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `Some text after the list.
# Heading`,
			}}},
		{name: "Text, a heading, and a paragraph before the list", input: `# Heading
Some text before the list.
This is a paragraph before the list.
- List item`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `# Heading
Some text before the list.
This is a paragraph before the list.` + "\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- List item`,
			}}},
		{name: "Text, a heading, and a paragraph after the list", input: `- List item
Some text after the list.
This is a paragraph after the list.
# Heading`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- List item\n",
			},
			&entity.LineElement{Type: entity.ElementKindText, Content: `Some text after the list.
This is a paragraph after the list.
# Heading`}}},
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
				Type:    entity.ElementKindListItem,
				Content: "- List item\n",
			},
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `Some text after the list.
This is a paragraph after the list.
# Heading`,
			}}},
		{name: "Text, a heading, a paragraph, and a blockquote before the list", input: `# Heading
Some text before the list.
This is a paragraph before the list.
> This is a blockquote before the list.
- List item`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `# Heading
Some text before the list.
This is a paragraph before the list.
> This is a blockquote before the list.` + "\n",
			},
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: `- List item`,
			}}},
		{name: "Text, a heading, a paragraph, and a blockquote after the list", input: `- List item
Some text after the list.
This is a paragraph after the list.
> This is a blockquote after the list.
# Heading`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindListItem,
				Content: "- List item\n",
			},
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `Some text after the list.
This is a paragraph after the list.
> This is a blockquote after the list.
# Heading`,
			}}},
		{name: "Text, a heading, a paragraph, and a blockquote before and after the list", input: `# Heading
Some text before the list.
This is a paragraph before the list.
> This is a blockquote before the list.
- List item
Some text after the list.
This is a paragraph after the list.
> This is a blockquote after the list.
# Heading`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `# Heading
Some text before the list.
This is a paragraph before the list.
> This is a blockquote before the list.` + "\n",
			},
			&entity.LineElement{Type: entity.ElementKindListItem, Content: "- List item\n"},
			&entity.LineElement{Type: entity.ElementKindText,
				Content: `Some text after the list.
This is a paragraph after the list.
> This is a blockquote after the list.
# Heading`}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseLineListItemElement(test.input,
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
