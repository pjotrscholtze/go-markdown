package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestCheckbox(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No checkbox", input: "~~Hello~~", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindText, Content: "~~Hello~~"},
		}},
		{name: "Empty checkbox", input: "[ ] Hello", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindCheckbox, Content: "[ ]"},
			&entity.LineElement{Type: entity.ElementKindText, Content: " Hello"},
		}},
		{name: "Checked checkbox", input: "[x] Hello", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindCheckbox, Content: "[x]"},
			&entity.LineElement{Type: entity.ElementKindText, Content: " Hello"},
		}},
		{name: "Multi checked checkbox", input: "[x] Hello [x] world", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindCheckbox, Content: "[x]"},
			&entity.LineElement{Type: entity.ElementKindText, Content: " Hello "},
			&entity.LineElement{Type: entity.ElementKindCheckbox, Content: "[x]"},
			&entity.LineElement{Type: entity.ElementKindText, Content: " world"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseLineCheckboxElement([]entity.MarkdownElement{
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
