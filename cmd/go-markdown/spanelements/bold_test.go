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
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindBold, Content: "__smile__"},
		{Type: "nontext", Content: "input"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}, {Type: "nontext", Content: "input"}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineBoldElementNormal(t *testing.T) {
	input := "Hello __smile__"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindBold, Content: "__smile__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.LineElement{}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.LineElement{
		{
			Type:    entity.ElementKindText,
			Content: "  ",
		},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementMultipleItalics(t *testing.T) {
	input := "Hello __smile____laugh__"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindBold, Content: "__smile__"},
		{Type: entity.ElementKindBold, Content: "__laugh__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello __smile__ "
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: " Hello "},
		{Type: entity.ElementKindBold, Content: "__smile__"},
		{Type: entity.ElementKindText, Content: " "},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementSpecialChars(t *testing.T) {
	input := "Hello @username __smile__"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello @username "},
		{Type: entity.ElementKindBold, Content: "__smile__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementInterspersed(t *testing.T) {
	input := "__smile__Hello__laugh__"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindBold, Content: "__smile__"},
		{Type: entity.ElementKindText, Content: "Hello"},
		{Type: entity.ElementKindBold, Content: "__laugh__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlyItalics(t *testing.T) {
	input := "__smile____laugh__"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindBold, Content: "__smile__"},
		{Type: entity.ElementKindBold, Content: "__laugh__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello World"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello __smile____laugh____heart__"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindBold, Content: "__smile__"},
		{Type: entity.ElementKindBold, Content: "__laugh__"},
		{Type: entity.ElementKindBold, Content: "__heart__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "@username#hashtag"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "__smile__" + strings.Repeat("b", 100000)
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: strings.Repeat("a", 100000)},
		{Type: entity.ElementKindBold, Content: "__smile__"},
		{Type: entity.ElementKindText, Content: strings.Repeat("b", 100000)},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello __smile____laugh____heart____grinning____rofl__"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindBold, Content: "__smile__"},
		{Type: entity.ElementKindBold, Content: "__laugh__"},
		{Type: entity.ElementKindBold, Content: "__heart__"},
		{Type: entity.ElementKindBold, Content: "__grinning__"},
		{Type: entity.ElementKindBold, Content: "__rofl__"},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementColon(t *testing.T) {
	input := "Hello__ smile "
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello__ smile "},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementColon2(t *testing.T) {
	input := "Hello__ smile, blah__ testing "
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello__ smile, blah__ testing "},
	}
	result := parseLineBoldElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
