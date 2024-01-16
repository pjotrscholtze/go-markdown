package spanelements

import (
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineItalicElementNormalNontext(t *testing.T) {
	input := "Hello *smile*"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
		&entity.LineElement{Type: "nontext", Content: "input"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input},
		&entity.LineElement{Type: "nontext", Content: "input"}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineItalicElementNormal(t *testing.T) {
	input := "Hello *smile*"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.MarkdownElement{}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{
			Type:    entity.ElementKindText,
			Content: "  ",
		},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementMultipleItalics(t *testing.T) {
	input := "Hello *smile**laugh*"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*laugh*"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello *smile* "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: " Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
		&entity.LineElement{Type: entity.ElementKindText, Content: " "},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementSpecialChars(t *testing.T) {
	input := "Hello @username *smile*"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello @username "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementInterspersed(t *testing.T) {
	input := "*smile*Hello*laugh*"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*laugh*"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementOnlyItalics(t *testing.T) {
	input := "*smile**laugh*"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*laugh*"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello World"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello *smile**laugh**heart*"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*laugh*"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*heart*"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "@username#hashtag"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "*smile*" + strings.Repeat("b", 100000)
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: strings.Repeat("a", 100000)},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
		&entity.LineElement{Type: entity.ElementKindText, Content: strings.Repeat("b", 100000)},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello *smile**laugh**heart**grinning**rofl*"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*smile*"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*laugh*"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*heart*"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*grinning*"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "*rofl*"},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementColon(t *testing.T) {
	input := "Hello* smile "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello* smile "},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementColon2(t *testing.T) {
	input := "Hello* smile, blah* testing "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello* smile, blah* testing "},
	}
	result := parseLineItalicElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
