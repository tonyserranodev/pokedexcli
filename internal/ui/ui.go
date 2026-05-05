// Package ui provides methods for styling pokedexcli
package ui

type Styles struct {
	Colors      map[string]string
	Decorations map[string]string
}

func NewStyles() Styles {
	return Styles{
		Colors: map[string]string{
			"re":      "\033[0m",
			"red":     "\033[31m",
			"green":   "\033[32m",
			"yellow":  "\033[33m",
			"blue":    "\033[34m",
			"magenta": "\033[35m",
			"cyan":    "\033[36m",
			"white":   "\033[37m",
		},
		Decorations: map[string]string{
			"bold":      "\033[22m",
			"italic":    "\033[23m",
			"underline": "\033[24m",
			"strike":    "\033[25m",
		},
	}
}
