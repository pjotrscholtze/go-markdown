package spanelements

import (
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineItalicAltElementNormalNontext(t *testing.T) {
	input := "Hello _smile_"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
		&entity.LineElement{Type: "nontext", Content: "input"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input},
		&entity.LineElement{Type: "nontext", Content: "input"}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
func TestParseLineItalicAltElementNormal(t *testing.T) {
	input := "Hello _smile_"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.MarkdownElement{}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{
			Type:    "text",
			Content: "  ",
		},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementMultipleItalics(t *testing.T) {
	input := "Hello _smile__laugh_"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_laugh_"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello _smile_ "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: " Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
		&entity.LineElement{Type: entity.ElementKindText, Content: " "},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementSpecialChars(t *testing.T) {
	input := "Hello @username _smile_"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello @username "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementInterspersed(t *testing.T) {
	input := "_smile_Hello_laugh_"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_laugh_"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementOnlyItalics(t *testing.T) {
	input := "_smile__laugh_"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_laugh_"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello World"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello _smile__laugh__heart_"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_laugh_"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_heart_"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "@username#hashtag"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "_smile_" + strings.Repeat("b", 100000)
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: strings.Repeat("a", 100000)},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
		&entity.LineElement{Type: entity.ElementKindText, Content: strings.Repeat("b", 100000)},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello _smile__laugh__heart__grinning__rofl_"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_smile_"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_laugh_"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_heart_"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_grinning_"},
		&entity.LineElement{Type: entity.ElementKindItalic, Content: "_rofl_"},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementColon(t *testing.T) {
	input := "Hello_ smile "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello_ smile "},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}

func TestParseLineItalicAltElementColon2(t *testing.T) {
	input := "Hello_ smile, blah_ testing "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello_ smile, blah_ testing "},
	}
	result := ParseLineItalicAltElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})
	if !equalResults(expectedOutput, result) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
