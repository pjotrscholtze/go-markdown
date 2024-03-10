package entity

import (
	"testing"
)

func TestHeaderMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect headerMarkdownElement
	}{

		{name: "No content", input: "#", expect: headerMarkdownElement{Content: "", HeadingLevel: 1}},
		{name: "Content", input: "# asdf", expect: headerMarkdownElement{Content: " asdf", HeadingLevel: 1}},
		{name: "Multiple content", input: "# This is a blockquote with leading and trailing spaces", expect: headerMarkdownElement{Content: " This is a blockquote with leading and trailing spaces", HeadingLevel: 1}},
		{name: "H2 Content", input: "## asdf", expect: headerMarkdownElement{Content: " asdf", HeadingLevel: 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewHeaderMarkdownElement(test.input)
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
