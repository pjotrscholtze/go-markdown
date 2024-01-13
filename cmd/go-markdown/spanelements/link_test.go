package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineLinkElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.LineElement
	}{
		{
			name:  "Empty link:",
			input: "[](http://example.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindLink,
				Content: "[](http://example.com)",
			},
			},
		},
		{
			name:  "No link:",
			input: "[]()",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindLink,
				Content: "[]()",
			},
			},
		},
		{
			name:  "Empty link and target:",
			input: "[]()",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindLink,
				Content: "[]()",
			},
			},
		},
		{
			name:  "Whitespace link:",
			input: "[ ]",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindText,
				Content: "[ ]",
			},
			},
		},
		{
			name:  "Whitespace target:",
			input: "[label]( )",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindLink,
				Content: "[label]( )",
			},
			},
		},
		{
			name:  "Non-ASCII characters in link label:",
			input: "[ラベル](http://example.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindLink,
				Content: "[ラベル](http://example.com)",
			},
			},
		},
		{
			name:  "Non-ASCII characters in link target:",
			input: "[label](http://例子.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindLink,
				Content: "[label](http://例子.com)",
			},
			},
		},
		{
			name:  "Special characters in link label:",
			input: "[*label*](http://example.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindLink,
				Content: "[*label*](http://example.com)",
			},
			},
		},
		{
			name:  "Special characters in link target:",
			input: "[label](http://ex*ample*.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindLink,
				Content: "[label](http://ex*ample*.com)",
			},
			},
		},
		{
			//
			name:  "Link surrounded by text:",
			input: " Visit [Google](http://google.com) for searching the internet.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: " Visit ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[Google](http://google.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for searching the internet.",
				},
			},
		},
		{
			name:  "Link at the start of a sentence:",
			input: "[Google](http://google.com) is a well-known search engine.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindLink,
					Content: "[Google](http://google.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " is a well-known search engine.",
				},
			},
		},
		{
			name:  "Link in the middle of a sentence:",
			input: " Search engines like [Google](http://google.com) are essential tools.",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindText,
				Content: " Search engines like ",
			},
				{
					Type:    entity.ElementKindLink,
					Content: "[Google](http://google.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " are essential tools.",
				},
			},
		},
		{
			name:  "Links with different capitalization:",
			input: " Use [Google](http://google.com) for searching, or [BING](http://bing.com) for shopping.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: " Use ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[Google](http://google.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for searching, or ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[BING](http://bing.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for shopping.",
				},
			},
		},
		{
			name:  "Links with special characters in the label:",
			input: " Use [Google](http://google.com) for searching, or [Bing!](http://bing.com) for fun.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: " Use ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[Google](http://google.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for searching, or ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[Bing!](http://bing.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for fun.",
				},
			},
		},
		{
			name:  "Links with special characters in the target:",
			input: " Use [Google](http://google.com) for searching, or [Bing](http://bing!.com) for fun.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: " Use ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[Google](http://google.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for searching, or ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[Bing](http://bing!.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for fun.",
				},
			},
		},
		{
			name:  "Links with non-ASCII characters in the label and target:",
			input: " Use [Google](http://google.com) for searching, or [Bing](http://bing.com) for shopping.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: " Use ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[Google](http://google.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for searching, or ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[Bing](http://bing.com)",
				},
				{
					Type:    entity.ElementKindText,
					Content: " for shopping.",
				},
			},
		},
		{
			name:  "Link inside a heading:",
			input: " # [Google](http://google.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindText,
				Content: " # ",
			}, {
				Type:    entity.ElementKindLink,
				Content: "[Google](http://google.com)",
			},
			},
		},
		{
			name:  "Link inside a blockquote:",
			input: " > Visit [Google](http://google.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindText,
				Content: " > Visit ",
			}, {
				Type:    entity.ElementKindLink,
				Content: "[Google](http://google.com)",
			},
			},
		},
		{
			name:  "Link inside a task list:",
			input: " - [x] Visit [Google](http://google.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindText,
				Content: " - [x] Visit ",
			}, {
				Type:    entity.ElementKindLink,
				Content: "[Google](http://google.com)",
			},
			},
		},
		{
			name:  "Link inside a strikethrough:",
			input: " ~~Visit [Google](http://google.com)~~",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindText,
				Content: " ~~Visit ",
			}, {
				Type:    entity.ElementKindLink,
				Content: "[Google](http://google.com)",
			}, {
				Type:    entity.ElementKindText,
				Content: "~~",
			},
			},
		},
		{
			name:  "Link inside a subscript:",
			input: " H~2~O + [CaCO3](http://example.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindText,
				Content: " H~2~O + ",
			}, {
				Type:    entity.ElementKindLink,
				Content: "[CaCO3](http://example.com)",
			},
			},
		},
		{
			name:  "Link inside a superscript:",
			input: " E = mc^2 + [eV](http://example.com)",
			expect: []entity.LineElement{{
				Type:    entity.ElementKindText,
				Content: " E = mc^2 + ",
			}, {
				Type:    entity.ElementKindLink,
				Content: "[eV](http://example.com)",
			},
			},
		},
		{
			name:  "With title",
			input: "This is [an example](http://example.com/ \"Title\") inline link.",
			expect: []entity.LineElement{
				{
					Type:    entity.ElementKindText,
					Content: "This is ",
				},
				{
					Type:    entity.ElementKindLink,
					Content: "[an example](http://example.com/ \"Title\")",
				},
				{
					Type:    entity.ElementKindText,
					Content: " inline link.",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineLinkElement([]entity.LineElement{{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
