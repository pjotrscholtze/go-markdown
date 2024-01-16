package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineImageElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{
		{
			name:  "Image with no alt text:",
			input: "![]()",
			expect: []entity.MarkdownElement{
				&entity.LineElement{
					Type:    entity.ElementKindImage,
					Content: "![]()",
				},
			},
		},
		{
			name:  "Image with empty alt text:",
			input: "![](image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![](image.jpg)"},
			},
		},
		{
			name:  "Image with space in alt text:",
			input: "![ ](image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![ ](image.jpg)"},
			},
		},
		{
			name:  "Image with special characters in alt text:",
			input: "![!@#$%^&*()](image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![!@#$%^&*()](image.jpg)"},
			},
		},
		{
			name:  "Image with relative path:",
			input: "![relative path](./images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![relative path](./images/image.jpg)"},
			},
		},
		{
			name:  "Image with absolute path:",
			input: "![absolute path](/images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![absolute path](/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (HTTP):",
			input: "![protocol](http://example.com/images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![protocol](http://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (HTTPS):",
			input: "![protocol](https://example.com/images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![protocol](https://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (FTP):",
			input: "![protocol](ftp://example.com/images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![protocol](ftp://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (SFTP):",
			input: "![protocol](sftp://example.com/images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![protocol](sftp://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (SSH):",
			input: "![protocol](ssh://example.com/images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![protocol](ssh://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (file):",
			input: "![protocol](file:///home/user/images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![protocol](file:///home/user/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (mailto):",
			input: "![protocol](mailto:user@example.com)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![protocol](mailto:user@example.com)"},
			},
		},
		{
			name:  "Image with protocol (tel):",
			input: "![protocol](tel:+1234567890)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![protocol](tel:+1234567890)"},
			},
		},
		{
			name:  "Image with invalid protocol:",
			input: "![invalid protocol](invalid://example.com/images/image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![invalid protocol](invalid://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image within a sentence:",
			input: "This is a sentence with an image ![image](image.jpg) at the end.",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: "This is a sentence with an image "},
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![image](image.jpg)"},
				&entity.LineElement{Type: entity.ElementKindText, Content: " at the end."},
			},
		},
		{
			name:  "Image followed by a link:",
			input: "![image](image.jpg)[link](https://example.com)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![image](image.jpg)"},
				&entity.LineElement{Type: entity.ElementKindText, Content: "[link](https://example.com)"},
			},
		},
		{
			name:  "Image within a heading:",
			input: "# Heading with an image ![image](image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: "# Heading with an image "},
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![image](image.jpg)"},
			},
		},
		{
			name:  "Image within a list item:",
			input: "- List item with an image ![image](image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: "- List item with an image "},
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![image](image.jpg)"},
			},
		},
		{
			name:  "Image within a blockquote:",
			input: "> Blockquote with an image ![image](image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: "> Blockquote with an image "},
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![image](image.jpg)"},
			},
		},
		{
			name:  "Image within a strikethrough:",
			input: "~~Strikethrough with an image ![image](image.jpg)~~",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: "~~Strikethrough with an image "},
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![image](image.jpg)"},
				&entity.LineElement{Type: entity.ElementKindText, Content: "~~"},
			},
		},
		{
			name:  "Image with a subscript:",
			input: "H~2~O ![image](image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: "H~2~O "},
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![image](image.jpg)"},
			},
		},
		{
			name:  "Image with a superscript:",
			input: "X^2^ ![image](image.jpg)",
			expect: []entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: "X^2^ "},
				&entity.LineElement{Type: entity.ElementKindImage, Content: "![image](image.jpg)"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineImageElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}

func equalResults(a, b []entity.MarkdownElement) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Kind() != b[i].Kind() {
			a := v.Kind()
			c := (b[i].Kind())
			_ = a
			_ = c
			return false
		}
		if v.AsMarkdownString() != b[i].AsMarkdownString() {
			return false
		}
	}
	return true
}
