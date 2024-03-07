package entity

import "testing"

func TestCheckboxMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect checkboxMarkdownElement
	}{

		{name: "Checked checkbox", input: "[x] Hello", expect: checkboxMarkdownElement{Content: "[x] Hello"}},
		{name: "Unchecked checkbox", input: "[ ] Hello", expect: checkboxMarkdownElement{Content: "[ ] Hello"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewCheckboxMarkdownElement(test.input)
			if got.AsMarkdownString() != test.expect.AsMarkdownString() {
				t.Errorf("AsMarkdownString() not the same. Expected %v, got %v", test.expect.AsMarkdownString(), got.AsMarkdownString())
			}
			if got.Kind() != test.expect.Kind() {
				t.Errorf("Kind() not the same. Expected %v, got %v", test.expect.Kind(), got.Kind())
			}
		})
	}
}
