package spanelements

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineBoldAltElementNormalNontext(t *testing.T) {
	input := "Hello **smile**"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "**smile**"},
		{Type: "nontext", Content: "input"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}, {Type: "nontext", Content: "input"}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineBoldAltElementNormal(t *testing.T) {
	input := "Hello **smile**"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "**smile**"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.LineElement{}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.LineElement{
		{
			Type:    "text",
			Content: "  ",
		},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementMultipleItalics(t *testing.T) {
	input := "Hello **smile****laugh**"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "**smile**"},
		{"bold", "**laugh**"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello **smile** "
	expectedOutput := []entity.LineElement{
		{"text", " Hello "},
		{"bold", "**smile**"},
		{"text", " "},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementSpecialChars(t *testing.T) {
	input := "Hello @username **smile**"
	expectedOutput := []entity.LineElement{
		{"text", "Hello @username "},
		{"bold", "**smile**"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementInterspersed(t *testing.T) {
	input := "**smile**Hello**laugh**"
	expectedOutput := []entity.LineElement{
		{"bold", "**smile**"},
		{"text", "Hello"},
		{"bold", "**laugh**"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementOnlyItalics(t *testing.T) {
	input := "**smile****laugh**"
	expectedOutput := []entity.LineElement{
		{"bold", "**smile**"},
		{"bold", "**laugh**"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.LineElement{
		{"text", "Hello World"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello **smile****laugh****heart**"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "**smile**"},
		{"bold", "**laugh**"},
		{"bold", "**heart**"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.LineElement{
		{"text", "@username#hashtag"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.LineElement{
		{"text", "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "**smile**" + strings.Repeat("b", 100000)
	expectedOutput := []entity.LineElement{
		{"text", strings.Repeat("a", 100000)},
		{"bold", "**smile**"},
		{"text", strings.Repeat("b", 100000)},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello **smile****laugh****heart****grinning****rofl**"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"bold", "**smile**"},
		{"bold", "**laugh**"},
		{"bold", "**heart**"},
		{"bold", "**grinning**"},
		{"bold", "**rofl**"},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementColon(t *testing.T) {
	input := "Hello** smile "
	expectedOutput := []entity.LineElement{
		{"text", "Hello** smile "},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineBoldAltElementColon2(t *testing.T) {
	input := "Hello** smile, blah** testing "
	expectedOutput := []entity.LineElement{
		{"text", "Hello** smile, blah** testing "},
	}
	result := parseLineBoldAltElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
