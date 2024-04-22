package entity

import "testing"

func TestImageMarkdownElement(t *testing.T) {
	emptyStr := RawTextMarkdownElement("")
	spaceStr := RawTextMarkdownElement(" ")
	relativePath := RawTextMarkdownElement("relative path")
	title := RawTextMarkdownElement("example title")
	tests := []struct {
		name   string
		input  string
		expect imageMarkdownElement
	}{

		{name: "No content", input: "![]()", expect: imageMarkdownElement{
			Alt:   []MarkdownElement{&emptyStr},
			Url:   "",
			Title: []MarkdownElement{},
		}},
		{name: "Image with empty alt text", input: "![](image.jpg)", expect: imageMarkdownElement{
			Alt:   []MarkdownElement{&emptyStr},
			Url:   "image.jpg",
			Title: []MarkdownElement{},
		}},
		{name: "Image with space in alt text", input: "![ ](image.jpg)", expect: imageMarkdownElement{
			Alt:   []MarkdownElement{&spaceStr},
			Url:   "image.jpg",
			Title: []MarkdownElement{},
		}},
		{name: "Image with relative path", input: "![relative path](./images/image.jpg)", expect: imageMarkdownElement{
			Alt:   []MarkdownElement{&relativePath},
			Url:   "./images/image.jpg",
			Title: []MarkdownElement{},
		}},
		{name: "Image with relative path", input: "![relative path](./images/image.jpg \"example title\")", expect: imageMarkdownElement{
			Alt:   []MarkdownElement{&relativePath},
			Url:   "./images/image.jpg",
			Title: []MarkdownElement{&title},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewImageMarkdownElement(test.input,
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
			if got.GetAlt() != test.expect.GetAlt() {
				t.Errorf("GetAlt() not the same. Expected %v, got %v", test.expect.GetAlt(), got.GetAlt())
			}
			if !(got.GetTitle() == nil && test.expect.GetTitle() == nil) && *got.GetTitle() != *test.expect.GetTitle() {
				t.Errorf("GetTitle() not the same. Expected %v, got %v", test.expect.GetTitle(), got.GetTitle())
			}
			if got.GetUrl() != test.expect.GetUrl() {
				t.Errorf("GetUrl() not the same. Expected %v, got %v", test.expect.GetUrl(), got.GetUrl())
			}
		})
	}
}
