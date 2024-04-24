package util

import (
	"regexp"
)

type Result struct {
	Type    string
	Content string
}

func FindPatternsAndNonPatterns(regex, input, matchType, nonMatchType string) []Result {
	results := make([]Result, 0)
	r := regexp.MustCompile("\r?\n")
	matches := r.FindAllStringSubmatchIndex(input, -1)
	before := 0
	if len(matches) > 1 {
		matches = matches[:len(matches)-1]
	}
	for _, match := range matches {
		line := input[before:match[1]]
		before = match[1]
		results = append(results, FindPatternsAndNonPatternsSingleLine(regex, line, matchType, nonMatchType)...)
	}
	line := input[before:]
	results = append(results, FindPatternsAndNonPatternsSingleLine(regex, line, matchType, nonMatchType)...)

	return results //findPatternsAndNonPatterns(regex, input, matchType, nonMatchType)
}

func FindPatternsAndNonPatternsSingleLine(regex, input, matchType, nonMatchType string) []Result {
	results := make([]Result, 0)
	if len(input) == 0 {
		return results
	}
	pattern := regexp.MustCompile(regex)
	matches := pattern.FindAllStringSubmatchIndex(input, -1)

	prev := 0
	for _, match := range matches {
		start, end := match[0], match[1]
		if len(input[prev:start]) > 0 {
			results = append(results, Result{
				Type:    nonMatchType,
				Content: input[prev:start],
			})
		}
		results = append(results, Result{
			Type:    matchType,
			Content: input[start:end],
		})
		prev = end
	}
	if prev < len(input) {
		results = append(results, Result{
			Type:    nonMatchType,
			Content: input[prev:],
		})
	}

	return results
}
