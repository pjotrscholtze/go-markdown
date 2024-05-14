package entity

import (
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

type tableElementMarkdownElement struct {
	Content string
}
type TableElementMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	Header() *TableRow
	AddRow(row TableRow)
	Rows() []TableRow
}
type TableRow struct {
	Cells []string
}

func (tr *TableRow) AddCell(cell string) {
	tr.Cells = append(tr.Cells, cell)
}

func (tr *TableRow) AsMarkdownString() string {
	return strings.Join(tr.Cells, "|")
}

func parseTableRow(input string) TableRow {
	return TableRow{
		Cells: strings.Split(input, "|"),
	}
}

func (bqme *tableElementMarkdownElement) Kind() string {
	return ElementKindTable
}
func (bqme *tableElementMarkdownElement) AsMarkdownString() string {
	return bqme.Content
}
func (bqme *tableElementMarkdownElement) Header() *TableRow {
	lines := util.SplitOnNewLine(bqme.Content)
	if len(lines) == 0 {
		return nil
	}
	tr := parseTableRow(lines[0])
	return &tr
}
func (bqme *tableElementMarkdownElement) AddRow(row TableRow) {
	bqme.Content += "\n" + row.AsMarkdownString()
}
func (bqme *tableElementMarkdownElement) Rows() []TableRow {
	lines := util.SplitOnNewLine(bqme.Content)
	if len(lines) < 3 {
		return nil
	}
	rows := []TableRow{}
	for _, line := range lines[2:] {
		rows = append(rows, parseTableRow(line))
	}
	return rows
}

func NewTableElementMarkdownElement(input string, parserFn func(input string) []MarkdownElement) TableElementMarkdownElement {
	// @todo properly parse the table here, then parseFn can be used...
	return &tableElementMarkdownElement{
		Content: input,
	}
}
