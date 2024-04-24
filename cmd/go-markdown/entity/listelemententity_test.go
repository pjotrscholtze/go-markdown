package entity

import (
	"testing"
)

func TestListEntityMarkdownElement(t *testing.T) {
	emptyList := RawTextMarkdownElement("- ")
	emptyList2 := RawTextMarkdownElement("+ ")
	helloWorld := RawTextMarkdownElement("+ Hello world")
	// 	helloWorldLoremIpsum := RawTextMarkdownElement(`- Hello world
	// - Lorem ipsum`)

	tests := []struct {
		name   string
		input  string
		expect listElementMarkdownElement
	}{
		{name: "No content", input: "- ", expect: listElementMarkdownElement{Content: []MarkdownElement{&emptyList}}},
		{name: "No content different symbol", input: "+ ", expect: listElementMarkdownElement{Content: []MarkdownElement{&emptyList2}}},
		{name: "Content", input: "+ Hello world", expect: listElementMarkdownElement{Content: []MarkdownElement{&helloWorld}}},

		// 		{name: "ContentMultiline", input: `- Hello world
		//   Complicated
		//   - testing abcd
		//   - testing 1234
		//   With different lists
		//   - testing world
		//   Not ending on a list
		// - Lorem ipsum
		// - Content here
		// - Hello world
		//   - testing 1234
		// - Lorem ipsum`, expect: listElementMarkdownElement{Content: []MarkdownElement{
		// 			&subListElementMarkdownElement{
		// 				MainContent:   []MarkdownElement{&helloWorldLoremIpsum},
		// 				SubListOffset: "  ",
		// 				Symbol:        "- ",
		// 				SubList:       []MarkdownElement{},
		// 			},
		// 		}}},
		// 		{name: "ContentMultiline", input: `- Hello world
		//   - testing 1234
		//   - testing abcd
		//   - testing hello
		//   - testing world
		// - Lorem ipsum
		// - Hello world
		//   - testing 1234
		// - Lorem ipsum`, expect: listElementMarkdownElement{Content: []MarkdownElement{&helloWorldLoremIpsum}}},
		// 		{name: "ContentMultiline", input: `- Hello world
		// testing 1234
		// - Lorem ipsum`, expect: listElementMarkdownElement{Content: []MarkdownElement{&helloWorldLoremIpsum}}},
		// 		{name: "ContentMultiline", input: `- Hello world
		// - Lorem ipsum`, expect: listElementMarkdownElement{Content: []MarkdownElement{&helloWorldLoremIpsum}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewListElementMarkdownElement(test.input,
				func(input string) []MarkdownElement {
					return []MarkdownElement{&LineElement{
						Type:    ElementKindText,
						Content: input,
					}}
				},
				func(input []MarkdownElement, parserFn func(input string) []MarkdownElement) []MarkdownElement {
					content := []MarkdownElement{}
					for _, entry := range input {
						content = append(content, parserFn(entry.AsMarkdownString())...)
					}
					return content
				},
			)
			gmd := got.AsMarkdownString()
			tmd := test.expect.AsMarkdownString()
			_ = gmd
			_ = tmd
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
