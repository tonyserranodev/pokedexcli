package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeAPIClient.GetPokemon(name)
	if err != nil {
		return err
	}

	baseExp := pokemon.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if shouldCatch(baseExp) {
		caughtMsg := fmt.Sprintf("%s was caught!", name)
		fmt.Println(cfg.Styles.Colorize(caughtMsg, "green"))

		cfg.pokedex.Caught[name] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}
	failedMsg := fmt.Sprintf(" %s escaped!", name)
	fmt.Println(cfg.Styles.Colorize(failedMsg, "red"))
	return nil
}

func shouldCatch(baseExp int) bool {
	thresh := 50.0
	return rand.Float64()*float64(baseExp) < thresh
}
