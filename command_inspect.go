package main

import (
	"errors"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, exists := cfg.pokedex.Caught[name]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	lines := cfg.Styles.FormatPokemon(pokemon)

	cfg.Styles.DrawBox(pokemon.Name, lines)
	return nil

}
