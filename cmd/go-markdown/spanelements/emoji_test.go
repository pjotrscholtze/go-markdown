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
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
		{Type: "nontext", Content: "input"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}, {Type: "nontext", Content: "input"}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineEmojiElementNormal(t *testing.T) {
	input := "Hello :smile:"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.LineElement{}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.LineElement{
		{
			Type:    entity.ElementKindText,
			Content: "  ",
		},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementMultipleEmojis(t *testing.T) {
	input := "Hello :smile::laugh:"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
		{Type: entity.ElementKindEmoji, Content: ":laugh:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello :smile: "
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: " Hello "},
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
		{Type: entity.ElementKindText, Content: " "},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementSpecialChars(t *testing.T) {
	input := "Hello @username :smile:"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello @username "},
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementInterspersed(t *testing.T) {
	input := ":smile:Hello:laugh:"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
		{Type: entity.ElementKindText, Content: "Hello"},
		{Type: entity.ElementKindEmoji, Content: ":laugh:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementOnlyEmojis(t *testing.T) {
	input := ":smile::laugh:"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
		{Type: entity.ElementKindEmoji, Content: ":laugh:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello World"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementMultipleDifferentEmojis(t *testing.T) {
	input := "Hello :smile::laugh::heart:"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
		{Type: entity.ElementKindEmoji, Content: ":laugh:"},
		{Type: entity.ElementKindEmoji, Content: ":heart:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "@username#hashtag"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and emojis. Let's see how well our function handles it."
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "This is a very long line of text that contains a mix of text and emojis. Let's see how well our function handles it."},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + ":smile:" + strings.Repeat("b", 100000)
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: strings.Repeat("a", 100000)},
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
		{Type: entity.ElementKindText, Content: strings.Repeat("b", 100000)},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementLargeNumberOfEmojis(t *testing.T) {
	input := "Hello :smile::laugh::heart::grinning::rofl:"
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello "},
		{Type: entity.ElementKindEmoji, Content: ":smile:"},
		{Type: entity.ElementKindEmoji, Content: ":laugh:"},
		{Type: entity.ElementKindEmoji, Content: ":heart:"},
		{Type: entity.ElementKindEmoji, Content: ":grinning:"},
		{Type: entity.ElementKindEmoji, Content: ":rofl:"},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementColon(t *testing.T) {
	input := "Hello: smile "
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello: smile "},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineEmojiElementColon2(t *testing.T) {
	input := "Hello: smile, blah: testing "
	expectedOutput := []entity.LineElement{
		{Type: entity.ElementKindText, Content: "Hello: smile, blah: testing "},
	}
	result := parseLineEmojiElement([]entity.LineElement{{Type: entity.ElementKindText, Content: input}})
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
