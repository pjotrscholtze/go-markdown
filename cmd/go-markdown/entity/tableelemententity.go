package entity

import (
	"strings"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/util"
)

type tableElementMarkdownElement struct {
	header TableRow
	sep    TableRow
	rows   []TableRow
}
type TableElementMarkdownElement interface {
	AsMarkdownString() string
	Kind() string
	Header() *TableRow
	AddRow(row TableRow)
	Rows() []TableRow
}
type TableRow struct {
	Cells []MarkdownElement
}

func (tr *TableRow) AddCell(cell MarkdownElement) {
	tr.Cells = append(tr.Cells, cell)
}

func (tr *TableRow) AsMarkdownString() string {
	items := []string{}
	for _, cell := range tr.Cells {
		items = append(items, cell.AsMarkdownString())
	}
	return strings.Join(items, "|")
}

func parseTableRow(input string, parserFn func(input string) []MarkdownElement) TableRow {
	cells := []MarkdownElement{}
	for _, cell := range strings.Split(input, "|") {
		cells = append(cells, &GroupElement{Contents: parserFn(cell)})
	}
	return TableRow{
		Cells: cells,
	}
}

func (bqme *tableElementMarkdownElement) Kind() string {
	return ElementKindTable
}
func (bqme *tableElementMarkdownElement) AsMarkdownString() string {
	content := []string{bqme.header.AsMarkdownString(), bqme.sep.AsMarkdownString()}
	for _, row := range bqme.rows {
		content = append(content, row.AsMarkdownString())
	}
	return strings.Join(content, "")
}
func (bqme *tableElementMarkdownElement) Header() *TableRow {
	return &bqme.header
}
func (bqme *tableElementMarkdownElement) AddRow(row TableRow) {
	bqme.rows = append(bqme.rows, row)
}
func (bqme *tableElementMarkdownElement) Rows() []TableRow {
	return bqme.rows
}

func NewTableElementMarkdownElement(input string, parserFn func(input string) []MarkdownElement) TableElementMarkdownElement {
	// @todo properly parse the table here, then parseFn can be used...
	lines := util.SplitOnNewLine(input)
	rows := []TableRow{}
	if len(lines) > 2 {
		for _, line := range lines[2:] {
			rows = append(rows, parseTableRow(line, parserFn))
		}
	}

	header := TableRow{}
	sep := TableRow{}
	if len(lines) > 0 {
		header = parseTableRow(lines[0], parserFn)
	}
	if len(lines) > 1 {
		sep = parseTableRow(lines[1], parserFn)
	}

	return &tableElementMarkdownElement{
		rows:   rows,
		header: header,
		sep:    sep,
	}
}
