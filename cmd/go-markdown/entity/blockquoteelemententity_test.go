package entity

import (
	"testing"
)

func TestBlockQuoteMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect blockQuoteMarkdownElement
	}{

		{name: "No content", input: ">", expect: blockQuoteMarkdownElement{Content: ""}},
		{name: "Content", input: "> asdf", expect: blockQuoteMarkdownElement{Content: " asdf"}},
		{name: "Multiple content", input: ">   This is a blockquote with leading and trailing spaces.", expect: blockQuoteMarkdownElement{Content: "   This is a blockquote with leading and trailing spaces."}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewBlockQuoteMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
