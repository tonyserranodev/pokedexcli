package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeAPIClient pokeapi.Client
	Next          *string
	Previous      *string
}

func startRepl(cfg *config) {

	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands(cfg)
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

func getCommands(cfg *config) map[string]cliCommand {
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
			name:        "explore",
			description: "List pokemon encounters for the given location name",
			callback:    commandExplore,
		},
	}
}
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
