package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineHighlightElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{
		{name: "Single word highlight:", input: "==test==", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHighlight,
				Content: "==test==",
			}}},
		{name: "Multiple words highlight:", input: "==this is a test==", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHighlight,
				Content: "==this is a test==",
			}}},
		{name: "Highlight within a sentence:", input: "This is a ==test== sentence.", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "This is a ",
			},
			&entity.LineElement{
				Type:    entity.ElementKindHighlight,
				Content: "==test==",
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: " sentence.",
			},
		}},
		{name: "Highlight without opening ==:", input: "test==", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "test==",
			}}},
		{name: "Highlight without closing ==:", input: "==test", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: "==test",
			}}},
		{name: "Highlight with leading and trailing spaces:", input: "== test ==", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindHighlight,
				Content: "== test ==",
			}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineHighlightElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
