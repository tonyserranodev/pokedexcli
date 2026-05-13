package main

import "errors"

func commandParty(cfg *config, args ...string) error {

	if len(cfg.Party) == 0 {
		return errors.New("your party is empty")
	}

	lines := cfg.Styles.FormatParty(cfg.Party)
	cfg.Styles.DrawBox("Your Party:", lines)
	return nil
}
