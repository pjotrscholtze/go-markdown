package entity

import "testing"

func TestParagraphMarkdownElement(t *testing.T) {
	emptyStr := RawTextMarkdownElement("")
	txtStr := RawTextMarkdownElement("lorem ipsum")
	tests := []struct {
		name   string
		input  string
		expect ParagraphMarkdownElement
	}{
		{name: "No content", input: "", expect: &paragraphMarkdownElement{Content: []MarkdownElement{&emptyStr}}},
		{name: "Content", input: "lorem ipsum", expect: &paragraphMarkdownElement{Content: []MarkdownElement{&txtStr}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewParagraphMarkdownElement(test.input,
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
