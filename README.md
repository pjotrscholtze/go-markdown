# Go-markdown
A Markdown Parser in Go, this repository contains a Go library for parsing
Markdown text and allowing you to manipulate a markdown tree in Go. It's
designed to be fast, thread-safe, and minimal in dependencies, making it
suitable for a wide range of applications.

## Features
- **Extendability**: The parser is easily extendable, since a parse tree can be
                     given, allowing easy incorporation of custom elements.
- **Performance**: Fast enough to render on-demand in most web applications
                   without having to cache the output.
- **Thread Safety**: Multiple parsers can run in different goroutines without
                     ill effect.
- **Minimal Dependencies**: Only depends on standard library packages in Go.
- **Large set of elements**: By default a large set of markdown elements can be
                             parsed allowing most documents to be parsed.

## Usage

To use this library in your Go project, follow these steps:

1. Install the library:
```
go get -u github.com/pjotrscholtze/go-markdown
```

2. Use the library in your code:
```
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
```


## Technical Documentation
Technical documentation can be found in the `docs` folder.

## Contributing
Contributions are welcome Please feel free to submit pull requests or report
issues.

## License
This project is licensed under the MIT License.
