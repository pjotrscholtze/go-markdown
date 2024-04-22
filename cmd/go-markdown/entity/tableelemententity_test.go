package entity

import (
	"testing"
)

func TestTableEntityMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect tableElementMarkdownElement
	}{
		{name: "No content", input: `| | |
|-|-|
| | |`, expect: tableElementMarkdownElement{Content: `| | |
|-|-|
| | |`}},
		{name: "Content", input: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`, expect: tableElementMarkdownElement{Content: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewTableElementMarkdownElement(test.input,
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
