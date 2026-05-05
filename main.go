package main

import (
	"time"

	"github.com/tonyserranodev/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	pokedex := pokeapi.NewPokedex()
	cfg := &config{
		pokeAPIClient:    pokeClient,
		pokedex:          pokedex,
		VisitedLocations: make(map[string]struct{}),
	}
	startRepl(cfg)
}
