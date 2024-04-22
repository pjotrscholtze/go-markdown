package blockelements

import (
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestCodeblockDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No content", input: "", expect: []entity.MarkdownElement{}},
		{name: "Simple example", input: strings.Join([]string{
			"```",
			"test",
			"```",
		}, "\n"), expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindCodeblock,
				Content: strings.Join([]string{
					"```",
					"test",
					"```",
				}, "\n"),
			},
		}},

		{name: "Example with surrounding text", input: strings.Join([]string{
			"before",
			"```",
			"test",
			"```",
			"after",
		}, "\n"), expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindText, Content: "before"},
			&entity.LineElement{
				Type: entity.ElementKindCodeblock,
				Content: strings.Join([]string{
					"```",
					"test",
					"```",
				}, "\n"),
			},
			&entity.LineElement{Type: entity.ElementKindText, Content: "after"},
		}},

		{name: "Example with surrounding text", input: strings.Join([]string{
			"before",
			"```",
			"test1",
			"```",
			"middle",
			"```",
			"test2",
			"```",
			"after",
		}, "\n"), expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindText, Content: "before"},
			&entity.LineElement{
				Type: entity.ElementKindCodeblock,
				Content: strings.Join([]string{
					"```",
					"test1",
					"```",
				}, "\n"),
			},
			&entity.LineElement{Type: entity.ElementKindText, Content: "middle"},
			&entity.LineElement{
				Type: entity.ElementKindCodeblock,
				Content: strings.Join([]string{
					"```",
					"test2",
					"```",
				}, "\n"),
			},
			&entity.LineElement{Type: entity.ElementKindText, Content: "after"},
		}},
		{name: "Example with language hint", input: strings.Join([]string{
			"```json",
			"['test']",
			"```",
		}, "\n"), expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindCodeblock,
				Content: strings.Join([]string{
					"```json",
					"['test']",
					"```",
				}, "\n"),
			},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseLineCodeblockElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}},
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
