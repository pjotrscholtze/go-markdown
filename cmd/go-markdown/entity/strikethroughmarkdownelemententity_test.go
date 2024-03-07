package entity

import "testing"

func TestStrikethroughMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect strikethroughMarkdownElement
	}{
		{name: "No content", input: "~~~~", expect: strikethroughMarkdownElement{Content: ""}},
		{name: "Content", input: "~~lorem ipsum~~", expect: strikethroughMarkdownElement{Content: "lorem ipsum"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewStrikethroughMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
