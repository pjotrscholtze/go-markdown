package parser

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/blockelements"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/spanelements"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{
		{name: "No content", input: "", expect: []entity.MarkdownElement{}},
		{name: "Text content", input: "Hello", expect: []entity.MarkdownElement{
			&entity.LineElement{Type: entity.ElementKindText, Content: "Hello"},
		}},
		{name: "Empty checkbox", input: "[ ] Hello", expect: []entity.MarkdownElement{
			entity.NewCheckboxMarkdownElement("[ ]",
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				}),
			&entity.LineElement{Type: entity.ElementKindText, Content: " Hello"},
		}},
		{name: "Heading", input: "# Hello", expect: []entity.MarkdownElement{
			entity.NewHeaderMarkdownElement("# Hello", func(input string) []entity.MarkdownElement {
				md := entity.NewRawTextMarkdownElement(input,
					func(input string) []entity.MarkdownElement {
						return []entity.MarkdownElement{&entity.LineElement{
							Type:    entity.ElementKindText,
							Content: input,
						}}
					})
				return []entity.MarkdownElement{&md}
			}),
		}},
		{name: "Multiline", input: `# Hello
[ ] Hello`, expect: []entity.MarkdownElement{
			entity.NewHeaderMarkdownElement("# Hello", func(input string) []entity.MarkdownElement {
				md := entity.NewRawTextMarkdownElement(input,
					func(input string) []entity.MarkdownElement {
						return []entity.MarkdownElement{&entity.LineElement{
							Type:    entity.ElementKindText,
							Content: input,
						}}
					})
				return []entity.MarkdownElement{&md}
			}),
			entity.NewCheckboxMarkdownElement("[ ]",
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				}),
			&entity.LineElement{Type: entity.ElementKindText, Content: " Hello"},
		}},
		{name: "Multiline with text content", input: `# Hello
Testing
[ ] Hello`, expect: []entity.MarkdownElement{
			entity.NewHeaderMarkdownElement("# Hello", func(input string) []entity.MarkdownElement {
				md := entity.NewRawTextMarkdownElement(input,
					func(input string) []entity.MarkdownElement {
						return []entity.MarkdownElement{&entity.LineElement{
							Type:    entity.ElementKindText,
							Content: input,
						}}
					})
				return []entity.MarkdownElement{&md}
			}),
			&entity.LineElement{Type: entity.ElementKindText, Content: "Testing\n"},
			entity.NewCheckboxMarkdownElement("[ ]",
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				}),
			&entity.LineElement{Type: entity.ElementKindText, Content: " Hello"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseString(test.input, []func(input []entity.MarkdownElement, parser func(input string) []entity.MarkdownElement) []entity.MarkdownElement{
				blockelements.ParseLineHeaderElement,
				spanelements.ParseLineCheckboxElement,
			})
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}

func TestParserNested(t *testing.T) {
	test := `# Testing _lorem_`
	got := ParseString(test, []func(input []entity.MarkdownElement, parser func(input string) []entity.MarkdownElement) []entity.MarkdownElement{
		spanelements.ParseLineCheckboxElement,
		blockelements.ParseLineHeaderElement,
		spanelements.ParseLineItalicAltElement,
	})
	if len(got) != 1 {
		t.Errorf("Expected single element, got %d", len(got))
	}
	if got[0].AsMarkdownString() != test {
		t.Errorf("Expected %v, got %v", test, got[0].AsMarkdownString())
	}
	if got[0].Kind() != entity.ElementKindHeader {
		t.Errorf("Expected header element, got %v", got[0].Kind())
	}
	e, ok := got[0].(entity.HeaderMarkdownElement)
	if !ok {
		t.Errorf("Was unable to cast header element to header element!")
	}
	if e.GetHeadingLevel() != 1 {
		t.Errorf("Header level was expected to be 1, got %d", e.GetHeadingLevel())
	}
	children := e.GetChildren()
	if len(children) != 2 {
		t.Errorf("Header children was expected to be 3, got %d", len(children))
	}

	if children[0].Kind() != entity.ElementKindText {
		t.Errorf("First header child was expected to be text, got %v", children[0].Kind())
	}
	if children[0].AsMarkdownString() != " Testing " {
		t.Errorf("First header child was expected to be ' Testing ', got %v", children[0].AsMarkdownString())
	}
	if children[1].Kind() != entity.ElementKindItalic {
		t.Errorf("First header child was expected to be text, got %v", children[1].Kind())
	}
	if children[1].AsMarkdownString() != "_lorem_" {
		t.Errorf("First header child was expected to be '_lorem_', got %v", children[1].AsMarkdownString())
	}
	it := children[1].(entity.ItalicMarkdownElement)
	if it.GetWrappingSymbolAsRune() != '_' {
		t.Errorf("First header child was expected to be 'lorem', got %v", children[1].AsMarkdownString())
	}
}

func equalResults(a, b []entity.MarkdownElement) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Kind() != b[i].Kind() {
			return false
		}
		if v.AsMarkdownString() != b[i].AsMarkdownString() {
			return false
		}
	}
	return true
}
