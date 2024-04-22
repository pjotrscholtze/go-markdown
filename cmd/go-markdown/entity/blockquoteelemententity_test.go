package entity

import (
	"testing"
)

func TestBlockQuoteMarkdownElement(t *testing.T) {
	emptyStr := RawTextMarkdownElement("")
	asdfStr := RawTextMarkdownElement(" asdf")
	multipleContentStr := RawTextMarkdownElement("   This is a blockquote with leading and trailing spaces.")
	tests := []struct {
		name   string
		input  string
		expect blockQuoteMarkdownElement
	}{

		{name: "No content", input: ">", expect: blockQuoteMarkdownElement{Content: []MarkdownElement{&emptyStr}}},
		{name: "Content", input: "> asdf", expect: blockQuoteMarkdownElement{Content: []MarkdownElement{&asdfStr}}},
		{name: "Multiple content", input: ">   This is a blockquote with leading and trailing spaces.", expect: blockQuoteMarkdownElement{Content: []MarkdownElement{&multipleContentStr}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewBlockQuoteMarkdownElement(test.input,
				func(input string) []MarkdownElement {
					return []MarkdownElement{&LineElement{
						Type:    ElementKindText,
						Content: input,
					}}
				})
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
