package entity

type MarkdownElement interface {
	AsMarkdownString() string
	Kind() string
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
