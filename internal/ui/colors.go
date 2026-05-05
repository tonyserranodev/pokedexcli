package ui

func (s Styles) Colorize(text, color string) string {
	return s.Colors[color] + text + s.Colors["re"]
}
