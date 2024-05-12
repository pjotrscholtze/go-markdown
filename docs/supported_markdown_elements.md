# Supported Markdown elements
This document serves as an overview of the various Markdown elements that are
supported within the parser. Each section below delves into the syntax, usage,
and examples for a wide range of Markdown features. Not all Markdown elements
are supported by Github, so some will not render properly on Github.

## Blockquote element
The blockquote element allows you to highlight a section of text as a quotation.
It's a great way to emphasize a statement or thought within your document.
```
> This is a block quote.
```

### Example
To make text bold, wrap it with double asterisks (`**`) or double underscores
(`__`). This is useful for drawing attention to important points or headings.
> Block quote here

## Boldmarkdown element
```
**Bold text**
__Also bold__
```

### Example
**Bold text**
__Also bold__

## Checkbox element
Checkboxes can be added using square brackets (`[ ]`) for unchecked and `x` for
checked, other symbols can be used to, but this is limited to a single symbol
per checkbox.
```
[ ] Not checked checkbox
[x] Checked checkbox
```

### Example
[ ] Not checked checkbox
[x] Checked checkbox

## Codeblock element
For displaying code snippets, enclose them between triple backticks (` ``` `).
```
` ` `
Code here
` ` `
```

### Example
```
Code here
```

## Emoji element
Emojis can be inserted directly into your text using their shortcode, such as
`:smile:` for a smiley face emoji. This adds a fun and visual element to your
content.
```
:happy:
```

### Example
:happy:

## Footnote element
Footnotes allow you to add additional information or references at the bottom of
your document. They are marked with a caret (`^`) followed by the footnote
number.
```
Footnote example in text [^1]

[^1] Refered to here
```

### Example
Footnote example in text [^1]

[^1] Refered to here

## Headers element
Headers are used to structure your document with different levels of importance.
Use the hash symbol (`#`) followed by one or more spaces for headers.
```
# H1
## H2
####### H7
```

### Example
# H1
## H2
####### H7

## Highlight element
To highlight text, wrap it with double equals signs (`==`). This is useful for
drawing attention to key points or terms.
```
Some text ==that is highlighted== in the middle.
```

### Example
Some text ==that is highlighted== in the middle.

## Horizontal line element
Horizontal lines can be added using various combinations of hyphens, asterisks,
or underscores. This creates a clear separation in your document.
```
---
----
----------
___
____
__________
***
****
**********
```

### Example
---
----
----------
___
____
__________
***
****
**********

## Image element
Images can be inserted with alternative text and a link to the image file. Use
the exclamation mark (`~`) followed by square brackets for the alt text and
parentheses for the image URL.
```
![Alt text here](image.jpg)
![](image.jpg)
![Alt text here](https://example.com/image.jpg)
![Alt text here](./images/image.jpg \"example title\")
```

### Example
![Alt text here](image.jpg)
![](image.jpg)
![Alt text here](https://example.com/image.jpg)
![Alt text here](./images/image.jpg \"example title\")

## Inline-code element
For inline code snippets, wrap the code with single backticks (`` ` ``). This
maintains the formatting of the code within a sentence.
```
Text here `some code here` more text here.
```

### Example
Text here `some code here` more text here.

## Italic element
To italicize text, wrap it with single asterisks (`*`) or single underscores
(`_`). This is useful for emphasizing words or phrases.
```
*smile*
_smile_
```

### Example
*smile*
_smile_

## Link element
Links can be created using square brackets for the link text and parentheses for
the URL. Optionally, you can add a title attribute for the link.
```
[Google](http://google.com)
[an example](http://example.com/ \"Title\")
```

### Example
[Google](http://google.com)
[an example](http://example.com/ \"Title\")

## List element
Lists can be created using either hyphens, plus signs, or numbers followed by
periods. Sublists can be nested using indentation.
```
- Item here
    - sub item here
- Item here again.

+ Different symbols can be used too.
    + again with subitems
+ And goin further with list items here
    - Mixing these the symbols used for nested lists is also supported.

1. Numbered list
    1.2. With sub items
```

### Example
- Item here
    - sub item here
- Item here again.

+ Different symbols can be used too.
    + again with subitems
+ And goin further with list items here
    - Mixing these the symbols used for nested lists is also supported.

1. Numbered list
    1.2. With sub items

## Strikethrough element
To strike through text, wrap it with double tildes (`~~`). This is useful for
indicating revisions or deletions.
```
Some ~~ strikethrough ~~ text
```

### Example
Some ~~ strikethrough ~~ text

## Table element
Tables can be created using pipes (`|`) to separate columns and dashes (`-`) to
create headers. This is ideal for organizing data or information.
```
| Header 1 | Header 2 |
|----------|----------|
| Cell     | content  |
| Multi    | row      |
```

### Example
| Header 1 | Header 2 |
|----------|----------|
| Cell     | content  |
| Multi    | row      |

## Termdefinition element
Term definitions can be added using the caret (`^`) followed by the term and its
definition. This is useful for glossaries or explanations.
```
^: Term definition text here.
```

### Example
^: Term definition text here.
