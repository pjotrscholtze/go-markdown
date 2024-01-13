package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseInlineCodeElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{
		{name: "No content", input: "", expect: []entity.MarkdownElement{}},
		{name: "Single word code", input: "`ipsum`", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindInlineCode,
				Content: "`ipsum`",
			}}},
		{name: "Single word code, with text surrounding", input: "lorem `ipsum` dolar", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "lorem ",
			},
			&entity.LineElement{
				Type:    "inlinecode",
				Content: "`ipsum`",
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: " dolar",
			}}},
		{name: "Multi word code, with text surrounding", input: "lorem `ipsum dolar` sit", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "lorem ",
			},
			&entity.LineElement{
				Type:    "inlinecode",
				Content: "`ipsum dolar`",
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: " sit",
			}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseInlineCodeElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
