package entity

import "testing"

func TestStrikethroughMarkdownElement(t *testing.T) {
	content := RawTextMarkdownElement("lorem ipsum")
	tests := []struct {
		name   string
		input  string
		expect strikethroughMarkdownElement
	}{
		{name: "No content", input: "~~~~", expect: strikethroughMarkdownElement{Content: []MarkdownElement{}}},
		{name: "Content", input: "~~lorem ipsum~~", expect: strikethroughMarkdownElement{Content: []MarkdownElement{&content}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewStrikethroughMarkdownElement(test.input,
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
