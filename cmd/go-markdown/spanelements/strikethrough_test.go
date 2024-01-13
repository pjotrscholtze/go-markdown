package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineStrikethroughElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "Single word strikethrough:", input: "~~Hello~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~Hello~~"},
		}},
		{name: "Multiple words strikethrough:", input: "~~Hello World~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~Hello World~~"},
		}},
		{name: "Strikethrough within a sentence:", input: "This is a ~~strikethrough~~ test.", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindText, Content: "This is a "},
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~strikethrough~~"},
			&entity.LineElement{Type: entity.ElementKindText, Content: " test."},
		}},
		{name: "Strikethrough with special characters:", input: "~~Hello, World!~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~Hello, World!~~"},
		}},
		{name: "Strikethrough with numbers:", input: "~~1234567890~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~1234567890~~"},
		}},
		{name: "Strikethrough with trailing space:", input: "~~Hello World ~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~Hello World ~~"},
		}},
		{name: "Strikethrough with leading space:", input: "~~ Hello World~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~ Hello World~~"},
		}},
		{name: "Strikethrough with multiple spaces:", input: "~~Hello    World~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~Hello    World~~"},
		}},
		{name: "Strikethrough with tabs:", input: "~~Hello   World~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindStrikethrough, Content: "~~Hello   World~~"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineStrikethroughElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
