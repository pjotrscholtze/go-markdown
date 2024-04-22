package blockelements

import (
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestTermDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No content", input: "", expect: []entity.MarkdownElement{}},
		{name: "Simple example", input: `^: Markdown: is a markup language`, expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindTermDefinitionLine, Content: "^: Markdown: is a markup language"},
		}},
		{name: "Surrounding text example", input: strings.Join([]string{
			"before",
			"^: Markdown: is a markup language",
			"after",
		}, "\n"), expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindText, Content: "before\n"},
			entity.NewTermDefinitionElementMarkdownElement("^: Markdown: is a markup language",
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				}),
			&entity.LineElement{Type: entity.ElementKindText, Content: "\nafter"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseLineTermDefinitionLineElement([]entity.MarkdownElement{
				&entity.LineElement{Type: "text", Content: test.input}},
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

func equalResults(a, b []entity.MarkdownElement) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Kind() != b[i].Kind() {
			return false
		}
		if v.AsMarkdownString() != b[i].AsMarkdownString() {
			return false
		}
	}
	return true
}
