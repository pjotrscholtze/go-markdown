// listelemententity_test.go
package entity

import (
	"testing"
)

func TestListItemEntityMarkdownElement(t *testing.T) {
	emptyList := RawTextMarkdownElement("- ")
	emptyList2 := RawTextMarkdownElement("+ ")
	helloWorld := RawTextMarkdownElement("+ Hello world")
	helloWorld1 := RawTextMarkdownElement("1. Hello world")
	helloWorld12 := RawTextMarkdownElement("1.2. Hello world")
	multiline := RawTextMarkdownElement("- Hello world\n  more content here")
	multilineNumberStart := RawTextMarkdownElement("1. Hello world\n   more content here")

	tests := []struct {
		name   string
		input  string
		expect listItemElementMarkdownElement
	}{
		{name: "No content", input: "- ", expect: listItemElementMarkdownElement{Content: []MarkdownElement{&emptyList}}},
		{name: "No content different symbol", input: "+ ", expect: listItemElementMarkdownElement{Content: []MarkdownElement{&emptyList2}}},
		{name: "Content", input: "+ Hello world", expect: listItemElementMarkdownElement{Content: []MarkdownElement{&helloWorld}}},
		{name: "Content 1.", input: "1. Hello world", expect: listItemElementMarkdownElement{Content: []MarkdownElement{&helloWorld1}}},
		{name: "Content 1.2.", input: "1.2. Hello world", expect: listItemElementMarkdownElement{Content: []MarkdownElement{&helloWorld12}}},
		{name: "Multi line", input: "- Hello world\n  more content here", expect: listItemElementMarkdownElement{Content: []MarkdownElement{&multiline}}},
		{name: "Multi line number start", input: "1. Hello world\n   more content here", expect: listItemElementMarkdownElement{Content: []MarkdownElement{&multilineNumberStart}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewListItemElementMarkdownElement(test.input,
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
