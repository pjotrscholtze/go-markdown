package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestHeadingLineDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No content", input: "", expect: []entity.MarkdownElement{}},
		{name: "Simple content", input: `Not a heading
# Heading with *emphasis*
## Heading with **strong emphasis**
### Heading with ~~strikethrough~~
#### Heading with ` + "`code`" + `
##### Heading with [link](https://example.com)
###### Heading 6
####### No longer a heading
Not a heading #afterwards`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "Not a heading",
			},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `# Heading with *emphasis*`},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `## Heading with **strong emphasis**`},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `### Heading with ~~strikethrough~~`},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `#### Heading with ` + "`code`" + ``},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `##### Heading with [link](https://example.com)`},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `###### Heading 6`},
			&entity.LineElement{Type: entity.ElementKindText,
				Content: `####### No longer a heading
Not a heading #afterwards`,
			},
		}},
		{name: "Spacing", input: `# Normal spacing
#No spacing at all
#      Extra spacing
# Trailing spacing      `, expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `# Normal spacing`},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `#No spacing at all`},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `#      Extra spacing`},
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `# Trailing spacing      `},
		}},
		{name: "Long heading", input: `# A very long heading that goes beyond the usual limit for a heading in Markdown`, expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindHeader, Content: `# A very long heading that goes beyond the usual limit for a heading in Markdown`},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseLineHeaderElement([]entity.MarkdownElement{
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
