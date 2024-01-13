package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineFootnoteElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{
		{
			name:  "Footnote at the beginning of a line:",
			input: "[^1] This is a footnote at the beginning of a line.",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^1]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: " This is a footnote at the beginning of a line.",
				},
			},
		},
		{
			name:  "Footnote within a sentence:",
			input: "This is a sentence with a footnote[^2] within it.",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This is a sentence with a footnote",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^2]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: " within it.",
				},
			},
		},
		{
			name:  "Footnote at the end of a paragraph:",
			input: "This is a paragraph with a footnote at the end[^3]",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This is a paragraph with a footnote at the end",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^3]",
				},
			},
		},
		{
			name:  "Multiple footnotes in one paragraph:",
			input: "This paragraph has multiple footnotes[^4], including one with a space in the label[^ 5 ].",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This paragraph has multiple footnotes",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^4]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: ", including one with a space in the label",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^ 5 ]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: ".",
				},
			},
		},
		{
			name:  "Nested footnotes:",
			input: "This footnote[^6] contains another footnote[^7].",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^6]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: " contains another footnote",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^7]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: ".",
				},
			},
		},
		{
			name:  "Footnote with special characters:",
			input: "This footnote[^8] contains special characters like @, #, $, %, ^, &, *, (, ), _, +, `, ~, -, =, {, }, [, ], |, :, ;, ', \", ,, ., /, ?, !, and .",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^8]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: " contains special characters like @, #, $, %, ^, &, *, (, ), _, +, `, ~, -, =, {, }, [, ], |, :, ;, ', \", ,, ., /, ?, !, and .",
				},
			},
		},
		{
			name:  "Footnote with HTML tags:",
			input: "This footnote[^9] contains HTML tags like `<b>`bold`</b>`, `<i>`italic`</i>`, and `<a href=\"http://example.com\">link</a>`.",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^9]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: " contains HTML tags like `<b>`bold`</b>`, `<i>`italic`</i>`, and `<a href=\"http://example.com\">link</a>`.",
				},
			},
		},
		{
			name:  "Footnote with Unicode characters:",
			input: "This footnote[^10] contains Unicode characters like ä, ö, ü, ß, é, è, ê, à",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^10]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: " contains Unicode characters like ä, ö, ü, ß, é, è, ê, à",
				},
			},
		},
		{
			name:  "Text footnote:",
			input: "This footnote[^hoi] contains Unicode characters like ä, ö, ü, ß, é, è, ê, à",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: "This footnote",
				},
				&entity.LineElement{
					Type:    entity.ElementKindFootnote,
					Content: "[^hoi]",
				},
				&entity.LineElement{
					Type:    entity.ElementKindText,
					Content: " contains Unicode characters like ä, ö, ü, ß, é, è, ê, à",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineFootnoteElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
