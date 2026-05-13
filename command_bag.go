package main

func commandBag(cfg *config, args ...string) error {
	lines := cfg.Styles.FormatBag(cfg.Bag)
	cfg.Styles.DrawBox("Your Bag:", lines)
	return nil
}
