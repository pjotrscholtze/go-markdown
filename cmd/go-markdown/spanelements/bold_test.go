package spanelements

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineBoldElementNormalNontext(t *testing.T) {
	input := "Hello __smile__"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "__smile__"},
		{Type: "nontext", Content: "input"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}, {Type: "nontext", Content: "input"}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineBoldElementNormal(t *testing.T) {
	input := "Hello __smile__"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "__smile__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.LineElement{}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.LineElement{
		{
			Type:    "text",
			Content: "  ",
		},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementMultipleItalics(t *testing.T) {
	input := "Hello __smile____laugh__"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "__smile__"},
		{"bold", "__laugh__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello __smile__ "
	expectedOutput := []entity.LineElement{
		{"text", " Hello "},
		{"bold", "__smile__"},
		{"text", " "},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementSpecialChars(t *testing.T) {
	input := "Hello @username __smile__"
	expectedOutput := []entity.LineElement{
		{"text", "Hello @username "},
		{"bold", "__smile__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementInterspersed(t *testing.T) {
	input := "__smile__Hello__laugh__"
	expectedOutput := []entity.LineElement{
		{"bold", "__smile__"},
		{"text", "Hello"},
		{"bold", "__laugh__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlyItalics(t *testing.T) {
	input := "__smile____laugh__"
	expectedOutput := []entity.LineElement{
		{"bold", "__smile__"},
		{"bold", "__laugh__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.LineElement{
		{"text", "Hello World"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello __smile____laugh____heart__"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "__smile__"},
		{"bold", "__laugh__"},
		{"bold", "__heart__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.LineElement{
		{"text", "@username#hashtag"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.LineElement{
		{"text", "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "__smile__" + strings.Repeat("b", 100000)
	expectedOutput := []entity.LineElement{
		{"text", strings.Repeat("a", 100000)},
		{"bold", "__smile__"},
		{"text", strings.Repeat("b", 100000)},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello __smile____laugh____heart____grinning____rofl__"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "__smile__"},
		{"bold", "__laugh__"},
		{"bold", "__heart__"},
		{"bold", "__grinning__"},
		{"bold", "__rofl__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementColon(t *testing.T) {
	input := "Hello__ smile "
	expectedOutput := []entity.LineElement{
		{"text", "Hello__ smile "},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementColon2(t *testing.T) {
	input := "Hello__ smile, blah__ testing "
	expectedOutput := []entity.LineElement{
		{"text", "Hello__ smile, blah__ testing "},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
