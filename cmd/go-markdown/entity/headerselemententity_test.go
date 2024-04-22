package entity

import (
	"testing"
)

func TestHeaderMarkdownElement(t *testing.T) {
	asdfStr := NewRawTextMarkdownElement(" asdf",
		func(input string) []MarkdownElement {
			return []MarkdownElement{&LineElement{
				Type:    ElementKindText,
				Content: input,
			}}
		})
	txtStr := NewRawTextMarkdownElement(" This is a blockquote with leading and trailing spaces",
		func(input string) []MarkdownElement {
			return []MarkdownElement{&LineElement{
				Type:    ElementKindText,
				Content: input,
			}}
		})
	tests := []struct {
		name   string
		input  string
		expect headerMarkdownElement
	}{

		{name: "No content", input: "#", expect: headerMarkdownElement{Content: []MarkdownElement{}, HeadingLevel: 1}},
		{name: "Content", input: "# asdf", expect: headerMarkdownElement{Content: []MarkdownElement{&asdfStr}, HeadingLevel: 1}},
		{name: "Multiple content", input: "# This is a blockquote with leading and trailing spaces", expect: headerMarkdownElement{Content: []MarkdownElement{&txtStr}, HeadingLevel: 1}},
		{name: "H2 Content", input: "## asdf", expect: headerMarkdownElement{Content: []MarkdownElement{&asdfStr}, HeadingLevel: 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewHeaderMarkdownElement(test.input, func(input string) []MarkdownElement {
				md := NewRawTextMarkdownElement(input,
					func(input string) []MarkdownElement {
						return []MarkdownElement{&LineElement{
							Type:    ElementKindText,
							Content: input,
						}}
					})
				return []MarkdownElement{&md}
			})
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
			if got.GetHeadingLevel() != test.expect.GetHeadingLevel() {
				t.Errorf("GetHeadingLevel() not the same. Expected %v, got %v", test.expect.GetHeadingLevel(), got.GetHeadingLevel())
			}
		})
	}
}
