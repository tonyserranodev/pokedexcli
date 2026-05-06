package ui

func (s Styles) Colorize(text, color string) string {
	code, ok := s.Colors[color]
	if !ok {
		return text
	}

	return code + text + s.Colors["re"]
}
