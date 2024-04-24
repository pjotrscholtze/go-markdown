package entity

import (
	"regexp"
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

type subListElementMarkdownElement struct {
	MainContent   []MarkdownElement
	SubListOffset string
	Symbol        string
	SubList       []MarkdownElement
}

func (bqme *subListElementMarkdownElement) Kind() string {
	return "subListElement"
}
func (bqme *subListElementMarkdownElement) AsMarkdownString() string {
	out := []string{}
	for _, entry := range bqme.SubList {
		out = append(out, entry.AsMarkdownString())
	}
	outStr := strings.Join(out, "")
	before := 0
	r := regexp.MustCompile("\r?\n") //.Split(inputString, -1)
	matches := r.FindAllIndex([]byte(outStr), -1)
	out = []string{}
	for _, match := range matches {
		line := outStr[before:match[1]]
		before = match[1]
		out = append(out, bqme.SubListOffset+line)
	}
	line := outStr[before:]
	out = append(out, bqme.SubListOffset+line)

	return bqme.Symbol + GlueToString(bqme.MainContent) + strings.Join(out, "")
}

type listElementMarkdownElement struct {
	Content []MarkdownElement
}
type ListElementMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

func (bqme *listElementMarkdownElement) Kind() string {
	return ElementKindList
}
func (bqme *listElementMarkdownElement) AsMarkdownString() string {
	return GlueToString(bqme.Content)
}

func findListItemSections(input string) []string {
	r := regexp.MustCompile("\r?\n") //.Split(inputString, -1)
	matches := r.FindAllIndex([]byte(input), -1)

	preLines := []string{}
	entries := []string{}
	before := 0
	for _, match := range matches {
		line := input[before:match[1]]
		before = match[1]
		if entry := util.FindPatternsAndNonPatternsSingleLine(`^([\d\-\*\.\+]+ )+.*`, line, "match", "")[0]; entry.Type == "match" {
			if len(preLines) > 0 {
				entries = append(entries, strings.Join(preLines, ""))
			}
			preLines = []string{line}
		} else {
			preLines = append(preLines, line)
		}

	}
	line := input[before:]
	if entry := util.FindPatternsAndNonPatternsSingleLine(`^([\d\-\*\.\+]+ )+.*$`, line, "match", "")[0]; entry.Type == "match" {
		if len(preLines) > 0 {
			entries = append(entries, strings.Join(preLines, ""))
		}
		preLines = []string{entry.Content}
	} else {
		preLines = append(preLines, entry.Content)
	}
	if len(preLines) > 0 {
		entries = append(entries, strings.Join(preLines, ""))
	}

	return entries
}
func getListItemSymbol(input string) (string, string) {
	r := regexp.MustCompile(`^([\d\-\*\.\+]+ )+`)
	match := r.FindAllIndex([]byte(input), -1)[0]
	symbol := input[:match[1]]
	content := input[match[1]:]
	return symbol, content
}

func parseSingleListItem(input string, parserFn func(input string) []MarkdownElement, selfFn func(input []MarkdownElement, parserFn func(input string) []MarkdownElement) []MarkdownElement) *subListElementMarkdownElement {
	symbol, content := getListItemSymbol(input)

	_ = symbol
	_ = content

	r := regexp.MustCompile("\r?\n")
	matches := r.FindAllStringSubmatchIndex(content, -1)
	before := 0
	if len(matches) > 1 {
		matches = matches[:len(matches)-1]
	}
	rWhitespace := regexp.MustCompile(`^\s+`)
	symbolLength := len(symbol)
	// sections := []string{}
	preLines := []string{}
	// out := listElementMarkdownElement{Content: []MarkdownElement{}}
	candidateSections := []string{}
	subStarted := false
	whitespace := ""
	for i, match := range matches {
		line := content[before:match[1]]
		whiteSpaceMatches := rWhitespace.FindAllIndex([]byte(line), -1)
		before = match[1]

		// Case 1: there is no whitespace at all merge with prev.
		// Case 2: there is white space merge with prev, except when there is a
		//         sublist.
		if i == 0 {
			preLines = append(preLines, line)
			continue
		}
		if !subStarted && (len(whiteSpaceMatches) > 0 && whiteSpaceMatches[0][1] == symbolLength) {
			candidateSections = append(candidateSections, strings.Join(preLines, ""))
			whitespace = line[:whiteSpaceMatches[0][1]]
			preLines = []string{}
			subStarted = true
		}
		preLines = append(preLines, line)
		// sections = append(sections, whiteSpaceMatches)
		_ = line
		_ = whiteSpaceMatches
	}
	preLines = append(preLines, content[before:])
	candidateSections = append(candidateSections, strings.Join(preLines, ""))

	// line := content[before:]
	stripped := ""
	if len(candidateSections) > 1 {
		subMatches := r.FindAllStringSubmatchIndex(candidateSections[1], -1)
		before = 0
		for _, match := range subMatches {
			line := candidateSections[1][before:match[1]]
			before = match[1]
			stripped += line[symbolLength:]
		}
	}
	strippedElement := RawTextMarkdownElement(stripped)
	_ = strippedElement
	// blockelements.ParseLineListElement([]MarkdownElement{
	// 	&strippedElement,
	// }, parserFn)

	return &subListElementMarkdownElement{
		Symbol:        symbol,
		MainContent:   parserFn(candidateSections[0]),
		SubListOffset: whitespace,
		SubList:       selfFn([]MarkdownElement{&strippedElement}, parserFn),
	}
}

func NewListElementMarkdownElement(input string, parserFn func(input string) []MarkdownElement, selfFn func(input []MarkdownElement, parserFn func(input string) []MarkdownElement) []MarkdownElement) ListElementMarkdownElement {
	content := []MarkdownElement{}
	for _, section := range findListItemSections(input) {
		content = append(content, parseSingleListItem(section, parserFn, selfFn))
		//
	}
	return &listElementMarkdownElement{
		Content: content,
	}
}
