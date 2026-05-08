package main

import "errors"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex.Caught) == 0 {
		return errors.New("your pokedex is empty")
	}

	lines := cfg.Styles.FormatPokedex(cfg.pokedex)
	cfg.Styles.DrawBox("Your Pokedex:", lines)
	return nil
}
