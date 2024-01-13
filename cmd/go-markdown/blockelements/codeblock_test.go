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
		expect []entity.LineElement
	}{

		{name: "No content", input: "", expect: []entity.LineElement{}},
		{name: "Simple example", input: strings.Join([]string{
			"```",
			"test",
			"```",
		}, "\n"), expect: []entity.LineElement{
			{
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
		}, "\n"), expect: []entity.LineElement{
			{Type: entity.ElementKindText, Content: "before"},
			{
				Type: entity.ElementKindCodeblock,
				Content: strings.Join([]string{
					"```",
					"test",
					"```",
				}, "\n"),
			},
			{Type: entity.ElementKindText, Content: "after"},
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
		}, "\n"), expect: []entity.LineElement{
			{Type: entity.ElementKindText, Content: "before"},
			{
				Type: entity.ElementKindCodeblock,
				Content: strings.Join([]string{
					"```",
					"test1",
					"```",
				}, "\n"),
			},
			{Type: entity.ElementKindText, Content: "middle"},
			{
				Type: entity.ElementKindCodeblock,
				Content: strings.Join([]string{
					"```",
					"test2",
					"```",
				}, "\n"),
			},
			{Type: entity.ElementKindText, Content: "after"},
		}},
		{name: "Example with language hint", input: strings.Join([]string{
			"```json",
			"['test']",
			"```",
		}, "\n"), expect: []entity.LineElement{
			{
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
			got := parseLineCodeblockElement([]entity.LineElement{{Type: entity.ElementKindText, Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
