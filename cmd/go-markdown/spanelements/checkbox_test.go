package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestCheckbox(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.LineElement
	}{

		{name: "No checkbox", input: "~~Hello~~", expect: []entity.LineElement{
			{Type: "text", Content: "~~Hello~~"},
		}},
		{name: "Empty checkbox", input: "[ ] Hello", expect: []entity.LineElement{
			{Type: "checkbox", Content: "[ ] Hello"},
		}},
		{name: "Checked checkbox", input: "[x] Hello", expect: []entity.LineElement{
			{Type: "checkbox", Content: "[x] Hello"},
		}},
		{name: "Multi checked checkbox", input: "[x] Hello [x] world", expect: []entity.LineElement{
			{Type: "checkbox", Content: "[x] Hello [x] world"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineCheckboxElement([]entity.LineElement{{Type: "text", Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
