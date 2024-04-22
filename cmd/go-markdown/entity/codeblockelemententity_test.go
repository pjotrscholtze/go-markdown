package entity

import (
	"strings"
	"testing"
)

func TestCodeBlockMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect codeBlockMarkdownElement
	}{

		{name: "No content", input: "``````", expect: codeBlockMarkdownElement{Content: ""}},
		{name: "Content", input: "```asdf```", expect: codeBlockMarkdownElement{Content: "asdf"}},
		{name: "Multiple content", input: "```This is a blockquote with leading and trailing spaces.```", expect: codeBlockMarkdownElement{Content: "This is a blockquote with leading and trailing spaces."}},
		{name: "Multiline content", input: strings.Join([]string{
			"```",
			"test",
			"asdf",
			"```",
		}, "\n"), expect: codeBlockMarkdownElement{Content: strings.Join([]string{
			"",
			"test",
			"asdf",
			"",
		}, "\n")}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewCodeBlockMarkdownElement(test.input,
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
