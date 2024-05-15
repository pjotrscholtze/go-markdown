package entity

import (
	"testing"
)

func TestTableEntityMarkdownElement(t *testing.T) {
	str := func(in string) MarkdownElement {
		s := RawTextMarkdownElement(in)
		return &s
	}
	tests := []struct {
		name   string
		input  string
		expect tableElementMarkdownElement
	}{
		{name: "No content", input: `| | |
|-|-|
| | |`, expect: tableElementMarkdownElement{
			header: TableRow{
				Cells: []MarkdownElement{str(""), str(" "), str(" "), str("\n")},
			},
			sep: TableRow{Cells: []MarkdownElement{
				str(""),
				str("-"),
				str("-"),
				str("\n"),
			}},
			rows: []TableRow{
				{
					Cells: []MarkdownElement{str(""), str(" "), str(" "), str("")},
				}}}},
		{name: "Content", input: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`, expect: tableElementMarkdownElement{
			header: TableRow{
				Cells: []MarkdownElement{str(""), str(" Column 1 "), str(" Column 2 "), str("\n")},
			},
			sep: TableRow{Cells: []MarkdownElement{
				str(""),
				str("----------"),
				str("----------"),
				str("\n"),
			}},
			rows: []TableRow{
				{
					Cells: []MarkdownElement{str(""), str(" A "), str(" B "), str("")},
				}}}},
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
