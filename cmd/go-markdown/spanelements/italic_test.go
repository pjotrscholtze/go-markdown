package spanelements

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineItalicElementNormalNontext(t *testing.T) {
	input := "Hello *smile*"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "*smile*"},
		{Type: "nontext", Content: "input"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}, {Type: "nontext", Content: "input"}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineItalicElementNormal(t *testing.T) {
	input := "Hello *smile*"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "*smile*"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.LineElement{}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.LineElement{
		{
			Type:    "text",
			Content: "  ",
		},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementMultipleItalics(t *testing.T) {
	input := "Hello *smile**laugh*"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "*smile*"},
		{"italic", "*laugh*"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello *smile* "
	expectedOutput := []entity.LineElement{
		{"text", " Hello "},
		{"italic", "*smile*"},
		{"text", " "},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementSpecialChars(t *testing.T) {
	input := "Hello @username *smile*"
	expectedOutput := []entity.LineElement{
		{"text", "Hello @username "},
		{"italic", "*smile*"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementInterspersed(t *testing.T) {
	input := "*smile*Hello*laugh*"
	expectedOutput := []entity.LineElement{
		{"italic", "*smile*"},
		{"text", "Hello"},
		{"italic", "*laugh*"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementOnlyItalics(t *testing.T) {
	input := "*smile**laugh*"
	expectedOutput := []entity.LineElement{
		{"italic", "*smile*"},
		{"italic", "*laugh*"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.LineElement{
		{"text", "Hello World"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello *smile**laugh**heart*"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "*smile*"},
		{"italic", "*laugh*"},
		{"italic", "*heart*"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.LineElement{
		{"text", "@username#hashtag"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.LineElement{
		{"text", "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "*smile*" + strings.Repeat("b", 100000)
	expectedOutput := []entity.LineElement{
		{"text", strings.Repeat("a", 100000)},
		{"italic", "*smile*"},
		{"text", strings.Repeat("b", 100000)},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello *smile**laugh**heart**grinning**rofl*"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "*smile*"},
		{"italic", "*laugh*"},
		{"italic", "*heart*"},
		{"italic", "*grinning*"},
		{"italic", "*rofl*"},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementColon(t *testing.T) {
	input := "Hello* smile "
	expectedOutput := []entity.LineElement{
		{"text", "Hello* smile "},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicElementColon2(t *testing.T) {
	input := "Hello* smile, blah* testing "
	expectedOutput := []entity.LineElement{
		{"text", "Hello* smile, blah* testing "},
	}
	result := parseLineItalicElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
