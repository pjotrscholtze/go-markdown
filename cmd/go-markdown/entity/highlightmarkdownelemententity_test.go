package entity

import "testing"

func TestHighlightMarkdownElement(t *testing.T) {
	emptyStr := RawTextMarkdownElement("")
	testStr := RawTextMarkdownElement("test")
	tests := []struct {
		name   string
		input  string
		expect highlightMarkdownElement
	}{

		{name: "No content", input: "====", expect: highlightMarkdownElement{Content: []MarkdownElement{&emptyStr}}},
		{name: "Some content", input: "==test==", expect: highlightMarkdownElement{Content: []MarkdownElement{&testStr}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewHighlightMarkdownElement(test.input,
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
