package entity

import "testing"

func TestFootnotemarkdownMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect footnoteMarkdownElement
	}{

		{name: "No content", input: "[^]", expect: footnoteMarkdownElement{Content: ""}},
		{name: "Some content", input: "[^test]", expect: footnoteMarkdownElement{Content: "test"}},
		{name: "Some content different type", input: "[^hello world]", expect: footnoteMarkdownElement{Content: "hello world"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewFootnoteMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
