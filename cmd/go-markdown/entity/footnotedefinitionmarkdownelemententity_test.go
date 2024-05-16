package entity

import "testing"

func TestFootnoteDefinitionmarkdownMarkdownElement(t *testing.T) {
	emptyStr := RawTextMarkdownElement("")
	testStr := RawTextMarkdownElement("test")
	helloWorld := RawTextMarkdownElement("hello world")
	tests := []struct {
		name   string
		input  string
		expect footnoteDefinitionMarkdownElement
	}{

		{name: "No content", input: "[^]: ", expect: footnoteDefinitionMarkdownElement{Content: []MarkdownElement{&emptyStr}}},
		{name: "Some content", input: "[^test]: ", expect: footnoteDefinitionMarkdownElement{Content: []MarkdownElement{&testStr}}},
		{name: "Some content different type", input: "[^hello world]: ", expect: footnoteDefinitionMarkdownElement{Content: []MarkdownElement{&helloWorld}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewFootnoteDefinitionMarkdownElement(test.input,
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
