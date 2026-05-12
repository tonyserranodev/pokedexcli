package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
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

	if !shouldCatch(baseExp) {
		failedMsg := fmt.Sprintf(" %s escaped!", name)
		fmt.Println(cfg.Styles.Colorize(failedMsg, "red"))
		return nil
	}

	caughtMsg := fmt.Sprintf("%s was caught!", name)
	fmt.Println(cfg.Styles.Colorize(caughtMsg, "green"))
	cfg.pokedex.Caught[name] = pokemon

	if len(cfg.Party) < 6 {
		cfg.Party = append(cfg.Party, name)
		fmt.Printf("%s was added to your party!\n", name)
	} else {

		fmt.Println("Your party is full. Release a pokemon from your party? (y/n)")
		line, _ := cfg.RL.Readline()
		words := cleanInput(line)

		if len(words) > 0 && words[0] == "y" {
			handlePartySwap(cfg, name)
		} else {
			fmt.Printf("%s was sent to your pokedex\n", name)
		}
	}

	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}

func shouldCatch(baseExp int) bool {
	thresh := 50.0
	return rand.Float64()*float64(baseExp) < thresh
}

func handlePartySwap(cfg *config, newPokemonName string) error {
	fmt.Println("Select a party member to release")
	for i, name := range cfg.Party {
		fmt.Printf("%d. %s\n", i+1, name)
	}

	line, err := cfg.RL.Readline()
	if err != nil {
		return err
	}

	words := cleanInput(line)
	if len(words) == 0 {
		return nil
	}

	opt := words[0]
	choice, err := strconv.Atoi(opt)
	if err != nil || choice < 1 || choice > len(cfg.Party) {
		fmt.Printf("Invalid choice. No changes made to your party")
	}

	index := choice - 1
	oldName := cfg.Party[index]
	cfg.Party[index] = newPokemonName

	fmt.Printf("Bye %s! %s joined your party\n", oldName, newPokemonName)
	return nil
}
