package entity

import (
	"testing"
)

func TestTermDefinitionEntityMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect termDefinitionElementMarkdownElement
	}{
		{name: "No content", input: `^: `, expect: termDefinitionElementMarkdownElement{Content: ""}},
		{name: "Content", input: `^: Markdown: is a markup language`, expect: termDefinitionElementMarkdownElement{Content: "Markdown: is a markup language"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewTermDefinitionElementMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
