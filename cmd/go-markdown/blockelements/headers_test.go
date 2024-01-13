package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestHeadingLineDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.LineElement
	}{

		{name: "No content", input: "", expect: []entity.LineElement{}},
		{name: "Simple content", input: `Not a heading
# Heading with *emphasis*
## Heading with **strong emphasis**
### Heading with ~~strikethrough~~
#### Heading with ` + "`code`" + `
##### Heading with [link](https://example.com)
###### Heading 6
####### No longer a heading
Not a heading #afterwards`, expect: []entity.LineElement{
			{
				Type:    entity.ElementKindText,
				Content: "Not a heading",
			},
			{Type: entity.ElementKindHeader, Content: `# Heading with *emphasis*`},
			{Type: entity.ElementKindHeader, Content: `## Heading with **strong emphasis**`},
			{Type: entity.ElementKindHeader, Content: `### Heading with ~~strikethrough~~`},
			{Type: entity.ElementKindHeader, Content: `#### Heading with ` + "`code`" + ``},
			{Type: entity.ElementKindHeader, Content: `##### Heading with [link](https://example.com)`},
			{Type: entity.ElementKindHeader, Content: `###### Heading 6`},
			{Type: entity.ElementKindText,
				Content: `####### No longer a heading
Not a heading #afterwards`,
			},
		}},
		{name: "Spacing", input: `# Normal spacing
#No spacing at all
#      Extra spacing
# Trailing spacing      `, expect: []entity.LineElement{
			{Type: entity.ElementKindHeader, Content: `# Normal spacing`},
			{Type: entity.ElementKindHeader, Content: `#No spacing at all`},
			{Type: entity.ElementKindHeader, Content: `#      Extra spacing`},
			{Type: entity.ElementKindHeader, Content: `# Trailing spacing      `},
		}},
		{name: "Long heading", input: `# A very long heading that goes beyond the usual limit for a heading in Markdown`, expect: []entity.LineElement{
			{Type: entity.ElementKindHeader, Content: `# A very long heading that goes beyond the usual limit for a heading in Markdown`},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineHeaderElement([]entity.LineElement{{Type: "text", Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
