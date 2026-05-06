package ui

func (s Styles) Decorate(text, decoration string) string {
	code, ok := s.Decorations[decoration]
	if !ok {
		return text
	}

	return code + text + s.Decorations["re"]
}
