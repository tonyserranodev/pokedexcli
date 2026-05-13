package main

import "errors"

func commandPC(cfg *config, args ...string) error {

	if len(cfg.PC) == 0 {
		return errors.New("your PC is empty")
	}

	//TODO implement paging through results if pc length is certain length
	lines := cfg.Styles.FormatPC(cfg.PC)
	cfg.Styles.DrawBox("Your PC:", lines)
	return nil
}
