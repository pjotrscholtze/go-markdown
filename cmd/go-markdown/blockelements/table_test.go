package blockelements

import (
	"testing"

	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func TestTableDefinition(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []entity.MarkdownElement
	}{

		{name: "No content", input: "", expect: []entity.MarkdownElement{}},

		{name: "Empty Table", input: `|  |  |
|---|---|
|  |  |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `|  |  |
|---|---|
|  |  |`,
			}}},

		{name: "Single Row, Single Column", input: `| Header |
|--------|
| Value |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Header |
|--------|
| Value |`,
			}}},

		{name: "No Headers", input: `|  |  |
|---|---|
| A | B |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `|  |  |
|---|---|
| A | B |`,
			}}},

		{name: "No Data", input: `| Header 1 | Header 2 |
|----------|----------|`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Header 1 | Header 2 |
|----------|----------|`,
			}}},

		{name: "Colspan", input: `| Column 1 | Column 2 | Column 3 | Column 4 |
|----------|----------|----------|----------|
| Cell 1 | Cell 2 |         | Cell 4 |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 | Column 3 | Column 4 |
|----------|----------|----------|----------|
| Cell 1 | Cell 2 |         | Cell 4 |`,
			}}},

		{name: "Rowspan", input: `| Column 1 | Column 2 | Column 3 |
|----------|----------|----------|
| Cell 1 | Cell 2 |         |
| Cell 4 | Cell 5 |         |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 | Column 3 |
|----------|----------|----------|
| Cell 1 | Cell 2 |         |
| Cell 4 | Cell 5 |         |`,
			}}},

		{name: "Complex Colspan and Rowspan", input: `| Column 1 | Column 2 | Column 3 | Column 4 |
|----------|----------|----------|----------|
| Cell 1 | Cell 2 |         | Cell 4 |
| Cell 5 | Cell 6 |         | Cell 8 |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 | Column 3 | Column 4 |
|----------|----------|----------|----------|
| Cell 1 | Cell 2 |         | Cell 4 |
| Cell 5 | Cell 6 |         | Cell 8 |`,
			}}},

		{name: "Table with Special Characters", input: `| Column 1 | Column 2 |
|----------|----------|
| ~~Test~~ | ` + "`Test`" + `  |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| ~~Test~~ | ` + "`Test`" + `  |`,
			}}},

		{name: "Table with Empty Cells", input: `| Column 1 | Column 2 |
|----------|----------|
|         | Test    |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
|         | Test    |`,
			}}},

		{name: "Table with Non-English Characters", input: `| Column 1 | Column 2 |
|----------|----------|
| ÄÖÜßäöüß | Test    |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| ÄÖÜßäöüß | Test    |`,
			}}},

		{name: "Table with Links", input: `| Column 1 | Column 2 |
|----------|----------|
| [Google](http://google.com) | Test    |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| [Google](http://google.com) | Test    |`,
			}}},

		{name: "Table with Multi-Line Cells", input: `| Column 1 | Column 2 |
|----------|----------|
| Line 1<br>Line 2 | Test    |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| Line 1<br>Line 2 | Test    |`,
			}}},

		{name: "Table with Long Text", input: `| Column 1 | Column 2 |
|----------|----------|
| This is a very long sentence that goes beyond the width of the table. | Test    |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| This is a very long sentence that goes beyond the width of the table. | Test    |`,
			}}},

		{name: "Table with HTML Tags", input: `| Column 1 | Column 2 |
|----------|----------|
| <b>Bold</b> | Test   |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| <b>Bold</b> | Test   |`,
			}}},

		{name: "Table with Different Number of Columns", input: `| Column 1 | Column 2 |
|----------|----------|
| A | B | C|`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| A | B | C|`,
			}}},

		{name: "Table with Duplicate Headers", input: `| Column 1 | Column 1 |
|----------|----------|
| A | B |`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 1 |
|----------|----------|
| A | B |`,
			}}},

		{name: "Table Surrounded by Text", input: `Before the table
| Column 1 | Column 2 |
|----------|----------|
| A | B |
After the table`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `Before the table`,
			},
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `After the table`,
			}}},

		{name: "Table Surrounded by Headings", input: `# Heading Before
| Column 1 | Column 2 |
|----------|----------|
| A | B |
## Heading After`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `# Heading Before`,
			},
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `## Heading After`,
			}}},

		{name: "Table Surrounded by Lists", input: `- Item before
| Column 1 | Column 2 |
|----------|----------|
| A | B |
- Item after`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `- Item before`,
			},
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `- Item after`,
			}}},

		{name: "Table Surrounded by Blockquotes", input: `> Quote before
| Column 1 | Column 2 |
|----------|----------|
| A | B |
> Quote after`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `> Quote before`,
			},
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `> Quote after`,
			}}},

		{name: "Table Surrounded by Links", input: `  [Link before](#before)
| Column 1 | Column 2 |
|----------|----------|
| A | B |
[Link after](#after)`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `  [Link before](#before)`,
			},
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `[Link after](#after)`,
			}}},

		{name: "Table Surrounded by Images", input: `![Image before](#before)
| Column 1 | Column 2 |
|----------|----------|
| A | B |
![Image after](#after)`, expect: []entity.MarkdownElement{
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `![Image before](#before)`,
			},
			&entity.LineElement{
				Type: entity.ElementKindTable,
				Content: `| Column 1 | Column 2 |
|----------|----------|
| A | B |`,
			},
			&entity.LineElement{
				Type:    entity.ElementKindText,
				Content: `![Image after](#after)`,
			}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseLineTableElement([]entity.MarkdownElement{
				&entity.LineElement{Type: entity.ElementKindText, Content: test.input}},
				func(input string) []entity.MarkdownElement {
					return []entity.MarkdownElement{&entity.LineElement{
						Type:    entity.ElementKindText,
						Content: input,
					}}
				})
			gmd := entity.GlueToString(got)
			emd := entity.GlueToString(test.expect)
			_ = gmd
			_ = emd

			if !equalResults(got, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, got)
			}
		})
	}
}
