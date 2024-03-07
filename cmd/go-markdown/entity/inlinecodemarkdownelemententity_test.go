package entity

import "testing"

func TestInlineCodeMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect inlineCodeMarkdownElement
	}{
		{name: "No content", input: "``", expect: inlineCodeMarkdownElement{Content: ""}},
		{name: "Single word code", input: "`ipsum`", expect: inlineCodeMarkdownElement{Content: "ipsum"}},
		{name: "Multi word code", input: "`ipsum ipsum`", expect: inlineCodeMarkdownElement{Content: "ipsum ipsum"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewInlineCodeMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
