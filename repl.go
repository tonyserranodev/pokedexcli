package main

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"

	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
	"github.com/tonyserranodev/pokedexcli/internal/trainer"
	"github.com/tonyserranodev/pokedexcli/internal/ui"
)

type config struct {
	Next             *string
	Previous         *string
	VisitedLocations map[string]struct{}
	pokeAPIClient    pokeapi.Client
	pokedex          trainer.Pokedex
	Party            []pokeapi.Pokemon
	PC               map[string]pokeapi.Pokemon
	Bag              map[string]int
	Styles           ui.Styles
	RL               *readline.Instance
}

func startRepl(cfg *config) {

	fmt.Println(pokemonTitle)
	commands := getCommands()
	l, err := readline.NewEx(&readline.Config{
		Prompt:          cfg.Styles.Render("Pokedex > ", "blue", "bold"),
		HistoryFile:     "/tmp/pokedex_history.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	// save readline instance in config for use in sub-menus
	cfg.RL = l

	for {
		updateCompleter(cfg)

		line, err := l.Readline()
		if err != nil {
			break
		}

		words := cleanInput(line)

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

		err = cmd.callback(cfg, args...)
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
		"party": {
			name:        "party",
			description: "List all pokemon in your party",
			callback:    commandParty,
		},
		"release": {
			name:        "release",
			description: "Release a member from your party",
			callback:    commandRelease,
		},
		"pc": {
			name:        "pc",
			description: "View pokemon in your PC",
			callback:    commandPC,
		},
		"bag": {
			name:        "bag",
			description: "View all items in your bag",
			callback:    commandBag,
		},
		"clear": {
			name:        "clear",
			description: "Clear the screen",
			callback:    commandClear,
		},
	}
}

func updateCompleter(cfg *config) {
	var items []readline.PrefixCompleterInterface

	// Build location items from VisitedLocations
	var locationItems []readline.PrefixCompleterInterface
	for name := range cfg.VisitedLocations {
		locationItems = append(locationItems, readline.PcItem(name))
	}

	// Build pokemon items from Pokedex
	var pokemonItems []readline.PrefixCompleterInterface
	for _, p := range cfg.pokedex.Caught {
		pokemonItems = append(pokemonItems, readline.PcItem(p.Name))
	}

	// assemble command tree
	items = append(items,
		readline.PcItem("explore", locationItems...),
		readline.PcItem("inspect", pokemonItems...),
		readline.PcItem("map"),
		readline.PcItem("mapb"),
		readline.PcItem("exit"),
		readline.PcItem("clear"),
	)
	cfg.RL.Config.AutoComplete = readline.NewPrefixCompleter(items...)
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
