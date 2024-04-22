package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestBlockquoteLineDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No content", input: "", expect: []entity.MarkdownElement{}},
		{name: "Single quote line", input: ">   This is a blockquote with leading and trailing spaces.", expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindBlockquote,
				Content: ">   This is a blockquote with leading and trailing spaces.",
			},
		}},
		{name: "Multi line quote", input: `>
>This is a blockquote with leading and trailing newlines.
>`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindBlockquote,
				Content: `>
>This is a blockquote with leading and trailing newlines.
>`,
			},
		}},
		{name: "Nested quote", input: `> Outer blockquote
>> Inner blockquote`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindBlockquote,
				Content: `> Outer blockquote
>> Inner blockquote`,
			},
		}},
		{name: "Quote surrounded with text", input: `Title here
> This is a blockquote.
Next line after blockquote.`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `Title here`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindBlockquote,
				Content: `> This is a blockquote.`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `Next line after blockquote.`,
			},
		}},
		{name: "Quote with special characters", input: `> This is a blockquote with special characters: !"#$%&'()*+,-./:;<=>?@[\]^_{|}~`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindBlockquote,
				Content: `> This is a blockquote with special characters: !"#$%&'()*+,-./:;<=>?@[\]^_{|}~`,
			},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseLineBlockquoteElement([]entity.MarkdownElement{
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
