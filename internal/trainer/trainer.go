// Package trainer manages the players state including their pokedex, party, pc
// and bag.
package trainer

import "github.com/tonyserranodev/pokedexcli/internal/pokeapi"

type Pokedex struct {
	Caught map[string]pokeapi.Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{
		Caught: make(map[string]pokeapi.Pokemon),
	}
}

func NewParty() []pokeapi.Pokemon {
	return []pokeapi.Pokemon{}
}

func NewPC() map[string]pokeapi.Pokemon {
	return map[string]pokeapi.Pokemon{}
}
