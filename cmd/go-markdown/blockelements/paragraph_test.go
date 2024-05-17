package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParagraph(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No content", input: "", expect: []entity.MarkdownElement{}},

		{name: "Content", input: `blah`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindParagraph,
				Content: `blah`,
			}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseParagraphElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}},
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				})
			gmd := entity.GlueToString(got)
			emd := entity.GlueToString(test.expect)
			_ = gmd
			_ = emd
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
