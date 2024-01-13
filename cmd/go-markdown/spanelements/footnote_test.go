package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineFootnoteElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.LineElement
	}{
		{
			name:  "Footnote at the beginning of a line:",
			input: "[^1] This is a footnote at the beginning of a line.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^1]",
				},
				{
					Type:    entity.ElementKindText,
					Content: " This is a footnote at the beginning of a line.",
				},
			},
		},
		{
			name:  "Footnote within a sentence:",
			input: "This is a sentence with a footnote[^2] within it.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This is a sentence with a footnote",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^2]",
				},
				{
					Type:    entity.ElementKindText,
					Content: " within it.",
				},
			},
		},
		{
			name:  "Footnote at the end of a paragraph:",
			input: "This is a paragraph with a footnote at the end[^3]",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This is a paragraph with a footnote at the end",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^3]",
				},
			},
		},
		{
			name:  "Multiple footnotes in one paragraph:",
			input: "This paragraph has multiple footnotes[^4], including one with a space in the label[^ 5 ].",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This paragraph has multiple footnotes",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^4]",
				},
				{
					Type:    entity.ElementKindText,
					Content: ", including one with a space in the label",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^ 5 ]",
				},
				{
					Type:    entity.ElementKindText,
					Content: ".",
				},
			},
		},
		{
			name:  "Nested footnotes:",
			input: "This footnote[^6] contains another footnote[^7].",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^6]",
				},
				{
					Type:    entity.ElementKindText,
					Content: " contains another footnote",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^7]",
				},
				{
					Type:    entity.ElementKindText,
					Content: ".",
				},
			},
		},
		{
			name:  "Footnote with special characters:",
			input: "This footnote[^8] contains special characters like @, #, $, %, ^, &, *, (, ), _, +, `, ~, -, =, {, }, [, ], |, :, ;, ', \", ,, ., /, ?, !, and .",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^8]",
				},
				{
					Type:    entity.ElementKindText,
					Content: " contains special characters like @, #, $, %, ^, &, *, (, ), _, +, `, ~, -, =, {, }, [, ], |, :, ;, ', \", ,, ., /, ?, !, and .",
				},
			},
		},
		{
			name:  "Footnote with HTML tags:",
			input: "This footnote[^9] contains HTML tags like `<b>`bold`</b>`, `<i>`italic`</i>`, and `<a href=\"http://example.com\">link</a>`.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^9]",
				},
				{
					Type:    entity.ElementKindText,
					Content: " contains HTML tags like `<b>`bold`</b>`, `<i>`italic`</i>`, and `<a href=\"http://example.com\">link</a>`.",
				},
			},
		},
		{
			name:  "Footnote with Unicode characters:",
			input: "This footnote[^10] contains Unicode characters like ä, ö, ü, ß, é, è, ê, à",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^10]",
				},
				{
					Type:    entity.ElementKindText,
					Content: " contains Unicode characters like ä, ö, ü, ß, é, è, ê, à",
				},
			},
		},
		{
			name:  "Text footnote:",
			input: "This footnote[^hoi] contains Unicode characters like ä, ö, ü, ß, é, è, ê, à",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				{
					Type:    entity.ElementKindFootnote,
					Content: "[^hoi]",
				},
				{
					Type:    entity.ElementKindText,
					Content: " contains Unicode characters like ä, ö, ü, ß, é, è, ê, à",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineFootnoteElement([]entity.LineElement{{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
