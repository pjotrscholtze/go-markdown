package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineImageElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.LineElement
	}{
		{
			name:  "Image with no alt text:",
			input: "![]()",
			expect: []entity.LineElement{
				{
					Type:    "image",
					Content: "![]()",
				},
			},
		},
		{
			name:  "Image with empty alt text:",
			input: "![](image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![](image.jpg)"},
			},
		},
		{
			name:  "Image with space in alt text:",
			input: "![ ](image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![ ](image.jpg)"},
			},
		},
		{
			name:  "Image with special characters in alt text:",
			input: "![!@#$%^&*()](image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![!@#$%^&*()](image.jpg)"},
			},
		},
		{
			name:  "Image with relative path:",
			input: "![relative path](./images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![relative path](./images/image.jpg)"},
			},
		},
		{
			name:  "Image with absolute path:",
			input: "![absolute path](/images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![absolute path](/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (HTTP):",
			input: "![protocol](http://example.com/images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![protocol](http://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (HTTPS):",
			input: "![protocol](https://example.com/images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![protocol](https://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (FTP):",
			input: "![protocol](ftp://example.com/images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![protocol](ftp://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (SFTP):",
			input: "![protocol](sftp://example.com/images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![protocol](sftp://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (SSH):",
			input: "![protocol](ssh://example.com/images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![protocol](ssh://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (file):",
			input: "![protocol](file:///home/user/images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![protocol](file:///home/user/images/image.jpg)"},
			},
		},
		{
			name:  "Image with protocol (mailto):",
			input: "![protocol](mailto:user@example.com)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![protocol](mailto:user@example.com)"},
			},
		},
		{
			name:  "Image with protocol (tel):",
			input: "![protocol](tel:+1234567890)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![protocol](tel:+1234567890)"},
			},
		},
		{
			name:  "Image with invalid protocol:",
			input: "![invalid protocol](invalid://example.com/images/image.jpg)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![invalid protocol](invalid://example.com/images/image.jpg)"},
			},
		},
		{
			name:  "Image within a sentence:",
			input: "This is a sentence with an image ![image](image.jpg) at the end.",
			expect: []entity.LineElement{
				{Type: "text", Content: "This is a sentence with an image "},
				{Type: "image", Content: "![image](image.jpg)"},
				{Type: "text", Content: " at the end."},
			},
		},
		{
			name:  "Image followed by a link:",
			input: "![image](image.jpg)[link](https://example.com)",
			expect: []entity.LineElement{
				{Type: "image", Content: "![image](image.jpg)"},
				{Type: "text", Content: "[link](https://example.com)"},
			},
		},
		{
			name:  "Image within a heading:",
			input: "# Heading with an image ![image](image.jpg)",
			expect: []entity.LineElement{
				{Type: "text", Content: "# Heading with an image "},
				{Type: "image", Content: "![image](image.jpg)"},
			},
		},
		{
			name:  "Image within a list item:",
			input: "- List item with an image ![image](image.jpg)",
			expect: []entity.LineElement{
				{Type: "text", Content: "- List item with an image "},
				{Type: "image", Content: "![image](image.jpg)"},
			},
		},
		{
			name:  "Image within a blockquote:",
			input: "> Blockquote with an image ![image](image.jpg)",
			expect: []entity.LineElement{
				{Type: "text", Content: "> Blockquote with an image "},
				{Type: "image", Content: "![image](image.jpg)"},
			},
		},
		{
			name:  "Image within a strikethrough:",
			input: "~~Strikethrough with an image ![image](image.jpg)~~",
			expect: []entity.LineElement{
				{Type: "text", Content: "~~Strikethrough with an image "},
				{Type: "image", Content: "![image](image.jpg)"},
				{Type: "text", Content: "~~"},
			},
		},
		{
			name:  "Image with a subscript:",
			input: "H~2~O ![image](image.jpg)",
			expect: []entity.LineElement{
				{Type: "text", Content: "H~2~O "},
				{Type: "image", Content: "![image](image.jpg)"},
			},
		},
		{
			name:  "Image with a superscript:",
			input: "X^2^ ![image](image.jpg)",
			expect: []entity.LineElement{
				{Type: "text", Content: "X^2^ "},
				{Type: "image", Content: "![image](image.jpg)"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineImageElement([]entity.LineElement{{Type: "text", Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}

func equalResults(a, b []entity.LineElement) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
