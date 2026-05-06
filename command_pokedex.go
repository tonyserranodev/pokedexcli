package main

func commandPokedex(cfg *config, args ...string) error {
	lines := cfg.Styles.FormatPokedex(cfg.pokedex)
	cfg.Styles.DrawBox("Your Pokedex:", lines)
	return nil
}
