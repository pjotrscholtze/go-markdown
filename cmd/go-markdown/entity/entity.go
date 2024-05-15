package entity

type MarkdownElement interface {
	AsMarkdownString() string
	Kind() string
}

type GroupElement struct {
	Contents []MarkdownElement
}

func (ge *GroupElement) Kind() string {
	return ElementKindGroup
}
func (ge *GroupElement) AsMarkdownString() string {
	return GlueToString(ge.Contents)
}

type LineElement struct {
	Type    string
	Content string
}

func (le *LineElement) Kind() string {
	return le.Type
}
func (le *LineElement) AsMarkdownString() string {
	return le.Content
}

func GlueToString(contents []MarkdownElement) string {
	out := ""
	for _, entry := range contents {
		out += entry.AsMarkdownString()
	}
	return out
}

// func (le *LineElement) AsMarkdownEntity() MarkdownElement {
// 	switch le.Type {
// 	case ElementKindBold:
// 		return NewBoldMarkdownElement(le.Content)

// 	case ElementKindCheckbox:
// 		return NewCheckboxMarkdownElement(le.Content)

// 	case ElementKindEmoji:
// 		return NewEmojiMarkdownElement(le.Content)

// 	case ElementKindFootnote:
// 		return NewFootnoteMarkdownElement(le.Content)

// 	case ElementKindHighlight:
// 		return NewHighlightMarkdownElement(le.Content)

// 	case ElementKindImage:
// 		return NewImageMarkdownElement(le.Content)

// 	case ElementKindItalic:
// 		return NewItalicMarkdownElement(le.Content)

// 	case ElementKindLink:
// 		return NewLinkMarkdownElement(le.Content)

// 	case ElementKindStrikethrough:
// 		return NewStrikethroughMarkdownElement(le.Content)

// 	case ElementKindInlineCode:
// 		return NewInlineCodeMarkdownElement(le.Content)

// 	case ElementKindText:
// 		md := NewRawTextMarkdownElement(le.Content)
// 		return &md

// 	case ElementKindHeader:
// 		return NewHeaderMarkdownElement(le.Content)

// 	case ElementKindBlockquote:
// 		return NewBlockQuoteMarkdownElement(le.Content)

// 	case ElementKindTable:
// 		return NewTableElementMarkdownElement(le.Content)

// 	case ElementKindList:
// 		return NewListElementMarkdownElement(le.Content)

// 	case ElementKindCodeblock:
// 		return NewCodeBlockMarkdownElement(le.Content)

// 	case ElementKindHorizontalLine:
// 		return NewHorizontalLineMarkdownElement(le.Content)

// 	case ElementKindTermDefinitionLine:
// 		return NewTermDefinitionElementMarkdownElement(le.Content)

// 	}
// 	return nil
// }
