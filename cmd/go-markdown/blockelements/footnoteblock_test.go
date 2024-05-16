package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseFootnoteDefinitionElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{
		{
			name:  "Footnote at the beginning of a line:",
			input: "[^1]: This is a footnote at the beginning of a line.",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindFootnoteDefinition,
					Content: "[^1]: ",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This is a footnote at the beginning of a line.",
				},
			},
		},
		{
			name:  "Multiple footnotes at the beginning of a line:",
			input: "[^1]: This is a footnote at the beginning of a line.\n[^testing]: This is a footnote at the beginning of a line too.",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindFootnoteDefinition,
					Content: "[^1]: ",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This is a footnote at the beginning of a line.\n",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnoteDefinition,
					Content: "[^testing]: ",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This is a footnote at the beginning of a line too.",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseFootnoteDefinitionElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}},
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				})
			if entity.GlueToString(got) != entity.GlueToString(test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
