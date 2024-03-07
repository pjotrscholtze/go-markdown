package entity

import (
	"testing"
)

func TestBoldMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect boldMarkdownElement
	}{

		{name: "No content", input: "____", expect: boldMarkdownElement{Content: "", WrappingSymbol: '_'}},
		{name: "Content", input: "__a__", expect: boldMarkdownElement{Content: "a", WrappingSymbol: '_'}},
		{name: "Multiple content", input: "__asdf asdf hello world__", expect: boldMarkdownElement{Content: "asdf asdf hello world", WrappingSymbol: '_'}},
		{name: "With spaces surrounded content", input: "__ asdf asdf hello world __", expect: boldMarkdownElement{Content: " asdf asdf hello world ", WrappingSymbol: '_'}},
		{name: "Only a space", input: "__ __", expect: boldMarkdownElement{Content: " ", WrappingSymbol: '_'}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewBoldMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
			if got.GetWrappingSymbolAsRune() != test.expect.GetWrappingSymbolAsRune() {
				t.Errorf("GetWrappingSymbolAsRune() not the same. Expected %v, got %v", test.expect.GetWrappingSymbolAsRune(), got.GetWrappingSymbolAsRune())
			}
		})
	}
}
