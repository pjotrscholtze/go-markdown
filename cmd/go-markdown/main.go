// package main

// import (
// 	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/mockdata"
// )

// func main() {
// 	print(mockdata.MockData1)
// 	print(mockdata.MockData2)
// 	print(mockdata.MockData3)
// 	print(mockdata.MockData4)
// 	print(mockdata.MockData5)
// 	// input := "abc:123:def:456:ghi"
// 	// pattern := regexp.MustCompile(`:[A-Za-z0-9]+:`)
// 	// matches := pattern.FindAllStringSubmatchIndex(input, -1)

// 	// for i := 0; i < len(matches); i++ {
// 	// 	start, end := matches[i][0], matches[i][1]
// 	// 	if i > 0 {
// 	// 		prevEnd := matches[i-1][1]
// 	// 		if start > prevEnd {
// 	// 			fmt.Printf("Non-match: %s\n", input[prevEnd:start])
// 	// 		}
// 	// 	}
// 	// 	fmt.Printf("Match: %s\n", input[start:end])
// 	// }

// 	// if len(matches) > 0 && matches[len(matches)-1][1] < len(input) {
// 	// 	fmt.Printf("Non-match: %s\n", input[matches[len(matches)-1][1]:])
// 	// }
// }

package main

import (
	"fmt"
	"strings"
)

func extractTables(input string) []string {
	lines := strings.Split(input, "\n")
	var tables []string
	var table []string
	var preLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, "|") {
			table = append(table, line)
		} else if len(table) > 0 {
			tables = append(tables, strings.Join(preLines, "\n"))
			tables = append(tables, strings.Join(table, "\n"))
			preLines = []string{line}
			table = nil
		} else {
			preLines = append(preLines, line)
		}
	}
	if len(table) > 0 {
		tables = append(tables, strings.Join(preLines, "\n"))
		tables = append(tables, strings.Join(table, "\n"))
	}
	if len(preLines) > 0 {
		tables = append(tables, strings.Join(preLines, "\n"))
	}
	return tables
}

func main() {
	input := `Before table:

| Column 1 | Column 2 | Column 3 |
|----------|:---------:|----------:|
| Cell 1 | Cell 2 | Cell 3 |
| Cell 4 | Cell 5 | Cell 6<br/>Second line of text |

After first table but before second table:

Another table:

| Column A | Column B | Column C |
|----------|:---------:|----------:|
| Cell 1 | Cell 2 | Cell 3 |
| Cell 4 | Cell 5 | Cell 6<br/>Second line of text |

After second table:`
	tables := extractTables(input)
	for i, table := range tables {
		fmt.Printf("Table %d:\n%s\n", i+1, table)
	}
}
