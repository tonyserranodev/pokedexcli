package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Next     *string
	Previous *string
}

func getCommands(cfg *Config) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List 20 pokemon locations. Call map again to list the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List 20 pokemon locations. Call mapb again to list the previous 20 locations",
			callback:    commandMapBack,
		},
	}
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands(cfg) {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	res, err := pokeapi.GetLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = res.Next
	cfg.Previous = res.Previous
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := pokeapi.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = res.Next
	cfg.Previous = res.Previous
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
