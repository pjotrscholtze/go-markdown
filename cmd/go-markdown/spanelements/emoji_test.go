package spanelements

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineEmojiElementNormalNontext(t *testing.T) {
	input := "Hello :smile:"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"emoji", ":smile:"},
		{Type: "nontext", Content: "input"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}, {Type: "nontext", Content: "input"}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineEmojiElementNormal(t *testing.T) {
	input := "Hello :smile:"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"emoji", ":smile:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.LineElement{}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.LineElement{
		{
			Type:    "text",
			Content: "  ",
		},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementMultipleEmojis(t *testing.T) {
	input := "Hello :smile::laugh:"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"emoji", ":smile:"},
		{"emoji", ":laugh:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello :smile: "
	expectedOutput := []entity.LineElement{
		{"text", " Hello "},
		{"emoji", ":smile:"},
		{"text", " "},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementSpecialChars(t *testing.T) {
	input := "Hello @username :smile:"
	expectedOutput := []entity.LineElement{
		{"text", "Hello @username "},
		{"emoji", ":smile:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementInterspersed(t *testing.T) {
	input := ":smile:Hello:laugh:"
	expectedOutput := []entity.LineElement{
		{"emoji", ":smile:"},
		{"text", "Hello"},
		{"emoji", ":laugh:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementOnlyEmojis(t *testing.T) {
	input := ":smile::laugh:"
	expectedOutput := []entity.LineElement{
		{"emoji", ":smile:"},
		{"emoji", ":laugh:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.LineElement{
		{"text", "Hello World"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementMultipleDifferentEmojis(t *testing.T) {
	input := "Hello :smile::laugh::heart:"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"emoji", ":smile:"},
		{"emoji", ":laugh:"},
		{"emoji", ":heart:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.LineElement{
		{"text", "@username#hashtag"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and emojis. Let's see how well our function handles it."
	expectedOutput := []entity.LineElement{
		{"text", "This is a very long line of text that contains a mix of text and emojis. Let's see how well our function handles it."},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + ":smile:" + strings.Repeat("b", 100000)
	expectedOutput := []entity.LineElement{
		{"text", strings.Repeat("a", 100000)},
		{"emoji", ":smile:"},
		{"text", strings.Repeat("b", 100000)},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementLargeNumberOfEmojis(t *testing.T) {
	input := "Hello :smile::laugh::heart::grinning::rofl:"
	expectedOutput := []entity.LineElement{
		{"text", "Hello "},
		{"emoji", ":smile:"},
		{"emoji", ":laugh:"},
		{"emoji", ":heart:"},
		{"emoji", ":grinning:"},
		{"emoji", ":rofl:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementColon(t *testing.T) {
	input := "Hello: smile "
	expectedOutput := []entity.LineElement{
		{"text", "Hello: smile "},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementColon2(t *testing.T) {
	input := "Hello: smile, blah: testing "
	expectedOutput := []entity.LineElement{
		{"text", "Hello: smile, blah: testing "},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: "text", Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
