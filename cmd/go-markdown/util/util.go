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
