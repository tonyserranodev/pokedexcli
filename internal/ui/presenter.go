package ui

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
)

func (s Styles) DrawBox(title string, lines []string) {
	width := utf8.RuneCountInString(title)
	for _, line := range lines {
		if utf8.RuneCountInString(line) > width {
			width = utf8.RuneCountInString(line)
		}
	}

	title = s.Render(capitalize(title), "green", "bold", "italic")
	fmt.Printf("┌ %s %s┐\n", title, strings.Repeat("-", width-visualWidth(title)))

	for _, line := range lines {
		padding := strings.Repeat(" ", width-visualWidth(line))

		fmt.Printf("| %s%s |\n", line, padding)
	}

	fmt.Printf("└%s┘\n", strings.Repeat("-", width+2))
}

func (s Styles) FormatPokemonInspect(p pokeapi.Pokemon) []string {
	lines := []string{
		fmt.Sprintf("%s: %d", s.Colorize("Height", "yellow"), p.Height),
		fmt.Sprintf("%s: %d", s.Colorize("Weight", "yellow"), p.Weight),
		"",
		s.Colorize("Stats:", "cyan"),
	}

	maxNameLen := 0
	for _, s := range p.Stats {
		if len(s.Stat.Name) > maxNameLen {
			maxNameLen = len(s.Stat.Name)
		}
	}

	fString := fmt.Sprintf(" %%-%ds- %%d", maxNameLen+1)
	for _, stat := range p.Stats {
		lines = append(lines, fmt.Sprintf(fString, s.Colorize(stat.Stat.Name, "yellow"), stat.BaseStat))
	}

	lines = append(lines, "", s.Colorize("Types:", "cyan"))
	for _, t := range p.Types {
		lines = append(lines, fmt.Sprintf(" - %s", s.Colorize(t.Type.Name, "yellow")))
	}
	return lines
}

func (s Styles) FormatPokedex(pokedex *pokeapi.Pokedex) []string {
	lines := []string{}
	for _, p := range pokedex.Caught {
		lines = append(lines, s.Colorize(fmt.Sprintf(" %s", capitalize(p.Name)), "yellow"))
	}
	return lines
}

func (s Styles) Render(text string, color string, decorations ...string) string {
	out := text

	for _, d := range decorations {
		if code, ok := s.Decorations[d]; ok {
			out = code + out
		}
	}

	if code, ok := s.Colors[color]; ok {
		out = code + out
	}
	return out + s.Colors["re"]
}

func visualWidth(text string) int {
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	plain := ansiRegex.ReplaceAllString(text, "")
	return utf8.RuneCountInString(plain)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
