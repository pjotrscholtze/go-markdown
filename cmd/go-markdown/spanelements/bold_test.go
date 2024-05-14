package spanelements

import (
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineBoldElementNormalNontext(t *testing.T) {
	input := "Hello __smile__"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
		&entity.LineElement{Type: "nontext", Content: "input"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input},
		&entity.LineElement{Type: "nontext", Content: "input"}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}
func TestParseLineBoldElementNormal(t *testing.T) {
	input := "Hello __smile__"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.MarkdownElement{}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{
			Type:    entity.ElementKindText,
			Content: "  ",
		},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementMultipleItalics(t *testing.T) {
	input := "Hello __smile____laugh__"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__laugh__"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello __smile__ "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: " Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
		&entity.LineElement{Type: entity.ElementKindText, Content: " "},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementSpecialChars(t *testing.T) {
	input := "Hello @username __smile__"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello @username "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementInterspersed(t *testing.T) {
	input := "__smile__Hello__laugh__"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__laugh__"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlyItalics(t *testing.T) {
	input := "__smile____laugh__"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__laugh__"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello World"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello __smile____laugh____heart__"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__laugh__"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__heart__"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "@username#hashtag"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "__smile__" + strings.Repeat("b", 100000)
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: strings.Repeat("a", 100000)},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
		&entity.LineElement{Type: entity.ElementKindText, Content: strings.Repeat("b", 100000)},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello __smile____laugh____heart____grinning____rofl__"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__smile__"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__laugh__"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__heart__"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__grinning__"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__rofl__"},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementColon(t *testing.T) {
	input := "Hello__ smile "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello__ smile "},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}

func TestParseLineBoldElementColon2(t *testing.T) {
	input := "Hello__ smile, blah__ testing "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "__ smile, blah__"},
		&entity.LineElement{Type: entity.ElementKindText, Content: " testing "},
	}
	result := ParseLineBoldElement([]entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: input}},
		func(input string) []entity.MarkdownElement {
			return []entity.MarkdownElement{&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: input,
			}}
		})

	if !equalResults(result, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, result)
	}
}
