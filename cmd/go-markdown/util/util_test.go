package util

import (
	"testing"
)

func TestSplitOnNewLine(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []string
	}{
		{
			name:   "No input",
			input:  "",
			expect: []string{},
		},
		{
			name:   "spaces input",
			input:  " ",
			expect: []string{" "},
		},
		{
			name:   "Single line",
			input:  "abcdefghijklmnopqrstuvwxyz",
			expect: []string{"abcdefghijklmnopqrstuvwxyz"},
		},
		{
			name: "Multiple lines",
			input: `abc
def`,
			expect: []string{"abc\n", "def"},
		},
		{
			name: "Multiple lines more",
			input: `abc
def
efg
hij`,
			expect: []string{"abc\n", "def\n", "efg\n", "hij"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SplitOnNewLine(test.input)
			if !equalResultsString(got, test.expect) {
				t.Errorf("Expected %d entries, got %d entries", len(got), len(test.expect))
			}
		})
	}
}

func TestFindPatternsAndNonPatterns(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []Result
	}{
		{
			name:   "No input",
			input:  "",
			expect: []Result{},
		},
		{
			name:   "spaces input",
			input:  " ",
			expect: []Result{{"Non-match", " "}},
		},
		{
			name:   "No matches",
			input:  "abcdefghijklmnopqrstuvwxyz",
			expect: []Result{{"Non-match", "abcdefghijklmnopqrstuvwxyz"}},
		},
		{
			name:   "Single match",
			input:  "abc:123:def",
			expect: []Result{{"Non-match", "abc"}, {"Match", ":123:"}, {"Non-match", "def"}},
		},
		{
			name:   "Multiple matches",
			input:  "abc:123:def:456:ghi",
			expect: []Result{{"Non-match", "abc"}, {"Match", ":123:"}, {"Non-match", "def"}, {"Match", ":456:"}, {"Non-match", "ghi"}},
		},
		{
			name:   "Multiple matches, no nonmatch in between",
			input:  "abc:123::456:ghi",
			expect: []Result{{"Non-match", "abc"}, {"Match", ":123:"}, {"Match", ":456:"}, {"Non-match", "ghi"}},
		},
		{
			name:   "Multiple matches spaces",
			input:  "ab c:123:d ef:456:gh     i",
			expect: []Result{{"Non-match", "ab c"}, {"Match", ":123:"}, {"Non-match", "d ef"}, {"Match", ":456:"}, {"Non-match", "gh     i"}},
		},
		{
			name:   "Single match, nothing after last match",
			input:  "abc:123:",
			expect: []Result{{"Non-match", "abc"}, {"Match", ":123:"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FindPatternsAndNonPatterns(`:[A-Za-z0-9]+:`, test.input, "Match", "Non-match")
			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}

func equalResults(a, b []Result) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func equalResultsString(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
