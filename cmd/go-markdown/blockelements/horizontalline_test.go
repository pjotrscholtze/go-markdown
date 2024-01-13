package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestHorizontalLineDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No content", input: "", expect: []entity.MarkdownElement{}},

		{name: "Dash line", input: `---`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHorizontalLine,
				Content: `---`,
			}}},
		{name: "Underscore line", input: `___`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHorizontalLine,
				Content: `___`,
			}}},
		{name: "* line", input: `***`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHorizontalLine,
				Content: `***`,
			}}},

		{name: "Dash line with spaces", input: `- - -`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHorizontalLine,
				Content: `- - -`,
			}}},
		{name: "Underscore line with spaces", input: `_ _ _`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHorizontalLine,
				Content: `_ _ _`,
			}}},
		{name: "* line with spaces", input: `* * *`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHorizontalLine,
				Content: `* * *`,
			}}},

		{name: "Dash line", input: `a
---
b`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `a
`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindHorizontalLine,
				Content: `---`,
			},
			&entity.LineElement{
				Type: entity.ElementKindText,
				Content: `
b`,
			}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineHorizontalLineElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
