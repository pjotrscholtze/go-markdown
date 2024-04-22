package entity

import "testing"

func TestItalicMarkdownElement(t *testing.T) {
	emptyStr := RawTextMarkdownElement("")
	ipsumStr := RawTextMarkdownElement("ipsum")
	ipsumIpsumStr := RawTextMarkdownElement("ipsum ipsum")
	tests := []struct {
		name   string
		input  string
		expect italicMarkdownElement
	}{
		{name: "No content", input: "__", expect: italicMarkdownElement{Content: []MarkdownElement{&emptyStr}, WrappingSymbol: '_'}},
		{name: "Single word code", input: "_ipsum_", expect: italicMarkdownElement{Content: []MarkdownElement{&ipsumStr}, WrappingSymbol: '_'}},
		{name: "Multi word code", input: "_ipsum ipsum_", expect: italicMarkdownElement{Content: []MarkdownElement{&ipsumIpsumStr}, WrappingSymbol: '_'}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewItalicMarkdownElement(test.input,
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
