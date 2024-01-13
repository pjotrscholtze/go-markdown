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
		expect []entity.LineElement
	}{

		{name: "No content", input: "", expect: []entity.LineElement{}},
		{name: "Simple example", input: `^: Markdown: is a markup language`, expect: []entity.LineElement{
			{Type: entity.ElementKindTermDefinitionLine, Content: "^: Markdown: is a markup language"},
		}},
		{name: "Surrounding text example", input: strings.Join([]string{
			"before",
			"^: Markdown: is a markup language",
			"after",
		}, "\n"), expect: []entity.LineElement{
			{Type: entity.ElementKindText, Content: "before\n"},
			{Type: entity.ElementKindTermDefinitionLine, Content: "^: Markdown: is a markup language"},
			{Type: entity.ElementKindText, Content: "\nafter"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineTermDefinitionLineElement([]entity.LineElement{{Type: "text", Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}

func equalResults(a, b []entity.LineElement) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
