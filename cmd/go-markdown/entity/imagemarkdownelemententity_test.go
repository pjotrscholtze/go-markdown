package entity

import "testing"

func TestImageMarkdownElement(t *testing.T) {
	title := "example title"
	tests := []struct {
		name   string
		input  string
		expect imageMarkdownElement
	}{

		{name: "No content", input: "![]()", expect: imageMarkdownElement{
			Alt:   "",
			Url:   "",
			Title: nil,
		}},
		{name: "Image with empty alt text", input: "![](image.jpg)", expect: imageMarkdownElement{
			Alt:   "",
			Url:   "image.jpg",
			Title: nil,
		}},
		{name: "Image with space in alt text", input: "![ ](image.jpg)", expect: imageMarkdownElement{
			Alt:   " ",
			Url:   "image.jpg",
			Title: nil,
		}},
		{name: "Image with relative path", input: "![relative path](./images/image.jpg)", expect: imageMarkdownElement{
			Alt:   "relative path",
			Url:   "./images/image.jpg",
			Title: nil,
		}},
		{name: "Image with relative path", input: "![relative path](./images/image.jpg \"example title\")", expect: imageMarkdownElement{
			Alt:   "relative path",
			Url:   "./images/image.jpg",
			Title: &title,
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewImageMarkdownElement(test.input)
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
