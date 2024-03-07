package entity

import "testing"

func TestLinkMarkdownElement(t *testing.T) {
	title := "example title"
	tests := []struct {
		name   string
		input  string
		expect linkMarkdownElement
	}{
		{name: "No content", input: "[]()", expect: linkMarkdownElement{
			Content: "",
			Url:     "",
			Title:   nil,
		}},
		{name: "With url", input: "[](http://example.com)", expect: linkMarkdownElement{
			Content: "",
			Url:     "http://example.com",
			Title:   nil,
		}},
		{name: "With url and content", input: "[asdf](http://example.com)", expect: linkMarkdownElement{
			Content: "asdf",
			Url:     "http://example.com",
			Title:   nil,
		}},
		{name: "With url, content, and title", input: "[asdf](http://example.com \"example title\")", expect: linkMarkdownElement{
			Content: "asdf",
			Url:     "http://example.com",
			Title:   &title,
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewLinkMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
			if got.GetContent() != test.expect.GetContent() {
				t.Errorf("GetContent() not the same. Expected %v, got %v", test.expect.GetContent(), got.GetContent())
			}
			if got.GetUrl() != test.expect.GetUrl() {
				t.Errorf("GetUrl() not the same. Expected %v, got %v", test.expect.GetUrl(), got.GetUrl())
			}
			if !(got.GetTitle() == nil && test.expect.GetTitle() == nil) && *got.GetTitle() != *test.expect.GetTitle() {
				t.Errorf("GetTitle() not the same. Expected %v, got %v", test.expect.GetTitle(), got.GetTitle())
			}
		})
	}
}
