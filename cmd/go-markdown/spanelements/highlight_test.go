package spanelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineHighlightElement(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.LineElement
	}{
		{name: "Single word highlight:", input: "==test==", expect: []entity.LineElement{{
			Type:    "highlight",
			Content: "==test==",
		}}},
		{name: "Multiple words highlight:", input: "==this is a test==", expect: []entity.LineElement{{
			Type:    "highlight",
			Content: "==this is a test==",
		}}},
		{name: "Highlight within a sentence:", input: "This is a ==test== sentence.", expect: []entity.LineElement{
			{
				Type:    "text",
				Content: "This is a ",
			},
			{
				Type:    "highlight",
				Content: "==test==",
			},
			{
				Type:    "text",
				Content: " sentence.",
			},
		}},
		{name: "Highlight without opening ==:", input: "test==", expect: []entity.LineElement{{
			Type:    "text",
			Content: "test==",
		}}},
		{name: "Highlight without closing ==:", input: "==test", expect: []entity.LineElement{{
			Type:    "text",
			Content: "==test",
		}}},
		{name: "Highlight with leading and trailing spaces:", input: "== test ==", expect: []entity.LineElement{{
			Type:    "highlight",
			Content: "== test ==",
		}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseLineHighlightElement([]entity.LineElement{{Type: "text", Content: test.input}})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
