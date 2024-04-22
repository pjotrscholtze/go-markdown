package spanelements

import (
	"strings"
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestParseLineBoldAltElementNormalNontext(t *testing.T) {
	input := "Hello **smile**"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
		&entity.LineElement{Type: "nontext", Content: "input"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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
func TestParseLineBoldAltElementNormal(t *testing.T) {
	input := "Hello **smile**"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementEmptyInput(t *testing.T) {
	input := ""
	expectedOutput := []entity.MarkdownElement{}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementWhitespaceInput(t *testing.T) {
	input := "  "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{
			Type:    entity.ElementKindText,
			Content: "  ",
		},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementMultipleItalics(t *testing.T) {
	input := "Hello **smile****laugh**"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**laugh**"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementLeadingTrailingSpaces(t *testing.T) {
	input := " Hello **smile** "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: " Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
		&entity.LineElement{Type: entity.ElementKindText, Content: " "},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementSpecialChars(t *testing.T) {
	input := "Hello @username **smile**"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello @username "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementInterspersed(t *testing.T) {
	input := "**smile**Hello**laugh**"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**laugh**"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementOnlyItalics(t *testing.T) {
	input := "**smile****laugh**"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**laugh**"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementOnlyText(t *testing.T) {
	input := "Hello World"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello World"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementMultipleDifferentItalics(t *testing.T) {
	input := "Hello **smile****laugh****heart**"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**laugh**"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**heart**"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementOnlySpecialChars(t *testing.T) {
	input := "@username#hashtag"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "@username#hashtag"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementLongLinesOfText(t *testing.T) {
	input := "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "This is a very long line of text that contains a mix of text and Italics. Let's see how well our function handles it."},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementVeryLongStrings(t *testing.T) {
	input := strings.Repeat("a", 100000) + "**smile**" + strings.Repeat("b", 100000)
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: strings.Repeat("a", 100000)},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
		&entity.LineElement{Type: entity.ElementKindText, Content: strings.Repeat("b", 100000)},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementLargeNumberOfItalics(t *testing.T) {
	input := "Hello **smile****laugh****heart****grinning****rofl**"
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello "},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**smile**"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**laugh**"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**heart**"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**grinning**"},
		&entity.LineElement{Type: entity.ElementKindBold, Content: "**rofl**"},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementColon(t *testing.T) {
	input := "Hello** smile "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello** smile "},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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

func TestParseLineBoldAltElementColon2(t *testing.T) {
	input := "Hello** smile, blah** testing "
	expectedOutput := []entity.MarkdownElement{
		&entity.LineElement{Type: entity.ElementKindText, Content: "Hello** smile, blah** testing "},
	}
	result := ParseLineBoldAltElement([]entity.MarkdownElement{
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
