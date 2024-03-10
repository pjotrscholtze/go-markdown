package entity

import (
	"testing"
)

func TestListEntityMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect listElementMarkdownElement
	}{
		{name: "No content", input: "- ", expect: listElementMarkdownElement{Content: "- "}},
		{name: "No content different symbol", input: "+ ", expect: listElementMarkdownElement{Content: "+ "}},
		{name: "Content", input: "+ Hello world", expect: listElementMarkdownElement{Content: "+ Hello world"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewListElementMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
