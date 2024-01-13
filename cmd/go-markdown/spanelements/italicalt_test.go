package spanelements

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineItalicAltElementNormalNontext(t *testing.T) {
	input := "Hello _smile_"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "_smile_"},
		{Type: "nontext", Content: "input"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}, {Type: "nontext", Content: "input"}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineItalicAltElementNormal(t *testing.T) {
	input := "Hello _smile_"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "_smile_"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.LineElement{}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.LineElement{
		{
			Type:    "text",
			Content: "  ",
		},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementMultipleItalics(t *testing.T) {
	input := "Hello _smile__laugh_"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "_smile_"},
		{"italic", "_laugh_"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello _smile_ "
	expectedOutput := []entity.LineElement{
		{"text", " Hello "},
		{"italic", "_smile_"},
		{"text", " "},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementSpecialChars(t *testing.T) {
	input := "Hello @username _smile_"
	expectedOutput := []entity.LineElement{
		{"text", "Hello @username "},
		{"italic", "_smile_"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementInterspersed(t *testing.T) {
	input := "_smile_Hello_laugh_"
	expectedOutput := []entity.LineElement{
		{"italic", "_smile_"},
		{"text", "Hello"},
		{"italic", "_laugh_"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementOnlyItalics(t *testing.T) {
	input := "_smile__laugh_"
	expectedOutput := []entity.LineElement{
		{"italic", "_smile_"},
		{"italic", "_laugh_"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.LineElement{
		{"text", "Hello World"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello _smile__laugh__heart_"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "_smile_"},
		{"italic", "_laugh_"},
		{"italic", "_heart_"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.LineElement{
		{"text", "@username#hashtag"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.LineElement{
		{"text", "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "_smile_" + strings.Repeat("b", 100000)
	expectedOutput := []entity.LineElement{
		{"text", strings.Repeat("a", 100000)},
		{"italic", "_smile_"},
		{"text", strings.Repeat("b", 100000)},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello _smile__laugh__heart__grinning__rofl_"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"italic", "_smile_"},
		{"italic", "_laugh_"},
		{"italic", "_heart_"},
		{"italic", "_grinning_"},
		{"italic", "_rofl_"},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementColon(t *testing.T) {
	input := "Hello_ smile "
	expectedOutput := []entity.LineElement{
		{"text", "Hello_ smile "},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementColon2(t *testing.T) {
	input := "Hello_ smile, blah_ testing "
	expectedOutput := []entity.LineElement{
		{"text", "Hello_ smile, blah_ testing "},
	}
	result := parseLineItalicAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
