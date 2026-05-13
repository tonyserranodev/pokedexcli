package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"

	"github.com/tonyserranodev/pokedexcli/internal/items"
	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	// name: args[0]
	// ball type: args[1]

	if len(args) != 2 {
		return errors.New("you must provide a pokemon name and a ball name")
	}

	allBalls := items.GetAvailableBalls()
	ballName := args[1]
	ball, ok := allBalls[ballName]
	if !ok {
		return errors.New("that ball doesnt exist")
	}

	if cfg.Bag[ballName] <= 0 {
		return fmt.Errorf("you dont have any %ss", ballName)
	}

	name := args[0]
	pokemon, err := cfg.pokeAPIClient.GetPokemon(name)
	if err != nil {
		return err
	}

	baseExp := pokemon.BaseExperience
	mult := ball.Multiplier

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	cfg.Bag[ballName]--
	if !shouldCatch(baseExp, mult) {
		failedMsg := fmt.Sprintf(" %s escaped!", name)
		fmt.Println(cfg.Styles.Colorize(failedMsg, "red"))
		return nil
	}

	caughtMsg := fmt.Sprintf("%s was caught!", name)
	fmt.Println(cfg.Styles.Colorize(caughtMsg, "green"))
	cfg.pokedex.Caught[name] = pokemon

	if len(cfg.Party) < 6 {
		cfg.Party = append(cfg.Party, pokemon)
		fmt.Printf("%s was added to your party!\n", name)
	} else {

		fmt.Println("Your party is full. Release a pokemon from your party? (y/n)")
		line, _ := cfg.RL.Readline()
		words := cleanInput(line)

		if len(words) > 0 && words[0] == "y" {
			handlePartySwap(cfg, pokemon)
		} else {
			fmt.Printf("%s was sent to your pokedex\n", name)
		}
	}

	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}

func shouldCatch(baseExp int, mult float64) bool {
	fmt.Printf("base exp: %d\n", baseExp)
	thresh := 35.0
	return rand.Float64()*float64(baseExp) < (thresh * mult)
}

func handlePartySwap(cfg *config, newPokemon pokeapi.Pokemon) error {
	fmt.Println("Select a party member to release")
	for i, pokemon := range cfg.Party {
		fmt.Printf("%d. %s\n", i+1, pokemon.Name)
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
	oldName := (cfg.Party)[index].Name
	(cfg.Party)[index] = newPokemon
	(cfg.PC)[newPokemon.Name] = newPokemon

	fmt.Printf("%s has been sent to your PC. %s has joined your party!\n", oldName, newPokemon.Name)
	return nil
}
