package entity

import "testing"

func TestLinkMarkdownElement(t *testing.T) {
	title := RawTextMarkdownElement("example title")
	emptyStr := RawTextMarkdownElement("")
	asdfStr := RawTextMarkdownElement("asdf")
	tests := []struct {
		name   string
		input  string
		expect linkMarkdownElement
	}{
		{name: "No content", input: "[]()", expect: linkMarkdownElement{
			Content: []MarkdownElement{&emptyStr},
			Url:     "",
			Title:   []MarkdownElement{},
		}},
		{name: "With url", input: "[](http://example.com)", expect: linkMarkdownElement{
			Content: []MarkdownElement{&emptyStr},
			Url:     "http://example.com",
			Title:   []MarkdownElement{},
		}},
		{name: "With url and content", input: "[asdf](http://example.com)", expect: linkMarkdownElement{
			Content: []MarkdownElement{&asdfStr},
			Url:     "http://example.com",
			Title:   []MarkdownElement{},
		}},
		{name: "With url, content, and title", input: "[asdf](http://example.com \"example title\")", expect: linkMarkdownElement{
			Content: []MarkdownElement{&asdfStr},
			Url:     "http://example.com",
			Title:   []MarkdownElement{&title},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewLinkMarkdownElement(test.input,
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
			if got.GetContent() != test.expect.GetContent() {
				t.Errorf("GetContent() not the same. Expected %v, got %v", test.expect.GetContent(), got.GetContent())
			}
			if got.GetUrl() != test.expect.GetUrl() {
				t.Errorf("GetUrl() not the same. Expected %v, got %v", test.expect.GetUrl(), got.GetUrl())
			}

			if (got.GetTitle() != nil && test.expect.GetTitle() != nil && (*got.GetTitle()) != (*test.expect.GetTitle())) || (got.GetTitle() == nil && test.expect.GetTitle() != nil) || (got.GetTitle() != nil && test.expect.GetTitle() == nil) {
				if got.GetTitle() != nil && test.expect.GetTitle() != nil {
					t.Errorf("GetTitle() not the same. Expected '%v', got '%v'", *test.expect.GetTitle(), *got.GetTitle())
				} else {
					t.Errorf("GetTitle() not the same. Expected %v, got %v", test.expect.GetTitle(), got.GetTitle())
				}
			}
		})
	}
}
