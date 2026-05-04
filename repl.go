package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeAPIClient pokeapi.Client
	pokedex       *pokeapi.Pokedex
	Next          *string
	Previous      *string
}

func startRepl(cfg *config) {

	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")

		reader.Scan()
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
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
		"explore": {
			name:        "explore <location>",
			description: "List pokemon encounters for the given location name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch a pokemon. The higher the pokemon experience the more difficult it will be to catch",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "List information for the given pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemon in your pokedex",
			callback:    commandPokedex,
		},
		"clear": {
			name:        "clear",
			description: "Clear the screen",
			callback:    commandClear,
		},
	}
}
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
