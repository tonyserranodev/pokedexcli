package main

import (
	"time"

	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
	"github.com/tonyserranodev/pokedexcli/internal/trainer"
	"github.com/tonyserranodev/pokedexcli/internal/ui"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeAPIClient: pokeClient,
		pokedex:       trainer.NewPokedex(),
		PC:            trainer.NewPC(),
		Party:         trainer.NewParty(),
		Bag: map[string]int{
			"pokeball":   10,
			"greatball":  10,
			"masterball": 10,
			"ultraball":  10,
		},
		VisitedLocations: make(map[string]struct{}),
		Styles:           ui.NewStyles(),
	}
	startRepl(cfg)
}
