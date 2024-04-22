package entity

import "testing"

func TestEmojiMarkdownElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect emojiMarkdownElement
	}{

		{name: "No content", input: "::", expect: emojiMarkdownElement{Content: ""}},
		{name: "Some smiley", input: ":smile:", expect: emojiMarkdownElement{Content: "smile"}},
		{name: "Another smiley", input: ":sad:", expect: emojiMarkdownElement{Content: "sad"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewEmojiMarkdownElement(test.input,
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
