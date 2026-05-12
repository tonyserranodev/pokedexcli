package main

import (
	"fmt"
	"slices"
	"strconv"
)

func commandRelease(cfg *config, args ...string) error {
	fmt.Println("Select a party member to release")
	for i, pokemon := range *cfg.Party {
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
	if err != nil || choice < 1 || choice > len(*cfg.Party) {
		fmt.Printf("Invalid choice. No changes made to your party")
	}

	index := choice - 1
	pokemon := (*cfg.Party)[index]
	*cfg.Party = slices.Delete(*cfg.Party, index, index+1)
	(*cfg.PC)[pokemon.Name] = pokemon

	fmt.Printf("%s has been sent to your pc!\n", pokemon.Name)
	return nil
}
