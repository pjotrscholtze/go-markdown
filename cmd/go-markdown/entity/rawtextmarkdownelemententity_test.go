package entity

import "testing"

func TestRawtextMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect RawTextMarkdownElement
	}{
		{name: "No content", input: "", expect: RawTextMarkdownElement("")},
		{name: "Content", input: "lorem ipsum", expect: RawTextMarkdownElement("lorem ipsum")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewRawTextMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
