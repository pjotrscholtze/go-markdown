package main

import (
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/parser"
)

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
	doc := parser.ParseString(input, parser.DefaultMarkdownOrder)
	_ = doc
}
