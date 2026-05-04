package main

import (
	"errors"
	"fmt"
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

	height := pokemon.Height
	weight := pokemon.Weight
	stats := pokemon.Stats
	types := pokemon.Types

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", height)
	fmt.Printf("Weight: %d\n", weight)
	fmt.Println("Stats:")

	for _, s := range stats {
		fmt.Printf("  %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	for _, t := range types {
		fmt.Println("Types:")
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	return nil

}
