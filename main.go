package main

import (
	"time"

	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
	"github.com/tonyserranodev/pokedexcli/internal/ui"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	pokedex := pokeapi.NewPokedex()
	cfg := &config{
		pokeAPIClient:    pokeClient,
		pokedex:          pokedex,
		Party:            *pokeapi.NewParty(),
		VisitedLocations: make(map[string]struct{}),
		Styles:           ui.NewStyles(),
	}
	startRepl(cfg)
}
