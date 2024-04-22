package entity

import (
	"testing"
)

func TestHorizontalLineMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect horizontalLineMarkdownElement
	}{
		{name: "Line", input: "--", expect: horizontalLineMarkdownElement{Content: "--"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewHorizontalLineMarkdownElement(test.input,
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
